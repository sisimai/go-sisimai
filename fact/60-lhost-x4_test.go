// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      __  ___  _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    \ \/ / || |  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____\  /| || |_ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____/  \|__   _|
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   /_/\_\  |_|  
//                                                               
import "testing"

func TestLhostX4(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
	}; EngineTest(t, "X4", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1002", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1003", 1, "5.1.2",   "",    "hostunknown",      true, 1, ""}},
		{{"1004", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1005", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1006", 1, "5.1.1",   "",    "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.9.221", "550", "suspend",         false, 1, ""}},
		{{"1009", 1, "5.1.1",   "",    "userunknown",      true, 1, ""}},
		{{"1010", 1, "5.1.2",   "",    "hostunknown",      true, 1, ""}},
		{{"1011", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1012", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1013", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1014", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1015", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1016", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1018", 1, "5.1.1",   "",    "userunknown",      true, 1, ""}},
		{{"1019", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1020", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1022", 1, "5.1.1",   "",    "userunknown",      true, 1, ""}},
		{{"1023", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1024", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
		{{"1025", 1, "5.1.1",   "",    "userunknown",      true, 1, ""}},
		{{"1026", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1027", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
	}; EngineTest(t, "X4", secretlist, false)
}

