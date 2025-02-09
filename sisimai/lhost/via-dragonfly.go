// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ______                              _____ _       
// | | |__   ___  ___| |_   / /  _ \ _ __ __ _  __ _  ___  _ __ |  ___| |_   _ 
// | | '_ \ / _ \/ __| __| / /| | | | '__/ _` |/ _` |/ _ \| '_ \| |_  | | | | |
// | | | | | (_) \__ \ |_ / / | |_| | | | (_| | (_| | (_) | | | |  _| | | |_| |
// |_|_| |_|\___/|___/\__/_/  |____/|_|  \__,_|\__, |\___/|_| |_|_|   |_|\__, |
//                                             |___/                     |___/ 
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisiaddr "sisimai/address"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from DragonFly: https://www.dragonflybsd.org/handbook/mta/
	InquireFor["DragonFly"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		// From: MAILER-DAEMON <>
		// To: kijitora@df.example.jp
		// Subject: Mail delivery failed
		if strings.Contains(bf.Headers["subject"][0], "Mail delivery failed") == false { return sis.RisingUnderway{} }
		proceedsto := false; for _, e := range bf.Headers["received"] {
			// Received: from MAILER-DAEMON
			//    id e070f
			//    by df.example.jp (DragonFly Mail Agent v0.13);
			//    Sun, 16 Jun 2024 18:15:07 +0900
			if strings.Contains(e, " (DragonFly Mail Agent") { proceedsto = true; break }
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Original message follows.", "Message headers follow"}
		startingof := map[string][]string{
			// https://github.com/corecode/dma/blob/ffad280aa40c242aa9a2cb9ca5b1b6e8efedd17e/mail.c#L84
			"message": []string{"This is the DragonFly Mail Agent "},
		}
		messagesof := map[string][]string{
			"expired": []string{
				// https://github.com/corecode/dma/blob/master/dma.c#L370C1-L374C19
				// dma.c:370| if (gettimeofday(&now, NULL) == 0 &&
				// dma.c:371|     (now.tv_sec - st.st_mtim.tv_sec > MAX_TIMEOUT)) {
				// dma.c:372|     snprintf(errmsg, sizeof(errmsg),
				// dma.c:373|          "Could not deliver for the last %d seconds. Giving up.",
				// dma.c:374|          MAX_TIMEOUT);
				// dma.c:375|     goto bounce;
				// dma.c:376| }
				"Could not deliver for the last ",
			},
			"hostunknown": []string{
				// net.c:663| snprintf(errmsg, sizeof(errmsg), "DNS lookup failure: host %s not found", host);
				"DNS lookup failure: host ",
			},
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
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// This is the DragonFly Mail Agent v0.13 at df.example.jp.
			//
			// There was an error delivering your mail to <kijitora@example.com>.
			//
			// email.example.jp [192.0.2.25] did not like our RCPT TO:
			// 552 5.2.2 <kijitora@example.com>: Recipient address rejected: Mailbox full
			//
			// Original message follows.
			if strings.HasPrefix(e, "There was an error delivering your mail to <") {
				// email.example.jp [192.0.2.25] did not like our RCPT TO:
				// 552 5.2.2 <kijitora@example.com>: Recipient address rejected: Mailbox full
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e[strings.Index(e, "<"):])
				recipients += 1

			} else {
				// Pick the error message
				v.Diagnosis += e

				// Pick the remote hostname, and the SMTP command
				// net.c:500| snprintf(errmsg, sizeof(errmsg), "%s [%s] did not like our %s:\n%s",
				if strings.Contains(e, " did not like our ") == false { continue }
				if len(v.Rhost) > 0                                   { continue }

				p := strings.SplitN(e, " ", 3)
				if strings.Index(p[0], ".") > 1 { v.Rhost = p[0] } else { v.Rhost = p[1] }
				v.Command = command.Find(e)
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

