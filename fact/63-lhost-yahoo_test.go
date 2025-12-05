// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"03",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"04",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"05",   1, "5.2.1",   "550", "userunknown",      true,  true, ""}},
		{{"06",   1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"07",   1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"08",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"09",   1, "5.0.932", "",    "notaccept",        true,  true, ""}},
		{{"10",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"11",   1, "5.1.8",   "501", "rejected",        false, false, ""}},
		{{"12",   1, "5.1.8",   "501", "rejected",        false, false, ""}},
		{{"13",   1, "5.0.930", "",    "systemerror",     false, false, ""}},
		{{"14",   1, "5.0.971", "554", "blocked",         false, false, ""}},
	}; EngineTest(t, "Yahoo", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1002", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1003", 1, "5.2.1",   "550", "userunknown",      true,  true, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1005", 1, "5.0.971", "554", "blocked",         false, false, ""}},
		{{"1006", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1007", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1008", 1, "5.0.932", "",    "notaccept",        true,  true, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1010", 1, "5.1.8",   "501", "rejected",        false, false, ""}},
		{{"1011", 1, "5.0.971", "554", "blocked",         false, false, ""}},
	}; EngineTest(t, "Yahoo", secretlist, false)
}

