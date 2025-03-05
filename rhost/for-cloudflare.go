// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.

//       _               _      ______ _                 _  __ _                
//  _ __| |__   ___  ___| |_   / / ___| | ___  _   _  __| |/ _| | __ _ _ __ ___ 
// | '__| '_ \ / _ \/ __| __| / / |   | |/ _ \| | | |/ _` | |_| |/ _` | '__/ _ \
// | |  | | | | (_) \__ \ |_ / /| |___| | (_) | |_| | (_| |  _| | (_| | | |  __/
// |_|  |_| |_|\___/|___/\__/_/  \____|_|\___/ \__,_|\__,_|_| |_|\__,_|_|  \___|

package rhost
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/moji"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["Cloudflare"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      Email Routing   https://developers.cloudflare.com/email-routing/postmaster/
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"blocked":     []string{"found on one or more DNSBLs"},
			"systemerror": []string{"Upstream error"},
		}

		for e := range messagesof {
			// Each key is an error reason name
			if moji.ContainsAny(fo.DiagnosticCode, messagesof[e]) { return e }
		}
		return ""
	}
}

