// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      _______                         _   
//  _ __| |__   ___  ___| |_   / /_   _|__ _ __   ___ ___ _ __ | |_ 
// | '__| '_ \ / _ \/ __| __| / /  | |/ _ \ '_ \ / __/ _ \ '_ \| __|
// | |  | | | | (_) \__ \ |_ / /   | |  __/ | | | (_|  __/ | | | |_ 
// |_|  |_| |_|\___/|___/\__/_/    |_|\___|_| |_|\___\___|_| |_|\__|
import "strings"
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["Tencent"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      https://service.mail.qq.com/detail/122
		if fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			"authfailure": []string{
				"spf check failed",         // https://service.mail.qq.com/detail/122/72
				"dmarc check failed",
			},
			"blocked": []string{
				"suspected bounce attacks", // https://service.mail.qq.com/detail/122/57
				"suspected spam ip",        // https://service.mail.qq.com/detail/122/66
				"connection denied",        // https://service.mail.qq.com/detail/122/170
			},
			"mesgtoobig": []string{
				"message too large",        // https://service.mail.qq.com/detail/122/168
			},
			"rejected": []string{
				"suspected spam",                   // https://service.mail.qq.com/detail/122/71
				"mail is rejected by recipients",   // https://service.mail.qq.com/detail/122/92
			},
			"spandetected": []string{
				"spam is embedded in the email",    // https://service.mail.qq.com/detail/122/59
				"mail content denied",              // https://service.mail.qq.com/detail/122/171
			},
			"speeding": []string{
				"mailbox unavailable or access denined", // https://service.mail.qq.com/detail/122/166
			},
			"suspend": []string{
				"is a deactivated mailbox", // http://service.mail.qq.com/cgi-bin/help?subtype=1&&id=20022&&no=1000742
			},
			"syntaxerror": []string{
				"bad address syntax", // https://service.mail.qq.com/detail/122/167
			},
			"toomanyconn": []string{
				"ip frequency limited",         // https://service.mail.qq.com/detail/122/172
				"domain frequency limited",     // https://service.mail.qq.com/detail/122/173
				"sender frequency limited",     // https://service.mail.qq.com/detail/122/174
				"connection frequency limited", // https://service.mail.qq.com/detail/122/175
			},
			"userunknown": []string{
				"mailbox not found",  // https://service.mail.qq.com/detail/122/169
			},
		}
		issuedcode := strings.ToLower(fo.DiagnosticCode)
		reasontext := ""

		for e := range messagesof {
			// The key name is a bounce reason name
			if sisimoji.ContainsAny(issuedcode, messagesof[e]) == false { continue }
			reasontext = e; break
		}
		return reasontext
	}
}

