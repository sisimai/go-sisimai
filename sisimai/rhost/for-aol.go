// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      ___         _ 
//  _ __| |__   ___  ___| |_   / / \   ___ | |
// | '__| '_ \ / _ \/ __| __| / / _ \ / _ \| |
// | |  | | | | (_) \__ \ |_ / / ___ \ (_) | |
// |_|  |_| |_|\___/|___/\__/_/_/   \_\___/|_|
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["Aol"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      Aol Mail: https://www.aol.com
		if fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"hostunknown": []string{"Host or domain name not found"},
			"notaccept":   []string{"type=MX: Malformed or unexpected name server reply"},
		}

		reasontext := ""
		for e := range messagesof {
			// Each key is an error reason name
			if sisimoji.ContainsAny(fo.DiagnosticCode, messagesof[e]) == false { continue }
			reasontext = e; break
		}
		return reasontext
	}
}

