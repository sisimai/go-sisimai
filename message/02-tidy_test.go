// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package message

//  _____         _      __                                        
// |_   _|__  ___| |_   / / __ ___   ___  ___ ___  __ _  __ _  ___ 
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
//   | |  __/\__ \ |_ / /| | | | | |  __/\__ \__ \ (_| | (_| |  __/
//   |_|\___||___/\__/_/ |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                                      |___/      
import "testing"
import "strings"

func TestTidy(t *testing.T) {
	fn := "sisimai/message.tidy"
	cx := 0
	ae := `This is a MIME-encapsulated message

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
	cv := tidy(&ae)
	cx++; if len(*cv) == 0 { t.Errorf("%s() returns empty", fn) }
	cx++; if strings.Contains(*cv, "Content-Type: text/plain") == false { t.Errorf("%s() does not contain Content-Type heaader", fn) }

	we := []struct {tidied string; examples []string}{
		{"Action: failed", []string{"Action: FAILED", "ACTION:   Failed"}},
		{"Arrival-Date: Sat, 3 Oct 2020 20:11:48 +0900", []string{"Arrival-DATE: Sat,      3 Oct 2020 20:11:48 +0900"}},
		{"Diagnostic-Code: smtp; 550 Host does not accept mail", []string{"Diagnostic-code:SMTP;550 Host does not accept mail"}},
		{"Final-Recipient: rfc822; neko@libsisimai.org", []string{"Final-recipient: RFC822;NEKO@libsisimai.org"}},
		{"Last-Attempt-Date: Sat, 3 Oct 2020 20:12:06 +0900", []string{"Last-Attempt-DATE:Sat, 3    Oct 2020 20:12:06 +0900"}},
		{"Original-Recipient: rfc822; neko@example.com", []string{"Original-recipient:rfc822;NEKO@example.com"}},
		{"Received-From-MTA: dns; localhost", []string{"Received-From-mta:    DNS; LocalHost"}},
		{"Remote-MTA: dns; mx.libsisimai.org", []string{"Remote-mta: DNS; mx.libsisimai.org"}},
		{"Reporting-MTA: dns; nyaan.example.jp", []string{"Reporting-mta: DNS;   nyaan.example.jp"}},
		{"Status: 5.0.0 (permanent failure)", []string{"STATUS:    5.0.0 (permanent failure)"}},
		{"X-Actual-Recipient: rfc822; neko@libsisimai.org", []string{"X-Actual-rEcipient:rfc822;NEKO@libsisimai.org"}},
		{"X-Original-Message-ID: <NEKOCHAN>", []string{"x-original-message-ID:     <NEKOCHAN>"}},
		{"Content-Type: text/plain", []string{"content-type:     TEXT/plain"}},
		{
			`Content-Type: message/delivery-status; charset=us-ascii; boundary="Neko-Nyaan-22=="`,
			[]string{
				`Content-Type:   message/xdelivery-status; charset=us-ascii; boundary="Neko-Nyaan-22=="`,
				`Content-Type: message/xdelivery-status;   charset=us-ascii; boundary="Neko-Nyaan-22=="`,
				`Content-Type: message/xdelivery-status; charset=us-ascii;   boundary="Neko-Nyaan-22=="`,
				`content-type: message/xdelivery-status; CharSet=us-ascii; Boundary="Neko-Nyaan-22=="`,
				`content-Type: Message/Xdelivery-Status; CharSet=us-ascii; Boundary="Neko-Nyaan-22=="`,
				`Content-type:message/xdelivery-status;CharSet=us-ascii;Boundary="Neko-Nyaan-22=="`,
			},
		},
	}

	for _, e := range we {
		for _, f := range e.examples {
			cx++; if cv = tidy(&f); strings.Contains(*cv, e.tidied) == false {
				t.Errorf("%s(%s) failed to tidy and returns [%s]", fn, f, *cv)
			}
		}
	}

	cw := ""
	cx += 1; if cv := tidy(&cw); *cv != "" { t.Errorf("%s() returns %s", fn, *cv) }

	t.Logf("The number of tests = %d", cx)
}

