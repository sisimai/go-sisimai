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

func TestIndexOnTheWay(t *testing.T) {
	fn := "sisimai/string.IndexOnTheWay"
	cx := 0
	ae := []struct {text string; find string; index int; expected int}{
		{"Date: Wed, 27 Apr 2022 15:45:12 +0900 (JST)", ":", 10, 25},
		{"From: MAILER-DAEMON@mx22.example.com (Mail Delivery System)", " ", 10, 36},
		{"To: postmaster@mx22.example.com (Postmaster)", ".", 10, 19},
		{"Subject: Postfix SMTP server: errors from localhost[127.0.0.1]", ":", 10, 28},
		{"Message-Id: <XMY2brWtcXzdFV-Id@mx22.example.com>", "Id", 10, 28},
	}
	for _, e := range ae {
		cx++; if cv := IndexOnTheWay(e.text, e.find, e.index); cv != e.expected {
			t.Errorf("%s(%s, %s, %d) returns %d", fn, e.text, e.find, e.index, cv)
		}
	}
	cx++; if cv := IndexOnTheWay("neko", ":", -1); cv != -1 { t.Errorf("%s(%s) returns %d", fn, "neko", cv) }
	cx++; if cv := IndexOnTheWay("neko", ":", 22); cv != -1 { t.Errorf("%s(%s) returns %d", fn, "neko", cv) }

	t.Logf("The number of tests = %d", cx)
}

