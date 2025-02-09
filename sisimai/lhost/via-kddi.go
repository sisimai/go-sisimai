// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ___  ______  ____ ___ 
// | | |__   ___  ___| |_   / / |/ /  _ \|  _ \_ _|
// | | '_ \ / _ \/ __| __| / /| ' /| | | | | | | | 
// | | | | | (_) \__ \ |_ / / | . \| |_| | |_| | | 
// |_|_| |_|\___/|___/\__/_/  |_|\_\____/|____/___|
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from au by KDDI: https://www.au.kddi.com
	InquireFor["KDDI"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		proceedsto := false
		senderlist := []string{"no-reply@.", ".dion.ne.jp"}
		replyslist := []string{"no-reply@app.auone-net.jp"}
		ISKDDI: for {
			replyto := ""; if len(bf.Headers["reply-to"]) > 0 { replyto = bf.Headers["reply-to"][0] }
			if sisimoji.ContainsAny(bf.Headers["from"][0], senderlist) { proceedsto = true; break ISKDDI }
			if sisimoji.ContainsAny(replyto, replyslist)               { proceedsto = true; break ISKDDI }

			for _, e := range bf.Headers["received"] {
				// Received: from ezweb.ne.jp (nx3oBP05-09.ezweb.ne.jp [59.135.39.233])
				if strings.Contains(e, "ezweb.ne.jp (") { proceedsto = true; break ISKDDI }
				if strings.Contains(e, ".au.com (")     { proceedsto = true; break ISKDDI }
			}
			break
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{
			"message": []string{"Your mail sent on:", "Your mail attempted to be delivered on:"},
		}
		messagesof := map[string][]string{
			"mailboxfull": []string{"As their mailbox is full"},
			"norelaying":  []string{"Due to the following SMTP relay error"},
			"hostunknown": []string{"As the remote domain doesnt exist"},
		}

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
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			if strings.Contains(e, " Could not be delivered to: <") {
				// Your mail sent on: Thu, 29 Apr 2010 11:04:47 +0900
				//     Could not be delivered to: <******@**.***.**>
				//     As their mailbox is full.
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				cv := sisiaddr.S3S4(e[strings.Index(e, "<"):])
				if sisiaddr.IsEmailAddress(cv) == false { continue }
				v.Recipient = cv
				recipients += 1

			} else if strings.Contains(e, "Your mail sent on: ") {
				// Your mail sent on: Thu, 29 Apr 2010 11:04:47 +0900
				v.Date = e[19:]

			} else {
				//     As their mailbox is full.
				if strings.HasPrefix(e, " ") { v.Diagnosis += e + " " }
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			e.Command   = command.Find(e.Diagnosis)

			if len(bf.Headers["x-spasign"]) > 0 && bf.Headers["x-spasign"][0] == "NG" {
				// Content-Type: text/plain; ..., X-SPASIGN: NG (spamghetti, au by KDDI)
				// Filtered recipient returns message that include 'X-SPASIGN' header
				e.Reason = "filtered"

			} else {
				// There is no X-SPASIGN: header in the bounce message
				if e.Command == "RCPT" { e.Reason = "userunkonwn"; continue }

				FINDREASON: for r := range messagesof {
					// The key name is a bounce reason name
					for _, f := range messagesof[r] {
						// Try to find an error message including lower-cased string listed in messagesof
						if strings.Contains(e.Diagnosis, f) == false { continue }
						e.Reason = r; break FINDREASON
					}
				}
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

