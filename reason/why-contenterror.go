// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//   ____            _             _   _____                     
//  / ___|___  _ __ | |_ ___ _ __ | |_| ____|_ __ _ __ ___  _ __ 
// | |   / _ \| '_ \| __/ _ \ '_ \| __|  _| | '__| '__/ _ \| '__|
// | |__| (_) | | | | ||  __/ | | | |_| |___| |  | | | (_) | |   
//  \____\___/|_| |_|\__\___|_| |_|\__|_____|_|  |_|  \___/|_|   

package reason
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/smtp/status"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["ContentError"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"duplicate header",
			"header size exceeds maximum permitted",
			"improper use of 8-bit data in message header",
			"message header size, or recipient list, exceeds policy limit",
			"message mime complexity exceeds the policy maximum",
			"routing loop detected -- too many received: headers",
			"this message contain invalid mime headers",
			"this message contain improperly-formatted binary content",
			"this message contain text that uses unnecessary base64 encoding",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "contenterror" or not
	ProbesInto["ContentError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is contenterror, false: is not contenterror
		if fo == nil                                        { return false }
		if fo.Reason == "contenterror"                      { return true  }
		if ProbesInto["SpamDetected"](fo) == true           { return false }
		if status.Name(fo.DeliveryStatus) == "contenterror" { return true  }
		return IncludedIn["ContentError"](strings.ToLower(fo.DiagnosticCode))
	}
}

