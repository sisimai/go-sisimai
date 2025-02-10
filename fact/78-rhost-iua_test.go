// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _       ___ _   _   _    
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_    |_ _| | | | / \   
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| || | | |/ _ \  
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| || |_| / ___ \ 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|   |___|\___/_/   \_\
import "testing"

func TestRhostIUA(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.0",   "550", "suspend",         false, ""}},
	}; EngineTest(t, "IUA", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "IUA", secretlist, false)
}

