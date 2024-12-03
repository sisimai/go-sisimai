// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ______                   _    __        __         _                             
// | | |__   ___  ___| |_   / / ___| ___   ___   __ _| | __\ \      / /__  _ __| | _____ _ __   __ _  ___ ___ 
// | | '_ \ / _ \/ __| __| / / |  _ / _ \ / _ \ / _` | |/ _ \ \ /\ / / _ \| '__| |/ / __| '_ \ / _` |/ __/ _ \
// | | | | | (_) \__ \ |_ / /| |_| | (_) | (_) | (_| | |  __/\ V  V / (_) | |  |   <\__ \ |_) | (_| | (_|  __/
// |_|_| |_|\___/|___/\__/_/  \____|\___/ \___/ \__, |_|\___| \_/\_/ \___/|_|  |_|\_\___/ .__/ \__,_|\___\___|
//                                              |___/                                   |_|                   
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Google Workspace: https://workspace.google.com/
	InquireFor["GoogleWorkspace"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }
		if strings.Contains(bf.Head["from"][0], "<mailer-daemon@googlemail.com>")  == false { return sis.RisingUnderway{} }
		if strings.Contains(bf.Head["subject"][0], "Delivery Status Notification") == false { return sis.RisingUnderway{} }
		if strings.Contains(bf.Body, "\nDiagnostic-Code:")                         == true  { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"}
		startingof := map[string][]string{
			"message": []string{"** "},
			"error":   []string{"The response was:", "The response from the remote server was:"},
		}
		messagesof := map[string][]string{
			"networkerror": []string{" had no relevant answers.", " responded with code NXDOMAIN"},
			"notaccept":    []string{"Null MX"},
			"userunknown":  []string{"because the address couldn't be found. Check for typos or unnecessary spaces and try again."},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
				v.Diagnosis = e + " "
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// ** Message not delivered **
			// You're sending this from a different address or alias using the 'Send mail as' feature.
			// The settings for your 'Send mail as' account are misconfigured or out of date. Check those settings and try resending.
			// Learn more here: https://support.google.com/mail/?p=CustomFromDenied
			// The response was:
			// Unspecified Error (SENT_SECOND_EHLO): Smtp server does not advertise AUTH capability
			if strings.HasPrefix(e, "Content-Type:") == true { continue }
			v.Diagnosis += e + " "
		}

		for recipients == 0 {
			// Pick the recipient address from the value of To: header of the original message after
			// Content-Type: message/rfc822 field
			p0 := strings.Index(emailparts[1], "\nTo:"); if p0 < 0 { break }
			p1 := sisimoji.IndexOnTheWay(emailparts[1], "\n", p0 + 2)
			dscontents[0].Recipient = sisiaddr.S3S4(emailparts[1][p0 + 4:p1])
			recipients++
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(strings.ReplaceAll(e.Diagnosis, "\n", " "))

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

