// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _      ___                                   ____  _____ ____  
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   / / \   _ __ ___   __ _ _______  _ __ / ___|| ____/ ___| 
//   | |/ _ \/ __| __| / /| | "_ \ / _ \/ __| __| / / _ \ | "_ ` _ \ / _` |_  / _ \| "_ \\___ \|  _| \___ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ / / ___ \| | | | | | (_| |/ / (_) | | | |___) | |___ ___) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__/_/_/   \_\_| |_| |_|\__,_/___\___/|_| |_|____/|_____|____/ 
import "testing"

func TestLhostAmazonSES(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.7.1",   "550", "securityerror",   false, 0, ""}},
		{{"02",   1, "5.3.0",   "550", "filtered",        false, 1, ""}},
		{{"03",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"05",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"06",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"07",   1, "5.7.6",   "550", "securityerror",   false, 0, ""}},
		{{"08",   1, "5.7.9",   "550", "securityerror",   false, 0, ""}},
		{{"09",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"10",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"11",   1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"12",   1, "2.6.0",   "250", "delivered",       false, 0, ""}},
		{{"13",   1, "2.6.0",   "250", "delivered",       false, 0, ""}},
		{{"14",   1, "5.7.1",   "554", "blocked",         false, 0, ""}},
		{{"15",   1, "5.7.1",   "554", "blocked",         false, 0, ""}},
		{{"16",   1, "5.7.1",   "521", "blocked",         false, 0, ""}},
		{{"17",   1, "4.4.2",   "421", "expired",         false, 0, ""}},
		{{"18",   1, "5.4.4",   "550", "hostunknown",      true, 1, ""}},
		{{"19",   1, "5.7.1",   "550", "suspend",         false, 1, ""}},
		{{"20",   1, "5.2.1",   "550", "suspend",         false, 1, ""}},
		{{"21",   1, "5.7.1",   "554", "norelaying",      false, 1, ""}},
	}; EngineTest(t, "AmazonSES", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1002", 1, "5.2.1",   "550", "filtered",        false, 1, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1005", 1, "5.7.1",   "550", "securityerror",   false, 0, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1007", 1, "5.4.7",   "",    "expired",         false, 0, ""}},
		{{"1008", 1, "5.1.2",   "",    "hostunknown",      true, 1, ""}},
		{{"1009", 1, "5.1.0",   "550", "userunknown",      true, 1, ""}},
		{{"1010", 1, "5.1.0",   "550", "userunknown",      true, 1, ""}},
		{{"1011", 1, "5.1.0",   "550", "userunknown",      true, 1, ""}},
		{{"1012", 1, "5.1.0",   "550", "userunknown",      true, 1, ""}},
		{{"1013", 1, "5.1.0",   "550", "userunknown",      true, 1, ""}},
		{{"1014", 1, "5.3.0",   "550", "filtered",        false, 1, ""}},
		{{"1015", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1016", 1, "",        "",    "feedback",        false, 1, "abuse"}},
		{{"1017", 1, "2.6.0",   "250", "delivered",       false, 0, ""}},
		{{"1018", 1, "2.6.0",   "250", "delivered",       false, 0, ""}},
		{{"1019", 1, "5.7.1",   "554", "blocked",         false, 0, ""}},
		{{"1020", 1, "4.4.2",   "421", "expired",         false, 0, ""}},
		{{"1021", 1, "5.4.4",   "550", "hostunknown",      true, 1, ""}},
		{{"1022", 1, "5.5.1",   "550", "blocked",         false, 0, ""}},
		{{"1023", 1, "5.7.1",   "550", "suspend",         false, 1, ""}},
		{{"1024", 1, "5.4.1",   "550", "userunknown",      true, 1, ""}},
		{{"1025", 1, "5.2.1",   "550", "suspend",         false, 1, ""}},
		{{"1026", 1, "5.7.1",   "554", "norelaying",      false, 1, ""}},
		{{"1027", 1, "5.2.2",   "552", "mailboxfull",     false, 1, ""}},
		{{"1028", 1, "5.4.7",   "",    "expired",         false, 0, ""}},
		{{"1029", 1, "5.1.0",   "550", "userunknown",      true, 1, ""}},
		{{"1030", 1, "2.6.0",   "250", "delivered",       false, 0, ""}},
		{{"1031", 1, "2.6.0",   "250", "delivered",       false, 0, ""}},
	}; EngineTest(t, "AmazonSES", secretlist, false)
}

