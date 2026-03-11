// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _         ____                 _           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      / ___|___  _   _ _ __(_) ___ _ __ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |   / _ \| | | | '__| |/ _ \ '__|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |__| (_) | |_| | |  | |  __/ |   
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \____\___/ \__,_|_|  |_|\___|_|   
import "testing"

func TestLhostCourier(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.0.0",   "550", "filtered",        false, 1, ""}},
		{{"03",   1, "5.7.1",   "550", "rejected",        false, 0, ""}},
		{{"04",   1, "5.0.0",   "",    "hostunknown",      true, 1, ""}},
	}; EngineTest(t, "Courier", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.0.0",   "550", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.0.0",   "550", "userunknown",      true, 1, ""}},
		{{"1003", 1, "5.7.1",   "550", "rejected",        false, 0, ""}},
		{{"1004", 1, "5.0.0",   "550", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.0.0",   "550", "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1009", 1, "5.0.0",   "550", "filtered",        false, 1, ""}},
		{{"1010", 1, "5.7.1",   "550", "rejected",        false, 0, ""}},
		{{"1011", 1, "5.0.0",   "",    "hostunknown",      true, 1, ""}},
		{{"1012", 1, "5.0.0",   "",    "hostunknown",      true, 1, ""}},
	}; EngineTest(t, "Courier", secretlist, false)
}

