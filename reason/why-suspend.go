// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                                 _ 
// / ___| _   _ ___ _ __   ___ _ __   __| |
// \___ \| | | / __| '_ \ / _ \ '_ \ / _` |
//  ___) | |_| \__ \ |_) |  __/ | | | (_| |
// |____/ \__,_|___/ .__/ \___|_| |_|\__,_|
//                 |_|                     

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReQUIT] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			" currently suspended",
			" temporary locked",
			"address no longer accepts mail",
			"archived recipient",
			"boite du destinataire archivee",
			"email account that you tried to reach is inactive",
			"inactive account",
			"invalid/inactive user",
			"is a deactivated mailbox", // http://service.mail.qq.com/cgi-bin/help?subtype=1&&id=20022&&no=1000742
			"is unavailable: user is terminated",
			"mailbox is frozen",
			"mailbox is inactive",
			"mailbox unavailable or access denied",
			"recipient rejected: temporarily inactive",
			"recipient suspend the service",
			"user or domain is disabled",
			"user suspended", // http://mail.163.com/help/help_spam_16.htm
			"vdelivermail: account is locked email bounced",
		}
		pairs := [][]string{
			[]string{"account ", "disabled"},
			[]string{"has been ", "suspended"},
			[]string{"mailbox ", "disabled"},
			[]string{"not ", "active"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReQUIT] = func(fo *siba.Fact) bool {
		if fo == nil                                       { return false }
		if fo.Reason == eb.ReQUIT || fo.ReplyCode == "525" { return true  }
		return IncludedIn[eb.ReQUIT](strings.ToLower(fo.DiagnosticCode))
	}
}

