// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _                  _____ ___ _   _____ _____ ____  
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      _ __ ___ |  ___|_ _| | |_   _| ____|  _ \ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| '_ ` _ \| |_   | || |   | | |  _| | |_) |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| | | | | |  _|  | || |___| | | |___|  _ < 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_| |_| |_|_|   |___|_____|_| |_____|_| \_\
import "testing"

func TestLhostmFILTER(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"03",   1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"04",   1, "5.4.1",   "550", "rejected",        false, false, ""}},
		{{"05",   1, "4.3.1",   "452", "systemfull",      false, false, ""}},
	}; EngineTest(t, "mFILTER", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1003", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1004", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1006", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1007", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1008", 1, "5.4.1",   "550", "rejected",        false, false, ""}},
		{{"1009", 1, "5.4.1",   "550", "rejected",        false, false, ""}},
		{{"1010", 1, "4.3.1",   "452", "systemfull",      false, false, ""}},
		{{"1011", 1, "5.6.0",   "550", "spamdetected",    false, false, ""}},
		{{"1012", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1013", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1014", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
	}; EngineTest(t, "mFILTER", secretlist, false)
}

