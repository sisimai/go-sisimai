// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  __  __                _____           ____  _       
// |  \/  | ___  ___  __ |_   _|__   ___ | __ )(_) __ _ 
// | |\/| |/ _ \/ __|/ _` || |/ _ \ / _ \|  _ \| |/ _` |
// | |  | |  __/\__ \ (_| || | (_) | (_) | |_) | | (_| |
// |_|  |_|\___||___/\__, ||_|\___/ \___/|____/|_|\__, |
//                   |___/                        |___/ 
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["MesgTooBig"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		index := []string{
			"exceeded maximum inbound message size",
			"line limit exceeded",
			"max message size exceeded",
			"message file too big",
			"message length exceeds administrative limit",
			"message size exceeds fixed limit",
			"message size exceeds fixed maximum message size",
			"message size exceeds maximum value",
			"message too big",
			"message too large for this ",
			"size limit",
			"taille limite du message atteinte",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "mesgtoobig" or not
	ProbesInto["MesgTooBig"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is mesgtoobig, false: is not mesgtoobig
		if fo.Reason == "mesgtoobig"                      { return true }

		statuscode := fo.DeliveryStatus
		tempreason := status.Name(statuscode)

		// Delivery status code points "mesgtoobig".
		// Status: 5.3.4
		// Diagnostic-Code: SMTP; 552 5.3.4 Error: message file too big
		// Diagnostic-Code: SMTP; 552 5.2.3 Message length exceeds administrative limit
		if tempreason == "mesgtoobig"                           { return true }
		if tempreason == "exceedlimit" || statuscode == "5.2.3" { return false }
		return IncludedIn["MesgTooBig"](strings.ToLower(fo.DiagnosticCode))
	}
}

