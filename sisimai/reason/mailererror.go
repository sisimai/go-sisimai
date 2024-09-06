// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  __  __       _ _           _____                     
// |  \/  | __ _(_) | ___ _ __| ____|_ __ _ __ ___  _ __ 
// | |\/| |/ _` | | |/ _ \ '__|  _| | '__| '__/ _ \| '__|
// | |  | | (_| | | |  __/ |  | |___| |  | | | (_) | |   
// |_|  |_|\__,_|_|_|\___|_|  |_____|_|  |_|  \___/|_|   
import "strings"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["MailerError"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			" || exit ",
			"procmail: ",
			"bin/procmail",
			"bin/maidrop",
			"command failed: ",
			"command died with status ",
			"command output:",
			"mailer error",
			"pipe to |/",
			"x-unix; ",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "mailererror" or not
	Truth["MailerError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is mailererror, false: is not mailererror
		return false
	}
}

