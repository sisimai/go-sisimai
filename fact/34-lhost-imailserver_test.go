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
		{{"01",   1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"02",   1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"03",   1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"04",   1, "5.9.340", "",    "expired",         false, false, ""}},
		{{"06",   1, "5.9.164", "550", "spamdetected",    false, false, ""}},

	}; EngineTest(t, "IMailServer", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.212", "",    "hostunknown",      true,  true, ""}},
		{{"1002", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1003", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1004", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1005", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1006", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1007", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1008", 1, "5.9.212", "",    "hostunknown",      true,  true, ""}},
		{{"1009", 1, "5.9.340", "",    "expired",         false, false, ""}},
		{{"1010", 1, "5.9.340", "",    "expired",         false, false, ""},
		 {"1010", 1, "5.9.340", "",    "expired",         false, false, ""}},
		{{"1011", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1012", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1013", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1014", 1, "5.9.212", "",    "hostunknown",      true,  true, ""}},
		{{"1015", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
//		{{"1016", 1, "5.9.340", "",    "expired",         false, false, ""}}, // Invalid Date: field
		{{"1017", 1, "5.9.340", "",    "expired",         false, false, ""}},
		{{"1018", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1019", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1020", 1, "5.9.301", "",    "onhold",          false, false, ""}},
		{{"1021", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1022", 1, "5.9.213", "",    "userunknown",      true,  true, ""},
		 {"1022", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1023", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1024", 1, "5.9.301", "",    "onhold",          false, false, ""}},
		{{"1025", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1026", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1027", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1028", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1029", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1030", 1, "5.9.212", "",    "hostunknown",      true,  true, ""}},
		{{"1031", 1, "5.9.212", "",    "hostunknown",      true,  true, ""}},
		{{"1032", 1, "5.9.301", "",    "onhold",          false, false, ""}},
		{{"1033", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1034", 1, "5.9.213", "",    "userunknown",      true,  true, ""}},
		{{"1035", 1, "5.9.164", "550", "spamdetected",    false, false, ""}},
		{{"1036", 1, "5.9.164", "550", "spamdetected",    false, false, ""}},
		{{"1037", 1, "5.9.164", "550", "spamdetected",    false, false, ""}},
		{{"1038", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
	}; EngineTest(t, "IMailServer", secretlist, false)
}

