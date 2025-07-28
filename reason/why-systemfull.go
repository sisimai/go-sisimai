// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____            _                 _____      _ _ 
// / ___| _   _ ___| |_ ___ _ __ ___ |  ___|   _| | |
// \___ \| | | / __| __/ _ \ '_ ` _ \| |_ | | | | | |
//  ___) | |_| \__ \ ||  __/ | | | | |  _|| |_| | | |
// |____/ \__, |___/\__\___|_| |_| |_|_|   \__,_|_|_|
//        |___/                                      

package reason
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn["SystemFull"] = func(argv1 string) bool {
		if argv1 == "" { return false }

		index := []string{
			"mail system full",
			"requested mail action aborted: exceeded storage allocation", // MS Exchange
		}
		return moji.ContainsAny(argv1, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["SystemFull"] = func(fo *sis.Fact) bool { return false }
}

