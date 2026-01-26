// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
// __     ___                ____       _            _           _ 
// \ \   / (_)_ __ _   _ ___|  _ \  ___| |_ ___  ___| |_ ___  __| |
//  \ \ / /| | '__| | | / __| | | |/ _ \ __/ _ \/ __| __/ _ \/ _` |
//   \ V / | | |  | |_| \__ \ |_| |  __/ ||  __/ (__| ||  __/ (_| |
//    \_/  |_|_|   \__,_|___/____/ \___|\__\___|\___|\__\___|\__,_|

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
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReEXEC] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"form of attachment has been used by recent viruses or other malware",
			"it has a potentially executable attachment",
			"virus detected",
			"virus phishing/malicious_url detected",
		}
		pairs := [][]string{
			[]string{"message was ", "ected", " virus"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReEXEC] = func(fo *siba.Fact) bool {
		if fo        == nil                               { return false }
		if fo.Reason == eb.ReEXEC                         { return true  }
		if slices.Contains(command.ExceptDATA, fo.Command){ return false }
		return IncludedIn[eb.ReEXEC](strings.ToLower(fo.DiagnosticCode))
	}
}

