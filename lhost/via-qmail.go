// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      __                    _ _ 
// | | |__   ___  ___| |_   / /_ _ _ __ ___   __ _(_) |
// | | '_ \ / _ \/ __| __| / / _` | '_ ` _ \ / _` | | |
// | | | | | (_) \__ \ |_ / / (_| | | | | | | (_| | | |
// |_|_| |_|\___/|___/\__/_/ \__, |_| |_| |_|\__,_|_|_|
//                              |_|                    

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/smtp/command"
import sisimoji "libsisimai.org/sisimai/string"
import sisiaddr "libsisimai.org/sisimai/address"

func init() {
	// Decode bounce messages from qmail: https://cr.yp.to/qmail.html
	InquireFor["qmail"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		// Pre process email headers and the body part of the message which generated by qmail.
		// see https://cr.yp.to/qmail.html
		//   e.g.) Received: (qmail 12345 invoked for bounce); 29 Apr 2009 12:34:56 -0000
		//         Subject: failure notice
		proceedsto := false
		relayedvia := [][]string{
			[]string{"(qmail ", "invoked for bounce)"},
			[]string{"(qmail ", "invoked from ", "network)"},
		}
		emailtitle := []string{
			"failure notice", // qmail-send.c:Subject: failure notice\n\
			"Failure Notice", // Yahoo
		}
		if sisimoji.EqualsAny(bf.Headers["subject"][0], emailtitle) { proceedsto = true }
		for _, e := range bf.Headers["received"] {
			// Received: (qmail 2222 invoked for bounce);29 Apr 2017 23:34:45 +0900
			// Received: (qmail 2202 invoked from network); 29 Apr 2018 00:00:00 +0900
			if proceedsto == true { break }
			if sisimoji.Aligned(e, relayedvia[0]) { proceedsto = true }
			if sisimoji.Aligned(e, relayedvia[1]) { proceedsto = true }
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{
			// qmail-send.c:qmail_puts(&qqt,*sender.s ? "--- Below this line is a copy of the message.\n\n" :...
			"--- Below this line is a copy of the message.",     // qmail-1.03
			"--- Below this line is a copy of the mail header.",
			"--- Below the next line is a copy of the message.", // The followings are the qmail clone
			"--- Mensaje original adjunto.",
			"Content-Type: message/rfc822",
			"Original message follows.",
		}
		startingof := map[string][]string{
			//  qmail-remote.c:248|    if (code >= 500) {
			//  qmail-remote.c:249|      out("h"); outhost(); out(" does not like recipient.\n");
			//  qmail-remote.c:265|  if (code >= 500) quit("D"," failed on DATA command");
			//  qmail-remote.c:271|  if (code >= 500) quit("D"," failed after I sent the message");
			//
			// Characters: K,Z,D in qmail-qmqpc.c, qmail-send.c, qmail-rspawn.c
			//  K = success, Z = temporary error, D = permanent error
			"error":   []string{"Remote host said:"},
			"message": []string{
				"Hi. This is the qmail", // qmail-send.c:Hi. This is the qmail-send program at ");
				"He/Her is not ",        // The followings are the qmail clone
				"unable to deliver your message to the following addresses",
				"Su mensaje no pudo ser entregado",
				"Sorry, we were unable to deliver your message to the following address",
				"This is the machine generated message from mail service",
				"This is the mail delivery agent at",
				"Unable to deliver message to the following address",
				"unable to deliver your message to the following addresses",
				"Unfortunately, your mail was not delivered to the following address:",
				"Your mail message to the following address",
				"Your message to the following addresses",
				"We're sorry.",
			},
			"rhost":   []string{"Giving up on ", "Connected to ", "remote host "},
		}
		commandset := map[string][]string{
			// Error text regular expressions which defined in qmail-remote.c
			// qmail-remote.c:225|  if (smtpcode() != 220) quit("ZConnected to "," but greeting failed");
			"CONN": []string{" but greeting failed."},

			// qmail-remote.c:231|  if (smtpcode() != 250) quit("ZConnected to "," but my name was rejected");
			"EHLO": []string{" but my name was rejected."},

			// qmail-remote.c:238|  if (code >= 500) quit("DConnected to "," but sender was rejected");
			// reason = rejected
			"MAIL": []string{" but sender was rejected."},

			// qmail-remote.c:249|  out("h"); outhost(); out(" does not like recipient.\n");
			// qmail-remote.c:253|  out("s"); outhost(); out(" does not like recipient.\n");
			// reason = userunknown
			"RCPT": []string{" does not like recipient."},

			// qmail-remote.c:265|  if (code >= 500) quit("D"," failed on DATA command");
			// qmail-remote.c:266|  if (code >= 400) quit("Z"," failed on DATA command");
			// qmail-remote.c:271|  if (code >= 500) quit("D"," failed after I sent the message");
			// qmail-remote.c:272|  if (code >= 400) quit("Z"," failed after I sent the message");
			"DATA": []string{" failed on DATA command", " failed after I sent the message"},
		}

		// qmail-send.c:922| ... (&dline[c],"I'm not going to try again; this message has been in the queue too long.\n")) nomem();
		// qmail-remote-fallback.patch
		hasexpired := "this message has been in the queue too long."
		onholdpair := []string{" does not like recipient.", "this message has been in the queue too long."}
		failonldap := map[string][]string{
			// qmail-ldap-1.03-20040101.patch:19817 - 19866
			"exceedlimit": []string{"The message exeeded the maximum size the user accepts"}, // 5.2.3
			"suspend":     []string{
				"Mailaddress is administrativly disabled",
				"Mailaddress is administrativley disabled",
				"Mailaddress is administratively disabled",
				"Mailaddress is administrativeley disabled",
			},  // 5.2.1
			"systemerror": []string{
				"Automatic homedir creator crashed",                // 4.3.0
				"Illegal value in LDAP attribute",                  // 5.3.5
				"LDAP attribute is not given but mandatory",        // 5.3.5
				"Timeout while performing search on LDAP server",   // 4.4.3
				"Too many results returned but needs to be unique", // 5.3.5
				"Permanent error while executing qmail-forward",    // 5.4.4
				"Temporary error in automatic homedir creation",    // 4.3.0 or 5.3.0
				"Temporary error while executing qmail-forward",    // 4.4.4
				"Temporary failure in LDAP lookup",                 // 4.4.3
				"Unable to contact LDAP server",                    // 4.4.3
				"Unable to login into LDAP server, bad credentials",// 4.4.3
			},
			"userunknown": []string{"Sorry, no mailbox here by that name"}, // 5.1.1
		}
		messagesof := map[string][]string{
			// qmail-local.c:589|  strerr_die1x(100,"Sorry, no mailbox here by that name. (#5.1.1)");
			// qmail-remote.c:253|  out("s"); outhost(); out(" does not like recipient.\n");
			"hostunknown": []string{"Sorry, I couldn't find any host "},
			// error_str.c:192|  X(EDQUOT,"disk quota exceeded")
			"mailboxfull": []string{"disk quota exceeded"},
			// qmail-qmtpd.c:233| ... result = "Dsorry, that message size exceeds my databytes limit (#5.3.4)";
			// qmail-smtpd.c:391| ... out("552 sorry, that message size exceeds my databytes limit (#5.3.4)\r\n"); return;
			"mesgtoobig":  []string{"Message size exceeds fixed maximum message size:"},
			// qmail-remote.c:68|  Sorry, I couldn't find any host by that name. (#4.1.2)\n"); zerodie();
			// qmail-remote.c:78|  Sorry, I couldn't find any host named ");
			"networkerror": []string{
				"Sorry, I wasn't able to establish an SMTP connection",
				"Sorry. Although I'm listed as a best-preference MX or A for that host",
			},
			"notaccept":   []string{
				// notqmail 1.08 returns the following error message when the destination MX is NullMX
				"Sorry, I couldn't find a mail exchanger or IP address",
			},
			"systemerror": []string{
				"bad interpreter: No such file or directory",
				"system error",
				"Unable to",
			},
			"systemfull":  []string{"Requested action not taken: mailbox unavailable (not enough free space)"},
			"userunknown": []string{"no mailbox here by that name"},
		}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		anotherone := []string{""}        // Keeping another error messages
		rightindex := uint8(0)            // The last index number of dscontents
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if sisimoji.ContainsAny(e, startingof["message"]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// <kijitora@example.jp>:
			// 192.0.2.153 does not like recipient.
			// Remote host said: 550 5.1.1 <kijitora@example.jp>... User Unknown
			// Giving up on 192.0.2.153.
			if strings.HasPrefix(e, "<") && sisimoji.Aligned(e, []string{"<", "@", ">:"}) {
				// <kijitora@example.jp>:
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					anotherone = append(anotherone, "")
					rightindex++
					v = &(dscontents[rightindex])
				}
				v.Recipient = sisiaddr.S3S4(e[1:strings.Index(e, ">:")])
				recipients += 1

			} else if len(dscontents) == int(recipients) {
				// Get error messages
				v.Diagnosis += e + " "
				if strings.HasPrefix(e, startingof["error"][0]) { anotherone[rightindex] = e }

				if v.Rhost != "" { continue }
				for _, r := range startingof["rhost"] {
					// Connected to 192.0.2.225 but sender was rejected.
					// Connected to 192.0.2.112 but my name was rejected.
					// Giving up on 192.0.2.135.
					// remote host 203.138.180.112 said:...
					p1 := strings.Index(e, r); if p1 < 0 { continue }
					cm := len(r)
					p2 := sisimoji.IndexOnTheWay(e, " ", p1 + cm + 1); if p2 < 0 { p2 = strings.LastIndex(e, ".") }
					v.Rhost = sisimoji.Sweep(e[p1 + cm:p2])
					break
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			for r := range commandset {
				// Get the last SMTP command
				if sisimoji.ContainsAny(e.Diagnosis, commandset[r]) { e.Command = r; break }
			}
			if e.Command == "" && strings.Contains(e.Diagnosis, "no SMTP connection got far enough") {
				// Sorry, no SMTP connection got far enough; most progress was RCPT TO response; ...
				e.Command = command.Find(e.Diagnosis)
			}

			if e.Command == "EHLO" || e.Command == "HELO" {
				// HELO | Connected to 192.0.2.135 but my name was rejected.
				e.Reason = "blocked"

			} else {
				// The error message includes any of patterns defined in the variable avobe
				if sisimoji.Aligned(e.Diagnosis, onholdpair) {
					// Need to be matched with error message pattens defined in sisimai/reason/*
					e.Reason = "onhold"

				} else {
					FINDREASON: for _, f := range []string{anotherone[j], e.Diagnosis} {
						// Check that the error message includes any of message patterns or not
						if e.Reason != "" { break    }
						if f == ""        { continue }
						for r := range messagesof {
							// The key name is a bounce reason name
							if sisimoji.ContainsAny(f, messagesof[r]) == false { continue }
							e.Reason = r
							break FINDREASON
						}

						for r := range failonldap {
							// The key name is a bounce reason name
							if sisimoji.ContainsAny(f, failonldap[r]) == false { continue }
							e.Reason = r
							break FINDREASON
						}
						if strings.Contains(f, hasexpired) { e.Reason = "expired" }
					}
				}
			}
			if e.Command == "" { e.Command = command.Find(e.Diagnosis) }
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

