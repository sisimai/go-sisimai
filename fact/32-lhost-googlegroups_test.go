// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"02",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"03",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"04",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"05",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"06",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"07",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"08",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"09",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"10",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"11",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"12",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"13",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"14",   1, "5.9.110", "",    "rejected",        false, 0, ""}},
	}; EngineTest(t, "GoogleGroups", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1002", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1003", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1004", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1005", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1006", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1007", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1008", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1009", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1010", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1011", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1012", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1013", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1014", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
		{{"1015", 1, "5.9.110", "",    "rejected",        false, 0, ""}},
	}; EngineTest(t, "GoogleGroups", secretlist, false)
}

