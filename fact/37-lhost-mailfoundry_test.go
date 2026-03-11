// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        __  __       _ _ _____                     _            
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  \/  | __ _(_) |  ___|__  _   _ _ __   __| |_ __ _   _ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |\/| |/ _` | | | |_ / _ \| | | | '_ \ / _` | '__| | | |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |  | | (_| | | |  _| (_) | |_| | | | | (_| | |  | |_| |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  |_|\__,_|_|_|_|  \___/ \__,_|_| |_|\__,_|_|   \__, |
//                                                                                                      |___/ 
import "testing"

func TestLhostMailFoundry(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"02",   1, "5.1.1",   "552", "mailboxfull",     false, 1, ""}},
	}; EngineTest(t, "MailFoundry", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1002", 1, "5.1.1",   "552", "mailboxfull",     false, 1, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1005", 1, "5.1.1",   "552", "mailboxfull",     false, 1, ""}},
	}; EngineTest(t, "MailFoundry", secretlist, false)
}

