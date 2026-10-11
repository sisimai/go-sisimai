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

func TestZeroPadding(t *testing.T) {
	fn := "moji.ZeroPadding"
	ae := []struct {n int; d int; x string}{
		{  0,  0, "0"},
		{  0,  1, "0"},
		{  0,  2, "00"},
		{  1,  0, "1"},
		{  1,  1, "1"},
		{  1,  2, "01"},
		{  2,  2, "02"},
		{  2,  3, "002"},
		{  2,  4, "0002"},
		{  2,  8, "00000002"},
		{ 25,  7, "0000025"},
		{ 25,  1, "25"},
		{ 25, 17, "00000025"},
	}
	cx := 0
	for _, e := range ae {
		cx++; if cv := ZeroPadding(e.n, e.d); cv != e.x { t.Errorf("%s(%d,%d) returns %s", fn, e.n, e.d, cv) }
	}

	t.Logf("The number of tests = %d", cx)
}

