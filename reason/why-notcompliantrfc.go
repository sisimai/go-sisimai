// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _       _    ____                      _ _             _   ____  _____ ____ 
// | \ | | ___ | |_ / ___|___  _ __ ___  _ __ | (_) __ _ _ __ | |_|  _ \|  ___/ ___|
// |  \| |/ _ \| __| |   / _ \| '_ ` _ \| '_ \| | |/ _` | '_ \| __| |_) | |_ | |    
// | |\  | (_) | |_| |__| (_) | | | | | | |_) | | | (_| | | | | |_|  _ <|  _|| |___ 
// |_| \_|\___/ \__|\____\___/|_| |_| |_| .__/|_|_|\__,_|_| |_|\__|_| \_\_|   \____|
//                                      |_|                                         

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReNRFC] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"duplicate header",
			"message is not rfc 5322 compliant",
			"rfc 1035 violation: recursive cname records for",
			"https://support.google.com/mail/?p=rfcmessagenoncompliant",
		}
		pairs := [][]string{[]string{" multiple ", " header"}}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReNRFC] = func(fo *siba.Fact) bool {
		if fo        == nil       { return false }
		if fo.Reason == eb.ReNRFC { return true  }
		return IncludedIn[eb.ReNRFC](strings.ToLower(fo.DiagnosticCode))
	}
}

