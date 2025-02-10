// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _           _   _   _       _                              
// | | | | ___  ___| |_| | | |_ __ | | ___ __   _____      ___ __  
// | |_| |/ _ \/ __| __| | | | '_ \| |/ / '_ \ / _ \ \ /\ / / '_ \ 
// |  _  | (_) \__ \ |_| |_| | | | |   <| | | | (_) \ V  V /| | | |
// |_| |_|\___/|___/\__|\___/|_| |_|_|\_\_| |_|\___/ \_/\_/ |_| |_|
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/smtp/status"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["HostUnknown"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"domain does not exist",
			"domain is not reachable",
			"domain must exist",
			"host or domain name not found",
			"host unknown",
			"host unreachable",
			"mail domain mentioned in email address is unknown",
			"name or service not known",
			"no such domain",
			"recipient address rejected: unknown domain name",
			"recipient domain must exist",
			"the account or domain may not exist",
			"unknown host",
			"unroutable address",
			"unrouteable address",
		}
		pairs := [][]string{
			[]string{"553 ", " does not exist"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "hostunknown" or not
	ProbesInto["HostUnknown"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is hostunknown, false: is not hostunknown
		if fo        == nil           { return false }
		if fo.Reason == "hostunknown" { return true  }

		issuedcode := strings.ToLower(fo.DiagnosticCode)
		if status.Name(fo.DeliveryStatus) == "hostunknown" {
			// To prevent classifying DNS errors as "HostUnknown"
			if IncludedIn["NetworkError"](issuedcode) == false { return true }

		} else {
			// Status: 5.1.2
			// Diagnostic-Code: SMTP; 550 Host unknown
			if IncludedIn["HostUnknown"](issuedcode)  == true  { return true }
		}
		return false
	}
}

