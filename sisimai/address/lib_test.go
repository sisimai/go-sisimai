// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//  _____         _      __        _     _                      ___ _ _     
// |_   _|__  ___| |_   / /_ _  __| | __| |_ __ ___  ___ ___   / / (_) |__  
//   | |/ _ \/ __| __| / / _` |/ _` |/ _` | '__/ _ \/ __/ __| / /| | | '_ \ 
//   | |  __/\__ \ |_ / / (_| | (_| | (_| | | |  __/\__ \__ \/ / | | | |_) |
//   |_|\___||___/\__/_/ \__,_|\__,_|\__,_|_|  \___||___/___/_/  |_|_|_.__/ 
import "testing"

var AE = []struct {testname string; argument string; expected string; displays string; comments string}{
	//{"test name", "the argument", "email address", "display name", "comment"}
	{"", `"Neko" <neko@example.jp>`, "neko@example.jp", "Neko", ""},
	{"", `"=?ISO-2022-JP?B?dummy?=" <nyan@example.jp>`, "nyan@example.jp", "=?ISO-2022-JP?B?dummy?=", ""},
	{"", `"N Y A N K O" <nyanko@example.jp>`, "nyanko@example.jp", "N Y A N K O", ""},
	{"", `"Shironeko Lui" <lui@example.jp>`, "lui@example.jp", "Shironeko Lui", ""},
	{"", "<aoi@example.jp>", "aoi@example.jp", "", ""},
	{"", "<may@example.jp> may@example.jp", "may@example.jp", "may@example.jp", ""},
	{"", "Odd-Eyes Aoki <aoki@example.jp>", "aoki@example.jp", "Odd-Eyes Aoki", ""},
	{"", "Mikeneko Shima <shima@example.jp> SHIMA@EXAMPLE.JP", "shima@example.jp", "Mikeneko Shima SHIMA@EXAMPLE.JP", ""},
	{"", "chosuke@neko <chosuke@example.jp>", "chosuke@example.jp", "chosuke@neko", ""},
	{"", "akari@chatora.neko <akari@example.jp>", "akari@example.jp", "akari@chatora.neko", ""},
	{"", "mari <mari@example.jp> mari@host.int", "mari@example.jp", "mari mari@host.int", ""},
	{"", "8suke@example.gov (Mayuge-Neko)", "8suke@example.gov", "8suke@example.gov", "(Mayuge-Neko)"},
	{"", "Shibainu Hachibe. (Harima-no-kami) 8be@example.gov", "8be@example.gov", "Shibainu Hachibe. 8be@example.gov", "(Harima-no-kami)"},
	{"", "neko(nyaan)chan@example.jp", "nekochan@example.jp", "nekochan@example.jp", "(nyaan)"},
	{"", "(nyaan)neko@example.jp", "neko@example.jp", "neko@example.jp", "(nyaan)"},
	{"", "neko(nyaan)@example.jp", "neko@example.jp", "neko@example.jp", "(nyaan)"},
	{"", "nora(nyaan)neko@example.jp(cat)", "noraneko@example.jp", "noraneko@example.jp", "(nyaan) (cat)"},
	{"", "<neko@example.com>:", "neko@example.com", ":", ""},
	{"", `"<neko@example.org>"`, "neko@example.org", "", ""},
	{"", `"neko@example.net"`, "neko@example.net", "neko@example.net", ""},
	{"", "'neko@example.edu'", "neko@example.edu", "'neko@example.edu'", ""},
	{"", "`neko@example.cat`", "neko@example.cat", "`neko@example.cat`", ""},
	{"", "[neko@example.gov]", "neko@example.gov", "[neko@example.gov]", ""},
	{"", "{neko@example.int}", "neko@example.int", "{neko@example.int}", ""},
	{"", `"neko.."@example.jp`, `"neko.."@example.jp`, `"neko.."@example.jp`, ""},
	{"", "Mail Delivery Subsystem <MAILER-DAEMON>", "MAILER-DAEMON", "Mail Delivery Subsystem", ""},
	{"", "postmaster", "postmaster", "postmaster", ""},
	{"", "neko.nyaan@example.com", "neko.nyaan@example.com", "neko.nyaan@example.com", ""},
	{"", "neko.nyaan+nyan@example.com", "neko.nyaan+nyan@example.com", "neko.nyaan+nyan@example.com", ""},
	{"", "neko-nyaan@example.com.", "neko-nyaan@example.com", "neko-nyaan@example.com.", ""},
	{"", "neko-nyaan@example.org.", "neko-nyaan@example.org", "neko-nyaan@example.org.", ""},
	{"", "n@example.com", "n@example.com", "n@example.com", ""},
	{"", `"neko.nyaan.@.nyaan.jp"@example.com`, `"neko.nyaan.@.nyaan.jp"@example.com`, `"neko.nyaan.@.nyaan.jp"@example.com`, ""},
	{"", `"neko nyaan"@example.org`, `"neko nyaan"@example.org`, `"neko nyaan"@example.org`, ""},
	{"", "neko@nyaan", "", "neko@nyaan", ""},
	{"", "neko(1)-nyaan(2)@exa(3)mp(4)le.j(5)p", "neko-nyaan@example.jp", "neko-nyaan@example.jp", "(1) (2) (3) (4) (5)"},
	{"", "#!$%&'*+-/=?^_`{}|~@example.org", "#!$%&'*+-/=?^_`{}|~@example.org", "#!$%&'*+-/=?^_`{}|~@example.org", ""},
	{"", `" "@example.org`, `" "@example.org`, `" "@example.org`, ""},
	{"", "neko@localhost", "", "neko@localhost", ""},
	{"", "neko@[IPv4:192.0.2.22]", "neko@[IPv4:192.0.2.22]", "neko@[IPv4:192.0.2.22]", ""},
	{"", "neko@[IPv6:2001:DB8::1]", "", "neko@[IPv6:2001:DB8::1]", ""},
	{"", "neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]", "neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]",
		 "neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]", ""},
}

func TestFind(t *testing.T) {
	fn := "sisimai/address.Find()"
	cx := 0

	for _, e := range AE {
		//if e.testname != "D" { continue }
		t.Run(e.testname, func(t *testing.T) {
			cv := Find(e.argument)
			if cv[0] != e.expected { t.Errorf("[%6d]: %s [0:address] is (%s) not (%s)", cx, fn, cv[0], e.expected) }; cx += 1
			if cv[1] != e.displays { t.Errorf("[%6d]: %s [1:display] is (%s) not (%s)", cx, fn, cv[1], e.displays) }; cx += 1
			if cv[2] != e.comments { t.Errorf("[%6d]: %s [2:comment] is (%s) not (%s)", cx, fn, cv[2], e.comments) }; cx += 1
		})
	}
	t.Logf("The number of tests = %d", cx)
}

func TestUndisclosed(t *testing.T) {
	fn := "sisimai/address.Undisclosed()"
	cx := 0
	ae := []struct {testname string; argument bool; expected string}{
		{"",  true, "undisclosed-recipient-in-headers@libsisimai.org.invalid"},
		{"", false, "undisclosed-sender-in-headers@libsisimai.org.invalid"},
	}

	for _, e := range ae {
		t.Run(e.testname, func(t *testing.T) {
			cv := Undisclosed(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s is (%s) not (%s)", cx, fn, cv, e.expected) }
			cx += 1
		})
	}
	t.Logf("The number of tests = %d", cx)
}

func TestFinal(t *testing.T) {
	fn := "sisimai/address.Final()"
	cx := 0
	ae := []struct{testname string; argument string; expected string}{
		{"", "", ""},
		{"", "<neko>", "<neko>"},
		{"", "<neko@example.jp>", "neko@example.jp"},
		{"",  "<neko@example.jp",  "neko@example.jp"},
		{"",  "neko@example.jp>",  "neko@example.jp"},
	}
	for _, e := range ae {
		t.Run(e.testname, func(t *testing.T) {
			cv := Final(e.argument)
			if cv != e.expected { t.Errorf("[%6d]: %s is (%s) not (%s)", cx, fn, cv, e.expected) }
			cx += 1
		})
	}
	t.Logf("The number of tests = %d", cx)
}

