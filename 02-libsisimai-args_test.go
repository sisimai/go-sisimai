// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sisimai

//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
import "testing"

func TestArgs(t *testing.T) {
	fn := "sis.DecodingArgs"
	cx := 0
	cv := Args()

	cx++; if cv.Delivered == true { t.Errorf("%s.Delivered is true", fn) }
	cx++; if cv.Vacation  == true { t.Errorf("%s.Vacation  is true", fn) }
	cx++; if cv.Callback0 != nil  { t.Errorf("%s.Callback0 is not nil: %v", fn, cv.Callback0) }
	cx++; if cv.Callback1 != nil  { t.Errorf("%s.Callback1 is not nil: %v", fn, cv.Callback1) }

	t.Logf("The number of tests = %d", cx)
}

