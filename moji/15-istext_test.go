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

func TestIsText(t *testing.T) {
	fn := "moji.IsText"
	cx := 0
	ae := []struct {text string; expected bool}{
		{"neko", true},
		{"",    false},
		{"nekochan-cat", true},
		{"\n\t\v\rneko", true},
		{"\x00nekocat", false},
		{"neko\x7fcat", false},
		{"neko    cat",  true},
		{"neko\x7ecat",  true},
		{"neko\xffcat", false},
	}

	for _, e := range ae {
		cx++; if cv := IsText(&(e.text)); cv != e.expected {
			t.Errorf("%s(%s) returns %t", fn, e.text, !e.expected)
		}
	}
	cx++; if cv := IsText(nil); cv == true { t.Errorf("%s(nil) returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}

