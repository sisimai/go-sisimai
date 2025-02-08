// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  __  __       _ _ _               _____      _ _ 
// |  \/  | __ _(_) | |__   _____  _|  ___|   _| | |
// | |\/| |/ _` | | | '_ \ / _ \ \/ / |_ | | | | | |
// | |  | | (_| | | | |_) | (_) >  <|  _|| |_| | | |
// |_|  |_|\__,_|_|_|_.__/ \___/_/\_\_|   \__,_|_|_|
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["MailboxFull"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"452 insufficient disk space",
			"account disabled temporarly for exceeding receiving limits",
			"account is exceeding their quota",
			"account is over quota",
			"account is temporarily over quota",
			"boite du destinataire pleine",
			"delivery failed: over quota",
			"disc quota exceeded",
			"diskspace quota",
			"does not have enough space",
			"exceeded storage allocation",
			"exceeding its mailbox quota",
			"full mailbox",
			"is over disk quota",
			"is over quota temporarily",
			"mail file size exceeds the maximum size allowed for mail delivery",
			"mail quota exceeded",
			"mailbox exceeded the local limit",
			"mailbox full",
			"mailbox has exceeded its disk space limit",
			"mailbox is full",
			"mailbox over quota",
			"mailbox quota usage exceeded",
			"mailbox size limit exceeded",
			"maildir over quota",
			"maildir delivery failed: userdisk quota ",
			"maildir delivery failed: domaindisk quota ",
			"mailfolder is full",
			"no space left on device",
			"not enough disk space",
			"not enough storage space in",
			"not sufficient disk space",
			"over the allowed quota",
			"quota exceeded",
			"quota violation for",
			"recipient reached disk quota",
			"recipient rejected: mailbox would exceed maximum allowed storage",
			"the recipient mailbox has exceeded its disk space limit",
			"the user's space has been used up",
			"the user you are trying to reach is over quota",
			"too much mail data", // @docomo.ne.jp
			"user has exceeded quota, bouncing mail",
			"user has too many messages on the server",
			"user is over quota",
			"user is over the quota",
			"user over quota",
			"user over quota. (#5.1.1)", // qmail-toaster
			"was automatically rejected: quota exceeded",
			"would be over the allowed quota",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "mailboxfull" or not
	ProbesInto["MailboxFull"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is mailboxfull, false: is not mailboxfull

		// Delivery status code points "mailboxfull".
		// Status: 4.2.2
		// Diagnostic-Code: SMTP; 450 4.2.2 <***@example.jp>... Mailbox Full
		if fo == nil                                       { return false }
		if fo.Reason == "mailboxfull"                      { return true  }
		if status.Name(fo.DeliveryStatus) == "mailboxfull" { return true  }
		return IncludedIn["MailboxFull"](strings.ToLower(fo.DiagnosticCode))
	}
}

