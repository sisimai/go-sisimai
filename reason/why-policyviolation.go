// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____       _ _          __     ___       _       _   _             
// |  _ \ ___ | (_) ___ _   \ \   / (_) ___ | | __ _| |_(_) ___  _ __  
// | |_) / _ \| | |/ __| | | \ \ / /| |/ _ \| |/ _` | __| |/ _ \| '_ \ 
// |  __/ (_) | | | (__| |_| |\ V / | | (_) | | (_| | |_| | (_) | | | |
// |_|   \___/|_|_|\___|\__, | \_/  |_|\___/|_|\__,_|\__|_|\___/|_| |_|
//                      |___/                                          

package reason
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
			"denied by policy",
			// http://kb.mimecast.com/Mimecast_Knowledge_Base/Administration_Console/Monitoring/Mimecast_SMTP_Error_Codes#554
			"email rejected due to security policies",
			"executable files are not allowed in compressed files",
			"for policy reasons",
			"header are not accepted",
			"header error",
			"illegal attachment on your message",
			"local policy",
			"message bounced due to organizational settings",
			"message given low priority",
			"message was rejected by organization policy",
			"message was blocked because its content presents a potential", // https://support.google.com/mail/answer/6590
			"messages with multiple addresses",
			"protocol violation",
			"we do not accept messages containing images or other attachments",
			"you're using a mass mailer",
		}
		pairs := [][]string{
			[]string{"you have exceeded the", "allowable number of posts without solving a captcha"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReWONT] = func(fo *siba.Fact) bool { return false }
}

