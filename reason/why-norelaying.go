// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _       ____      _             _             
// | \ | | ___ |  _ \ ___| | __ _ _   _(_)_ __   __ _ 
// |  \| |/ _ \| |_) / _ \ |/ _` | | | | | '_ \ / _` |
// | |\  | (_) |  _ <  __/ | (_| | |_| | | | | | (_| |
// |_| \_|\___/|_| \_\___|_|\__,_|\__, |_|_| |_|\__, |
//                                |___/         |___/ 

package reason
import "slices"
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
	IncludedIn[eb.RePASS] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"as a relay",
			"email address is not verified.",
			"insecure mail relay",
			"is not permitted to relay through this server without authentication",
			"mail server requires authentication when attempting to send to a non-local e-mail address", // MailEnable
			"no relaying",
			"not a gateway",
			"not allowed to relay through this machine",
			"not an open relay, so get lost",
			"not local host",
			"relay not permitted",
			"relaying denied", // Sendmail
			"relaying mail to ",
			"specified domain is not allowed",
			"that domain isn't in my list of allowed rcpthost",
			"this system is not configured to relay mail",
			"unable to relay ",
			"we don't handle mail for",
		}
		pairs := [][]string{
			[]string{"relay ", "denied"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.RePASS] = func(fo *siba.Fact) bool {
		if fo         == nil       { return false }
		if fo.Reason  == eb.RePASS { return true  }
		if slices.Contains([]string{eb.ReSAFE, eb.RePROC, eb.Re___0}, fo.Reason) { return false }
		if slices.Contains([]string{eb.CeCONN, eb.CeEHLO, eb.CeHELO}, fo.Command){ return false }
		return IncludedIn[eb.RePASS](strings.ToLower(fo.DiagnosticCode))
	}
}

