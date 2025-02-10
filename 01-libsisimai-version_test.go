// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sisimai

//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
import "testing"
import "strings"

func TestVersion(t *testing.T) {
	fn := "sisimai.Version"
	cx := 0
	cv := Version()

	cx++; if cv == ""                              { t.Errorf("%s() returns empty", fn)                 }
	cx++; if strings.HasPrefix(cv, "v5.") == false { t.Errorf("%s() returns invalid value :%s", fn, cv) }
	t.Logf("The number of tests = %d", cx)
}

