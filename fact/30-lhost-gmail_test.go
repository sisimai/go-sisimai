// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"03",   1, "5.7.0",   "554", "filtered",        false,  true, ""}},
		{{"04",   1, "5.7.1",   "554", "blocked",         false, false, ""}},
		{{"05",   1, "5.7.1",   "550", "securityerror",   false, false, ""}},
		{{"06",   1, "4.2.2",   "450", "mailboxfull",     false, false, ""}},
		{{"07",   1, "5.0.976", "500", "failedstarttls",  false, false, ""}},
		{{"08",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"09",   1, "4.0.947", "",    "expired",         false, false, ""}},
		{{"10",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"11",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"15",   1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"16",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"17",   1, "4.0.947", "",    "expired",         false, false, ""}},
		{{"18",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"19",   1, "5.0.922", "",    "mailboxfull",     false, false, ""}},
	}; EngineTest(t, "Gmail", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1002", 1, "5.2.1",   "550", "suspend",         false,  true, ""}},
		{{"1003", 1, "4.0.947", "",    "expired",         false, false, ""}},
		{{"1004", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1005", 1, "4.0.947", "",    "expired",         false, false, ""}},
		{{"1006", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1007", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1008", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1009", 1, "4.0.947", "",    "expired",         false, false, ""}},
		{{"1010", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1011", 1, "4.2.2",   "450", "mailboxfull",     false, false, ""}},
		{{"1012", 1, "4.0.947", "",    "expired",         false, false, ""}},
		{{"1013", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1014", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1015", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1016", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1017", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1018", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1019", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1020", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1021", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1022", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1023", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1024", 1, "5.0.0",   "553", "blocked",         false, false, ""}},
		{{"1025", 1, "5.7.0",   "554", "filtered",        false,  true, ""}},
		{{"1026", 1, "5.0.910", "550", "filtered",        false,  true, ""}},
		{{"1027", 1, "5.7.1",   "550", "securityerror",   false, false, ""}},
		{{"1028", 1, "5.0.976", "500", "failedstarttls",  false, false, ""}},
		{{"1029", 1, "5.0.901", "",    "onhold",          false, false, ""}},
		{{"1030", 1, "5.7.1",   "554", "blocked",         false, false, ""}},
		{{"1031", 1, "5.7.1",   "550", "blocked",         false, false, ""}},
		{{"1032", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1033", 1, "5.0.971", "",    "blocked",         false, false, ""}},
		{{"1034", 1, "4.0.947", "",    "expired",         false, false, ""}},
		{{"1035", 1, "4.0.971", "",    "blocked",         false, false, ""}},
		{{"1036", 1, "4.0.971", "",    "blocked",         false, false, ""}},
		{{"1037", 1, "5.0.971", "",    "blocked",         false, false, ""}},
		{{"1038", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1039", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1040", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1041", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1042", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1043", 1, "5.0.911", "550", "userunknown",      true,  true, ""}},
		{{"1044", 1, "5.0.972", "",    "policyviolation", false, false, ""}},
		{{"1045", 1, "5.0.947", "",    "expired",         false, false, ""}},
		{{"1046", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1047", 1, "5.1.1",   "550", "userunknown",      true,  true, ""},
		 {"1047", 2, "5.1.1",   "550", "userunknown",      true,  true, ""},
		 {"1047", 3, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1048", 1, "5.0.922", "",    "mailboxfull",     false, false, ""}},
	}; EngineTest(t, "Gmail", secretlist, false)
}

