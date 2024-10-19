// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ____  __       _ _ _____                     _            
// | | |__   ___  ___| |_   / /  \/  | __ _(_) |  ___|__  _   _ _ __   __| |_ __ _   _ 
// | | '_ \ / _ \/ __| __| / /| |\/| |/ _` | | | |_ / _ \| | | | '_ \ / _` | '__| | | |
// | | | | | (_) \__ \ |_ / / | |  | | (_| | | |  _| (_) | |_| | | | | (_| | |  | |_| |
// |_|_| |_|\___/|___/\__/_/  |_|  |_|\__,_|_|_|_|  \___/ \__,_|_| |_|\__,_|_|   \__, |
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from MailFoundry: https://www.barracuda.com/
	InquireFor["MailFoundry"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		proceedsto := false
		ISMF: for {
			// Subject: Message delivery has failed
			if bf.Head["subject"][0] != "Message delivery has failed" { break ISMF }
			for _, e := range bf.Head["received"] {
				// Received: From localhost (127.0.0.1) by smtp9.mf.example.ne.jp (MAILFOUNDRY) id ...
				if strings.Contains(e, "(MAILFOUNDRY) id") { proceedsto = true; break ISMF }
			}
			break ISMF
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{
			"message": []string{"Unable to deliver message to:"},
			"error":   []string{"Delivery failed for the following reason:"},
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
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// Unable to deliver message to: <kijitora@example.org>
			// Delivery failed for the following reason:
			// Server mx22.example.org[192.0.2.222] failed with: 550 <kijitora@example.org> No such user here
			//
			// This has been a permanent failure.  No further delivery attempts will be made.
			if strings.HasPrefix(e, "Unable to deliver message to: <") && strings.Contains(e, "@") {
				// Unable to deliver message to: <kijitora@example.org>
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e[strings.Index(e, "<"):])
				recipients += 1

			} else {
				// Error messages
				if e == startingof["error"][0] { v.Diagnosis = e; continue }
				if v.Diagnosis == ""         { continue }
				if strings.HasPrefix(e, "-") { continue }
				v.Diagnosis += " " + e
			}

		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Set default values stored in "permessage" if each value in "dscontents" is empty.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

