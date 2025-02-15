// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      ____  __       _ _ __  __                _           _ ____  __  __ _____ ____  
// | | |__   ___  ___| |_   / /  \/  | __ _(_) |  \/  | __ _ _ __ ___| |__   __ _| / ___||  \/  |_   _|  _ \ 
// | | '_ \ / _ \/ __| __| / /| |\/| |/ _` | | | |\/| |/ _` | '__/ __| '_ \ / _` | \___ \| |\/| | | | | |_) |
// | | | | | (_) \__ \ |_ / / | |  | | (_| | | | |  | | (_| | |  \__ \ | | | (_| | |___) | |  | | | | |  __/ 
// |_|_| |_|\___/|___/\__/_/  |_|  |_|\__,_|_|_|_|  |_|\__,_|_|  |___/_| |_|\__,_|_|____/|_|  |_| |_| |_|    

package lhost
import "fmt"
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc2045"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from Trustwave Secure Email Gateway: https://www.trustwave.com/en-us/services/email-security/
	InquireFor["MailMarshalSMTP"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }
		if strings.HasPrefix(bf.Headers["subject"][0], `Undeliverable Mail: "`) == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"'+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++"}
		if cv := rfc2045.Boundary(bf.Headers["content-type"][0], 1); cv != "" { boundaries = append(boundaries, cv) }

		startingof := map[string][]string{
			"message": []string{"Your message:"},
			"error":   []string{"Could not be delivered because of"},
			"rcpts":   []string{"The following recipients were affected:"},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		endoferror := false               // Flag for the end of error messages
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }

			// Your message:
			//    From:    originalsender@example.com
			//    Subject: ...
			//
			// Could not be delivered because of
			//
			// 550 5.1.1 User unknown
			//
			// The following recipients were affected:
			//    neko@example.com
			if strings.HasPrefix(e, "    ") && strings.IndexByte(e, '@') > 0 {
				// The following recipients were affected:
				//    neko@example.com
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = e[4:]
				recipients += 1

			} else {
				// Get error message lines
				if e == startingof["error"][0] {
					// Could not be delivered because of
					//
					// 550 5.1.1 User unknown
					v.Diagnosis = e

				} else if v.Diagnosis != "" && endoferror == false {
					// Append error mesages to v.Diagnosis
					if strings.HasPrefix(e, startingof["rcpts"][0]) { endoferror = true; continue }
					v.Diagnosis += " " + e

				} else {
					// Additional Information
					// ======================
					// Original Sender:    <originalsender@example.com>
					// Sender-MTA:         <10.11.12.13>
					// Remote-MTA:         <10.0.0.1>
					// Reporting-MTA:      <relay.xxxxxxxxxxxx.com>
					// MessageName:        <B549996730000.000000000001.0003.mml>
					// Last-Attempt-Date:  <16:21:07 seg, 22 Dezembro 2014>
					if strings.HasPrefix(e, "Original Sender: ") {
						// Original Sender:    <originalsender@example.com>
						// Use this line instead of "From" header of the original message.
						emailparts[1] += fmt.Sprintf("From: %s\n", sisimoji.Select(e, "<", ">", 0))

					} else if strings.HasPrefix(e, "Sender-MTA: ") {
						// Sender-MTA:         <10.11.12.13>
						v.Lhost = sisimoji.Select(e, "<", ">", 0)

					} else if strings.HasPrefix(e, "Reporting-MTA: ") {
						// Reporting-MTA:      <relay.xxxxxxxxxxxx.com>
						v.Rhost = sisimoji.Select(e, "<", ">", 0)

					} else if strings.Contains(e, " From:") || strings.Contains(e, " Subject:") {
						//    From:    originalsender@example.com
						//    Subject: ...
						p1 := strings.Index(e, " From:"); if p1 < 0 { p1 = strings.Index(e, " Subject:") }
						p2 := strings.IndexByte(e, ':')
						emailparts[1] += fmt.Sprintf("%s: %s\n", e[p1 + 1:p2], sisimoji.Sweep(e[p2 + 1:]))
					}
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

