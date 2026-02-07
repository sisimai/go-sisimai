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
import "libsisimai.org/sisimai/v5/smtp/command"

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
			"domain isn't in my list of allowed rcpthost",
			"email address is not verified.",
			"insecure mail relay",
			"no relaying",
			"not a gateway",
			"not local host",
			"open relay",
			"relay not permitted",
			"relay prohibition",
			"relaying denied", // Sendmail
			"relaying mail to ",
			"send to a non-local e-mail address", // MailEnable
			"specified domain is not allowed",
			"unable to relay ",
			"we don't handle mail for",
		}
		pairs := [][]string{
			[]string{"relay ", "denied"},
			[]string{"n", "t ", "to relay"},
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
		if slices.Contains(command.BeforeRCPT, fo.Command) == true               { return false }
		if slices.Contains([]string{eb.ReSAFE, eb.RePROC, eb.Re___0}, fo.Reason) { return false }
		return IncludedIn[eb.RePASS](strings.ToLower(fo.DiagnosticCode))
	}
}

