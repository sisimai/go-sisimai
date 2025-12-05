// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      _____                   _ __  __ _                
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   |_   _| __ ___ _ __   __| |  \/  (_) ___ _ __ ___  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| || '__/ _ \ '_ \ / _` | |\/| | |/ __| '__/ _ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| || | |  __/ | | | (_| | |  | | | (__| | | (_) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_||_|  \___|_| |_|\__,_|_|  |_|_|\___|_|  \___/ 
import "testing"

func TestLhostTrendMicro(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"02",   1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"03",   1, "5.0.911", "",    "userunknown",      true,  true, ""}},
	}; EngineTest(t, "TrendMicro", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1005", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1007", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1008", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1009", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1010", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1011", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1012", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1013", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1014", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1015", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1016", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1017", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
		{{"1018", 1, "5.0.911", "",    "userunknown",      true,  true, ""}},
	}; EngineTest(t, "TrendMicro", secretlist, false)
}

