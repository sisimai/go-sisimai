// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
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
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"02",   1, "5.2.1",   "550", "filtered",        false, 1, ""}},
		{{"03",   1, "5.3.5",   "550", "systemerror",     false, 0, ""}},
		{{"04",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"05",   1, "4.4.2",   "421", "expired",         false, 0, ""}},
		{{"07",   1, "4.4.2",   "421", "expired",         false, 0, ""}},
		{{"08",   1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
	}; EngineTest(t, "AmazonWorkMail", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, 1, ""}},
		{{"1002", 1, "5.2.1",   "550", "filtered",        false, 1, ""}},
		{{"1003", 1, "5.3.5",   "550", "systemerror",     false, 0, ""}},
		{{"1004", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
		{{"1005", 1, "4.4.2",   "421", "expired",         false, 0, ""}},
		{{"1006", 1, "5.2.2",   "550", "mailboxfull",     false, 1, ""}},
	}; EngineTest(t, "AmazonWorkMail", secretlist, false)
}

