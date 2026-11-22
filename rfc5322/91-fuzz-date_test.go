//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
package rfc5322
import "testing"

// % go test -tags=fuzz -fuzz=FuzzDate -fuzztime=1m ./rfc5322
func FuzzDate(f *testing.F) {
	datestring := []string{
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
	for _, e := range datestring { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, date string) { _ = Date(date) })
}

