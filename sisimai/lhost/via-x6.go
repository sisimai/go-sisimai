// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ____  ____   
// | | |__   ___  ___| |_   / /\ \/ / /_  
// | | '_ \ / _ \/ __| __| / /  \  / '_ \ 
// | | | | | (_) \__ \ |_ / /   /  \ (_) |
// |_|_| |_|\___/|___/\__/_/   /_/\_\___/ 
import "strings"
import "sisimai/sis"
import "sisimai/rfc1123"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Unknown MTA #6
	InquireFor["X6"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Headers) == 0 { return sis.RisingUnderway{} }
		if len(bf.Payload) == 0 { return sis.RisingUnderway{} }
		if strings.HasPrefix(bf.Headers["subject"][0], "There was an error sending your mail") == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"The attachment contains the original mail headers"}
		startingof := map[string][]string{"message": []string{"We had trouble delivering your message. Full details follow:"}}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// We had trouble delivering your message. Full details follow:
			//
			// Subject: 'Nyaan'
			// Date: 'Thu, 29 Apr 2012 23:34:45 +0000'
			//
			// 1 error(s):
			//
			// SMTP Server <mta2.example.jp> rejected recipient <kijitora@example.jp> 
			//   (Error following RCPT command). It responded as follows: [550 5.1.1 User unknown]
			p1 := strings.Index(e, "The following recipients returned permanent errors: ")
			p2 := strings.Index(e, "SMTP Server <")
			p3 := 0
			if p1 == 0 || p2 == 0 {
				// SMTP Server <mta2.example.jp> rejected recipient <kijitora@example.jp>
				// The following recipients returned permanent errors: neko@example.jp.
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				if p1 == 0 { p3 = strings.Index(e, ": ") } else { p3 = strings.LastIndex(e, " <") }
				cv := sisiaddr.S3S4(e[p3:]); if sisiaddr.IsEmailAddress(cv) == false { continue }

				v.Recipient  = cv
				v.Diagnosis += " " + e
				recipients  += 1

			} else if strings.HasPrefix(e, "Date: ") {
				// Date: 'Thu, 29 Apr 2012 23:34:45 +0000'
				v.Date = strings.Trim(e[6:], "'")
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			e.Command   = command.Find(e.Diagnosis)
			e.Rhost     = rfc1123.Find(e.Diagnosis)
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

