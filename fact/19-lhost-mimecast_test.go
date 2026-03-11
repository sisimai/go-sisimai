// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        __  __ _                              _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  \/  (_)_ __ ___   ___  ___ __ _ ___| |_ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |\/| | | '_ ` _ \ / _ \/ __/ _` / __| __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |  | | | | | | | |  __/ (_| (_| \__ \ |_ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  |_|_|_| |_| |_|\___|\___\__,_|___/\__|
import "testing"

func TestLhostMimecast(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.4.1",   "",    "userunknown",      true, 1, ""}},
		{{"02",   1, "5.7.54",  "550", "norelaying",      false, 1, ""}},
	}; EngineTest(t, "Mimecast", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.4.1",   "",    "userunknown",      true, 1, ""}},
		{{"1002", 1, "4.4.4",   "",    "networkerror",    false, 0, ""}},
		{{"1003", 1, "5.1.1",   "",    "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.4.14",  "554", "networkerror",    false, 0, ""}},
		{{"1005", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.7.54",  "550", "norelaying",      false, 1, ""}},
		{{"1007", 1, "5.7.1",   "550", "blocked",         false, 0, ""}},
		{{"1008", 1, "5.2.1",   "550", "suspend",         false, 1, ""}},
	}; EngineTest(t, "Mimecast", secretlist, false)
}

