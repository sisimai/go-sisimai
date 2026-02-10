// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _       _      _                      _   
// | \ | | ___ | |_   / \   ___ ___ ___ _ __ | |_ 
// |  \| |/ _ \| __| / _ \ / __/ __/ _ \ '_ \| __|
// | |\  | (_) | |_ / ___ \ (_| (_|  __/ |_) | |_ 
// |_| \_|\___/ \__/_/   \_\___\___\___| .__/ \__|
//                                     |_|        

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
	IncludedIn[eb.Re00MX] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"destination seem to reject all mails", // OpenSMTPD/smtp/mta.c
			"no mx found for ",                     // OpenSMTPD/smtp/mta.c
			"does not accept mail",                 // Sendmail, iCloud
			"mail receiving disabled",
			"name server: .: host not found",       // Sendmail
			"no mx record found for domain=",       // Oath(Yahoo!)
			"no route for current request",
			"null mx",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.Re00MX] = func(fo *siba.Fact) bool {
		if fo        == nil                                      { return false }
		if fo.Reason == eb.Re00MX                                { return true  }
		if slices.Contains([]string{"521", "556"}, fo.ReplyCode) { return true  }
		if slices.Contains(command.BeforeRCPT, fo.Command)       { return false }
		return IncludedIn[eb.Re00MX](strings.ToLower(fo.DiagnosticCode))
	}
}

