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

func TestMakeFlat(t *testing.T) {
	fn := "sisimai/rfc2045.MakeFlat"
	cx := 0
	ct := `multipart/report; report-type=delivery-status; boundary="NekoNyaan--------1"` 
	ae := `--NekoNyaan--------1
Content-Type: multipart/related; boundary="NekoNyaan--------2"

--NekoNyaan--------2
Content-Type: multipart/alternative; boundary="NekoNyaan--------3"

--NekoNyaan--------3
Content-Type: text/plain; charset="UTF-8"
Content-Transfer-Encoding: base64

c2lyb25la28K

--NekoNyaan--------3
Content-Type: text/html; charset="UTF-8"
Content-Transfer-Encoding: base64

PGh0bWw+CjxoZWFkPgogICAgPHRpdGxlPk5la28gTnlhYW48L3RpdGxlPgo8L2hl
YWQ+Cjxib2R5PgogICAgPGgxPk5la28gTnlhYW48L2gxPgo8L2JvZHk+CjwvaHRt
bD4K

--NekoNyaan--------2
Content-Type: image/jpg

/9j/4AAQSkZJRgABAQEBLAEsAAD/7VaWUGhvdG9zaG9wIDMuMAA4QklNBAwAAAAA
Vk4AAAABAAAArwAAAQAAAAIQAAIQAAAAVjIAGAAB/9j/7gAOQWRvYmUAZAAAAAAB
/9sAhAAGBAQEBQQGBQUGCQYFBgkLCAYGCAsMCgoLCgoMEAwMDAwMDBAMDAwMDAwM
DAwMDAwMDAwMDAwMDAwMDAwMDAwMAQcHBw0MDRgQEBgUDg4OFBQODg4OFBEMDAwM
DBERDAwMDAwMEQwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAz/wAARCAEAAK8D
AREAAhEBAxEB/90ABAAW/8QBogAAAAcBAQEBAQAAAAAAAAAABAUDAgYBAAcICQoL

--NekoNyaan--------2
Content-Type: message/delivery-status

Reporting-MTA: dns; example.jp
Received-From-MTA: dns; neko.example.jp
Arrival-Date: Thu, 11 Oct 2018 23:34:45 +0900 (JST)

Final-Recipient: rfc822; kijitora@example.jp
Action: failed
Status: 5.1.1
Diagnostic-Code: User Unknown

--NekoNyaan--------2
Content-Type: message/rfc822

Received: ...

--NekoNyaan--------2--
`
	cv, ce := MakeFlat(ct, &ae)
	cx++; if len(*cv) == 0      { t.Errorf("%s(%s, %s) returns empty", fn, ct, ae[:20]) }
	cx++; if len(ce)  != 0      { t.Errorf("%s(%s, %s) returns error: %v", fn, ct, ae[:20], ce) }
	cx++; if len(*cv) > len(ae) { t.Errorf("%s(%s, %s) returns too short", fn, ct, ae[:20]) }
	cx++; if strings.Contains(*cv, "sironeko")  == false { t.Errorf("%s(%s, %s) does not contain sironeko", fn, ct, ae[:20]) }
	cx++; if strings.Contains(*cv, "<html>")    == true  { t.Errorf("%s(%s, %s) contains <html>", fn, ct, ae[:20]) }
	cx++; if strings.Contains(*cv, "4AAQSkZJR") == true  { t.Errorf("%s(%s, %s) contains 4AAQSkZJ", fn, ct, ae[:20]) }
	cx++; if strings.Contains(*cv, "kijitora@") == false { t.Errorf("%s(%s, %s) does not contain kijitora@", fn, ct, ae[:20]) }
	cx++; if strings.Contains(*cv, "Received:") == false { t.Errorf("%s(%s, %s) does not contain Received:", fn, ct, ae[:20]) }

	ae = ""; cx++; if cv, _ = MakeFlat("", &ae); cv!= nil { t.Errorf("%s('', '')  did not return nil", fn) }
	ae = "2";cx++; if cv, _ = MakeFlat("", &ae); cv!= nil { t.Errorf("%s('', '2') did not return nil", fn) }
	ae = ""; cx++; if cv, _ = MakeFlat("2", &ae);cv!= nil { t.Errorf("%s('2', '') did not return nil", fn) }

	t.Logf("The number of tests = %d", cx)
}

