// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____       _           _           _ 
// |  _ \ ___ (_) ___  ___| |_ ___  __| |
// | |_) / _ \| |/ _ \/ __| __/ _ \/ _` |
// |  _ <  __/| |  __/ (__| ||  __/ (_| |
// |_| \_\___|/ |\___|\___|\__\___|\__,_|
//          |__/                         

package reason
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool):         true if the argument includes one or more error message pattern
	IncludedIn["Rejected"] = func(argv1 string) bool {
		if argv1 == "" { return false }

		isnot := []string{
			"5.1.0 address rejected",
			"recipient address rejected",
			"sender ip address rejected",
		}
		index := []string{
			"access denied (in reply to mail from command)",
			"access denied (sender blacklisted)",
			"address rejected",
			"administrative prohibition",
			"batv failed to verify",   // SoniWall
			"batv validation failure", // SoniWall
			"backscatter protection detected an invalid or expired email address", // MDaemon
			"because the sender isn't on the recipient's list of senders to accept mail from",
			"bogus mail from",          // IMail - block empty sender
			"can't determine purported responsible address",
			"connections not accepted from servers without a valid sender domain",
			"denied [bouncedeny]",      // McAfee
			"denied by secumail valid-address-filter",
			"delivery not authorized, message refused",
			"does not exist e2110",
			"domain of sender address ",
			"email address is on senderfilterconfig list",
			"emetteur invalide",
			"empty envelope senders not allowed",
			"envelope blocked - ",
			"error: no third-party dsns",   // SpamWall - block empty sender
			"from: domain is invalid. please provide a valid from:",
			"fully qualified email address required",   // McAfee
			"invalid domain, see <url:",
			"invalid sender",
			"is not a registered gateway user",
			"mail from not owned by user",
			"message rejected: email address is not verified",
			"mx records for ",
			"null sender is not allowed",
			"recipient addresses rejected : access denied",
			"recipient not accepted. (batv: no tag",
			"returned mail not accepted here",
			"rfc 1035 violation: recursive cname records for",
			"rule imposed mailbox access for",  // MailMarshal
			"sending this from a different address or alias using the ",
			"sender address has been blacklisted",
			"sender email address rejected",
			"sender is in my black list",
			"sender is spammer",
			"sender not pre-approved",
			"sender rejected",
			"sender domain is empty",
			"sender verify failed",     // Exim callout
			"sender was rejected",      // qmail
			"spam reporting address",   // SendGrid|a message to an address has previously been marked as Spam by the recipient.
			"syntax error: empty email address",
			"the message has been rejected by batv defense",
			"this server does not accept mail from",
			"transaction failed unsigned dsn for",
			"unroutable sender address",
			"you are not allowed to post to this mailing list",
			"you are sending to/from an address that has been blacklisted",
			"your access to submit messages to this e-mail system has been rejected",
			"your email address has been blacklisted",  // MessageLabs
		}
		if moji.ContainsAny(argv1, isnot) { return false }
		if moji.ContainsAny(argv1, index) { return true  }
		return false
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress
	//   Returns:
	//     - (bool):         true if a reason is the reason defined in this file
	ProbesInto["Rejected"] = func(fo *sis.Fact) bool {
		if fo        == nil        { return false }
		if fo.Reason == "rejected" { return true  }

		tempreason := status.Name(fo.DeliveryStatus)
		if tempreason == "rejected" { return true } // Delivery status code points "rejected"
		if tempreason == ""         { tempreason = "undefined" }

		// Check the value of Diagnosic-Code: field with patterns
		if issuedcode := strings.ToLower(fo.DiagnosticCode); fo.Command == "MAIL" {
			// The session was rejected at 'MAIL FROM' command
			if IncludedIn["Rejected"](issuedcode) == true { return true }

		} else if fo.Command == "DATA" && tempreason != "userunknown" {
			// The session was rejected at 'DATA' command except "userunknown"
			if IncludedIn["Rejected"](issuedcode) == true { return true }

		} else if IsExplicit(tempreason) == false || slices.Contains([]string{"securityerror", "systemerror"}, tempreason) {
			// Try to match with message patterns when the temporary reason is "onhold", "undefined",
			// "securityerror", or "systemerror"
			if IncludedIn["Rejected"](issuedcode) == true { return true }
		}
		return false
	}
}

