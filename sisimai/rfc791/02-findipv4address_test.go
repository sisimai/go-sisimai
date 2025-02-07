// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc791

//  _____         _      ______  _____ ____ _____ ___  _ 
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___  / _ \/ |
//   | |/ _ \/ __| __| / /| |_) | |_ | |      / / (_) | |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___  / / \__, | |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|/_/    /_/|_|
import "testing"

func TestFindIPv4Address(t *testing.T) {
	fn := "sisimai/rfc791.FindIPv4Address"
	cx := 0
	ae := []struct {text string; expected string}{
		{"host smtp.example.jp 127.0.0.4 SMTP error from remote mail server", "127.0.0.4"},
		{"mx.example.jp (192.0.2.2) reason: 550 5.2.0 Mail rejete.", "192.0.2.2"},
		{"Client host [192.0.2.49] blocked using cbl.abuseat.org (state 13).", "192.0.2.49"},
		{"127.0.0.1", "127.0.0.1"},
	}
	ce := []string{"365.31.7.1", "a.b.c.d", ""}

	for _, e := range ae {
		cv := FindIPv4Address(&e.text)
		cx++; if len(cv) == 0          { t.Errorf("%s(%s) returns empty", fn, e.text)     } 
		cx++; if cv[0] != e.expected   { t.Errorf("%s(%s) returns %s", fn, e.text, cv[0]) }
		cx++; if !IsIPv4Address(cv[0]) { t.Errorf("IsIPv4Address(%s) is false", cv[0])    }
	}
	for _, e := range ce {
		cv := FindIPv4Address(&e)
		cx++; if len(cv) != 0 { t.Errorf("%s(%s) returns %v", fn, e, cv)  } 
	}
	cx++; if cv := FindIPv4Address(nil); len(cv) > 0 { t.Errorf("%s(nil) returns %v", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

