// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        __  __         _     __           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  \/  | ___   / \   / _| ___  ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |\/| |/ __| / _ \ | |_ / _ \/ _ \
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |  | | (__ / ___ \|  _|  __/  __/
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  |_|\___/_/   \_\_|  \___|\___|
//                                                                                      
import "testing"

func TestLhostMcAfee(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"04",   1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"05",   1, "5.9.210", "550", "filtered",        false, 1, ""}},
	}; EngineTest(t, "McAfee", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1009", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
	}; EngineTest(t, "McAfee", secretlist, false)
}

