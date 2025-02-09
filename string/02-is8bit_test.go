// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string

//  _____         _      __   _        _             
// |_   _|__  ___| |_   / /__| |_ _ __(_)_ __   __ _ 
//   | |/ _ \/ __| __| / / __| __| '__| | '_ \ / _` |
//   | |  __/\__ \ |_ / /\__ \ |_| |  | | | | | (_| |
//   |_|\___||___/\__/_/ |___/\__|_|  |_|_| |_|\__, |
//                                             |___/ 
import "testing"

func TestIs8Bit(t *testing.T) {
	fn := "sisimai/string.Is8Bit"
	ae := []string{"nekochan", "Suzu", ""}
	je := []string{"ニャーン", "道綱"}
	cx := 0

	for _, e := range ae { cx++; if Is8Bit(&e) == true  { t.Errorf("%s(%s) returns true", fn, e) } }
	for _, e := range je { cx++; if Is8Bit(&e) == false { t.Errorf("%s(%s) returns false", fn, e) } }

	t.Logf("The number of tests = %d", cx)
}

