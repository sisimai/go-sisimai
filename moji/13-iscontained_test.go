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

func TestIsContained(t *testing.T) {
	fn := "moji.IsContained"
	cx := 0
	ae := []struct {text string; list []string; expected bool}{
		{"neko", []string{"nekochan", "kijitora"}, true},
		{"postmaster@", []string{"From: <postmaster@example.jp>", "mailer-daemon"}, true},
		{"Feb ", []string{"Date: Fri,  2 Feb 2018 18:30:22 +0900 (JST)", "Sat, 14 Jun 2025 05:53:47 +0900 (JST)"}, true},
		{"Failed", []string{"Subject: Delivery failure", "Subject: Postmaster notify"}, false},
	}

	for _, e := range ae {
		cx++; if cv := IsContained(e.text, e.list); cv != e.expected {
			t.Errorf("%s(%s, %v) returns %t", fn, e.text, e.list, e.expected)
		}
	}
	if cv := IsContained("", []string{});    cv == true { t.Errorf("%s('', []) returns true", fn) }
	if cv := IsContained("", []string{"2"}); cv == true { t.Errorf("%s('', [2]) returns true", fn) }
	if cv := IsContained("2", []string{});   cv == true { t.Errorf("%s(2, []) returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}

