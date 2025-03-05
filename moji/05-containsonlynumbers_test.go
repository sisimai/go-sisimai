// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package moji

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   ___ (_|_)
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \| | |
//   | |  __/\__ \ |_ / /| | | | | | (_) | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\___// |_|
//                                     |__/   
import "testing"

func TestContainsOnlyNumbers(t *testing.T) {
	fn := "moji.ContainsOnlyNumbers"
	cx := 0
	et := []string{"1", "23", "456", "78910"}
	ef := []string{"A", "B1", "C12", "34D5E"}

	for _, e := range et {
		cx++; if cv := ContainsOnlyNumbers(e); cv == false { t.Errorf("%s(%s) returns false", fn, e) }
	}
	for _, e := range ef {
		cx++; if cv := ContainsOnlyNumbers(e); cv == true  { t.Errorf("%s(%s) returns true",  fn, e) }
	}
	cx++; if cv := ContainsOnlyNumbers("");    cv == true  { t.Errorf("%s(%s) returns true",  fn, "") }

	t.Logf("The number of tests = %d", cx)
}

