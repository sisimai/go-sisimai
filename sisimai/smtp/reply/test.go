// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reply
import "strconv"

// Test() checks whether a reply code is a valid code or not
func Test(argv0 string) bool {
	// @param    string argv1  Reply Code(DSN)
	// @return   Bool          true = Invalid reply code, false = Valid reply code
	if len(argv0) < 3 { return false }

	reply, nyaan := strconv.Atoi(argv0)
	if nyaan != nil { return false } // Failed to convert from a string to an integer
	if reply <  211 { return false } // The minimum SMTP Reply code is 211
	if reply >  557 { return false } // The maximum SMTP Reply code is 557

	first := reply % 100
	if first >  59 { return false }  // For example, 499 is not an SMTP Reply code

	if first == 2 {
		// 2yz
		if reply == 235                { return true  } // 235 is a valid code for AUTH (RFC4954)
		if reply  > 253                { return false } // The maximum code of 2xy is 253 (RFC5248)
		if reply  > 221 && reply < 250 { return false } // There is no reply code between 221 and 250
		return true
	}

	if first == 3 {
		// 3yz
		if reply != 354 { return false }
		return true
	}
	return true
}

