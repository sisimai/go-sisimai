// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      ____     __        _                
// | | |__   ___  ___| |_   / /\ \   / /__ _ __(_)_______  _ __  
// | | '_ \ / _ \/ __| __| / /  \ \ / / _ \ '__| |_  / _ \| '_ \ 
// | | | | | (_) \__ \ |_ / /    \ V /  __/ |  | |/ / (_) | | | |
// |_|_| |_|\___/|___/\__/_/      \_/ \___|_|  |_/___\___/|_| |_|

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from Verizon: https://www.verizon.com/
	InquireFor["Verizon"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		proceedsto := uint8(0)
		if strings.Contains(bf.Headers["from"][0], "post_master@vtext.com")              { proceedsto = 1 }
		if sisimoji.Aligned(bf.Headers["from"][0], []string{"sysadmin@", ".vzwpix.com"}) { proceedsto = 1 }
		if proceedsto == 0 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Original Message:", "Message details:"}
		nooriginal := false
		startingof := map[string][]string{"message": []string{"Error: "}}
		messagesof := map[string][]string{
			"userunknown": []string{"550 - Requested action not taken: no such user here", "No valid recipients"},
		}

		if strings.Contains(bf.Payload, boundaries[1]) {
			// Message details:
			//   Subject: Test message
			//   Sent date: Wed Apr 29 23:34:45 GMT 2013
			//   MAIL FROM: nyaaaaaaan@neko.example.org
			//   RCPT TO: may-be-straycat-nyaaaaaan@vtext.com
			//   From: sironeko-nekochan-nyaaaaaaaan-nyan@sabineko.example.com
			// Convert strings above to RFC822 email headers
			nooriginal = true
			bf.Payload = strings.Replace(bf.Payload, "Sent date: ", "Date: ", 1)
			bf.Payload = strings.Replace(bf.Payload, "MAIL FROM: ", "Return-Path: ", 1)
			bf.Payload = strings.Replace(bf.Payload, "RCPT TO: ",   "To: ", 1)
		}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
			}
			if readcursor & indicators["deliverystatus"] == 0 || e == "" { continue }

			v.Diagnosis += e
		}

		if nooriginal { emailparts[1] = strings.ReplaceAll(emailparts[1], "\n  ", "\n") }
		p1 := strings.Index(emailparts[1], "\nTo: "); if p1 < 1 { return sis.RisingUnderway{} }
		p2 := sisimoji.IndexOnTheWay(emailparts[1], "\n", p1 + 5)
		dscontents[0].Recipient = emailparts[1][p1 + 5:p2]

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			FINDREASON: for r := range messagesof {
				// The key name is a bounce reason name
				for _, f := range messagesof[r] {
					// Try to find an error message including lower-cased string listed in messagesof
					if strings.Contains(e.Diagnosis, f) { e.Reason = r; break FINDREASON }
				}
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

