// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _____ _  _    __   _  _     ______                        __  __ _____  _    
// |  _ \|  ___/ ___|___ /| || |  / /_ | || |   / /  _ \ _____      _____ _ __|  \/  |_   _|/ \   
// | |_) | |_ | |     |_ \| || |_| '_ \| || |_ / /| |_) / _ \ \ /\ / / _ \ '__| |\/| | | | / _ \  
// |  _ <|  _|| |___ ___) |__   _| (_) |__   _/ / |  __/ (_) \ V  V /  __/ |  | |  | | | |/ ___ \ 
// |_| \_\_|   \____|____/   |_|  \___/   |_|/_/  |_|   \___/ \_/\_/ \___|_|  |_|  |_| |_/_/   \_\

package rfc3464
import "strings"

func init() {
	// ReturnedBy["PowerMTA"] returns a []string which is compatible with the value returned from rfc1894.Field().
	//   Arguments:
	//     - mesg (string): Line of the error message.
	//   Returns:
	//     - ([]string): []string{"field-name", "value-type", "value", "field-group", "comment"}
	//   See:
	//     - https://bird.com/email/power-mta
	ReturnedBy["PowerMTA"] = func(mesg string) []string {
		if mesg == "" || strings.Contains(mesg, ": ") == false { return []string{} }

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
		lhs,rhs, _ := strings.Cut(mesg, ":") // []string{"Final-Recipient", " rfc822; <neko@example.jp>"}
		xfieldname := strings.ToLower(lhs)   // "final-recipient"
		xef, nyaan := fieldgroup[xfieldname]; if nyaan == false { return []string{} }
		xfieldlist := []string{"", "", strings.TrimSpace(rhs), xef, "", "PowerMTA"}

		// - 0: Field-Name
		// - 1: Sub Type: RFC822, DNS, X-Unix, and so on)
		// - 2: Value
		// - 3: Field Group(addr, code, date, host, stat, text)
		// - 4: Comment
		// - 5: 3rd Party MTA-Name
		switch xfieldname {
			// X-PowerMTA-BounceCategory: bad-mailbox
			// Set the bounce reason picked from the value of the field
			case "x-powermta-bouncecategory":
				xfieldlist[0] = xfieldname
				if len(messagesof[xfieldlist[2]]) > 0 {
					// "reason:mailboxfull"; the 5th value supposed to be assigned to "Reason" member
					// of sis.DeliveryMatter{} struct.
					xfieldlist[4] = "reason:" + messagesof[xfieldlist[2]]
				}
			// X-PowerMTA-VirtualMTA: mx22.neko.example.jp
			case "x-powermta-virtualmta": xfieldlist[0] = "Reporting-MTA"
		}

		return xfieldlist
	}
}

