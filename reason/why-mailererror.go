// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  __  __       _ _           _____                     
// |  \/  | __ _(_) | ___ _ __| ____|_ __ _ __ ___  _ __ 
// | |\/| |/ _` | | |/ _ \ '__|  _| | '__| '__/ _ \| '__|
// | |  | | (_| | | |  __/ |  | |___| |  | | | (_) | |   
// |_|  |_|\__,_|_|_|\___|_|  |_____|_|  |_|  \___/|_|   

package reason
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool):         true if the argument includes one or more error message pattern
	IncludedIn["MailerError"] = func(argv1 string) bool {
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
		return moji.ContainsAny(argv1, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress
	//   Returns:
	//     - (bool):         true if a reason is the reason defined in this file
	ProbesInto["MailerError"] = func(fo *sis.Fact) bool { return false }
}

