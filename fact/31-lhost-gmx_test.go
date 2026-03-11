// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _         ____ __  ____  __
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      / ___|  \/  \ \/ /
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |  _| |\/| |\  / 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | |  | |/  \ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \____|_|  |_/_/\_\
import "testing"

func TestLhostGMX(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.2.2",   "",    "mailboxfull",     false, 1, ""}},
		{{"02",   1, "5.1.1",   "",    "userunknown",      true, 1, ""}},
		{{"03",   1, "5.2.1",   "",    "userunknown",      true, 1, ""},
		 {"03",   2, "5.2.2",   "",    "mailboxfull",     false, 1, ""}},
		{{"04",   1, "5.9.340", "",    "expired",         false, 0, ""}},
	}; EngineTest(t, "GMX", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1002", 1, "5.1.1",   "",    "userunknown",      true, 1, ""}},
		{{"1003", 1, "5.2.2",   "",    "mailboxfull",     false, 1, ""}},
		{{"1004", 1, "5.2.1",   "",    "userunknown",      true, 1, ""},
		 {"1004", 2, "5.2.2",   "",    "mailboxfull",     false, 1, ""}},
	}; EngineTest(t, "GMX", secretlist, false)
}

