// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____       _       _     _           _ _           _ 
// |  _ \ __ _| |_ ___| |   (_)_ __ ___ (_) |_ ___  __| |
// | |_) / _` | __/ _ \ |   | | '_ ` _ \| | __/ _ \/ _` |
// |  _ < (_| | ||  __/ |___| | | | | | | | ||  __/ (_| |
// |_| \_\__,_|\__\___|_____|_|_| |_| |_|_|\__\___|\__,_|

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
	IncludedIn[eb.ReRATE] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"has exceeded the max emails per hour ",
			"please try again slower",
			"receiving mail at a rate that prevents additional messages from being delivered",
			"temporarily deferred due to unexpected volume or user complaints",
			"throttling failure: ",
			"too many errors from your ip",         // Free.fr
			"too many recipients",                  // ntt docomo
			"trop de connexions, ",
			"we have already made numerous attempts to deliver this message",
		}
		pairs := [][]string{
			[]string{"connection ", "limit"},
			[]string{"exceeded ", "allowable number of posts without solving a captcha"},
			[]string{"temporarily", "rate limited"},
			[]string{"throttled ", "postmaster.comcast.net"},
			[]string{"too many con", "s"},
			[]string{"too many ", "sessions "}, // Sendmail(daemon.c), comcast.net
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReRATE] = func(fo *siba.Fact) bool {
		if fo == nil                                   { return false }
		if fo.Reason == eb.ReRATE                      { return true  }
		if status.Name(fo.DeliveryStatus) == eb.ReRATE { return true  }
		return IncludedIn[eb.ReRATE](strings.ToLower(fo.DiagnosticCode))
	}
}

