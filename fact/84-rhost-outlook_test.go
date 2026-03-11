// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _         ___        _   _             _    
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_      / _ \ _   _| |_| | ___   ___ | | __
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| | | | | | | __| |/ _ \ / _ \| |/ /
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |_| | |_| | |_| | (_) | (_) |   < 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|     \___/ \__,_|\__|_|\___/ \___/|_|\_\
import "testing"

func TestRhostOutlook(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.5.0",   "554", "hostunknown",      true, 1, ""}},
		{{"04",   1, "5.1.1",   "550", "userunknown",      true, 1, ""},
		 {"04",   2, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"06",   1, "4.4.7",   "",    "expired",         false, 0, ""}},
		{{"07",   1, "4.4.7",   "",    "expired",         false, 0, ""}},
		{{"08",   1, "5.5.0",   "550", "userunknown",      true, 1, ""}},
		{{"09",   1, "5.5.0",   "550", "requireptr",      false, 0, ""}},
	}; EngineTest(t, "Outlook", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "Outlook", secretlist, false)
}

