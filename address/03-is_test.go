// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//  _____         _      __        _     _                     ___     
// |_   _|__  ___| |_   / /_ _  __| | __| |_ __ ___  ___ ___  |_ _|___ 
//   | |/ _ \/ __| __| / / _` |/ _` |/ _` | '__/ _ \/ __/ __|  | |/ __|
//   | |  __/\__ \ |_ / / (_| | (_| | (_| | | |  __/\__ \__ \_ | |\__ \
//   |_|\___||___/\__/_/ \__,_|\__,_|\__,_|_|  \___||___/___(_)___|___/
import "testing"

func TestIsIncluded(t *testing.T) {
	fn := "sisimai/address.IsIncluded"
	cx := 0
	ae := []struct {testname string; argument string; expected bool}{
		{"", "<neko@example.jp>", true},
		{"", "<n@e.jp>", true},
		{"", "Kijitora neko@example.jp (Nekochan)", true},
		{"", "Sironeko <siro@example.jp> (Meow)", true},
		{"", "<mailer-daemon>", false},
		{"", "<neko@chan>", false},
		{"", "", false},
	}

	for _, e := range ae {
		t.Run(e.testname, func(t *testing.T) {
			cv := IsIncluded(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s(%s) is (%v) not (%v)", cx, fn, e.argument, cv, e.expected) }
			cx += 1
		})
	}
	t.Logf("The number of tests = %d", cx)
}

func TestIsMailerDaemon(t *testing.T) {
	fn := "sisimai/address.IsMailerDaemon"
	cx := 0

	for _, e := range TestPostmaster {
		t.Run("", func(t *testing.T) {
			cv := IsMailerDaemon(e)
			if cv == false { t.Errorf("[%6d]: %s(%s) is (false) not (true)", cx, fn, e) }; cx++
		})
	}
	t.Logf("The number of tests = %d", cx)
}

