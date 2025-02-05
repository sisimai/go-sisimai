// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
import "testing"

func TestHEADERTABLE(t *testing.T) {
	fn := "sisimai/rfc5322.HEADERTABLE"
	cx := 0
	cv := HEADERTABLE()

	cx++; if len(cv) == 0 { t.Errorf("%s() returns empty", fn) }
	cx++; if len(cv) != 6 { t.Errorf("%s() returns %d elements", fn, len(cv)) }
	for e := range cv {
		cx++; if len(cv[e]) == 0 { t.Errorf("%s[%s] have no element", fn, e) }
		for _, f := range cv[e] {
			cx++; if f == "" { t.Errorf("%s[%s] is empty string", fn, e) }
		}
	}

	t.Logf("The number of tests = %d", cx)
}

func TestHEADERFIELDS(t *testing.T) {
	fn := "sisimai/rfc5322.HEADERFIELDS"
	cx := 0

	for _, e := range []string{"messageid", "subject", "listid", "date", "addresser", "recipient"} {
		cv := HEADERFIELDS(e)
		cx++; if len(cv) == 0 { t.Errorf("%s(%s) returns empty list", fn, e) }
	}
	cx++; if cv := HEADERFIELDS("");     len(cv) > 0 { t.Errorf("%s() returns %v", fn, cv) }
	cx++; if cv := HEADERFIELDS("neko"); len(cv) > 0 { t.Errorf("%s(neko) returns %v", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

