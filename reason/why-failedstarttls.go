// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____     _ _          _ ____ _____  _    ____ _____ _____ _     ____  
// |  ___|_ _(_) | ___  __| / ___|_   _|/ \  |  _ \_   _|_   _| |   / ___| 
// | |_ / _` | | |/ _ \/ _` \___ \ | | / _ \ | |_) || |   | | | |   \___ \ 
// |  _| (_| | | |  __/ (_| |___) || |/ ___ \|  _ < | |   | | | |___ ___) |
// |_|  \__,_|_|_|\___|\__,_|____/ |_/_/   \_\_| \_\|_|   |_| |_____|____/ 

package reason
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/sis"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool):         true if the argument includes one or more error message pattern
	IncludedIn["FailedSTARTTLS"] = func(argv1 string) bool {
		if argv1 == "" { return false }

		index := []string{
			"starttls is required to send mail",
			"tls required but not supported", // SendGrid:the recipient mailserver does not support TLS or have a valid certificate
		}
		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress
	//   Returns:
	//     - (bool):         true if a reason is the reason defined in this file
	ProbesInto["FailedSTARTTLS"] = func(fo *sis.Fact) bool {
		if fo == nil                                                    { return false }
		if fo.Reason == "failedstarttls" || fo.Command == "STARTTLS"    { return true  }
		if slices.Contains([]string{"523", "524", "538"}, fo.ReplyCode) { return true  }
		return IncludedIn["FailedSTARTTLS"](strings.ToLower(fo.DiagnosticCode))
	}
}

