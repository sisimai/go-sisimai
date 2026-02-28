// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _         ___                   ____  __  __ _____ ____  ____  
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      / _ \ _ __   ___ _ __ / ___||  \/  |_   _|  _ \|  _ \ 
//   | |/ _ \/ __| __| / /| | "_ \ / _ \/ __| __|____| | | | "_ \ / _ \ "_ \\___ \| |\/| | | | | |_) | | | |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | |_) |  __/ | | |___) | |  | | | | |  __/| |_| |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \___/| .__/ \___|_| |_|____/|_|  |_| |_| |_|   |____/ 
//                                                         |_|                                              
import "testing"

func TestLhostOpenSMTPD(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""},
		 {"02",   2, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"04",   1, "5.9.341", "",    "networkerror",    false, 0, ""}},
		{{"05",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"06",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"10",   1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"11",   1, "5.7.26",  "550", "authfailure",     false, 0, ""}},
		{{"12",   1, "5.9.215", "",    "notaccept",        true, 1, ""}},
		{{"13",   1, "4.7.0",   "421", "badreputation",   false, 0, ""}},
		{{"14",   1, "5.7.25",  "550", "requireptr",      false, 0, ""}},
		{{"15",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"16",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"17",   1, "5.1.1",   "550", "userunknown",      true, 1, ""},
		 {"17",   2, "5.2.2",   "552", "mailboxfull",     false, 1, ""}},
	}; EngineTest(t, "OpenSMTPD", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.2.1",   "550", "filtered",        false, 1, ""}},
		{{"1003", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1004", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1007", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""},
		 {"1008", 2, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1009", 1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"1010", 1, "5.9.341", "",    "networkerror",    false, 0, ""}},
		{{"1011", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1012", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""},
		 {"1012", 2, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1013", 1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"1014", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1015", 1, "5.9.341", "",    "networkerror",    false, 0, ""}},
		{{"1016", 1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"1017", 1, "5.7.26",  "550", "authfailure",     false, 0, ""}},
		{{"1018", 1, "5.9.215", "",    "notaccept",        true, 1, ""}},
	}; EngineTest(t, "OpenSMTPD", secretlist, false)
}

