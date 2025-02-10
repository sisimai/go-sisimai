// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _     __        __         _                             
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   \ \      / /__  _ __| | _____ _ __   __ _  ___ ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|___\ \ /\ / / _ \| '__| |/ / __| '_ \ / _` |/ __/ _ \
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____\ V  V / (_) | |  |   <\__ \ |_) | (_| | (_|  __/
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \_/\_/ \___/|_|  |_|\_\___/ .__/ \__,_|\___\___|
//                                                                              |_|                   
import "testing"

func TestLhostGoogleWorkspace(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.918", "",    "rejected",        false, ""}},
	}; EngineTest(t, "GoogleWorkspace", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.918", "",    "rejected",        false, ""}},
	}; EngineTest(t, "GoogleWorkspace", secretlist, false)
}

