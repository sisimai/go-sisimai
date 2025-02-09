// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      _______     _           
// | | |__   ___  ___| |_   / /__  /___ | |__   ___  
// | | '_ \ / _ \/ __| __| / /  / // _ \| '_ \ / _ \ 
// | | | | | (_) \__ \ |_ / /  / /| (_) | | | | (_) |
// |_|_| |_|\___/|___/\__/_/  /____\___/|_| |_|\___/ 
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"
import sisiaddr "libsisimai.org/sisimai/address"

func init() {
	// Decode bounce messages from Zoho Mail: https://www.zoho.com/mail/'
	InquireFor["Zoho"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true    { return sis.RisingUnderway{} }
		if len(bf.Headers["x-zohomail"]) == 0 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"\n\nReceived:"}
		startingof := map[string][]string{
			"message": []string{"This message was created automatically by mail delivery"},
		}
		messagesof := map[string][]string{
			"expired": []string{"Host not reachable"},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, true)
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

			// This message was created automatically by mail delivery software.
			// A message that you sent could not be delivered to one or more of its recip=
			// ients. This is a permanent error.=20
			//
			// kijitora@example.co.jp Invalid Address, ERROR_CODE :550, ERROR_CODE :5.1.=
			// 1 <kijitora@example.co.jp>... User Unknown

			// This message was created automatically by mail delivery software.
			// A message that you sent could not be delivered to one or more of its recipients. This is a permanent error.
			//
			// shironeko@example.org Invalid Address, ERROR_CODE :550, ERROR_CODE :Requested action not taken: mailbox unavailable
			if sisimoji.Aligned(e, []string{"@", " ", "ERROR_CODE :"}) || strings.HasPrefix(e, "[Status: ") {
				// kijitora@example.co.jp Invalid Address, ERROR_CODE :550, ERROR_CODE :5.1.=
				// [Status: Error, Address: <kijitora@6kaku.example.co.jp>, ResponseCode 421, , Host not reachable.]
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient  = sisiaddr.S3S4(e)
				v.Diagnosis += " " + e
				recipients  += 1

			} else {
				// Error message which does not include a recipient email address
				if strings.HasPrefix(e, "-----") { continue }
				v.Diagnosis += " " + e
			}
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
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}
