// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ___   _       _            
// | | |__   ___  ___| |_   / / \ | | ___ | |_ ___  ___ 
// | | '_ \ / _ \/ __| __| / /|  \| |/ _ \| __/ _ \/ __|
// | | | | | (_) \__ \ |_ / / | |\  | (_) | ||  __/\__ \
// |_|_| |_|\___/|___/\__/_/  |_| \_|\___/ \__\___||___/
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import sisiaddr "libsisimai.org/sisimai/address"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from HCL Notes (Formerly IBM Notes(Formerly Lotus Notes))
	InquireFor["Notes"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }
		if strings.HasPrefix(bf.Headers["subject"][0], "Undeliverable message") == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"------- Returned Message --------"}
		startingof := map[string][]string{"message": []string{"------- Failure Reasons "} }
		messagesof := map[string][]string{
			"networkerror": []string{"Message has exceeded maximum hop count"},
			"userunknown":  []string{
				"User not listed in public Name & Address Book",
				"ディレクトリのリストにありません",
			},
		}
		dscontents := []sis.DeliveryMatter{{}}
		notdecoded := []sis.NotDecoded{}
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
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// ------- Failure Reasons  --------
			//
			// User not listed in public Name & Address Book
			// kijitora@notes.example.jp
			//
			// ------- Returned Message --------
			if strings.Contains(e, "@") && strings.Contains(e, " ") == false {
				// kijitora@notes.example.jp
				if rfc5322.IsEmailAddress(e) == false { continue }
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = e
				recipients += 1

			} else {
				// Get error messages
				if strings.HasPrefix(e, "-") == false { v.Diagnosis += " " + e }
			}
		}

		for recipients == 0 {
			// Pick an email address from "To:" header of the original message
			p0 := strings.Index(emailparts[1], "\n\n");  if p0 < 0 { p0 = len(emailparts[1]) }
			p1 := strings.Index(emailparts[1], "\nTo:"); if p1 < 0 || p1 > p0 { break }
			p2 := strings.Index(emailparts[1][p1 + 4:], "\n")
			cv := sisiaddr.S3S4(emailparts[1][p1 + 4:p1 + p2 + 4])
			dscontents[0].Recipient = cv; recipients++; break
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			FINDREASON: for r := range messagesof {
				// The key name is a bounce reason name
				for _, f := range messagesof[r] {
					// Try to find an error message including lower-cased string listed in messagesof
					if strings.Contains(e.Diagnosis, f) == false { continue }
					e.Reason = r; break FINDREASON
				}
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1], Errors: notdecoded }
	}
}

