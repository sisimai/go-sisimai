// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        ____               _       _             ____  _____ ____  
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  _ \ ___  ___ ___(_)_   _(_)_ __   __ _/ ___|| ____/ ___| 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |_) / _ \/ __/ _ \ \ \ / / | '_ \ / _` \___ \|  _| \___ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____|  _ <  __/ (_|  __/ |\ V /| | | | | (_| |___) | |___ ___) |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_| \_\___|\___\___|_| \_/ |_|_| |_|\__, |____/|_____|____/ 
//                                                                                       |___/                   
import "testing"

func TestLhostReceivingSES(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"02",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "4.0.0",   "450", "onhold",          false, ""}},
		{{"04",   1, "5.2.2",   "552", "mailboxfull",     false, ""}},
		{{"05",   1, "5.3.4",   "552", "mesgtoobig",      false, ""}},
		{{"06",   1, "5.6.1",   "500", "spamdetected",    false, ""}},
		{{"07",   1, "5.2.0",   "550", "filtered",        false, ""}},
		{{"08",   1, "5.2.3",   "552", "exceedlimit",     false, ""}},
	}; EngineTest(t, "ReceivingSES", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.2.3",   "552", "exceedlimit",     false, ""}},
	}; EngineTest(t, "ReceivingSES", secretlist, false)
}

