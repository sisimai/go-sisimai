// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _           __  __                    _ 
// | | | | __ _ ___|  \/  | _____   _____  __| |
// | |_| |/ _` / __| |\/| |/ _ \ \ / / _ \/ _` |
// |  _  | (_| \__ \ |  | | (_) \ V /  __/ (_| |
// |_| |_|\__,_|___/_|  |_|\___/ \_/ \___|\__,_|

package reason
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/command"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern
	IncludedIn[eb.ReMOVE] = func(mesg string) bool {
		if mesg == "" { return false }
		index := []string{" has been replaced by "}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReMOVE] = func(fo *siba.Fact) bool {
		if fo        == nil                                { return false }
		if fo.Reason == eb.ReMOVE                          { return true  }
		if slices.Contains(command.BeforeRCPT, fo.Command) { return false }
		return IncludedIn[eb.ReMOVE](strings.ToLower(fo.DiagnosticCode))
	}
}

