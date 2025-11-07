// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____           __  __                    ____                  
// |_   _|__   ___ |  \/  | __ _ _ __  _   _ / ___|___  _ __  _ __  
//   | |/ _ \ / _ \| |\/| |/ _` | '_ \| | | | |   / _ \| '_ \| '_ \ 
//   | | (_) | (_) | |  | | (_| | | | | |_| | |__| (_) | | | | | | |
//   |_|\___/ \___/|_|  |_|\__,_|_| |_|\__, |\____\___/|_| |_|_| |_|
//                                     |___/                        

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
	IncludedIn[eb.ReCONN] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"all available ips are at maximum connection limit", // SendGrid
			"connection rate limit exceeded",
			"exceeds per-domain connection limit for",
			"has exceeded the max emails per hour ",
			"throttling failure: daily message quota exceeded",
			"throttling failure: maximum sending rate exceeded",
			"too many connections",
			"too many connections from your host.", // Microsoft
			"too many concurrent smtp connections", // Microsoft
			"too many errors from your ip",         // Free.fr
			"too many recipients",                  // ntt docomo
			"too many smtp sessions for this host", // Sendmail(daemon.c)
			"trop de connexions, ",
			"we have already made numerous attempts to deliver this message",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReCONN] = func(fo *siba.Fact) bool {
		if fo == nil                                   { return false }
		if fo.Reason == eb.ReCONN                      { return true  }
		if status.Name(fo.DeliveryStatus) == eb.ReCONN { return true  }
		return IncludedIn[eb.ReCONN](strings.ToLower(fo.DiagnosticCode))
	}
}

