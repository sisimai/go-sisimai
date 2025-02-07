// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ______                  _             
// | | |__   ___  ___| |_   / /  _ \  ___  _ __ ___ (_)_ __   ___  
// | | '_ \ / _ \/ __| __| / /| | | |/ _ \| '_ ` _ \| | '_ \ / _ \ 
// | | | | | (_) \__ \ |_ / / | |_| | (_) | | | | | | | | | | (_) |
// |_|_| |_|\___/|___/\__/_/  |____/ \___/|_| |_| |_|_|_| |_|\___/ 
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/status"
import sisiaddr "sisimai/address"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from HCL Domino
	InquireFor["Domino"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Headers) == 0 { return sis.RisingUnderway{} }
		if len(bf.Payload) == 0 { return sis.RisingUnderway{} }

		proceedsto := false; for {
			if strings.HasPrefix(bf.Headers["subject"][0], "DELIVERY FAILURE:") { proceedsto = true }
			if strings.HasPrefix(bf.Headers["subject"][0], "DELIVERY_FAILURE:") { proceedsto = true }
			break
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{"message": []string{"Your message"}}
		messagesof := map[string][]string{
			"filtered":    []string{"Cannot route mail to user"},
			"systemerror": []string{"Several matches found in Domino Directory"},
			"userunknown": []string{
				"not listed in Domino Directory",
				"not listed in public Name & Address Book",
				"no se encuentra en el Directorio de Domino",
				"non répertorié dans l'annuaire Domino",
				"Domino ディレクトリには見つかりません",
			},
		}
		exceptutf8 := map[string][][]string{
			// Try to match with the order of each elements for non utf-8 encoded error message
			// such as ISO-8859-1
			"userunknown": [][]string{
				[]string{"non r", "pertori", "dans l'annuaire Domino"}, // "non répertorié dans l'annuaire Domino",
			},
		}

		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		subjecttxt := ""                  // The value of "Subject:"
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

			// Your message
			//
			//   Subject: Test Bounce
			//
			// was not delivered to:
			//
			//   kijitora@example.net
			//
			// because:
			//
			//   User some.name (kijitora@example.net) not listed in Domino Directory
			//
			if e == "was not delivered to:" {
				// was not delivered to:
				//   kijitora@example.net
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = e
				recipients += 1

			} else if strings.HasPrefix(e, "  ") && strings.Contains(e, "@") && strings.Index(e[3:], " ") < 0 {
				// Continued from the line "was not delivered to:"
				//   kijitora@example.net
				v.Recipient = sisiaddr.S3S4(e[2:])

			} else if e == "because:" {
				// because:
				//   User some.name (kijitora@example.net) not listed in Domino Directory
				v.Diagnosis = e

			} else {
				// Error messages, Subject:, other fields defined in RFC3464
				if v.Diagnosis == "because:" {
					// Continued line of the like starts with "because:"
					v.Diagnosis = e

				} else if strings.HasPrefix(e, "  Subject: ") {
					// Subject: Nyaan
					subjecttxt = e[11:]

				} else {
					// Other filelds defined in RFC3464
					f := rfc1894.Match(e); if f      < 1 { continue }
					o := rfc1894.Field(e); if len(o) < 1 { continue }
					z := fieldtable[o[0]]; if len(z) < 1 { continue }

					if o[3] == "code" {
						// Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
						if v.Spec      == "" { v.Spec = o[1] }
						if v.Diagnosis == "" { v.Diagnosis = o[2] }

					} else {
						// Other DSN fields defined in RFC3464
						v.Update(v.AsRFC1894(o[0]), o[2]); if f != 1 { continue }

						// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
						permessage[z] = o[2]
						if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }
					}
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Set default values stored in "permessage" if each value in "dscontents" is empty.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			e.Recipient = sisiaddr.S3S4(e.Recipient)
			for _, z := range keystrings {
				// Do not set an empty string into each member of DeliveryMatter{}
				if len(v.Select(z))    > 0 { continue }
				if len(permessage[z]) == 0 { continue }
				e.Update(z, permessage[z])
			}

			FINDREASON: for r := range messagesof {
				// The key name is a bounce reason name
				for _, f := range messagesof[r] {
					// Try to find an error message including lower-cased string listed in messagesof
					if strings.Contains(e.Diagnosis, f) == false { continue }
					e.Reason = r

					if e.Status == "" { status.Code(r, false) }
					break FINDREASON
				}
			}

			EXCEPTUTF8: for r := range exceptutf8 {
				// The key name is a bounce reason name
				if e.Reason != "" { break EXCEPTUTF8 }
				for _, f := range exceptutf8[r] {
					// Try to find an error message including lower-cased string listed in messagesof
					if sisimoji.Aligned(e.Diagnosis, f) == false { continue }
					e.Reason = r
					break EXCEPTUTF8
				}
			}
		}

		if emailparts[1] == "" {
			// The original message is empty
			if strings.Contains(emailparts[1], "\nTo:") == false {
				// Set "To:" header into the original message
				emailparts[1] += fmt.Sprintf("To: <%s>\n", dscontents[0].Recipient)
			}
			if strings.Contains(emailparts[1], "\nSubject:") == false {
				// Set "subjecttxt" as a Subject if there is no original message in the bounce mail.
				emailparts[1] += "Subject: " + subjecttxt + "\n"
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

