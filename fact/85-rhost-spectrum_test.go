// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _       ____                  _                        
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_    / ___| _ __   ___  ___| |_ _ __ _   _ _ __ ___  
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|___\___ \| '_ \ / _ \/ __| __| '__| | | | '_ ` _ \ 
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____|__) | |_) |  __/ (__| |_| |  | |_| | | | | | |
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|   |____/| .__/ \___|\___|\__|_|   \__,_|_| |_| |_|
//                                                          |_|                                       
import "testing"

func TestRhostSpectrum(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
		{{"01",   1, "5.1.0",   "550", "ratelimited",     false, 0, ""}},
	}; EngineTest(t, "Spectrum", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, Toxic, AnotherOne
	}; EngineTest(t, "Spectrum", secretlist, false)
}

