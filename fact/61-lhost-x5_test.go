// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      __  ______  
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    \ \/ / ___| 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____\  /|___ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____/  \ ___) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   /_/\_\____/ 
//                                                              
import "testing"

func TestLhostX5(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "X5", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "X5", secretlist, false)
}

