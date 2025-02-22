// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package status

//  _____         _      __             _           __   _        _             
// |_   _|__  ___| |_   / /__ _ __ ___ | |_ _ __   / /__| |_ __ _| |_ _   _ ___ 
//   | |/ _ \/ __| __| / / __| '_ ` _ \| __| '_ \ / / __| __/ _` | __| | | / __|
//   | |  __/\__ \ |_ / /\__ \ | | | | | |_| |_) / /\__ \ || (_| | |_| |_| \__ \
//   |_|\___||___/\__/_/ |___/_| |_| |_|\__| .__/_/ |___/\__\__,_|\__|\__,_|___/
//                                         |_|                                  
import "testing"
import "strings"

var ReasonList = []string{
	"authfailure", "badreputation", "blocked", "contenterror", "exceedlimit", "expired", "failedstarttls",
	"filtered", "hasmoved", "hostunknown", "mailboxfull", "mailererror", "mesgtoobig", "networkerror",
	"notaccept", "onhold", "rejected", "norelaying", "spamdetected", "virusdetected",  "policyviolation",
	"securityerror", "speeding", "suppressed", "suspend", "requireptr", "notcompliantrfc", "systemerror",
	"systemfull", "toomanyconn", "userunknown", "syntaxerror",
}
var StatusList = []string{
	"2.1.5",
	"4.1.6", "4.1.7", "4.1.8", "4.1.9", "4.2.1", "4.2.2", "4.2.3", "4.2.4", "4.3.1", "4.3.2", "4.3.3",
	"4.3.5", "4.4.1", "4.4.2", "4.4.4", "4.4.5", "4.4.6", "4.4.7", "4.5.3", "4.5.5", "4.6.0", "4.6.2",
	"4.6.5", "4.7.1", "4.7.2", "4.7.5", "4.7.6", "4.7.7",
	"5.1.0", "5.1.1", "5.1.2", "5.1.3", "5.1.4", "5.1.6", "5.1.7", "5.1.8", "5.1.9", "5.2.0", "5.2.1",
	"5.2.2", "5.2.3", "5.2.4", "5.3.0", "5.3.1", "5.3.2", "5.3.3", "5.3.4", "5.3.5", "5.4.0", "5.4.3",
	"5.5.3", "5.5.4", "5.5.5", "5.5.6", "5.6.0", "5.6.1", "5.6.2", "5.6.3", "5.6.5", "5.6.6", "5.6.7",
	"5.6.8", "5.6.9", "5.7.0", "5.7.1", "5.7.2", "5.7.3", "5.7.4", "5.7.5", "5.7.6", "5.7.7", "5.7.8",
	"5.7.9",
}
var SMTPErrors = []string{
	"smtp; 2.1.5 250 OK",
	"smtp;550 5.2.2 <mikeneko@example.co.jp>... Mailbox Full",
	"smtp; 550 5.1.1 Mailbox does not exist",
	"smtp; 550 5.1.1 Mailbox does not exist",
	"smtp; 450 4.0.0 Temporary failure",
	"smtp; 552 5.2.2 Mailbox full",
	"smtp; 552 5.3.4 Message too large",
	"smtp; 500 5.6.1 Message content rejected",
	"smtp; 550 5.2.0 Message Filtered",
	"550 5.1.1 <kijitora@example.jp>... User Unknown",
	"SMTP; 552-5.7.0 This message was blocked because its content presents a potential",
	"SMTP; 550 5.1.1 Requested action not taken: mailbox unavailable",
	"SMTP; 550 5.7.1 IP address blacklisted by recipient",
	"SMTP; 550 5.7.25 The ip address sending this message does not have a ptr record setup",
	"smtp; 550-5.7.1 This message is not RFC 5322 compliant. There are multiple Subject 550-5.7.1 headers",
}

func TestCode(t *testing.T) {
	fn := "sisimai/smtp/status.Code"
	cx := 0

	for _, e := range ReasonList {
		if e == "hasmoved" || e == "hostunknown" || e == "userunknown" {
			cx++; if cv := Code(e, true);  cv != "" { t.Errorf("%s(%s, true) returns (%s)", fn, e, cv) }
			cx++; if cv := Code(e, false); strings.HasPrefix(cv, "5.0.9") == false {
				t.Errorf("%s(%s, true) returns (%s)", fn, e, cv)
			}
		} else {
			cx++; if cv := Code(e, true);  strings.HasPrefix(cv, "4.0.9") == false {
				t.Errorf("%s(%s, true) returns (%s)", fn, e, cv)
			}
			cx++; if cv := Code(e, false); strings.HasPrefix(cv, "5.0.9") == false {
				t.Errorf("%s(%s, true) returns (%s)", fn, e, cv)
			}
		}
	}
	cx++; if cv := Code("neko", true); cv  != "" { t.Errorf("%s(neko, true) returns (%s)", fn, cv) }
	cx++; if cv := Code("neko", false); cv != "" { t.Errorf("%s(neko, false) returns (%s)", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

func TestName(t *testing.T) {
	fn := "sisimai/smtp/status.Name"
	cx := 0

	for _, e := range StatusList {
		cx++; if cv := Name(e); cv == "" { t.Errorf("%s(%s) returns empty", fn, e) }
	}
	cx++; if cv := Name("123"); cv != "" { t.Errorf("%s(123) returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

func TestTest(t *testing.T) {
	fn := "sisimai/smtp/status.Test"
	cx := 0
	ae := []string{
		"3.14", "9.99", "5.0.3.2", "1.0.0", "3.1.4", "6.7.8", "5.-1.0", "5.12.0", "5.2.-2",
		"5.2.2202", "8.8.8.8", "192.0.2.25", "", "nekochan",
	}

	for _, e := range StatusList {
		cx++; if cv := Test(e); cv == false { t.Errorf("%s(%s) returns false", fn, e) }
	}
	for _, e := range ae {
		cx++; if cv := Test(e); cv == true  { t.Errorf("%s(%s) returns true",  fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

func TestFind(t *testing.T) {
	fn := "sisimai/smtp/status.Test"
	cx := 0

	for _, e := range SMTPErrors {
		cx++; if cv := Find(e, ""); cv == "" { t.Errorf("%s(%s) returns empty", fn, e) }
	}
	cx++; if cv := Find("", "");    cv != "" { t.Errorf("%s('') returns (%s)", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

func TestPrefer(t *testing.T) {
	fn := "sisimai/smtp/status.Prefer"
	ae := []struct {lhs string; rhs string; rep string; exp string}{
		{"", "", "", ""},
		{"5.2.2", "", "", "5.2.2"},
		{"5.0.0", "5.1.1", "", "5.1.1"},
		{"5.2.0", "5.2.1", "", "5.2.1"},
		{"4.4.7", "4.2.2", "", "4.2.2"},
		{"5.7.8", "4.4.0", "550", "5.7.8"},
		{"4.2.1", "5.7.0", "421", "4.2.1"},
		{"5.7", "5.7.26", "421", "5.7.26"},
		{"5.7.26", "5.7", "421", "5.7.26"},
	}
	cx := 0

	for _, e := range ae {
		cx++; if cv := Prefer(e.lhs, e.rhs, e.rep); cv != e.exp {
			t.Errorf("%s(%s, %s, %s) returns %s", fn, e.lhs, e.rhs, e.rep, cv)
		}
	}

	t.Logf("The number of tests = %d", cx)
}

