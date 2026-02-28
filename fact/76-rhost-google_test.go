// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _         ____                   _      
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_      / ___| ___   ___   __ _| | ___ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |  _ / _ \ / _ \ / _` | |/ _ \
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |_| | (_) | (_) | (_| | |  __/
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|     \____|\___/ \___/ \__, |_|\___|
//                                                                        |___/        
import "testing"

func TestRhostGoogle(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.2.1",   "550", "suspend",         false, 1, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.7.26",  "550", "authfailure",     false, 0, ""}},
		{{"04",   1, "5.7.26",  "550", "authfailure",     false, 0, ""}},
		{{"05",   1, "5.2.2",   "552", "mailboxfull",     false, 1, ""}},
		{{"06",   1, "5.7.25",  "550", "requireptr",      false, 0, ""}},
		{{"07",   1, "5.2.1",   "550", "suspend",         false, 1, ""}},
		{{"08",   1, "5.7.1",   "550", "notcompliantrfc", false, 0, ""}},
	}; EngineTest(t, "Google", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "Google", secretlist, false)
}

