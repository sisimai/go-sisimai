// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _          _                _      
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_       / \   _ __  _ __ | | ___ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____ / _ \ | '_ \| '_ \| |/ _ \
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____/ ___ \| |_) | |_) | |  __/
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|   /_/   \_\ .__/| .__/|_|\___|
//                                                            |_|   |_|           
import "testing"

func TestRhostApple(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.6",   "550", "hasmoved",         true, ""}},
		{{"02",   1, "5.7.1",   "554", "authfailure",     false, ""}},
		{{"03",   1, "5.2.2",   "552", "mailboxfull",     false, ""}},
		{{"04",   1, "5.1.1",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Apple", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "Apple", secretlist, false)
}

