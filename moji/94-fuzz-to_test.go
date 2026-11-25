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

// % go test -tags=fuzz -fuzz=FuzzToLF -fuzztime=1m ./moji
func FuzzToLF(f *testing.F) {
	stringlist := []string{
		"nekochan\r\n",
		"nyaan?\r\r\r",
		"kijitora??\n",
	}
	for _, e := range stringlist { f.Add(e) }
	f.Fuzz(func(t *testing.T, text string) { _ = ToLF(&text) })
}

// % go test -tags=fuzz -fuzz=FuzzToPlain -fuzztime=1m ./moji
func FuzzToPlain(f *testing.F) {
	hypertexts := []string{
		"<html><head><title>nekochan</title></head><body>nyaan?</body></html>",
		"<body style = 'font-size: 22px;'>nyaan?</body>",
	}
	for _, e := range hypertexts { f.Add(e) }
	f.Fuzz(func(t *testing.T, html string) { _ = ToPlain(&html) })
}

// % go test -tags=fuzz -fuzz=FuzzZeroPadding -fuzztime=1m ./moji
func FuzzZeroPadding(f *testing.F) {
	numberlist := []int{1,22,330,4000}
	for _, e := range numberlist { f.Add(e) }
	f.Fuzz(func(t *testing.T, args int) { _ = ZeroPadding(args, 22) })
}

