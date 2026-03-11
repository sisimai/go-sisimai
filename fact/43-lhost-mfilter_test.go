// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _                  _____ ___ _   _____ _____ ____  
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      _ __ ___ |  ___|_ _| | |_   _| ____|  _ \ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| '_ ` _ \| |_   | || |   | | |  _| | |_) |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| | | | | |  _|  | || |___| | | |___|  _ < 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_| |_| |_|_|   |___|_____|_| |_____|_| \_\
import "testing"

func TestLhostmFILTER(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"04",   1, "5.4.1",   "550", "rejected",        false, 0, ""}},
		{{"05",   1, "4.3.1",   "452", "systemfull",      false, 0, ""}},
	}; EngineTest(t, "mFILTER", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1003", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1004", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1007", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.4.1",   "550", "rejected",        false, 0, ""}},
		{{"1009", 1, "5.4.1",   "550", "rejected",        false, 0, ""}},
		{{"1010", 1, "4.3.1",   "452", "systemfull",      false, 0, ""}},
		{{"1011", 1, "5.6.0",   "550", "spamdetected",    false, 0, ""}},
		{{"1012", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1013", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1014", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "mFILTER", secretlist, false)
}

