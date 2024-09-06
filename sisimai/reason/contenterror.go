// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//   ____            _             _   _____                     
//  / ___|___  _ __ | |_ ___ _ __ | |_| ____|_ __ _ __ ___  _ __ 
// | |   / _ \| '_ \| __/ _ \ '_ \| __|  _| | '__| '__/ _ \| '__|
// | |__| (_) | | | | ||  __/ | | | |_| |___| |  | | | (_) | |   
//  \____\___/|_| |_|\__\___|_| |_|\__|_____|_|  |_|  \___/|_|   
//                                                               
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"
import sisimoji "sisimai/string"

func init() {
	// Try to match that the given text and message patterns
	Match["ContentError"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
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
	Truth["ContentError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is contenterror, false: is not contenterror
		return false
	}
}

