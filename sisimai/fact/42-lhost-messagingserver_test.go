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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.2.0",   "",    "mailboxfull",     false, ""}},
		{{"03",   1, "5.7.1",   "550", "filtered",        false, ""},
		 {"03",   2, "5.7.1",   "550", "filtered",        false, ""}},
		{{"04",   1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"05",   1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"06",   1, "5.2.1",   "550", "filtered",        false, ""}},
		{{"07",   1, "4.4.7",   "",    "expired",         false, ""}},
		{{"08",   1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"09",   1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"10",   1, "5.1.10",  "",    "notaccept",        true, ""}},
		{{"11",   1, "5.1.8",   "501", "rejected",        false, ""}},
		{{"12",   1, "4.2.2",   "",    "mailboxfull",     false, ""}},
	}; EngineTest(t, "MessagingServer", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"1002", 1, "5.0.0",   "",    "mailboxfull",     false, ""}},
		{{"1003", 1, "5.7.1",   "550", "filtered",        false, ""},
		 {"1003", 1, "5.7.1",   "550", "filtered",        false, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1005", 1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"1006", 1, "5.7.1",   "550", "filtered",        false, ""}},
		{{"1007", 1, "5.2.0",   "",    "mailboxfull",     false, ""}},
		{{"1008", 1, "5.2.1",   "550", "filtered",        false, ""}},
		{{"1009", 1, "5.0.0",   "",    "mailboxfull",     false, ""}},
		{{"1010", 1, "5.2.0",   "",    "mailboxfull",     false, ""}},
		{{"1011", 1, "4.4.7",   "",    "expired",         false, ""}},
		{{"1012", 1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"1013", 1, "4.2.2",   "",    "mailboxfull",     false, ""}},
		{{"1014", 1, "4.2.2",   "",    "mailboxfull",     false, ""}},
		{{"1015", 1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"1016", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1017", 1, "5.1.10",  "",    "notaccept",        true, ""}},
		{{"1018", 1, "5.1.8",   "501", "rejected",        false, ""}},
		{{"1019", 1, "4.2.2",   "",    "mailboxfull",     false, ""}},
	}; EngineTest(t, "MessagingServer", secretlist, false)
}

