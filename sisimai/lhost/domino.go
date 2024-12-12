// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ______                  _             
// | | |__   ___  ___| |_   / /  _ \  ___  _ __ ___ (_)_ __   ___  
// | | '_ \ / _ \/ __| __| / /| | | |/ _ \| '_ ` _ \| | '_ \ / _ \ 
// | | | | | (_) \__ \ |_ / / | |_| | (_) | | | | | | | | | | (_) |
// |_|_| |_|\___/|___/\__/_/  |____/ \___/|_| |_| |_|_|_| |_|\___/ 
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
		if len(bf.Head)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Body)            == 0 { return sis.RisingUnderway{} }

		proceedsto := false; for {
			if strings.HasPrefix(bf.Head["subject"][0], "DELIVERY FAILURE:") { proceedsto = true }
			if strings.HasPrefix(bf.Head["subject"][0], "DELIVERY_FAILURE:") { proceedsto = true }
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

		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
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
					f := rfc1894.Match(e); if f > 0 {
						// "e" matched with any field defined in RFC3464
						o := rfc1894.Field(e); if len(o) < 1 { continue }
						z := fieldtable[o[0]]; if len(z) < 1 { continue }

						if o[3] == "code" {
							// Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
							if v.Spec      == "" { v.Spec = o[1] }
							if v.Diagnosis == "" { v.Diagnosis = o[2] }

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
		}

		// Set "subjecttxt" as a Subject if there is no original message in the bounce mail.
		if strings.Contains(emailparts[1], "\nSubject:") == false { emailparts[1] += "Subject: " + subjecttxt + "\n" }

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

