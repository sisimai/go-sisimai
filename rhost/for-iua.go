// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      _____ _   _   _    
//  _ __| |__   ___  ___| |_   / /_ _| | | | / \   
// | '__| '_ \ / _ \/ __| __| / / | || | | |/ _ \  
// | |  | | | | (_) \__ \ |_ / /  | || |_| / ___ \ 
// |_|  |_| |_|\___/|___/\__/_/  |___|\___/_/   \_\

package rhost
import "strings"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (string): Bounce reason name or an empty string.
	ReturnedBy["IUA"] = func(fo *siba.Fact) string {
		// - https://www.i.ua/
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		errorcodes := map[string]string{
			// https://mail.i.ua/err/$(CODE)/
			"1":  "norelaying",  // The use of SMTP as mail gate is forbidden.
			"2":  "userunknown", // User is not found.
			"3":  "suspend",     // Mailbox was not used for more than 3 months
			"4":  "mailboxfull", // Mailbox is full.
			"5":  "toomanyconn", // Letter sending limit is exceeded.
			"6":  "norelaying",  // Use SMTP of your provider to send mail.
			"7":  "blocked",     // Wrong value if command HELO/EHLO parameter.
			"8":  "rejected",    // Couldn't check sender address.
			"9":  "blocked",     // IP-address of the sender is blacklisted.
			"10": "filtered",    // Not in the list Mail address management.
		}
		issuedcode := strings.ToLower(fo.DiagnosticCode)
		codenumber := moji.Select(issuedcode, ".i.ua/err/", "/", 0); if codenumber == "" { return "" }
		return errorcodes[codenumber]
	}
}

