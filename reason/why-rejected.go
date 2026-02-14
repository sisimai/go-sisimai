// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
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
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReFROM] = func(mesg string) bool {
		if mesg == "" { return false }

		isnot := []string{
			"5.1.0 address rejected",
			"ip address ",
			"recipient address rejected",
		}
		index := []string{
			"access denied (in reply to mail from command)",
			"administrative prohibition",
			"all recipient addresses rejected : access denied",
			"backscatter protection detected an invalid or expired email address", // MDaemon
			"badsendermx", // BadSenderMX
			"by non-member to a members-only list",
			"can't determine purported responsible address",
			"connections not accepted from servers without a valid sender domain",
			"denied by secumail valid-address-filter", // SecuMail
			"domain of sender address ",
			"email address is on senderfilterconfig list",
			"emetteur invalide",
			"empty email address",
			"empty envelope senders not allowed",
			"from: domain is invalid. please provide a valid from:",
			"fully qualified email address required",   // McAfee
			"has an outgoing mail suspension",
			"invalid sender",
			"is not a registered gateway user",
			"mail from not owned by user",
			"mailfrom domain is listed in spamhaus",
			"not member article from ", // FML
			"null sender is not allowed",
			"returned mail not accepted here",
			"sending this from a different address or alias using the ",
			"sender is spammer",
			"sender not pre-approved",
			"sender domain is empty",
			"sender domain listed at ",
			"sender verify failed",     // Exim callout
			"sendernoa",                // SenderNoA
			"server does not accept mail from",
			"spam reporting address",   // SendGrid|a message to an address has previously been marked as Spam by the recipient.
			"too many spam complaints",
			"unroutable sender address",
			"you are not allowed to post to this mailing list",
			"your access to submit messages to this e-mail system has been rejected",
			"your email address has been blacklisted",  // MessageLabs
		}
		pairs := [][]string{
			[]string{"after end of data:", ".", " does not exist"},
			[]string{"after mail from:", ".", " does not exist"},
			[]string{"domain ", " is a dead domain"},
			[]string{"email address ", "is not "},
			[]string{"reject mail from ", "@"}, // FML
			[]string{"send", "blacklisted"},
			[]string{"sender", " rejected"},
			[]string{"sender is", " list"},
		}
		if moji.ContainsAny(mesg, isnot) { return false }
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReFROM] = func(fo *siba.Fact) bool {
		if fo         == nil       { return false }
		if fo.Reason  == eb.ReFROM { return true  }

		tempreason := status.Name(fo.DeliveryStatus)
		if tempreason == eb.ReFROM { return true } // Delivery status code points Rejected.
		if tempreason == ""        { tempreason = eb.Re___0 }

		// Check the value of Diagnosic-Code: field with patterns
		if issuedcode := strings.ToLower(fo.DiagnosticCode); fo.Command == eb.CeMAIL {
			// The session was Rejected at "MAIL FROM" command
			if IncludedIn[eb.ReFROM](issuedcode) == true { return true }

		} else if fo.Command == eb.CeDATA && tempreason != eb.ReUSER {
			// The session was rejected at "DATA" command except UserUnknown.
			if IncludedIn[eb.ReFROM](issuedcode) == true { return true }

		} else if IsExplicit(tempreason) == false || slices.Contains([]string{eb.ReSAFE, eb.RePROC}, tempreason) {
			// Try to match with message patterns when the temporary reason is OnHold, Undefined,
			// SecurityError, or SystemError.
			if IncludedIn[eb.ReFROM](issuedcode) == true { return true }
		}
		return false
	}
}

