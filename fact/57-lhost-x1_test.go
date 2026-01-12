// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"02",   1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"03",   1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"04",   1, "5.9.340", "",    "expired",         false, false, ""}},
	}; EngineTest(t, "X1", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"1002", 1, "5.9.210", "",    "filtered",        false, false, ""},
		 {"1002", 2, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"1003", 1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"1004", 1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"1005", 1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"1006", 1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"1007", 1, "5.9.340", "",    "expired",         false, false, ""}},
		{{"1008", 1, "5.9.221", "",    "suspend",         false,  true, ""}},
	}; EngineTest(t, "X1", secretlist, false)
}

