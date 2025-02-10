// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"03",   1, "5.2.2",   "550", "mailboxfull",     false, ""},
		 {"03",   2, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"04",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"05",   1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"06",   1, "5.4.4",   "",    "notaccept",        true, ""}},
	}; EngineTest(t, "Aol", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "Aol", secretlist, false)
}

