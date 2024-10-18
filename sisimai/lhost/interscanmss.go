// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//       _               _      _____       _            ____                  __  __ ____ ____  
//  _ __| |__   ___  ___| |_   / /_ _|_ __ | |_ ___ _ __/ ___|  ___ __ _ _ __ |  \/  / ___/ ___| 
// | '__| '_ \ / _ \/ __| __| / / | || '_ \| __/ _ \ '__\___ \ / __/ _` | '_ \| |\/| \___ \___ \ 
// | |  | | | | (_) \__ \ |_ / /  | || | | | ||  __/ |   ___) | (_| (_| | | | | |  | |___) |__) |
// |_|  |_| |_|\___/|___/\__/_/  |___|_| |_|\__\___|_|  |____/ \___\__,_|_| |_|_|  |_|____/____/ 
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Trend Micro InterScan Messaging Security Suite
	// https://www.trendmicro.com/en_us/business/products/user-protection/sps/email-and-collaboration/interscan-messaging.html
	InquireFor["InterScanMSS"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		proceedsto := false
		emailtitle := bf.Head["subject"][0]
		titletable := []string{
			"Mail could not be delivered",
			"メッセージを配信できません。",
			"メール配信に失敗しました",
		}

		if strings.HasPrefix(bf.Head["from"][0], `"InterScan`) { proceedsto = true }
		if sisimoji.ContainsAny(emailtitle, titletable)        { proceedsto = true }
		if proceedsto == false { return sis.RisingUnderway{} }

		boundaries := []string{"Content-Type: message/rfc822"}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if e == "" { continue }

			p1 := strings.Index(e, " <<< ") // Sent <<< ...
			p2 := strings.Index(e, " >>> ") // Received >>> ...
			if strings.Index(e, "@") > 1 && strings.Index(e, " <") > 1 && 
			   (p1 > 1 || p2 > 1 || strings.Contains(e, "Unable to deliver ")) {
				// Sent <<< RCPT TO:<kijitora@example.co.jp>
				// Received >>> 550 5.1.1 <kijitora@example.co.jp>... user unknown
				// Received >>> 550 5.1.1 unknown user.
				// Unable to deliver message to <kijitora@neko.example.jp>
				// Unable to deliver message to <neko@example.jp> (and other recipients in the same domain).
				p3 := strings.LastIndex(e, "<")
				p4 := strings.LastIndex(e, ">")
				cr := sisiaddr.Find(e[p3:p4 + 1])
				if len(cr) == 0 || sisiaddr.IsEmailAddress(cr[0]) == false { continue }

				if len(v.Recipient) > 0 && strings.Contains(cr[0], v.Recipient) == false {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				if strings.Contains(e, "Unable to deliver ") { v.Diagnosis = e }

				v.Recipient = sisiaddr.S3S4(cr[0])
				recipients  = uint8(len(dscontents))
			}

			if strings.HasPrefix(e, "Sent <<< ") {
				// Sent <<< RCPT TO:<kijitora@example.co.jp>
				v.Command = command.Find(e)

			} else if strings.HasPrefix(e, "Received >>> ") {
				// Received >>> 550 5.1.1 <kijitora@example.co.jp>... user unknown
				v.Diagnosis = e[strings.Index(e, " >>> ") + 4:]

			} else if p1 > 0 || p2 > 0 {
				// Error messages are not written in English
				if strings.Contains(e, " >>> ") { v.Command = command.Find(e) }
				p3 := strings.Index(e, " <<< "); if p3 < 0 { continue }
				v.Diagnosis = e[p3 + 4:]
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up error messages in e.Diagnosis, set the value of e.Reason
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			if strings.Contains(e.Diagnosis, "Unable to deliver") { e.Reason = "userunknown" }
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

