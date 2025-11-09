// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      _______                         _   
//  _ __| |__   ___  ___| |_   / /_   _|__ _ __   ___ ___ _ __ | |_ 
// | '__| '_ \ / _ \/ __| __| / /  | |/ _ \ '_ \ / __/ _ \ '_ \| __|
// | |  | | | | (_) \__ \ |_ / /   | |  __/ | | | (_|  __/ | | | |_ 
// |_|  |_| |_|\___/|___/\__/_/    |_|\___|_| |_|\___\___|_| |_|\__|

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
	ReturnedBy["Tencent"] = func(fo *siba.Fact) string {
		// - https://service.mail.qq.com/detail/122
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		messagesof := map[string][]string{
			eb.ReAUTH: []string{ // AuthFailure
				"spf check failed",         // https://service.mail.qq.com/detail/122/72
				"dmarc check failed",
			},
			eb.ReBLOC: []string{ // Blocked
				"suspected bounce attacks", // https://service.mail.qq.com/detail/122/57
				"suspected spam ip",        // https://service.mail.qq.com/detail/122/66
				"connection denied",        // https://service.mail.qq.com/detail/122/170
			},
			eb.ReSIZE: []string{ // MesgTooBig
				"message too large",        // https://service.mail.qq.com/detail/122/168
			},
			eb.ReFROM: []string{ // Rejected
				"suspected spam",                   // https://service.mail.qq.com/detail/122/71
				"mail is rejected by recipients",   // https://service.mail.qq.com/detail/122/92
			},
			eb.ReSPAM: []string{ // SpamDetected
				"spam is embedded in the email",    // https://service.mail.qq.com/detail/122/59
				"mail content denied",              // https://service.mail.qq.com/detail/122/171
			},
			eb.ReFAST: []string{ // Speeding
				"mailbox unavailable or access denined", // https://service.mail.qq.com/detail/122/166
			},
			eb.ReQUIT: []string{ // Suspend
				"is a deactivated mailbox", // http://service.mail.qq.com/cgi-bin/help?subtype=1&&id=20022&&no=1000742
			},
			eb.ReCOMM: []string{ // SyntaxError
				"bad address syntax", // https://service.mail.qq.com/detail/122/167
			},
			eb.ReCONN: []string{ // TooManyConn
				"ip frequency limited",         // https://service.mail.qq.com/detail/122/172
				"domain frequency limited",     // https://service.mail.qq.com/detail/122/173
				"sender frequency limited",     // https://service.mail.qq.com/detail/122/174
				"connection frequency limited", // https://service.mail.qq.com/detail/122/175
			},
			eb.ReUSER: []string{ // UserUnknown
				"mailbox not found",  // https://service.mail.qq.com/detail/122/169
			},
		}
		issuedcode := strings.ToLower(fo.DiagnosticCode); for e := range messagesof {
			// The key name is a bounce reason name
			if moji.ContainsAny(issuedcode, messagesof[e]) { return e }
		}
		return ""
	}
}

