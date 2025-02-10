// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _       ___       _            ____                  __  __ ____ ____  
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    |_ _|_ __ | |_ ___ _ __/ ___|  ___ __ _ _ __ |  \/  / ___/ ___| 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| || '_ \| __/ _ \ '__\___ \ / __/ _` | '_ \| |\/| \___ \___ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| || | | | ||  __/ |   ___) | (_| (_| | | | | |  | |___) |__) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   |___|_| |_|\__\___|_|  |____/ \___\__,_|_| |_|_|  |_|____/____/ 
import "testing"

func TestLhostInterScanMSS(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"03",   1, "5.0.911", "",    "userunknown",      true, ""}},
	}; EngineTest(t, "InterScanMSS", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1005", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1007", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1008", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1009", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1010", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1011", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1012", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1013", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1014", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1015", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1016", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1017", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1018", 1, "5.0.911", "",    "userunknown",      true, ""}},
	}; EngineTest(t, "InterScanMSS", secretlist, false)
}

