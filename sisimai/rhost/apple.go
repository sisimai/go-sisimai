// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      ___                _      
//  _ __| |__   ___  ___| |_   / / \   _ __  _ __ | | ___ 
// | '__| '_ \ / _ \/ __| __| / / _ \ | '_ \| '_ \| |/ _ \
// | |  | | | | (_) \__ \ |_ / / ___ \| |_) | |_) | |  __/
// |_|  |_| |_|\___/|___/\__/_/_/   \_\ .__/| .__/|_|\___|
//                                    |_|   |_|           
import "strings"
import "sisimai/sis"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["Apple"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		if fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"authfailure": []string{
				// - 554 5.7.1 Your message was rejected due to example.jp's DMARC policy.
				//   See https://support.apple.com/en-us/HT204137 for
				// - 554 5.7.1 [HME1] This message was blocked for failing both SPF and DKIM authentication
				//   checks. See https://support.apple.com/en-us/HT204137 for mailing best practices
				"s dmarc policy",
				"blocked for failing both spf and dkim autentication checks",
			},
			"blocked": []string{
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
			"hasmoved": []string{
				// - 550 5.1.6 recipient no longer on server: *****@icloud.com
				"recipient no longer on server",
			},
			"mailboxfull": []string{
				// - 552 5.2.2 <****@icloud.com>: user is over quota (in reply to RCPT TO command)
				"user is over quota",
			},
			"norelaying": []string{
				// - 554 5.7.1 <*****@icloud.com>: Relay access denied
				"relay access denied",
			},
			"notaccept": []string{"host/domain does not accept mail"},
			"policyviolation": []string{
				// - 550 5.7.1 [CS01] Message rejected due to local policy.
				//   Please visit https://support.apple.com/en-us/HT204137
				"due to local policy",
			},
			"rejected": []string{
				// - 450 4.1.8 <kijitora@example.jp>: Sender address rejected: Domain not found
				"sender address rejected",
			},
			"speeding": []string{
				// - 421 4.7.1 Messages to ****@icloud.com deferred due to excessive volume.
				//   Try again later - https://support.apple.com/en-us/HT204137
				"due to excessive volume",
			},
			"userunknown": []string{
				// - 550 5.1.1 <****@icloud.com>: inactive email address (in reply to RCPT TO command)
				// - 550 5.1.1 unknown or illegal alias: ****@icloud.com
				"inactive email address",
				"user does not exist",
				"unknown or illegal alias",
			},
		}

		issuedcode := strings.ToLower(fo.DiagnosticCode)
		reasontext := ""
		for e := range messagesof {
			// Each key is an error reason name
			for _, f := range messagesof[e] {
				// Try to match each SMTP reply code, status code, error message
				if strings.Contains(issuedcode, f) == false { continue }
				reasontext = e; break
			}
		}
		return reasontext
	}
}

