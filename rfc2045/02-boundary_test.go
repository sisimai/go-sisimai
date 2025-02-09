// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045

//  _____         _      ______  _____ ____ ____   ___  _  _  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___ \ / _ \| || || ___| 
//   | |/ _ \/ __| __| / /| |_) | |_ | |     __) | | | | || ||___ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ / __/| |_| |__   _|__) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_____|\___/   |_||____/ 
import "testing"
import "strings"

func TestBoundary(t *testing.T) {
	fn := "sisimai/rfc2045.Boundary"
	cx := 0
	ae := []string{
		`multipart/report; boundary="000000000000288e03056cdb87e8"; report-type=delivery-status`,
		`multipart/alternative; boundary="000000000000ad35ae059fee8fb4"`,
		`multipart/mixed; Boundary="0000000000009d969b05f5a54484"`,
		`multipart/related; BOUNDARY="000000000000a5a7020593e667ad"`,
	}

	for _, e := range ae {
		for _, f := range []int{-2, -1, 0, 1, 2} {
			cv := Boundary(e, f)
			cx++; if strings.Contains(cv, "000000") == false {
				t.Errorf("%s(%s, %d) returns %s", fn, e, f, cv)
			}

			if f < 0 {
				cx++; if strings.HasPrefix(cv, "000000") == false {
					t.Errorf("%s(%s, %d) returns: %s", fn, e, f, cv)
				}
			} else {
				cx++; if strings.HasPrefix(cv, "--") == false {
					t.Errorf("%s(%s, %d) does not start with --: %s", fn, e, f, cv)
				}
				if f > 0 {
					cx++; if strings.HasSuffix(cv, "--") == false {
						t.Errorf("%s(%s, %d) does not end with --: %s", fn, e, f, cv)
					}
				}
			}
		}
	}
	if cv := Boundary("", 22); cv != "" { t.Errorf("%s('', 22) returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

