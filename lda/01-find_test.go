// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lda

//  _____         _      ___     ____    _    
// |_   _|__  ___| |_   / / |   |  _ \  / \   
//   | |/ _ \/ __| __| / /| |   | | | |/ _ \  
//   | |  __/\__ \ |_ / / | |___| |_| / ___ \ 
//   |_|\___||___/\__/_/  |_____|____/_/   \_\
import "testing"
import "libsisimai.org/sisimai/sis"

func TestFind(t *testing.T) {
	fn := "sisimai/lda.Find"
	cx := 0
	ae := [][]string{
		[]string{`x-unix; procmail: Couldn't create "/var/spool/mail/neko" id:`, "DATA"},
		[]string{"vdelivermail: account is locked email bounced kijitora@neko2.example.co.jp", "DATA"},
		[]string{"maildrop: maildir over quota.", ""},
		[]string{`"|IFS=' ' && exec /usr/local/bin/procmail -f- || exit 75 #kijitora"`, "DATA"},
		[]string{`554 "|IFS=' ' && exec /usr/local/bin/procmail -f- || exit 75 #kijitora"... Service unavailable`, "DATA"},
		[]string{"mail.local: unknown user: kijitora", ""},
	}
	cv := &sis.Fact{}

	cx++; if Find(nil) != "" { t.Errorf("%s(nil) returns true", fn) }
	for _, e := range ae {
		cv.DiagnosticCode = e[0]
		cv.Command = e[1]

		if Find(cv) == "" { t.Errorf("%s(%s) returns empty", fn, e[0][0:10]) }
	}
	t.Logf("The number of tests = %d", cx)
}

