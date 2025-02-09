// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____         _      __                              
// |_   _|__  ___| |_   / / __ ___  __ _ ___  ___  _ __  
//   | |/ _ \/ __| __| / / '__/ _ \/ _` / __|/ _ \| '_ \ 
//   | |  __/\__ \ |_ / /| | |  __/ (_| \__ \ (_) | | | |
//   |_|\___||___/\__/_/ |_|  \___|\__,_|___/\___/|_| |_|
import "testing"

func TestFind(t *testing.T) {
	fn := "sisimai/reason.Find"
	cx := 0
	cv := Find(nil)

	cx++; if cv != "" { t.Errorf("%s(nil) returns %s", fn, cv) }
	t.Logf("The number of tests = %d", cx)
}

func TestAnotherOne(t *testing.T) {
	fn := "sisimai/reason.anotherone"
	cx := 0
	cv := anotherone(nil)

	cx++; if cv != "" { t.Errorf("%s(nil) returns %s", fn, cv) }
	t.Logf("The number of tests = %d", cx)
}

