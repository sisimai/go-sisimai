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
import "strings"

func FuzzRise(f *testing.F) {
	headerseed := []string{
		"Return-Path: <>",
		"From: Mail Delivery Subsystem <postmaster@libsisimai.org>",
		"To: Kijitora-Neko <kijitora@libsisimai.net>",
		"Subject: Returned mail: see transcript for details",
		"Date: Fri,  2 Feb 2018 18:30:22 +0900 (JST)",
		"Message-Id: <neko-michitsuna-20180202183022@nyaan.jp>",
		`Content-Type: multipart/report; report-type=delivery-status; boundary="nekochan"`,
		"MIME-Version: 1.0",
		"Auto-Submitted: auto-replied",
	}
	receivedby := [][]string{
		[]string{
			"from mx.libsisimai.org (ip-192-0-2-25.ap-northeast-1.compute.internal [192.0.2.25]) ",
			"by smtp.nyaan.jp (8.14.1/8.14.1) with ESMTPS id m8I8s45D007047 ",
			"for <kijitora@libsisimai.net>; ",
			"Fri,  2 Feb 2018 18:30:22 +0900 (JST)",
		},
		[]string{
			"from hub.libsisimai.org (ip-192-0-2-22.us-east-1.compute.internal [192.0.2.22]) ",
			"by mx2.nyaan.jp (Postfix) with ESMTPS id K1bgZJ3Hn6zmQTy6 ",
			"for <kijitora@libsisimai.net>; ",
			"Fri,  2 Feb 2018 18:30:22 +0900 (EST)",
		},
		[]string{"(qmail 3016 invoked for bounce); 25 Jun 2024 10:36:44 -0000"},
	}
	bodyasseed := "Nyaan?"

	for _, e := range headerseed { f.Add([]byte(e + "\n"))                                    }
	for _, e := range receivedby { f.Add([]byte("Received: " + strings.Join(e, "\n") + "\n")) }
	f.Add([]byte("\n" + bodyasseed + "\n"))
	f.Fuzz(func(t *testing.T, data []byte) { _, _ = Rise(string(data), nil) })
}

