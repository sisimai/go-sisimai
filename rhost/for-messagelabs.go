// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ____  __                                _          _         
//  _ __| |__   ___  ___| |_   / /  \/  | ___  ___ ___  __ _  __ _  ___| |    __ _| |__  ___ 
// | '__| '_ \ / _ \/ __| __| / /| |\/| |/ _ \/ __/ __|/ _` |/ _` |/ _ \ |   / _` | '_ \/ __|
// | |  | | | | (_) \__ \ |_ / / | |  | |  __/\__ \__ \ (_| | (_| |  __/ |__| (_| | |_) \__ \
// |_|  |_| |_|\___/|___/\__/_/  |_|  |_|\___||___/___/\__,_|\__, |\___|_____\__,_|_.__/|___/
//                                                           |___/                           

package rhost
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/moji"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress
	//   Returns:
	//     - (string):       Bounce reason name or an empty string
	ReturnedBy["MessageLabs"] = func(fo *sis.Fact) string {
		// - https://www.broadcom.com/products/cybersecurity/email
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"securityerror": []string{"Please turn on SMTP Authentication in your mail client"},
			"userunknown":   []string{"542 ", " Rejected", "No such user"},
		}

		for e := range messagesof {
			// Each key is an error reason name
			if moji.ContainsAny(fo.DiagnosticCode, messagesof[e]) { return e }
		}
		return ""
	}
}

