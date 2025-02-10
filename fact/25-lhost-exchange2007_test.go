// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        _____          _                            ____   ___   ___ _____ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     | ____|_  _____| |__   __ _ _ __   __ _  ___|___ \ / _ \ / _ \___  |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____|  _| \ \/ / __| '_ \ / _` | '_ \ / _` |/ _ \ __) | | | | | | | / / 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |___ >  < (__| | | | (_| | | | | (_| |  __// __/| |_| | |_| |/ /  
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_____/_/\_\___|_| |_|\__,_|_| |_|\__, |\___|_____|\___/ \___//_/   
//                                                                                     |___/                             
import "testing"

func TestLhostExchange2007(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"03",   1, "5.2.3",   "550", "exceedlimit",     false, ""}},
		{{"04",   1, "5.7.1",   "550", "securityerror",   false, ""}},
		{{"05",   1, "4.4.1",   "",    "expired",         false, ""}},
		{{"06",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"07",   1, "5.1.1",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Exchange2007", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1002", 1, "5.2.3",   "550", "exceedlimit",     false, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1005", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1006", 1, "5.2.3",   "550", "exceedlimit",     false, ""}},
		{{"1007", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1008", 1, "5.7.1",   "550", "securityerror",   false, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1010", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1011", 1, "5.2.3",   "550", "exceedlimit",     false, ""}},
		{{"1012", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1013", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1014", 1, "4.2.0",   "",    "systemerror",     false, ""}},
		{{"1015", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1016", 1, "5.2.3",   "550", "exceedlimit",     false, ""}},
		{{"1017", 1, "5.1.10",  "550", "userunknown",      true, ""}},
		{{"1018", 1, "5.1.10",  "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Exchange2007", secretlist, false)
}

