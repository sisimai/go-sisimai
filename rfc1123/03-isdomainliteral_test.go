// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc1123

//  _____         _      ______  _____ ____ _ _ ____  _____ 
// |_   _|__  ___| |_   / /  _ \|  ___/ ___/ / |___ \|___ / 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   | | | __) | |_ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___| | |/ __/ ___) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_|_|_____|____/ 
import "testing"

func TestIsDomainLiteral(t *testing.T) {
	fn := "sisimai/rfc1123.IsDomainLiteral"
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

