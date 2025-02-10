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

func TestToUTF8(t *testing.T) {
	fn := "sisimai/string.ToUTF8"
	cx := 0
	cw := ""

	if cv, ce := ToUTF8([]byte(cw), "iso-2022-jp"); fn != "" {
		cx++; if ce != nil { t.Errorf("%s(%s) returns error: %s", fn, cw, ce) }
		cx++; if cv != ""  { t.Errorf("%s(%s) returns %s", fn, cw, cv) }
	}

	if cv, ce := ToUTF8([]byte("ネコ"), "utf-8"); fn != "" {
		cx++; if cv != "ネコ" { t.Errorf("%s(%s, utf-8) returns %s", fn, "ネコ", cv) }
		cx++; if ce != nil    { t.Errorf("%s(%s, utf-8) returns error: %s", fn, "ネコ", ce) }
	}

	if cv, ce := ToUTF8([]byte("Neko"), "us-ascii"); fn != "" {
		cx++; if cv != "Neko" { t.Errorf("%s(%s, us-ascii) returns %s", fn, "Neko", cv) }
		cx++; if ce != nil    { t.Errorf("%s(%s, us-ascii) returns error: %s", fn, "Neko", ce) }
	}

	t.Logf("The number of tests = %d", cx)
}

