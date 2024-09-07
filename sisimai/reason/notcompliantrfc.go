// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _       _    ____                      _ _             _   ____  _____ ____ 
// | \ | | ___ | |_ / ___|___  _ __ ___  _ __ | (_) __ _ _ __ | |_|  _ \|  ___/ ___|
// |  \| |/ _ \| __| |   / _ \| '_ ` _ \| '_ \| | |/ _` | '_ \| __| |_) | |_ | |    
// | |\  | (_) | |_| |__| (_) | | | | | | |_) | | | (_| | | | | |_|  _ <|  _|| |___ 
// |_| \_|\___/ \__|\____\___/|_| |_| |_| .__/|_|_|\__,_|_| |_|\__|_| \_\_|   \____|
//                                      |_|                                         
import "strings"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["NotCompliantRFC"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			"this message is not rfc 5322 compliant",
			"https://support.google.com/mail/?p=rfcmessagenoncompliant",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "notcompliantrfc" or not
	Truth["NotCompliantRFC"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is notcompliantrfc, false: is not notcompliantrfc
		if fo.Reason == "notcompliantrfc" { return true }
		return Match["NotCompliantRFC"](strings.ToLower(fo.DiagnosticCode))
	}
}

