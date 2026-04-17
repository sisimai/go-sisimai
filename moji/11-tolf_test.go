// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package moji

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   ___ (_|_)
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \| | |
//   | |  __/\__ \ |_ / /| | | | | | (_) | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\___// |_|
//                                     |__/   
import "testing"
import "bytes"

func TestToLF(t *testing.T) {
	fn := "moji.ToLF"
	cx := 0
	cw := [][]byte{
		[]byte("nekochan\r\ncat\r\nkijitora"),
		[]byte("nekochan\rcats\rkijitora\r\r"),
	}
	for _, e := range cw {
		v := ToLF(e);
		cx++; if bytes.Contains(v, []byte("\r\n")) == true { t.Errorf("%s(%s) contains CRLF", fn, v) }
		cx++; if bytes.Contains(v, []byte("\r"))   == true { t.Errorf("%s(%s) contains CR",   fn, v) }
	}
	ce := []byte("");
	cx++; if v := ToLF(ce); len(v) > 0 { t.Errorf("%s() returns %s", fn, v) }

	t.Logf("The number of tests = %d", cx)
}

