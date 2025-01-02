// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _          _                          _     
// |_   _|__  ___| |_   / / | |__   ___  ___| |_       / \   _ __ ___   __ ___   _(_)___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____ / _ \ | '_ ` _ \ / _` \ \ / / / __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____/ ___ \| | | | | | (_| |\ V /| \__ \
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   /_/   \_\_| |_| |_|\__,_| \_/ |_|___/
import "testing"

func TestLhostAmavis(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "5.7.0",   "554", "spamdetected",    false, ""}},
	}
	EngineTest(t, "Amavis", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.7.0",   "554", "spamdetected",    false, ""}},
	}; EngineTest(t, "Amavis", secretlist, false)
}


