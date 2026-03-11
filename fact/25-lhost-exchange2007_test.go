// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"03",   1, "5.2.3",   "550", "emailtoolarge",   false, 0, ""}},
		{{"04",   1, "5.7.1",   "550", "securityerror",   false, 0, ""}},
		{{"05",   1, "4.4.1",   "",    "expired",         false, 0, ""}},
		{{"06",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"07",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "Exchange2007", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.2.3",   "550", "emailtoolarge",   false, 0, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1006", 1, "5.2.3",   "550", "emailtoolarge",   false, 0, ""}},
		{{"1007", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1008", 1, "5.7.1",   "550", "securityerror",   false, 0, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1010", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1011", 1, "5.2.3",   "550", "emailtoolarge",   false, 0, ""}},
		{{"1012", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1013", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1014", 1, "4.2.0",   "",    "systemerror",     false, 0, ""}},
		{{"1015", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1016", 1, "5.2.3",   "550", "emailtoolarge",   false, 0, ""}},
		{{"1017", 1, "5.1.10",  "550", "userunknown",      true, 1, ""}},
		{{"1018", 1, "5.1.10",  "550", "userunknown",      true, 1, ""}},
		{{"1019", 1, "5.4.317", "550", "failedstarttls",  false, 0, ""}},
		{{"1020", 1, "5.7.23",  "550", "authfailure",     false, 0, ""}},
		{{"1021", 1, "5.7.509", "550", "authfailure",     false, 0, ""}},
		{{"1022", 1, "5.4.317", "550", "failedstarttls",  false, 0, ""}},
		{{"1023", 1, "5.4.317", "550", "failedstarttls",  false, 0, ""}},
		{{"1024", 1, "5.4.318", "550", "systemerror",     false, 0, ""}},
		{{"1025", 1, "5.1.351", "550", "userunknown",      true, 1, ""}},
		{{"1026", 1, "4.2.0",   "",    "systemerror",     false, 0, ""}},
		{{"1027", 1, "5.4.3",   "550", "systemerror",     false, 0, ""}},
		{{"1028", 1, "5.7.520", "550", "securityerror",   false, 0, ""}},
		{{"1029", 1, "5.7.1",   "550", "policyviolation", false, 0, ""}},
		{{"1030", 1, "5.4.317", "550", "expired",         false, 0, ""}},
		{{"1031", 1, "5.1.351", "550", "filtered",        false, 1, ""}},
		{{"1032", 1, "5.0.350", "550", "norelaying",      false, 1, ""}},
		{{"1033", 1, "5.0.350", "550", "norelaying",      false, 1, ""}},
		{{"1034", 1, "5.7.193", "550", "rejected",        false, 0, ""}},
	}; EngineTest(t, "Exchange2007", secretlist, false)
}

