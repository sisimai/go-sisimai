// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact
//  _____         _      ___ _               _        __  __                           _             ____                           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  \/  | ___  ___ ___  __ _  __ _(_)_ __   __ _/ ___|  ___ _ ____   _____ _ __ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |\/| |/ _ \/ __/ __|/ _` |/ _` | | '_ \ / _` \___ \ / _ \ '__\ \ / / _ \ '__|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |  | |  __/\__ \__ \ (_| | (_| | | | | | (_| |___) |  __/ |   \ V /  __/ |   
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  |_|\___||___/___/\__,_|\__, |_|_| |_|\__, |____/ \___|_|    \_/ \___|_|   
//                                                                               |___/         |___/                                
import "testing"

func TestLhostMessagingServer(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"02",   1, "5.2.0",   "",    "mailboxfull",     false,  true, ""}},
		{{"03",   1, "5.7.1",   "550", "filtered",        false,  true, ""},
		 {"03",   2, "5.7.1",   "550", "filtered",        false,  true, ""}},
		{{"04",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"05",   1, "5.4.4",   "",    "hostunknown",      true,  true, ""}},
		{{"06",   1, "5.2.1",   "550", "filtered",        false,  true, ""}},
		{{"07",   1, "4.4.7",   "",    "expired",         false, false, ""}},
		{{"08",   1, "5.0.0",   "550", "filtered",        false,  true, ""}},
		{{"09",   1, "5.0.0",   "550", "userunknown",      true,  true, ""}},
		{{"10",   1, "5.1.10",  "",    "notaccept",        true,  true, ""}},
		{{"11",   1, "5.1.8",   "501", "rejected",        false, false, ""}},
		{{"12",   1, "4.2.2",   "",    "mailboxfull",     false, false, ""}},
	}; EngineTest(t, "MessagingServer", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.4.4",   "",    "hostunknown",      true,  true, ""}},
		{{"1002", 1, "5.0.0",   "",    "mailboxfull",     false,  true, ""}},
		{{"1003", 1, "5.7.1",   "550", "filtered",        false,  true, ""},
		 {"1003", 1, "5.7.1",   "550", "filtered",        false,  true, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1005", 1, "5.4.4",   "",    "hostunknown",      true,  true, ""}},
		{{"1006", 1, "5.7.1",   "550", "filtered",        false,  true, ""}},
		{{"1007", 1, "5.2.0",   "",    "mailboxfull",     false,  true, ""}},
		{{"1008", 1, "5.2.1",   "550", "filtered",        false,  true, ""}},
		{{"1009", 1, "5.0.0",   "",    "mailboxfull",     false,  true, ""}},
		{{"1010", 1, "5.2.0",   "",    "mailboxfull",     false,  true, ""}},
		{{"1011", 1, "4.4.7",   "",    "expired",         false, false, ""}},
		{{"1012", 1, "5.0.0",   "550", "filtered",        false,  true, ""}},
		{{"1013", 1, "4.2.2",   "",    "mailboxfull",     false, false, ""}},
		{{"1014", 1, "4.2.2",   "",    "mailboxfull",     false, false, ""}},
		{{"1015", 1, "5.0.0",   "550", "filtered",        false,  true, ""}},
		{{"1016", 1, "5.0.0",   "550", "userunknown",      true,  true, ""}},
		{{"1017", 1, "5.1.10",  "",    "notaccept",        true,  true, ""}},
		{{"1018", 1, "5.1.8",   "501", "rejected",        false, false, ""}},
		{{"1019", 1, "4.2.2",   "",    "mailboxfull",     false, false, ""}},
	}; EngineTest(t, "MessagingServer", secretlist, false)
}

