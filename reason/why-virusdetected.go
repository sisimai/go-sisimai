// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
// __     ___                ____       _            _           _ 
// \ \   / (_)_ __ _   _ ___|  _ \  ___| |_ ___  ___| |_ ___  __| |
//  \ \ / /| | '__| | | / __| | | |/ _ \ __/ _ \/ __| __/ _ \/ _` |
//   \ V / | | |  | |_| \__ \ |_| |  __/ ||  __/ (__| ||  __/ (_| |
//    \_/  |_|_|   \__,_|___/____/ \___|\__\___|\___|\__\___|\__,_|

package reason
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/command"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn["VirusDetected"] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"it has a potentially executable attachment",
			"the message was rejected because it contains prohibited virus or spam content",
			"this form of attachment has been used by recent viruses or other malware",
			"virus detected",
			"virus phishing/malicious_url detected",
			"your message was infected with a virus",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["VirusDetected"] = func(fo *siba.Fact) bool {
		if fo        == nil                                { return false }
		if fo.Reason == "virusdetected"                    { return true  }
		if slices.Contains(command.ExceptDATA, fo.Command) { return false }
		return IncludedIn["VirusDetected"](strings.ToLower(fo.DiagnosticCode))
	}
}

