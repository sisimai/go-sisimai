// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____       _ _          __     ___       _       _   _             
// |  _ \ ___ | (_) ___ _   \ \   / (_) ___ | | __ _| |_(_) ___  _ __  
// | |_) / _ \| | |/ __| | | \ \ / /| |/ _ \| |/ _` | __| |/ _ \| '_ \ 
// |  __/ (_) | | | (__| |_| |\ V / | | (_) | | (_| | |_| | (_) | | | |
// |_|   \___/|_|_|\___|\__, | \_/  |_|\___/|_|\__,_|\__|_|\___/|_| |_|
//                      |___/                                          

package reason
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool):         true if the argument includes one or more error message pattern
	IncludedIn["PolicyViolation"] = func(argv1 string) bool {
		if argv1 == "" { return false }

		index := []string{
			"an illegal attachment on your message",
			"because the recipient is not accepting mail with ", // AOL Phoenix
			"by non-member to a members-only list",
			"closed mailing list",
			"denied by policy",
			"email not accepted for policy reasons",
			// http://kb.mimecast.com/Mimecast_Knowledge_Base/Administration_Console/Monitoring/Mimecast_SMTP_Error_Codes#554
			"email rejected due to security policies",
			"header are not accepted",
			"header error",
			"local policy violation",
			"message bounced due to organizational settings",
			"message given low priority",
			"message not accepted for policy reasons",
			"message rejected due to local policy",
			"messages with multiple addresses",
			"rejected for policy reasons",
			"protocol violation",
			"the email address used to send your message is not subscribed to this group",
			"the message was rejected by organization policy",
			"this message was blocked because its content presents a potential",
			"we do not accept messages containing images or other attachments",
			"you're using a mass mailer",
		}
		pairs := [][]string{
			[]string{"you have exceeded the", "allowable number of posts without solving a captcha"},
		}
		return moji.ContainsAny(argv1, index) || moji.AlignedAny(argv1, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress
	//   Returns:
	//     - (bool):         true if a reason is the reason defined in this file
	ProbesInto["PolicyViolation"] = func(fo *sis.Fact) bool { return false }
}

