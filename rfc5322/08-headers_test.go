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
import "os"

func TestHeaders(t *testing.T) {
	fn := "rfc5322.Headers"
	cx := 0
	ae := []string{
		"../set-of-emails/maildir/bsd/lhost-opensmtpd-17.eml",
		"../set-of-emails/maildir/bsd/lhost-postfix-78.eml",
	}
	for _, ef := range ae {
		bx, _ := os.ReadFile(ef); if len(bx) == 0 {
			cx++; t.Fatalf("os.ReadFile(%s) returns an empty string", ef)
		}
		et := string(bx); em, _ := mail.ReadMessage(strings.NewReader(et))
		cv := Headers(&em.Header, false)

		cx++; if len(cv) == 0 { t.Errorf("%s() returns empty", fn) }
		for e := range cv {
			cx++; if e == ""          { t.Errorf("%s() returns an empty key", fn) }
			cx++; if len(cv[e]) == 0  { t.Errorf("%s()[%s] have no elements", fn, e) }
			cx++; if cv[e][0]   == "" { t.Errorf("%s()[%s][0] is empty", fn, e) }
		}
		for _, e := range []string{"from", "received", "message-id", "content-type", "subject"} {
			cx++; if cv[e][0]   == "" { t.Errorf("%s()[%s][0] is empty", fn, e) }
		}
		for _, e := range []string{"dkim-signature", "authentication-results", "received-spf"} {
			cx++; if len(cv[e]) != 0  { t.Errorf("%s()[%s] is not empty", fn, e) }
		}
	}

	t.Logf("The number of tests = %d", cx)
}

