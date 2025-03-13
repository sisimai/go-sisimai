// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package moji

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   ___ (_|_)
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \| | |
//   | |  __/\__ \ |_ / /| | | | | | (_) | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\___// |_|
//                                     |__/   
import "testing"
import "strings"

func TestToLF(t *testing.T) {
	fn := "moji.ToLF"
	cx := 0
	cw := []string{
		"nekochan\r\ncat\r\nkijitora",
		"nekochan\rcats\rkijitora\r\r",
	}
	for _, e := range cw {
		ToLF(&e);
		cx++; if strings.Contains(e, "\r\n") == true { t.Errorf("%s(%s) contains CRLF", fn, e) }
		cx++; if strings.Contains(e, "\r")   == true { t.Errorf("%s(%s) contains CR",   fn, e) }
	}
	ce := "";
	cx++; if ToLF(&ce); ce != "" { t.Errorf("%s() returns %s", fn, ce) }

	t.Logf("The number of tests = %d", cx)
}

