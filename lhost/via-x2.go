// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      ____  ______  
// | | |__   ___  ___| |_   / /\ \/ /___ \ 
// | | '_ \ / _ \/ __| __| / /  \  /  __) |
// | | | | | (_) \__ \ |_ / /   /  \ / __/ 
// |_|_| |_|\___/|___/\__/_/   /_/\_\_____|

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"
import sisiaddr "libsisimai.org/sisimai/address"

func init() {
	// Decode bounce messages from Unknown MTA #2
	InquireFor["X2"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true { return sis.RisingUnderway{} }

		proceedsto := false
		emailtitle := []string{"Delivery failure", "failure delivery", "failed delivery"}
		if strings.Contains(bf.Headers["from"][0], "MAILER-DAEMON@")   { proceedsto = true }
		if sisimoji.HasPrefixAny(bf.Headers["subject"][0], emailtitle) { proceedsto = true }
		if proceedsto == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"--- Original message follows."}
		startingof := map[string][]string{
			"message": []string{"Unable to deliver message to the following address"},
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
			if readcursor & indicators["deliverystatus"] == 0 || e == "" { continue }

			// Message from example.com.
			// Unable to deliver message to the following address(es).
			//
			// <kijitora@example.com>:
			// This user doesn't have a example.com account (kijitora@example.com) [0]
			if strings.HasPrefix(e, "<") && sisimoji.Aligned(e, []string{"<", "@", ">:"}) {
				// <kijitora@example.com>:
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = sisiaddr.S3S4(e[:strings.IndexByte(e, ':')])
				recipients += 1

			} else {
				// Error message lines like the following:
				// This user doesn't have a example.com account (kijitora@example.com) [0]
				v.Diagnosis += " " + e
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

