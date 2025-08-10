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

func TestDecodeB(t *testing.T) {
	fn := "rfc2045.DecodeB"
	cx := 0
	be := []string{"44OL44Oj44O844Oz", "6YGT57ax"}
	jp := []string{"ニャーン", "道綱"}

	for j, e := range be {
		for _, f := range []string{"", "utf-8"} {
			cv, ce := DecodeB(e)
			cx++; if cv != jp[j] { t.Errorf("%s(%s, %s) returns %s", fn, e, f, cv) }
			cx++; if ce != nil   { t.Errorf("%s(%s, %s) returns error: %s", fn, e, f, ce) }
		}
	}
	if cv, _ := DecodeB(""); cv != "" { t.Errorf("%s('') returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

func TestDecodeQ(t *testing.T) {
	fn := "rfc2045.DecodeQ"
	cx := 0
	be := []string{"=E3=83=8B=E3=83=A3=E3=83=BC=E3=83=B3", "=E9=81=93=E7=B6=B1"}
	jp := []string{"ニャーン", "道綱"}
	cw := `I will be traveling for work on July 10-31.  During that time I will have i=
ntermittent access to email and phone, and I will respond to your message a=
s promptly as possible.

Please contact our Client Service Support Team (information below) if you n=
eed immediate assistance on regular account matters, or contact my colleagu=
e Neko Nyaan (neko@example.org; +0-000-000-0000) for all other needs.`

	for j, e := range be {
		cv, ce := DecodeQ(e)
		cx++; if cv != jp[j] { t.Errorf("%s(%s) returns %s", fn, e, cv) }
		cx++; if ce != nil   { t.Errorf("%s(%s) returns error: %s", fn, e, ce) }
	}
	if cv, _ := DecodeQ(cw); strings.Contains(cv, "=\n") { t.Errorf("%s(%s) returns %s", fn, cw[:10], cv) }
	if cv, _ := DecodeQ(""); cv != ""                    { t.Errorf("%s('') returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

