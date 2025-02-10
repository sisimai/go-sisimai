// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc1123

//  _____         _      ______  _____ ____ _ _ ____  _____ 
// |_   _|__  ___| |_   / /  _ \|  ___/ ___/ / |___ \|___ / 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   | | | __) | |_ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___| | |/ __/ ___) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_|_|_____|____/ 
import "testing"

func TestIsInternetHost(t *testing.T) {
	fn := "sisimai/rfc1123.IsInternetHost"
	cx := 0

	hostnames0 := []string{
		"",
		"127.0.0.1",
		"cat",
		"neko",
		"nyaan.22",
		"mx0.example.22",
		"mx0.example.jp-",
		"mx--0.example.jp",
		"mx..0.example.jp",
		"mx0.example.jp/neko",
		"mx22.nyaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaan.jp",
	}
	hostnames1 := []string{
		"localhost",
		"mx1.example.jp",
		"mx1.example.jp.",
		"a.jp",
		"mx22.nyaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaan.jp",
	}

	for _, e := range hostnames0 {
		cx++; if cv := IsInternetHost(e); cv == true  { t.Errorf("%s(%s) returns true",  fn, e) }
	}
	for _, e := range hostnames1 {
		cx++; if cv := IsInternetHost(e); cv == false { t.Errorf("%s(%s) returns false", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

