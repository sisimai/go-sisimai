// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      ______  _       _       _          
// | | |__   ___  ___| |_   / / __ )(_) __ _| | ___ | |__   ___ 
// | | '_ \ / _ \/ __| __| / /|  _ \| |/ _` | |/ _ \| '_ \ / _ \
// | | | | | (_) \__ \ |_ / / | |_) | | (_| | | (_) | |_) |  __/
// |_|_| |_|\___/|___/\__/_/  |____/|_|\__, |_|\___/|_.__/ \___|
//                                     |___/                    
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from BIGLOBE: https://www.biglobe.ne.jp
	InquireFor["Biglobe"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true                                 { return sis.RisingUnderway{} }
		if strings.Contains(bf.Headers["from"][0], "postmaster@") == false { return sis.RisingUnderway{} }
		if strings.Index(bf.Headers["subject"][0], "Returned mail:") != 0  { return sis.RisingUnderway{} }

		jpprovider := []string{"biglobe", "inacatv", "tmtv", "ttv"}
		proceedsto := false; for _, e := range jpprovider {
			// The From: header should contain one of domain above
			if strings.Contains(bf.Headers["from"][0], "@" + e + ".ne.jp") { proceedsto = true; break }
		}
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822"}
		startingof := map[string][]string{
			"message": []string{"   ----- The following addresses had delivery problems -----"},
			"error":   []string{"   ----- Non-delivered information -----"},
		}
		messagesof := map[string][]string{
			"filtered":    []string{"Mail Delivery Failed... User unknown"},
			"mailboxfull": []string{"The number of messages in recipient's mailbox exceeded the local limit."},
		}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
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

			// This is a MIME-encapsulated message.
			//
			// ----_Biglobe000000/00000.biglobe.ne.jp
			// Content-Type: text/plain; charset="iso-2022-jp"
			//
			//    ----- The following addresses had delivery problems -----
			// ********@***.biglobe.ne.jp
			//
			//    ----- Non-delivered information -----
			// The number of messages in recipient's mailbox exceeded the local limit.
			//
			// ----_Biglobe000000/00000.biglobe.ne.jp
			// Content-Type: message/rfc822
			//
			if strings.Index(e, "@") > 1 && strings.Contains(e, " ") == false {
				//    ----- The following addresses had delivery problems -----
				// ********@***.biglobe.ne.jp
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				if rfc5322.IsEmailAddress(e) == false { continue }
				v.Recipient = e
				recipients += 1

			} else {
				// The boundary string or the error messages
				if strings.Contains(e, "--") { continue }
				v.Diagnosis += e + " "
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

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

