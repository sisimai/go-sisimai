//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
package sisimai
import "testing"

func FuzzRise(f *testing.F) {
	headerseed := []string{
		"From: Mail Delivery Subsystem <postmaster@libsisimai.org>",
		"To: Kijitora-Neko <kijitora@libsisimai.net>",
		"Subject: Returned mail: see transcript for details",
		"Date: Fri,  2 Feb 2018 18:30:22 +0900 (JST)",
		"Message-Id: <neko>",
		"MIME-Version: 1.0",
	}
	bodyasseed := "Nyaan?"

	for _, e := range headerseed { f.Add([]byte(e + "\n")) }
	f.Add([]byte("\n" + bodyasseed + "\n"))
	f.Fuzz(func(t *testing.T, data []byte) { _, _ = Rise(string(data), nil) })
}

