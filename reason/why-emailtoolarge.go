// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____                 _ _ _____           _                         
// | ____|_ __ ___   __ _(_) |_   _|__   ___ | |    __ _ _ __ __ _  ___ 
// |  _| | '_ ` _ \ / _` | | | | |/ _ \ / _ \| |   / _` | '__/ _` |/ _ \
// | |___| | | | | | (_| | | | | | (_) | (_) | |__| (_| | | | (_| |  __/
// |_____|_| |_| |_|\__,_|_|_| |_|\___/ \___/|_____\__,_|_|  \__, |\___|
//                                                           |___/      

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
			"message size exceeds fixed ",
			"message size exceeds maximum value",
			"message too big",
			"message too large",
			"size limit",
			"taille limite du message atteinte",
		}
		pairs := [][]string{
			[]string{"message ", " exceeds ", "limit"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReSIZE] = func(fo *siba.Fact) bool {
		// Delivery status code points EmailTooLarge.
		// Status: 5.3.4
		// Diagnostic-Code: SMTP; 552 5.3.4 Error: message file too big
		// Diagnostic-Code: SMTP; 552 5.2.3 Message length exceeds administrative limit
		if fo        == nil       { return false }
		if fo.Reason == eb.ReSIZE { return true  }

		if status.Name(fo.DeliveryStatus) == eb.ReSIZE { return true }
		return IncludedIn[eb.ReSIZE](strings.ToLower(fo.DiagnosticCode))
	}
}

