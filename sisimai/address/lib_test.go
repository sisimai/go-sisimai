// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//  _____         _      __        _     _                      ___ _ _     
// |_   _|__  ___| |_   / /_ _  __| | __| |_ __ ___  ___ ___   / / (_) |__  
//   | |/ _ \/ __| __| / / _` |/ _` |/ _` | '__/ _ \/ __/ __| / /| | | '_ \ 
//   | |  __/\__ \ |_ / / (_| | (_| | (_| | | |  __/\__ \__ \/ / | | | |_) |
//   |_|\___||___/\__/_/ \__,_|\__,_|\__,_|_|  \___||___/___/_/  |_|_|_.__/ 
import "testing"

func TestUndisclosed(t *testing.T) {
	fn := "Undisclosed"
	cx := 0
	ae := []struct {testname string; argument bool; expected string}{
		{"(true)",  true,  "undisclosed-recipient-in-headers@libsisimai.org.invalid"},
		{"(false)", false, "undisclosed-sender-in-headers@libsisimai.org.invalid"},
	}

	for _, e := range ae {
		t.Run(fn + e.testname, func(t *testing.T) {
			cx += 1
			cv := Undisclosed(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s() = (%s)", cx, fn, cv) }
		})
	}
	t.Logf("The number of tests = %d", cx)
}

func TestFinal(t *testing.T) {
	fn := "Final"
	cx := 0
	ae := []struct{testname string; argument string; expected string}{
		{"('')", "", ""},
		{"(<neko>)", "<neko>", "<neko>"},
		{"(<neko@example.jp>)", "<neko@example.jp>", "neko@example.jp"},
		{"(<neko@example.jp)",  "<neko@example.jp",  "neko@example.jp"},
		{"(neko@example.jp>)",  "neko@example.jp>",  "neko@example.jp"},
	}
	for _, e := range ae {
		t.Run(fn + e.testname, func(t *testing.T) {
			cx += 1
			cv := Final(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s() = (%s)", cx, fn, cv) }
		})
	}
	t.Logf("The number of tests = %d", cx)
}

