// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      _______          _                            ____   ___   ___ _____ 
// | | |__   ___  ___| |_   / / ____|_  _____| |__   __ _ _ __   __ _  ___|___ \ / _ \ / _ \___ / 
// | | '_ \ / _ \/ __| __| / /|  _| \ \/ / __| '_ \ / _` | '_ \ / _` |/ _ \ __) | | | | | | ||_ \ 
// | | | | | (_) \__ \ |_ / / | |___ >  < (__| | | | (_| | | | | (_| |  __// __/| |_| | |_| |__) |
// |_|_| |_|\___/|___/\__/_/  |_____/_/\_\___|_| |_|\__,_|_| |_|\__, |\___|_____|\___/ \___/____/ 
//                                                              |___/                             

package lhost
import "fmt"
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/smtp/status"
import sisimoji "libsisimai.org/sisimai/string"
import sisiaddr "libsisimai.org/sisimai/address"

func init() {
	// Decode bounce messages from Microsoft Exchange Server 2003: https://www.microsoft.com/microsoft-365/exchange/email
	InquireFor["Exchange2003"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		// X-MS-TNEF-Correlator: <00000000000000000000000000000000000000@example.com>
		// X-Mailer: Internet Mail Service (5.5.1960.3)
		// X-MS-Embedded-Report:
		proceedsto := false
		if len(bf.Headers["x-ms-embedded-report"]) > 0 { proceedsto = true }
		for proceedsto == false {
			// Check X-Mailer, X-MimeOLE, and Received headers
			tryto := []string{
				"Internet Mail Service (",                           // X-Mailer:
				"Microsoft Exchange Server Internet Mail Connector", // X-Mailer:
				"Produced By Microsoft Exchange",                    // X-MimeOLE:
				" with Internet Mail Service (",                     // Received:
			}
			if len(bf.Headers["x-mailer"]) > 0 {
				// X-Mailer:  Microsoft Exchange Server Internet Mail Connector Version 4.0.994.63
				// X-Mailer: Internet Mail Service (5.5.2232.9)
				if strings.HasPrefix(bf.Headers["x-mailer"][0], tryto[0]) { proceedsto = true; break }
				if strings.HasPrefix(bf.Headers["x-mailer"][0], tryto[1]) { proceedsto = true; break }
			}

			if len(bf.Headers["x-mimeole"]) > 0 {
				// X-MimeOLE: Produced By Microsoft Exchange V6.5
				if strings.HasPrefix(bf.Headers["x-mimeole"][0],tryto[2]) { proceedsto = true; break }
			}

			for _, e := range bf.Headers["received"] {
				// Received: by ***.**.** with Internet Mail Service (5.5.2657.72)
				if strings.Contains(e, tryto[3]) { proceedsto = true; break }
			}
			break
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{
			"message": []string{"Your message"},
			"error":   []string{"did not reach the following recipient(s):"},
		}
		errorcodes := map[string][]string{
			"userunknown":  []string{"000C05A6", /* Unknown Recipient   */ },
			"filtered":     []string{"000C0595", /* Ambiguous Recipient */ },
			"systemerror":  []string{
				"00010256", // Too many recipients.
				"000D06B5", // No proxy for recipient (non-smtp mail?)
			},
			"networkerror": []string{
				"00120270", // Too Many Hops
			},
			"contenterror": []string{
				"00050311", // Conversion to Internet format failed
				"000502CC", // Conversion to Internet format failed
			},
			"securityerror":[]string{
				"000B0981", // 502 Server does not support AUTH
			},
			"onhold":       []string{
				"000B099C", // Host Unknown, Message exceeds size limit, ...
				"000B09AA", // Unable to relay for, Message exceeds size limit,...
				"000B09B6", // Error messages by remote MTA
			},
		}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)              // Points the current cursor position
		recipients := uint8(0)              // The number of 'Final-Recipient' header
		statuspart := false                 // Flag, true if it has read the delivery status part
		connvalues := 0                     // Counter, 3 if it has got the all values of connheader
		connheader := [3]string{"", "", ""} // [To:, Subject:, Date:]
		rightindex := uint8(0)              // The last index number of dscontents
		anotherone := []string{""}          // Keeping another error messages
		msexchange := []bool{false}         // Flag, true if "MSEXCH:" text has been appeared
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 || statuspart == true { continue }

			if connvalues == len(connheader) {
				// did not reach the following recipient(s):
				//
				// kijitora@example.co.jp on Thu, 29 Apr 2007 16:51:51 -0500
				//     The recipient name is not recognized
				//     The MTS-ID of the original message is: c=jp;a= ;p=neko
				// ;l=EXCHANGE000000000000000000
				//     MSEXCH:IMS:KIJITORA CAT:EXAMPLE:EXCHANGE 0 (000C05A6) Unknown Recipient
				// mikeneko@example.co.jp on Thu, 29 Apr 2007 16:51:51 -0500
				//     The recipient name is not recognized
				//     The MTS-ID of the original message is: c=jp;a= ;p=neko
				// ;l=EXCHANGE000000000000000000
				//     MSEXCH:IMS:KIJITORA CAT:EXAMPLE:EXCHANGE 0 (000C05A6) Unknown Recipient
				if sisimoji.Aligned(e, []string{"@", " on "}) {
					// kijitora@example.co.jp on Thu, 29 Apr 2007 16:51:51 -0500
					//   kijitora@example.com on 4/29/99 9:19:59 AM
					if len(v.Recipient) > 0 {
						// There are multiple recipient addresses in the message body.
						dscontents = append(dscontents, sis.DeliveryMatter{})
						anotherone = append(anotherone, "")
						msexchange = append(msexchange, false)
						rightindex++
						v = &(dscontents[rightindex])
					}
					p1 := strings.Index(strings.ToLower(e), "smtp="); if p1 < 0 { p1 = 0 } else { p1 += 5 }
					p2 := strings.Index(e, " on ")

					v.Recipient = sisiaddr.S3S4(e[p1:p2])
					recipients += 1
					msexchange[rightindex] = false

				} else if strings.HasPrefix(e, " ") && strings.Contains(e, "MSEXCH:") {
					//     MSEXCH:IMS:KIJITORA CAT:EXAMPLE:EXCHANGE 0 (000C05A6) Unknown Recipient
					v.Diagnosis += e[strings.Index(e, "MSEXCH:"):]

				} else {
					if msexchange[rightindex] == true { continue }
					if strings.HasPrefix(v.Diagnosis, "MSEXCH:") {
						// continued from the previous line starting with "MSEXCH:"
						msexchange[rightindex] = true
						statuspart   = true
						v.Diagnosis += " " + e

					} else {
						// Error messages in the body part
						anotherone[rightindex] += " " + e
					}
				}
			} else {
				// Your message
				//
				//  To:      shironeko@example.jp
				//  Subject: ...
				//  Sent:    Thu, 29 Apr 2010 18:14:35 +0000
				//
				if strings.HasPrefix(e, "  To:  ") || strings.HasPrefix(e, "      To: ") {
					//  To:      shironeko@example.jp
					if connheader[0] != "" { continue }
					connheader[0] = strings.Trim(e[strings.Index(e, "To: ") + 4:], " ")
					connvalues++

				} else if strings.HasPrefix(e, "      Subject: ") || strings.HasPrefix(e, "  Subject: ") {
					//  Subject: ...
					if connheader[1] != "" { continue }
					connheader[1] = strings.Trim(e[strings.Index(e, "Subject: ") + 9:], " ")
					connvalues++

				} else if strings.HasPrefix(e, "  Sent: ") || strings.HasPrefix(e, "      Sent: ") {
					//  Sent:    Thu, 29 Apr 2010 18:14:35 +0000
					//  Sent:    4/29/99 9:19:59 AM
					if connheader[2] != "" { continue }
					connheader[2] = strings.Trim(e[strings.Index(e, "Sent: ") + 6:], " ")
					connvalues++
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			if sisimoji.Aligned(e.Diagnosis, []string{"MSEXCH:", "(", ")"}) {
				//     MSEXCH:IMS:KIJITORA CAT:EXAMPLE:EXCHANGE 0 (000C05A6) Unknown Recipient
				capturedcode := sisimoji.Select(e.Diagnosis, "(", ")", 0)
				errormessage := e.Diagnosis[strings.Index(e.Diagnosis, ")") + 1:]

				FINDREASON: for r := range errorcodes {
					// The key name is a bounce reason name
					for _, f := range errorcodes[r] {
						// Find the captured code from the error code table
						if capturedcode != f { continue }
						e.Reason = r
						e.Status = status.Code(r, false)
						break FINDREASON
					}
				}
				e.Diagnosis = errormessage
			}

			// Could not detect the reason from the value of "diagnosis", copy alternative error message 
			if e.Reason != "" || anotherone[j] == "" { continue }
			e.Diagnosis = anotherone[j] + " " + e.Diagnosis
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
		}

		if emailparts[1] == "" {
			// When original message is not included in the bounce message
			emailparts[1] += fmt.Sprintf("From: %s\n", connheader[0])
			emailparts[1] += fmt.Sprintf("Subject: %s\n", connheader[2])
			emailparts[1] += fmt.Sprintf("Date: %s\n", connheader[1])
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

