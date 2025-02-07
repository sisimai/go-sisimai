// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//   ___        _   _       _     _ 
//  / _ \ _ __ | | | | ___ | | __| |
// | | | | '_ \| |_| |/ _ \| |/ _` |
// | |_| | | | |  _  | (_) | | (_| |
//  \___/|_| |_|_| |_|\___/|_|\__,_|
import "sisimai/sis"
import "sisimai/smtp/status"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["OnHold"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		return false
	}

	// The bounce reason is "onhold" or not
	ProbesInto["OnHold"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is onhold, false: is not onhold
		if fo.Reason == "onhold"                      { return true }
		if status.Name(fo.DeliveryStatus) == "onhold" { return true }
		return false
	}
}

