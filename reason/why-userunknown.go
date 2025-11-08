// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _               _   _       _                              
// | | | |___  ___ _ __| | | |_ __ | | ___ __   _____      ___ __  
// | | | / __|/ _ \ '__| | | | '_ \| |/ / '_ \ / _ \ \ /\ / / '_ \ 
// | |_| \__ \  __/ |  | |_| | | | |   <| | | | (_) \ V  V /| | | |
//  \___/|___/\___|_|   \___/|_| |_|_|\_\_| |_|\___/ \_/\_/ |_| |_|

package reason
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
	IncludedIn[eb.ReUSER] = func(mesg string) bool {
		if mesg == "" { return false }

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
			"weil die adresse nicht gefunden wurde oder keine e-mails empfangen kann",
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
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReUSER] = func(fo *siba.Fact) bool {
		if fo        == nil       { return false }
		if fo.Reason == eb.ReUSER { return true  }

		tempreason := status.Name(fo.DeliveryStatus); if tempreason == eb.ReQUIT { return false }
		issuedcode := strings.ToLower(fo.DiagnosticCode)

		if tempreason == eb.ReUSER {
			// *.1.1 = 'Bad destination mailbox address'
			//   Status: 5.1.1
			//   Diagnostic-Code: SMTP; 550 5.1.1 <***@example.jp>:
			//     Recipient address rejected: User unknown in local recipient table
			for _, e := range []string{eb.ReRELA, eb.ReBLOC, eb.ReFULL, eb.ReMOVE, eb.ReFROM, eb.Re00MX} {
				// Check the value of "Diagnostic-Code" with other error patterns.
				if IncludedIn[e](issuedcode) { return false }
			}
			return true

		} else {
			// The reason name found by fo.DeliveryStatus is not UserUnknown, or is empty
			// When the SMTP command is not "RCPT", the session rejected by other reason, maybe.
			if fo.Command == "RCPT" && IncludedIn[eb.ReUSER](issuedcode) { return true }
		}
		return false
	}
}

