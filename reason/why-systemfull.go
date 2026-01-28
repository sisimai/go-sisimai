// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____            _                 _____      _ _ 
// / ___| _   _ ___| |_ ___ _ __ ___ |  ___|   _| | |
// \___ \| | | / __| __/ _ \ '_ ` _ \| |_ | | | | | |
//  ___) | |_| \__ \ ||  __/ | | | | |  _|| |_| | | |
// |____/ \__, |___/\__\___|_| |_| |_|_|   \__,_|_|_|
//        |___/                                      

package reason
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReDISK] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"exceeded storage allocation", // MS Exchange
			"mail system full",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReDISK] = func(fo *siba.Fact) bool { return false }
}

