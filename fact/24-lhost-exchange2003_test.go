// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        _____          _                            ____   ___   ___ _____ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     | ____|_  _____| |__   __ _ _ __   __ _  ___|___ \ / _ \ / _ \___ / 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____|  _| \ \/ / __| '_ \ / _` | '_ \ / _` |/ _ \ __) | | | | | | ||_ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |___ >  < (__| | | | (_| | | | | (_| |  __// __/| |_| | |_| |__) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_____/_/\_\___|_| |_|\__,_|_| |_|\__, |\___|_____|\___/ \___/____/ 
//                                                                                     |___/                             
import "testing"

func TestLhostExchange2003(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"02",   1, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"02",   2, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"03",   1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"04",   1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"05",   1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"07",   1, "5.9.213", "",    "userunknown",      true, 1, ""}},
	}; EngineTest(t, "Exchange2003", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1003", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1009", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1010", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1011", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1012", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1013", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1014", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1015", 1, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015", 2, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015", 3, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015", 4, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015", 5, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015", 6, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015", 7, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015", 8, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015", 9, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015",10, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1015",11, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1016", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1017", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1018", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1019", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1020", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1021", 1, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1021", 2, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1022", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1023", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1024", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1025", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1026", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1027", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1028", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1029", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1030", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1031", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1032", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1033", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
	}; EngineTest(t, "Exchange2003", secretlist, false)
}

