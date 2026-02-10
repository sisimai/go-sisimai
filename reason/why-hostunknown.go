// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _           _   _   _       _                              
// | | | | ___  ___| |_| | | |_ __ | | ___ __   _____      ___ __  
// | |_| |/ _ \/ __| __| | | | '_ \| |/ / '_ \ / _ \ \ /\ / / '_ \ 
// |  _  | (_) \__ \ |_| |_| | | | |   <| | | | (_) \ V  V /| | | |
// |_| |_|\___/|___/\__|\___/|_| |_|_|\_\_| |_|\___/ \_/\_/ |_| |_|

package reason
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"
import "libsisimai.org/sisimai/v5/smtp/command"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReHOST] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"couldn't find any host ", // qmail-remote.c:78
			"dns server returned answer with no data",
			"domain is not reachable",
			"domain mentioned in email address is unknown",
			"domain must exist",
			"domain name not found",
			"host or domain name not found",
			"host unknown",
			"host unreachable",
			"illegal host/domain name found",
			"name or service not known",
			"no such domain",
			"recipient address rejected: unknown domain name",
			"responded with code nxdomain",
			"unknown host",
		}
		pairs := [][]string{
			[]string{"domain ", "not exist"},
			[]string{"host ", " not found"},
			[]string{"unrout", "able ", "address"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file
	ProbesInto[eb.ReHOST] = func(fo *siba.Fact) bool {
		if fo        == nil                                { return false }
		if fo.Reason == eb.ReHOST                          { return true  }
		if slices.Contains(command.BeforeRCPT, fo.Command) { return false }

		issuedcode := strings.ToLower(fo.DiagnosticCode)
		if status.Name(fo.DeliveryStatus) == eb.ReHOST {
			// To prevent classifying DNS errors as HostUnknown.
			if IncludedIn[eb.ReINET](issuedcode) == false { return true }

		} else {
			// Status: 5.1.2
			// Diagnostic-Code: SMTP; 550 Host unknown
			if IncludedIn[eb.ReHOST](issuedcode)  == true { return true }
		}
		return false
	}
}

