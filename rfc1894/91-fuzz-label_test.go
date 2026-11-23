//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ______  _____ ____ _  ___  ___  _  _   
// |_   _|__  ___| |_   / /  _ \|  ___/ ___/ |( _ )/ _ \| || |  
//   | |/ _ \/ __| __| / /| |_) | |_ | |   | |/ _ \ (_) | || |_ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___| | (_) \__, |__   _|
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_|\___/  /_/   |_|  
package rfc1894
import "testing"

// % go test -tags=fuzz -fuzz=FuzzLabel -fuzztime=1m ./rfc1894
func FuzzLabel(f *testing.F) {
	fieldindex := []string{
		"Action: failed",
		"Arrival-Date: Thu, 29 Apr 2015 23:34:45 +0000",
		"Diagnostic-Code: smtp; 550 5.1.1 <neko@example.co.jp>: Recipient address rejected: User unknown in virtual mailbox table",
		"Final-Recipient: rfc822;neko@example.co.jp",
		"Last-Attempt-Date: Thu, 29 Apr 2010 23:34:45 +0900 (JST)",
		"Original-Recipient: rfc822;neko@example.co.jp",
		"Received-From-MTA: smtp; mail.example.com ([127.0.0.1])",
		"Remote-MTA: dns; 127.0.0.1",
		"Reporting-MTA: dns; neko1.example.com",
		"Status: 5.1.1",
		"X-Virus-Status: Clean",
		"X-Actual-Recipient: RFC822; userunknown@example.co.jp",
	}

	for _, e := range fieldindex { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, text string) { _ = Label(text) })
}

