// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____            _              _ 
// | ____|_  ___ __ (_)_ __ ___  __| |
// |  _| \ \/ / '_ \| | '__/ _ \/ _` |
// | |___ >  <| |_) | | | |  __/ (_| |
// |_____/_/\_\ .__/|_|_|  \___|\__,_|
//            |_|                     

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
	IncludedIn[eb.ReTIME] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"connection timed out",
			"could not find a gateway for",
			"delivery attempts will continue to be",
			"delivery expired",
			"delivery time expired",
			"failed to deliver to domain ",
			"giving up on",
			"have been failing for a long time",
			"has been delayed",
			"it has not been collected after",
			"message expired, ",
			"message timed out",
			"retry time not reached for any host after a long failure period",
			"server did not respond",
			"this message has been in the queue too long",
			"unable to deliver message after multiple retries",
			"was not reachable within the allowed queue period",
			"your message could not be delivered for more than",
		}
		pairs := [][]string{
			[]string{"could not be delivered for", " days"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReTIME] = func(fo *siba.Fact) bool { return false }
}

