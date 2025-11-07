// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ___  ______  ____ ___ 
//  _ __| |__   ___  ___| |_   / / |/ /  _ \|  _ \_ _|
// | '__| '_ \ / _ \/ __| __| / /| ' /| | | | | | | | 
// | |  | | | | (_) \__ \ |_ / / | . \| |_| | |_| | | 
// |_|  |_| |_|\___/|___/\__/_/  |_|\_\____/|____/___|

package rhost
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (string): Bounce reason name or an empty string.
	ReturnedBy["KDDI"] = func(fo *siba.Fact) string {
		// - https://www.au.com/support/service/internet/trouble/mail/01/
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string]string{
			eb.ReFILT: "550 : user unknown", // The response was: 550 : User unknown
			eb.ReUSER: ">: user unknown",    // The response was: 550 <...>: User unknown
		}
		issuedcode := strings.ToLower(fo.DiagnosticCode); for e := range messagesof {
			// The key name is a bounce reason name
			if strings.Contains(issuedcode, messagesof[e]) { return e }
		}
		return ""
	}
}

