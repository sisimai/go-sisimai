// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045

//  _____         _      ______  _____ ____ ____   ___  _  _  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___ \ / _ \| || || ___| 
//   | |/ _ \/ __| __| / /| |_) | |_ | |     __) | | | | || ||___ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ / __/| |_| |__   _|__) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_____|\___/   |_||____/ 
import "testing"

func TestParameter(t *testing.T) {
	fn := "sisimai/rfc2045.Parameter"
	cx := 0
	ae := []struct {head string; attr string; value string}{
		{"text/plain charset=iso-2022-jp", "charset", "iso-2022-jp"},
		{"text/plain; charset=iso-8859-1", "charset", "iso-8859-1"},
		{"text/plain; charset=iso-8859-15", "charset", "iso-8859-15"},
		{`text/html; CHARSET="UTF-8"`, "charset", "utf-8"},
		{`text/html; CharSet="Windows-1252"`, "charset", "windows-1252"},
		{"text/plain; charset=unicode-1-1-utf-7", "charset", "unicode-1-1-utf-7"},
		{`text/plain; name="deliveryproblems.txt"`, "name", "deliveryproblems.txt"},
		{`text/plain; charset="us-ascii"; format=flowed`, "format", "flowed"},
		{`multipart/report; report-type=delivery-status; boundary="Nekochan22"`, "report-type", "delivery-status"},
		{`multipart/report; report-type=delivery-status; boundary="Nekochan22"`, "boundary", "Nekochan22"},
	}

	for _, e := range ae {
		cx++; if cv := Parameter(e.head, e.attr); cv != e.value {
			t.Errorf("%s(%s, %s) returns %s", fn, e.head, e.attr, cv)
		}
	}
	if cv := Parameter("", "");    cv != "" { t.Errorf("%s('', '') returns %s", fn, cv) }
	if cv := Parameter("cat", ""); cv == "" { t.Errorf("%s('cat', '') returns %s", fn, cv) }
	if cv := Parameter("", "cat"); cv != "" { t.Errorf("%s('', 'cat') returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

