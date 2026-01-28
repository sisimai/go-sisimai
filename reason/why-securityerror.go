// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                       _ _         _____                     
// / ___|  ___  ___ _   _ _ __(_) |_ _   _| ____|_ __ _ __ ___  _ __ 
// \___ \ / _ \/ __| | | | '__| | __| | | |  _| | '__| '__/ _ \| '__|
//  ___) |  __/ (__| |_| | |  | | |_| |_| | |___| |  | | | (_) | |   
// |____/ \___|\___|\__,_|_|  |_|\__|\__, |_____|_|  |_|  \___/|_|   
//                                   |___/                           

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
	IncludedIn[eb.ReSAFE] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"account not subscribed to ses",
			"authentication credentials invalid",
			"authentication failure",
			"authentication required",
			"authentication turned on in your email client",
			"unauthenticated senders not allowed",
			"verification failure",
			"you are not authorized to send mail, authentication is required",
			"you don't authenticate or the domain isn't in my list of allowed rcpthosts",
		}
		pairs := [][]string{
			[]string{"authentication failed; server ", " said: "}, // Postfix
			[]string{"authentification invalide", "305"},
			[]string{"authentification requise", "402"},
			[]string{"user ", " is not authorized to perform ses:sendrawemail on resource"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReSAFE] = func(fo *siba.Fact) bool { return false }
}

