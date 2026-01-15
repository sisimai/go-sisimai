// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ __ ___  __ _ ___  ___  _ __  
// | '__/ _ \/ _` / __|/ _ \| '_ \ 
// | | |  __/ (_| \__ \ (_) | | | |
// |_|  \___|\__,_|___/\___/|_| |_|

// Package "reason" provides functions for detecting the bounce reason by matching many error message
// patterns defined in why-*.go files.
package reason

import "slices"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"

// Keep each function (pointer) defined in reason/why-*.go to check/detect the bounce reason.
// % grep '^func init' ./reason/why-*.go | wc -l
var IncludedIn = make(map[string]func(string) bool, 35)
var ProbesInto = make(map[string]func(*siba.Fact) bool, 35)
var Availables = map[string]string{
	eb.ReAUTH: "Email rejected due to SPF, DKIM, DMARC failure",
	eb.ReREPU: "Email rejected due to an IP address reputation",
	eb.ReBLOC: "Email rejected due to client IP address or a hostname",
	eb.ReBODY: "Email rejected due to a header format of the email",
	eb.ReSENT: "Email delivered successfully",
	eb.ReSIZE: "Email rejected due to an email size is too big for a destination mail server",
	eb.ReTIME: "Delivery time has expired due to a connection failure",
	eb.ReTTLS: "Email delivery failed due to STARTTLS related problem",
	eb.ReFEED: "Email forwarded to the sender as a complaint message from your mailbox provider",
	eb.ReFILT: "Email rejected due to a header content after SMTP DATA command",
	eb.ReMOVE: "Email rejected due to users mailbox has moved and is not forwarded automatically",
	eb.ReHOST: "Delivery failed due to a domain part of a recipients email address does not exist",
	eb.ReFULL: "Email rejected due to a recipients mailbox is full",
	eb.ReUNIX: "Email returned due to a mailer program has not exited successfully",
	eb.ReNETW: "SMTP connection failed due to DNS look up failure or other network problems",
	eb.ReRELA: "Email rejected due to a connected host did not accept relaying",
	eb.Re00MX: "Delivery failed due to a destination mail server does not accept any email",
	eb.ReNRFC: "Email rejected due to non-compliance with RFC",
	eb.Re___1: "Sisimai could not decided the reason due to there is no (or less) detailed information for judging the reason",
	eb.RePOLI: "Email rejected due to policy violation on a destination host",
	eb.ReFROM: "Email rejected due to a senders email address (envelope from)",
	eb.ReQPTR: "Email rejected due to missing PTR record or having invalid PTR record",
	eb.ReRATE: "Rejected due to exceeding a rate limit: sending too fast or too many concurrency connections",
	eb.ReSECU: "Email rejected due to security violation was detected on a destination host",
	eb.ReSPAM: "Email rejected by spam filter running on the remote host",
	eb.ReSUPP: "Email was not delivered due to being listed in suppression list on MTA",
	eb.ReQUIT: "Email rejected due to a recipient account is being suspended",
	eb.ReCOMM: "Email rejected due to syntax error at sent commands in SMTP session",
	eb.ReSYSE: "Email returned due to system error on the remote host",
	eb.ReDISK: "Email rejected due to a destination mail servers disk is full",
	eb.Re___0: "Sisimai could not detect an error reason",
	eb.ReUSER: "Email rejected due to a local part of a recipients email address does not exist",
	eb.ReAWAY: "Email replied automatically due to a recipient is out of office",
	eb.ReEXEC: "Email rejected due to a virus scanner on a destination host",
}
var classorder = [][]string{
	[]string{
		eb.ReFULL, eb.ReSIZE, eb.ReQUIT, eb.ReMOVE, eb.ReRELA, eb.ReAUTH, eb.ReUSER, eb.ReFILT, eb.ReQPTR,
		eb.ReNRFC, eb.ReREPU, eb.ReBODY, eb.ReFROM, eb.ReHOST, eb.ReSPAM, eb.ReRATE, eb.ReBLOC,
	},
	[]string{
		eb.ReFULL, eb.ReAUTH, eb.ReREPU, eb.ReSPAM, eb.ReEXEC, eb.RePOLI, eb.ReRELA, eb.ReSYSE, eb.ReNETW,
		eb.ReQUIT, eb.ReBODY, eb.ReDISK, eb.Re00MX, eb.ReTIME, eb.ReTTLS, eb.ReSECU, eb.ReSUPP, eb.ReUNIX,
	},
}

// IsExplicit returns false when the argument is empty or is Undefined or is OnHold.
//   Arguments:
//     - name (string): Reason name.
//   Returns:
//     - (bool): true if the reason is an explicit, false otherwise.
func IsExplicit(name string) bool {
	return !(name == "" || name == eb.Re___0 || name == eb.Re___1)
}

// ShouldBeRetried returns true if the argument is a reason listed in the table defined in this function.
//   Arguments:
//     - name (string): Reason name.
//   Returns:
//     - (bool): true if the reason is listed in the table.
func ShouldBeRetried(name string) bool {
	cv := []string{eb.Re___0, eb.Re___1, eb.ReSYSE, eb.ReSECU, eb.ReTIME, eb.ReNETW, eb.ReHOST, eb.ReUSER}
	return slices.Contains(cv, name) || name == ""
}

