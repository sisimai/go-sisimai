// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                                               _ 
// / ___| _   _ _ __  _ __  _ __ ___  ___ ___  ___  __| |
// \___ \| | | | '_ \| '_ \| '__/ _ \/ __/ __|/ _ \/ _` |
//  ___) | |_| | |_) | |_) | | |  __/\__ \__ \  __/ (_| |
// |____/ \__,_| .__/| .__/|_|  \___||___/___/\___|\__,_|
//             |_|   |_|                                 

package reason
import "strings"
import "libsisimai.org/sisimai/v5/sis"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool):         true if the argument includes one or more error message pattern
	IncludedIn["Suppressed"] = func(argv1 string) bool { return false }

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress
	//   Returns:
	//     - (bool):         true if a reason is the reason defined in this file
	ProbesInto["Suppressed"] = func(fo *sis.Fact) bool {
		if fo        == nil          { return false }
		if fo.Reason == "suppressed" { return true  }
		return IncludedIn["Suppressed"](strings.ToLower(fo.DiagnosticCode))
	}
}

