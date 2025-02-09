// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      _______          _   
// |_   _|__  ___| |_   / /  ___|_ _  ___| |_ 
//   | |/ _ \/ __| __| / /| |_ / _` |/ __| __|
//   | |  __/\__ \ |_ / / |  _| (_| | (__| |_ 
//   |_|\___||___/\__/_/  |_|  \__,_|\___|\__|
import "testing"

func TestRise(t *testing.T) {
	fn := "sisimai/fact.Rise"
	fs := "sis.NotDecoded"
	cx := 0

	cv, ce := Rise(nil, "", nil)
	cx++; if len(cv)  > 0 { t.Errorf("%s(nil) returns %v", fn, cv) }
	cx++; if len(ce) == 0 { t.Errorf("%s(nil) did not return errors", fn) }
	cx++; if len(ce) != 1 { t.Errorf("%s(nil) return errors: %v", fn, ce) }

	for _, e := range ce {
		cx++; if e.EmailFile != ""    { t.Errorf("%s.EmailFile is not empty: %s", fs, e.EmailFile) }
		cx++; if e.Email("cat") == "" { t.Errorf("%s.Email(cat) returns empty", fs) }
		cx++; if e.EmailFile != "cat" { t.Errorf("%s.EmailFile is not `cat`: %s", fs, e.EmailFile) }
		cx++; if e.BecauseOf == ""    { t.Errorf("%s.BecauseOf is empty", fs) }
		cx++; if e.CalledOff == false { t.Errorf("%s.CalledOff returns false", fs) }
		cx++; if e.DecodedBy != ""    { t.Errorf("%s.DecodedBy is not empty: %s", fs, e.DecodedBy) }
		cx++; if e.WhoCalled == ""    { t.Errorf("%s.WhoCalled is empty", fs) }
		cx++; if e.Timestamp.IsZero() { t.Errorf("%s.Timestamp is nil", fs) }
		cx++; if e.Error()   == ""    { t.Errorf("%s.Error() returns empty", fs) }
		cx++; if e.Label()   == ""    { t.Errorf("%s.Label() returns empty", fs) }
	}

	t.Logf("The number of tests = %d", cx)
}

