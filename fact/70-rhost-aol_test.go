// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _          _         _ 
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_       / \   ___ | |
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____ / _ \ / _ \| |
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____/ ___ \ (_) | |
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|   /_/   \_\___/|_|
import "testing"

func TestRhostAol(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.4.4",   "",    "hostunknown",      true, 1, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"03",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""},
		 {"03",   2, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"04",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"05",   1, "5.4.4",   "",    "hostunknown",      true, 1, ""}},
		{{"06",   1, "5.4.4",   "",    "notaccept",        true, 1, ""}},
	}; EngineTest(t, "Aol", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "Aol", secretlist, false)
}

