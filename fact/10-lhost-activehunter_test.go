// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _                   _   _           _                 _            
// |_   _|__  ___| |_   / / | |__   ___  ___| |_       __ _  ___| |_(_)_   _____| |__  _   _ _ __ | |_ ___ _ __ 
//   | |/ _ \/ __| __| / /| | "_ \ / _ \/ __| __|____ / _` |/ __| __| \ \ / / _ \ "_ \| | | | "_ \| __/ _ \ "__|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| (_| | (__| |_| |\ V /  __/ | | | |_| | | | | ||  __/ |   
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \__,_|\___|\__|_| \_/ \___|_| |_|\__,_|_| |_|\__\___|_|   
import "testing"

func TestLhostActivehunter(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.0.910", "550", "filtered",        false, ""}},
	}
	EngineTest(t, "Activehunter", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.3.0"  , "553", "filtered",        false, ""}},
		{{"1004", 1, "5.7.17",  "550", "filtered",        false, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1007", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1008", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1010", 1, "5.3.0",   "553", "filtered",        false, ""}},
		{{"1011", 1, "5.7.17",  "550", "filtered",        false, ""}},
		{{"1012", 1, "5.1.1",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Activehunter", secretlist, false)
}

