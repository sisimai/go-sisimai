// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ___                _      
//  _ __| |__   ___  ___| |_   / / \   _ __  _ __ | | ___ 
// | '__| '_ \ / _ \/ __| __| / / _ \ | '_ \| '_ \| |/ _ \
// | |  | | | | (_) \__ \ |_ / / ___ \| |_) | |_) | |  __/
// |_|  |_| |_|\___/|___/\__/_/_/   \_\ .__/| .__/|_|\___|
//                                    |_|   |_|           

package rhost
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (string): Bounce reason name or an empty string.
	ReturnedBy["Apple"] = func(fo *siba.Fact) string {
		// - Postmaster information for iCloud Mail: https://support.apple.com/en-us/102322
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			eb.ReAUTH: []string{ // AuthFailure
				// - 554 5.7.1 Your message was rejected due to example.jp's DMARC policy.
				//   See https://support.apple.com/en-us/HT204137 for
				// - 554 5.7.1 [HME1] This message was blocked for failing both SPF and DKIM authentication
				//   checks. See https://support.apple.com/en-us/HT204137 for mailing best practices
				"s dmarc policy",
				"blocked for failing both spf and dkim autentication checks",
			},
			eb.ReBLOC: []string{ // Blocked
				// - 550 5.7.0 Blocked - see https://support.proofpoint.com/dnsbl-lookup.cgi?ip=192.0.1.2
				// - 550 5.7.1 Your email was rejected due to having a domain present in the Spamhaus
				//   DBL -- see https://www.spamhaus.org/dbl/
				// - 550 5.7.1 Mail from IP 192.0.2.1 was rejected due to listing in Spamhaus SBL.
				//   For details please see http://www.spamhaus.org/query/bl?ip=x.x.x.x
				// - 554 ****-smtpin001.me.com ESMTP not accepting connections
				"rejected due to having a domain present in the spamhaus",
				"rejected due to listing in spamhaus",
				"blocked - see https://support.proofpoint.com/dnsbl-lookup",
				"not accepting connections",
			},
			eb.ReMOVE: []string{ // HasMoved
				// - 550 5.1.6 recipient no longer on server: *****@icloud.com
				"recipient no longer on server",
			},
			eb.ReFULL: []string{ // MailboxFull
				// - 552 5.2.2 <****@icloud.com>: user is over quota (in reply to RCPT TO command)
				"user is over quota",
			},
			eb.ReRELA: []string{ // NoRelaying
				// - 554 5.7.1 <*****@icloud.com>: Relay access denied
				"relay access denied",
			},
			eb.Re00MX: []string{"host/domain does not accept mail"}, // NotAccept
			eb.ReWONT: []string{ // PolicyViolation
				// - 550 5.7.1 [CS01] Message rejected due to local policy.
				//   Please visit https://support.apple.com/en-us/HT204137
				"due to local policy",
			},
			eb.ReRATE: []string{ // RateLimited
				// - 421 4.7.1 Messages to ****@icloud.com deferred due to excessive volume.
				//   Try again later - https://support.apple.com/en-us/HT204137
				"due to excessive volume",
			},
			eb.ReFROM: []string{ // Rejected
				// - 450 4.1.8 <kijitora@example.jp>: Sender address rejected: Domain not found
				"sender address rejected",
			},
			eb.ReQUIT: []string{ // Suspend
				// - https://support.apple.com/guide/icloud/stop-using-or-reactivate-addresses-mm3adb030cbf/icloud
				// - 550 5.1.1 <****@icloud.com>: inactive email address (in reply to RCPT TO command)
				"inactive email address",
			},
			eb.ReUSER: []string{ // UserUnknown
				// - 550 5.1.1 <****@icloud.com>: inactive email address (in reply to RCPT TO command)
				// - 550 5.1.1 unknown or illegal alias: ****@icloud.com
				"user does not exist",
				"unknown or illegal alias",
			},
		}

		issuedcode := strings.ToLower(fo.DiagnosticCode); for e := range messagesof {
			// Each key is an error reason name
			if moji.ContainsAny(issuedcode, messagesof[e]) { return e }
		}
		return ""
	}
}

