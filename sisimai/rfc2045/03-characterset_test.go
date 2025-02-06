// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045

//  _____         _      ______  _____ ____ ____   ___  _  _  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___ \ / _ \| || || ___| 
//   | |/ _ \/ __| __| / /| |_) | |_ | |     __) | | | | || ||___ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ / __/| |_| |__   _|__) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_____|\___/   |_||____/ 
import "testing"

func TestCharacterSet(t *testing.T) {
	fn := "sisimai/rfc2045.CharacterSet"
	cx := 0
	ae := []struct {head string; char string}{
		{"=?UTF-8?B?6YGT57ax44OL44Oj44O8?=", "UTF-8"},
		{"=?ISO-2022-JP?B?GyRCRjs5SyVLJWMhPBsoQg==?=", "ISO-2022-JP"},
		{"=?EUC-JP?B?xru5y6XLpeOhvA==?=", "EUC-JP"},
		{"=?SHIFT_JIS?B?k7mNaoNqg4OBWw==?=", "SHIFT_JIS"},
		{"=?utf-8?q?=E9=81=93=E7=B6=B1=E3=83=8B=E3=83=A3=E3=83=BC?=", "UTF-8"},
		{"=?iso-2022-jp?q?=1B=24BF=3B9K=25K=25c=21=3C=1B=28B?=", "ISO-2022-JP"},
		{"=?euc-jp?q?=C6=BB=B9=CB=A5=CB=A5=E3=A1=BC?=", "EUC-JP"},
		{"=?shift_jis?q?=93=B9=8Dj=83j=83=83=81=5B?=", "SHIFT_JIS"},
	}
	for _, e := range ae {
		cx++; if cv := CharacterSet(e.head); cv != e.char { t.Errorf("%s(%s) returns %s", fn, e.head, cv) }
	}
	if cv := CharacterSet(""); cv != "" { t.Errorf("%s('') returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

