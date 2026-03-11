// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      __  _______ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    \ \/ /___ / 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____\  /  |_ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____/  \ ___) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   /_/\_\____/ 
import "testing"

func TestLhostX3(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.3.0",   "553", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"03",   1, "5.3.0",   "553", "userunknown",      true, 1, ""}},
		{{"05",   1, "5.9.300", "",    "undefined",       false, 0, ""}},
		{{"06",   1, "5.2.2",   "552", "mailboxfull",     false, 1, ""}},
	}; EngineTest(t, "X3", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.3.0",   "553", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.9.300", "",    "undefined",       false, 0, ""}},
		{{"1003", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1004", 1, "5.3.0",   "553", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.9.300", "",    "undefined",       false, 0, ""}},
		{{"1006", 1, "5.3.0",   "553", "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1008", 1, "5.3.0",   "553", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "X3", secretlist, false)
}

