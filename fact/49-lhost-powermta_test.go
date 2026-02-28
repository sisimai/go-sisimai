// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        ____                        __  __ _____  _    
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  _ \ _____      _____ _ __|  \/  |_   _|/ \   
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |_) / _ \ \ /\ / / _ \ '__| |\/| | | | / _ \  
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____|  __/ (_) \ V  V /  __/ |  | |  | | | |/ ___ \ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|   \___/ \_/\_/ \___|_|  |_|  |_| |_/_/   \_\
//                                                                                                   
import "testing"

func TestLhostPowerMTA(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.2.1",   "550", "suspend",         false, 1, ""}},
		{{"02",   1, "5.0.0",   "554", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.2.1",   "550", "suspend",         false, 1, ""}},
	}; EngineTest(t, "PowerMTA", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.0.0",   "554", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.2.1",   "550", "suspend",         false, 1, ""}},
	}; EngineTest(t, "PowerMTA", secretlist, false)
}

