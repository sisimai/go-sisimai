// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _____         _      ___ _               _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|
import "testing"

func TestIndicators(t *testing.T) {
	cx := 0

	cx++; if HereIsDeliveryStatus != 2 { t.Errorf("HereIsDeliveryStatus = %d", HereIsDeliveryStatus) }
	cx++; if HereIsMessageRFC822  != 4 { t.Errorf("HereIsMessageRFC822  = %d", HereIsMessageRFC822)  }
	t.Logf("The number of tests = %d", cx)
}

