// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _       ____      _             _             
// | \ | | ___ |  _ \ ___| | __ _ _   _(_)_ __   __ _ 
// |  \| |/ _ \| |_) / _ \ |/ _` | | | | | '_ \ / _` |
// | |\  | (_) |  _ <  __/ | (_| | |_| | | | | | (_| |
// |_| \_|\___/|_| \_\___|_|\__,_|\__, |_|_| |_|\__, |
//                                |___/         |___/ 
import "strings"
import "sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["NoRelaying"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
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
			"relay access denied",
			"relay denied",
			"relaying mail to ",
			"relay not permitted",
			"relaying denied", // Sendmail
			"relaying mail to ",
			"specified domain is not allowed",
			"that domain isn't in my list of allowed rcpthost",
			"this system is not configured to relay mail",
			"unable to relay for",
			"we don't handle mail for",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "norelaying" or not
	ProbesInto["NoRelaying"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is norelaying, false: is not norelaying
		if fo.Reason == "norelaying" { return true }
		if fo.Reason == "securityerror" || fo.Reason == "systemerror" || fo.Reason == "undefined" { return false }
		if fo.SMTPCommand == "CONN"     || fo.SMTPCommand == "EHLO"   || fo.SMTPCommand == "HELO" { return false }
		return IncludedIn["NoRelaying"](strings.ToLower(fo.DiagnosticCode))
	}
}

