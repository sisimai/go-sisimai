// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"03",   1, "5.2.2",   "550", "mailboxfull",     false, ""},
		 {"03",   2, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"04",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"05",   1, "5.0.932", "",    "notaccept",        true, ""}},
		{{"06",   1, "5.0.912", "",    "hostunknown",      true, ""}},
		{{"07",   1, "5.0.910", "550", "filtered",        false, ""}},
		{{"08",   1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"09",   1, "5.1.8",   "501", "rejected",        false, ""}},
		{{"10",   1, "4.0.947", "",    "expired",         false, ""}},
	}; EngineTest(t, "MailRu", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false, ""},
		 {"1004", 2, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"1005", 1, "5.0.910", "",    "filtered",        false, ""}},
		{{"1006", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1007", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1008", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1009", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1010", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1011", 1, "5.1.8",   "501", "rejected",        false, ""}},
	}; EngineTest(t, "MailRu", secretlist, false)
}

