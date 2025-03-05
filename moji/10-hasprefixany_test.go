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

func TestHasPrefixAny(t *testing.T) {
	fn := "moji.HasPrefixAny"
	cx := 0
	ae := []struct {text string; list []string; expected bool}{
		{"nekochan", []string{"cats", "kijitora", "nekochan"}, true},
		{"From: <postmaster@example.jp>", []string{"@", "From"}, true},
		{"Date: Fri,  2 Feb 2018 18:30:22 +0900 (JST)", []string{"Feb", "Apr"}, false},
		{"Subject: Delivery failure", []string{"Subject:", "Failed"}, true},
	}

	for _, e := range ae {
		cx++; if cv := HasPrefixAny(e.text, e.list); cv != e.expected {
			t.Errorf("%s(%s, %v) returns %t", fn, e.text, e.list, e.expected)
		}
	}
	if cv := HasPrefixAny("", []string{});    cv == true { t.Errorf("%s('', []) returns true", fn) }
	if cv := HasPrefixAny("", []string{"2"}); cv == true { t.Errorf("%s('', [2]) returns true", fn) }
	if cv := HasPrefixAny("2", []string{});   cv == true { t.Errorf("%s(2, []) returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}

