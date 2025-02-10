// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        ____                                      _       
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     | __ )  __ _ _ __ _ __ __ _  ___ _   _  __| | __ _ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____|  _ \ / _` | '__| '__/ _` |/ __| | | |/ _` |/ _` |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_) | (_| | |  | | | (_| | (__| |_| | (_| | (_| |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |____/ \__,_|_|  |_|  \__,_|\___|\__,_|\__,_|\__,_|
import "testing"

func TestLhostBarracuda(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"02",   1, "5.7.1",   "550", "spamdetected",    false, ""}},
	}; EngineTest(t, "Barracuda", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001" ,1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1002" ,1, "5.7.1",   "550", "spamdetected",    false, ""}},
	}; EngineTest(t, "Barracuda", secretlist, false)
}

