// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____                       _ _         _____                     
// / ___|  ___  ___ _   _ _ __(_) |_ _   _| ____|_ __ _ __ ___  _ __ 
// \___ \ / _ \/ __| | | | '__| | __| | | |  _| | '__| '__/ _ \| '__|
//  ___) |  __/ (__| |_| | |  | | |_| |_| | |___| |  | | | (_) | |   
// |____/ \___|\___|\__,_|_|  |_|\__|\__, |_____|_|  |_|  \___/|_|   
//                                   |___/                           
import "strings"
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Try to match that the given text and message patterns
	Match["SecurityError"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			"account not subscribed to ses",
			"authentication credentials invalid",
			"authentication failure",
			"authentication required",
			"authentication turned on in your email client",
			"executable files are not allowed in compressed files",
			"insecure mail relay",
			"recipient address rejected: access denied",
			"sorry, you don't authenticate or the domain isn't in my list of allowed rcpthosts",
			"starttls is required to send mail",
			"tls required but not supported", // SendGrid:the recipient mailserver does not support TLS or have a valid certificate
			"unauthenticated senders not allowed",
			"verification failure",
			"you are not authorized to send mail, authentication is required",
		}
		pairs := [][]string{
			[]string{"authentication failed; server ", " said: "}, // Postfix
			[]string{"authentification invalide", "305"},
			[]string{"authentification requise", "402"},
			[]string{"domain ", " is a dead domain"},
			[]string{"user ", " is not authorized to perform ses:sendrawemail on resource"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "securityerror" or not
	Truth["SecurityError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is securityerror, false: is not securityerror
		return false
	}
}

