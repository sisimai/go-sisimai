// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _      _                      _    _____                     
// | \ | | ___| |___      _____  _ __| | _| ____|_ __ _ __ ___  _ __ 
// |  \| |/ _ \ __\ \ /\ / / _ \| '__| |/ /  _| | '__| '__/ _ \| '__|
// | |\  |  __/ |_ \ V  V / (_) | |  |   <| |___| |  | | | (_) | |   
// |_| \_|\___|\__| \_/\_/ \___/|_|  |_|\_\_____|_|  |_|  \___/|_|   
import "strings"
import "libsisimai.org/sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["NetworkError"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

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

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "networkerror" or not
	ProbesInto["NetworkError"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is networkerror, false: is not networkerror
		return false
	}
}

