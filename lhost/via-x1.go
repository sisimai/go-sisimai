// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      ____  ___ 
// | | |__   ___  ___| |_   / /\ \/ / |
// | | '_ \ / _ \/ __| __| / /  \  /| |
// | | | | | (_) \__ \ |_ / /   /  \| |
// |_|_| |_|\___/|___/\__/_/   /_/\_\_|

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"
import sisiaddr "libsisimai.org/sisimai/address"

func init() {
	// Decode bounce messages from Unknown MTA #1
	InquireFor["X1"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true                                             { return sis.RisingUnderway{} }
		if strings.HasPrefix(bf.Headers["subject"][0], "Returned Mail: ")     == false { return sis.RisingUnderway{} }
		if strings.HasPrefix(bf.Headers["from"][0], `"Mail Deliver System" `) == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Received: from "}
		startingof := map[string][]string{"message": []string{"The original message was received at "}}

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
			if readcursor & indicators["deliverystatus"] == 0 || e == "" { continue }

			// The original message was received at Thu, 29 Apr 2010 23:34:45 +0900 (JST)
			// from shironeko@example.jp
			//
			// ---The following addresses had delivery errors---
			//
			// kijitora@example.co.jp [User unknown]
			if sisimoji.Aligned(e, []string{"@", " [", "]"}) {
				// kijitora@example.co.jp [User unknown]
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				p1 := strings.IndexByte(e, ' ')
				cv := sisiaddr.S3S4(e[:p1]); if rfc5322.IsEmailAddress(cv) == false { continue }
				v.Recipient  = cv
				v.Diagnosis += " " + e
				recipients  += 1

			} else {
				// The original message was received at Thu, 29 Apr 2010 23:34:45 +0900 (JST) 
				// from shironeko@example.jp
				if strings.HasPrefix(e, "---") && strings.HasSuffix(e, "---") {
					// ---The following addresses had delivery errors---
					e = strings.Trim(e, "---"); e += ":"
					v.Diagnosis += "."
				}
				v.Diagnosis += " " + e
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Pick the date string from the error message.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			if e.Date == "" {
				// The original message was received at Thu, 29 Apr 2010 23:34:45 +0900 (JST)
				// from shironeko@example.jp
				e.Date = strings.Trim(sisimoji.Select(e.Diagnosis, " at ", "from", 0), " ")
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

