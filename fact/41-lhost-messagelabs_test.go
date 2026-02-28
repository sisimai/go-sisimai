// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        __  __                                _          _         
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  \/  | ___  ___ ___  __ _  __ _  ___| |    __ _| |__  ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |\/| |/ _ \/ __/ __|/ _` |/ _` |/ _ \ |   / _` | '_ \/ __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |  | |  __/\__ \__ \ (_| | (_| |  __/ |__| (_| | |_) \__ \
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  |_|\___||___/___/\__,_|\__, |\___|_____\__,_|_.__/|___/
//                                                                               |___/                           
import "testing"

func TestLhostMessageLabs(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "MessageLabs", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.0.0",   "550", "securityerror",   false, 0, ""}},
		{{"1003", 1, "5.0.0",   "",    "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.0.0",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "MessageLabs", secretlist, false)
}

