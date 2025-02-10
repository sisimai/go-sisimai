// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc1123

//  _____         _      ______  _____ ____ _ _ ____  _____ 
// |_   _|__  ___| |_   / /  _ \|  ___/ ___/ / |___ \|___ / 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   | | | __) | |_ \ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___| | |/ __/ ___) |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|_|_|_____|____/ 
import "testing"

func TestFind(t *testing.T) {
	fn := "sisimai/rfc1123.Find"
	cx := 0
	ae := []string{
		"<neko@example.jp>: host neko.example.jp[192.0.2.2] said: 550 5.7.1 This message was not accepted due to domain (libsisimai.org) owner DMARC policy",
		"neko.example.jp[192.0.2.232]: server refused to talk to me: 421 Service not available, closing transmission channel",
		"... while talking to neko.example.jp.: <<< 554 neko.example.jp ESMTP not accepting connections",
		"host neko.example.jp [192.0.2.222]: 500 Line limit exceeded",
		"Google tried to deliver your message, but it was rejected by the server for the recipient domain nyaan.jp by neko.example.jp. [192.0.2.2].",
		"Delivery failed for the following reason: Server neko.example.jp[192.0.2.222] failed with: 550 <kijitora@example.jp> No such user here",
		"Remote system: dns;neko.example.jp (TCP|17.111.174.65|48044|192.0.2.225|25) (neko.example.jp ESMTP SENDMAIL-VM)",
		"SMTP Server <neko.example.jp> rejected recipient <cat@libsisimai.org> (Error following RCPT command). It responded as follows: [550 5.1.1 User unknown]",
		"Reporting-MTA:      <neko.example.jp>",
		"cat@example.jp:000000:<cat@example.jp> : 192.0.2.250 : neko.example.jp:[192.0.2.153] : 550 5.1.1 <cat@example.jp>... User Unknown  in RCPT TO",
		"Generating server: neko.example.jp",
		"Server di generazione: neko.example.jp",
		"Serveur de génération : neko.example.jp",
		"Genererande server: neko.example.jp",
		"neko.example.jp [192.0.2.25] did not like our RCPT TO: 550 5.1.1 <cat@example.jp>: Recipient address rejected: User unknown",
		"neko.example.jp [192.0.2.79] did not like our final DATA: 554 5.7.9 Message not accepted for policy reasons",
	}

	for _, e := range ae {
		cv := Find(e)
		cx++; if cv == ""                    { t.Errorf("%s(%s) returns an empty text", fn, e)   }
		cx++; if cv != "neko.example.jp"     { t.Errorf("%s(%s) returns hostname %s", fn, e, cv) }
		cx++; if IsInternetHost(cv) == false { t.Errorf("IsInternetHost(%s) returns false", cv)  }
	}
	cx++; if cv := Find(""); cv != "" { t.Errorf("%s() returns %s", fn, cv) }

	t.Logf("The number of tests = %d", cx)
}

