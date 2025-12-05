// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _          _                                __        __         _    __  __       _ _ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_       / \   _ __ ___   __ _ _______  _ _\ \      / /__  _ __| | _|  \/  | __ _(_) |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____ / _ \ | '_ ` _ \ / _` |_  / _ \| '_ \ \ /\ / / _ \| '__| |/ / |\/| |/ _` | | |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____/ ___ \| | | | | | (_| |/ / (_) | | | \ V  V / (_) | |  |   <| |  | | (_| | | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|   /_/   \_\_| |_| |_|\__,_/___\___/|_| |_|\_/\_/ \___/|_|  |_|\_\_|  |_|\__,_|_|_|
import "testing"

func TestLhostAmazonWorkMail(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"02",   1, "5.2.1",   "550", "filtered",        false,  true, ""}},
		{{"03",   1, "5.3.5",   "550", "systemerror",     false, false, ""}},
		{{"04",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"05",   1, "4.4.2",   "421", "expired",         false, false, ""}},
		{{"07",   1, "4.4.2",   "421", "expired",         false, false, ""}},
		{{"08",   1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
	}; EngineTest(t, "AmazonWorkMail", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true,  true, ""}},
		{{"1002", 1, "5.2.1",   "550", "filtered",        false,  true, ""}},
		{{"1003", 1, "5.3.5",   "550", "systemerror",     false, false, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
		{{"1005", 1, "4.4.2",   "421", "expired",         false, false, ""}},
		{{"1006", 1, "5.2.2",   "550", "mailboxfull",     false,  true, ""}},
	}; EngineTest(t, "AmazonWorkMail", secretlist, false)
}

