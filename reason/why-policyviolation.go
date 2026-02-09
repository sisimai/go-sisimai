// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____       _ _          __     ___       _       _   _             
// |  _ \ ___ | (_) ___ _   \ \   / (_) ___ | | __ _| |_(_) ___  _ __  
// | |_) / _ \| | |/ __| | | \ \ / /| |/ _ \| |/ _` | __| |/ _ \| '_ \ 
// |  __/ (_) | | | (__| |_| |\ V / | | (_) | | (_| | |_| | (_) | | | |
// |_|   \___/|_|_|\___|\__, | \_/  |_|\___/|_|\__,_|\__|_|\___/|_| |_|
//                      |___/                                          

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReWONT] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"because the recipient is not accepting mail with ", // AOL Phoenix
			"closed mailing list",
			"delivery not authorized, message refused",
			"denied by policy",
			// http://kb.mimecast.com/Mimecast_Knowledge_Base/Administration_Console/Monitoring/Mimecast_SMTP_Error_Codes#554
			"email rejected due to security policies",
			"for policy reasons",
			"local policy violation",
			"message bounced due to organizational settings",
			"message given low priority",
			"message was rejected by organization policy",
			"protocol violation",
			"support.google.com/a/answer/172179",
			"you're using a mass mailer",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReWONT] = func(fo *siba.Fact) bool {
		if fo        == nil                         { return false }
		if fo.Reason == eb.ReWONT                   { return true  }
		if fo.Command != "" && fo.Command != "DATA" { return false }
		return IncludedIn[eb.ReWONT](strings.ToLower(fo.DiagnosticCode))
	}
}

