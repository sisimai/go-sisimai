// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      ____   __    _                ___            
//  _ __| |__   ___  ___| |_   / /\ \ / /_ _| |__   ___   ___|_ _|_ __   ___ 
// | '__| '_ \ / _ \/ __| __| / /  \ V / _` | '_ \ / _ \ / _ \| || '_ \ / __|
// | |  | | | | (_) \__ \ |_ / /    | | (_| | | | | (_) | (_) | || | | | (__ 
// |_|  |_| |_|\___/|___/\__/_/     |_|\__,_|_| |_|\___/ \___/___|_| |_|\___|
import "strings"
import "sisimai/sis"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["YahooInc"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      https://senders.yahooinc.com/smtp-error-codes
		//           https://smtpfieldmanual.com/provider/yahoo
		//           https://www.postmastery.com/yahoo-postmaster/
		if fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"authfailure": []string{
				// - 550 5.7.9 This mail has been blocked because the sender is unauthenticated. Yahoo
				//   requires all senders to authenticate with either SPF or DKIM.
				"yahoo requires all senders to authenticate with either spf or dkim",
			},
			"blocked": []string{
				// - 553 5.7.1 [BL21] Connections will not be accepted from 192.0.2.25,
				//   because the ip is in Spamhaus's list; see http://postmaster.yahoo.com/550-bl23.html
				// - 553 5.7.1 [BL23] Connections not accepted from IP addresses on Spamhaus XBL;
				//   see http://postmaster.yahoo.com/errors/550-bl23.html [550]",
				" because the ip is in spamhaus's list;",
				"not accepted from ip addresses on spamhaus xbl",
			},
			"norelaying": []string{
				// - 550 relaying denied for <***@yahoo.com>
				"relaying denied for ",
			},
			"notcomplaintrfc": []string{"headers are not rfc compliant"},
			"policyviolation": []string{
				// - 554 Message not allowed - [PH01] Email not accepted for policy reasons.
				//   Please visit https://postmaster.yahooinc.com/error-codes
				// - 554 5.7.9 Message not accepted for policy reasons. 
				//   See https://postmaster.yahooinc.com/error-codes
				"not accepted for policy reasons",
			},
			"rejected": []string{
				// Observed the following error message since around March 2024:
				//
				// - 421 4.7.0 [TSS04] Messages from 192.0.2.25 temporarily deferred due to unexpected
				//   volume or user complaints - 4.16.55.1;
				//   see https://postmaster.yahooinc.com/error-codes (in reply to MAIL FROM command))
				//
				// However, the same error message is returned even for domains that are considered to
				// have a poor reputation without SPF, DKIM, or DMARC settings, or for other reasons.
				// It seems that the error message is not as granular as Google's.
				"temporarily deferred due to unexpected volume or user complaints",

				// - 451 Message temporarily deferred due to unresolvable RFC.5321 from domain.
				//   See https://senders.yahooinc.com/error-codes//unresolvable-from-domain
				"due to unresolvable rfc.5321 domain",

				// - 553 5.7.2 [TSS09] All messages from 192.0.2.25 will be permanently deferred;
				//   Retrying will NOT succeed. See https://postmaster.yahooinc.com/error-codes
				// - 553 5.7.2 [TSS11] All messages from 192.0.2.25 will be permanently deferred;
				//   Retrying will NOT succeed. See https://postmaster.yahooinc.com/error-codes
				" will be permanently deferred",
			},
			"speeding": []string{
				// - 450 User is receiving mail too quickly
				"user is receiving mail too quickly",
			},
			"suspend": []string{
				// - 554 delivery error: dd ****@yahoo.com is no longer valid.
				// - 554 30 Sorry, your message to *****@aol.jp cannot be delivered.
				//   This mailbox is disabled (554.30)
				" is no longer valid.",
				"this mailbox is disabled",
			},
			"syntaxerror": []string{
				// - 501 Syntax error in parameters or arguments
				"syntax error in parameters or arguments",
			},
			"toomanyconn": []string{
				// - 421 Max message per connection reached, closing transmission channel
				"max message per connection reached",
			},
			"userunknown": []string{
				// - 554 delivery error: dd This user doesn't have a yahoo.com account (***@yahoo.com)
				// - 552 1 Requested mail action aborted, mailbox not found (in reply to end of DATA command)
				"dd this user doesn't have a ",
				"mailbox not found",
			},
		}

		issuedcode := strings.ToLower(fo.DiagnosticCode)
		reasontext := ""

		for e := range messagesof {
			// The key name is a bounce reason name
			for _, f := range messagesof[e] {
				// Try to find the text listed in messagesof from the error message
				if strings.Contains(issuedcode, f) == false { continue }
				reasontext = e; break
			}
		}
		return reasontext
	}
}

