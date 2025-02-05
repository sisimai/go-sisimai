// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _     __   __              _           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   \ \ / /_ _ _ __   __| | _____  __
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|___\ V / _` | '_ \ / _` |/ _ \ \/ /
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| | (_| | | | | (_| |  __/>  < 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|\__,_|_| |_|\__,_|\___/_/\_\
//                                                                                  
import "testing"

func TestLhostYandex(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.2.1",   "550", "userunknown",      true, ""},
		 {"02",   2, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"03",   1, "4.4.1",   "",    "expired",         false, ""}},
	}; EngineTest(t, "Yandex", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1002", 1, "5.2.1",   "550", "userunknown",      true, ""},
		 {"1002", 2, "5.2.2",   "550", "mailboxfull",     false, ""}},
	}; EngineTest(t, "Yandex", secretlist, false)
}

