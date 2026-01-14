// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        ____                              _____ _       
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  _ \ _ __ __ _  __ _  ___  _ __ |  ___| |_   _ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| | | | '__/ _` |/ _` |/ _ \| '_ \| |_  | | | | |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | | | (_| | (_| | (_) | | | |  _| | | |_| |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |____/|_|  \__,_|\__, |\___/|_| |_|_|   |_|\__, |
//                                                                    |___/                     |___/ 
import "testing"

func TestLhostDragonFly(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.7.26",  "550", "authfailure",     false, false, ""}},
		{{"02",   1, "5.7.509", "550", "authfailure",     false, false, ""}},
		{{"03",   1, "5.7.9",   "554", "policyviolation", false, false, ""}},
		{{"04",   1, "5.9.212", "",    "hostunknown",      true,  true, ""}},
		{{"05",   1, "5.7.26",  "550", "authfailure",     false, false, ""}},
		{{"06",   1, "5.7.25",  "550", "requireptr",      false, false, ""}},
		{{"07",   1, "5.6.0",   "550", "contenterror",    false, false, ""}},
		{{"08",   1, "5.2.3",   "552", "emailtoolarge",   false, false, ""}},
		{{"09",   1, "5.2.1",   "550", "userunknown",      true,  true, ""}},
		{{"10",   1, "5.1.6",   "550", "hasmoved",         true,  true, ""}},
		{{"11",   1, "5.1.2",   "550", "hostunknown",      true,  true, ""}},
		{{"12",   1, "5.2.2",   "552", "mailboxfull",     false,  true, ""}},
		{{"13",   1, "5.3.0",   "554", "mailererror",     false, false, ""}},
		{{"14",   1, "5.3.4",   "554", "emailtoolarge",   false, false, ""}},
		{{"15",   1, "5.7.0",   "550", "norelaying",      false,  true, ""}},
		{{"16",   1, "5.3.2",   "521", "notaccept",        true,  true, ""}},
		{{"17",   1, "5.0.0",   "550", "onhold",          false, false, ""}},
		{{"18",   1, "5.7.0",   "550", "securityerror",   false, false, ""}},
		{{"19",   1, "5.7.1",   "551", "securityerror",   false, false, ""}},
		{{"20",   1, "5.7.0",   "550", "spamdetected",    false, false, ""}},
		{{"21",   1, "5.7.13",  "525", "suspend",         false,  true, ""}},
		{{"22",   1, "5.1.3",   "501", "userunknown",      true,  true, ""}},
		{{"23",   1, "5.3.0",   "554", "systemerror",     false, false, ""}},
		{{"24",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"25",   1, "5.7.0",   "550", "virusdetected",   false, false, ""}},
		{{"26",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"27",   1, "5.7.13",  "525", "suspend",         false,  true, ""}},
		{{"28",   1, "5.2.2",   "552", "mailboxfull",     false,  true, ""}},
		{{"29",   1, "5.9.340", "",    "expired",         false, false, ""}},
		{{"30",   1, "5.9.340", "",    "expired",         false, false, ""}},
	}; EngineTest(t, "DragonFly", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.7.26",  "550", "authfailure",     false, false, ""}},
		{{"1002", 1, "5.9.340", "",    "expired",         false, false, ""}},
	}; EngineTest(t, "DragonFly", secretlist, false)
}

