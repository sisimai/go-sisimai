// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _     __   __    _                ___            
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_   \ \ / /_ _| |__   ___   ___|_ _|_ __   ___ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|___\ V / _` | '_ \ / _ \ / _ \| || '_ \ / __|
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| | (_| | | | | (_) | (_) | || | | | (__ 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|    |_|\__,_|_| |_|\___/ \___/___|_| |_|\___|
import "testing"

func TestRhostYahooInc(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"02",   1, "4.7.0",   "421", "rejected",        false, ""}},
		{{"03",   1, "5.0.0",   "554", "userunknown",      true, ""}},
	}; EngineTest(t, "YahooInc", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "YahooInc", secretlist, false)
}

