// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      _____ __  __       _ _ ____                           
// | | |__   ___  ___| |_   / /_ _|  \/  | __ _(_) / ___|  ___ _ ____   _____ _ __ 
// | | '_ \ / _ \/ __| __| / / | || |\/| |/ _` | | \___ \ / _ \ '__\ \ / / _ \ '__|
// | | | | | (_) \__ \ |_ / /  | || |  | | (_| | | |___) |  __/ |   \ V /  __/ |   
// |_|_| |_|\___/|___/\__/_/  |___|_|  |_|\__,_|_|_|____/ \___|_|    \_/ \___|_|   
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from Progress iMail Server: https://community.progress.com/s/products/imailserver
	InquireFor["IMailServer"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		// X-Mailer: <SMTP32 v8.22>
		proceedsto := false
		if strings.HasPrefix(bf.Headers["subject"][0], "Undeliverable Mail ") { proceedsto = true }
		if len(bf.Headers["x-mailer"]) > 0 && strings.HasPrefix(bf.Headers["x-mailer"][0], "<SMTP32 v") { proceedsto = true }
		if proceedsto == false { return sis.RisingUnderway{} }

		boundaries := []string{"Original message follows."}
		startingof := map[string][]string{"error": []string{"Body of message generated response:"}}
		messagesof := map[string][]string{
			"hostunknown":   []string{"Unknown host"},
			"userunknown":   []string{"Unknown user", "Invalid final delivery userid"},
			"mailboxfull":   []string{"User mailbox exceeds allowed size"},
			"virusdetected": []string{"Requested action not taken: virus detected"},
			"spamdetected":  []string{"Blacklisted URL in message"},
			"expired":       []string{"Delivery failed "},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		alternates := ""                  // Other error message strings
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			p0 := strings.Index(e, ": ")
			if (p0 > 8 && sisimoji.Aligned(e, []string{": ", "@"})) || strings.HasPrefix(e, "undeliverable ") {
				// Unknown user: kijitora@example.com
				// undeliverable to kijitora@example.com
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Diagnosis = e
				v.Recipient = sisiaddr.Find(e)[0]
				recipients += 1

			} else {
				// Other error messages
				if strings.Contains(e, startingof["error"][0]) {
					// Body of message generated response:
					alternates = e

				} else {
					// Error message after "Body of message generated response:" line
					if alternates == "" { continue }
					alternates += " " + e
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])

			if alternates != "" {
				// Copy the alternative error message to e.Diagnosis
				e.Diagnosis = alternates + " " + e.Diagnosis
				e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			}
			e.Diagnosis = sisimoji.Sweep(strings.ReplaceAll(e.Diagnosis, "\n", " "))
			e.Command   = command.Find(e.Diagnosis)

			FINDREASON: for r := range messagesof {
				// The key name is a bounce reason name
				for _, f := range messagesof[r] {
					// Try to find an error message including lower-cased string listed in messagesof
					if strings.Contains(e.Diagnosis, f) == false { continue }
					e.Reason = r; break FINDREASON
				}
			}
		}

		if strings.Contains(emailparts[1], "\nFrom: ") == false {
			// Set pseudo From: header into the original message
			emailparts[1] = fmt.Sprintf("From: %s\n%s\n", bf.Headers["to"][0], emailparts[1])
		}
		if strings.Contains(emailparts[1], "\nTo: ") == false {
			// Set pseudo To: header into the original message
			emailparts[1] = fmt.Sprintf("To: %s\n%s\n", dscontents[0].Recipient, emailparts[1])
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

