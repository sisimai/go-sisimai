// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _     __   __    _                 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   \ \ / /_ _| |__   ___   ___  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|___\ V / _` | '_ \ / _ \ / _ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| | (_| | | | | (_) | (_) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|\__,_|_| |_|\___/ \___/ 
//                                                                              
import "testing"

func TestLhostYahoo(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"03",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"04",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"05",   1, "5.2.1",   "550", "userunknown",      true, 1, ""}},
		{{"06",   1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"07",   1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"08",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"09",   1, "5.9.215", "",    "notaccept",        true, 1, ""}},
		{{"10",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"11",   1, "5.1.8",   "501", "rejected",        false, 0, ""}},
		{{"12",   1, "5.1.8",   "501", "rejected",        false, 0, ""}},
		{{"13",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"14",   1, "5.9.134", "554", "blocked",         false, 0, ""}},
	}; EngineTest(t, "Yahoo", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1003", 1, "5.2.1",   "550", "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.9.134", "554", "blocked",         false, 0, ""}},
		{{"1006", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1008", 1, "5.9.215", "",    "notaccept",        true, 1, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1010", 1, "5.1.8",   "501", "rejected",        false, 0, ""}},
		{{"1011", 1, "5.9.134", "554", "blocked",         false, 0, ""}},
	}; EngineTest(t, "Yahoo", secretlist, false)
}

