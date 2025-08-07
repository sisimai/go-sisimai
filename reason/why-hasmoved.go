// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _           __  __                    _ 
// | | | | __ _ ___|  \/  | _____   _____  __| |
// | |_| |/ _` / __| |\/| |/ _ \ \ / / _ \/ _` |
// |  _  | (_| \__ \ |  | | (_) \ V /  __/ (_| |
// |_| |_|\__,_|___/_|  |_|\___/ \_/ \___|\__,_|

package reason
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern
	IncludedIn["HasMoved"] = func(mesg string) bool {
		if mesg == "" { return false }
		index := []string{" has been replaced by "}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["HasMoved"] = func(fo *sis.Fact) bool {
		if fo        == nil        { return false }
		if fo.Reason == "hasmoved" { return true  }
		return IncludedIn["HasMoved"](strings.ToLower(fo.DiagnosticCode))
	}
}

