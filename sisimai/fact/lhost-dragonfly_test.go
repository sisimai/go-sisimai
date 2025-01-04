// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.7.26",  "550", "authfailure",     false, ""}},
		{{"02",   1, "5.7.509", "550", "authfailure",     false, ""}},
		{{"03",   1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"04",   1, "5.0.912", "",    "hostunknown",      true, ""}},
		{{"05",   1, "5.7.26",  "550", "authfailure",     false, ""}},
		{{"06",   1, "5.7.25",  "550", "requireptr",      false, ""}},
		{{"07",   1, "5.6.0",   "550", "contenterror",    false, ""}},
		{{"08",   1, "5.2.3",   "552", "exceedlimit",     false, ""}},
		{{"09",   1, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"10",   1, "5.1.6",   "550", "hasmoved",         true, ""}},
		{{"11",   1, "5.1.2",   "550", "hostunknown",      true, ""}},
		{{"12",   1, "5.2.2",   "552", "mailboxfull",     false, ""}},
		{{"13",   1, "5.3.0",   "554", "mailererror",     false, ""}},
		{{"14",   1, "5.3.4",   "554", "mesgtoobig",      false, ""}},
		{{"15",   1, "5.7.0",   "550", "norelaying",      false, ""}},
		{{"16",   1, "5.3.2",   "521", "notaccept",        true, ""}},
		{{"17",   1, "5.0.0",   "550", "onhold",          false, ""}},
		{{"18",   1, "5.7.0",   "550", "securityerror",   false, ""}},
		{{"19",   1, "5.7.1",   "551", "securityerror",   false, ""}},
		{{"20",   1, "5.7.0",   "550", "spamdetected",    false, ""}},
		{{"21",   1, "5.7.13",  "525", "suspend",         false, ""}},
		{{"22",   1, "5.1.3",   "501", "userunknown",      true, ""}},
		{{"23",   1, "5.3.0",   "554", "systemerror",     false, ""}},
		{{"24",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"25",   1, "5.7.0",   "550", "virusdetected",   false, ""}},
		{{"26",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"27",   1, "5.7.13",  "525", "suspend",         false, ""}},
		{{"28",   1, "5.2.2",   "552", "mailboxfull",     false, ""}},
		{{"29",   1, "5.0.947", "",    "expired",         false, ""}},
		{{"30",   1, "5.0.947", "",    "expired",         false, ""}},
	}; EngineTest(t, "DragonFly", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.7.26",  "550", "authfailure",     false, ""}},
		{{"1002", 1, "5.0.947", "",    "expired",         false, ""}},
	}; EngineTest(t, "DragonFly", secretlist, false)
}

