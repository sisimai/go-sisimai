// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package failure

//  _____         _      __             _           ____       _ _                
// |_   _|__  ___| |_   / /__ _ __ ___ | |_ _ __   / / _| __ _(_) |_   _ _ __ ___ 
//   | |/ _ \/ __| __| / / __| '_ ` _ \| __| '_ \ / / |_ / _` | | | | | | '__/ _ \
//   | |  __/\__ \ |_ / /\__ \ | | | | | |_| |_) / /|  _| (_| | | | |_| | | |  __/
//   |_|\___||___/\__/_/ |___/_| |_| |_|\__| .__/_/ |_|  \__,_|_|_|\__,_|_|  \___|
//                                         |_|                                    
import "testing"

var SoftBounce = []string{
	"blocked", "contenterror", "exceedlimit", "expired", "failedstarttls", "filtered", "mailboxfull",
	"mailererror", "mesgtoobig", "networkerror", "norelaying", "rejected", "securityerror",
	"spamdetected", "suspend", "systemerror", "systemfull", "toomanyconn", "undefined", "onhold",
}
var HardBounce = []string{"userunknown", "hostunknown", "hasmoved", "notaccept"}
var IsntBounce = []string{"delivered", "feedback", "vacation"}
var IsntErrors = []string{"smtp; 2.1.5 250 OK"}
var TempErrors = []string{
	"smtp; 450 4.0.0 Temporary failure",
	"smtp; 554 4.4.7 Message expired: unable to deliver in 840 minutes.<421 4.4.2 Connection timed out>",
	"SMTP; 450 4.7.1 Access denied. IP name lookup failed [192.0.2.222]",
	"smtp; 451 4.7.650 The mail server [192.0.2.25] has been",
	"4.4.1 (Persistent transient failure - routing/network: no answer from host)",
};
var PermErrors = []string{
	"smtp;550 5.2.2 <mikeneko@example.co.jp>... Mailbox Full",
	"smtp; 550 5.1.1 Mailbox does not exist",
	"smtp; 550 5.1.1 Mailbox does not exist",
	"smtp; 552 5.2.2 Mailbox full",
	"smtp; 552 5.3.4 Message too large",
	"smtp; 500 5.6.1 Message content rejected",
	"smtp; 550 5.2.0 Message Filtered",
	"550 5.1.1 <kijitora@example.jp>... User Unknown",
	"SMTP; 552-5.7.0 This message was blocked because its content presents a potential",
	"SMTP; 550 5.1.1 Requested action not taken: mailbox unavailable",
	"SMTP; 550 5.7.1 IP address blacklisted by recipient",
};

func TestIsPermanent(t *testing.T) {
	fn := "sisimai/smtp/failure.IsPermanent"
	cx := 0

	for _, e := range PermErrors {
		cx++; if cv := IsPermanent(e); cv == false { t.Errorf("%s(%s) returns false", fn, e) }
	}
	cx++; if IsPermanent("") == true { t.Errorf("%s(%s) returns true", fn, "") }

	t.Logf("The number of tests = %d", cx)
}

func TestIsTemporary(t *testing.T) {
	fn := "sisimai/smtp/failure.IsTemporary"
	cx := 0

	for _, e := range TempErrors {
		cx++; if cv := IsTemporary(e); cv == false { t.Errorf("%s(%s) returns false", fn, e) }
	}
	cx++; if IsTemporary("") == true { t.Errorf("%s(%s) returns true", fn, "") }

	t.Logf("The number of tests = %d", cx)
}

