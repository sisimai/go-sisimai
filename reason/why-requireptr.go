// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                  _          ____ _____ ____  
// |  _ \ ___  __ _ _   _(_)_ __ ___|  _ \_   _|  _ \ 
// | |_) / _ \/ _` | | | | | '__/ _ \ |_) || | | |_) |
// |  _ <  __/ (_| | |_| | | | |  __/  __/ | | |  _ < 
// |_| \_\___|\__, |\__,_|_|_|  \___|_|    |_| |_| \_\
//               |_|                                  

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReQPTR] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"access denied. ip name lookup failed",
			"cannot find your hostname",
			"cannot resolve your address.",
			"client host rejected: cannot find your hostname", // Yahoo!
			"sender ip reverse lookup rejected",
			"the corresponding forward dns entry does not point to the sending ip", // Google
			"unresolvable relay host name",
		}
		pairs := [][]string{
			[]string{"domain "," mismatches client ip"},
			[]string{"dns lookup failure: ", " try again later"},
			[]string{"ptr", "record"},
			[]string{"reverse", " dns"},
			[]string{"server access ", " forbidden by invalid rdns record of your mail server"},
			[]string{"service permits ", " unverifyable sending ips"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReQPTR] = func(fo *siba.Fact) bool {
		if fo == nil                                   { return false }
		if fo.Reason == eb.ReQPTR                      { return true  }
		if status.Name(fo.DeliveryStatus) == eb.ReQPTR { return true  }
		return IncludedIn[eb.ReQPTR](strings.ToLower(fo.DiagnosticCode))
	}
}

