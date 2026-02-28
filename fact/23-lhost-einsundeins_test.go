// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        _  ___   _ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     / |( _ ) / |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |/ _ \/\ |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| | (_>  < |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|\___/\/_|
import "testing"

func TestLhostEinsUndEins(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"02",   1, "5.9.161", "",    "emailtoolarge",   false, 0, ""}},
		{{"03",   1, "5.2.0",   "550", "spamdetected",    false, 0, ""}},
	}; EngineTest(t, "EinsUndEins", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1002", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1003", 1, "5.9.161", "",    "emailtoolarge",   false, 0, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.4.1",   "550", "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.4.1",   "550", "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.4.1",   "550", "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.4.1",   "550", "userunknown",      true, 1, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1010", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1011", 1, "5.4.1",   "550", "userunknown",      true, 1, ""}},
		{{"1012", 1, "5.4.1",   "550", "userunknown",      true, 1, ""}},
		{{"1013", 1, "5.4.1",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "EinsUndEins", secretlist, false)
}

