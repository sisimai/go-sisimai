// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____              _             _____                     
// / ___| _   _ _ __ | |_ __ ___  _| ____|_ __ _ __ ___  _ __ 
// \___ \| | | | '_ \| __/ _` \ \/ /  _| | '__| '__/ _ \| '__|
//  ___) | |_| | | | | || (_| |>  <| |___| |  | | | (_) | |   
// |____/ \__, |_| |_|\__\__,_/_/\_\_____|_|  |_|  \___/|_|   
//        |___/                                               
import "strconv"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["SyntaxError"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		return false
	}

	// The bounce reason is "syntaxerror" or not
	Truth["SyntaxError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is syntaxerror, false: is not syntaxerror
		if fo.Reason == "syntaxerror" { return true }

		reply, nyaan := strconv.ParseUint(fo.ReplyCode, 10, 16); if nyaan != nil { return false }
		if reply > 400 && reply < 408 { return true }
		if reply > 500 && reply < 508 { return true }
		return false
	}
}

