//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
package sisimai
import "testing"
import "strings"

func FuzzRise(f *testing.F) {
	headerseed := []string{
		"From MAILER-DAEMON Fri Feb  2 18:30:22 2018",
		"Return-Path: <>",
		"From: Mail Delivery Subsystem <postmaster@libsisimai.org>",
		"To: Kijitora-Neko <kijitora@libsisimai.net>",
		"Subject: Returned mail: see transcript for details",
		"Date: Fri,  2 Feb 2018 18:30:22 +0900 (JST)",
		"Message-Id: <neko-michitsuna-20180202183022@nyaan.jp>",
		`Content-Type: multipart/report; report-type=delivery-status; boundary="1CDEA1EB158C.1110549095/mail.example.co.jp"`,
		"Content-Transfer-Encoding: 7bit",
		"MIME-Version: 1.0",
		"Auto-Submitted: auto-replied",
		"Content-Language: en-US",
		"X-Mailer: Internet Mail Service (5.5.1960.3)",
		"X-Failed-Recipients: chashiro@libsisimai.org",
		"X-AHMAILID: 20110429T233445+0900-0000-0000D0-AAA-FF",
		"X-MS-Embedded-Report: ",
		"X-MimeOLE: Produced By Microsoft MimeOLE V15.4.3555.308",
		"X-MX-Bounce: mta/src/queue/bounce",
	}
	receivedby := [][]string{
		[]string{
			"from mx.libsisimai.org (ip-192-0-2-25.ap-northeast-1.compute.internal [192.0.2.25]) ",
			"by smtp.nyaan.jp (8.14.1/8.14.1) with ESMTPS id m8I8s45D007047 ",
			"for <kijitora@libsisimai.net>; ",
			"Fri,  2 Feb 2018 18:30:22 +0900 (JST)",
		},
		[]string{
			"from hub.libsisimai.org (ip-192-0-2-22.us-east-1.compute.internal [192.0.2.22]) ",
			"by mx2.nyaan.jp (Postfix) with ESMTPS id K1bgZJ3Hn6zmQTy6 ",
			"for <kijitora@libsisimai.net>; ",
			"Fri,  2 Feb 2018 18:30:22 +0900 (EST)",
		},
		[]string{"(qmail 3016 invoked for bounce); 25 Jun 2024 10:36:44 -0000"},
	}
	emailbody1 := `
This is a MIME-encapsulated message.

--1CDEA1EB158C.1110549095/mail.example.co.jp
Content-Description: Notification
Content-Type: text/plain; charset=us-ascii

This is the mail system at host mail.example.co.jp.

I'm sorry to have to inform you that your message could not
be delivered to one or more recipients. It's attached below.

For further assistance, please send mail to <postmaster>

If you do so, please include this problem report. You can
delete your own text from the attached returned message.

                   The mail system

<kijitora@hotmail.example.com>: host mx3.hotmail.com[198.51.100.2] said: 554
    Transaction failed (in reply to MAIL FROM command)

--1CDEA1EB158C.1110549095/mail.example.co.jp
Content-Description: Delivery report
Content-Type: message/delivery-status

Reporting-MTA: dns; mail.example.co.jp
X-Postfix-Queue-ID: 1CDEA1EB158C
X-Postfix-Sender: rfc822; neko-nyaan@example.org
Arrival-Date: Thu, 29 Apr 2017 23:34:45 +0900 (JST)

Final-Recipient: rfc822; kijitora@hotmail.example.com
Original-Recipient: rfc822;kijitora@hotmail.example.com
Action: failed
Status: 5.0.0
Remote-MTA: dns; mx3.hotmail.com
Diagnostic-Code: smtp; 554 Transaction failed

--1CDEA1EB158C.1110549095/mail.example.co.jp
Content-Description: Undelivered Message
Content-Type: message/rfc822
Content-Transfer-Encoding: 7bit

Received: from localhost (localhost.localdomain [127.0.0.1])
	by mail.example.co.jp (Postfix) with ESMTP id 1CDEA1EB158C
	for <kijitora@hotmail.example.com>; Thu, 29 Apr 2017 23:34:45 +0900 (JST)
Received: from mail.example.co.jp ([127.0.0.1])
	by localhost (smtp-1.example.co.jp [127.0.0.1]) (amavisd-new, port 10024)
	with ESMTP id VicMejKaknys for <kijitora@hotmail.example.com>;
	Thu, 29 Apr 2017 23:34:45 +0900 (JST)
Received: from localhost (localhost.localdomain [127.0.0.1])
	by mail.example.co.jp (Postfix) with ESMTP id 13D116269C08
	for <kijitora@hotmail.example.com>; Thu, 29 Apr 2017 23:34:45 +0900 (JST)
MIME-Version: 1.0
Content-Type: text/plain; charset=ISO-2022-JP
Content-Transfer-Encoding: quoted-printable
From: nekochan@example.org
To: kijitora@hotmail.example.com
Subject: Nyaan
Message-Id: <20170429233445.13D116269C08@mail.example.co.jp>
Date: Thu, 29 Apr 2017 23:34:45 +0900 (JST)

Nyaan

--1CDEA1EB158C.1110549095/mail.example.co.jp--
`

	for _, e := range headerseed { f.Add([]byte(e + "\n"))                                    }
	for _, e := range receivedby { f.Add([]byte("Received: " + strings.Join(e, "\n") + "\n")) }
	f.Add([]byte("\n" + emailbody1 + "\n"))
	f.Fuzz(func(t *testing.T, data []byte) { _, _ = Rise(string(data), nil) })
}

