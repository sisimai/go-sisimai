// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"02",   1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"11",   1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"12",   1, "",        "",    "feedback",        false,  true, "opt-out"}},
		{{"14",   1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"15",   1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"16",   1, "",        "",    "feedback",        false,  true, "abuse"},
		 {"16",   2, "",        "",    "feedback",        false,  true, "abuse"},
		 {"16",   3, "",        "",    "feedback",        false,  true, "abuse"},
		 {"16",   4, "",        "",    "feedback",        false,  true, "abuse"},
		 {"16",   5, "",        "",    "feedback",        false,  true, "abuse"},
		 {"16",   6, "",        "",    "feedback",        false,  true, "abuse"},
		 {"16",   7, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"17",   1, "",        "",    "feedback",        false,  true, "abuse"},
		 {"17",   2, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"18",   1, "",        "",    "feedback",        false, false, "auth-failure"}},
		{{"19",   1, "",        "",    "feedback",        false, false, "auth-failure"}},
		{{"20",   1, "",        "",    "feedback",        false, false, "auth-failure"}},
		{{"21",   1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"25",   1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"26",   1, "",        "",    "feedback",        false,  true, "opt-out"}},
	}; EngineTest(t, "ARF", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1002", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1003", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1004", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1005", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1005", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1007", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1009", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1010", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1011", 1,"",         "",    "feedback",        false,  true,  "opt-out"}},
		{{"1012", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1015", 1,"",         "",    "feedback",        false,  true,  "abuse"}},
		{{"1016", 1,"",         "",    "feedback",        false, false,  "auth-failure"}},
	}; EngineTest(t, "ARF", secretlist, false)
}

