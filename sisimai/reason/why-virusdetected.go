// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

// __     ___                ____       _            _           _ 
// \ \   / (_)_ __ _   _ ___|  _ \  ___| |_ ___  ___| |_ ___  __| |
//  \ \ / /| | '__| | | / __| | | |/ _ \ __/ _ \/ __| __/ _ \/ _` |
//   \ V / | | |  | |_| \__ \ |_| |  __/ ||  __/ (__| ||  __/ (_| |
//    \_/  |_|_|   \__,_|___/____/ \___|\__\___|\___|\__\___|\__,_|
import "strings"
import "sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["VirusDetected"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		index := []string{
			"it has a potentially executable attachment",
			"the message was rejected because it contains prohibited virus or spam content",
			"this form of attachment has been used by recent viruses or other malware",
			"virus detected",
			"virus phishing/malicious_url detected",
			"your message was infected with a virus",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "virusdetected" or not
	ProbesInto["VirusDetected"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is virusdetected, false: is not virusdetected
		if fo.Reason == "virusdetected"                 { return true  }
		if fo.Command == "CONN"                         { return false }
		if fo.Command == "EHLO" || fo.Command == "HELO" { return false }
		if fo.Command == "MAIL" || fo.Command == "RCPT" { return false }
		return IncludedIn["VirusDetected"](strings.ToLower(fo.DiagnosticCode))
	}
}

