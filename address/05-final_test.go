// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//  _____         _      __        _     _                      ___ _ _     
// |_   _|__  ___| |_   / /_ _  __| | __| |_ __ ___  ___ ___   / / (_) |__  
//   | |/ _ \/ __| __| / / _` |/ _` |/ _` | '__/ _ \/ __/ __| / /| | | '_ \ 
//   | |  __/\__ \ |_ / / (_| | (_| | (_| | | |  __/\__ \__ \/ / | | | |_) |
//   |_|\___||___/\__/_/ \__,_|\__,_|\__,_|_|  \___||___/___/_/  |_|_|_.__/ 
import "testing"

func TestFinal(t *testing.T) {
	fn := "sisimai/address.Final()"
	cx := 0
	ae := []struct{testname string; argument string; expected string}{
		{"", "", ""},
		{"", "<neko>", "<neko>"},
		{"", "<neko@example.jp>", "neko@example.jp"},
		{"",  "<neko@example.jp",  "neko@example.jp"},
		{"",  "neko@example.jp>",  "neko@example.jp"},
	}
	for _, e := range ae {
		t.Run(e.testname, func(t *testing.T) {
			cv := Final(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s is (%s) not (%s)", cx, fn, cv, e.expected) }
			cx += 1
		})
	}

	t.Logf("The number of tests = %d", cx)
}

