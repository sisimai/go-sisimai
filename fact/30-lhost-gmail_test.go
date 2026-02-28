// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _         ____                 _ _ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_      / ___|_ __ ___   __ _(_) |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |  _| '_ ` _ \ / _` | | |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| |_| | | | | | | (_| | | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \____|_| |_| |_|\__,_|_|_|
import "testing"

func TestLhostGmail(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"03",   1, "5.7.0",   "554", "filtered",        false, 1, ""}},
		{{"04",   1, "5.7.1",   "554", "blocked",         false, 0, ""}},
		{{"05",   1, "5.7.1",   "550", "securityerror",   false, 0, ""}},
		{{"06",   1, "4.2.2",   "450", "mailboxfull",     false, 0, ""}},
		{{"07",   1, "5.9.350", "500", "failedstarttls",  false, 0, ""}},
		{{"08",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"09",   1, "4.9.340", "",    "expired",         false, 0, ""}},
		{{"10",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"11",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"15",   1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"16",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"17",   1, "4.9.340", "",    "expired",         false, 0, ""}},
		{{"18",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"19",   1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
	}; EngineTest(t, "Gmail", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1002", 1, "5.2.1",   "550", "suspend",         false, 1, ""}},
		{{"1003", 1, "4.9.340", "",    "expired",         false, 0, ""}},
		{{"1004", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1005", 1, "4.9.340", "",    "expired",         false, 0, ""}},
		{{"1006", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1007", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1008", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1009", 1, "4.9.340", "",    "expired",         false, 0, ""}},
		{{"1010", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1011", 1, "4.2.2",   "450", "mailboxfull",     false, 0, ""}},
		{{"1012", 1, "4.9.340", "",    "expired",         false, 0, ""}},
		{{"1013", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1014", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1015", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1016", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1017", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1018", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1019", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1020", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1021", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1022", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1023", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1024", 1, "5.0.0",   "553", "blocked",         false, 0, ""}},
		{{"1025", 1, "5.7.0",   "554", "filtered",        false, 1, ""}},
		{{"1026", 1, "5.9.210", "550", "filtered",        false, 1, ""}},
		{{"1027", 1, "5.7.1",   "550", "securityerror",   false, 0, ""}},
		{{"1028", 1, "5.9.350", "500", "failedstarttls",  false, 0, ""}},
		{{"1029", 1, "5.9.301", "",    "onhold",          false, 0, ""}},
		{{"1030", 1, "5.7.1",   "554", "blocked",         false, 0, ""}},
		{{"1031", 1, "5.7.1",   "550", "blocked",         false, 0, ""}},
		{{"1032", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1033", 1, "5.9.134", "",    "blocked",         false, 0, ""}},
		{{"1034", 1, "4.9.340", "",    "expired",         false, 0, ""}},
		{{"1035", 1, "4.9.134", "",    "blocked",         false, 0, ""}},
		{{"1036", 1, "4.9.134", "",    "blocked",         false, 0, ""}},
		{{"1037", 1, "5.9.134", "",    "blocked",         false, 0, ""}},
		{{"1038", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1039", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1040", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1041", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1042", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1043", 1, "5.9.213", "550", "userunknown",      true, 1, ""}},
		{{"1044", 1, "5.9.131", "",    "ratelimited",     false, 0, ""}},
		{{"1045", 1, "5.9.340", "",    "expired",         false, 0, ""}},
		{{"1046", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1047", 1, "5.1.1",   "550", "userunknown",      true, 1, ""},
		 {"1047", 2, "5.1.1",   "550", "userunknown",      true, 1, ""},
		 {"1047", 3, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1048", 1, "5.9.220", "",    "mailboxfull",     false, 0, ""}},
	}; EngineTest(t, "Gmail", secretlist, false)
}

