// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _       ____                 _  ____      _     _ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    / ___|  ___ _ __   __| |/ ___|_ __(_) __| |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|___\___ \ / _ \ '_ \ / _` | |  _| '__| |/ _` |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____|__) |  __/ | | | (_| | |_| | |  | | (_| |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   |____/ \___|_| |_|\__,_|\____|_|  |_|\__,_|
//                                                                                             
import "testing"

func TestLhostSendGrid(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "5.0.947", "",    "expired",         false, ""}},
	}; EngineTest(t, "SendGrid", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.0.947", "",    "expired",         false, ""}},
		{{"1004", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1005", 1, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"1006", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1007", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1008", 1, "5.0.911", "554", "userunknown",      true, ""}},
		{{"1009", 1, "5.0.911", "550", "userunknown",      true, ""}},
	}; EngineTest(t, "SendGrid", secretlist, false)
}

