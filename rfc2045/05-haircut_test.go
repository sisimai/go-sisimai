// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045

//  _____         _      ______  _____ ____ ____   ___  _  _  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___ \ / _ \| || || ___| 
//   | |/ _ \/ __| __| / /| |_) | |_ | |     __) | | | | || ||___ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ / __/| |_| |__   _|__) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_____|\___/   |_||____/ 
import "testing"

func TestHairCut(t *testing.T) {
	fn := "sisimai/rfc2045.haircut"
	cx := 0
	ae := `Content-Description: "error-message"
Content-Type: text/plain; charset="UTF-8"
Content-Transfer-Encoding: quoted-printable

This is the mail delivery agent at messagelabs.com.

I was unable to deliver your message to the following addresses:

neko@dest.example.net

Reason: 550 neko@dest.example.net... No such user`

	cv := haircut(&ae, false)
	cx++; if len(cv) == 0 { t.Errorf("%s(%s) returns an empty list", fn, ae[:20]) }
	cx++; if cv[0] != `text/plain; charset="utf-8"` { t.Errorf("%s(%s)[0] returns %s", fn, ae[:20], cv[0]) }
	cx++; if cv[1] != "quoted-printable"            { t.Errorf("%s(%s)[1] returns %s", fn, ae[:20], cv[1]) }
	cx++; if cv[2] == ""                            { t.Errorf("%s(%s)[2] returns empty", fn, ae[:20])     }

	t.Logf("The number of tests = %d", cx)
}

