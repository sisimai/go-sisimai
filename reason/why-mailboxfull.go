// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  __  __       _ _ _               _____      _ _ 
// |  \/  | __ _(_) | |__   _____  _|  ___|   _| | |
// | |\/| |/ _` | | | '_ \ / _ \ \/ / |_ | | | | | |
// | |  | | (_| | | | |_) | (_) >  <|  _|| |_| | | |
// |_|  |_|\__,_|_|_|_.__/ \___/_/\_\_|   \__,_|_|_|

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReFULL] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"452 insufficient disk space",
			"account disabled temporarly for exceeding receiving limits",
			"boite du destinataire pleine",
			"exceeded storage allocation",
			"full mailbox",
			"mailbox exceeds allowed size",
			"mailbox size limit exceeded",
			"mailbox would exceed maximum allowed storage",
			"mailfolder is full",
			"no space left on device",
			"not sufficient disk space",
			"quota violation for",
			"too much mail data", // @docomo.ne.jp
			"user has exceeded quota, bouncing mail",
			"user has too many messages on the server",
			"user's space has been used up",
		}
		pairs := [][]string{
			[]string{"account is ", " quota"},
			[]string{"disk", "quota"},
			[]string{"enough ", " space"},
			[]string{"mailbox ", "exceeded", " limit"},
			[]string{"mailbox ", "full"},
			[]string{"mailbox ", "quota"},
			[]string{"maildir ", "quota"},
			[]string{"over ", "quota"},
			[]string{"quota ", "exceeded"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReFULL] = func(fo *siba.Fact) bool {
		// Delivery status code points MailboxFull.
		// Status: 4.2.2
		// Diagnostic-Code: SMTP; 450 4.2.2 <***@example.jp>... Mailbox Full
		if fo == nil                                   { return false }
		if fo.Reason == eb.ReFULL                      { return true  }
		if status.Name(fo.DeliveryStatus) == eb.ReFULL { return true  }
		return IncludedIn[eb.ReFULL](strings.ToLower(fo.DiagnosticCode))
	}
}

