// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _         ____ ____        _ _       
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_      / ___/ ___| _   _(_) |_ ___ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |  _\___ \| | | | | __/ _ \
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |_| |___) | |_| | | ||  __/
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|     \____|____/ \__,_|_|\__\___|
import "testing"

func TestRhostGSuite(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.0",   "550", "userunknown",      true,  true, ""}},
		{{"02",   1, "5.0.0",   "",    "userunknown",      true,  true, ""}},
		{{"03",   1, "4.0.0",   "",    "notaccept",       false, false, ""}},
		{{"04",   1, "4.0.0",   "",    "networkerror",    false, false, ""}},
		{{"05",   1, "4.0.0",   "",    "networkerror",    false, false, ""}},
		{{"06",   1, "4.4.1",   "",    "expired",         false, false, ""}},
		{{"07",   1, "4.4.1",   "",    "expired",         false, false, ""}},
		{{"08",   1, "5.0.0",   "550", "filtered",        false,  true, ""}},
		{{"09",   1, "5.0.0",   "550", "userunknown",      true,  true, ""}},
		{{"10",   1, "4.0.0",   "",    "notaccept",       false, false, ""}},
		{{"11",   1, "5.1.8",   "501", "rejected",        false, false, ""}},
		{{"12",   1, "5.0.0",   "",    "spamdetected",    false, false, ""}},
		{{"13",   1, "4.0.0",   "",    "networkerror",    false, false, ""}},
		{{"14",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
	}; EngineTest(t, "GSuite", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "GSuite", secretlist, false)
}

