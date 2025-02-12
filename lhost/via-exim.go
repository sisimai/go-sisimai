// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      _______      _           
// | | |__   ___  ___| |_   / / ____|_  _(_)_ __ ___  
// | | '_ \ / _ \/ __| __| / /|  _| \ \/ / | '_ ` _ \ 
// | | | | | (_) \__ \ |_ / / | |___ >  <| | | | | | |
// |_|_| |_|\___/|___/\__/_/  |_____/_/\_\_|_| |_| |_|

package lhost
import "fmt"
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc1123"
import "libsisimai.org/sisimai/rfc1894"
import "libsisimai.org/sisimai/rfc2045"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/smtp/reply"
import "libsisimai.org/sisimai/smtp/status"
import "libsisimai.org/sisimai/smtp/command"
import sisiaddr "libsisimai.org/sisimai/address"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from Exim Internet Mailer: https://www.exim.org/
	InquireFor["Exim"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		// X-Failed-Recipients: kijitora@example.ed.jp
		thirdparty := false
		proceedsto := uint8(0)
		messageidv := bf.Headers["message-id"][0]
		emailtitle := []string{
			"Delivery Status Notification",
			"Mail delivery failed",
			"Mail failure",
			"Message frozen",
			"Warning: message ",
			"error(s) in forwarding or filtering",
		}
		if strings.Contains(bf.Headers["from"][0], "Mail Delivery System") { proceedsto++ }
		for messageidv != "" {
			// Message-Id: <E1P1YNN-0003AD-Ga@example.org>
			if strings.Index(messageidv, "<") !=  0 { break }
			if strings.Index(messageidv, "-") !=  8 { break }
			if strings.Index(messageidv, "@") != 18 { break }
			proceedsto++; break
		}
		for _, e := range emailtitle {
			// Subject: Mail delivery failed: returning message to sender
			// Subject: Mail delivery failed
			// Subject: Message frozen
			if strings.Contains(bf.Headers["subject"][0], e) == false { continue }
			proceedsto++; break
		}

		for {
			// Exim clones of the third Parties
			// 1. McAfee Saas (Formerly MXLogic)
			if len(bf.Headers["x-mx-bounce"])    > 0 { thirdparty = true; break }
			if len(bf.Headers["x-mxl-hash"])     > 0 { thirdparty = true; break }
			if len(bf.Headers["x-mxl-notehash"]) > 0 { thirdparty = true; break }
			if strings.Contains(messageidv, "<mxl~") { thirdparty = true; break }
			break
		}
		if proceedsto < 2 && thirdparty == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{
			// deliver.c:6423|          if (bounce_return_body) fprintf(f,
			// deliver.c:6424|"------ This is a copy of the message, including all the headers. ------\n");
			// deliver.c:6425|          else fprintf(f,
			// deliver.c:6426|"------ This is a copy of the message's headers. ------\n");
			"------ This is a copy of the message, including all the headers. ------",
			"Content-Type: message/rfc822",
			"Included is a copy of the message header:\n-----------------------------------------", // MXLogic
		}
		startingof := map[string][]string{
			// Error text strings which are defined in exim/src/deliver.c
			//
			// deliver.c:6292| fprintf(f,
			// deliver.c:6293|"This message was created automatically by mail delivery software.\n");
			// deliver.c:6294|        if (to_sender)
			// deliver.c:6295|          {
			// deliver.c:6296|          fprintf(f,
			// deliver.c:6297|"\nA message that you sent could not be delivered to one or more of its\n"
			// deliver.c:6298|"recipients. This is a permanent error. The following address(es) failed:\n");
			// deliver.c:6299|          }
			// deliver.c:6300|        else
			// deliver.c:6301|          {
			// deliver.c:6302|          fprintf(f,
			// deliver.c:6303|"\nA message sent by\n\n  <%s>\n\n"
			// deliver.c:6304|"could not be delivered to one or more of its recipients. The following\n"
			// deliver.c:6305|"address(es) failed:\n", sender_address);
			// deliver.c:6306|          }
			"alias":          []string{" an undisclosed address"},
			"command":        []string{"SMTP error from remote ", "LMTP error after "},
			"deliverystatus": []string{"Content-Type: message/delivery-status"},
			"frozen":         []string{" has been frozen", " was frozen on arrival"},
			"message":        []string{
				"This message was created automatically by mail delivery software.",
				"A message that you sent was rejected by the local scannning code",
				"A message that you sent contained one or more recipient addresses ",
				"A message that you sent could not be delivered to all of its recipients",
				" has been frozen",
				" was frozen on arrival",
				" router encountered the following error(s):",
			},
		}
		messagesof := map[string][]string{
			// find exim/ -type f -exec grep 'message = US' {} /dev/null \;
			// route.c:1158|  DEBUG(D_uid) debug_printf("getpwnam() returned NULL (user not found)\n");
			"userunknown": []string{"user not found"},
			// transports/smtp.c:3524|  addr->message = US"all host address lookups failed permanently";
			// routers/dnslookup.c:331|  addr->message = US"all relevant MX records point to non-existent hosts";
			// route.c:1826|  uschar *message = US"Unrouteable address";
			"hostunknown": []string{
				"all host address lookups failed permanently",
				"all relevant MX records point to non-existent hosts",
				"Unrouteable address",
			},
			// transports/appendfile.c:2567|  addr->user_message = US"mailbox is full";
			// transports/appendfile.c:3049|  addr->message = string_sprintf("mailbox is full "
			// transports/appendfile.c:3050|  "(quota exceeded while writing to file %s)", filename);
			"mailboxfull": []string{
				"mailbox is full",
				"error: quota exceed",
			},
			// routers/dnslookup.c:328|  addr->message = US"an MX or SRV record indicated no SMTP service";
			// transports/smtp.c:3502|  addr->message = US"no host found for existing SMTP connection";
			"notaccept": []string{
				"an MX or SRV record indicated no SMTP service",
				"no host found for existing SMTP connection",
			},
			// parser.c:666| *errorptr = string_sprintf("%s (expected word or \"<\")", *errorptr);
			// parser.c:701| if(bracket_count++ > 5) FAILED(US"angle-brackets nested too deep");
			// parser.c:738| FAILED(US"domain missing in source-routed address");
			// parser.c:747| : string_sprintf("malformed address: %.32s may not follow %.*s",
			"syntaxerror": []string{
				"angle-brackets nested too deep",
				`expected word or "<"`,
				"domain missing in source-routed address",
				"malformed address:",
			},
			// deliver.c:5614|  addr->message = US"delivery to file forbidden";
			// deliver.c:5624|  addr->message = US"delivery to pipe forbidden";
			// transports/pipe.c:1156|  addr->user_message = US"local delivery failed";
			"systemerror": []string{
				"delivery to file forbidden",
				"delivery to pipe forbidden",
				"local delivery failed",
				"LMTP error after ",
			},
			// deliver.c:5425|  new->message = US"Too many \"Received\" headers - suspected mail loop";
			"contenterror": []string{`Too many "Received" headers`},
		}
		delayedfor := []string{
			// retry.c:902|  addr->message = (addr->message == NULL)? US"retry timeout exceeded" :
			// deliver.c:7475|  "No action is required on your part. Delivery attempts will continue for\n"
			// smtp.c:3508|  US"retry time not reached for any host after a long failure period" :
			// smtp.c:3508|  US"all hosts have been failing for a long time and were last tried "
			//                 "after this message arrived";
			// deliver.c:7459|  print_address_error(addr, f, US"Delay reason: ");
			// deliver.c:7586|  "Message %s has been frozen%s.\nThe sender is <%s>.\n", message_id,
			// receive.c:4021|  moan_tell_someone(freeze_tell, NULL, US"Message frozen on arrival",
			// receive.c:4022|  "Message %s was frozen on arrival by %s.\nThe sender is <%s>.\n",
			"retry timeout exceeded",
			"No action is required on your part",
			"retry time not reached for any host after a long failure period",
			"all hosts have been failing for a long time and were last tried",
			"Delay reason: ",
			"has been frozen",
			"was frozen on arrival by ",
		}

		if strings.Contains(bf.Payload, "\n----- This ") {
			// There are extremely rare cases where there are only five hyphens.
			// https://github.com/sisimai/set-of-emails/blob/master/maildir/bsd/lhost-exim-05.eml
			// ----- This is a copy of the message, including all the headers. ------
			bf.Payload = strings.Replace(bf.Payload, "\n----- This ", "\n------ This ", 1)
		}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)              // Points the current cursor position
		nextcursor := uint8(0)
		recipients := 0                     // The number of 'Final-Recipient' header
		anotherone := []string{""}          // Keeping another error messages
		rightindex := uint8(0)              // The last index number of dscontents
		boundary00 := ""                    // Boundary sting
		v          := &(dscontents[len(dscontents) - 1])

		if bf.Headers["content-type"][0] != "" {
			// Get the boundary string and set regular expression for matching with the boundary string.
			boundary00 = rfc2045.Boundary(bf.Headers["content-type"][0], 0)
		}

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				for _, f := range startingof["message"] {
					// Check the message defined in startingof["message"], ["frozen"]
					if strings.Contains(e, f) == false { continue }
					readcursor |= indicators["deliverystatus"]

					for _, g := range startingof["frozen"] {
						// Goes to the next loop if the string does not contain "frozen" message
						if strings.Contains(e, g) == false { continue }
					}
				}
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// This message was created automatically by mail delivery software.
			//
			// A message that you sent could not be delivered to one or more of its
			// recipients. This is a permanent error. The following address(es) failed:
		    //
			//  kijitora@example.jp
			//    SMTP error from remote mail server after RCPT TO:<kijitora@example.jp>:
			//    host neko.example.jp [192.0.2.222]: 550 5.1.1 <kijitora@example.jp>... User Unknown
			cv := ""
			ce := false
			for {
				// Check whether the line matches the following conditions or not
				if strings.HasPrefix(e, "  ") == false  { break } // The line should start with "  " (2 spaces)
				if strings.Index(e, "@") < 2            { break } // "@" should be included (email)
				if strings.Index(e, ".") < 2            { break } // "." should be included (domain part)
				if strings.Contains(e, "pipe to |")     { break } // Exclude "pipe to /path/to/prog" line
				if e[2:3] == " "                        { break } // The 3rd character is " "
				if thirdparty == false && e[2:3] == "<" { break } // MXLogic returns "  <neko@example.jp>: ..."

				ce = true; break
			}

			if ce == true || sisimoji.ContainsAny(e, startingof["alias"]) {
				// The line is including an email address
				if sisimoji.ContainsAny(e, startingof["alias"]) {
					// The line does not include an email address
					// deliver.c:4549|  printed = US"an undisclosed address";
					//   an undisclosed address
					//     (generated from kijitora@example.jp)
					cv = e[2:]

				} else {
					//   kijitora@example.jp
					//   sabineko@example.jp: forced freeze
					//   mikeneko@example.jp <nekochan@example.org>: ...
					p1 := strings.Index(e, "<")
					p2 := strings.Index(e, ">:")
					if p1 > 1 && p2 > 1 {
						// There are an email address and an error message in the line
						// parser.c:743| while (bracket_count-- > 0) if (*s++ != '>')
						// parser.c:744|   {
						// parser.c:745|   *errorptr = s[-1] == 0
						// parser.c:746|     ? US"'>' missing at end of address"
						// parser.c:747|     : string_sprintf("malformed address: %.32s may not follow %.*s",
						// parser.c:748|     s-1, (int)(s - US mailbox - 1), mailbox);
						// parser.c:749|   goto PARSE_FAILED;
						// parser.c:750|   }
						cv = sisiaddr.S3S4(e[p1:p2 + 1])
						v.Diagnosis = sisimoji.Sweep(e[p2 + 1:])

					} else {
						// There is an email address only in the line, such as "  kijitora@example.jp"
						// --- OR ---
						// "  kijitora@example.jp: forced freeze"
						p3 := strings.LastIndex(e, ": "); if p3 < 0 { p3 = len(e) }
						cv  = sisiaddr.S3S4(e[2:p3])
					}
					if rfc5322.IsEmailAddress(cv) == false { continue }
				}

				if len(v.Recipient) > 0 && cv != v.Recipient {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					anotherone = append(anotherone, "")
					rightindex++
					v = &(dscontents[rightindex])
				}
				v.Recipient = cv
				recipients++

			} else if strings.Contains(e, " (generated from ") || strings.Contains(e, " generated by ") {
				//     (generated from kijitora@example.jp)
				//  pipe to |/bin/echo "Some pipe output"
				//    generated by userx@myhost.test.ex
				for _, f := range strings.Split(e, " ") {
					// Get the alias adress
					if strings.Contains(f, "@") == false { continue }
					v.Alias = sisiaddr.S3S4(f); break
				}
			} else {
				if sisimoji.ContainsAny(e, startingof["frozen"]) {
					// Message *** has been frozen by the system filter.
					// Message *** was frozen on arrival by ACL.
					anotherone[rightindex] = e + " "

				} else if boundary00 != "" {
					// --NNNNNNNNNN-eximdsn-MMMMMMMMMM
					// Content-type: message/delivery-status
					// ...
					if rfc1894.Match(e) > 0 {
						// "e" matched with any field defined in RFC3464
						o := rfc1894.Field(e); if len(o) == 0 { continue }

						if o[3] == "addr" {
							// Final-Recipient: rfc822;|/bin/echo "Some pipe output"
							if o[0]   != "final-recipient" { continue }
							if v.Spec != ""                { continue }
							if strings.Contains(o[2], "@") { v.Spec = "SMTP" } else { v.Spec = "X-UNIX" }

						} else if o[3] == "code" {
							// Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
							v.Spec      = strings.ToUpper(o[1])
							v.Diagnosis = o[2]

						} else {
							// Other DSN fields defined in RFC3464
							v.Update(v.AsRFC1894(o[0]), o[2])
						}
					} else {
						// There are other error messages?
						if nextcursor > 0 { continue }

						// Content-type: message/delivery-status
						if strings.HasPrefix(e, startingof["deliverystatus"][0]) { nextcursor = 1 }
						if strings.HasPrefix(e, " ") { anotherone[rightindex] += e + " " }
					}
				} else {
					// There is no boundary string in "boundary00"
					if len(dscontents) == recipients {
						// This line is an error message
						v.Diagnosis += " " + e

					} else {
						// Error message when the email address above does not include '@' and 
						// domain part such as
						//   pipe to |/path/to/prog ...
						//     generated by kijitora@example.com
						if strings.HasPrefix(e, "  ") == false { continue }
						v.Diagnosis += " " + e 
					}
				}
			}
		}

		if recipients > 0 {
			// Check "an undisclosed address", "unroutable address"
			for j, _ := range dscontents {
				// Replace the recipient address with the value of "alias"
				e := &(dscontents[j])
				if e.Alias == "" { continue }
				if e.Recipient == "" || strings.Contains(e.Recipient, "@") == false {
					// The value of "recipient" is empty or does not include "@"
					e.Recipient = e.Alias
				}
			}
		} else {
			// Fallback for getting recipient addresses
			if len(bf.Headers["x-failed-recipients"]) > 0 {
				// X-Failed-Recipients: kijitora@example.jp
				rcptinhead := strings.Split(bf.Headers["x-failed-recipients"][0], ",")
				for j, _ := range rcptinhead { rcptinhead[j] = strings.Trim(rcptinhead[j], " ") }
				recipients  = len(rcptinhead)

				for _, e := range rcptinhead {
					// Insert each recipient address into "dscontents"
					dscontents[len(dscontents) - 1].Recipient = e
					if len(dscontents) == recipients { continue }
					dscontents = append(dscontents, sis.DeliveryMatter{})
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		// Get the name of the local MTA
		// Received: from marutamachi.example.org (c192128.example.net [192.0.2.128])
		receivedby := bf.Headers["received"][len(bf.Headers["received"]) - 1]
		recvdtoken := rfc5322.Received(receivedby)

		for j, _ := range dscontents {
			// Check the error message, the rhost, the lhost, and the smtp command.
			e := &(dscontents[j])

			if e.Diagnosis == "" && boundary00 != "" {
				// Empty Diagnostic-Code: or error message
				//
				// --NNNNNNNNNN-eximdsn-MMMMMMMMMM
				// Content-type: message/delivery-status
				//
				// Reporting-MTA: dns; the.local.host.name
				//
				// Action: failed
				// Final-Recipient: rfc822;/a/b/c
				// Status: 5.0.0
				//
				// Action: failed
				// Final-Recipient: rfc822;|/p/q/r
				// Status: 5.0.0
				e.Diagnosis = dscontents[0].Diagnosis
				if e.Spec        == "" { e.Spec = dscontents[0].Spec   }
				if anotherone[0] != "" {
					if len(anotherone) <= j { anotherone = append(anotherone, "") }
					anotherone[j] = anotherone[0]
				}
			}

			if len(anotherone) > j && anotherone[j] != "" {
				// Copy alternative error message
				if e.Diagnosis == "" { e.Diagnosis = anotherone[j] }

				if strings.HasPrefix(e.Diagnosis, "-") || strings.HasSuffix(e.Diagnosis, "__") {
					// Override the value of diagnostic code message
					e.Diagnosis = anotherone[j]

				} else if len(e.Diagnosis) < len(anotherone[j]) {
					// Override the value of diagnostic code message with the value of alterrors
					// because the latter includes the former.
					anotherone[j] = sisimoji.Squeeze(anotherone[j], " ")
					if strings.Contains(strings.ToLower(anotherone[j]), strings.ToLower(e.Diagnosis)) {
						// anotherone[j] contains the same error message stored in e.Diagnosis
						e.Diagnosis = anotherone[j]
					}
				}
			}

			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			p1 := strings.Index(e.Diagnosis, "__"); if p1 > 1 { e.Diagnosis = e.Diagnosis[0:p1] }

			if e.Rhost == "" { e.Rhost = rfc1123.Find(e.Diagnosis) }
			if e.Lhost == "" { e.Lhost = recvdtoken[0] }

			if e.Command == "" {
				// Get the SMTP command name for the session
				for _, f := range startingof["command"] {
					// Find the last SMTP command from e.Diagnosis
					if strings.Contains(e.Diagnosis, f) == false              { continue }
					e.Command = command.Find(e.Diagnosis); if e.Command == "" { continue }
					break
				}

				// Detect the bounce reason by using the SMTP command
				if e.Command == "EHLO" || e.Command == "HELO" {
					// HELO | Connected to 192.0.2.135 but my name was rejected.
					e.Reason = "blocked"

				} else if e.Command == "MAIL" {
					// MAIL | Connected to 192.0.2.135 but sender was rejected.
					e.Reason = "onhold"

				} else {
					// Find any error message string defined in "messagesof" from e.Diagnosis
					for r := range messagesof {
						// The key is a bounce reason name
						if sisimoji.ContainsAny(e.Diagnosis, messagesof[r]) == false { continue }
						e.Reason = r; break
					}

					if e.Reason == "" {
						// The reason is "expired", or "mailererror"
						if sisimoji.ContainsAny(e.Diagnosis, delayedfor) == true {
							// The reason is "expired"
							e.Reason = "expired"

						} else {
							// The reason is "mailererror"
							if strings.Contains(e.Diagnosis, "pipe to |") { e.Reason = "mailererror" }
						}
					}
				}
			}

			// Prefer the value of smtp reply code in Diagnostic-Code: field
			// See set-of-emails/maildir/bsd/exim-20.eml
			//
			//   Action: failed
			//   Final-Recipient: rfc822;userx@test.ex
			//   Status: 5.0.0
			//   Remote-MTA: dns; 127.0.0.1
			//   Diagnostic-Code: smtp; 450 TEMPERROR: retry timeout exceeded
			//
			// The value of "Status:" indicates permanent error but the value of SMTP reply code in
			// Diagnostic-Code: field is "TEMPERROR"!!!!
			re := e.Reason
			cr := reply.Find(e.Diagnosis, e.Status)
			cs := status.Find(e.Diagnosis, cr)
			cv := ""

			if strings.HasPrefix(cr, "4") || re == "expired" || re == "mailboxfull" {
				// Set the pseudo status code as a temporary error
				cv = status.Code(re, true)

			} else {
				// Set the pseudo status code as a permanent error
				cv = status.Code(re, false)
			}
			if e.ReplyCode == "" { e.ReplyCode = cr }
			if e.Status    == "" { e.Status = status.Prefer(cs, cv, cr) }
		}

		for emailparts[1] == "" {
			// There is no original message in the warning message like the follwing:
			//
			// This message was created automatically by mail delivery software.
			// A message that you sent has not yet been delivered to one or more of its
			// recipients after more than 24 hours on the queue on neko.example.com.
			//
			// The message identifier is:     neko222-nyaaan-22
			// The subject of the message is: Nyaan
			// The date of the message is:    Thu, 22 Apr 2016 23:34:45 +0900
			//
			// The address to which the message has not yet been delivered is:
			//
			//   kijitora@example.co.jp
			//     host mta-2.example.co.jp [192.0.2.222]
			//     Delay reason: SMTP error from remote mail server after MAIL FROM:<sironeko@neko.example.com> SIZE=1024:
			//     450 service permits 2 unverifyable sending IPs - neko.example.com is not 203.0.113.222
			//
			// No action is required on your part. Delivery attempts will continue for
			// some time, and this warning may be repeated at intervals if the message
			// remains undelivered. Eventually the mail delivery software will give up,
			// and when that happens, the message will be returned to you.
			emailparts[1] += fmt.Sprintf("To: <%s>\n", dscontents[0].Recipient)

			p1 := strings.Index(bf.Payload, "The date of the message is: ");    if p1 < 0 { break }
			p2 := sisimoji.IndexOnTheWay(bf.Payload, "\n", p1);                 if p2 < 0 { break }
			emailparts[1] += fmt.Sprintf("Date: %s\n", bf.Payload[p1 + 31:p2])

			p1  = strings.Index(bf.Payload, "The subject of the message is: "); if p1 < 0 { break }
			p2  = sisimoji.IndexOnTheWay(bf.Payload, "\n", p1);                 if p2 < 0 { break }
			emailparts[1] += fmt.Sprintf("Subject: %s\n", bf.Payload[p1 + 31:p2])
			break
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

