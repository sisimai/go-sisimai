// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____       _ _          __     ___       _       _   _             
// |  _ \ ___ | (_) ___ _   \ \   / (_) ___ | | __ _| |_(_) ___  _ __  
// | |_) / _ \| | |/ __| | | \ \ / /| |/ _ \| |/ _` | __| |/ _ \| '_ \ 
// |  __/ (_) | | | (__| |_| |\ V / | | (_) | | (_| | |_| | (_) | | | |
// |_|   \___/|_|_|\___|\__, | \_/  |_|\___/|_|\__,_|\__|_|\___/|_| |_|
//                      |___/                                          
import "strings"
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Try to match that the given text and message patterns
	Match["PolicyViolation"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
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

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "policyviolation" or not
	Truth["PolicyViolation"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is policyviolation, false: is not policyviolation
		return false
	}
}

