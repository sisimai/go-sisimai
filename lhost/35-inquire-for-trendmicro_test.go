// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
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
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/rfc5322"

func TestInquire35(t *testing.T) {
	en := "TrendMicro"
	fn := "lhost.InquireFor[" + en + "]"
	ae := []string{
		"lhost-trendmicro-01", "lhost-trendmicro-02", "lhost-trendmicro-03", 
	}
	cv := InquireFor[en](nil) 
	cx := 0
	cx++; if cv != nil { t.Errorf("%s(nil) did not return nil", fn) }

	for _, e := range ae {
		ef    := filepath.Join("..", "set-of-emails", "maildir", "bsd", e); ef += ".eml"
		eb, _ := os.ReadFile(ef); ee := string(eb)
		eo, _ := mail.ReadMessage(strings.NewReader(ee))
		bo, _ := io.ReadAll(eo.Body)
		bf    := &siba.BeforeFact{
			Headers: rfc5322.Headers(&eo.Header),
			Payload: bo,
		}

		cv = InquireFor[en](bf)
		cx++; if cv == nil                    { t.Fatalf("%s(%s) returns nil", fn, e) }
		cx++; if len(cv.Digest) < 1           { t.Errorf("%s(%s).Digest is empty", fn, e) }
		cx++; if cv.Digest[0].Agent     != "" { t.Errorf("%s(%s).Digest.Agent is not empty", fn, e) }
		cx++; if cv.Digest[0].Recipient == "" { t.Errorf("%s(%s).Digest.Recipient is empty", fn, e) }
		cx++; if cv.RFC822 != ""              { t.Errorf("%s(%s).RFC822 is not empty (%d)", fn, e, len(cv.RFC822)) }
	}

	t.Logf("The number of tests = %d", cx)
}

