// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"02",   1, "5.0.910", "",    "filtered",        false, false, ""},
		 {"02",   2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"02",   3, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"03",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"04",   1, "5.0.922", "",    "mailboxfull",     false, false, ""}},
		{{"05",   1, "4.1.9",   "",    "expired",         false, false, ""}},
		{{"06",   1, "4.4.1",   "",    "expired",         false, false, ""}},
		{{"07",   1, "5.4.14",  "554", "networkerror",    false, false, ""}},
	}; EngineTest(t, "X2", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.7.1",   "554", "norelaying",      false,  true, ""}},
		{{"1002", 1, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"1003", 1, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"1004", 1, "5.0.910", "",    "filtered",        false, false, ""},
		 {"1004", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1004", 3, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"1005", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1006", 1, "5.1.2",   "",    "hostunknown",      true,  true, ""}},
		{{"1007", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1008", 1, "4.4.1",   "",    "expired",         false, false, ""}},
		{{"1009", 1, "5.0.922", "",    "mailboxfull",     false, false, ""}},
		{{"1010", 1, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1011", 1, "5.0.922", "",    "mailboxfull",     false, false, ""},
		 {"1011", 2, "5.0.922", "",    "mailboxfull",     false, false, ""}},
		{{"1012", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1012", 2, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1013", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1013", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1013", 3, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1013", 4, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1014", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1014", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1014", 3, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1014", 4, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1015", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1015", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1015", 3, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1015", 4, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1015", 5, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1015", 6, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1016", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1016", 2, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1017", 1, "5.0.910", "",    "filtered",        false, false, ""},
		 {"1017", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1017", 3, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"1018", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1018", 2, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1019", 1, "5.0.922", "",    "mailboxfull",     false, false, ""}},
		{{"1020", 1, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"1021", 1, "5.0.910", "",    "filtered",        false, false, ""},
		 {"1021", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1021", 3, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"1022", 1, "5.0.910", "",    "filtered",        false, false, ""}},
		{{"1023", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1023", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1023", 3, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1023", 4, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1023", 5, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1023", 6, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1024", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1024", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1024", 3, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1024", 4, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1025", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1025", 2, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1025", 3, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1025", 4, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1026", 1, "5.0.921", "",    "suspend",         false,  true, ""},
		 {"1026", 2, "5.0.921", "",    "suspend",         false,  true, ""}},
		{{"1027", 1, "5.0.922", "",    "mailboxfull",     false, false, ""},
		 {"1027", 2, "5.0.922", "",    "mailboxfull",     false, false, ""}},
		{{"1028", 1, "4.4.1",   "",    "expired",         false, false, ""}},
		{{"1029", 1, "4.1.9",   "",    "expired",         false, false, ""}},
		{{"1030", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1031", 1, "5.4.14",  "554", "networkerror",    false, false, ""}},
		{{"1032", 1, "5.4.14",  "554", "networkerror",    false, false, ""}},
		{{"1033", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
	}; EngineTest(t, "X2", secretlist, false)
}

