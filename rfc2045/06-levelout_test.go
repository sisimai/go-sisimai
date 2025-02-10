// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045

//  _____         _      ______  _____ ____ ____   ___  _  _  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___ \ / _ \| || || ___| 
//   | |/ _ \/ __| __| / /| |_) | |_ | |     __) | | | | || ||___ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ / __/| |_| |__   _|__) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_____|\___/   |_||____/ 
import "testing"

func TestLevelOut(t *testing.T) {
	fn := "sisimai/rfc2045.levelout"
	cx := 0
	ct := `multipart/mixed; boundary="b0Nvs+XKfKLLRaP/Qo8jZhQPoiqeWi3KWPXMgw=="`
	ae := `--b0Nvs+XKfKLLRaP/Qo8jZhQPoiqeWi3KWPXMgw==
Content-Description: "error-message"
Content-Type: text/plain; charset="UTF-8"
Content-Transfer-Encoding: quoted-printable

This is the mail delivery agent at messagelabs.com.

I was unable to deliver your message to the following addresses:

neko@dest.example.net

Reason: 550 neko@dest.example.net... No such user

The message subject was: Re: BOAS FESTAS!
The message date was: Tue, 23 Dec 2014 20:39:24 +0000
The message identifier was: DB/3F-17375-60D39495
The message reference was: server-5.tower-143.messagelabs.com!1419367172!32=
691968!1

Please do not reply to this email as it is sent from an unattended mailbox.
Please visit www.messagelabs.com/support for more details
about this error message and instructions to resolve this issue.


--b0Nvs+XKfKLLRaP/Qo8jZhQPoiqeWi3KWPXMgw==
Content-Type: message/delivery-status

Reporting-MTA: dns; server-15.bemta-3.messagelabs.com
Arrival-Date: Tue, 23 Dec 2014 20:39:34 +0000

`
	cv, ce := levelout(ct, &ae)
	for _, e := range cv {
		cx++; if len(cv) == 0 { t.Errorf("%s(%s) returns an empty list", fn, ae[:20]) }
		cx++; if len(ce)  > 0 { t.Errorf("%s(%s) returns error: %v", fn, ae[:20], ce) }
		cx++; if e[0]   == "" { t.Errorf("%s(%s)[0] is empty", fn, ae[:20]) }
		cx++; if e[2]   == "" { t.Errorf("%s(%s)[2] is empty", fn, ae[:20]) }
	}

	t.Logf("The number of tests = %d", cx)
}

