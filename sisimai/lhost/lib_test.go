// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _____         _      ___ _               _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|
import "testing"

func TestINDICATORS(t *testing.T) {
	fn := "sisimai/lhost.INDICATORS"
	cv := INDICATORS()
	cx := 0

	cx++; if cv == nil                 { t.Fatalf("%s() = nil", fn) }
	cx++; if cv["deliverystatus"] != 2 { t.Errorf("%s(deliverystatus) = %d", fn, cv["deliverystatus"]) }
	cx++; if cv["message-rfc822"] != 4 { t.Errorf("%s(message-rfc822) = %d", fn, cv["message-rfc822"]) }
	t.Logf("The number of tests = %d", cx)
}

func TestINDEX(t *testing.T) {
	fn := "sisimai/lhost.INDEX"
	cv := INDEX()
	cx := 0

	cx++; if len(cv) == 0 { t.Fatalf("%s() = empty", fn) }
	for _, e := range cv {
		cx++; if InquireFor[e] == nil { t.Errorf("%s(%s) = nil", fn, e) }
	}
	t.Logf("The number of tests = %d", cx)
}

