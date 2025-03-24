// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _____         _      ___ _               _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|
import "testing"

func TestIndicators(t *testing.T) {
	fn := "lhost.Indicators"
	cv := Indicators
	cx := 0

	cx++; if cv == nil                 { t.Fatalf("%s = nil", fn) }
	cx++; if cv["deliverystatus"] != 2 { t.Errorf("%s[deliverystatus] = %d", fn, cv["deliverystatus"]) }
	cx++; if cv["message-rfc822"] != 4 { t.Errorf("%s[message-rfc822] = %d", fn, cv["message-rfc822"]) }
	t.Logf("The number of tests = %d", cx)
}

