// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//       _               _      ______  _        __             _   
//  _ __| |__   ___  ___| |_   / / __ )(_) __ _ / _| ___   ___ | |_ 
// | '__| '_ \ / _ \/ __| __| / /|  _ \| |/ _` | |_ / _ \ / _ \| __|
// | |  | | | | (_) \__ \ |_ / / | |_) | | (_| |  _| (_) | (_) | |_ 
// |_|  |_| |_|\___/|___/\__/_/  |____/|_|\__, |_|  \___/ \___/ \__|
//                                        |___/                     
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from Bigfoot: http://www.bigfoot.com
	InquireFor["Bigfoot"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		proceedsto := false
		if strings.Contains(bf.Head["from"][0], "@bigfoot.com>") {
			// From: Mail Delivery Subsystem <MAILER-DAEMON@bigfoot.com>
			proceedsto = true

		} else {
			for _, e := range bf.Head["received"] {
				// Received: by mx.bigfoot.com with SMTP id ...
				if strings.Contains(e, ".bigfoot.com ") == false { continue }
				proceedsto = true; break
			}
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/partial"}
		startingof := map[string][]string{"message": []string{"   ----- Transcript of session follows -----"}}

		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		readslices := []string{""}        // Copy each line for later reference
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		thecommand := ""                  // An SMTP command name begins with the string ">>>"
		esmtpreply := ""                  // Reply messages from the remote server on an SMTP session
		v          := &(dscontents[len(dscontents) - 1])

		for j, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			readslices = append(readslices, e) // Save the current line for the next loop

			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			f := rfc1894.Match(e); if f > 0 {
				// "e" matched with any field defined in RFC3464
				o := rfc1894.Field(e); if len(o) == 0 { continue }
				z := fieldtable[o[0]]
				v  = &(dscontents[len(dscontents) - 1])

				if o[3] == "addr" {
					// Final-Recipient: rfc822; kijitora@example.jp
					// X-Actual-Recipient: rfc822; kijitora@example.co.jp
					if o[0] == "final-recipient" {
						// Final-Recipient: rfc822; kijitora@example.jp
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
					// Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
					v.Spec = o[1]
					v.Diagnosis = o[2]

				} else {
					// Other DSN fields defined in RFC3464
					v.Set(o[0], o[2]); if f != 1 { continue }

					// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
					permessage[z] = o[2]
					if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }
				}
			} else {
				// The line does not begin with a DSN field defined in RFC3464
				if e[0:1] != " " {
					//    ----- Transcript of session follows -----
					// >>> RCPT TO:<destinaion@example.net>
					// <<< 553 Invalid recipient destinaion@example.net (Mode: normal)
					if strings.HasPrefix(e, ">>> ") {
						// >>> DATA
						thecommand = command.Find(e)

					} else if strings.HasPrefix(e, "<<< ") {
						// <<< Response from the MTA
						esmtpreply = e[4:]
					}
				} else {
					// Continued line of the value of Diagnostic-Code field
					if strings.HasPrefix(readslices[j], "Diagnostic-Code:") == false { continue }
					if strings.HasPrefix(e, " ")                            == false { continue }
					v.Diagnosis += fmt.Sprintf(" %s", sisimoji.Sweep(e))
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Set default values stored in "permessage" if each value in "dscontents" is empty.
			e := &(dscontents[j])
			for _, z := range keystrings {
				// Do not set an empty string into each member of DeliveryMatter{}
				if len(v.Get(z))       > 0 { continue }
				if len(permessage[z]) == 0 { continue }
				e.Set(z, permessage[z])
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			e.Command   = thecommand; if e.Command == "" && esmtpreply != "" { e.Command = "EHLO" }
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

