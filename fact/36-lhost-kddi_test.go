// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        _  ______  ____ ___ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     | |/ /  _ \|  _ \_ _|
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| ' /| | | | | | | | 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| . \| |_| | |_| | | 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|\_\____/|____/___|
import "testing"

func TestLhostKDDI(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"02",   1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"03",   1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
	}; EngineTest(t, "KDDI", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1002", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1003", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
	}; EngineTest(t, "KDDI", secretlist, false)
}

