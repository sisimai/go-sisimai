// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//   ___        _   _       _     _ 
//  / _ \ _ __ | | | | ___ | | __| |
// | | | | '_ \| |_| |/ _ \| |/ _` |
// | |_| | | | |  _  | (_) | | (_| |
//  \___/|_| |_|_| |_|\___/|_|\__,_|

package reason
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn["OnHold"] = func(mesg string) bool { return false }

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["OnHold"] = func(fo *sis.Fact) bool {
		if fo        == nil                           { return false }
		if fo.Reason == "onhold"                      { return true  }
		if status.Name(fo.DeliveryStatus) == "onhold" { return true  }
		return false
	}
}

