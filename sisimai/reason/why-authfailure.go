// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//     _         _   _     _____     _ _                
//    / \  _   _| |_| |__ |  ___|_ _(_) |_   _ _ __ ___ 
//   / _ \| | | | __| '_ \| |_ / _` | | | | | | '__/ _ \
//  / ___ \ |_| | |_| | | |  _| (_| | | | |_| | | |  __/
// /_/   \_\__,_|\__|_| |_|_|  \__,_|_|_|\__,_|_|  \___|
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"
import sisimoji "sisimai/string"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["AuthFailure"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"//spf.pobox.com",
			"bad spf records for",
			"dmarc policy",
			"please inspect your spf settings",
			"sender policy framework (spf) fail",
			"sender policy framework violation",
			"spf (sender policy framework) domain authentication fail",
			"spf check: fail",
		}
		pairs := [][]string{
			[]string{" is not allowed to send mail.", "_401"},
			[]string{"is not allowed to send from <", " per it's spf record"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "authfailure" or not
	ProbesInto["AuthFailure"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is authfailure, false: is not authfailure
		if fo.Reason == "authfailure"                      { return true }
		if status.Name(fo.DeliveryStatus) == "authfailure" { return true }
		return IncludedIn["AuthFailure"](strings.ToLower(fo.DiagnosticCode))
	}
}

