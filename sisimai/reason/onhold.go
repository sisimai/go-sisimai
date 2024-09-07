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
	// Try to match that the given text and message patterns
	Match["OnHold"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		return false
	}

	// The bounce reason is "onhold" or not
	Truth["OnHold"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is onhold, false: is not onhold
		if fo.Reason == "onhold"                      { return true }
		if status.Name(fo.DeliveryStatus) == "onhold" { return true }
		return false
	}
}

