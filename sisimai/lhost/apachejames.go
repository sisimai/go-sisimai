// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ___                     _              _                           
// | | |__   ___  ___| |_   / / \   _ __   __ _  ___| |__   ___    | | __ _ _ __ ___   ___  ___ 
// | | '_ \ / _ \/ __| __| / / _ \ | '_ \ / _` |/ __| '_ \ / _ \_  | |/ _` | '_ ` _ \ / _ \/ __|
// | | | | | (_) \__ \ |_ / / ___ \| |_) | (_| | (__| | | |  __/ |_| | (_| | | | | | |  __/\__ \
// |_|_| |_|\___/|___/\__/_/_/   \_\ .__/ \__,_|\___|_| |_|\___|\___/ \__,_|_| |_| |_|\___||___/
//                                 |_|                                                          
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from James: https://james.apache.org/
	InquireFor["ApacheJames"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		proceedsto := false; ISJAMES: for {
			// Subject:     [BOUNCE]
			// Message-Id:  JavaMail.
			if bf.Head["subject"][0] == "[BOUNCE]"                      { proceedsto = true; break ISJAMES }
			if strings.Contains(bf.Head["message-id"][0], ".JavaMail.") { proceedsto = true; break ISJAMES }
			for _, e := range bf.Head["received"] {
				// Received: from localhost ([127.0.0.1])
				//    by mx.example.org (JAMES SMTP Server 2.3.2) with SMTP ID 220...
				if strings.Contains(e, "JAMES SMTP Server") == true     { proceedsto = true; break ISJAMES }
			}
			break
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{
			// apache-james-2.3.2/src/java/org/apache/james/transport/mailets/
			//   AbstractNotify.java|124:  out.println("Error message below:");
			//   AbstractNotify.java|128:  out.println("Message details:");
			"message": []string{"Message details:"},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)                  // Points the current cursor position
		recipients := uint8(0)                  // The number of 'Final-Recipient' header
		alternates := [4]string{"", "", "", ""} // [Envelope-From, Header-From, Date, Subject]
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) {
					// Message details:
					//   Subject: Nyaaan
					readcursor |= indicators["deliverystatus"]
					continue
				}
				if e != "" { v.Diagnosis += e + " " }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// Message details:
			//   Subject: Nyaaan
			//   Sent date: Thu Apr 29 01:20:50 JST 2015
			//   MAIL FROM: shironeko@example.jp
			//   RCPT TO: kijitora@example.org
			//   From: Neko <shironeko@example.jp>
			//   To: kijitora@example.org
			//   Size (in bytes): 1024
			//   Number of lines: 64
			if strings.HasPrefix(e, "  RCPT TO: ") {
				//   RCPT TO: kijitora@example.org
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = e[12:]
				recipients += 1

			} else if strings.HasPrefix(e, "  Sent date: ") {
				//   Sent date: Thu Apr 29 01:20:50 JST 2015
				v.Date        = e[13:]
				alternates[2] = v.Date

			} else if strings.HasPrefix(e, "  Subject: ") {
				//   Subject: Nyaaan
				alternates[3] = e[11:]

			} else if strings.HasPrefix(e, "  MAIL FROM: ") {
				//   MAIL FROM: shironeko@example.jp
				alternates[0] = e[13:]

			} else if strings.HasPrefix(e, "  From: ") {
				//   From: Neko <shironeko@example.jp>
				alternates[1] = e[8:]
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		if emailparts[1] == "" {
			// The original message is empty
			if alternates[1] != "" { emailparts[1] += fmt.Sprintf("From: %s\n", alternates[1]) }
			if alternates[2] != "" { emailparts[1] += fmt.Sprintf("Date: %s\n", alternates[2]) }
		}
		if strings.Contains(emailparts[1], "Return-Path: ") == false {
			// Set the envelope from address as a Return-Path: header
			if alternates[0] != "" { emailparts[1] += fmt.Sprintf("Return-Path: <%s>\n", alternates[0]) }
		}
		if strings.Contains(emailparts[1], "\nSubject: ") == false {
			// There is no Subject field in the original message
			if alternates[3] != "" { emailparts[1] += fmt.Sprintf("Subject: %s\n", alternates[3]) }
		}

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

