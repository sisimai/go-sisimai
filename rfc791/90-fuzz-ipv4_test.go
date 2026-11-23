//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ______  _____ ____ _____ ___  _ 
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___  / _ \/ |
//   | |/ _ \/ __| __| / /| |_) | |_ | |      / / (_) | |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___  / / \__, | |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|/_/    /_/|_|
package rfc791
import "testing"

// % go test -tags=fuzz -fuzz=FuzzIsIPv4Address -fuzztime=1m ./rfc791
func FuzzIsIPv4Address(f *testing.F) {
	ip4address := []string{
		"0.0.0.0",
		"1.2.3.4",
		"127.0.0.1",
		"192.0.2.25",
		"255.255.255.255",
		"13597846544569",
		"3.14",
		"3.14",
		"8.17.1",
		"5.6.7.8.9",
		"123.456.78.9",
	}

	for _, e := range ip4address { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, addr string) { _ = IsIPv4Address(addr) })
}

// % go test -tags=fuzz -fuzz=FuzzFindIPv4Address -fuzztime=1m ./rfc791
func FuzzFindIPv4Address(f *testing.F) {
	ip4address := []string{
		"host smtp.example.jp 127.0.0.4 SMTP error from remote mail server",
		"mx.example.jp (192.0.2.2) reason: 550 5.2.0 Mail rejete.",
		"Client host [192.0.2.49] blocked using cbl.abuseat.org (state 13).",
		"127.0.0.1",
	}

	for _, e := range ip4address { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, addr string) { _ = FindIPv4Address(addr) })
}

