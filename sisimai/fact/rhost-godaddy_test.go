// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _         ____       ____            _     _       
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_      / ___| ___ |  _ \  __ _  __| | __| |_   _ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |  _ / _ \| | | |/ _` |/ _` |/ _` | | | |
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |_| | (_) | |_| | (_| | (_| | (_| | |_| |
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|     \____|\___/|____/ \__,_|\__,_|\__,_|\__, |
//                                                                                          |___/ 
import "testing"

func TestRhostGoDaddy(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"02",   1, "5.1.3",   "553", "blocked",         false, ""}},
		{{"03",   1, "5.1.1",   "550", "speeding",        false, ""}},
	}; EngineTest(t, "GoDaddy", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "GoDaddy", secretlist, false)
}

