// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ __ ___  __ _ ___  ___  _ __  
// | '__/ _ \/ _` / __|/ _ \| '_ \ 
// | | |  __/ (_| \__ \ (_) | | | |
// |_|  \___|\__,_|___/\___/|_| |_|

package reason
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/smtp/status"

// Find detects and returns the bounce reason.
//   Arguments:
//     - fo (*sis.Fact): Decoded data in progress.
//   Returns:
//     - (string): Bounce reason name or an empty string.
func Find(fo *sis.Fact) string {
	// Return the reason text already decided except the reason matched with the name checked by
	// reason.ShouldBeRetried() function.
	if fo == nil { return "" }
	if fo.Reason != "" && ShouldBeRetried(fo.Reason) == false { return fo.Reason   }
	if strings.HasPrefix(fo.DeliveryStatus, "2.")    == true  { return "delivered" }

	reasontext := ""; if fo.DiagnosticType == "SMTP" || fo.DiagnosticType == "" {
		// Diagnostic-Code: SMTP; ... or empty value
		for _, e := range classorder[0] {
			// Check the values of Diagnostic-Code: and Status: fields using reason.ProbesInto[*]()
			// function of each child class in reason/why-*.go
			if ProbesInto[e](fo) { reasontext = strings.ToLower(e); break }
		}
	}

	if reasontext == "" || reasontext == "undefined" {
		// The bounce reason is not detected yet at the code block above
		reasontext  = anotherone(fo) // Try to find a reason name using anotherone()
		if reasontext == "undefined"                  { reasontext = ""          }
		if reasontext == "" && fo.Action == "delayed" { reasontext = "expired"   }
		if reasontext != ""                           { return reasontext        }

		issuedcode := strings.ToLower(fo.DiagnosticCode)
		if IncludedIn["Vacation"](issuedcode)         { reasontext = "vacation"  }
		if reasontext == "" && issuedcode != ""       { reasontext = "onhold"    }
		if reasontext == ""                           { reasontext = "undefined" }
	}
	return reasontext
}

// anotherone detects the other bounce reason, is a fall back function for Find().
//   Arguments:
//     - fo (*sis.Fact): Decoded data in progress.
//   Returns:
//     - (string): Bounce reason name or an empty string.
func anotherone(fo *sis.Fact) string {
	if fo                    == nil  { return ""        }
	if IsExplicit(fo.Reason) == true { return fo.Reason }

	issuedcode := strings.ToLower(fo.DiagnosticCode)
	reasontext := status.Name(fo.DeliveryStatus)

	if ShouldBeRetried(reasontext) || fo.DiagnosticType == "SMTP" {
		// - The value of the reason is not decided yet by the fo.DeliveryStatus.
		// - Try to find the bounce reason by the fo.DeliveryStatus or fo.DianosticCode when the
		//   ShouldBeRetried(reasontext) returns true or fo.DiagnosticType is "SMTP".
		for _, e := range classorder[1] {
			// Trying to match with other patterns in reason/why-*.go
			if IncludedIn[e](issuedcode) == true { return strings.ToLower(e) }
		}
		if reasontext != "" { return reasontext }

		if forsubject := ""; len(fo.DeliveryStatus) > 3 {
			// Check the first 3 characters of fo.DeliveryStatus
			forsubject = fo.DeliveryStatus[0:3]
			if forsubject == "5.6" || forsubject == "4.6"   { return "contenterror"  }
			if forsubject == "5.7" || forsubject == "4.7"   { return "securityerror" }
		}
		if strings.HasPrefix(fo.DiagnosticType, "X-UNIX")   { return "mailererror"   }
		if ProbesInto["SyntaxError"](fo) == true            { return "syntaxerror"   }
		if fo.Action == "delayed" || fo.Action == "expired" { return "expired"       }
		if fo.Command == "EHLO"   || fo.Command == "HELO"   { return "blocked"       }
	}
	return reasontext
}

