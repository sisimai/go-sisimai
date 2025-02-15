// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      ______                 _ _ 
// | | |__   ___  ___| |_   / / ___|_ __ ___   __ _(_) |
// | | '_ \ / _ \/ __| __| / / |  _| '_ ` _ \ / _` | | |
// | | | | | (_) \__ \ |_ / /| |_| | | | | | | (_| | | |
// |_|_| |_|\___/|___/\__/_/  \____|_| |_| |_|\__,_|_|_|

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/rfc1123"
import "libsisimai.org/sisimai/smtp/status"
import sisiaddr "libsisimai.org/sisimai/address"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from Gmail: https://mail.google.com/
	InquireFor["Gmail"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		// Google Mail
		// From: Mail Delivery Subsystem <mailer-daemon@googlemail.com>
		// Received: from vw-in-f109.1e100.net [74.125.113.109] by ...
		//
		// * Check the body part
		//   This is an automatically generated Delivery Status Notification
		//   Delivery to the following recipient failed permanently:
		//
		//        recipient-address-here@example.jp
		//
		//   Technical details of permanent failure:
		//   Google tried to deliver your message, but it was rejected by the
		//   recipient domain. We recommend contacting the other email provider
		//   for further information about the cause of this error. The error
		//   that the other server returned was:
		//   550 550 <recipient-address-heare@example.jp>: User unknown (state 14).
		//
		//   -- OR --
		//   THIS IS A WARNING MESSAGE ONLY.
		//
		//   YOU DO NOT NEED TO RESEND YOUR MESSAGE.
		//
		//   Delivery to the following recipient has been delayed:
		//
		//        mailboxfull@example.jp
		//
		//   Message will be retried for 2 more day(s)
		//
		//   Technical details of temporary failure:
		//   Google tried to deliver your message, but it was rejected by the recipient
		//   domain. We recommend contacting the other email provider for further infor-
		//   mation about the cause of this error. The error that the other server re-
		//   turned was: 450 450 4.2.2 <mailboxfull@example.jp>... Mailbox Full (state 14).
		//
		//   -- OR --
		//
		//   Delivery to the following recipient failed permanently:
		//
		//        userunknown@example.jp
		//
		//   Technical details of permanent failure:=20
		//   Google tried to deliver your message, but it was rejected by the server for=
		//    the recipient domain example.jp by mx.example.jp. [192.0.2.59].
		//
		//   The error that the other server returned was:
		//   550 5.1.1 <userunknown@example.jp>... User Unknown
		if strings.Contains(bf.Headers["from"][0], "<mailer-daemon@googlemail.com>")  == false { return sis.RisingUnderway{} }
		if strings.Contains(bf.Headers["subject"][0], "Delivery Status Notification") == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"----- Original message -----", "----- Message header follows -----"}
		startingof := map[string][]string{
			"message": []string{"Delivery to the following recipient"},
			"error":   []string{"The error that the other server returned was:"},
		}
		messagesof := map[string][]string{
			"expired": []string{
				"DNS Error: Could not contact DNS servers",
				"Delivery to the following recipient has been delayed",
				"The recipient server did not accept our requests to connect",
			},
			"hostunknown": []string{
				"DNS Error: Domain name not found",
				"DNS Error: DNS server returned answer with no data",
			},
		}
		statetable := map[string][2]string{
			// Technical details of permanent failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 500 Remote server does not support TLS (state 6).
			"6":  [2]string{"MAIL", "failedstarttls"},

			// https://www.google.td/support/forum/p/gmail/thread?tid=08a60ebf5db24f7b&hl=en
			// Technical details of permanent failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 535 SMTP AUTH failed with the remote server. (state 8).
			"8":  [2]string{"AUTH", "systemerror"},

			// https://www.google.co.nz/support/forum/p/gmail/thread?tid=45208164dbca9d24&hl=en
			// Technical details of temporary failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 454 454 TLS missing certificate: error:0200100D:system library:fopen:Permission denied (//4.3.0) (state 9).
			"9":  [2]string{"AUTH", "failedstarttls"},

			// https://www.google.com/support/forum/p/gmail/thread?tid=5cfab8c76ec88638&hl=en
			// Technical details of permanent failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 500 Remote server does not support SMTP Authenticated Relay (state 12).
			"12": [2]string{"AUTH", "relayingdenied"},

			// Technical details of permanent failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 550 550 5.7.1 <****@gmail.com>... Access denied (state 13).
			"13": [2]string{"EHLO", "blocked"},

			// Technical details of permanent failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 550 550 5.1.1 <******@*********.**>... User Unknown (state 14).
			// 550 550 5.2.2 <*****@****.**>... Mailbox Full (state 14).
			"14": [2]string{"RCPT", "userunknown"},

			// https://www.google.cz/support/forum/p/gmail/thread?tid=7090cbfd111a24f9&hl=en
			// Technical details of permanent failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 550 550 5.7.1 SPF unauthorized mail is prohibited. (state 15).
			// 554 554 Error: no valid recipients (state 15).
			"15": [2]string{"DATA", "filtered"},

			// https://www.google.com/support/forum/p/Google%20Apps/thread?tid=0aac163bc9c65d8e&hl=en
			// Technical details of permanent failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 550 550 <****@***.**> No such user here (state 17).
			// 550 550 //5.1.0 Address rejected ***@***.*** (state 17).
			"17": [2]string{"DATA", "filtered"},

			// Technical details of permanent failure:
			// Google tried to deliver your message, but it was rejected by the recipient domain.
			// We recommend contacting the other email provider for further information about the
			// cause of this error. The error that the other server returned was:
			// 550 550 Unknown user *****@***.**.*** (state 18).
			"18": [2]string{"DATA", "filtered"},
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
			if readcursor & indicators["deliverystatus"] == 0 || e == "" { continue }

			// Technical details of permanent failure:=20
			// Google tried to deliver your message, but it was rejected by the recipient =
			// domain. We recommend contacting the other email provider for further inform=
			// ation about the cause of this error. The error that the other server return=
			// ed was: 554 554 5.7.0 Header error (state 18).
			//
			// -- OR --
			//
			// Technical details of permanent failure:=20
			// Google tried to deliver your message, but it was rejected by the server for=
			// the recipient domain example.jp by mx.example.jp. [192.0.2.49].
			//
			// The error that the other server returned was:
			// 550 5.1.1 <userunknown@example.jp>... User Unknown
			if strings.HasPrefix(e, " ") && strings.Index(e, "@") > 0 {
				// Delivery to the following recipient failed permanently:
				//
				//      userunknown@example.jp
				//
				// Technical details of permanent failure:
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				cv := sisiaddr.S3S4(strings.Trim(e, " "))
				if rfc5322.IsEmailAddress(cv) == true { v.Recipient = cv; recipients++ }

			} else {
				// Error message lines except "    neko@example.jp" line
				v.Diagnosis += e + " "
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			e.Rhost     = rfc1123.Find(e.Diagnosis)

			if cv := sisimoji.Select(e.Diagnosis, " (state ", ")", 0); len(statetable[cv]) > 0 {
				// Find "(state 18)" and pick "18" as a key of statetable
				e.Command = statetable[cv][0]
				e.Reason  = statetable[cv][1]
			}

			if e.Reason == "" {
				// There is no state code in the error message
				FINDREASON: for r := range messagesof {
					// The key name is a bounce reason name
					for _, f := range messagesof[r] {
						// Try to find an error message including lower-cased string listed in messagesof
						if strings.Contains(e.Diagnosis, f) { e.Reason = r; break FINDREASON }
					}
				}
			}
			if e.Reason == "" { continue }

			// Set a pseudo status code and override the bounce reason
			e.Status = status.Find(e.Diagnosis, e.ReplyCode);  if e.Status == "" { continue }
			if strings.Contains(e.Status, ".0") == false { e.Reason = status.Name(e.Status) }
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

