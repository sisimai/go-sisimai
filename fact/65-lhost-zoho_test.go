// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _       _____     _           
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    |__  /___ | |__   ___  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____ / // _ \| '_ \ / _ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____/ /| (_) | | | | (_) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   /____\___/|_| |_|\___/ 
import "testing"

func TestLhostZoho(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"02",   1, "5.2.1",   "550", "filtered",        false,  true, ""},
		 {"02",   2, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"03",   1, "5.9.210", "550", "filtered",        false,  true, ""}},
		{{"04",   1, "4.9.340", "421", "expired",         false, false, ""}},
		{{"05",   1, "4.9.340", "421", "expired",         false, false, ""}},
	}; EngineTest(t, "Zoho", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1002", 1, "5.2.1",   "550", "filtered",        false,  true, ""},
		 {"1002", 2, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1003", 1, "5.9.210", "550", "filtered",        false,  true, ""}},
		{{"1004", 1, "4.9.340", "421", "expired",         false, false, ""}},
	}; EngineTest(t, "Zoho", secretlist, false)
}

