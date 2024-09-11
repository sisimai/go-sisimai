// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//       _               _      ___        _   _           _                 _            
//  _ __| |__   ___  ___| |_   / / \   ___| |_(_)_   _____| |__  _   _ _ __ | |_ ___ _ __ 
// | '__| '_ \ / _ \/ __| __| / / _ \ / __| __| \ \ / / _ \ '_ \| | | | '_ \| __/ _ \ '__|
// | |  | | | | (_) \__ \ |_ / / ___ \ (__| |_| |\ V /  __/ | | | |_| | | | | ||  __/ |   
// |_|  |_| |_|\___/|___/\__/_/_/   \_\___|\__|_| \_/ \___|_| |_|\__,_|_| |_|\__\___|_|   
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from QUALITIA Active!hunter
	InquireFor["Activehunter"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		// From: MAILER-DAEMON
		// Subject: FAILURE NOTICE :
		if len(bf.Head["x-ahmailid"]) == 0 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{
			"message": []string{"  ----- The following addresses had permanent fatal errors -----"},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		recipients := uint8(0)     // The number of 'Final-Recipient' header
		readcursor := uint8(0)     // Points the current cursor position
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
			if len(e)                                    == 0 { continue }

			//  ----- The following addresses had permanent fatal errors -----
			//
			// >>> kijitora@example.org <kijitora@example.org>
			//
			//  ----- Transcript of session follows -----
			// 550 sorry, no mailbox here by that name (#5.1.1 - chkusr)
			if strings.HasPrefix(e, ">>> ") && strings.Contains(e, "@") {
				// >>> kijitora@example.org <kijitora@example.org>
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e[0:strings.Index(e, ">")])
				recipients += 1

			} else {
				//  ----- Transcript of session follows -----
				// 550 sorry, no mailbox here by that name (#5.1.1 - chkusr)
				cr := []rune(e[0:1])
				if cr[0] < 48 || cr[0] > 122 { continue } // 48 = '0', 122 = 'z'
				if len(v.Diagnosis) > 0      { continue }
				v.Diagnosis = e
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents { 
			// Remove leading or/and trailing spaces, redandant spaces from the error messaage
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

