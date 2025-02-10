// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _         ____          
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_      / ___|_____  __
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |   / _ \ \/ /
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |__| (_) >  < 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|     \____\___/_/\_\
import "testing"

func TestRhostCox(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.0",   "550", "blocked",         false, ""},
		 {"01",   2, "5.1.0",   "550", "blocked",         false, ""}},
	}; EngineTest(t, "Cox", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "Cox", secretlist, false)
}

