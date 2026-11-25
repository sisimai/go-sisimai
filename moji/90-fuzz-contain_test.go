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

// % go test -tags=fuzz -fuzz=FuzzContainsOnlyNumbers -fuzztime=1m ./moji
func FuzzContainsOnlyNumbers(f *testing.F) {
	numberlist := []string{"1", "23", "456", "78910", "A", "B1", "C12", "34D5E"}
	for _, e := range numberlist { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, text string) { _ = ContainsOnlyNumbers(text) })
}

// % go test -tags=fuzz -fuzz=FuzzIsContained -fuzztime=1m ./moji
func FuzzIsContained(f *testing.F) {
	neko := []string{"neko", "22", "cat", "nyaan"}
	list := []string{"nekochan", "kijitora"}
	for _, e := range neko { f.Add(e) }
	f.Fuzz(func(t *testing.T, text string) { _ = IsContained(text, list) })
}

// % go test -tags=fuzz -fuzz=FuzzContainsAny -fuzztime=1m ./moji
func FuzzContainsAny(f *testing.F) {
	neko := []string{"neko", "22", "cat", "nyaan"}
	list := []string{"nekochan", "kijitora"}
	for _, e := range neko { f.Add(e) }
	f.Fuzz(func(t *testing.T, text string) { _ = ContainsAny(text, list) })
}

// % go test -tags=fuzz -fuzz=FuzzHasPrefixAny -fuzztime=1m ./moji
func FuzzHasPrefixAny(f *testing.F) {
	neko := []string{"neko", "22", "cat", "nyaan"}
	list := []string{"nekochan", "kijitora"}
	for _, e := range neko { f.Add(e) }
	f.Fuzz(func(t *testing.T, text string) { _ = HasPrefixAny(text, list) })
}

