// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _        _____                         ____ _____ _____ 
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_     |  ___| __ __ _ _ __   ___ ___|  _ \_   _|_   _|
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |_ | '__/ _` | '_ \ / __/ _ \ |_) || |   | |  
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____|  _|| | | (_| | | | | (_|  __/  __/ | |   | |  
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|    |_|  |_|  \__,_|_| |_|\___\___|_|    |_|   |_|  
import "testing"

func TestRhostFrancePTT(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.5.0",   "550", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.2.0",   "550", "spamdetected",    false, 0, ""}},
		{{"04",   1, "5.2.0",   "550", "spamdetected",    false, 0, ""}},
		{{"05",   1, "5.5.0",   "550", "suspend",         false, 1, ""}},
		{{"06",   1, "4.0.0",   "",    "blocked",         false, 0, ""}},
		{{"07",   1, "4.0.0",   "421", "ratelimited",     false, 0, ""}},
		{{"08",   1, "4.2.0",   "421", "systemerror",     false, 0, ""}},
		{{"10",   1, "5.5.0",   "550", "blocked",         false, 0, ""}},
		{{"11",   1, "4.2.1",   "421", "requireptr",      false, 0, ""}},
		{{"12",   1, "5.7.1",   "554", "policyviolation", false, 0, ""}},
	}; EngineTest(t, "FrancePTT", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "FrancePTT", secretlist, false)
}

