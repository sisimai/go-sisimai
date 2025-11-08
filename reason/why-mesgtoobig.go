// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  __  __                _____           ____  _       
// |  \/  | ___  ___  __ |_   _|__   ___ | __ )(_) __ _ 
// | |\/| |/ _ \/ __|/ _` || |/ _ \ / _ \|  _ \| |/ _` |
// | |  | |  __/\__ \ (_| || | (_) | (_) | |_) | | (_| |
// |_|  |_|\___||___/\__, ||_|\___/ \___/|____/|_|\__, |
//                   |___/                        |___/ 

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
	IncludedIn[eb.ReSIZE] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"exceeded maximum inbound message size",
			"exceeded the maximum incoming message size",
			"line limit exceeded",
			"max message size exceeded",
			"message file too big",
			"message length exceeds administrative limit",
			"message size exceeds fixed limit",
			"message size exceeds fixed maximum message size",
			"message size exceeds maximum value",
			"message too big",
			"message too large for this ",
			"size limit",
			"taille limite du message atteinte",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReSIZE] = func(fo *siba.Fact) bool {
		// Delivery status code points MesgTooBig.
		// Status: 5.3.4
		// Diagnostic-Code: SMTP; 552 5.3.4 Error: message file too big
		// Diagnostic-Code: SMTP; 552 5.2.3 Message length exceeds administrative limit
		if fo        == nil       { return false }
		if fo.Reason == eb.ReSIZE { return true  }

		tempreason    := status.Name(fo.DeliveryStatus)
		if tempreason == eb.ReSIZE                                 { return true  }
		if tempreason == eb.ReXLIM || fo.DeliveryStatus == "5.2.3" { return false }
		return IncludedIn[eb.ReSIZE](strings.ToLower(fo.DiagnosticCode))
	}
}

