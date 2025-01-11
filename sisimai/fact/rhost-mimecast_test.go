// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      __    _               _        __  __ _                              _   
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_     |  \/  (_)_ __ ___   ___  ___ __ _ ___| |_ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|____| |\/| | | '_ ` _ \ / _ \/ __/ _` / __| __|
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ ||_____| |  | | | | | | | |  __/ (_| (_| \__ \ |_ 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|    |_|  |_|_|_| |_| |_|\___|\___\__,_|___/\__|
import "testing"

func TestRhostMimecast(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.0.0",   "554", "policyviolation", false, ""}},
		{{"02",   1, "5.0.0",   "554", "virusdetected",   false, ""}},
	}; EngineTest(t, "Mimecast", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
	}; EngineTest(t, "Mimecast", secretlist, false)
}

