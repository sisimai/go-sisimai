// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sisimai

//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
import "testing"

func TestReason(t *testing.T) {
	fn := "sisimai.Reason"
	cx := 0
	cv := Reason()

	cx++; if len(cv) ==  0 { t.Errorf("%s() returned an empty list", fn) }
	cx++; if len(cv) != 34 { t.Errorf("%s() returned invalid elements: %d", fn, len(cv)) }
	for e := range cv {
		cx++; if e == ""     { t.Errorf("%s returns an empty key", fn) }
		cx++; if cv[e] == "" { t.Errorf("%s[%s] is empty", fn, cv[e])  }
	}

	t.Logf("The number of tests = %d", cx)
}


