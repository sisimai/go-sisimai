// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      _______     _           
//  _ __| |__   ___  ___| |_   / /__  /___ | |__   ___  
// | '__| '_ \ / _ \/ __| __| / /  / // _ \| '_ \ / _ \ 
// | |  | | | | (_) \__ \ |_ / /  / /| (_) | | | | (_) |
// |_|  |_| |_|\___/|___/\__/_/  /____\___/|_| |_|\___/ 

package rhost
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (string): Bounce reason name or an empty string.
	ReturnedBy["Zoho"] = func(fo *siba.Fact) string {
		// - Zoho Mail: https://www.zoho.com/mail/
		// - Reasons an email is marked as Spam: https://www.zoho.com/mail/help/spam-reason.html
		// - https://github.com/zoho/zohodesk-oas/blob/main/v1.0/EmailFailureAlert.json
		// - Zoho SMTP Error Codes | SMTP Field Manual: https://smtpfieldmanual.com/provider/zoho
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			eb.ReAUTH: []string{
				// - <*******@zoho.com>: host smtpin.zoho.com[204.141.33.23] said: 550 5.7.1 Email
				//   rejected per DMARC policy for zoho.com
				"Email rejected per DMARC policy",
			},
			eb.ReBLOC: []string{
				// - mx.zoho.com[204.141.33.44]:25, delay=1202, delays=1200/0/0.91/0.30, dsn=4.7.1,
				//   status=deferred (host mx.zoho.com[204.141.33.44] said:
				//   451 4.7.1 Greylisted, try again after some time (in reply to RCPT TO command))
				"Greylisted, try again after some time",
			},
			eb.ReFROM: []string{
				// - <*******@zoho.com>: host smtpin.zoho.com[204.141.33.23] said: 554 5.7.1 Email
				//   cannot be delivered. Reason: Email flagged as Spam. (in reply to RCPT TO command)
				// - <***@zoho.com>: host mx.zoho.com[136.143.183.44] said: 541 5.4.1 Mail rejected
				//   by destination domain (in reply to RCPT TO command)
				"Email cannot be delivered. Reason: Email flagged as Spam",
				"Mail rejected by destination domain",
			},
			eb.ReWONT: []string{
				// - <*******@zoho.com>: host smtpin.zoho.com[204.141.33.23] said: 554 5.7.7 Email
				//   policy violation detected (in reply to end of DATA command)
				"Email policy violation detected",
				"Mailbox delivery restricted by policy error",
			},
			eb.ReSYSE: []string{
				// - https://github.com/zoho/zohodesk-oas/blob/main/v1.0/EmailFailureAlert.json#L168
				//   452 4.3.1 Temporary System Error
				"Temporary System Error",
			},
			eb.ReUSER: []string{
				// - <*******@zoho.com>: host smtpin.zoho.com[204.141.33.23] said:
				//   550 5.1.1 User does not exist - <***@zoho.com> (in reply to RCPT TO command)
				// - 552 5.1.1 <****@zoho.com> Mailbox delivery failure policy error
				"User does not exist",
			},
			eb.ReEXEC: []string{
				// - 552 5.7.1 virus **** detected by Zoho Mail
				" detected by Zoho Mail",
			},
		}

		for e := range messagesof {
			// Each key is an error reason name
			if moji.ContainsAny(fo.DiagnosticCode, messagesof[e]) { return e }
		}
		return ""
	}
}

