// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _       _      _                      _   
// | \ | | ___ | |_   / \   ___ ___ ___ _ __ | |_ 
// |  \| |/ _ \| __| / _ \ / __/ __/ _ \ '_ \| __|
// | |\  | (_) | |_ / ___ \ (_| (_|  __/ |_) | |_ 
// |_| \_|\___/ \__/_/   \_\___\___\___| .__/ \__|
//                                     |_|        

package reason
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool):         true if the argument includes one or more error message pattern
	IncludedIn["NotAccept"] = func(argv1 string) bool {
		if argv1 == "" { return false }

		index := []string{
			"does not accept mail (nullmx)",
			"host/domain does not accept mail", // iCloud
			"host does not accept mail",        // Sendmail
			"mail receiving disabled",
			"name server: .: host not found",   // Sendmail
			"no mx record found for domain=",   // Oath(Yahoo!)
			"no route for current request",
			"smtp protocol returned a permanent error",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress
	//   Returns:
	//     - (bool):         true if a reason is the reason defined in this file
	ProbesInto["NotAccept"] = func(fo *sis.Fact) bool {
		if fo        == nil                                     { return false }
		if fo.Reason == "notaccept"                             { return true  }
		if moji.EqualsAny(fo.ReplyCode, []string{"521", "556"}) { return true  }
		return IncludedIn["NotAccept"](strings.ToLower(fo.DiagnosticCode))
	}
}

