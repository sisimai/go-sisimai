// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc1894

//  _____         _      ______  _____ ____ _  ___  ___  _  _   
// |_   _|__  ___| |_   / /  _ \|  ___/ ___/ |( _ )/ _ \| || |  
//   | |/ _ \/ __| __| / /| |_) | |_ | |   | |/ _ \ (_) | || |_ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___| | (_) \__, |__   _|
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_|\___/  /_/   |_|  
import "testing"

func TestFIELDINDEX(t *testing.T) {
	fn := "sisimai/rfc1894.FIELDINDEX"
	cx := 0
	cv := FIELDINDEX()

	cx++; if len(cv) ==  0 { t.Errorf("%s() returns empty", fn) }
	cx++; if len(cv) != 12 { t.Errorf("%s() returns %d elements", fn, len(cv)) }
	for _, e := range cv {
		cx++; if e == "" { t.Errorf("%s() includes an empty string", fn) }
	}

	t.Logf("The number of tests = %d", cx)
}

func TestFIELDTABLE(t *testing.T) {
	fn := "sisimai/rfc1894.FIELDTABLE"
	cx := 0
	ae := FIELDTABLE()

	for e := range ae {
		cx++; if e     == "" { t.Errorf("%s() have an empty key", fn) }
		cx++; if ae[e] == "" { t.Errorf("%s() Key:%s have an empty value", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

