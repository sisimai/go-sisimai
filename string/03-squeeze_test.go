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

func TestSqueeze(t *testing.T) {
	fn := "sisimai/string.Squeeze"
	cx := 0
	ae := []struct {text string; char string; expected string}{
		{"neko		meow	cat", "	", "neko	meow	cat"},
		{"neko      meow   cat", " ", "neko meow cat"},
		{"neko//////meow///cat", "/", "neko/meow/cat"},
		{"neko::meow:::::::cat", ":", "neko:meow:cat"},
		{"nekochan", "", "nekochan"},
		{"nekonekopoint", "neko", "nekopoint"},
		{"", "?", ""},
	}
	for _, e := range ae {
		cx++; if Squeeze(e.text, e.char) != e.expected { t.Errorf("%s(%s, %s) returns %s", fn, e.text, e.char, e.expected) }
	}
	t.Logf("The number of tests = %d", cx)
}

