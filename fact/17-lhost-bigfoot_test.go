// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        ____  _        __             _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     | __ )(_) __ _ / _| ___   ___ | |_ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____|  _ \| |/ _` | |_ / _ \ / _ \| __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_) | | (_| |  _| (_) | (_) | |_ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |____/|_|\__, |_|  \___/ \___/ \__|
//                                                            |___/                     
import "testing"

func TestLhostBigfoot(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"02",   1, "5.7.1",   "553", "userunknown",      true, ""}},
	}; EngineTest(t, "Bigfoot", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1002", 1, "5.7.1",   "553", "userunknown",      true, ""}},
	}; EngineTest(t, "Bigfoot", secretlist, false)
}

