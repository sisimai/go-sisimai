// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//     _         _   _     _____     _ _                
//    / \  _   _| |_| |__ |  ___|_ _(_) |_   _ _ __ ___ 
//   / _ \| | | | __| '_ \| |_ / _` | | | | | | '__/ _ \
//  / ___ \ |_| | |_| | | |  _| (_| | | | |_| | | |  __/
// /_/   \_\__,_|\__|_| |_|_|  \__,_|_|_|\__,_|_|  \___|

package reason
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn["AuthFailure"] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"//spf.pobox.com",
			"bad spf records for",
			"dmarc policy",
			"doesn't meet the required authentication level",
			"please inspect your spf settings",
			"sender policy framework (spf) fail",
			"sender policy framework violation",
			"spf (sender policy framework) domain authentication fail",
			"spf check: fail",
			"the 5322.From address doesn't meet the authentication requirements defined for the sender",
		}
		pairs := [][]string{
			[]string{" is not allowed to send mail.", "_401"},
			[]string{"is not allowed to send from <", " per it's spf record"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["AuthFailure"] = func(fo *sis.Fact) bool {
		if fo == nil                                       { return false }
		if fo.Reason == "authfailure"                      { return true  }
		if status.Name(fo.DeliveryStatus) == "authfailure" { return true  }
		return IncludedIn["AuthFailure"](strings.ToLower(fo.DiagnosticCode))
	}
}

