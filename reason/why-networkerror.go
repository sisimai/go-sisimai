// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _      _                      _    _____                     
// | \ | | ___| |___      _____  _ __| | _| ____|_ __ _ __ ___  _ __ 
// |  \| |/ _ \ __\ \ /\ / / _ \| '__| |/ /  _| | '__| '__/ _ \| '__|
// | |\  |  __/ |_ \ V  V / (_) | |  |   <| |___| |  | | | (_) | |   
// |_| \_|\___|\__| \_/\_/ \___/|_|  |_|\_\_____|_|  |_|  \___/|_|   

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
	IncludedIn[eb.ReINET] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"address family mismatch on destination mxs", // OpenSMTPD/smtpd/mta.c
			"all routes to destination blocked",          // OpenSMTPD/smtpd/mta.c
			"bad dns lookup error code",                  // OpenSMTPD/smtpd/mta.c
			"could not connect and send the mail to",
			"could not contact dns servers",
			"could not retrieve source address",          // OpenSMTPD/smtpd/mta.c
			"dns records for the destination computer could not be found",
			"establish an smtp connection",
			"exceeded maximum hop count",                 // Courier
			"host is unreachable",
			"host name lookup failure",
			"host not found, try again",
			"listed as a best-preference mx",
			"loop detected",                              // OpenSMTPD/smtpd/mta.c
			"maximum forwarding loop count exceeded",
			"network error on destination mxs",           // OpenSMTPD/smtpd/mta.c
			"no relevant answers",
			"temporary failure in mx lookup",             // OpenSMTPD/smtpd/mta.c
			"too many hops",
			"unable to resolve route ",
			"unrouteable mail domain",
		}
		pairs := [][]string{
			[]string{"malformed", "name server reply"},
			[]string{"mail ", "loop"},
			[]string{"message ", "loop"},
			[]string{"no ", "route to"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReINET] = func(fo *siba.Fact) bool { return false }
}

