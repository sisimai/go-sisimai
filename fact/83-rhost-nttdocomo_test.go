// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _        _   _ _____ _____ ____   ___   ____ ___  __  __  ___  
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_     | \ | |_   _|_   _|  _ \ / _ \ / ___/ _ \|  \/  |/ _ \ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____|  \| | | |   | | | | | | | | | |  | | | | |\/| | | | |
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |\  | | |   | | | |_| | |_| | |__| |_| | |  | | |_| |
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|    |_| \_| |_|   |_| |____/ \___/ \____\___/|_|  |_|\___/ 
import "testing"

func TestRhostNTTDOCOMO(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.2.0",   "550", "filtered",        false, 1, ""}},
		{{"02",   1, "5.0.0",   "550", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.0.0",   "550", "userunknown",      true, 1, ""}},
	}; EngineTest(t, "NTTDOCOMO", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "NTTDOCOMO", secretlist, false)
}

