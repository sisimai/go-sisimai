// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____              _             _____                     
// / ___| _   _ _ __ | |_ __ ___  _| ____|_ __ _ __ ___  _ __ 
// \___ \| | | | '_ \| __/ _` \ \/ /  _| | '__| '__/ _ \| '__|
//  ___) | |_| | | | | || (_| |>  <| |___| |  | | | (_) | |   
// |____/ \__, |_| |_|\__\__,_/_/\_\_____|_|  |_|  \___/|_|   
//        |___/                                               

package reason
import "strconv"
import "libsisimai.org/sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["SyntaxError"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		return false
	}

	// The bounce reason is "syntaxerror" or not
	ProbesInto["SyntaxError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is syntaxerror, false: is not syntaxerror
		if fo        == nil           { return false }
		if fo.Reason == "syntaxerror" { return true  }

		reply, nyaan := strconv.ParseUint(fo.ReplyCode, 10, 16); if nyaan != nil { return false }
		if (reply > 400 && reply < 408) || (reply > 500 && reply < 508)          { return true  }
		return false
	}
}

