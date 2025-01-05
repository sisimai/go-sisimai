// Copyright (C) 2020-2022,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      __               _    __ _      
// | | |__   ___  ___| |_   / / __   ___  ___| |_ / _(_)_  __
// | | '_ \ / _ \/ __| __| / / '_ \ / _ \/ __| __| |_| \ \/ /
// | | | | | (_) \__ \ |_ / /| |_) | (_) \__ \ |_|  _| |>  < 
// |_|_| |_|\___/|___/\__/_/ | .__/ \___/|___/\__|_| |_/_/\_\
//                           |_|                             
import "fmt"
import "strings"
import "strconv"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import "sisimai/smtp/command"
import "sisimai/smtp/transcript"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Postfix https://www.postfix.org/
	InquireFor["Postfix"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Headers)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Payload)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Headers["x-aol-ip"]) > 0 { return sis.RisingUnderway{} } // X-AOL-IP: 192.0.2.1

		proceedsto := uint8(0)
		if strings.Index(bf.Headers["subject"][0], "SMTP server: errors from ") > 0 {
			// src/smtpd/smtpd_chat.c:|337: post_mail_fprintf(notice, "Subject: %s SMTP server: errors from %s",
			// src/smtpd/smtpd_chat.c:|338:   var_mail_name, state->namaddr);
			proceedsto = 2

		} else if bf.Headers["subject"][0] == "Undelivered Mail Returned to Sender" {
			// Subject: Undelivered Mail Returned to Sender
			proceedsto = 1
		}
		if proceedsto == 0 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"}
		startingof := map[string][][]string{
			// Postfix manual - bounce(5) - http://www.postfix.org/bounce.5.html
			"message": [][]string{
				[]string{"The ", "Postfix "},           // The Postfix program, The Postfix on <os> program
				[]string{"The ", "mail system"},        // The mail system
				[]string{"The ", "program"},            // The <name> pogram
				[]string{"This is the", "Postfix"},     // This is the Postfix program
				[]string{"This is the", "mail system"}, // This is the mail system at host <hostname>
			},
		}

		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		nomessages := false               // Delivery report unavailable
		anotherset := map[string]string{} // Another error information
		commandset := []string{}          // "in reply to * command" list
		v          := &(dscontents[len(dscontents) - 1])

		if proceedsto == 2 {
			// The message body starts with "Transcript of session follows."
			transcript := transcript.Rise(emailparts[0], "In:", "Out:")
			if len(transcript) == 0 { return sis.RisingUnderway{} }

			for _, e := range transcript {
				// Pick email addresses, error messages, and the last SMTP command.
				v  = &(dscontents[len(dscontents) - 1])
				p := e.Response

				if e.Command == "EHLO" || e.Command == "HELO" {
					// Use the argument of EHLO/HELO command as a value of "lhost"
					v.Lhost = e.Argument

				} else if e.Command == "MAIL" {
					// Set the argument of "MAIL" command to pseudo To: header of the original message
					if len(emailparts[1]) == 0 { emailparts[1] += fmt.Sprintf("To: %s\n", e.Argument) }

				} else if e.Command == "RCPT" {
					// RCPT TO: <...>
					if len(v.Recipient) > 0 {
						// There are multiple recipient addresses in the transcript of the session
						dscontents = append(dscontents, sis.DeliveryMatter{})
						v = &(dscontents[len(dscontents) - 1])
					}
					v.Recipient = e.Argument
					recipients += 1
				}
				reply, nyaan := strconv.ParseUint(p.Reply, 10, 16)
				if nyaan != nil || reply < 400 { continue }

				commandset = append(commandset, e.Command)
				if len(v.Diagnosis) == 0 { v.Diagnosis = strings.Join(p.Text, " ") }
				if len(v.ReplyCode) == 0 { v.ReplyCode = p.Reply  }
				if len(v.Status)    == 0 { v.Status    = p.Status }
			}
		} else {
			// The message body is a general bounce mail message of Postfix
			fieldtable := rfc1894.FIELDTABLE()
			readcursor := uint8(0)     // Points the current cursor position
			readslices := []string{""} // Copy each line for later reference

			for j, e := range(strings.Split(emailparts[0], "\n")) {
				// Read error messages and delivery status lines from the head of the email to the
				// previous line of the beginning of the original message.
				readslices = append(readslices, e) // Save the current line for the next loop

				if readcursor == 0 {
					// Beginning of the bounce message or message/delivery-status part
					for _, a := range startingof["message"] {
						if sisimoji.Aligned(e, a) { readcursor |= indicators["deliverystatus"]; break }
					}
					continue
				}
				if readcursor & indicators["deliverystatus"] == 0 { continue }

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
						if strings.ToUpper(o[1]) == "X-POSTFIX" { v.Spec = "SMTP" }
						v.Diagnosis = o[2]

					} else {
						// Other DSN fields defined in RFC3464
						v.Update(v.AsRFC1894(o[0]), o[2]); if f != 1 { continue }

						// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
						permessage[z] = o[2]
						if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }
					}
				} else {
					// If you do so, please include this problem report. You can
					// delete your own text from the attached returned message.
					//
					//           The mail system
					//
					// <userunknown@example.co.jp>: host mx.example.co.jp[192.0.2.153] said: 550
					// 5.1.1 <userunknown@example.co.jp>... User Unknown (in reply to RCPT TO command)
					if strings.HasPrefix(readslices[j], "Diagnostic-Code:") && strings.HasPrefix(e, " ") {
						// Continued line of the value of Diagnostic-Code field
						v.Diagnosis += fmt.Sprintf(" %s", sisimoji.Sweep(e))
						readslices[j + 1] = "Diagnostic-Code: " + e

					} else if sisimoji.Aligned(e, []string{"X-Postfix-Sender:", "rfac822;", "@"}) {
						// X-Postfix-Sender: rfc822; shironeko@example.org
						emailparts[1] += fmt.Sprintf("X-Postfix-Sender: %s\n", strings.Trim(strings.SplitN(e, ";", 2)[1], " "))

					} else {
						// Alternative error messages and recipients
						if strings.Contains(e, "(in reply to ") || strings.Contains(e, "command)") {
							// Find an SMTP command from a string like the following:
							// 5.1.1 <userunknown@example.co.jp>... User Unknown (in reply to RCPT TO
							cv := command.Find(e)
							if len(cv) > 0                      { commandset = append(commandset, cv) }
							if len(anotherset["diagnosis"]) > 0 { anotherset["diagnosis"] += " " + e  }

						} else if sisimoji.Aligned(e, []string{"<", "@", ">", "(expanded from <", "):"}) {
							// <r@example.ne.jp> (expanded from <kijitora@example.org>): user ...
							// OR
							// <kijitora@exmaple.jp>: ...
							p1 := strings.Index(e, "> ")
							p2 := strings.Index(e[p1:], "(expanded from ")
							p3 := strings.Index(e[p2 + 14:], ">):")
							anotherset["recipient"] = sisiaddr.S3S4(e[0:p1])
							anotherset["alias"]     = sisiaddr.S3S4(e[p2 + 15:])
							anotherset["diagnosis"] = e[p3 + 3:]

						} else if strings.HasPrefix(e, "<") && sisimoji.Aligned(e, []string{"<", "@", ">:"}) {
							// <kijitora@exmaple.jp>: ...
							anotherset["recipient"] = sisiaddr.S3S4(e[0:strings.Index(e, ">") + 1])
							anotherset["diagnosis"] = e[strings.Index(e, ">:") + 2:]

						} else if strings.Contains(e, "--- Delivery report unavailable ---") {
							// postfix-3.1.4/src/bounce/bounce_notify_util.c
							// bounce_notify_util.c:602|if (bounce_info->log_handle == 0
							// bounce_notify_util.c:602||| bounce_log_rewind(bounce_info->log_handle)) {
							// bounce_notify_util.c:602|if (IS_FAILURE_TEMPLATE(bounce_info->template)) {
							// bounce_notify_util.c:602|    post_mail_fputs(bounce, "");
							// bounce_notify_util.c:602|    post_mail_fputs(bounce, "\t--- delivery report unavailable ---");
							// bounce_notify_util.c:602|    count = 1;              /* xxx don't abort */
							// bounce_notify_util.c:602|}
							// bounce_notify_util.c:602|} else {
							nomessages = true

						} else {
							// Get the error message continued from the previous line
							if len(anotherset["diagnosis"]) == 0 { continue }
							if strings.HasPrefix(e, "    ") { anotherset["diagnosis"] += " " + strings.Trim(e[4:], " ") }
						}
					}
				} // End of message/delivery-status
			}
		}

		if recipients == 0 {
			// Fallback: get a recipient address from error messages
			if len(anotherset["recipient"]) > 0 {
				// Set a recipient address saved in "anotherset"
				v.Recipient = anotherset["recipient"]
				recipients += 1

			} else {
				// Get a recipient address from message/rfc822 part if the delivery report was unavailable:
				// "--- Delivery report unavailable ---"
				for {
					p1 := strings.Index(emailparts[1], "\nTo: ")
					if nomessages == false         { break }
					if p1 < 1                      { break }
					if p1 + 6 > len(emailparts[1]) { break }

					// Try to get a recipient address from To: field in the original message at message/rfc822 part
					p2 := sisimoji.IndexOnTheWay(emailparts[1], "\n", p1 + 1)
					cv := emailparts[1][p1 + 5:p2 + 1]
					dscontents[len(dscontents) - 1].Recipient = sisiaddr.S3S4(cv)
					recipients += 1
					break
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

			if len(anotherset["diagnosis"]) > 0 {
				// Copy alternative error message to e.Diagnosis
				anotherset["diagnosis"] = sisimoji.Sweep(anotherset["diagnosis"])
				if len(e.Diagnosis) == 0 { e.Diagnosis = anotherset["diagnosis"] }

				if sisimoji.ContainsOnlyNumbers(e.Diagnosis) {
					// Override the value of diagnostic code message when the value of e.Diagnosis
					// contains only numbers
					e.Diagnosis = anotherset["diagnosis"]

				} else {
					// More detailed error message is in anotherset
					as := "" // The value of SMTP Status Code picked from anotherset["diagnosis"]
					ar := "" // The value of SMTP  Reply Code picked from anotherset["diagnosis"]

					if len(e.Status) == 0 || strings.HasSuffix(e.Status, ".0.0") {
						// Check the value of D.S.N. in "anotherset"
						// The delivery status code is neither an empty nor *.0.0
						as = status.Find(anotherset["diagnosis"], e.ReplyCode)
						if len(as) > 0 && strings.HasSuffix(as, ".0.0") == false { e.Status = as }
					}

					if len(e.ReplyCode) == 0 || strings.HasSuffix(e.ReplyCode, "00") {
						// Check the value of the SMTP reply code in anotherset
						// The SMTP reply code is neither an empty nor *00 
						ar = reply.Find(anotherset["diagnosis"], e.Status)
						if len(ar) > 0 && strings.HasSuffix(ar, "00") == false { e.ReplyCode = ar }
					}

					for {
						// Replace e.Diagnosis with the value of anotherset["diagnosis"] when all
						// the following conditions have not matched.
						if len(as) + len(ar) == 0                                  { break }
						if len(anotherset["diagnosis"]) < len(e.Diagnosis)         { break }
						if strings.Index(anotherset["diagnosis"], e.Diagnosis) < 0 { break }

						e.Diagnosis = anotherset["diagnosis"]
						break
					}
				}
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			if len(commandset) > 0 {
				// Set an SMTP command name picked from the error message
				e.Command = commandset[0]

			} else {
				// There is no SMTP command
				e.Command = command.Find(e.Diagnosis)
				if len(e.Command) == 0 {
					// <kijitora@example.org>: host r2.example.org[198.51.100.18] refused to talk to me:
					if strings.Contains(e.Diagnosis, "refused to talk to me:") { e.Command = "HELO" }
				}
			}
			if e.Spec != "" { continue }
			if sisimoji.Aligned(e.Diagnosis, []string{"host ", " said:"}) { e.Spec = "SMTP" }
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

