// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//  _____         _      __        _     _                     ___     
// |_   _|__  ___| |_   / /_ _  __| | __| |_ __ ___  ___ ___  |_ _|___ 
//   | |/ _ \/ __| __| / / _` |/ _` |/ _` | '__/ _ \/ __/ __|  | |/ __|
//   | |  __/\__ \ |_ / / (_| | (_| | (_| | | |  __/\__ \__ \_ | |\__ \
//   |_|\___||___/\__/_/ \__,_|\__,_|\__,_|_|  \___||___/___(_)___|___/
import "testing"
import "strings"

func TestIsQuotedAddress(t *testing.T) {
	fn := "sisimai/address.IsQuotedAddress"
	cx := 0
	ae := []struct {testname string; argument string; expected bool}{
		{"", `"neko@example@jp"@example.org`, true},
		{"", `"neko\ miaow"@example.co.uk`, true},
		{"", "neko@example.jp", false},
	}

	for _, e := range ae {
		t.Run(e.testname, func(t *testing.T) {
			cv := IsQuotedAddress(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s(%s) is (%v) not (%v)", cx, fn, e.argument, cv, e.expected) }
			cx += 1
		})
	}
	t.Logf("The number of tests = %d", cx)
}

func TestIsIncluded(t *testing.T) {
	fn := "sisimai/address.IsIncluded"
	cx := 0
	ae := []struct {testname string; argument string; expected bool}{
		{"", "<neko@example.jp>", true},
		{"", "<n@e.jp>", true},
		{"", "Kijitora neko@example.jp (Nekochan)", true},
		{"", "Sironeko <siro@example.jp> (Meow)", true},
		{"", "<mailer-daemon>", false},
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

func TestIsComment(t *testing.T) {
	fn := "sisimai/address.IsComment"
	cx := 0
	ae := []struct {testname string; argument string; expected bool}{
		{"", "", false},
		{"", "<neko@example.jp>", false},
		{"", "(Kijitora neko)", true},
		{"", "(Siro(neko))(Meow)", true},
	}

	for _, e := range ae {
		t.Run(e.testname, func(t *testing.T) {
			cv := IsComment(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s(%s) is (%v) not (%v)", cx, fn, e.argument, cv, e.expected) }
			cx += 1
		})
	}
	t.Logf("The number of tests = %d", cx)
}

func TestIsDomainLiteral(t *testing.T) {
	fn := "sisimai/address.IsDomainLiteral"
	cx := 0
	ae := []struct {testname string; argument string; expected bool}{
		{"", "", false},
		{"", "<neko@example.jp>", false},
		{"", "<neko@[IPv4:192.0.2.25]>", true},
		{"", "neko@[IPv4:192.0.2.25]", true},
		{"", "neko@[Neko:192.0.2.25]", false},
		{"", "neko@[IPv6:192.0.2.25]", false},
		{"", "neko@[IPv6:2001:DB8::1]", false},
		{"", "neko@[IPv5:2001:DB8::1]", false},
		{"", "<neko@[IPv6:2001:DB8::1]>", false},
		{"", "neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]", true},
		{"", "<neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]>", true},
	}

	for _, e := range ae {
		t.Run(e.testname, func(t *testing.T) {
			cv := IsDomainLiteral(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s(%s) is (%v) not (%v)", cx, fn, e.argument, cv, e.expected) }
			cx += 1
		})
	}
	t.Logf("The number of tests = %d", cx)
}

func TestIsEmailAddress(t *testing.T) {
	fn := "sisimai/address.IsEmailAddress"
	cx := 0

	for _, e := range TestEmailAddrs {
		t.Run(e.testname, func(t *testing.T) {
			cv := IsEmailAddress(e.expected)
			if strings.Contains(e.expected, "@") && cv == false {
				t.Errorf("[%6d]: %s(%s) is (false) not (true)", cx, fn, e.expected)
			}; cx++
		})
	}
	t.Logf("The number of tests = %d", cx)
}

