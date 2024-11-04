// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _               _   _       _                              
// | | | |___  ___ _ __| | | |_ __ | | ___ __   _____      ___ __  
// | | | / __|/ _ \ '__| | | | '_ \| |/ / '_ \ / _ \ \ /\ / / '_ \ 
// | |_| \__ \  __/ |  | |_| | | | |   <| | | | (_) \ V  V /| | | |
//  \___/|___/\___|_|   \___/|_| |_|_|\_\_| |_|\___/ \_/\_/ |_| |_|
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"
import sisimoji "sisimai/string"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["UserUnknown"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		index := []string{
			"#5.1.1 bad address",
			"550 address invalid",
			"5.1.0 address rejected.",
			"address does not exist",
			"address not present in directory",
			"address unknown",
			"can't accept user",
			"does not exist.",
			"destination addresses were unknown",
			"destination server rejected recipients",
			"email address does not exist",
			"email address could not be found",
			"invalid address",
			"invalid mailbox",
			"invalid mailbox path",
			"invalid recipient",
			"is not a known user",
			"is not a valid mailbox",
			"is not an active address at this host",
			"mailbox does not exist",
			"mailbox invalid",
			"mailbox is inactive",
			"mailbox is unavailable",
			"mailbox not present",
			"mailbox not found",
			"mailbox unavaiable",
			"nessun utente simile in questo indirizzo",
			"no account by that name here",
			"no existe dicha persona",
			"no existe ese usuario ",
			"no mail box available for this user",
			"no mailbox by that name is currently available",
			"no mailbox found",
			"no such address here",
			"no such mailbox",
			"no such person at this address",
			"no such recipient",
			"no such user",
			"no thank you rejected: account unavailable",
			"no valid recipients, bye",
			"not a valid recipient",
			"not a valid user here",
			"not a local address",
			"not email addresses",
			"recipient address rejected. (in reply to rcpt to command)",
			"recipient address rejected: access denied",
			"recipient address rejected: invalid user",
			"recipient address rejected: invalid-recipient",
			"recipient address rejected: unknown user",
			"recipient address rejected: userunknown",
			"recipient does not exist",
			"recipient is not accepted",
			"recipient is not local",
			"recipient not exist",
			"recipient not found",
			"recipient not ok",
			"recipient refuses to accept your mail",
			"recipient unknown",
			"requested action not taken: mailbox unavailable",
			"resolver.adr.recipient notfound",
			"sorry, user unknown",
			"sorry, badrcptto",
			"sorry, no mailbox here by that name",
			"sorry, your envelope recipient has been denied",
			"that domain or user isn't in my list of allowed rcpthosts",
			"the email account that you tried to reach does not exist",
			"the following recipients was undeliverable",
			"the user's email name is not found",
			"there is no one at this address",
			"this address no longer accepts mail",
			"this email address is wrong or no longer valid",
			"this recipient is in my badrecipientto list",
			"this recipient is not in my validrcptto list",
			"this spectator does not exist",
			"unknown mailbox",
			"unknown recipient",
			"unknown user",
			"user does not exist",
			"user missing home directory",
			"user not active",
			"user not exist",
			"user not found",
			"user not known",
			"user unknown",
			"utilisateur inconnu !",
			"vdeliver: invalid or unknown virtual user",
			"your envelope recipient is in my badrcptto list",
		}
		pairs := [][]string{
			[]string{"<", "> not found"},
			[]string{"<", ">... blocked by "},
			[]string{"account ", " does not exist at the organization"},
			[]string{"adresse d au moins un destinataire invalide. invalid recipient.", "416"},
			[]string{"adresse d au moins un destinataire invalide. invalid recipient.", "418"},
			[]string{"bad", "recipient"},
			[]string{"mailbox ", "does not exist"},
			[]string{"mailbox ", "unavailable or access denied"},
			[]string{"no ", " in name directory"},
			[]string{"non", "existent user"},
			[]string{"rcpt <", " does not exist"},
			[]string{"rcpt (", "t exist "},
			[]string{"recipient ", " was not found in"},
			[]string{"recipient address rejected: user ", "  does not exist"},
			[]string{"recipient address rejected: user unknown in ", "  table"},
			[]string{"said: 550-5.1.1 ", " user unknown "},
			[]string{"said: 550 5.1.1 ", " user unknown "},
			[]string{"this user doesn't have a ", " account"},
			[]string{"unknown e", "mail address"},
			[]string{"unknown local", "part"},
			[]string{"user ", " was not found"},
			[]string{"user ", " does not exist"},
			[]string{"user (", ") unknown"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "userunknown" or not
	ProbesInto["UserUnknown"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is userunknown, false: is not userunknown
		if fo.Reason == "userunknown" { return true }

		tempreason := status.Name(fo.DeliveryStatus); if tempreason == "suspend" { return false }
		issuedcode := strings.ToLower(fo.DiagnosticCode)

		if tempreason == "userunknown" {
			// *.1.1 = 'Bad destination mailbox address'
			//   Status: 5.1.1
			//   Diagnostic-Code: SMTP; 550 5.1.1 <***@example.jp>:
			//     Recipient address rejected: User unknown in local recipient table
			prematches := []string{"NoRelaying", "Blocked", "MailboxFull", "HasMoved", "Rejected", "NotAccept"}
			matchother := false

			for _, e := range prematches {
				// Check the value of "Diagnostic-Code" with other error patterns.
				if IncludedIn[e](issuedcode) == false { continue }
				matchother = true; break
			}
			if matchother == false { return true } // Did not match with other message patterns

		} else {
			// The reason name found by fo.DeliveryStatus is not "userunknown", or is empty
			if fo.Command == "RCPT" {
				// When the SMTP command is not "RCPT", the session rejected by other reason, maybe.
				if IncludedIn["UserUnknown"](strings.ToLower(fo.DiagnosticCode)) { return true }
			}
		}
		return false
	}
}

