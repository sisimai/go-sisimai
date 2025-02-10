// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
import "testing"

func TestLONGFIELDS(t *testing.T) {
	fn := "sisimai/rfc5322.LONGFIELDS"
	cx := 0
	cv := LONGFIELDS()

	cx++; if len(cv) == 0 { t.Errorf("%s() returns empty", fn) }
	cx++; if len(cv) != 4 { t.Errorf("%s() returns %d elements", fn, len(cv)) }
	for e := range cv {
		cx++; if cv[e] == false { t.Errorf("%s[%s] is false", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

