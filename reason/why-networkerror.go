// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _      _                      _    _____                     
// | \ | | ___| |___      _____  _ __| | _| ____|_ __ _ __ ___  _ __ 
// |  \| |/ _ \ __\ \ /\ / / _ \| '__| |/ /  _| | '__| '__/ _ \| '__|
// | |\  |  __/ |_ \ V  V / (_) | |  |   <| |___| |  | | | (_) | |   
// |_| \_|\___|\__| \_/\_/ \___/|_|  |_|\_\_____|_|  |_|  \___/|_|   

package reason
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn["NetworkError"] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"could not connect and send the mail to",
			"dns records for the destination computer could not be found",
			"hop count exceeded - possible mail loop",
			"host is unreachable",
			"host name lookup failure",
			"host not found, try again",
			"mail forwarding loop for ",
			"malformed name server reply",
			"malformed or unexpected name server reply",
			"maximum forwarding loop count exceeded",
			"message looping",
			"message probably in a routing loop",
			"no route to host",
			"too many hops",
			"unable to resolve route ",
			"unrouteable mail domain",
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["NetworkError"] = func(fo *siba.Fact) bool { return false }
}

