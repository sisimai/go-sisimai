// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____         _      __                              
// |_   _|__  ___| |_   / / __ ___  __ _ ___  ___  _ __  
//   | |/ _ \/ __| __| / / '__/ _ \/ _` / __|/ _ \| '_ \ 
//   | |  __/\__ \ |_ / /| | |  __/ (_| \__ \ (_) | | | |
//   |_|\___||___/\__/_/ |_|  \___|\__,_|___/\___/|_| |_|
import "testing"
import "slices"
import "libsisimai.org/sisimai/v5/eb"

var ae = []string{
	eb.ReAUTH, eb.ReREPU, eb.ReBLOC, eb.ReBODY, eb.ReTIME, eb.ReTTLS, eb.ReFILT, eb.ReFULL, eb.ReUNIX,
	eb.ReSIZE, eb.ReINET, eb.ReNRFC, eb.RePOLI, eb.ReRELA, eb.ReFROM, eb.ReQPTR, eb.ReSECU, eb.ReSPAM,
	eb.ReRATE, eb.ReSUPP, eb.ReQUIT, eb.ReCOMM, eb.ReSYSE, eb.ReDISK, eb.ReEXEC, eb.Re___0, eb.Re___1,
	eb.ReHOST, eb.ReUSER, eb.ReMOVE, eb.Re00MX, eb.ReSENT, eb.ReAWAY, eb.ReFEED,
}

func TestAvailables(t *testing.T) {
	fn := "reason.Availables"
	cx := 0
	cv := Availables

	cx++; if len(cv) ==  0 { t.Errorf("%s is empty", fn) }
	cx++; if len(cv) != 34 { t.Errorf("%s includes invalid elements: %d", fn, len(cv)) }
	for e := range cv {
		cx++; if e == ""     { t.Errorf("%s returned an empty key", fn) }
		cx++; if cv[e] == "" { t.Errorf("%s[%s] is empty", fn, cv[e]) }
		cx++; if slices.Contains(ae, e) == false {
			t.Errorf("%s() returns invalid reason name: %s", fn, e)
		}
		cx++; if ProbesInto[e](nil) == true { t.Errorf("ProbesInto[%s](nil) returns true", e) }
		cx++; if IncludedIn[e](" ") == true { t.Errorf("IncludedIn[%s](' ') returns true", e) }
	}

	t.Logf("The number of tests = %d", cx)
}

func TestIsExplicit(t *testing.T) {
	fn := "reason.IsExplicit"
	cx := 0

	for _, e := range ae {
		if e == eb.Re___1 || e == eb.Re___0 { continue }
		cx++; if cv := IsExplicit(e); cv == false { t.Errorf("%s(%s) returns false", fn, e) }
	}
	cx++; if IsExplicit("")        == true { t.Errorf("%s() returns true", fn) }
	cx++; if IsExplicit(eb.Re___1) == true { t.Errorf("%s(OnHold) returns true", fn) }
	cx++; if IsExplicit(eb.Re___0) == true { t.Errorf("%s(Undefined) returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}

func TestShouldBeRetried(t *testing.T) {
	fn := "reason.ShouldBeRetried"
	re := []string{eb.Re___0, eb.Re___1, eb.ReSYSE, eb.ReSECU, eb.ReTIME, eb.ReINET, eb.ReHOST, eb.ReUSER}
	cx := 0

	for _, e := range re {
		cx++; if ShouldBeRetried(e) == false  { t.Errorf("%s(%s) returns false", fn, e) }
	}
	for _, e := range ae {
		if slices.Contains(re, e) { continue }
		cx++; if ShouldBeRetried(e) == true   { t.Errorf("%s(%s) returns true",  fn, e) }
	}
	cx++; if ShouldBeRetried("")     == false { t.Errorf("%s('') returns false", fn) }
	cx++; if ShouldBeRetried("neko") == true  { t.Errorf("%s(neko) returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}

