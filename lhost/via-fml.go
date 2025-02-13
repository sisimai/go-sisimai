// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      _______ __  __ _     
// | | |__   ___  ___| |_   / /  ___|  \/  | |    
// | | '_ \ / _ \/ __| __| / /| |_  | |\/| | |    
// | | | | | (_) \__ \ |_ / / |  _| | |  | | |___ 
// |_|_| |_|\___/|___/\__/_/  |_|   |_|  |_|_____|

package lhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Decode bounce messages from fml mailing list server/manager: https://www.fml.org
	InquireFor["FML"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if bf == nil || bf.Empty() == true                        { return sis.RisingUnderway{} }
		if len(bf.Headers["x-mlserver"])                     == 0 { return sis.RisingUnderway{} }
		if strings.Index(bf.Headers["from"][0], "-admin@")    < 1 { return sis.RisingUnderway{} }
		if strings.Index(bf.Headers["message-id"][0], ".FML") < 2 { return sis.RisingUnderway{} }

		boundaries := []string{"Original mail as follows:"}
		errortitle := map[string][]string{
			"rejected": []string{
				" are not member",
				"NOT MEMBER article from ",
				"reject mail ",
				"Spam mail from a spammer is rejected",
			},
			"systemerror": []string{
				"fml system error message",
				"Loop Alert: ",
				"Loop Back Warning: ",
				"WARNING: UNIX FROM Loop",
			},
			"securityerror": []string{"Security Alert"},
		}
		errortable := map[string][]string{
			"rejected": []string{
				" header may cause mail loop",
				"NOT MEMBER article from ",
				"reject mail from ",
				"reject spammers:",
				"You are not a member of this mailing list",
			},
			"systemerror": []string{
				" has detected a loop condition so that",
				"Duplicated Message-ID",
				"Loop Back Warning:",
			},
			"securityerror": []string{"Security alert:"},
		}
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if len(e) == 0 { continue }
			if cv := sisimoji.Select(e, "<", ">", 0); cv != "" {
				// You are not a member of this mailing list <neko-meeting@example.org>.
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = cv
				v.Diagnosis = e
				recipients += 1

			} else {
				// If you know the general guide of this list, please send mail with the mail body
				v.Diagnosis += e
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			for f := range errortable {
				// The key is a bounce reason name
				if sisimoji.ContainsAny(e.Diagnosis, errortable[f]) == false { continue }
				e.Reason = f; break
			}
			if e.Reason != "" { continue }

			for f := range errortitle {
				// The key is a bounce reason name
				if sisimoji.ContainsAny(bf.Headers["subject"][0], errortitle[f]) == false { continue }
				e.Reason = f; break
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

