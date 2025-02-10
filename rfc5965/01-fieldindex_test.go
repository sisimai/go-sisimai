// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5965

//  _____         _      ______  _____ ____ ____  ___   __  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|/ _ \ / /_| ___| 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ (_) | '_ \___ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) \__, | (_) |__) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/  /_/ \___/____/ 
import "testing"

func TestFIELDINDEX(t *testing.T) {
	fn := "sisimai/rfc5965.FIELDINDEX"
	cx := 0
	cv := FIELDINDEX()

	cx++; if len(cv) ==  0 { t.Errorf("%s() returns empty", fn) }
	cx++; if len(cv) != 13 { t.Errorf("%s() returns empty", fn) }
	for _, e := range cv {
		cx++; if e == "" { t.Errorf("%s() includes an empty string", fn) }
	}

	t.Logf("The number of tests = %d", cx)
}

