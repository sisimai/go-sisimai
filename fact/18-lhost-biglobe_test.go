// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        ____  _       _       _          
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     | __ )(_) __ _| | ___ | |__   ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____|  _ \| |/ _` | |/ _ \| '_ \ / _ \
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_) | | (_| | | (_) | |_) |  __/
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |____/|_|\__, |_|\___/|_.__/ \___|
//                                                            |___/                    
import "testing"

func TestLhostBiglobe(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
	}; EngineTest(t, "Biglobe", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1002", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1003", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1004", 1, "5.9.220", "",    "mailboxfull",     false, false, ""}},
		{{"1005", 1, "5.9.210", "",    "filtered",        false, false, ""}},
		{{"1006", 1, "5.9.210", "",    "filtered",        false, false, ""}},
	}; EngineTest(t, "Biglobe", secretlist, false)
}

