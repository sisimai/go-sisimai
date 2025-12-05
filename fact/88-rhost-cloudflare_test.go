// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _         ____ _                 _  __ _                
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_      / ___| | ___  _   _  __| |/ _| | __ _ _ __ ___ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |   | |/ _ \| | | |/ _` | |_| |/ _` | '__/ _ \
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |___| | (_) | |_| | (_| |  _| | (_| | | |  __/
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|     \____|_|\___/ \__,_|\__,_|_| |_|\__,_|_|  \___|
import "testing"

func TestRhostCloudflare(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "4.3.0",   "421", "systemerror",     false, false, ""}},
	}; EngineTest(t, "Cloudflare", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "Cloudflare", secretlist, false)
}

