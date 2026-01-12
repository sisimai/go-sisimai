// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ __ ___  __ _ ___  ___  _ __  
// | '__/ _ \/ _` / __|/ _ \| '_ \ 
// | | |  __/ (_| \__ \ (_) | | | |
// |_|  \___|\__,_|___/\___/|_| |_|

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/smtp/status"

// Find detects and returns the bounce reason.
//   Arguments:
//     - fo (*siba.Fact): Decoded data in progress.
//   Returns:
//     - (string): Bounce reason name or an empty string.
func Find(fo *siba.Fact) string {
	// Return the reason text already decided except the reason matched with the name checked by
	// reason.ShouldBeRetried() function.
	if fo == nil { return "" }
	if fo.Reason != "" && ShouldBeRetried(fo.Reason) == false { return fo.Reason }
	if strings.HasPrefix(fo.DeliveryStatus, "2.")    == true  { return eb.ReSENT }

	reasontext := ""; if fo.DiagnosticType == "SMTP" || fo.DiagnosticType == "" {
		// Diagnostic-Code: SMTP; ... or empty value
		for _, e := range classorder[0] {
			// Check the values of Diagnostic-Code: and Status: fields using reason.ProbesInto[*]()
			// function of each child class in reason/why-*.go
			if ProbesInto[e](fo) { reasontext = e; break }
		}
	}

	if reasontext == "" || reasontext == eb.Re___0 {
		// The bounce reason is not detected yet at the code block above
		reasontext = anotherone(fo) // Try to find a reason name using anotherone()
		if reasontext == eb.Re___0                    { reasontext = ""        }
		if reasontext == "" && fo.Action == eb.AeSTAY { reasontext = eb.ReTIME }
		if reasontext != ""                           { return reasontext      }

		issuedcode := strings.ToLower(fo.DiagnosticCode)
		if IncludedIn[eb.ReAWAY](issuedcode)          { reasontext = eb.ReAWAY }
		if reasontext == "" && issuedcode != ""       { reasontext = eb.Re___1 }
		if reasontext == ""                           { reasontext = eb.Re___0 }
	}
	return reasontext
}

// anotherone detects the other bounce reason, is a fall back function for Find().
//   Arguments:
//     - fo (*siba.Fact): Decoded data in progress.
//   Returns:
//     - (string): Bounce reason name or an empty string.
func anotherone(fo *siba.Fact) string {
	if fo                    == nil  { return ""        }
	if IsExplicit(fo.Reason) == true { return fo.Reason }

	issuedcode := strings.ToLower(fo.DiagnosticCode)
	reasontext := status.Name(fo.DeliveryStatus)

	if ShouldBeRetried(reasontext) || fo.DiagnosticType != "SMTP" {
		// - The value of the reason is not decided yet by the fo.DeliveryStatus.
		// - Try to find the bounce reason by the fo.DeliveryStatus or fo.DianosticCode when the
		//   ShouldBeRetried(reasontext) returns true or fo.DiagnosticType is "SMTP".
		for _, e := range classorder[1] {
			// Trying to match with other patterns in reason/why-*.go
			if IncludedIn[e](issuedcode) == true { return e }
		}
		if reasontext != "" { return reasontext }

		if forsubject := ""; len(fo.DeliveryStatus) > 3 {
			// Check the first 3 characters of fo.DeliveryStatus
			forsubject = fo.DeliveryStatus[0:3]
			if forsubject == "5.6" || forsubject == "4.6"     { return eb.ReBODY }
			if forsubject == "5.7" || forsubject == "4.7"     { return eb.ReSECU }
		}
		if strings.HasPrefix(fo.DiagnosticType, "X-UNIX")     { return eb.ReUNIX }
		if ProbesInto[eb.ReCOMM](fo) == true                  { return eb.ReCOMM }
		if fo.Action == eb.AeSTAY                             { return eb.ReTIME }
		if fo.Command == eb.CeEHLO || fo.Command == eb.CeHELO { return eb.ReBLOC }
	}
	return reasontext
}

