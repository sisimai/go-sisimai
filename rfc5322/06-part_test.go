// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
import "testing"
import "strings"

func TestPart(t *testing.T) {
	fn := "sisimai/rfc5322.Part"
	cx := 0
	ae := `
This is a MIME-encapsulated message

The original message was received at Thu, 9 Apr 2014 23:34:45 +0900
from localhost [127.0.0.1]

   ----- The following addresses had permanent fatal errors -----
<kijitora@example.net>
    (reason: 551 not our customer)

   ----- Transcript of session follows -----
... while talking to mx-0.neko.example.jp.:
<<< 450 busy - please try later
... while talking to mx-1.neko.example.jp.:
>>> DATA
<<< 551 not our customer
550 5.1.1 <kijitora@example.net>... User unknown
<<< 503 need RCPT command [data]

Content-Type: message/delivery-status
Reporting-MTA: dns; mx.example.co.jp
Received-From-MTA: DNS; localhost
Arrival-Date: Thu, 9 Apr 2014 23:34:45 +0900

Final-Recipient: RFC822; kijitora@example.net
Action: failed
Status: 5.1.6
Remote-MTA: DNS; mx-s.neko.example.jp
Diagnostic-Code: SMTP; 551 not our customer
Last-Attempt-Date: Thu, 9 Apr 2014 23:34:45 +0900

Content-Type: message/rfc822
Return-Path: <shironeko@mx.example.co.jp>
Received: from mx.example.co.jp (localhost [127.0.0.1])
	by mx.example.co.jp (8.13.9/8.13.1) with ESMTP id fffff000000001
	for <kijitora@example.net>; Thu, 9 Apr 2014 23:34:45 +0900
Received: (from shironeko@localhost)
	by mx.example.co.jp (8.13.9/8.13.1/Submit) id fff0000000003
	for kijitora@example.net; Thu, 9 Apr 2014 23:34:45 +0900
Date: Thu, 9 Apr 2014 23:34:45 +0900
Message-Id: <0000000011111.fff0000000003@mx.example.co.jp>
content-type:       text/plain
MIME-Version: 1.0
From: Shironeko <shironeko@example.co.jp>
To: Kijitora <shironeko@example.co.jp>
Subject: Nyaaaan

Nyaaan
`
	cw := []string{"Final-Recipient", "Action", "Status", "Remote-MTA", "Diagnostic-Code", "Last-Attempt-Date"}
	for _, bo := range []bool{true, false} {
		cv := Part(&ae, []string{"Content-Type: message/rfc822"}, bo)
		cx++; if len(cv) == 0  { t.Errorf("%s(%s) returns empty", fn, ae[0:10]) }
		cx++; if len(cv) != 2  { t.Errorf("%s(%s) returns invalid elements: %d", fn, ae[0:10], len(cv)) }
		cx++; if cv[0]   == "" { t.Errorf("%s(%s)[0] is empty", fn, ae[0:10]) }
		cx++; if cv[1]   == "" { t.Errorf("%s(%s)[0] is empty", fn, ae[0:10]) }

		for _, e := range cw {
			cx++; if strings.Contains(cv[0], "\n" + e + ": ") == false { t.Errorf("%s()[0] does not include %s", fn, e) }
		}
		for _, e := range []string{"Subject", "Return-Path", "From", "Date", "From", "To"} {
			cx++; if strings.Contains(cv[1], "\n" + e + ": ") == false { t.Errorf("%s()[1] does not include %s", fn, e) }
		}

		ce := ""
		cv  = Part(&ce, []string{}, bo)
		cx++; if len(cv) == 0               { t.Errorf("%s() returns empty", fn) }
		cx++; if len(cv) != 2               { t.Errorf("%s() returns invalid elements: %d", fn, len(cv)) }
		cx++; if cv[0] != "" || cv[1] != "" { t.Errorf("%s() contains invalid string: %s %s", fn, cv[0], cv[1]) }

		cv  = Part(&ce, []string{"Content-Type: message/rfc822"}, bo)
		cx++; if len(cv) == 0               { t.Errorf("%s() returns empty", fn) }
		cx++; if len(cv) != 2               { t.Errorf("%s() returns invalid elements: %d", fn, len(cv)) }
		cx++; if cv[0] != "" || cv[1] != "" { t.Errorf("%s() contains invalid string: %s %s", fn, cv[0], cv[1]) }

		ce  = "Dummy message body"
		cv  = Part(&ce, []string{}, bo)
		cx++; if len(cv) == 0               { t.Errorf("%s() returns empty", fn) }
		cx++; if len(cv) != 2               { t.Errorf("%s() returns invalid elements: %d", fn, len(cv)) }
		cx++; if cv[0] != "" || cv[1] != "" { t.Errorf("%s() contains invalid string: %s %s", fn, cv[0], cv[1]) }
	}

	t.Logf("The number of tests = %d", cx)
}

