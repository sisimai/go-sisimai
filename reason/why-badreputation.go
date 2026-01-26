// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____            _ ____                  _        _   _             
// | __ )  __ _  __| |  _ \ ___ _ __  _   _| |_ __ _| |_(_) ___  _ __  
// |  _ \ / _` |/ _` | |_) / _ \ '_ \| | | | __/ _` | __| |/ _ \| '_ \ 
// | |_) | (_| | (_| |  _ <  __/ |_) | |_| | || (_| | |_| | (_) | | | |
// |____/ \__,_|\__,_|_| \_\___| .__/ \__,_|\__\__,_|\__|_|\___/|_| |_|
//                             |_|                                     

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReFAMA] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"has been temporarily rate limited due to ip reputation",
			"ip/domain reputation problems",
			"likely suspicious due to the very low reputation",
			"none/bad reputation", // t-online.de
			"poor email reputation score",
			"sending mta's poor reputation",
			"temporarily deferred due to unexpected volume or user complaints", // Yahoo Inc.
		}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReFAMA] = func(fo *siba.Fact) bool {
		if fo == nil              { return false }
		if fo.Reason == eb.ReFAMA { return true  }
		return IncludedIn[eb.ReFAMA](strings.ToLower(fo.DiagnosticCode))
	}
}

