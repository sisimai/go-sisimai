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
import "net/mail"

func TestDate(t *testing.T) {
	fn := "sisimai/rfc5322.Date"
	cx := 0
	ae := []string{
		"Sun, 24 Jun 2018 06:28:11 +0200 (CEST)",
		"Sun, 03 Oct 2010 21:11:37 +0000",
		"Mon, 24 Nov 2014 14:23:04 +0300 (MSK)",
		"Mon, 20 Sep 2021 21:32:59 +0200 (GMT+02:00)",
		"Tue, 029 Apr 2019 23:34:45 -0800 (PST)",
		"Wed, 3 May 2007 23:34:45",
		"    Thu, 29 Apr 2008 23:45:10 -0500",
		"Thu, 29 Apr 2009 23:34:45",
		"Thu, 2 May 2024 17:48:55 +0000 (UTC)",
		"Thu, 17 Jul 2017 23:34:45 +0500 (PKT)",
		"Thu, 29 Apr 2019 23:34:45 +0900 (JST)",
		"Thu, 29 Apr 2019 23:34:45 -0800 (PST)",
		"Thu, 29 Apr 2022 23:34:45 +0000",
		"Thu, 29 Apr 2022 23:34:45 +0000 (GMT)",
		"Thu, 29 Apr 2013 23:34:45 GMT",
		"Fri,  5 Aug 2022 05:22:50 +0900 (JST)",
		"Sat, 06 Jul 2013 23:34:45 JST",
	}

	for _, e := range ae {
		cv := Date(e)
		cx++; if cv      == ""  { t.Errorf("%s(%s) returns empty", fn, e) }
		cx++; if cv[3:4] != "," { t.Errorf("%s(%s)[3] is not `,`", fn, e) }
		cx++; if strings.HasPrefix(cv, " ") == true { t.Errorf("%s(%s) begins with ` `", fn, e) }

		ct, ce := mail.ParseDate(cv)
		cx++; if ce != nil   { t.Errorf("net/mail.ParseDate(%s) returns error: %s", cv, ce) }
		cx++; if ct.IsZero() { t.Errorf("net/mail.ParseDate(%s).IsZero() is true", cv) }
	}
	cx++; if cv := Date("");     cv != "" { t.Errorf("%s() returns %s", fn, cv) }
	cx++; if cv := Date("Neko"); cv != "" { t.Errorf("%s() returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

