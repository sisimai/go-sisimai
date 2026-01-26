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
			" - spam",
			"//www.spamhaus.org/help/help_spam_16.htm",
			"//dsbl.org/help/help_spam_16.htm",
			"//mail.163.com/help/help_spam_16.htm",
			"554 5.7.0 reject, id=",
			"appears to be unsolicited",
			"blacklisted url in message",
			"block for spam",
			"blocked by policy: no spam please",
			"blocked by spamassassin",                      // rejected by SpamAssassin
			"blocked for abuse. see http://att.net/blocks", // AT&T
			"content filter rejection",
			"denied due to spam list",
			"is classified as spam and is rejected",
			"listed in work.drbl.imedia.ru",
			"mail content denied",            // http://service.mail.qq.com/cgi-bin/help?subtype=1&&id=20022&&no=1000726
			"may consider spam",
			"message content rejected",
			"message filtered",
			"message is being rejected as it seems to be a spam",
			"message rejected because of unacceptable content",
			"message rejected for policy reasons",
			"message was rejected for possible spam/virus content",
			"our email server thinks this email is spam",
			"our system has detected that this message is ",
			"reject bulk.advertising",
			"rejecting banned content",
			"rejecting mail content",
			"spam check",
			"spam content",
			"spam detected",
			"spam email",
			"spam message",
			"spam not accepted",
			"spam refused",
			"spam rejection",
			"spam-like",
			"spambouncer identified spam", // SpamBouncer identified SPAM
			"spamming not allowed",
			"too many spam complaints",
			"too much spam.",              // Earthlink
			"this message was rejected by recurrent pattern detection system",
			"we dont accept spam",
			"your email breaches local uribl policy",
			"your message has been temporarily blocked by our filter",
			"your message failed several antispam checks",
		}
		pairs := [][]string{
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
			[]string{"message ", "spamprofiler"},
			[]string{"probab", " spam"},
			[]string{"rejected by ", " (spam)"},
			[]string{"rejected due to spam ", "classification"},
			[]string{"rejected due to spam ", "content"},
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

