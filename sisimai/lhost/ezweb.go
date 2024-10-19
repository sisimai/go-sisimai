// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      _______ _____             _     
// | | |__   ___  ___| |_   / / ____|__  /_      _____| |__  
// | | '_ \ / _ \/ __| __| / /|  _|   / /\ \ /\ / / _ \ '_ \ 
// | | | | | (_) \__ \ |_ / / | |___ / /_ \ V  V /  __/ |_) |
// |_|_| |_|\___/|___/\__/_/  |_____/____| \_/\_/ \___|_.__/ 
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from au EZweb: https://www.au.com/mobile/
	InquireFor["EZweb"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		proceedsto := 0; for {
			// Pre-process email headers of NON-STANDARD bounce message au by EZweb, as known as ezweb.ne.jp.
			//   Subject: Mail System Error - Returned Mail
			//   From: <Postmaster@ezweb.ne.jp>
			//   Message-Id: <20110000000000.F000000@lsean.ezweb.ne.jp>
			if strings.Contains(bf.Head["from"][0], "Postmaster@ezweb.ne.jp") { proceedsto++ }
			if strings.Contains(bf.Head["from"][0], "Postmaster@au.com")      { proceedsto++ }
			if bf.Head["subject"][0] == "Mail System Error - Returned Mail"   { proceedsto++ }
			for _, e := range bf.Head["received"] {
				//   Received: from ezweb.ne.jp (wmflb12na02.ezweb.ne.jp [222.15.69.197])
				//   Received: from nmomta.auone-net.jp ([aaa.bbb.ccc.ddd]) by ...
				if strings.Contains(e, "ezweb.ne.jp (EZweb Mail) with") || strings.Contains(e, ".au.com (") {
					proceedsto++
					break
				}
			}
			if strings.Contains(bf.Head["message-id"][0], ".ezweb.ne.jp>") { proceedsto++ }
			if strings.Contains(bf.Head["message-id"][0], ".au.com>")      { proceedsto++ }
			break
		}
		if proceedsto < 2 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"--------------------------------------------------", "Content-Type: message/rfc822"}
		startingof := map[string][]string{
			"message": []string{"The user(s) ", "Your message ", "Each of the following", "<"},
		}
		messagesof := map[string][]string{
			//"notaccept": []string{"The following recipients did not receive this message:"},
			"mailboxfull": []string{"The user(s) account is temporarily over quota"},
			"suspend":     []string{
				// http://www.naruhodo-au.kddi.com/qa3429203.html
				// The recipient may be unpaid user...?
				"The user(s) account is disabled.",
				"The user(s) account is temporarily limited.",
			},
			"expired": []string{
				// Your message was not delivered within 0 days and 1 hours.
				// Remote host is not responding.
				"Your message was not delivered within ",
			},
			"onhold": []string{"Each of the following recipients was rejected by a remote mail server"},
		}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		substrings := []string{}          // All the values of "messagesof"
		alternates := ""                  // Other error message strings
		v          := &(dscontents[len(dscontents) - 1])

		// Add all the values of messagesof into substrings
		for e := range messagesof { for _, f := range messagesof[e] { substrings = append(substrings, f) } }

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				for _, f := range startingof["message"] {
					if strings.Contains(e, f) { readcursor |= indicators["deliverystatus"]; break }
				}
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// The user(s) account is disabled.
			//
			// <***@ezweb.ne.jp>: 550 user unknown (in reply to RCPT TO command)
			//
			//  -- OR --
			// Each of the following recipients was rejected by a remote
			// mail server.
			//
			//    Recipient: <******@ezweb.ne.jp>
			//    >>> RCPT TO:<******@ezweb.ne.jp>
			//    <<< 550 <******@ezweb.ne.jp>: User unknown
			if sisimoji.Aligned(e, []string{"<", "@", ">"}) && 
			   (strings.Index(e, "Recipient: <") > 1 || strings.HasPrefix(e, "<")) {
				// Recipient: <******@ezweb.ne.jp> OR <***@ezweb.ne.jp>: 550 user unknown ...
				p1 := strings.Index(e, "<")
				p2 := strings.Index(e, ">")

				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e[p1 + 1:p2])
				recipients += 1

			} else {
				f := rfc1894.Match(e); if f > 0 {
					// "e" matched with any field defined in RFC3464
					o := rfc1894.Field(e); if len(o) == 0 { continue }
					v.Set(o[0], o[2])

				} else {
					// The line does not begin with a DSN field defined in RFC3464
					if sisimoji.Is8Bit(&e) == true { continue }
					if strings.Contains(e, " >>> ") {
						//    >>> RCPT TO:<******@ezweb.ne.jp>
						v.Command = command.Find(e)

					} else {
						// Check the error message
						isincluded := false
						for _, r := range substrings {
							// Try to find that the line contains any error message text
							if strings.Contains(e, r) == false { continue }
							v.Diagnosis += " " + e
							isincluded   = true
						}
						if isincluded == false { alternates += " " + e }
					}
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Check each value of DeliveryMatter
			e := &(dscontents[j])

			if alternates != "" {
				// Copy alternative error message to e.Diagnosis
				if e.Diagnosis == "" {
					// There is no error message in v.Diagnosis
					e.Diagnosis = alternates

				} else if strings.HasPrefix(e.Diagnosis, "-") || strings.HasSuffix(e.Diagnosis, "__") {
					// Override the value of v.Diagsnosis
					e.Diagnosis = alternates
				}
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			if len(bf.Head["x-spasign"]) > 0 && bf.Head["x-spasign"][0] == "NG" {
				// Content-Type: text/plain; ..., X-SPASIGN: NG (spamghetti, au by EZweb)
				// Filtered recipient returns message that include 'X-SPASIGN' header
				e.Reason = "filtered"

			} else {
				// There is no X-SPASIGN header or the value of the header is not "NG"
				if e.Command == "RCPT" {
					// Set "userunknown" when the remote server rejected after RCPT command.
					e.Reason = "userunknown"

				} else {
					// The last SMTP command is not "RCPT"
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
			if e.Reason == ""                                { continue }
			if strings.Contains(e.Recipient, "@ezweb.ne.jp") { continue }
			if strings.Contains(e.Recipient, "@au.com")      { continue }
			e.Reason = "userunknown"
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

