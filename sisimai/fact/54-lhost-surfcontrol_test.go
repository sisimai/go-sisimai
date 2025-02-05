// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _       ____              __  ____            _             _ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    / ___| _   _ _ __ / _|/ ___|___  _ __ | |_ _ __ ___ | |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|___\___ \| | | | '__| |_| |   / _ \| '_ \| __| '__/ _ \| |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____|__) | |_| | |  |  _| |__| (_) | | | | |_| | | (_) | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   |____/ \__,_|_|  |_|  \____\___/|_| |_|\__|_|  \___/|_|
//                                                                                                         
import "testing"

func TestLhostSurfControl(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"02",   1, "5.0.0",   "554", "systemerror",     false, ""}},
		{{"03",   1, "5.0.0",   "554", "systemerror",     false, ""}},
	}; EngineTest(t, "SurfControl", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.0",   "550", "filtered",      false, ""}},
		{{"1002", 1, "5.0.0",   "550", "filtered",      false, ""}},
		{{"1003", 1, "5.0.0",   "550", "filtered",      false, ""}},
		{{"1004", 1, "5.0.0",   "554", "systemerror",   false, ""}},
		{{"1005", 1, "5.0.0",   "554", "systemerror",   false, ""}},
	}; EngineTest(t, "SurfControl", secretlist, false)
}

