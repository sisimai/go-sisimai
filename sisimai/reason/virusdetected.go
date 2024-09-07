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
	// Try to match that the given text and message patterns
	Match["VirusDetected"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
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
	Truth["VirusDetected"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is virusdetected, false: is not virusdetected
		if fo.Reason == "virusdetected"                         { return true  }
		if fo.SMTPCommand == "CONN"                             { return false }
		if fo.SMTPCommand == "EHLO" || fo.SMTPCommand == "HELO" { return false }
		if fo.SMTPCommand == "MAIL" || fo.SMTPCommand == "RCPT" { return false }
		return Match["VirusDetected"](strings.ToLower(fo.DiagnosticCode))
	}
}

