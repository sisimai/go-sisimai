// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        _   _       _            
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     | \ | | ___ | |_ ___  ___ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____|  \| |/ _ \| __/ _ \/ __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |\  | (_) | ||  __/\__ \
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_| \_|\___/ \__\___||___/
import "testing"

func TestLhostNotes(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"02",   1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"03",   1, "5.9.213", "",    "userunknown",      true, 1, ""}},
	}; EngineTest(t, "Notes", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"1003", 1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"1004", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1005", 1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"1006", 1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"1007", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1009", 1, "5.9.213", "",    "userunknown",      true, 1, ""}},
		{{"1010", 1, "5.9.341", "",    "networkerror",    false, 0, ""}},
	}; EngineTest(t, "Notes", secretlist, false)
}

