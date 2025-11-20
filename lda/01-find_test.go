// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lda

//  _____         _      ___     ____    _    
// |_   _|__  ___| |_   / / |   |  _ \  / \   
//   | |/ _ \/ __| __| / /| |   | | | |/ _ \  
//   | |  __/\__ \ |_ / / | |___| |_| / ___ \ 
//   |_|\___||___/\__/_/  |_____|____/_/   \_\
import "testing"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"

func TestFind(t *testing.T) {
	fn := "lda.Find"
	cx := 0
	ae := [][]string{
		[]string{`x-unix; procmail: Couldn't create "/var/spool/mail/neko" id:`, eb.CeDATA},
		[]string{"vdelivermail: account is locked email bounced kijitora@neko2.example.co.jp", eb.CeDATA},
		[]string{"maildrop: maildir over quota.", ""},
		[]string{`"|IFS=' ' && exec /usr/local/bin/procmail -f- || exit 75 #kijitora"`, eb.CeDATA},
		[]string{`554 "|IFS=' ' && exec /usr/local/bin/procmail -f- || exit 75 #kijitora"... Service unavailable`, eb.CeDATA},
		[]string{"mail.local: unknown user: kijitora", ""},
	}
	cv := &siba.Fact{}

	cx++; if Find(nil) != "" { t.Errorf("%s(nil) returns true", fn) }
	for _, e := range ae {
		cv.DiagnosticCode = e[0]
		cv.Command = e[1]

		if Find(cv) == "" { t.Errorf("%s(%s) returns empty", fn, e[0][0:10]) }
	}
	t.Logf("The number of tests = %d", cx)
}

