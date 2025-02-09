// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        _____              _                 _    
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  ___|_ _  ___ ___| |__   ___   ___ | | __
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |_ / _` |/ __/ _ \ '_ \ / _ \ / _ \| |/ /
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____|  _| (_| | (_|  __/ |_) | (_) | (_) |   < 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|  \__,_|\___\___|_.__/ \___/ \___/|_|\_\
//                                                                                              
import "testing"

func TestLhostFacebook(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "Facebook", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.1.1",   "550", "userunknown",      true, ""}},
	}; EngineTest(t, "Facebook", secretlist, false)
}

