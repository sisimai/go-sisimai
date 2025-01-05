// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        __  ____  ___                _      
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  \/  \ \/ / |    ___   __ _(_) ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |\/| |\  /| |   / _ \ / _` | |/ __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |  | |/  \| |__| (_) | (_| | | (__ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  |_/_/\_\_____\___/ \__, |_|\___|
//                                                                           |___/        
import "testing"

func TestLhostMXLogic(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "5.0.910", "550", "filtered",        false, ""}},
	}; EngineTest(t, "MXLogic", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1007", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1008", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1010", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1011", 1, "5.0.910", "550", "filtered",        false, ""}},
	}; EngineTest(t, "MXLogic", secretlist, false)
}

