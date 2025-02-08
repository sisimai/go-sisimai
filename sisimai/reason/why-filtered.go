// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____ _ _ _                    _ 
// |  ___(_) | |_ ___ _ __ ___  __| |
// | |_  | | | __/ _ \ '__/ _ \/ _` |
// |  _| | | | ||  __/ | |  __/ (_| |
// |_|   |_|_|\__\___|_|  \___|\__,_|
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"
import sisimoji "sisimai/string"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["Filtered"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"because the recipient is only accepting mail from specific email addresses", // AOL Phoenix
			"bounced address", // SendGrid|a message to an address has previously been Bounced.
			"due to extended inactivity new mail is not currently being accepted for this mailbox",
			"has restricted sms e-mail", // AT&T
			"is not accepting any mail",
			"message filtered",
			"message rejected due to user rules",
			"not found recipient account",
			"refused due to recipient preferences", // Facebook
			"resolver.rst.notauthorized", // Microsoft Exchange
			"this account is protected by",
			"user not found", // Filter on MAIL.RU
			"user refuses to receive this mail",
			"user reject",
			"we failed to deliver mail because the following address recipient id refuse to receive mail", // Willcom
			"you have been blocked by the recipient",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "filtered" or not
	ProbesInto["Filtered"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is filtered, false: is not filtered
		if fo        == nil        { return false }
		if fo.Reason == "filtered" { return true  }

		tempreason := status.Name(fo.DeliveryStatus); if tempreason == "suspend" { return false }
		issuedcode := strings.ToLower(fo.DiagnosticCode)

		if tempreason == "filtered" {
			// The value of delivery status code points "filtered".
			if IncludedIn["UserUnknown"](issuedcode) || IncludedIn["Filtered"](issuedcode) { return true }

		} else {
			// The value of "Reason" is not "filtered" when the value of "fo.Command" is an SMTP
			// command to be sent before the SMTP DATA command because all the MTAs read the headers
			// and the entire message body after the DATA command.
			if sisimoji.EqualsAny(fo.Command, []string{"CONN", "EHLO", "HELO", "MAIL", "RCPT"}) { return false }
			if IncludedIn["Filtered"](issuedcode) || IncludedIn["UserUnknown"](issuedcode)      { return true  }
		}
		return false
	}
}

