// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""},
		 {"02",   2, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"03",   1, "5.0.912", "",    "hostunknown",      true,  true, ""}},
		{{"04",   1, "5.0.944", "",    "networkerror",    false, false, ""}},
		{{"05",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"06",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"10",   1, "5.0.912", "",    "hostunknown",      true,  true, ""}},
		{{"11",   1, "5.7.26",  "550", "authfailure",     false, false, ""}},
		{{"12",   1, "5.0.932", "",    "notaccept",        true,  true, ""}},
		{{"13",   1, "4.7.0",   "421", "badreputation",   false, false, ""}},
		{{"14",   1, "5.7.25",  "550", "requireptr",      false, false, ""}},
		{{"15",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"16",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"17",   1, "5.1.1",   "550", "userunknown",      true,  true, ""},
		 {"17",   2, "5.2.2",   "552", "mailboxfull",     false,  true, ""}},
	}; EngineTest(t, "OpenSMTPD", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1002", 1, "5.2.1",   "550", "filtered",        false,  true, ""}},
		{{"1003", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1004", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1005", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1006", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1007", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1008", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""},
		 {"1008", 2, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1009", 1, "5.0.912", "",    "hostunknown",      true,  true, ""}},
		{{"1010", 1, "5.0.944", "",    "networkerror",    false, false, ""}},
		{{"1011", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1012", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""},
		 {"1012", 2, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1013", 1, "5.0.912", "",    "hostunknown",      true,  true, ""}},
		{{"1014", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1015", 1, "5.0.944", "",    "networkerror",    false, false, ""}},
		{{"1016", 1, "5.0.912", "",    "hostunknown",      true,  true, ""}},
		{{"1017", 1, "5.7.26",  "550", "authfailure",     false, false, ""}},
		{{"1018", 1, "5.0.932", "",    "notaccept",        true,  true, ""}},
	}; EngineTest(t, "OpenSMTPD", secretlist, false)
}

