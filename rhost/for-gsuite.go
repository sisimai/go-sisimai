// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ______ ____        _ _       
//  _ __| |__   ___  ___| |_   / / ___/ ___| _   _(_) |_ ___ 
// | '__| '_ \ / _ \/ __| __| / / |  _\___ \| | | | | __/ _ \
// | |  | | | | (_) \__ \ |_ / /| |_| |___) | |_| | | ||  __/
// |_|  |_| |_|\___/|___/\__/_/  \____|____/ \__,_|_|\__\___|
// Google Workspace (formerly G Suite) https://workspace.google.com/

package rhost
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (string): Bounce reason name or an empty string.
	ReturnedBy["GSuite"] = func(fo *siba.Fact) string {
		// - https://workspace.google.com/
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			eb.ReHOST: []string{" responded with code NXDOMAIN", "Domain name not found"},
			eb.ReNETW: []string{" had no relevant answers.", "responded with code NXDOMAIN", "Domain name not found"},
			eb.Re00MX: []string{"Null MX"},
			eb.ReUSER: []string{"because the address couldn't be found. Check for typos or unnecessary spaces and try again."},
		}
		statuscode := ""; if fo.DeliveryStatus != "" { statuscode = string(fo.DeliveryStatus[0]) }
		esmtpreply := ""; if fo.ReplyCode      != "" { esmtpreply = string(fo.ReplyCode[0])      }

		for e := range messagesof {
			// The key is a bounce reason name
			if moji.ContainsAny(fo.DiagnosticCode, messagesof[e]) == false { continue }
			if e == eb.ReNETW && (statuscode == "5" || esmtpreply == "5")  { continue }
			if e == eb.ReHOST && (statuscode == "4" || statuscode == "")   { continue }
			if e == eb.ReHOST && (esmtpreply == "4" || esmtpreply == "")   { continue }
			return e
		}
		return ""
	}
}

