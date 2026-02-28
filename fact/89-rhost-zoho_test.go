// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _                _           
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_     _______ | |__   ___  
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|___|_  / _ \| '_ \ / _ \ 
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____/ / (_) | | | | (_) |
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|   /___\___/|_| |_|\___/ 
import "testing"

func TestRhostZoho(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.7.7",   "554", "policyviolation", false, 0, ""}},
		{{"03",   1, "5.7.1",   "554", "rejected",        false, 0, ""}},
		{{"04",   1, "5.4.1",   "",    "rejected",        false, 0, ""}},
	}; EngineTest(t, "Zoho", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "Zoho", secretlist, false)
}

