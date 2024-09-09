// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      ___  ______  ____ ___ 
//  _ __| |__   ___  ___| |_   / / |/ /  _ \|  _ \_ _|
// | '__| '_ \ / _ \/ __| __| / /| ' /| | | | | | | | 
// | |  | | | | (_) \__ \ |_ / / | . \| |_| | |_| | | 
// |_|  |_| |_|\___/|___/\__/_/  |_|\_\____/|____/___|
import "strings"
import "sisimai/sis"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["KDDI"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		if fo.DiagnosticCode == "" { return "" }

		messagesof := map[string]string{
			"filtered":    "550 : user unknown", // The response was: 550 : User unknown
			"userunknown": ">: user unknown",    // The response was: 550 <...>: User unknown
		}
		issuedcode := strings.ToLower(fo.DiagnosticCode)
		reasontext := ""

		for e := range messagesof {
			// The key name is a bounce reason name
			if strings.Contains(issuedcode, messagesof[e]) == false { continue }
			reasontext = e; break
		}
		return reasontext
	}
}

