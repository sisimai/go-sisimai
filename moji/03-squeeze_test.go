// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package moji

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   ___ (_|_)
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \| | |
//   | |  __/\__ \ |_ / /| | | | | | (_) | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\___// |_|
//                                     |__/   
import "testing"

func TestSqueeze(t *testing.T) {
	fn := "moji.Squeeze"
	cx := 0
	ae := []struct {text string; char byte; expected string}{
		{"neko		meow	cat", '	', "neko	meow	cat"},
		{"neko      meow   cat", ' ', "neko meow cat"},
		{"neko//////meow///cat", '/', "neko/meow/cat"},
		{"neko::meow:::::::cat", ':', "neko:meow:cat"},
		{"nekochan", ' ', "nekochan"},
		{"", '?', ""},
	}
	for _, e := range ae {
		cv := e.text
		cx++; if Squeeze(&cv, e.char); cv != e.expected { t.Errorf("%s(%s, %c) returns %s", fn, e.text, e.char, e.expected) }
	}
	t.Logf("The number of tests = %d", cx)
}

