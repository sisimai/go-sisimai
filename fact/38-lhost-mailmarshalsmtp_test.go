// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        __  __       _ _ __  __                _           _ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  \/  | __ _(_) |  \/  | __ _ _ __ ___| |__   __ _| |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |\/| |/ _` | | | |\/| |/ _` | '__/ __| '_ \ / _` | |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |  | | (_| | | | |  | | (_| | |  \__ \ | | | (_| | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  |_|\__,_|_|_|_|  |_|\__,_|_|  |___/_| |_|\__,_|_|
import "testing"

func TestLhostMailMarshalSMTP(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "MailMarshalSMTP", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.3.0",   "553", "filtered",        false, ""},
		 {"1001", 2, "5.3.0",   "553", "filtered",        false, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "MailMarshalSMTP", secretlist, false)
}

