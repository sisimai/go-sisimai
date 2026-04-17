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

func TestInquire53(t *testing.T) {
	en := "Sendmail"
	fn := "lhost.InquireFor[" + en + "]"
	ae := []string{
		"lhost-sendmail-01", "lhost-sendmail-02", "lhost-sendmail-03", "lhost-sendmail-04", "lhost-sendmail-05", 
		"lhost-sendmail-06", "lhost-sendmail-07", "lhost-sendmail-08", "lhost-sendmail-09", "lhost-sendmail-10",
		"lhost-sendmail-11", "lhost-sendmail-12", "lhost-sendmail-13", "lhost-sendmail-15", "lhost-sendmail-16",
		"lhost-sendmail-17", "lhost-sendmail-18", "lhost-sendmail-19", "lhost-sendmail-20", "lhost-sendmail-21",
		"lhost-sendmail-22", "lhost-sendmail-24", "lhost-sendmail-25", "lhost-sendmail-26", "lhost-sendmail-27",
		"lhost-sendmail-28", "lhost-sendmail-29", "lhost-sendmail-30", "lhost-sendmail-31", "lhost-sendmail-32",
		"lhost-sendmail-33", "lhost-sendmail-34", "lhost-sendmail-35", "lhost-sendmail-36", "lhost-sendmail-37",
		"lhost-sendmail-38", "lhost-sendmail-39", "lhost-sendmail-40", "lhost-sendmail-41", "lhost-sendmail-42",
		"lhost-sendmail-43", "lhost-sendmail-44", "lhost-sendmail-45", "lhost-sendmail-46", "lhost-sendmail-47",
		"lhost-sendmail-48", "lhost-sendmail-49", "lhost-sendmail-50", "lhost-sendmail-51", "lhost-sendmail-52",
		"lhost-sendmail-53", "lhost-sendmail-54", "lhost-sendmail-55", "lhost-sendmail-56", "lhost-sendmail-57",
		"lhost-sendmail-58", "lhost-sendmail-59", "lhost-sendmail-60", 
		"lhost-sendmail-14",
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

		if e == "lhost-sendmail-14" { continue } // TODO: lhost-sendmail-14 returns nil

		cv = InquireFor[en](bf)
		cx++; if cv == nil                    { t.Fatalf("%s(%s) returns nil", fn, e) }
		cx++; if len(cv.Digest) < 1           { t.Errorf("%s(%s).Digest is empty", fn, e) }
		cx++; if cv.Digest[0].Agent     != "" { t.Errorf("%s(%s).Digest.Agent is not empty", fn, e) }
		cx++; if cv.Digest[0].Recipient == "" { t.Errorf("%s(%s).Digest.Recipient is empty", fn, e) }
		cx++; if cv.RFC822 == ""              { t.Errorf("%s(%s).RFC822 is empty (%d)", fn, e, len(cv.RFC822)) }
	}

	t.Logf("The number of tests = %d", cx)
}

