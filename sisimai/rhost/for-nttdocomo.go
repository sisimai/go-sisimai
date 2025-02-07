// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      ___   _ _____ _____ ____   ___   ____ ___  __  __  ___  
//  _ __| |__   ___  ___| |_   / / \ | |_   _|_   _|  _ \ / _ \ / ___/ _ \|  \/  |/ _ \ 
// | '__| '_ \ / _ \/ __| __| / /|  \| | | |   | | | | | | | | | |  | | | | |\/| | | | |
// | |  | | | | (_) \__ \ |_ / / | |\  | | |   | | | |_| | |_| | |__| |_| | |  | | |_| |
// |_|  |_| |_|\___/|___/\__/_/  |_| \_| |_|   |_| |____/ \___/ \____\___/|_|  |_|\___/ 
import "strings"
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["NTTDOCOMO"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		messagesof := map[string][]string{
			"mailboxfull": []string{"552 too much mail data"},
			"toomanyconn": []string{"552 too many recipients"},
			"syntaxerror": []string{"503 bad sequence of commands", "504 command parameter not implemented"},
		}
		statuscode := fo.DeliveryStatus
		issuedcode := strings.ToLower(fo.DiagnosticCode)
		reasontext := ""

		// Check the value of Status: field, an SMTP Reply Code, and the SMTP Command
		if statuscode == "5.1.1" || statuscode == "5.0.911" {
			//    ----- Transcript of session follows -----
			// ... while talking to mfsmax.docomo.ne.jp.:
			// >>> RCPT To:<***@docomo.ne.jp>
			// <<< 550 Unknown user ***@docomo.ne.jp
			// 550 5.1.1 <***@docomo.ne.jp>... User unknown
			// >>> DATA
			// <<< 503 Bad sequence of commands
			reasontext = "userunknown"

		} else if statuscode == "5.2.0" {
			//    ----- The following addresses had permanent fatal errors -----
			// <***@docomo.ne.jp>
			// (reason: 550 Unknown user ***@docomo.ne.jp)
			// 
			//    ----- Transcript of session follows -----
			// ... while talking to mfsmax.docomo.ne.jp.:
			// >>> DATA
			// <<< 550 Unknown user ***@docomo.ne.jp
			// 554 5.0.0 Service unavailable
			// ...
			// Final-Recipient: RFC822; ***@docomo.ne.jp
			// Action: failed
			// Status: 5.2.0
			reasontext = "filtered"

		} else {
			// The value of "Diagnostic-Code:" field is not empty
			for e := range messagesof {
				// The key name is a bounce reason name
				if sisimoji.ContainsAny(issuedcode, messagesof[e]) == false { continue }
				reasontext = e; break
			}
		}
		if reasontext != "" { return reasontext }

		// A bounce reason did not decide from the status code, the error message.
		if statuscode == "5.0.0" {
			// Status: 5.0.0
			if fo.Command == "RCPT" {
				// Your message to the following recipients cannot be delivered:
				//
				// <***@docomo.ne.jp>:
				// mfsmax.docomo.ne.jp [203.138.181.112]:
				// >>> RCPT TO:<***@docomo.ne.jp>
				// <<< 550 Unknown user ***@docomo.ne.jp
				// ...
				//
				// Final-Recipient: rfc822; ***@docomo.ne.jp
				// Action: failed
				// Status: 5.0.0
				// Remote-MTA: dns; mfsmax.docomo.ne.jp [203.138.181.112]
				// Diagnostic-Code: smtp; 550 Unknown user ***@docomo.ne.jp
				reasontext = "userunknown"

			} else if fo.Command == "DATA" {
				// <***@docomo.ne.jp>: host mfsmax.docomo.ne.jp[203.138.181.240] said:
				// 550 Unknown user ***@docomo.ne.jp (in reply to end of DATA
				// command)
				// ...
				// Final-Recipient: rfc822; ***@docomo.ne.jp
				// Original-Recipient: rfc822;***@docomo.ne.jp
				// Action: failed
				// Status: 5.0.0
				// Remote-MTA: dns; mfsmax.docomo.ne.jp
				// Diagnostic-Code: smtp; 550 Unknown user ***@docomo.ne.jp
				reasontext = "rejected"

			} else {
				// Rejected by other SMTP commands: AUTH, MAIL,
				//   もしもこのブロックを通過するNTTドコモからのエラーメッセージを見つけたら
				//   https://github.com/sisimai/p5-sisimai/issues からご連絡ねがいます。
				//
				//   If you found a error message from mfsmax.docomo.ne.jp which passes this block,
				//   please open an issue at https://github.com/sisimai/p5-sisimai/issues .
			}
		} else {
			// Status: field is neither 5.0.0 nor values defined in code above
			//   もしもこのブロックを通過するNTTドコモからのエラーメッセージを見つけたら
			//   https://github.com/sisimai/p5-sisimai/issues からご連絡ねがいます。
			//
			//   If you found a error message from mfsmax.docomo.ne.jp which passes this block,
			//   please open an issue at https://github.com/sisimai/p5-sisimai .
		}
		return reasontext
	}
}

