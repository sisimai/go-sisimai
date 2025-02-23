// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _         ___        _   _             _    
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      / _ \ _   _| |_| | ___   ___ | | __
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| | | | | | | __| |/ _ \ / _ \| |/ /
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | |_| | |_| | (_) | (_) |   < 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \___/ \__,_|\__|_|\___/ \___/|_|\_\
import "testing"

func TestLhostOutlook(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "5.5.0",   "554", "hostunknown",      true, ""}},
		{{"04",   1, "5.1.1",   "550", "userunknown",      true, ""},
		 {"04",   2, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"06",   1, "4.4.7",   "",    "expired",         false, ""}},
		{{"07",   1, "4.4.7",   "",    "expired",         false, ""}},
		{{"08",   1, "5.5.0",   "550", "userunknown",      true, ""}},
		{{"09",   1, "5.5.0",   "550", "requireptr",      false, ""}},
	}; EngineTest(t, "Outlook", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1002", 1, "5.5.0",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.5.0",   "550", "userunknown",      true, ""}},
		{{"1007", 1, "5.5.0",   "550", "requireptr",      false, ""}},
		{{"1008", 1, "5.2.2",   "552", "mailboxfull",     false, ""}},
		{{"1016", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1017", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1018", 1, "5.5.0",   "554", "hostunknown",      true, ""}},
		{{"1019", 1, "5.1.1",   "550", "userunknown",      true, ""},
		 {"1019", 2, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1023", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1024", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1025", 1, "5.5.0",   "550", "filtered",        false, ""}},
		{{"1026", 1, "5.5.0",   "550", "filtered",        false, ""}},
		{{"1027", 1, "5.5.0",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Outlook", secretlist, false)
}

