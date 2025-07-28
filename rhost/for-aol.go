// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ___         _ 
//  _ __| |__   ___  ___| |_   / / \   ___ | |
// | '__| '_ \ / _ \/ __| __| / / _ \ / _ \| |
// | |  | | | | (_) \__ \ |_ / / ___ \ (_) | |
// |_|  |_| |_|\___/|___/\__/_/_/   \_\___/|_|

package rhost
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress.
	//   Returns:
	//     - (string): Bounce reason name or an empty string.
	ReturnedBy["Aol"] = func(fo *sis.Fact) string {
		// - Aol Mail: https://www.aol.com
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"hostunknown": []string{"Host or domain name not found"},
			"notaccept":   []string{"type=MX: Malformed or unexpected name server reply"},
		}

		for e := range messagesof {
			// Each key is an error reason name
			if moji.ContainsAny(fo.DiagnosticCode, messagesof[e]) { return e }
		}
		return ""
	}
}

