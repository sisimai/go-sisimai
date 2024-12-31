// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//  _____         _      __        _     _                     _____ _           _ 
// |_   _|__  ___| |_   / /_ _  __| | __| |_ __ ___  ___ ___  |  ___(_)_ __   __| |
//   | |/ _ \/ __| __| / / _` |/ _` |/ _` | '__/ _ \/ __/ __| | |_  | | '_ \ / _` |
//   | |  __/\__ \ |_ / / (_| | (_| | (_| | | |  __/\__ \__ \_|  _| | | | | | (_| |
//   |_|\___||___/\__/_/ \__,_|\__,_|\__,_|_|  \___||___/___(_)_|   |_|_| |_|\__,_|
import "testing"

func TestFind(t *testing.T) {
	fn := "sisimai/address.Find()"
	cx := 0

	for _, e := range TestEmailAddrs {
		//if e.testname != "D" { continue }
		t.Run(e.testname, func(t *testing.T) {
			cv := Find(e.argument)
			if len(cv) != 3        { t.Errorf("[%6d]: %s does not have 3 elements not (%d)", cx, fn, len(cv))      }; cx++
			if cv[0] != e.expected { t.Errorf("[%6d]: %s [0:address] is (%s) not (%s)", cx, fn, cv[0], e.expected) }; cx++
			if cv[1] != e.displays { t.Errorf("[%6d]: %s [1:display] is (%s) not (%s)", cx, fn, cv[1], e.displays) }; cx++
			if cv[2] != e.comments { t.Errorf("[%6d]: %s [2:comment] is (%s) not (%s)", cx, fn, cv[2], e.comments) }; cx++
		})
	}
	for _, e := range TestPostmaster {
		t.Run("", func(t *testing.T) {
			cv := Find(e)
			if len(cv) != 3 { t.Errorf("[%6d]: %s does not have 3 elements not (%d)", cx, fn, len(cv)) }; cx++
			if cv[0] == ""  { t.Errorf("[%6d]: %s [0:address] is empty", cx, fn)                       }; cx++
		})
	}

	for _, e := range TestNotAnEmail {
		t.Run("", func(t *testing.T) {
			cv := Find(e)
			if len(cv) != 3 { t.Errorf("[%6d]: %s does not have 3 elements not (%d)", cx, fn, len(cv)) }; cx++
			if cv[0] != ""  { t.Errorf("[%6d]: %s [0:address] is (%s) not empty", cx, fn, cv[0])       }; cx++
		})
	}
	t.Logf("The number of tests = %d", cx)
}

