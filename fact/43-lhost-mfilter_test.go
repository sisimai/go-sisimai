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
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.910", "550", "filtered",        false, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "5.0.910", "550", "filtered",        false, ""}},
		{{"04",   1, "5.4.1",   "550", "rejected",        false, ""}},
	}; EngineTest(t, "mFILTER", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1004", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1006", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1007", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1008", 1, "5.4.1",   "550", "rejected",        false, ""}},
		{{"1009", 1, "5.4.1",   "550", "rejected",        false, ""}},
	}; EngineTest(t, "mFILTER", secretlist, false)
}

