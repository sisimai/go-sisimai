// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//   ____            _             _   _____                     
//  / ___|___  _ __ | |_ ___ _ __ | |_| ____|_ __ _ __ ___  _ __ 
// | |   / _ \| '_ \| __/ _ \ '_ \| __|  _| | '__| '__/ _ \| '__|
// | |__| (_) | | | | ||  __/ | | | |_| |___| |  | | | (_) | |   
//  \____\___/|_| |_|\__\___|_| |_|\__|_____|_|  |_|  \___/|_|   

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReBODY] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"header size exceeds maximum permitted",
			"improper use of 8-bit data in message header",
			"message contain invalid mime headers",
			"message contain improperly-formatted binary content",
			"message contain text that uses unnecessary base64 encoding",
			"message header size, or recipient list, exceeds policy limit",
			"message mime complexity exceeds the policy maximum",
			"routing loop detected -- too many received: headers",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file
	ProbesInto[eb.ReBODY] = func(fo *siba.Fact) bool {
		if fo == nil                                   { return false }
		if fo.Reason == eb.ReBODY                      { return true  }
		if ProbesInto[eb.ReSPAM](fo) == true           { return false }
		if status.Name(fo.DeliveryStatus) == eb.ReBODY { return true  }
		return IncludedIn[eb.ReBODY](strings.ToLower(fo.DiagnosticCode))
	}
}

