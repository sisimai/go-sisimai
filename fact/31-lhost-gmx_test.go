// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _         ____ __  ____  __
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      / ___|  \/  \ \/ /
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |  _| |\/| |\  / 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | |  | |/  \ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \____|_|  |_/_/\_\
import "testing"

func TestLhostGMX(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01", 1, "5.2.2",   "",    "mailboxfull",     false, ""}},
		{{"02", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"03", 1, "5.2.1",   "",    "userunknown",      true, ""},
		 {"03", 2, "5.2.2",   "",    "mailboxfull",     false, ""}},
		{{"04", 1, "5.0.947", "",    "expired",         false, ""}},
	}; EngineTest(t, "GMX", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.947", "",    "expired",         false, ""}},
		{{"1002", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1003", 1, "5.2.2",   "",    "mailboxfull",     false, ""}},
		{{"1004", 1, "5.2.1",   "",    "userunknown",      true, ""},
		 {"1004", 2, "5.2.2",   "",    "mailboxfull",     false, ""}},
	}; EngineTest(t, "GMX", secretlist, false)
}

