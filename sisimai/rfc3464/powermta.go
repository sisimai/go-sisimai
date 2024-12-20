// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc3464

//  ____  _____ ____ _____ _  _    __   _  _     ______                        __  __ _____  _    
// |  _ \|  ___/ ___|___ /| || |  / /_ | || |   / /  _ \ _____      _____ _ __|  \/  |_   _|/ \   
// | |_) | |_ | |     |_ \| || |_| '_ \| || |_ / /| |_) / _ \ \ /\ / / _ \ '__| |\/| | | | / _ \  
// |  _ <|  _|| |___ ___) |__   _| (_) |__   _/ / |  __/ (_) \ V  V /  __/ |  | |  | | | |/ ___ \ 
// |_| \_\_|   \____|____/   |_|  \___/   |_|/_/  |_|   \___/ \_/\_/ \___|_|  |_|  |_| |_/_/   \_\
import "strings"

func init() {
	// Returns []string which is compatible with the value returned from rfc1894.Field()
	ReturnedBy["PowerMTA"] = func(argv1 string) []string {
		// @param    string argv1   A line of the error message
		// @return   []string       []string{"field-name", "value-type", "value", "field-group", "comment"}
		// @see      https://bird.com/email/power-mta
		if argv1 == "" || strings.Contains(argv1, ": ") == false { return []string{} }

		fieldgroup := map[string]string{
			"x-powermta-virtualmta":     "host", // X-PowerMTA-VirtualMTA: mx22.neko.example.jp
			"x-powermta-bouncecategory": "text", // X-PowerMTA-BounceCategory: bad-mailbox
		}
		messagesof := map[string]string{
			"bad-domain":          "hostunknown",
			"bad-mailbox":         "userunknown",
			"inactive-mailbox":    "disabled",
			"message-expired":     "expired",
			"no-answer-from-host": "networkerror",
			"policy-related":      "policyviolation",
			"quota-issues":        "mailboxfull",
			"routing-errors":      "systemerror",
			"spam-related":        "spamdetected",
		}
		fieldparts := strings.SplitN(argv1, ":", 2)  // []string{"Final-Recipient", " rfc822; <neko@example.jp>"}
		xfieldname := strings.ToLower(fieldparts[0]) // "final-recipient"
		xef, nyaan := fieldgroup[xfieldname]; if nyaan == false { return []string{} }
		xfieldlist := []string{"", "", strings.TrimSpace(fieldparts[1]), xef, "", "PowerMTA"}

		// - 0: Field-Name
		// - 1: Sub Type: RFC822, DNS, X-Unix, and so on)
		// - 2: Value
		// - 3: Field Group(addr, code, date, host, stat, text)
		// - 4: Comment
		// - 5: 3rd Party MTA-Name
		if xfieldname == "x-powermta-bouncecategory" {
			// X-PowerMTA-BounceCategory: bad-mailbox
			// Set the bounce reason picked from the value of the field
			xfieldlist[0] = xfieldname
			if len(messagesof[xfieldlist[2]]) > 0 {
				// "reason:mailboxfull"; the 5th value supposed to be assigned to "Reason" member
				// of sis.DeliveryMatter{} struct.
				xfieldlist[4] = "reason:" + messagesof[xfieldlist[2]]
			}
		} else if xfieldname == "x-powermta-virtualmta" {
			// X-PowerMTA-VirtualMTA: mx22.neko.example.jp
			xfieldlist[0] = "Reporting-MTA"
		}

		return xfieldlist
	}
}

