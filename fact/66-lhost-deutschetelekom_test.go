// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _            _            _            _          _       _      _                   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_       __| | ___ _   _| |_ ___  ___| |__   ___| |_ ___| | ___| | _____  _ __ ___  
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____ / _` |/ _ \ | | | __/ __|/ __| '_ \ / _ \ __/ _ \ |/ _ \ |/ / _ \| '_ ` _ \ 
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____| (_| |  __/ |_| | |_\__ \ (__| | | |  __/ ||  __/ |  __/   < (_) | | | | | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \__,_|\___|\__,_|\__|___/\___|_| |_|\___|\__\___|_|\___|_|\_\___/|_| |_| |_|
import "testing"

func TestLhostDeutscheTelekom(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.2.2",   "552", "mailboxfull",     false, -1, ""}},
		{{"02",   1, "5.2.2",   "552", "mailboxfull",     false, -1, ""},
		 {"02",   2, "5.1.1",   "550", "userunknown",      true, -1, ""}},
		{{"03",   1, "5.9.212", "",    "hostunknown",      true, -1, ""}},
	}; EngineTest(t, "DeutscheTelekom", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"1001", 1, "5.2.2",   "552", "mailboxfull",     false, -1, ""}},
	}; EngineTest(t, "DeutscheTelekom", secretlist, false)
}

