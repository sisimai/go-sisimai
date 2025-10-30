// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____                       _ _     _           _ _   
// | ____|_  _____ ___  ___  __| | |   (_)_ __ ___ (_) |_ 
// |  _| \ \/ / __/ _ \/ _ \/ _` | |   | | '_ ` _ \| | __|
// | |___ >  < (_|  __/  __/ (_| | |___| | | | | | | | |_ 
// |_____/_/\_\___\___|\___|\__,_|_____|_|_| |_| |_|_|\__|

package reason
import "strings"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn["ExceedLimit"] = func(mesg string) bool {
		if mesg == "" { return false }
		index := []string{"message header size exceeds limit", "message too large"}
		return moji.ContainsAny(mesg, index)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["ExceedLimit"] = func(fo *siba.Fact) bool {
		// Status: 5.2.3
		// Diagnostic-Code: SMTP; 552 5.2.3 Message size exceeds fixed maximum message size
		if fo == nil                                       { return false }
		if fo.Reason == "exceedlimit"                      { return true  }
		if status.Name(fo.DeliveryStatus) == "exceedlimit" { return true  }
		return IncludedIn["ExceedLimit"](strings.ToLower(fo.DiagnosticCode))
	}
}

