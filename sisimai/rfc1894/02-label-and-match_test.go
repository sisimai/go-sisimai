// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc1894

//  _____         _      ______  _____ ____ _  ___  ___  _  _   
// |_   _|__  ___| |_   / /  _ \|  ___/ ___/ |( _ )/ _ \| || |  
//   | |/ _ \/ __| __| / /| |_) | |_ | |   | |/ _ \ (_) | || |_ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___| | (_) \__, |__   _|
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_|\___/  /_/   |_|  
import "testing"
import "strings"
import sisimoji "sisimai/string"

var FieldList = []string{
	"Reporting-MTA: dns; mx.example.jp",
	"Received-From-MTA: DNS; localhost.localdomain",
	"Remote-MTA: dns; mail.neko.example.org",
	"Arrival-Date: Thu, 29 Apr 2009 23:45:00 +0900",
	"Final-Recipient: RFC822; kijitora@example.com",
	"Action: failed",
	"Status: 5.1.8",
	"Diagnostic-Code: SMTP; 553 5.1.8 <httpd@host1.mx.example.jp>... Domain of sender address httpd@host1.mx.example.jp does not exist",
	"Last-Attempt-Date: Thu, 29 Apr 2009 23:45:00 +0900",
	"Original-Recipient: rfc822;michitsuna@example.org",
	"X-Actual-Recipient: X-Unix; |/var/adm/sm.bin/neko",
}
var LowerList = []string{
	"reporting-mta",
	"received-from-mta",
	"remote-mta",
	"arrival-date",
	"final-recipient",
	"action",
	"status",
	"diagnostic-code",
	"last-attempt-date",
	"original-recipient",
	"x-actual-recipient",
}

func TestLabel(t *testing.T) {
	fn := "sisimai/rfc1894.Label"
	cx := 0

	for j, e := range FieldList {
		cx++; if cv := Label(e); cv != LowerList[j] { t.Errorf("%s(%s) returns %s", fn, e, cv) }
	}

	t.Logf("The number of tests = %d", cx)
}

func TestMatch(t *testing.T) {
	fn := "sisimai/rfc1894.Match"
	cx := 0

	for _, e := range FieldList {
		cx++; if cv := Match(e); cv == 0 { t.Errorf("%s(%s) returns 0", fn, e) }
	}
	for _, e := range []string{"Subject: neko", "From: postmaster", "Return-Path: <>", "To: root"} {
		cx++; if cv := Match(e); cv != 0 { t.Errorf("%s(%s) returns %d", fn, e, cv) }
	}

	t.Logf("The number of tests = %d", cx)
}

func TestField(t *testing.T) {
	fn := "sisimai/rfc1894.Field"
	cx := 0

	for _, e := range FieldList {
		cv := Field(e)
		cx++; if len(cv) == 0 { t.Errorf("%s(%s) returns an empty list", fn, e) }
		cx++; if len(cv) != 5 { t.Errorf("%s(%s) did not return 5 elements", fn, e) }
		cx++; if sisimoji.EqualsAny(cv[0], LowerList) == false { t.Errorf("%s(%s)[0] is %s", fn, e, cv[0]) }

		if strings.Contains(e, ";") {
			cx++; if cv[1] == "" { t.Errorf("%s(%s)[1] is empty", fn, e) }
			cx++; if sisimoji.ContainsAny(strings.ToLower(cv[1]), []string{"rfc822", "smtp", "dns", "x-unix"}) == false {
				t.Errorf("%s(%s)[1] is invalid subtype: %s", fn, e, cv[1])
			}
			cx++; if cv[2] == "" { t.Errorf("%s(%s)[2] is empty", fn, e) }
			cx++; if sisimoji.EqualsAny(cv[3], []string{"addr", "code", "date", "host", "list", "stat", "text"}) == false {
				t.Errorf("%s(%s)[3] is invalid group: %s", fn, e, cv[3])
			}
			cx++; if cv[4] != "" { t.Errorf("%s(%s)[4] is not empty: %s", fn, e, cv[4]) }
		}
	}
	for _, e := range []string{"Subject: neko", "From: postmaster", "Return-Path: <>", "To: root"} {
		cx++; if cv := Field(e); len(cv) != 0 { t.Errorf("%s(%s) returns %v", fn, e, cv) }
	}

	t.Logf("The number of tests = %d", cx)
}

