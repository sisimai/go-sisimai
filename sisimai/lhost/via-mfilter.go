// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      __         _____ ___ _   _____ _____ ____  
// | | |__   ___  ___| |_   / / __ ___ |  ___|_ _| | |_   _| ____|  _ \ 
// | | '_ \ / _ \/ __| __| / / '_ ` _ \| |_   | || |   | | |  _| | |_) |
// | | | | | (_) \__ \ |_ / /| | | | | |  _|  | || |___| | | |___|  _ < 
// |_|_| |_|\___/|___/\__/_/ |_| |_| |_|_|   |___|_____|_| |_____|_| \_\
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisiaddr "sisimai/address"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from Digital Arts m-FILTER: https://www.daj.jp/bs/mf/
	InquireFor["mFILTER"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true                { return sis.RisingUnderway{} }
		if len(bf.Headers["x-mailer"]) < 1                { return sis.RisingUnderway{} }
		if bf.Headers["x-mailer"][0]  != "m-FILTER"       { return sis.RisingUnderway{} }
		if bf.Headers["subject"][0]   != "failure notice" { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"-------original message", "-------original mail info"}
		startingof := map[string][]string{
			"command": []string{"-------SMTP command"},
			"error":   []string{"-------server message"},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		readcursor := uint8(0)              // Points the current cursor position
		recipients := 0                     // The number of 'Final-Recipient' header
		markingset := [2]bool{false, false} // [diganosis, command]
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.Index(e, "@") > 1 && strings.Contains(e, " ") == false && sisiaddr.IsEmailAddress(e) {
					// This line contains an email address only: "kijitora@example.jp"
					readcursor |= indicators["deliverystatus"]
				}
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// このメールは「m-FILTER」が自動的に生成して送信しています。
			// メールサーバーとの通信中、下記の理由により
			// このメールは送信できませんでした。
			//
			// 以下のメールアドレスへの送信に失敗しました。
			// kijitora@example.jp
			//
			//
			// -------server message
			// 550 5.1.1 unknown user <kijitora@example.jp>
			//
			// -------SMTP command
			// DATA
			//
			// -------original message
			if strings.Contains(e, "@") && strings.Contains(e, " ") == false {
				// 以下のメールアドレスへの送信に失敗しました。
				// kijitora@example.jp
				if sisiaddr.IsEmailAddress(e) == false { continue }
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = e
				recipients += 1

			} else {
				// Get error messages and the last SMTP command
				if e == startingof["error"][0] { markingset[0] = true; continue }
				if e == startingof["command"][0] {
					// -------SMTP command
					markingset[0] = false
					markingset[1] = true
					continue
				}

				if markingset[0] == true && v.Diagnosis == "" {
					// -------server message
					// 550 5.1.1 unknown user <kijitora@example.jp>
					v.Diagnosis = e; markingset[0] = false

				} else if markingset[1] == true && command.Test(e) {
					// -------SMTP command
					// DATA
					v.Command = e; markingset[1] = false
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

