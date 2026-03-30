// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sisimai

//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
import "testing"

func BenchmarkRise(b *testing.B) {
	path := "./tmp/all-the-emails"
	args := Args()
	args.Delivered = true
	args.Vacation  = true

	for j := 0; j < b.N; j++ {
		Rise(path, args)
	}
}

func BenchmarkDump(b *testing.B) {
	path := "./tmp/all-the-emails"
	args := Args()
	args.Delivered = true
	args.Vacation  = true

	for j := 0; j < b.N; j++ {
		Dump(path, args)
	}
}

