// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                      _ _             
// / ___| _ __   ___  ___  __| (_)_ __   __ _ 
// \___ \| '_ \ / _ \/ _ \/ _` | | '_ \ / _` |
//  ___) | |_) |  __/  __/ (_| | | | | | (_| |
// |____/| .__/ \___|\___|\__,_|_|_| |_|\__, |
//       |_|                            |___/ 

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
	IncludedIn[eb.ReFAST] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"mail sent from your IP address has been temporarily rate limited",
			"please try again slower",
			"receiving mail at a rate that prevents additional messages from being delivered",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReFAST] = func(fo *siba.Fact) bool {
		// Action: failed
		// Status: 4.7.1
		// Remote-MTA: dns; smtp.example.jp
		// Diagnostic-Code: smtp; 451 4.7.1 <mx.example.org[192.0.2.2]>: Client host rejected: Please try again slower
		if fo        == nil       { return false }
		if fo.Reason == eb.ReFAST { return true  }
		return IncludedIn[eb.ReFAST](strings.ToLower(fo.DiagnosticCode))
	}
}

