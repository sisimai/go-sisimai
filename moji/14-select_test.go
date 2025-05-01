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

func TestSelect(t *testing.T) {
	fn := "moji.Select"
	cx := 0
	ae := []struct {arg string; b string; u string; s int; exp string}{
		{"From: <neko@example.jp> Kijitora", "<", ">", 0, "neko@example.jp"},
		{"Diagnostic-Code: smtp;550 5.1.1 <neko@example.jp>... User Unknown", " ", ";", 4, "smtp"},
		{"From: <neko@example.jp>\nTo: <cat@example.org>\n ", "\nTo: ", "\n", 0, "<cat@example.org>"},
		{"Status: 4.4.0 (undefined routing status)", " (", ")", 1, "undefined routing status"},
		{"550-5.7.26 The MAIL FROM domain [email.example.jp] has an SPF", " [", "] ", 10, "email.example.jp"},
		{LHS + "nekochan" + RHS, "", "", 0, "nekochan"},
		{LHS + "nekochan:", "", ":", 0, "nekochan"},
		{":nekochan" + RHS, ":", "", 0, "nekochan"},
	}
	je := []struct {arg string; b string; u string; s int; exp string}{
		{"From: <neko@example.jp> Kijitora", "(", ">", 0, ""},
		{"From: <neko@example.jp> Kijitora", "<", ")", 0, ""},
		{"", "", "", -1, "" },
		{"n", "", "", -1, "" },
		{"n", "e", "", -1, "" },
		{"n", "e", "k", -1, "" },
		{"n", "e", "k",  0, "" },
	}

	for _, e := range ae {
		cx++; if cv := Select(e.arg, e.b, e.u, e.s); cv != e.exp {
			t.Errorf("%s(%s..., %s, %s, %d) returns [%s]", fn, e.arg[e.s:e.s + 10], e.b, e.u, e.s, cv)
		}
		cx++; if cv := Select(e.arg, e.b, e.u, e.s * 10 + 50); cv == e.exp {
			t.Errorf("%s(%s..., %s, %s, %d) returns [%s]", fn, e.arg[e.s:e.s + 10], e.b, e.u, e.s, e.exp)
		}
	}

	for _, e := range je {
		cx++; if cv := Select(e.arg, e.b, e.u, e.s); cv != e.exp {
			t.Errorf("%s(%s..., %s, %s, %d) returns [%s]", fn, e.arg[e.s:e.s + 10], e.b, e.u, e.s, cv)
		}
	}

	t.Logf("The number of tests = %d", cx)
}

