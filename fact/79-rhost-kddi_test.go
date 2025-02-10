// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _        _  ______  ____ ___ 
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_     | |/ /  _ \|  _ \_ _|
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| ' /| | | | | | | | 
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| . \| |_| | |_| | | 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|    |_|\_\____/|____/___|
import "testing"

func TestRhostKDDI(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.2.0",   "550", "filtered",        false, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "KDDI", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "KDDI", secretlist, false)
}


