// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//     _         _   _     _____     _ _                
//    / \  _   _| |_| |__ |  ___|_ _(_) |_   _ _ __ ___ 
//   / _ \| | | | __| '_ \| |_ / _` | | | | | | '__/ _ \
//  / ___ \ |_| | |_| | | |  _| (_| | | | |_| | | |  __/
// /_/   \_\__,_|\__|_| |_|_|  \__,_|_|_|\__,_|_|  \___|

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
	IncludedIn[eb.ReAUTH] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"//spf.pobox.com",
			"5322.From address doesn't meet the authentication requirements",
			"bad spf records for",
			"dmarc policy",
			"doesn't meet the required authentication level",
			"please inspect your spf settings",
			"sender policy framework",
			"spf check: fail",
		}
		pairs := [][]string{
			[]string{"spf: ", " is not allowed to send "},
			[]string{"is not allowed to send ", " spf "},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReAUTH] = func(fo *siba.Fact) bool {
		if fo == nil                                   { return false }
		if fo.Reason == eb.ReAUTH                      { return true  }
		if status.Name(fo.DeliveryStatus) == eb.ReAUTH { return true  }
		return IncludedIn[eb.ReAUTH](strings.ToLower(fo.DiagnosticCode))
	}
}

