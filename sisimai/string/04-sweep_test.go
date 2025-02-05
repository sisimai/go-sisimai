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

func TestSweep(t *testing.T) {
	fn := "sisimai/string.Sweep"
	cx := 0
	ae := []struct {arg string; exp string}{
		{" neko		meow	cat ", "neko meow cat"},
		{"neko      meow   cat --nekochan kijitora", "neko meow cat"},
		{"", ""},
	}
	for _, e := range ae {
		cx++; if cv := Sweep(e.arg); cv != e.exp { t.Errorf("%s(%s) returns %s", fn, e.arg, cv) }
	}

	t.Logf("The number of tests = %d", cx)
}

