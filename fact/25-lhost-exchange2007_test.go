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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"03",   1, "5.2.3",   "550", "exceedlimit",     false, false, ""}},
		{{"04",   1, "5.7.1",   "550", "securityerror",   false, false, ""}},
		{{"05",   1, "4.4.1",   "",    "expired",         false, false, ""}},
		{{"06",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"07",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
	}; EngineTest(t, "Exchange2007", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1002", 1, "5.2.3",   "550", "exceedlimit",     false, false, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1005", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1006", 1, "5.2.3",   "550", "exceedlimit",     false, false, ""}},
		{{"1007", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1008", 1, "5.7.1",   "550", "securityerror",   false, false, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1010", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1011", 1, "5.2.3",   "550", "exceedlimit",     false, false, ""}},
		{{"1012", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1013", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1014", 1, "4.2.0",   "",    "systemerror",     false, false, ""}},
		{{"1015", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1016", 1, "5.2.3",   "550", "exceedlimit",     false, false, ""}},
		{{"1017", 1, "5.1.10",  "550", "userunknown",      true,  true, ""}},
		{{"1018", 1, "5.1.10",  "550", "userunknown",      true,  true, ""}},
		{{"1019", 1, "5.4.317", "550", "failedstarttls",  false, false, ""}},
		{{"1020", 1, "5.7.23",  "550", "authfailure",     false, false, ""}},
		{{"1021", 1, "5.7.509", "550", "authfailure",     false, false, ""}},
		{{"1022", 1, "5.4.317", "550", "failedstarttls",  false, false, ""}},
		{{"1023", 1, "5.4.317", "550", "failedstarttls",  false, false, ""}},
		{{"1024", 1, "5.4.318", "550", "systemerror",     false, false, ""}},
		{{"1025", 1, "5.1.351", "550", "userunknown",      true,  true, ""}},
		{{"1026", 1, "4.2.0",   "",    "systemerror",     false, false, ""}},
		{{"1027", 1, "5.4.3",   "550", "systemerror",     false, false, ""}},
		{{"1028", 1, "5.7.520", "550", "securityerror",   false, false, ""}},
		{{"1029", 1, "5.7.1",   "550", "policyviolation", false, false, ""}},
		{{"1030", 1, "5.4.317", "550", "expired",         false, false, ""}},
		{{"1031", 1, "5.1.351", "550", "filtered",        false,  true, ""}},
	}; EngineTest(t, "Exchange2007", secretlist, false)
}

