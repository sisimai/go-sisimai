// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ______ ____        _ _       
//  _ __| |__   ___  ___| |_   / / ___/ ___| _   _(_) |_ ___ 
// | '__| '_ \ / _ \/ __| __| / / |  _\___ \| | | | | __/ _ \
// | |  | | | | (_) \__ \ |_ / /| |_| |___) | |_| | | ||  __/
// |_|  |_| |_|\___/|___/\__/_/  \____|____/ \__,_|_|\__\___|
// Google Workspace (formerly G Suite) https://workspace.google.com/

package rhost
import "libsisimai.org/sisimai/sis"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["GSuite"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      https://workspace.google.com/
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"hostunknown":  []string{" responded with code NXDOMAIN", "Domain name not found"},
			"networkerror": []string{" had no relevant answers.", "responded with code NXDOMAIN", "Domain name not found"},
			"notaccept":    []string{"Null MX"},
			"userunknown":  []string{"because the address couldn't be found. Check for typos or unnecessary spaces and try again."},
		}
		statuscode := ""; if fo.DeliveryStatus != "" { statuscode = string(fo.DeliveryStatus[0]) }
		esmtpreply := ""; if fo.ReplyCode      != "" { esmtpreply = string(fo.ReplyCode[0])      }

		for e := range messagesof {
			// The key is a bounce reason name
			if sisimoji.ContainsAny(fo.DiagnosticCode, messagesof[e]) == false { continue }
			if e == "networkerror" && (statuscode == "5" || esmtpreply == "5") { continue }
			if e == "hostunknown"  && (statuscode == "4" || statuscode == "")  { continue }
			if e == "hostunknown"  && (esmtpreply == "4" || esmtpreply == "")  { continue }
			return e
		}
		return ""
	}
}

