// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      ____  ___ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   / /\ \/ / |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __| / /  \  /| |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ / /   /  \| |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__/_/   /_/\_\_|
//                                                            
import "testing"

func TestLhostX1(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"02",   1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"03",   1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"04",   1, "5.9.340", "",    "expired",         false, 0, ""}},
	}; EngineTest(t, "X1", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1002", 1, "5.9.210", "",    "filtered",        false, 0, ""},
		 {"1002", 2, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1003", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1004", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1005", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1006", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1007", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1008", 1, "5.9.221", "",    "suspend",         false, 1, ""}},
	}; EngineTest(t, "X1", secretlist, false)
}

