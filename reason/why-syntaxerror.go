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
import "libsisimai.org/sisimai/v5/sis"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn["SyntaxError"] = func(mesg string) bool { return false }

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["SyntaxError"] = func(fo *sis.Fact) bool {
		if fo        == nil           { return false }
		if fo.Reason == "syntaxerror" { return true  }

		reply, nyaan := strconv.ParseUint(fo.ReplyCode, 10, 16); if nyaan != nil { return false }
		if (reply > 400 && reply < 408) || (reply > 500 && reply < 508)          { return true  }
		return false
	}
}

