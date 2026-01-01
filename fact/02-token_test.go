// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      _______          _   
// |_   _|__  ___| |_   / /  ___|_ _  ___| |_ 
//   | |/ _ \/ __| __| / /| |_ / _` |/ __| __|
//   | |  __/\__ \ |_ / / |  _| (_| | (__| |_ 
//   |_|\___||___/\__/_/  |_|  \__,_|\___|\__|
import "testing"

func TestToken(t *testing.T) {
	fn := "fact.token"
	es := "envelope-sender@example.jp"
	er := "envelope-recipient@example.org"
	to := "239aa35547613b2fa94f40c7f35f4394e99fdd88"
	cx := 0

	cx++; if token(es, er, 1) != to { t.Errorf("%s(%s, %s, 1) returns %s", fn, es, er, token(es,er, 1)) }
	cx++; if token("", "", 0) != "" { t.Errorf("%s('', '', 0) returns %s", fn, token("", "", 0)) }
	cx++; if token(es, "", 0) != "" { t.Errorf("%s(%s, '', 0) returns %s", fn, es, token("", "", 0)) }
	cx++; if token("", er, 0) != "" { t.Errorf("%s('', %s, 0) returns %s", fn, er, token("", "", 0)) }

	t.Logf("The number of tests = %d", cx)
}

