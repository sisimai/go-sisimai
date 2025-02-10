// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string

//  _____         _      __   _        _             
// |_   _|__  ___| |_   / /__| |_ _ __(_)_ __   __ _ 
//   | |/ _ \/ __| __| / / __| __| '__| | '_ \ / _` |
//   | |  __/\__ \ |_ / /\__ \ |_| |  | | | | | (_| |
//   |_|\___||___/\__/_/ |___/\__|_|  |_|_| |_|\__, |
//                                             |___/ 
import "testing"
import "strings"

func TestToLF(t *testing.T) {
	fn := "sisimai/string.ToLF"
	cx := 0
	cw := []string{
		"nekochan\r\ncat\r\nkijitora",
		"nekochan\rcats\rkijitora\r\r",
	}
	for _, e := range cw {
		cv := ToLF(&e);
		cx++; if strings.Contains(*cv, "\r\n") == true { t.Errorf("%s(%s) contains CRLF", fn, *cv) }
		cx++; if strings.Contains(*cv, "\r")   == true { t.Errorf("%s(%s) contains CR",   fn, *cv) }
	}
	ce := "";
	cx++; if cv := ToLF(&ce); *cv != "" { t.Errorf("%s() returns %s", fn, *cv) }

	t.Logf("The number of tests = %d", cx)
}

