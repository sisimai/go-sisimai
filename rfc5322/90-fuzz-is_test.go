//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
package rfc5322
import "testing"

// % go test -tags=fuzz -fuzz=FuzzIsEmailAddress -fuzztime=1m ./rfc5322
func FuzzIsEmailAddress(f *testing.F) {
	for _, e := range TestEmailAddrs { f.Add(e.argument + "\n") }
	f.Fuzz(func(t *testing.T, addr string) { _ = IsEmailAddress(addr) })
}

// % go test -tags=fuzz -fuzz=FuzzIsQuotedAddress -fuzztime=1m ./rfc5322
func FuzzIsQuotedAddress(f *testing.F) {
	for _, e := range TestEmailAddrs { f.Add(e.argument + "\n") }
	f.Fuzz(func(t *testing.T, addr string) { _ = IsQuotedAddress(addr) })
}

// % go test -tags=fuzz -fuzz=FuzzIsComment -fuzztime=1m ./rfc5322
func FuzzIsComment(f *testing.F) {
	for _, e := range TestEmailAddrs { f.Add(e.comments + "\n") }
	f.Fuzz(func(t *testing.T, text string) { _ = IsComment(text) })
}
