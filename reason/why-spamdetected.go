// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                        ____       _            _           _ 
// / ___| _ __   __ _ _ __ ___ |  _ \  ___| |_ ___  ___| |_ ___  __| |
// \___ \| '_ \ / _` | '_ ` _ \| | | |/ _ \ __/ _ \/ __| __/ _ \/ _` |
//  ___) | |_) | (_| | | | | | | |_| |  __/ ||  __/ (__| ||  __/ (_| |
// |____/| .__/ \__,_|_| |_| |_|____/ \___|\__\___|\___|\__\___|\__,_|
//       |_|                                                          

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
	IncludedIn[eb.ReSPAM] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"blacklisted url in message",
			"block for spam",
			"blocked by policy: no spam please",
			"blocked by spamassassin",  // rejected by SpamAssassin
			"classified as spam and is rejected",
			"content filter rejection",
			"denied due to spam list",
			"identified spam",          // 554 SpamBouncer identified SPAM, message permanently rejected (#5.3.0)
			"may consider spam",
			"message content rejected",
			"message has been temporarily blocked by our filter",
			"message is being rejected as it seems to be a spam",
			"message was rejected by recurrent pattern detection system",
			"our email server thinks this email is spam",
			"reject bulk.advertising",
			"spam check",
			"spam content ",
			"spam detected",
			"spam email",
			"spam-like header",
			"spam message",
			"spam not accepted",
			"spam refused",
			"spamming not allowed",
			"unsolicited ",
			"your email breaches local uribl policy",
		}
		pairs := [][]string{
			[]string{"accept", " spam"},
			[]string{"appears", " to ", "spam"},
			[]string{"bulk", "mail"},
			[]string{"considered", " spam"},
			[]string{"contain", " spam"},
			[]string{"detected", " spam"},
			[]string{"greylisted", " please try again in"},
			[]string{"mail score (", " over "},
			[]string{"mail rejete. mail rejected. ", "506"},
			[]string{"message ", "as spam"},
			[]string{"message ", "like spam"},
			[]string{"probab", " spam"},
			[]string{"refused by", " spamprofiler"},
			[]string{"reject", " content"},
			[]string{"reject, id=", "spam"},
			[]string{"rejected by ", " (spam)"},
			[]string{"rejected due to spam ", "classification"},
			[]string{"rule imposed as ", " is blacklisted on"},
			[]string{"score", "spam"},
			[]string{"spam ", "block"},
			[]string{"spam ", "filter"},
			[]string{"spam ", " exceeded"},
			[]string{"spam ", "score"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReSPAM] = func(fo *siba.Fact) bool {
		if fo == nil || fo.DeliveryStatus == ""           { return false }
		if fo.Reason == eb.ReSPAM                         { return true  }
		if status.Name(fo.DeliveryStatus) == eb.ReSPAM    { return true  }
		if slices.Contains(command.ExceptDATA, fo.Command){ return false }
		return IncludedIn[eb.ReSPAM](strings.ToLower(fo.DiagnosticCode))
	}
}

