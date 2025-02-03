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

func TestToken(t *testing.T) {
	fn := "sisimai/string.Token"
	es := "envelope-sender@example.jp"
	er := "envelope-recipient@example.org"
	to := "239aa35547613b2fa94f40c7f35f4394e99fdd88"
	cx := 0

	cx++; if Token(es, er, 1) != to { t.Errorf("%s(%s, %s, 1) returns %s", fn, es, er, Token(es,er, 1)) }
	cx++; if Token("", "", 0) != "" { t.Errorf("%s('', '', 0) returns %s", fn, Token("", "", 0)) }
	cx++; if Token(es, "", 0) != "" { t.Errorf("%s(%s, '', 0) returns %s", fn, es, Token("", "", 0)) }
	cx++; if Token("", er, 0) != "" { t.Errorf("%s('', %s, 0) returns %s", fn, er, Token("", "", 0)) }

	t.Logf("The number of tests = %d", cx)
}

func TestIs8Bit(t *testing.T) {
	fn := "sisimai/string.Is8Bit"
	ae := []string{"nekochan", "Suzu", ""}
	je := []string{"ニャーン", "道綱"}
	cx := 0

	for _, e := range ae { cx++; if Is8Bit(&e) == true  { t.Errorf("%s(%s) returns true", fn, e) } }
	for _, e := range je { cx++; if Is8Bit(&e) == false { t.Errorf("%s(%s) returns false", fn, e) } }

	t.Logf("The number of tests = %d", cx)
}

func TestSqueeze(t *testing.T) {
	fn := "sisimai/string.Sweep"
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

