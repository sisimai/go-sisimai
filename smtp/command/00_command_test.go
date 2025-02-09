// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package command

//  _____         _      __             _           __                                            _ 
// |_   _|__  ___| |_   / /__ _ __ ___ | |_ _ __   / /__ ___  _ __ ___  _ __ ___   __ _ _ __   __| |
//   | |/ _ \/ __| __| / / __| '_ ` _ \| __| '_ \ / / __/ _ \| '_ ` _ \| '_ ` _ \ / _` | '_ \ / _` |
//   | |  __/\__ \ |_ / /\__ \ | | | | | |_| |_) / / (_| (_) | | | | | | | | | | | (_| | | | | (_| |
//   |_|\___||___/\__/_/ |___/_| |_| |_|\__| .__/_/ \___\___/|_| |_| |_|_| |_| |_|\__,_|_| |_|\__,_|
//                                         |_|                                                      
import "testing"

var SMTPErrors = map[string][]string{
	"HELO": []string{
		"lost connection with mx.example.jp[192.0.2.2] while performing the HELO handshake",
		"SMTP error from remote mail server after HELO mx.example.co.jp:",
	},
	"EHLO": []string{
		"SMTP error from remote mail server after EHLO neko.example.com:",
	},
	"MAIL": []string{
		"452 4.3.2 Connection rate limit exceeded. (in reply to MAIL FROM command)",
		"5.1.8 (Server rejected MAIL FROM address)",
		"5.7.1 Access denied (in reply to MAIL FROM command)",
		"SMTP error from remote mail server after MAIL FROM:<shironeko@example.jp> SIZE=1543:",
	},
	"RCPT": []string{
		"550 5.1.1 <DATA@MAIL.EXAMPLE.JP>... User Unknown  in RCPT TO",
		"550 user unknown (in reply to RCPT TO command)",
		">>> RCPT To:<mikeneko@example.co.jp>",
		"most progress was RCPT TO response; remote host 192.0.2.32 said: 550 Unknown user MAIL@example.ne.jp",
		"SMTP error from remote mail server after RCPT TO:<kijitora@example.jp>:",
	},
	"DATA": []string{
		"Email rejected per DMARC policy for libsisimai.org (in reply to end of DATA command)",
		"SMTP Server <192.0.2.223> refused to accept your message (DATA), with the following error message",
	},
}
var IsntErrors = []string{
	"nekochan",
	"ニャーン?",
	"HELOWORLD!!!",
	"Sendmail 8.17.1",
	"Database Server",
	"",
}

func TestTest(t *testing.T) {
	fn := "sisimai/smtp/command.Test"
	cx := 0

	for e := range SMTPErrors {
		for _, f := range SMTPErrors[e] {
			cx++; if cv := Test(f); cv == false { t.Errorf("%s(%s) returns false", fn, f) }
		}
	}
	cx++; if cv := Test("");    cv == true { t.Errorf("%s(%s) returns true", fn, "") }
	cx++; if cv := Test("cat"); cv == true { t.Errorf("%s(%s) returns true", fn, "cat") }

	t.Logf("The number of tests = %d", cx)
}

func TestFind(t *testing.T) {
	fn := "sisimai/smtp/command.Find"
	cx := 0

	for e := range SMTPErrors {
		for _, f := range SMTPErrors[e] {
			cx++; if cv := Find(f); cv != e { t.Errorf("%s(%s) returns (%s)", fn, f, cv) }
		}
	}
	for _, e := range IsntErrors {
		cx++; if cv := Find(e); cv != "" { t.Errorf("%s(%s) returns (%s)", fn, e, cv) }
	}
	t.Logf("The number of tests = %d", cx)
}

