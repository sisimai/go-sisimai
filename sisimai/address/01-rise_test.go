// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//  _____         _      __        _     _                     ____  _          
// |_   _|__  ___| |_   / /_ _  __| | __| |_ __ ___  ___ ___  |  _ \(_)___  ___ 
//   | |/ _ \/ __| __| / / _` |/ _` |/ _` | '__/ _ \/ __/ __| | |_) | / __|/ _ \
//   | |  __/\__ \ |_ / / (_| | (_| | (_| | | |  __/\__ \__ \_|  _ <| \__ \  __/
//   |_|\___||___/\__/_/ \__,_|\__,_|\__,_|_|  \___||___/___(_)_| \_\_|___/\___|
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
var TestPostmaster = []string{
	"mailer-daemon@example.jp", 
	"MAILER-DAEMON@example.cat",
	"Mailer-Daemon <postmaster@example.org>",
	"MAILER-DAEMON",
	"postmaster",
	"postmaster@example.org",
}
var TestNotAnEmail = []string{"1", "neko", `neko%example.jp`, ""}

func TestRise(t *testing.T) {
	on := "EmailAddress"
	cx := 0

	for _, e := range TestEmailAddrs {
		t.Run(e.testname, func(t *testing.T) {
			cw := Find(e.argument)
			cv := Rise(cw)

			if cv.Void() == true {
				// EmailAddress{} is empty
				if cv.Address != "" { t.Errorf("[%6d]: %s.Address is (%s) not empty", cx, on, cv.Address) }; cx++
				if cv.User    != "" { t.Errorf("[%6d]: %s.User is (%s) not empty", cx, on, cv.User)       }; cx++
				if cv.Host    != "" { t.Errorf("[%6d]: %s.Host is (%s) not empty", cx, on, cv.Host)       }; cx++
				if cv.Verp    != "" { t.Errorf("[%6d]: %s.Verp is (%s) not empty", cx, on, cv.Verp)       }; cx++
				if cv.Alias   != "" { t.Errorf("[%6d]: %s.Alias is (%s) not empty", cx, on, cv.Alias)     }; cx++
				if cv.Name    != "" { t.Errorf("[%6d]: %s.Name is (%s) not empty", cx, on, cv.Name)       }; cx++
				if cv.Comment != "" { t.Errorf("[%6d]: %s.Comment is (%s) not empty", cx, on, cv.Comment) }; cx++

			} else {
				// EmailAddress{} is not empty
				if cv.Address != e.expected && cv.Alias == "" && cv.Verp == "" {
					t.Errorf("[%6d]: %s.Address is (%s) not (%s)", cx, on, cv.Address, e.expected); cx++
				}; cx++

				if strings.HasPrefix(e.expected, cv.User) == false {
					t.Errorf("[%6d]: %s.User is (%s), does not have the local part of (%s)", cx, on, cv.User, e.expected)
				}; cx++

				if strings.HasSuffix(e.expected, cv.Host) == false {
					t.Errorf("[%6d]: %s.Host is (%s), does not have the domain part of (%s)", cx, on, cv.Host, e.expected)
				}; cx++

				if cv.Verp != "" {
					if strings.Count(cv.Address, "=") < 1 { t.Errorf("[%6d]: %s.Address does not include '=' (%s)", cx, on, cv.Address) }; cx++
					if strings.Count(cv.Verp, "=")    > 0 { t.Errorf("[%6d]: %s.Verp includes '=' (%s)", cx, on, cv.Verp) }; cx++
				}

				if cv.Alias != "" {
					if strings.Count(cv.Address, "+") < 1 { t.Errorf("[%6d]: %s.Address does not include '+' (%s)", cx, on, cv.Address) }; cx++
					if strings.Count(cv.Alias, "+")   > 0 { t.Errorf("[%6d]: %s.Alias includes '+' (%s)", cx, on, cv.Alias) }; cx++
				}

				if cv.Name    != e.displays { t.Errorf("[%6d]: %s.Name is (%s) not (%s)", cx, on, cv.Name, e.displays)    }; cx++
				if cv.Comment != e.comments { t.Errorf("[%6d]: %s.Comment is (%s) not (%s)", cx, on, cv.Name, e.comments) }; cx++
			}
		})
	}

	for _, e := range TestPostmaster {
		t.Run("", func(t *testing.T) {
			cv := Rise([3]string{e, "", ""})
			if cv.Void() == true { t.Errorf("[%6d]: %s.Void is true not false (%s)", cx, on, e) }; cx++
		})
	}

	for _, e := range TestNotAnEmail {
		t.Run("", func(t *testing.T) {
			cv := Rise([3]string{e, "", ""})
			if cv.Void() == false { t.Errorf("[%6d]: %s.Void is false not true (%s)", cx, on, e) }; cx++
		})
	}
	t.Logf("The number of tests = %d", cx)
}

