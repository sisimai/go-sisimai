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

var Ae = []string{
	"=?UTF-8?B?6YGT57ax44OL44Oj44O8?=",
	"=?ISO-2022-JP?B?GyRCRjs5SyVLJWMhPBsoQg==?=",
	"=?EUC-JP?B?xru5y6XLpeOhvA==?=",
	"=?SHIFT_JIS?B?k7mNaoNqg4OBWw==?=",
	"=?UTF-8?Q?=E9=81=93=E7=B6=B1=E3=83=8B=E3=83=A3=E3=83=BC",
	"=?ISO-2022-JP?Q?=1B=24BF=3B9K=25K=25c=21=3C=1B=28B",
	"=?EUC-JP?Q?=C6=BB=B9=CB=A5=CB=A5=E3=A1=BC",
	"=?SHIFT_JIS?Q?=93=B9=8Dj=83j=83=83=81=5B",
}
var Je = []string{
	"ASCII TEXT",
	"道綱ニャー",
	"ニュースレター",
}

func TestIsEncoded(t *testing.T) {
	fn := "sisimai/rfc2045.IsEncoded"
	cx := 0

	for _, e := range Ae {
		cx++; if IsEncoded(e) == false { t.Errorf("%s(%s) returns false", fn, e) }
	}
	for _, e := range Je {
		cx++; if IsEncoded(e) == true  { t.Errorf("%s(%s) returns true",  fn, e) }
	}
	if IsEncoded("") == true { t.Errorf("%s('') returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}

func TestDecodeH(t *testing.T) {
	fn := "sisimai/rfc2045.DecodeH"
	cx := 0

	for _, e := range Ae {
		if strings.Contains(e, "?Q?=") { continue }
		cv, ce := DecodeH(e)
		cx++; if cv == e   { t.Errorf("%s(%s) returns %s", fn, e, cv) }
		cx++; if ce != nil { t.Errorf("%s(%s) returns error: %s", fn, e, ce) }
	}
	for _, e := range Je {
		cv, ce := DecodeH(e)
		cx++; if cv != e   { t.Errorf("%s(%s) returns %s", fn, e, cv) }
		cx++; if ce != nil { t.Errorf("%s(%s) returns error: %s", fn, e, ce) }
	}

	cw := "何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。"
	be := `=?utf-8?B?5L2V44Gn44KC6JaE5pqX44GE44GY44KB44GY44KB44GX44Gf5omA?=
=?utf-8?B?44Gn44OL44Oj44O844OL44Oj44O85rOj44GE44Gm44GE44Gf5LqL?=
=?utf-8?B?44Gg44GR44Gv6KiY5oa244GX44Gm44GE44KL44CC?=`
	cx++; if cv, _ := DecodeH(be); cv != cw { t.Errorf("%s(%s) returns %s", fn, be[:22], cv) }
	cx++; if cv, _ := DecodeH(""); cv != "" { t.Errorf("%s('') returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

func TestDecodeB(t *testing.T) {
	fn := "sisimai/rfc2045.DecodeB"
	cx := 0
	be := []string{"44OL44Oj44O844Oz", "6YGT57ax"}
	jp := []string{"ニャーン", "道綱"}

	for j, e := range be {
		for _, f := range []string{"", "utf-8"} {
			cv, ce := DecodeB(e, f)
			cx++; if cv != jp[j] { t.Errorf("%s(%s, %s) returns %s", fn, e, f, cv) }
			cx++; if ce != nil   { t.Errorf("%s(%s, %s) returns error: %s", fn, e, f, ce) }
		}
	}
	if cv, _ := DecodeB("", ""); cv != "" { t.Errorf("%s('') returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

func TestDecodeQ(t *testing.T) {
	fn := "sisimai/rfc2045.DecodeQ"
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

