//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package moji

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   ___ (_|_)
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \| | |
//   | |  __/\__ \ |_ / /| | | | | | (_) | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\___// |_|
//                                     |__/   
import "testing"

// % go test -tags=fuzz -fuzz=FuzzAligned -fuzztime=1m ./moji
func FuzzAligned(f *testing.F) {
	stringlist := []string{
		"[host smtp.example.jp 127.0.0.4 SMTP error from remote mail server]",
		"<mx.example.jp (192.0.2.2) reason: 550 5.2.0 Mail rejete.>",
		"{Client host [192.0.2.49] blocked using cbl.abuseat.org (state 13).}",
		"[127.0.0.1]",
	}
	for _, e := range stringlist { f.Add(e) }
	f.Fuzz(func(t *testing.T, text string) { _ = Aligned(text, []string{":", ".jp"}) })
}

// % go test -tags=fuzz -fuzz=FuzzAlignedAny -fuzztime=1m ./moji
func FuzzAlignedAny(f *testing.F) {
	stringlist := []string{
		"[host smtp.example.jp 127.0.0.4 SMTP error from remote mail server]",
		"<mx.example.jp (192.0.2.2) reason: 550 5.2.0 Mail rejete.>",
		"{Client host [192.0.2.49] blocked using cbl.abuseat.org (state 13).}",
		"[127.0.0.1]",
	}
	for _, e := range stringlist { f.Add(e) }
	f.Fuzz(func(t *testing.T, text string) { _ = AlignedAny(text, [][]string{[]string{":", ".jp"}}) })
}
