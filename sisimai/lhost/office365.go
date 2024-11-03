// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      _____   __  __ _          _____  __  ____  
// | | |__   ___  ___| |_   / / _ \ / _|/ _(_) ___ ___|___ / / /_| ___| 
// | | '_ \ / _ \/ __| __| / / | | | |_| |_| |/ __/ _ \ |_ \| '_ \___ \ 
// | | | | | (_) \__ \ |_ / /| |_| |  _|  _| | (_|  __/___) | (_) |__) |
// |_|_| |_|\___/|___/\__/_/  \___/|_| |_| |_|\___\___|____/ \___/____/ 
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/status"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Microsoft 365: https://office.microsoft.com/
	InquireFor["Office365"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		// X-MS-Exchange-Message-Is-Ndr:
		// X-Microsoft-Antispam-PRVS: <....@...outlook.com>
		// X-Exchange-Antispam-Report-Test: UriScan:;
		// X-Exchange-Antispam-Report-CFA-Test:
		// X-MS-Exchange-CrossTenant-OriginalArrivalTime: 29 Apr 2015 23:34:45.6789 (JST)
		// X-MS-Exchange-CrossTenant-FromEntityHeader: Hosted
		// X-MS-Exchange-Transport-CrossTenantHeadersStamped: ...
		proceedsto := uint8(0)
		relayhosts := []string{".outbound.protection.outlook.com", ".prod.outlook.com"}
		if strings.Contains(bf.Head["subject"][0], "Undeliverable:") ||
		   strings.Contains(bf.Head["subject"][0], "Onbestelbaar:")  ||
		   strings.Contains(bf.Head["subject"][0], "Não_entregue:")  { proceedsto++ }

		if len(bf.Head["x-ms-exchange-message-is-ndr"]) > 0 ||
		   len(bf.Head["x-microsoft-antispam-prvs"])    > 0 ||
		   len(bf.Head["x-exchange-antispam-report-test"])     > 0 ||
		   len(bf.Head["x-exchange-antispam-report-cfa-test"]) > 0 ||
		   len(bf.Head["x-ms-exchange-crosstenant-originalarrivaltime"])     > 0 ||
		   len(bf.Head["x-ms-exchange-crosstenant-fromentityheader"])        > 0 ||
		   len(bf.Head["x-ms-exchange-transport-crosstenantheadersstamped"]) > 0 { proceedsto++ }

		// Received: from aaaa00-bb0-ccc.outbound.protection.outlook.com ([192.0.2.1])
		//           by 192.0.2.4 with MailEnable ESMTP; Thu, 29 Apr 2015 23:34:45 +0900
		// Message-ID: <00000000-0000-0000-0000-000000000000@*.*.prod.outlook.com>
		for _, e := range bf.Head["received"] { if sisimoji.ContainsAny(e, relayhosts) { proceedsto++; break } }
		if sisimoji.ContainsAny(bf.Head["message-id"][0], relayhosts) == true          { proceedsto++          }
		if proceedsto < 2 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822", "Original message headers:"}
		commandset := map[string][]string{"RCPT": []string{"unknown recipient or mailbox unavailable ->", "@"}}
		startingof := map[string][]string{
			"eoe": []string{
				"Original message headers:", "Original Message Headers:",
				"Message Hops",
				"alhos originais da mensagem:",
				"Oorspronkelijke berichtkoppen:",
			},
			"error": []string{
				"Diagnostic information for administrators:",
				"Diagnostische gegevens voor beheerders:",
				"Error Details",
				"stico para administradores:",
			},
			"lhost": []string{
				"Generating server: ",
				"Bronserver: ",
				"Servidor de origem: ",
			},
			"message": []string{
				" rejected your message to the following e",
				"Delivery has failed to these recipients or groups:",
				"Falha na entrega a estes destinat",
				"Original Message Details",
				"Uw bericht kan niet worden bezorgd bij de volgende geadresseerden of groepen:",
			},
			"rfc3464": []string{"Content-Type: message/delivery-status"},
		}

		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		endoferror := false               // Flag for the end of error messages
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if sisimoji.ContainsAny(e, startingof["message"]) {readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// kijitora@example.com<mailto:kijitora@example.com>
			// The email address wasn't found at the destination domain. It might
			// be misspelled or it might not exist any longer. Try retyping the
			// address and resending the message.
			//
			// Original Message Details
			// Created Date:   4/29/2017 6:40:30 AM
			// Sender Address: neko@example.jp
			// Recipient Address:      kijitora@example.org
			// Subject:        Nyaan
			p1 := strings.Index(e, "<mailto:")
			p2 := strings.Index(e, "Recipient Address: ")
			if p1 > 1 || p2 == 0 {
				// kijitora@example.com<mailto:kijitora@example.com>
				// Recipient Address:      kijitora-nyaan@neko.kyoto.example.jp
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(strings.Trim(e[strings.Index(e, ":") + 1:], "< >"))
				recipients += 1

			} else if sisimoji.HasPrefixAny(e, startingof["lhost"]) {
				// Generating server: FFFFFFFFFFFF.e0.prod.outlook.com
				z := "reporting-mta"
				permessage[z] = e[strings.Index(e, ": ") + 2:]
				if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }

			} else {
				// Get error messages
				if endoferror {
					// After "Content-Type: message/delivery-status" line
					f := rfc1894.Match(e); if f      < 1 { continue }
					o := rfc1894.Field(e); if len(o) < 1 { continue }
					if len(fieldtable[o[0]]) == 0        { continue }

					z := fieldtable[o[0]]
					if v.Diagnosis == "" {
						// Capture "Diagnostic-Code:" field because no error messages have been captured
						v.Update(o[0], o[2])

					} else {
						// Do not capture "Diagnostic-Code:" field because error message have already
						// been captured
						if o[0] == "diagnostic-code" { continue }
						if o[0] == "final-recipient" { continue }
						v.Update(o[0], o[2]); if f != 1 { continue }
					}
					// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
					permessage[z] = o[2]
					if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }

				} else {
					// Before "Original message headers:" line
					// Diagnostic information for administrators:
					if sisimoji.ContainsAny(e, startingof["error"]) { v.Diagnosis = e; continue }

					// kijitora@example.com
					// Remote Server returned '550 5.1.10 RESOLVER.ADR.RecipientNotFound; Recipien=
					// t not found by SMTP address lookup'
					if v.Diagnosis != "" {
						// The error message text have already captured
						if sisimoji.ContainsAny(e, startingof["eoe"]) { 
							// Do not append error messages after "Content-Type: message/delivery-status" line
							endoferror = true
							continue
						}
						v.Diagnosis += " " + e

					} else {
						// The error message text has not been captured yet
						if strings.HasPrefix(e, startingof["rfc3464"][0]) { endoferror = true }
					}
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Set default values stored in "permessage" if each value in "dscontents" is empty.
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			for _, z := range keystrings {
				// Do not set an empty string into each member of DeliveryMatter{}
				if len(v.Select(z))    > 0 { continue }
				if len(permessage[z]) == 0 { continue }
				e.Update(z, permessage[z])
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			if e.Status == "" || strings.HasSuffix(e.Status, ".0.0") {
				// There is no value of Status header or the value is 5.0.0, 4.0.0
				if cv := status.Find(e.Diagnosis, e.ReplyCode); status.Test(cv) { e.Status = cv }
			}
			if sisimoji.Aligned(e.Diagnosis, commandset["RCPT"]) { e.Command = "RCPT" }
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

