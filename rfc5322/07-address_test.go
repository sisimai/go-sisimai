// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
import "testing"
import "strings"

var TestEmailAddrs = []struct {testname string; argument string; expected string; displays string; comments string}{
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
	{"", "neko(miaow)chan@example.jp", "nekochan@example.jp", "nekochan@example.jp", "(miaow)"},
	{"", "(miaow)neko@example.jp", "neko@example.jp", "neko@example.jp", "(miaow)"},
	{"", "neko(miaow)@example.jp", "neko@example.jp", "neko@example.jp", "(miaow)"},
	{"", "nora(miaow)neko@example.jp(cat)", "noraneko@example.jp", "noraneko@example.jp", "(miaow) (cat)"},
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
	{"", "neko.miaow@example.com", "neko.miaow@example.com", "neko.miaow@example.com", ""},
	{"", "neko.miaow+nyan@example.com", "neko.miaow+nyan@example.com", "neko.miaow+nyan@example.com", ""},
	{"", "neko-miaow@example.com.", "neko-miaow@example.com", "neko-miaow@example.com.", ""},
	{"", "neko-miaow@example.org.", "neko-miaow@example.org", "neko-miaow@example.org.", ""},
	{"", "n@example.com", "n@example.com", "n@example.com", ""},
	{"", `"neko.miaow.@.esmtp.jp"@example.com`, `"neko.miaow.@.esmtp.jp"@example.com`, `"neko.miaow.@.esmtp.jp"@example.com`, ""},
	{"", `"neko miaow"@example.org`, `"neko miaow"@example.org`, `"neko miaow"@example.org`, ""},
	{"", "neko@miaow", "", "neko@miaow", ""},
	{"", "neko(1)-miaow(2)@exa(3)mp(4)le.j(5)p", "neko-miaow@example.jp", "neko-miaow@example.jp", "(1) (2) (3) (4) (5)"},
	{"", "#!$%&'*-/=?^_`{}|~@example.org", "#!$%&'*-/=?^_`{}|~@example.org", "#!$%&'*-/=?^_`{}|~@example.org", ""},
	{"", `" "@example.org`, `" "@example.org`, `" "@example.org`, ""},
	{"", "neko@localhost", "neko@localhost", "neko@localhost", ""},
	{"", "neko@[IPv4:192.0.2.22]", "neko@[IPv4:192.0.2.22]", "neko@[IPv4:192.0.2.22]", ""},
	{"", "neko@[IPv6:2001:DB8::1]", "", "neko@[IPv6:2001:DB8::1]", ""},
	{"", "neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]", "neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]",
		 "neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]", ""},
}

func TestIsEmailAddress(t *testing.T) {
	fn := "sisimai/rfc5322.IsEmailAddress"
	cx := 0

	for _, e := range TestEmailAddrs {
		t.Run(e.testname, func(t *testing.T) {
			cv := IsEmailAddress(e.expected)
			if strings.Contains(e.expected, "@") && cv == false {
				t.Errorf("[%6d]: %s(%s) is (false) not (true)", cx, fn, e.expected)
			}; cx++
		})
	}
	cw := "nyaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaan@example.jp"
	cx++; if IsEmailAddress(cw) == true { t.Errorf("%s(%s) returns true", fn, cw[0:25]) }

	cw  = "neko@nyaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaan.jp"
	cx++; if IsEmailAddress(cw) == true { t.Errorf("%s(%s) returns true", fn, cw[0:25]) }

	t.Logf("The number of tests = %d", cx)
}

func TestIsQuotedAddress(t *testing.T) {
	fn := "sisimai/rfc5322.IsQuotedAddress"
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

func TestIsComment(t *testing.T) {
	fn := "sisimai/rfc5322.IsComment"
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

