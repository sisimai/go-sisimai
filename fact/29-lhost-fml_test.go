// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        _____ __  __ _     
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  ___|  \/  | |    
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |_  | |\/| | |    
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____|  _| | |  | | |___ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|   |_|  |_|_____|
import "testing"

func TestLhostFML(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"02",   1, "5.0.918", "",    "rejected",        false, ""}},
		{{"03",   1, "5.0.930", "",    "systemerror",     false, ""}},
	}; EngineTest(t, "FML", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.930", "",    "systemerror",     false, ""}},
		{{"1002", 1, "5.0.918", "",    "rejected",        false, ""}},
	}; EngineTest(t, "FML", secretlist, false)
}

