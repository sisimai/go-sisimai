// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _            _                           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_         | | __ _ _ __ ___   ___  ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____ _  | |/ _` | '_ ` _ \ / _ \/ __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | (_| | | | | | |  __/\__ \
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \___/ \__,_|_| |_| |_|\___||___/
import "testing"

func TestLhostApacheJames(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.210", "550", "filtered",        false,  true, ""}},
	}; EngineTest(t, "ApacheJames", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.210", "550", "filtered",        false,  true, ""}},
		{{"1002", 1, "5.9.210", "550", "filtered",        false,  true, ""}},
		{{"1003", 1, "5.9.210", "550", "filtered",        false,  true, ""}},
		{{"1004", 1, "5.9.301", "",    "onhold",          false, false, ""}},
		{{"1005", 1, "5.9.301", "",    "onhold",          false, false, ""}},
	}; EngineTest(t, "ApacheJames", secretlist, false)
}

