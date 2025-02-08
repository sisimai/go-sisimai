// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____                       _ _     _           _ _   
// | ____|_  _____ ___  ___  __| | |   (_)_ __ ___ (_) |_ 
// |  _| \ \/ / __/ _ \/ _ \/ _` | |   | | '_ ` _ \| | __|
// | |___ >  < (_|  __/  __/ (_| | |___| | | | | | | | |_ 
// |_____/_/\_\___\___|\___|\__,_|_____|_|_| |_| |_|_|\__|
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["ExceedLimit"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }
		index := []string{"message header size exceeds limit", "message too large"}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "exceedlimit" or not
	ProbesInto["ExceedLimit"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is exceedlimit, false: is not exceedlimit

		// Status: 5.2.3
		// Diagnostic-Code: SMTP; 552 5.2.3 Message size exceeds fixed maximum message size
		if fo == nil                                       { return false }
		if fo.Reason == "exceedlimit"                      { return true  }
		if status.Name(fo.DeliveryStatus) == "exceedlimit" { return true  }
		return IncludedIn["ExceedLimit"](strings.ToLower(fo.DiagnosticCode))
	}
}

