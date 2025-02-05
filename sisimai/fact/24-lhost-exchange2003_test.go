// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"02",   1, "5.0.911", "",    "userunknown",      true, ""},
		 {"02",   2, "5.0.911", "",    "userunknown",      true, ""}},
		{{"03",   1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"04",   1, "5.0.910", "",    "filtered",        false, ""}},
		{{"05",   1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"07",   1, "5.0.911", "",    "userunknown",      true, ""}},
	}; EngineTest(t, "Exchange2003", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1002", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1003", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1004", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1005", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1006", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1007", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1008", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1009", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1010", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1011", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1012", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1013", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1014", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1015", 1, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015", 2, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015", 3, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015", 4, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015", 5, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015", 6, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015", 7, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015", 8, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015", 9, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015",10, "5.0.911", "",    "userunknown",      true, ""},
		 {"1015",11, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1016", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1017", 1, "5.0.910", "",    "filtered",        false, ""}},
		{{"1018", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1019", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1020", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1021", 1, "5.0.911", "",    "userunknown",      true, ""},
		 {"1021", 2, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1022", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1023", 1, "5.0.910", "",    "filtered",        false, ""}},
		{{"1024", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1025", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1026", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1027", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1028", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1029", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1030", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1031", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1032", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1033", 1, "5.0.911", "",    "userunknown",      true, ""}},
	}; EngineTest(t, "Exchange2003", secretlist, false)
}

