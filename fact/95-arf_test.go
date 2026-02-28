// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___    ____  _____ 
// |_   _|__  ___| |_   / / \  |  _ \|  ___|
//   | |/ _ \/ __| __| / / _ \ | |_) | |_   
//   | |  __/\__ \ |_ / / ___ \|  _ <|  _|  
//   |_|\___||___/\__/_/_/   \_\_| \_\_|    
import "testing"

func TestARF(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"02",   1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"11",   1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"12",   1, "",        "",    "feedback",        false, 1, "opt-out"}},
		{{"14",   1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"15",   1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"16",   1, "",        "",    "feedback",        false, 1, "abuse"},
		 {"16",   2, "",        "",    "feedback",        false, 1, "abuse"},
		 {"16",   3, "",        "",    "feedback",        false, 1, "abuse"},
		 {"16",   4, "",        "",    "feedback",        false, 1, "abuse"},
		 {"16",   5, "",        "",    "feedback",        false, 1, "abuse"},
		 {"16",   6, "",        "",    "feedback",        false, 1, "abuse"},
		 {"16",   7, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"17",   1, "",        "",    "feedback",        false, 1, "abuse"},
		 {"17",   2, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"18",   1, "",        "",    "feedback",        false, 0, "auth-failure"}},
		{{"19",   1, "",        "",    "feedback",        false, 0, "auth-failure"}},
		{{"20",   1, "",        "",    "feedback",        false, 0, "auth-failure"}},
		{{"21",   1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"25",   1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"26",   1, "",        "",    "feedback",        false, 1, "opt-out"}},
	}; EngineTest(t, "ARF", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1002", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1003", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1004", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1005", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1005", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1007", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1009", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1010", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1011", 1,"",         "",    "feedback",        false, 1, "opt-out"}},
		{{"1012", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1015", 1,"",         "",    "feedback",        false, 1, "abuse"}},
		{{"1016", 1,"",         "",    "feedback",        false, 0, "auth-failure"}},
	}; EngineTest(t, "ARF", secretlist, false)
}

