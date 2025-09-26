// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _____         _      ___ _               _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|
import "testing"
import "os"
import "io"
import "strings"
import "net/mail"
import "path/filepath"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/rfc5322"

func TestInquire32(t *testing.T) {
	en := "GoogleGroups"
	fn := "lhost.InquireFor[" + en + "]"
	ae := []string{
		"lhost-googlegroups-01", "lhost-googlegroups-02", "lhost-googlegroups-03", "lhost-googlegroups-04",
		"lhost-googlegroups-05", "lhost-googlegroups-06", "lhost-googlegroups-07", "lhost-googlegroups-08",
		"lhost-googlegroups-09", "lhost-googlegroups-10", "lhost-googlegroups-11", "lhost-googlegroups-12",
		"lhost-googlegroups-13", "lhost-googlegroups-14", 
	}
	cv := InquireFor[en](nil) 
	cx := 0
	cx++; if cv != nil { t.Errorf("%s(nil) did not return nil", fn) }

	for _, e := range ae {
		ef    := filepath.Join("..", "set-of-emails", "maildir", "bsd", e); ef += ".eml"
		eb, _ := os.ReadFile(ef); ee := string(eb)
		eo, _ := mail.ReadMessage(strings.NewReader(ee))
		bo, _ := io.ReadAll(eo.Body)
		bf    := &sis.BeforeFact{
			Headers: rfc5322.Headers(&eo.Header),
			Payload: string(bo),
		}

		cv = InquireFor[en](bf)
		cx++; if cv == nil                    { t.Errorf("%s(%s) returns nil", fn, e) }
		cx++; if len(cv.Digest) < 1           { t.Errorf("%s(%s).Digest is empty", fn, e) }
		cx++; if cv.Digest[0].Agent     != "" { t.Errorf("%s(%s).Digest.Agent is not empty", fn, e) }
		cx++; if cv.Digest[0].Recipient == "" { t.Errorf("%s(%s).Digest.Recipient is empty", fn, e) }
		cx++; if cv.RFC822 == ""              { t.Errorf("%s(%s).RFC822 is empty (%d)", fn, e, len(cv.RFC822)) }
	}

	t.Logf("The number of tests = %d", cx)
}

