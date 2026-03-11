// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _          _         _ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_       / \   ___ | |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____ / _ \ / _ \| |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____/ ___ \ (_) | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   /_/   \_\___/|_|
import "testing"

func TestLhostAol(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "Aol", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.4.4",   "",    "hostunknown",      true, 1, ""}},
		{{"1002", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1003", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""},
		 {"1003", 2, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""},
		 {"1004", 2, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.2.2",   "552", "mailboxfull",     false, 1, ""}},
		{{"1008", 1, "5.7.1",   "554", "filtered",        false, 1, ""}},
		{{"1009", 1, "5.7.1",   "554", "policyviolation", false, 0, ""}},
		{{"1010", 1, "5.7.1",   "554", "filtered",        false, 1, ""}},
		{{"1011", 1, "5.7.1",   "554", "filtered",        false, 1, ""}},
		{{"1012", 1, "5.2.2",   "552", "mailboxfull",     false, 1, ""}},
		{{"1013", 1, "5.2.2",   "552", "mailboxfull",     false, 1, ""}},
		{{"1014", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "Aol", secretlist, false)
}

