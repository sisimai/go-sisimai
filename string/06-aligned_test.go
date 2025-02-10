// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string

//  _____         _      __   _        _             
// |_   _|__  ___| |_   / /__| |_ _ __(_)_ __   __ _ 
//   | |/ _ \/ __| __| / / __| __| '__| | '_ \ / _` |
//   | |  __/\__ \ |_ / /\__ \ |_| |  | | | | | (_| |
//   |_|\___||___/\__/_/ |___/\__|_|  |_|_| |_|\__, |
//                                             |___/ 
import "testing"

func TestAligned(t *testing.T) {
	fn := "sisimai/string.Aligned"
	cw := "Final-Recipient: rfc822; <neko@example.jp>"
	cx := 0

	cx++; if Aligned(cw, []string{"rfc822", "<", "@", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", " ", "@", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", ";", "<", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"Final-", ":", ";", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", "[", "@", ">"}) == true  { t.Errorf("%s(%s) returns true",  fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", "<", "@", " "}) == true  { t.Errorf("%s(%s) returns true",  fn, cw) }

	cx++; if Aligned("", []string{})          == true { t.Errorf("%s(%s) returns true",  fn, "") }
	cx++; if Aligned("neko", []string{})      == true { t.Errorf("%s(%s) returns true",  fn, "neko") }
	cx++; if Aligned("", []string{"neko"})    == true { t.Errorf("%s(%s) returns true",  fn, "") }
	cx++; if Aligned("cat", []string{"neko"}) == true { t.Errorf("%s(%s) returns true",  fn, "cat") }

	t.Logf("The number of tests = %d", cx)
}

