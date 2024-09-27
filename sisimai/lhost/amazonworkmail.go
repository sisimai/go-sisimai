// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//       _               _      ___                                __        __         _    __  __       _ _ 
//  _ __| |__   ___  ___| |_   / / \   _ __ ___   __ _ _______  _ _\ \      / /__  _ __| | _|  \/  | __ _(_) |
// | '__| '_ \ / _ \/ __| __| / / _ \ | '_ ` _ \ / _` |_  / _ \| '_ \ \ /\ / / _ \| '__| |/ / |\/| |/ _` | | |
// | |  | | | | (_) \__ \ |_ / / ___ \| | | | | | (_| |/ / (_) | | | \ V  V / (_) | |  |   <| |  | | (_| | | |
// |_|  |_| |_|\___/|___/\__/_/_/   \_\_| |_| |_|\__,_/___\___/|_| |_|\_/\_/ \___/|_|  |_|\_\_|  |_|\__,_|_|_|
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from Amazon WorkMail: https://aws.amazon.com/workmail/
	InquireFor["AmazonWorkMail"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		// X-Mailer: Amazon WorkMail
		// X-Original-Mailer: Amazon WorkMail
		// X-Ses-Outgoing: 2016.01.14-54.240.27.159
		proceedsto := uint8(0)
		mailername := ""
		if len(bf.Head["x-original-mailer"])            > 0 { mailername = bf.Head["x-original-mailer"][0] }
		if mailername != "" && len(bf.Head["x-mailer"]) > 0 { mailername = bf.Head["x-mailer"][0]          }
		if len(bf.Head["x-ses-outgoing"]) > 0 { proceedsto++ }
		if mailername == "Amazon WorkMail"    { proceedsto++ }
		if proceedsto < 2 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"}
		startingof := map[string][]string{
			"message": []string{"Technical report:"},
		}
		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
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
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

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
					// # Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
					v.Spec = o[1]
					v.Diagnosis = o[2]

				} else {
					// Other DSN fields defined in RFC3464
					v.Set(o[0], o[2]); if f != 1 { continue }

					// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
					permessage[z] = o[2]
					if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }
				}
			}

			// <!DOCTYPE HTML><html>
			// <head>
			// <meta name="Generator" content="Amazon WorkMail v3.0-2023.77">
			// <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
			if strings.HasPrefix(e, "<!DOCTYPE HTML><html>") { break }
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Set default values stored in "permessage" if each value in "dscontents" is empty.
			e := &(dscontents[j])
			for _, z := range keystrings {
				// Do not set an empty string into each member of DeliveryMatter{}
				if len(v.Get(z))       > 0 { continue }
				if len(permessage[z]) == 0 { continue }
				e.Set(z, permessage[z])
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			if strings.Index(e.Status, ".0.0") > 0 || strings.Index(e.Status, ".1.0") > 0 {
				// Get other status value from the error message
				// 5.1.0 - Unknown address error 550-'5.7.1 ...
				ce := status.Find(e.Diagnosis, "")
				if ce != "" { e.Status = ce }
			}

			// 554 4.4.7 Message expired: unable to deliver in 840 minutes.
			// <421 4.4.2 Connection timed out>
			e.ReplyCode = reply.Find(e.Diagnosis, e.Status)
			if e.Reason == "" { e.Reason = status.Name(e.Status) }
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

