// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _         ____                   _       ____                           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      / ___| ___   ___   __ _| | ___ / ___|_ __ ___  _   _ _ __  ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |  _ / _ \ / _ \ / _` | |/ _ \ |  _| '__/ _ \| | | | '_ \/ __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | (_) | (_) | (_| | |  __/ |_| | | | (_) | |_| | |_) \__ \
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \____|\___/ \___/ \__, |_|\___|\____|_|  \___/ \__,_| .__/|___/
//                                                                      |___/                             |_|        
import "testing"

func TestLhostGoogleGroups(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"02",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"03",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"04",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"05",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"06",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"07",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"08",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"09",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"10",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"11",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"12",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"13",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"14",   1, "5.0.918", "",    "rejected",        false, ""}},
	}; EngineTest(t, "GoogleGroups", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1002", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1003", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1004", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1005", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1006", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1007", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1008", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1009", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1010", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1011", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1012", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1013", 1, "5.0.918", "",    "rejected",        false, ""}},
		{{"1014", 1, "5.0.918", "",    "rejected",        false, ""}},
	}; EngineTest(t, "GoogleGroups", secretlist, false)
}

