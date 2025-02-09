// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
import "testing"

func TestFIELDINDEX(t *testing.T) {
	fn := "sisimai/rfc5322.FIELDINDEX"
	cx := 0
	cv := FIELDINDEX()

	cx++; if len(cv) ==  0 { t.Errorf("%s() returns empty", fn) }
	cx++; if len(cv) != 15 { t.Errorf("%s() returns empty", fn) }
	for _, e := range cv {
		cx++; if e == "" { t.Errorf("%s() includes an empty string", fn) }
	}

	t.Logf("The number of tests = %d", cx)
}

