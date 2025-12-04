// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      __  ____   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_    \ \/ / /_  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____\  / '_ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____/  \ (_) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   /_/\_\___/ 
import "testing"

func TestLhostX6(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.4.6",   "554", "networkerror",    false, false, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
	}; EngineTest(t, "X6", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1002", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1003", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1004", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1005", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1006", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1007", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1008", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1009", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1010", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1011", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1012", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1013", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1014", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1015", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1016", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1017", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1018", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1019", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1020", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1021", 1, "5.4.6",   "554", "networkerror",    false, false, ""}},
		{{"1022", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1023", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1024", 1, "5.4.6",   "554", "networkerror",    false, false, ""}},
		{{"1025", 1, "5.7.1",   "550", "norelaying",      false,  true, ""}},
		{{"1026", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1027", 1, "5.0.970", "550", "securityerror",   false, false, ""}},
	}; EngineTest(t, "X6", secretlist, false)
}

