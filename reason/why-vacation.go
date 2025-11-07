// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
// __     __              _   _             
// \ \   / /_ _  ___ __ _| |_(_) ___  _ __  
//  \ \ / / _` |/ __/ _` | __| |/ _ \| '_ \ 
//   \ V / (_| | (_| (_| | |_| | (_) | | | |
//    \_/ \__,_|\___\__,_|\__|_|\___/|_| |_|

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
	IncludedIn[eb.ReAWAY] = func(mesg string) bool {
		if mesg == "" { return false }
		index := []string{"i am away on vacation", "i am away until", "i am out of the office", "i will be traveling for work on"}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReAWAY] = func(fo *siba.Fact) bool { return false }
}

