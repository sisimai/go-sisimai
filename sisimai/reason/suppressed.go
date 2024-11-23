// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____                                               _ 
// / ___| _   _ _ __  _ __  _ __ ___  ___ ___  ___  __| |
// \___ \| | | | '_ \| '_ \| '__/ _ \/ __/ __|/ _ \/ _` |
//  ___) | |_| | |_) | |_) | | |  __/\__ \__ \  __/ (_| |
// |____/ \__,_| .__/| .__/|_|  \___||___/___/\___|\__,_|
//             |_|   |_|                                 
import "strings"
import "sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["Suppresed"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		return false
	}

	// The bounce reason is "suppresed" or not
	ProbesInto["Suppressed"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is suppressed, false: is not suppressed
		if fo.Reason == "suppressed" { return true }
		return IncludedIn["Suppressed"](strings.ToLower(fo.DiagnosticCode))
	}
}

