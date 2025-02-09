// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      _____        _   _             _    
//  _ __| |__   ___  ___| |_   / / _ \ _   _| |_| | ___   ___ | | __
// | '__| '_ \ / _ \/ __| __| / / | | | | | | __| |/ _ \ / _ \| |/ /
// | |  | | | | (_) \__ \ |_ / /| |_| | |_| | |_| | (_) | (_) |   < 
// |_|  |_| |_|\___/|___/\__/_/  \___/ \__,_|\__|_|\___/ \___/|_|\_\
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["Outlook"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      https://technet.microsoft.com/en-us/library/bb232118
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"hostunknown": []string{"The mail could not be delivered to the recipient because the domain is not reachable"},
			"userunknown": []string{"Requested action not taken: mailbox unavailable"},
		}
		for e := range messagesof {
			// Each key is an error reason name
			if sisimoji.ContainsAny(fo.DiagnosticCode, messagesof[e]) { return e }
		}
		return ""
	}
}

