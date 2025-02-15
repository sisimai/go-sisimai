// Copyright (C) 20242-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      ____  __                           _             ____                           
// | | |__   ___  ___| |_   / /  \/  | ___  ___ ___  __ _  __ _(_)_ __   __ _/ ___|  ___ _ ____   _____ _ __ 
// | | '_ \ / _ \/ __| __| / /| |\/| |/ _ \/ __/ __|/ _` |/ _` | | '_ \ / _` \___ \ / _ \ '__\ \ / / _ \ '__|
// | | | | | (_) \__ \ |_ / / | |  | |  __/\__ \__ \ (_| | (_| | | | | | (_| |___) |  __/ |   \ V /  __/ |   
// |_|_| |_|\___/|___/\__/_/  |_|  |_|\___||___/___/\__,_|\__, |_|_| |_|\__, |____/ \___|_|    \_/ \___|_|   
//                                                        |___/         |___/                                

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc791"
import "libsisimai.org/sisimai/rfc1894"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/smtp/reply"
import "libsisimai.org/sisimai/smtp/status"
import sisimoji "libsisimai.org/sisimai/string"
import sisiaddr "libsisimai.org/sisimai/address"

func init() {
	// Decode bounce messages from Oracle Communications Messaging Server
	InquireFor["MessagingServer"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		// @see https://docs.oracle.com/en/industries/communications/messaging-server/index.html
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		proceedsto := false
		if strings.Contains(bf.Headers["content-type"][0], "Boundary_(ID_")       { proceedsto = true }
		if strings.HasPrefix(bf.Headers["subject"][0], "Delivery Notification: ") { proceedsto = true }
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822", "\nReturn-Path: "}
		startingof := map[string][]string{"message": []string{"This report relates to a message you sent with the following header fields:"}}
		messagesof := map[string][]string{"hostunknown": []string{"Illegal host/domain name found"}}

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

			// --Boundary_(ID_0000000000000000000000)
			// Content-type: text/plain; charset=us-ascii
			// Content-language: en-US
			//
			// This report relates to a message you sent with the following header fields:
			//
			//   Message-id: <CD8C6134-C312-41D5-B083-366F7FA1D752@me.example.com>
			//   Date: Fri, 21 Nov 2014 23:34:45 +0900
			//   From: Shironeko <shironeko@me.example.com>
			//   To: kijitora@example.jp
			//   Subject: Nyaaaaaaaaaaaaaaaaaaaaaan
			//
			// Your message cannot be delivered to the following recipients:
			//
			//   Recipient address: kijitora@example.jp
			//   Reason: Remote SMTP server has rejected address
			//   Diagnostic code: smtp;550 5.1.1 <kijitora@example.jp>... User Unknown
			//   Remote system: dns;mx.example.jp (TCP|17.111.174.67|47323|192.0.2.225|25) (6jo.example.jp ESMTP SENDMAIL-VM)
			if sisimoji.Aligned(e, []string{"  Recipient address: ", "@", "."}) ||
			   sisimoji.Aligned(e, []string{"  Original address: ",  "@", "."}) {
				//   Recipient address: @smtp.example.net:kijitora@server
				//   Original address: kijitora@example.jp
				cv := sisiaddr.S3S4(e[strings.Index(e, ": ") + 2:])
				if rfc5322.IsEmailAddress(cv) == false { continue }

				if len(v.Recipient) > 0 && cv != v.Recipient {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = cv
				recipients += 1

			} else if strings.HasPrefix(e, "  Date: ") {
				//   Date: Fri, 21 Nov 2014 23:34:45 +0900
				v.Date = e[strings.IndexByte(e, ':') + 2:]

			} else if strings.HasPrefix(e, "  Reason: ") {
				//   Reason: Remote SMTP server has rejected address
				v.Diagnosis = e[strings.IndexByte(e, ':') + 2:]

			} else if strings.HasPrefix(e, "  Diagnostic code: ") {
				//   Diagnostic code: smtp;550 5.1.1 <kijitora@example.jp>... User Unknown
				e  = strings.Replace(strings.Trim(e, " "), "Diagnostic code:", "Diagnostic-Code:", 1)
				f := rfc1894.Field(e); if len(f) == 0 { continue }
				v.Spec       = f[1]
				v.ReplyCode  = reply.Find(f[2], "")
				v.Status     = status.Find(f[2], v.ReplyCode)
				v.Diagnosis += " " + f[2]

			} else if strings.HasPrefix(e, "  Remote system: ") {
				//   Remote system: dns;mx.example.jp (TCP|17.111.174.67|47323|192.0.2.225|25)
				//     (6jo.example.jp ESMTP SENDMAIL-VM)
				v.Rhost = sisimoji.Select(e, ";", " (", 0); if v.Rhost == "" { continue }

				if cv := strings.Split(sisimoji.Select(e, " (", ")", 0), "|"); len(cv) == 5 {
					// (TCP|17.111.174.67|47323|192.0.2.225|25)
					if cv[0] != "TCP" || strings.IndexByte(v.Rhost, '.') > 0 { continue }
					if rfc791.IsIPv4Address(cv[1]) { v.Lhost = cv[1] }
					if rfc791.IsIPv4Address(cv[3]) { v.Rhost = cv[3] }
				}
			} else {
				// Original-envelope-id: 0NFC009FLKOUVMA0@mr21p30im-asmtp004.me.com
				// Reporting-MTA: dns;mr21p30im-asmtp004.me.com (tcp-daemon)
				// Arrival-date: Thu, 29 Apr 2014 23:34:45 +0000 (GMT)
				//
				// Original-recipient: rfc822;kijitora@example.jp
				// Final-recipient: rfc822;kijitora@example.jp
				// Action: failed
				// Status: 5.1.1 (Remote SMTP server has rejected address)
				// Remote-MTA: dns;mx.example.jp (TCP|17.111.174.67|47323|192.0.2.225|25)
				//  (6jo.example.jp ESMTP SENDMAIL-VM)
				// Diagnostic-code: smtp;550 5.1.1 <kijitora@example.jp>... User Unknown
				if strings.HasPrefix(e, "Status: ") {
					// Status: 5.1.1 (Remote SMTP server has rejected address)
					if v.Status    == "" { v.Status = status.Find(e, v.ReplyCode)      }
					if v.Diagnosis == "" { v.Diagnosis = e[strings.IndexByte(e, '('):] }

				} else if strings.HasPrefix(e, "Arrival-Date: ") {
					// Arrival-date: Thu, 29 Apr 2014 23:34:45 +0000 (GMT)
					if v.Date == "" { v.Date = e[strings.IndexByte(e, ':') + 2:] }

				} else if strings.HasPrefix(e, "Reporting-MTA: ") {
					// Reporting-MTA: dns;mr21p30im-asmtp004.me.com (tcp-daemon)
					if strings.IndexByte(v.Lhost, '.') > 0 { continue        }
					if cv := rfc1894.Field(e); len(cv) > 0 { v.Lhost = cv[2] }
				}
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
					if strings.Contains(e.Diagnosis, f) { e.Reason = r; break FINDREASON }
				}
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

