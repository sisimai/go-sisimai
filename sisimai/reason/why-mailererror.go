// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  __  __       _ _           _____                     
// |  \/  | __ _(_) | ___ _ __| ____|_ __ _ __ ___  _ __ 
// | |\/| |/ _` | | |/ _ \ '__|  _| | '__| '__/ _ \| '__|
// | |  | | (_| | | |  __/ |  | |___| |  | | | (_) | |   
// |_|  |_|\__,_|_|_|\___|_|  |_____|_|  |_|  \___/|_|   
import "strings"
import "sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["MailerError"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			" || exit ",
			"procmail: ",
			"bin/procmail",
			"bin/maidrop",
			"command failed: ",
			"command died with status ",
			"command output:",
			"mailer error",
			"pipe to |/",
			"x-unix; ",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "mailererror" or not
	ProbesInto["MailerError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is mailererror, false: is not mailererror
		return false
	}
}

