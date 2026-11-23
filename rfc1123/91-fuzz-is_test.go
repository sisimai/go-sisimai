//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ______  _____ ____ _ _ ____  _____ 
// |_   _|__  ___| |_   / /  _ \|  ___/ ___/ / |___ \|___ / 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   | | | __) | |_ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___| | |/ __/ ___) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_|_|_____|____/ 
package rfc1123
import "testing"

// % go test -tags=fuzz -fuzz=FuzzIsInternetHost -fuzztime=1m ./rfc1123
func FuzzIsInternetHost(f *testing.F) {
	inet4hosts := []string{
		"",
		"127.0.0.1",
		"cat",
		"neko",
		".net",
		"-273.15",
		"nyaan.22",
		"mx0.example.22",
		"mx0.example.jp-",
		"mx--0.example.jp",
		"mx++0.example.jp",
		"mx__0.example.jp",
		"mx..0.example.jp",
		"mx0.example.jp/neko",
		"mx22.nyaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaan.jp",
		"mx25.jpppppppppppppppppppppppppppppppppppppppppppppppppppppppppppppppp",
		"localhost",
		"localhost6",
		"mx1.example.jp",
		"mx1.example.jp.",
		"a.jp",
		"mx22.nyaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaan.jp",
	}

	for _, e := range inet4hosts { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, host string) { _ = IsInternetHost(host) })
}

// % go test -tags=fuzz -fuzz=FuzzIsDomainLiteral -fuzztime=1m ./rfc1123
func FuzzIsDomainLiteral(f *testing.F) {
	domainlist := []string{
		"<neko@example.jp>",
		"<nekochan@example.org>",
		"<neko@[IPv4:192.0.2.25]>",
		"neko@[IPv4:192.0.2.25]",
		"neko@[Neko:192.0.2.25]",
		"neko@[IPv6:192.0.2.25]",
		"neko@[IPv6:2001:DB8::1]",
		"neko@[IPv5:2001:DB8::1]",
		"<neko@[IPv6:2001:DB8::1]>",
		"neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]",
		"<neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]>",
	}
	for _, e := range domainlist { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, host string) { _ = IsDomainLiteral(host) })
}

