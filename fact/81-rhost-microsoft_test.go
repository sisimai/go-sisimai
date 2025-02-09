// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _        __  __ _                           __ _   
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_     |  \/  (_) ___ _ __ ___  ___  ___  / _| |_ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |\/| | |/ __| '__/ _ \/ __|/ _ \| |_| __|
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |  | | | (__| | | (_) \__ \ (_) |  _| |_ 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|    |_|  |_|_|\___|_|  \___/|___/\___/|_|  \__|
import "testing"

func TestRhostMicrosoft(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.7.606", "550", "blocked",         false, ""}},
		{{"02",   1, "5.4.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "5.1.10",  "550", "userunknown",      true, ""}},
		{{"04",   1, "5.7.509", "550", "authfailure",     false, ""}},
		{{"05",   1, "4.7.650", "451", "badreputation",   false, ""}},
	}; EngineTest(t, "Microsoft", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "Microsoft", secretlist, false)
}

