// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package moji

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   ___ (_|_)
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \| | |
//   | |  __/\__ \ |_ / /| | | | | | (_) | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\___// |_|
//                                     |__/   
import "testing"

func TestIs7bit(t *testing.T) {
	fn := "moji.Is7bit"
	ct := []string{"\n", "	", "", "", "\x7f"}
	b7 := []string{"[cat]", "(neko)", "/nyaaan???"}
	b8 := []string{"[猫]", "(ねこ)", "/ニャーン？"}
	cx := 0

	for _, e := range b7 {
		// ASCII Characters
		cx++; if Is7bit(e, false) == false { t.Errorf("%s(%s, false) returns false", fn, e) }

		for _, c := range ct {
			// Append each control character
			cv := e + c
			cx++; if Is7bit(cv, false) == false { t.Errorf("%s(%s, false) returns false", fn, cv) }
			cx++; if Is7bit(cv, true)  == true  { t.Errorf("%s(%s + %x, true) returns true", fn, cv, c) }
		}
	}

	for _, e := range b8 {
		// UTF-8 Characters
		cx++; if Is7bit(e, false) == true { t.Errorf("%s(%s, false) returns true", fn, e) }
		cx++; if Is7bit(e, true)  == true { t.Errorf("%s(%s, true) returns true", fn, e) }

		for _, c := range ct {
			// Append each control character
			cv := e + c
			cx++; if Is7bit(cv, false) == true { t.Errorf("%s(%s, false) returns true", fn, cv) }
			cx++; if Is7bit(cv, true)  == true { t.Errorf("%s(%s, true) returns true", fn, cv) }
		}
	}

	t.Logf("The number of tests = %d", cx)
}

