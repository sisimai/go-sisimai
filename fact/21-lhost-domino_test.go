// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        ____                  _             
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  _ \  ___  _ __ ___ (_)_ __   ___  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| | | |/ _ \| '_ ` _ \| | '_ \ / _ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | (_) | | | | | | | | | | (_) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |____/ \___/|_| |_| |_|_|_| |_|\___/ 
import "testing"

func TestLhostDomino(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"02",   1, "5.0.0",   "",    "userunknown",      true, ""}},
		{{"03",   1, "5.0.0",   "",    "networkerror",    false, ""}},
	}; EngineTest(t, "Domino", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1002", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1003", 1, "5.0.0",   "",    "userunknown",      true, ""}},
		{{"1004", 1, "5.0.0",   "",    "userunknown",      true, ""}},
		{{"1005", 1, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1006", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1007", 1, "5.0.0",   "",    "userunknown",      true, ""}},
		{{"1008", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1009", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1010", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1011", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1012", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1013", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1014", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1015", 1, "5.0.0",   "",    "networkerror",    false, ""}},
		{{"1016", 1, "5.0.0",   "",    "systemerror",     false, ""}},
		{{"1017", 1, "5.0.0",   "",    "userunknown",      true, ""}},
		{{"1019", 1, "5.0.0",   "",    "userunknown",      true, ""}},
	}; EngineTest(t, "Domino", secretlist, false)
}

