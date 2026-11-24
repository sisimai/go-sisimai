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

// % go test -tags=fuzz -fuzz=FuzzIsText -fuzztime=1m ./moji
func FuzzIsText(f *testing.F) {
	asciichars := []string{"\x00", "nekochan", "\x7F", "\r\n", "nyaan\x1F", "ネコ", "🐈"}
	for _, e := range asciichars { f.Add(e) }
	f.Fuzz(func(t *testing.T, text string) { _ = IsText(text) })
}

// % go test -tags=fuzz -fuzz=FuzzIs7bit -fuzztime=1m ./moji
func FuzzIs7bit(f *testing.F) {
	asciichars := []string{"\x00", "nekochan", "\x7F", "\r\n", "nyaan\x1F", "ネコ", "🐈"}
	for _, e := range asciichars { f.Add(e) }
	f.Fuzz(func(t *testing.T, text string) { _ = Is7bit(text, false) })
}

