// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _       ___ __  __       _ _ ____                           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    |_ _|  \/  | __ _(_) / ___|  ___ _ ____   _____ _ __ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| || |\/| |/ _` | | \___ \ / _ \ '__\ \ / / _ \ '__|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| || |  | | (_| | | |___) |  __/ |   \ V /  __/ |   
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   |___|_|  |_|\__,_|_|_|____/ \___|_|    \_/ \___|_|   
import "testing"

func TestLhostIMailServer(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"02",   1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"03",   1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"04",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"06",   1, "5.9.164", "550", "spamdetected",    false, 0, ""}},

	}; EngineTest(t, "IMailServer", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"1002", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1003", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1006", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"1009", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1010", 1, "5.9.340", "",    "expired",         false, 0, ""},
		 {"1010", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1011", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1012", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1013", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1014", 1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"1015", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
//		{{"1016", 1, "5.9.340", "",    "expired",         false, 0, ""}}, // Invalid Date: field
		{{"1017", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1018", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1019", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1020", 1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"1021", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1022", 1, "5.9.213", "",    "userunknown",      true, 1, ""},
		 {"1022", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1023", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1024", 1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"1025", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1026", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1027", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1028", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1029", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1030", 1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"1031", 1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"1032", 1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"1033", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1034", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1035", 1, "5.9.164", "550", "spamdetected",    false, 0, ""}},
		{{"1036", 1, "5.9.164", "550", "spamdetected",    false, 0, ""}},
		{{"1037", 1, "5.9.164", "550", "spamdetected",    false, 0, ""}},
		{{"1038", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
	}; EngineTest(t, "IMailServer", secretlist, false)
}

