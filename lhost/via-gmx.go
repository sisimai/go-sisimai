// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      ______ __  ____  __
// | | |__   ___  ___| |_   / / ___|  \/  \ \/ /
// | | '_ \ / _ \/ __| __| / / |  _| |\/| |\  / 
// | | | | | (_) \__ \ |_ / /| |_| | |  | |/  \ 
// |_|_| |_|\___/|___/\__/_/  \____|_|  |_/_/\_\

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/smtp/command"
import sisiaddr "libsisimai.org/sisimai/address"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from GMX: https://gmx.net/
	InquireFor["GMX"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		// Envelope-To: <kijitora@mail.example.com>
		// X-GMX-Antispam: 0 (Mail was not recognized as spam); Detail=V3;
		// X-GMX-Antivirus: 0 (no virus found)
		// X-UI-Out-Filterresults: unknown:0;
		if len(bf.Headers["x-gmx-antispam"]) == 0 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"--- The header of the original message is following. ---"}
		startingof := map[string][]string{
			"message": []string{"This message was created automatically by mail delivery software"},
		}
		messagesof := map[string][]string{
			"expired": []string{"delivery retry timeout exceeded"},
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
			if readcursor & indicators["deliverystatus"] == 0 || e == "" { continue }

			// This message was created automatically by mail delivery software.
			//
			// A message that you sent could not be delivered to one or more of
			// its recipients. This is a permanent error. The following address
			// failed:
			//
			// "shironeko@example.jp":
			// SMTP error from remote server after RCPT command:
			// host: mx.example.jp
			// 5.1.1 <shironeko@example.jp>... User Unknown
			if (strings.IndexByte(e, '@') > 1 && strings.HasPrefix(e, `"`)) || strings.HasPrefix(e, "<") {
				// "shironeko@example.jp":
				// ---- OR ----
				// <kijitora@6jo.example.co.jp>
				//
				// Reason:
				// delivery retry timeout exceeded
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(strings.Trim(e, ":"))
				recipients += 1

			} else if strings.HasPrefix(e, "SMTP error ") {
				// SMTP error from remote server after RCPT command:
				v.Command = command.Find(e)

			} else if strings.HasPrefix(e, "host: ") {
				// host: mx.example.jp
				v.Rhost = e[6:]

			} else {
				// Get error messages
				if e != "" { v.Diagnosis += e + " " }
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(strings.ReplaceAll(e.Diagnosis, "\n", " "))

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

