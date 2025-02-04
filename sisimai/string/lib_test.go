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

func TestContainsOnlyNumbers(t *testing.T) {
	fn := "sisimai/string.ContainsOnlyNumbers"
	cx := 0
	et := []string{"1", "23", "456", "78910"}
	ef := []string{"A", "B1", "C12", "34D5E"}

	for _, e := range et {
		cx++; if cv := ContainsOnlyNumbers(e); cv == false { t.Errorf("%s(%s) returns false", fn, e) }
	}
	for _, e := range ef {
		cx++; if cv := ContainsOnlyNumbers(e); cv == true  { t.Errorf("%s(%s) returns true",  fn, e) }
	}
	cx++; if cv := ContainsOnlyNumbers("");    cv == true  { t.Errorf("%s(%s) returns true",  fn, "") }

	t.Logf("The number of tests = %d", cx)
}

func TestAligned(t *testing.T) {
	fn := "sisimai/string.Aligned"
	cw := "Final-Recipient: rfc822; <neko@example.jp>"
	cx := 0

	cx++; if Aligned(cw, []string{"rfc822", "<", "@", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", " ", "@", ">"}) == false { t.Errorf("%s(%s) returns false",  fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", ";", "<", ">"}) == false { t.Errorf("%s(%s) returns false",  fn, cw) }
	cx++; if Aligned(cw, []string{"Final-", ":", ";", ">"}) == false { t.Errorf("%s(%s) returns false",  fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", "[", "@", ">"}) == true  { t.Errorf("%s(%s) returns true",  fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", "<", "@", " "}) == true  { t.Errorf("%s(%s) returns true",  fn, cw) }

	cx++; if Aligned("", []string{})          == true { t.Errorf("%s(%s) returns true",  fn, "") }
	cx++; if Aligned("neko", []string{})      == true { t.Errorf("%s(%s) returns true",  fn, "neko") }
	cx++; if Aligned("", []string{"neko"})    == true { t.Errorf("%s(%s) returns true",  fn, "") }
	cx++; if Aligned("cat", []string{"neko"}) == true { t.Errorf("%s(%s) returns true",  fn, "cat") }

	t.Logf("The number of tests = %d", cx)
}

