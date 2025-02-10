// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.3.0",   "553", "userunknown",      true, ""}},
		{{"02",   1, "5.0.947", "",    "expired",         false, ""}},
		{{"03",   1, "5.3.0",   "553", "userunknown",      true, ""}},
		{{"05",   1, "5.0.900", "",    "undefined",       false, ""}},
		{{"06",   1, "5.2.2",   "552", "mailboxfull",     false, ""}},
	}; EngineTest(t, "X3", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.3.0",   "553", "userunknown",      true, ""}},
		{{"1002", 1, "5.0.900", "",    "undefined",       false, ""}},
		{{"1003", 1, "5.0.947", "",    "expired",         false, ""}},
		{{"1004", 1, "5.3.0",   "553", "userunknown",      true, ""}},
		{{"1005", 1, "5.0.900", "",    "undefined",       false, ""}},
		{{"1006", 1, "5.3.0",   "553", "userunknown",      true, ""}},
		{{"1007", 1, "5.0.947", "",    "expired",         false, ""}},
		{{"1008", 1, "5.3.0",   "553", "userunknown",      true, ""}},
	}; EngineTest(t, "X3", secretlist, false)
}

