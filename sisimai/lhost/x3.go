// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ____  _______ 
// | | |__   ___  ___| |_   / /\ \/ /___ / 
// | | '_ \ / _ \/ __| __| / /  \  /  |_ \ 
// | | | | | (_) \__ \ |_ / /   /  \ ___) |
// |_|_| |_|\___/|___/\__/_/   /_/\_\____/ 
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Unknown MTA #3
	InquireFor["X3"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }
		if strings.HasPrefix(bf.Head["from"][0],    "Mail Delivery System")         == false { return sis.RisingUnderway{} }
		if strings.HasPrefix(bf.Head["subject"][0], "Delivery status notification") == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{
			"message": []string{"      This is an automatically generated Delivery Status Notification."},
		}
		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
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

			// ============================================================================
			//      This is an automatically generated Delivery Status Notification.
			//
			// Delivery to the following recipients failed permanently:
			//
			//   * kijitora@example.com
			//
			//
			// ============================================================================
			//                             Technical details:
			//
			// SMTP:RCPT host 192.0.2.8: 553 5.3.0 <kijitora@example.com>... No such user here
			//
			//
			// ============================================================================
			if sisimoji.Aligned(e, []string{"  * ", "@", "."}) {
				//   * kijitora@example.com
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e[strings.Index(e, "  * ") + 3:])
				recipients += 1

			} else {
				// Error message lines
				if strings.HasPrefix(e, "SMTP:") {
					// SMTP:RCPT host 192.0.2.8: 553 5.3.0 <kijitora@example.com>... No such user here
					v.Command    = command.Find(e[4:10])
					v.Spec       = "SMTP"
					v.Diagnosis += " " + e

				} else if strings.HasPrefix(e, "Routing: ") {
					// Routing: Could not find a gateway for kijitora@example.co.jp
					v.Diagnosis += " " + e

				} else if f := rfc1894.Match(e); f > 0 {
					// "e" matched with any field defined in RFC3464
					o := rfc1894.Field(e); if len(o) == 0 { continue }
					z := fieldtable[o[0]]

					if o[3] == "addr" {
						// Final-Recipient: rfc822; kijitora@example.jp
						// X-Actual-Recipient: rfc822; kijitora@example.co.jp
						if o[0] == "final-recipient" {
							// Final-Recipient: rfc822; kijitora@example.jp
							if v.Recipient == o[2] { continue }
							if len(v.Recipient) > 0 {
								// There are multiple recipient addresses in the message body.
								dscontents = append(dscontents, sis.DeliveryMatter{})
								v = &(dscontents[len(dscontents) - 1])
							}
							v.Recipient = o[2]
							recipients += 1

						} else {
							// X-Actual-Recipient: rfc822; kijitora@example.co.jp
							v.Alias = o[2]
						}
					} else if o[3] == "code" {
						// # Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
						v.Spec = o[1]
						v.Diagnosis += " " + o[2]

					} else {
						// Other DSN fields defined in RFC3464
						v.Update(o[0], o[2]); if f != 1 { continue }

						// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
						permessage[z] = o[2]
						if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }
					}
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			for _, z := range keystrings {
				// Do not set an empty string into each member of DeliveryMatter{}
				if len(v.Select(z))    > 0 { continue }
				if len(permessage[z]) == 0 { continue }
				e.Update(z, permessage[z])
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
		}
		if emailparts[1] == "" {
			// The original message part is empty, set "To:", "Date:" headers
			emailparts[1] += fmt.Sprintf("To: %s\n", dscontents[0].Recipient)
			emailparts[1] += fmt.Sprintf("Date: %s\n", dscontents[0].Date)
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

