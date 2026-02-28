// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        __  __       _ _ ____        
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  \/  | __ _(_) |  _ \ _   _ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |\/| |/ _` | | | |_) | | | |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |  | | (_| | | |  _ <| |_| |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  |_|\__,_|_|_|_| \_\\__,_|
//                                                                                 
import "testing"

func TestLhostMailRu(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"03",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""},
		 {"03",   2, "5.2.1",   "550", "userunknown",      true, 1, ""}},
		{{"04",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"05",   1, "5.9.215", "",    "notaccept",        true, 1, ""}},
		{{"06",   1, "5.9.212", "",    "hostunknown",      true, 1, ""}},
		{{"07",   1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"08",   1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"09",   1, "5.1.8",   "501", "rejected",        false, 0, ""}},
		{{"10",   1, "5.9.340", "",    "expired",         false, 0, ""}},
	}; EngineTest(t, "MailRu", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1003", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""},
		 {"1004", 2, "5.2.1",   "550", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.9.210", "",    "filtered",        false, 0, ""}},
		{{"1006", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1007", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1009", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1010", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1011", 1, "5.1.8",   "501", "rejected",        false, 0, ""}},
	}; EngineTest(t, "MailRu", secretlist, false)
}

