// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _       _    ____                      _ _             _   ____  _____ ____ 
// | \ | | ___ | |_ / ___|___  _ __ ___  _ __ | (_) __ _ _ __ | |_|  _ \|  ___/ ___|
// |  \| |/ _ \| __| |   / _ \| '_ ` _ \| '_ \| | |/ _` | '_ \| __| |_) | |_ | |    
// | |\  | (_) | |_| |__| (_) | | | | | | |_) | | | (_| | | | | |_|  _ <|  _|| |___ 
// |_| \_|\___/ \__|\____\___/|_| |_| |_| .__/|_|_|\__,_|_| |_|\__|_| \_\_|   \____|
//                                      |_|                                         

package reason
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool):         true if the argument includes one or more error message pattern
	IncludedIn["NotCompliantRFC"] = func(argv1 string) bool {
		if argv1 == "" { return false }

		index := []string{
			"duplicate header",
			"this message is not rfc 5322 compliant",
			"https://support.google.com/mail/?p=rfcmessagenoncompliant",
		}
		pairs := [][]string{
			[]string{" multiple ", " header"},
		}
		return moji.ContainsAny(argv1, index) || moji.AlignedAny(argv1, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress
	//   Returns:
	//     - (bool):         true if a reason is the reason defined in this file
	ProbesInto["NotCompliantRFC"] = func(fo *sis.Fact) bool {
		if fo        == nil               { return false }
		if fo.Reason == "notcompliantrfc" { return true  }
		return IncludedIn["NotCompliantRFC"](strings.ToLower(fo.DiagnosticCode))
	}
}

