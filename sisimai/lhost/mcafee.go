// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ____  __         _     __           
// | | |__   ___  ___| |_   / /  \/  | ___   / \   / _| ___  ___ 
// | | '_ \ / _ \/ __| __| / /| |\/| |/ __| / _ \ | |_ / _ \/ _ \
// | | | | | (_) \__ \ |_ / / | |  | | (__ / ___ \|  _|  __/  __/
// |_|_| |_|\___/|___/\__/_/  |_|  |_|\___/_/   \_\_|  \___|\___|
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from McAfee Email Appliance'
	InquireFor["McAfee"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head)                 == 0 { return sis.RisingUnderway{} }
		if len(bf.Body)                 == 0 { return sis.RisingUnderway{} }
		if len(bf.Head["x-nai-header"]) == 0 { return sis.RisingUnderway{} }

		// X-NAI-Header: Modified by McAfee Email and Web Security Virtual Appliance
		if strings.Contains(bf.Head["x-nai-header"][0], "Modified by McAfee") == false { return sis.RisingUnderway{} }
		if bf.Head["subject"][0] != "Delivery Status"                                  { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{"message": []string{"--- The following addresses had delivery problems ---"}}
		messagesof := map[string][]string{"userunknown": []string{" User not exist", " unknown.", "550 Unknown user ", "No such user"}}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		readslices := []string{""}        // Copy each line for later reference
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		issuedcode := ""                  // Alternative error message
		v          := &(dscontents[len(dscontents) - 1])

		for j, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			readslices = append(readslices, e) // Save the current line for the next loop

			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.Contains(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// Content-Type: text/plain; name="deliveryproblems.txt"
			//
			//    --- The following addresses had delivery problems ---
			//
			// <user@example.com>   (User unknown user@example.com)
			//
			// --------------Boundary-00=_00000000000000000000
			// Content-Type: message/delivery-status; name="deliverystatus.txt"
			if sisimoji.Aligned(e, []string{"<", "@", ">", "(", ")"}) {
				// <kijitora@example.co.jp>   (Unknown user kijitora@example.co.jp)
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e[strings.Index(e, "<"):strings.Index(e, ">") + 1])
				issuedcode  = e[strings.Index(e, ")") + 1:]
				recipients += 1

			} else if rfc1894.Match(e) > 0 {
				// The line includes any field defined in RFC3464
				o := rfc1894.Field(e)
				if len(o) == 0 {
					// Fallback code for empty value or invalid formatted value
					if strings.HasPrefix(e, "Original-Recipient: ") {
						// - Original-Recipient: <kijitora@example.co.jp>
						v.Alias = sisiaddr.S3S4(e[strings.Index(e, ":") + 1:])
					}
					continue
				}
				// Other DSN fields defined in RFC3464
				v.Set(o[0], o[2])

			} else {
				// Continued line of the value of Diagnostic-Code field
				if strings.HasPrefix(readslices[j], "Diagnostic-Code:") == false { continue }
				if strings.HasPrefix(e, " ")                            == false { continue }
				v.Diagnosis += " " + sisimoji.Sweep(e)
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Set default values stored in "permessage" if each value in "dscontents" is empty.
			e := &(dscontents[j])
			f := e.Diagnosis; if f == "" { f = issuedcode }
			e.Diagnosis = sisimoji.Sweep(f)

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

