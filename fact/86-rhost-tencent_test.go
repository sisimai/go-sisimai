// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _      _____                         _   
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_   |_   _|__ _ __   ___ ___ _ __ | |_ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |/ _ \ '_ \ / __/ _ \ '_ \| __|
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |  __/ | | | (_|  __/ | | | |_ 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|    |_|\___|_| |_|\___\___|_| |_|\__|
import "testing"

func TestRhostTencent(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.0",   "550", "toomanyconn",     false, ""}},
		{{"02",   1, "5.0.0",   "550", "toomanyconn",     false, ""}},
		{{"03",   1, "5.0.0",   "550", "authfailure",     false, ""}},
	}; EngineTest(t, "Tencent", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "Tencent", secretlist, false)
}

