// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _     __     __        _                
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   \ \   / /__ _ __(_)_______  _ __  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|___\ \ / / _ \ '__| |_  / _ \| '_ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____\ V /  __/ |  | |/ / (_) | | | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \_/ \___|_|  |_/___\___/|_| |_|
import "testing"

func TestLhostVerizon(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"02",   1, "5.0.911", "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Verizon", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.911", "",    "userunknown",      true, ""}},
		{{"1002", 1, "5.0.911", "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Verizon", secretlist, false)
}

