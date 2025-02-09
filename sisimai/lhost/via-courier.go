// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ______                 _           
// | | |__   ___  ___| |_   / / ___|___  _   _ _ __(_) ___ _ __ 
// | | '_ \ / _ \/ __| __| / / |   / _ \| | | | '__| |/ _ \ '__|
// | | | | | (_) \__ \ |_ / /| |__| (_) | |_| | |  | |  __/ |   
// |_|_| |_|\___/|___/\__/_/  \____\___/ \__,_|_|  |_|\___|_|   
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from Courier MTA: https://www.courier-mta.org/
	InquireFor["Courier"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		proceedsto := true; for {
			// Subject: NOTICE: mail delivery status.
			// Message-ID: <courier.4D025E3A.00001792@5jo.example.org>
			if strings.Contains(bf.Headers["from"][0],    "Courier mail server at ")       { break }
			if strings.Contains(bf.Headers["subject"][0], "NOTICE: mail delivery status.") { break }
			if strings.Contains(bf.Headers["subject"][0], "WARNING: delayed mail.")        { break }
			if strings.HasPrefix(bf.Headers["message-id"][0], "<courier.")                 { break }
			proceedsto = false; break
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"}
		startingof := map[string][]string{
			// https://www.courier-mta.org/courierdsn.html
			// courier/module.dsn/dsn*.txt
			"message": []string{"DELAYS IN DELIVERING YOUR MESSAGE", "UNDELIVERABLE MAIL"},
		}
		messagesof := map[string][]string{
			// courier/module.esmtp/esmtpclient.c:526| hard_error(del, ctf, "No such domain.");
			"hostunknown":  []string{"No such domain."},
			// courier/module.esmtp/esmtpclient.c:531| hard_error(del, ctf,
			// courier/module.esmtp/esmtpclient.c:532|  "This domain's DNS violates RFC 1035.");
			"systemerror":  []string{"This domain's DNS violates RFC 1035."},
			// courier/module.esmtp/esmtpclient.c:535| soft_error(del, ctf, "DNS lookup failed.");
			"networkerror": []string{"DNS lookup failed."},
		}

		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		readslices := []string{""}        // Copy each line for later reference
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		thecommand := ""                  // An SMTP command name begins with the string ">>>"
		v          := &(dscontents[len(dscontents) - 1])

		for j, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			readslices = append(readslices, e) // Save the current line for the next loop

			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if sisimoji.ContainsAny(e, startingof["message"]) { readcursor |= indicators["deliverystatus"] }
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
					v.Update(v.AsRFC1894(o[0]), o[2]); if f != 1 { continue }

					// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
					permessage[z] = o[2]
					if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }
				}
			} else {
				// The line does not begin with a DSN field defined in RFC3464
				//
				// This is a delivery status notification from marutamachi.example.org,
				// running the Courier mail server, version 0.65.2.
				//
				// The original message was received on Sat, 11 Dec 2010 12:19:57 +0900
				// from [127.0.0.1] (c10920.example.com [192.0.2.20])
				//
				// ---------------------------------------------------------------------------
				//
				//                           UNDELIVERABLE MAIL
				//
				// Your message to the following recipients cannot be delivered:
				//
				// <kijitora@example.co.jp>:
				//    mx.example.co.jp [74.207.247.95]:
				// >>> RCPT TO:<kijitora@example.co.jp>
				// <<< 550 5.1.1 <kijitora@example.co.jp>... User Unknown
				//
				// ---------------------------------------------------------------------------
				if strings.HasPrefix(e, ">>> ") {
					// >>> DATA
					thecommand = command.Find(e)

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
				if len(v.Select(z))    > 0 { continue }
				if len(permessage[z]) == 0 { continue }
				e.Update(z, permessage[z])
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			FINDREASON: for r := range messagesof {
				// The key name is a bounce reason name
				for _, f := range messagesof[r] {
					// Try to find an error message including lower-cased string listed in messagesof
					if strings.Contains(e.Diagnosis, f) == false { continue }
					e.Reason = r; break FINDREASON
				}
			}
			e.Command = thecommand
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

