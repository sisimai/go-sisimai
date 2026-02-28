// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      ____  ______  
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   / /\ \/ /___ \ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __| / /  \  /  __) |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ / /   /  \ / __/ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__/_/   /_/\_\_____|
//                                                                
import "testing"

func TestLhostX2(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"02",   1, "5.9.210", "",    "filtered",        false, 0, ""},
		 {"02",   2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"02",   3, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"03",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"04",   1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"05",   1, "4.1.9",   "",    "expired",         false, 0, ""}},
		{{"06",   1, "4.4.1",   "",    "networkerror",    false, 0, ""}},
		{{"07",   1, "5.4.14",  "554", "networkerror",    false, 0, ""}},
	}; EngineTest(t, "X2", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.7.1",   "554", "norelaying",      false, 1, ""}},
		{{"1002", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1003", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1004", 1, "5.9.210", "",    "filtered",        false, 0, ""},
		 {"1004", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1004", 3, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1005", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1006", 1, "5.1.2",   "",    "hostunknown",      true, 1, ""}},
		{{"1007", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1008", 1, "4.4.1",   "",    "networkerror",    false, 0, ""}},
		{{"1009", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1010", 1, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1011", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""},
		 {"1011", 2, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1012", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1012", 2, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1013", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1013", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1013", 3, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1013", 4, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1014", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1014", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1014", 3, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1014", 4, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1015", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1015", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1015", 3, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1015", 4, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1015", 5, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1015", 6, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1016", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1016", 2, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1017", 1, "5.9.210", "",    "filtered",        false, 0, ""},
		 {"1017", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1017", 3, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1018", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1018", 2, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1019", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1020", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1021", 1, "5.9.210", "",    "filtered",        false, 0, ""},
		 {"1021", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1021", 3, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1022", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1023", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1023", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1023", 3, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1023", 4, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1023", 5, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1023", 6, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1024", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1024", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1024", 3, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1024", 4, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1025", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1025", 2, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1025", 3, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1025", 4, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1026", 1, "5.9.221", "",    "suspend",         false, 1, ""},
		 {"1026", 2, "5.9.221", "",    "suspend",         false, 1, ""}},
		{{"1027", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""},
		 {"1027", 2, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1028", 1, "4.4.1",   "",    "networkerror",    false, 0, ""}},
		{{"1029", 1, "4.1.9",   "",    "expired",         false, 0, ""}},
		{{"1030", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1031", 1, "5.4.14",  "554", "networkerror",    false, 0, ""}},
		{{"1032", 1, "5.4.14",  "554", "networkerror",    false, 0, ""}},
		{{"1033", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "X2", secretlist, false)
}

