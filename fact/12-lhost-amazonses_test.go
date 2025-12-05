// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "5.7.1",   "550", "securityerror",   false, false, ""}},
		{{"02",   1, "5.3.0",   "550", "filtered",        false,  true, ""}},
		{{"03",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"05",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"06",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"07",   1, "5.7.6",   "550", "securityerror",   false, false, ""}},
		{{"08",   1, "5.7.9",   "550", "securityerror",   false, false, ""}},
		{{"09",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"10",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"11",   1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"12",   1, "2.6.0",   "250", "delivered",       false, false, ""}},
		{{"13",   1, "2.6.0",   "250", "delivered",       false, false, ""}},
		{{"14",   1, "5.7.1",   "554", "blocked",         false, false, ""}},
		{{"15",   1, "5.7.1",   "554", "blocked",         false, false, ""}},
		{{"16",   1, "5.7.1",   "521", "blocked",         false, false, ""}},
		{{"17",   1, "4.4.2",   "421", "expired",         false, false, ""}},
		{{"18",   1, "5.4.4",   "550", "hostunknown",      true,  true, ""}},
		{{"19",   1, "5.7.1",   "550", "suspend",         false,  true, ""}},
		{{"20",   1, "5.2.1",   "550", "suspend",         false,  true, ""}},
		{{"21",   1, "5.7.1",   "554", "norelaying",      false,  true, ""}},
	}; EngineTest(t, "AmazonSES", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1002", 1, "5.2.1",   "550", "filtered",        false,  true, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1005", 1, "5.7.1",   "550", "securityerror",   false, false, ""}},
		{{"1006", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1007", 1, "5.4.7",   "",    "expired",         false, false, ""}},
		{{"1008", 1, "5.1.2",   "",    "hostunknown",      true,  true, ""}},
		{{"1009", 1, "5.1.0",   "550", "userunknown",      true,  true, ""}},
		{{"1010", 1, "5.1.0",   "550", "userunknown",      true,  true, ""}},
		{{"1011", 1, "5.1.0",   "550", "userunknown",      true,  true, ""}},
		{{"1012", 1, "5.1.0",   "550", "userunknown",      true,  true, ""}},
		{{"1013", 1, "5.1.0",   "550", "userunknown",      true,  true, ""}},
		{{"1014", 1, "5.3.0",   "550", "filtered",        false,  true, ""}},
		{{"1015", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1016", 1, "",        "",    "feedback",        false,  true, "abuse"}},
		{{"1017", 1, "2.6.0",   "250", "delivered",       false, false, ""}},
		{{"1018", 1, "2.6.0",   "250", "delivered",       false, false, ""}},
		{{"1019", 1, "5.7.1",   "554", "blocked",         false, false, ""}},
		{{"1020", 1, "4.4.2",   "421", "expired",         false, false, ""}},
		{{"1021", 1, "5.4.4",   "550", "hostunknown",      true,  true, ""}},
		{{"1022", 1, "5.5.1",   "550", "blocked",         false, false, ""}},
		{{"1023", 1, "5.7.1",   "550", "suspend",         false,  true, ""}},
		{{"1024", 1, "5.4.1",   "550", "userunknown",      true,  true, ""}},
		{{"1025", 1, "5.2.1",   "550", "suspend",         false,  true, ""}},
		{{"1026", 1, "5.7.1",   "554", "norelaying",      false,  true, ""}},
		{{"1027", 1, "5.2.2",   "552", "mailboxfull",     false,  true, ""}},
		{{"1028", 1, "5.4.7",   "",    "expired",         false, false, ""}},
		{{"1029", 1, "5.1.0",   "550", "userunknown",      true,  true, ""}},
		{{"1030", 1, "2.6.0",   "250", "delivered",       false, false, ""}},
		{{"1031", 1, "2.6.0",   "250", "delivered",       false, false, ""}},
	}; EngineTest(t, "AmazonSES", secretlist, false)
}

