// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      _____                   ____  __  __ _____ ____  ____  
// | | |__   ___  ___| |_   / / _ \ _ __   ___ _ __ / ___||  \/  |_   _|  _ \|  _ \ 
// | | '_ \ / _ \/ __| __| / / | | | '_ \ / _ \ '_ \\___ \| |\/| | | | | |_) | | | |
// | | | | | (_) \__ \ |_ / /| |_| | |_) |  __/ | | |___) | |  | | | | |  __/| |_| |
// |_|_| |_|\___/|___/\__/_/  \___/| .__/ \___|_| |_|____/|_|  |_| |_| |_|   |____/ 
//                                 |_|                                              
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from penSMTPD: https://www.opensmtpd.org/
	InquireFor["OpenSMTPD"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		proceedsto := uint8(0)
		ISOPENSMTPD: for {
			if strings.Contains(bf.Headers["subject"][0], "Delivery status notification") { proceedsto++ }
			if strings.Contains(bf.Headers["from"][0], "Mailer Daemon <")                 { proceedsto++ }

			if len(bf.Headers["received"]) == 0 { break ISOPENSMTPD }
			for _, e := range bf.Headers["received"] {
				// Received: from localhost (localhost [local]);
				//   by localhost (OpenSMTPD) with ESMTPA id 1e2a9eaa;
				//   for <kijitora@example.jp>;
				if strings.Contains(e, " (OpenSMTPD) with ") { proceedsto++; break ISOPENSMTPD }
			}
			break ISOPENSMTPD
		}
		if proceedsto == 0 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"    Below is a copy of the original message:"}
		startingof := map[string][]string{
			// http://www.openbsd.org/cgi-bin/man.cgi?query=smtpd&sektion=8
			// opensmtpd-5.4.2p1/smtpd/
			//   bounce.c/317://define NOTICE_INTRO \
			//   bounce.c/318:    "    Hi!\n\n"    \
			//   bounce.c/319:    "    This is the MAILER-DAEMON, please DO NOT REPLY to this e-mail.\n"
			//   bounce.c/320:
			//   bounce.c/321:const char *notice_error =
			//   bounce.c/322:    "    An error has occurred while attempting to deliver a message for\n"
			//   bounce.c/323:    "    the following list of recipients:\n\n";
			//   bounce.c/324:
			//   bounce.c/325:const char *notice_warning =
			//   bounce.c/326:    "    A message is delayed for more than %s for the following\n"
			//   bounce.c/327:    "    list of recipients:\n\n";
			//   bounce.c/328:
			//   bounce.c/329:const char *notice_warning2 =
			//   bounce.c/330:    "    Please note that this is only a temporary failure report.\n"
			//   bounce.c/331:    "    The message is kept in the queue for up to %s.\n"
			//   bounce.c/332:    "    You DO NOT NEED to re-send the message to these recipients.\n\n";
			//   bounce.c/333:
			//   bounce.c/334:const char *notice_success =
			//   bounce.c/335:    "    Your message was successfully delivered to these recipients.\n\n";
			//   bounce.c/336:
			//   bounce.c/337:const char *notice_relay =
			//   bounce.c/338:    "    Your message was relayed to these recipients.\n\n";
			//   bounce.c/339:
			"message": []string{"    This is the MAILER-DAEMON, please DO NOT REPLY to this"},
		}
		messagesof := map[string][]string{
			// smtpd/queue.c:221|  envelope_set_errormsg(&evp, "Envelope expired");
			// smtpd/mta.c:1013|  relay->failstr = "Could not retrieve credentials";
			"expired":       []string{"Envelope expired"},
			"securityerror": []string{"Could not retrieve credentials"},
			"hostunknown":   []string{
				// smtpd/mta.c:976|  relay->failstr = "Invalid domain name";
				// smtpd/mta.c:980|  relay->failstr = "Domain does not exist";
				"Invalid domain name",
				"Domain does not exist",
			},
			"networkerror":  []string{
				//  smtpd/mta.c:972|  relay->failstr = "Temporary failure in MX lookup";
				"Address family mismatch on destination MXs",
				"All routes to destination blocked",
				"bad DNS lookup error code",
				"Could not retrieve source address",
				"Loop detected",
				"Network error on destination MXs",
				"No valid route to remote MX",
				"No valid route to destination",
				"Temporary failure in MX lookup",
			},
			"notaccept": []string{
				// smtp/mta.c:1085|  relay->failstr = "Destination seem to reject all mails";
				"Destination seem to reject all mails",
				"No MX found for domain",
				"No MX found for destination",
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

			//    Hi!
			//
			//    This is the MAILER-DAEMON, please DO NOT REPLY to this e-mail.
			//
			//    An error has occurred while attempting to deliver a message for
			//    the following list of recipients:
			//
			// kijitora@example.jp: 550 5.2.2 <kijitora@example>... Mailbox Full
			//
			//    Below is a copy of the original message:
			if sisimoji.Aligned(e, []string{"@", " "}) {
				// kijitora@example.jp: 550 5.2.2 <kijitora@example.jp>... Mailbox Full
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = e[:strings.Index(e, ":")]
				v.Diagnosis = e[strings.Index(e, ":") + 1:]
				recipients += 1
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

