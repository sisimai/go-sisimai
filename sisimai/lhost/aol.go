// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//       _               _      ___         _ 
//  _ __| |__   ___  ___| |_   / / \   ___ | |
// | '__| '_ \ / _ \/ __| __| / / _ \ / _ \| |
// | |  | | | | (_) \__ \ |_ / / ___ \ (_) | |
// |_|  |_| |_|\___/|___/\__/_/_/   \_\___/|_|
import "fmt"
import "slices"
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from
	InquireFor["Aol"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head)             == 0 { return sis.RisingUnderway{} }
		if len(bf.Body)             == 0 { return sis.RisingUnderway{} }
		if len(bf.Head["x-aol-ip"]) == 0 { return sis.RisingUnderway{} }

		// X-AOL-IP: 192.0.2.135
		// X-AOL-VSS-INFO: 5600.1067/98281
		// X-AOL-VSS-CODE: clean
		// x-aol-sid: 3039ac1afc14546fb98a0945
		// X-AOL-SCOLL-EIL: 1
		// x-aol-global-disposition: G
		// x-aol-sid: 3039ac1afd4d546fb97d75c6
		// X-BounceIO-Id: 9D38DE46-21BC-4309-83E1-5F0D788EFF1F.1_0
		// X-Outbound-Mail-Relay-Queue-ID: 07391702BF4DC
		// X-Outbound-Mail-Relay-Sender: rfc822; shironeko@aol.example.jp
		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{"message": []string{"Content-Type: message/delivery-status"}}
		messagesof := map[string][]string{
			"hostunknown": []string{"Host or domain name not found"},
			"notaccept":   []string{"type=MX: Malformed or unexpected name server reply"},
		}

		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		readslices := []string{""}        // Copy each line for later reference
		recipients := uint8(0)            // The number of 'Final-Recipient' header
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
					// # Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
					v.Spec = o[1]
					v.Diagnosis = o[2]

				} else {
					// Other DSN fields defined in RFC3464
					v.Set(o[0], o[2]); if f != 1 { continue }

					// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
					permessage[z] = o[2]
					if slices.Contains(keystrings, z) == false { keystrings = append(keystrings, z) }
				}
			} else {
				// Continued line of the value of Diagnostic-Code field
				if strings.HasPrefix(readslices[j], "Diagnostic-Code:") == false { continue }
				if strings.HasPrefix(e, " ")                            == false { continue }
				v.Diagnosis += fmt.Sprintf(" %s", sisimoji.Sweep(e))
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
			e.Diagnosis = sisimoji.Sweep(strings.ReplaceAll(e.Diagnosis, "\n", " "))

			FINDREASON: for r := range messagesof {
				// The key name is a bounce reason name
				for _, f := range messagesof[r] {
					// Try to find an error message including lower-cased string listed in messagesof
					if strings.Contains(e.Diagnosis, f) == false { continue }
					e.Reason = r; break FINDREASON
				}
			}
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

