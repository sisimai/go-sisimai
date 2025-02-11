// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _       _    ____                      _ _             _   ____  _____ ____ 
// | \ | | ___ | |_ / ___|___  _ __ ___  _ __ | (_) __ _ _ __ | |_|  _ \|  ___/ ___|
// |  \| |/ _ \| __| |   / _ \| '_ ` _ \| '_ \| | |/ _` | '_ \| __| |_) | |_ | |    
// | |\  | (_) | |_| |__| (_) | | | | | | |_) | | | (_| | | | | |_|  _ <|  _|| |___ 
// |_| \_|\___/ \__|\____\___/|_| |_| |_| .__/|_|_|\__,_|_| |_|\__|_| \_\_|   \____|
//                                      |_|                                         
import "strings"
import "libsisimai.org/sisimai/sis"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["NotCompliantRFC"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"duplicate header",
			"this message is not rfc 5322 compliant",
			"https://support.google.com/mail/?p=rfcmessagenoncompliant",
		}
		pairs := [][]string{
			[]string{" multiple ", " header"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "notcompliantrfc" or not
	ProbesInto["NotCompliantRFC"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is notcompliantrfc, false: is not notcompliantrfc
		if fo        == nil               { return false }
		if fo.Reason == "notcompliantrfc" { return true  }
		return IncludedIn["NotCompliantRFC"](strings.ToLower(fo.DiagnosticCode))
	}
}

