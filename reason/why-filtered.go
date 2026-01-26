// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____ _ _ _                    _ 
// |  ___(_) | |_ ___ _ __ ___  __| |
// | |_  | | | __/ _ \ '__/ _ \/ _` |
// |  _| | | | ||  __/ | |  __/ (_| |
// |_|   |_|_|\__\___|_|  \___|\__,_|

package reason
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"
import "libsisimai.org/sisimai/v5/smtp/command"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReFILT] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"bounced address", // SendGrid|a message to an address has previously been Bounced.
			"due to extended inactivity new mail is not currently being accepted for this mailbox",
			"has restricted sms e-mail", // AT&T
			"is not accepting any mail",
			"message filtered",
			"message rejected due to user rules",
			"not found recipient account",
			"recipient id refuse to receive mail", // Willcom
			"recipient is only accepting mail from specific email addresses", // AOL Phoenix
			"refused due to recipient preferences", // Facebook
			"resolver.rst.notauthorized", // Microsoft Exchange
			"this account is protected by",
			"user not found", // Filter on MAIL.RU
			"user refuses to receive this mail",
			"user reject",
			"you have been blocked by the recipient",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReFILT] = func(fo *siba.Fact) bool {
		if fo        == nil       { return false }
		if fo.Reason == eb.ReFILT { return true  }

		tempreason := status.Name(fo.DeliveryStatus); if tempreason == eb.ReQUIT { return false }
		issuedcode := strings.ToLower(fo.DiagnosticCode)

		if tempreason == eb.ReFILT {
			// The value of delivery status code points Filtered.
			if IncludedIn[eb.ReUSER](issuedcode) || IncludedIn[eb.ReFILT](issuedcode) { return true }

		} else {
			// The value of "Reason" is not Filtered when the value of "fo.Command" is an SMTP
			// command to be sent before the SMTP DATA command because all the MTAs read the headers
			// and the entire message body after the DATA command.
			if slices.Contains(command.ExceptDATA, fo.Command) { return false }
			if IncludedIn[eb.ReFILT](issuedcode) || IncludedIn[eb.ReUSER](issuedcode) { return true }
		}
		return false
	}
}

