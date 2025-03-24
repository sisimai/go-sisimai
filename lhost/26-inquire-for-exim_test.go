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
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"

func TestInquire26(t *testing.T) {
	en := "Exim"
	fn := "lhost.InquireFor[" + en + "]"
	ae := []string{
		"lhost-exim-01", "lhost-exim-02", "lhost-exim-03", "lhost-exim-04", "lhost-exim-05",
		"lhost-exim-06", "lhost-exim-07", "lhost-exim-08", "lhost-exim-29", "lhost-exim-30",
		"lhost-exim-31", "lhost-exim-32", "lhost-exim-33", "lhost-exim-34", "lhost-exim-35",
		"lhost-exim-36", "lhost-exim-37", "lhost-exim-38", "lhost-exim-39", "lhost-exim-40",
		"lhost-exim-41", "lhost-exim-42", "lhost-exim-43", "lhost-exim-44", "lhost-exim-45",
		"lhost-exim-46", "lhost-exim-47", "lhost-exim-48", "lhost-exim-49", "lhost-exim-50",
		"lhost-exim-51", "lhost-exim-52", "lhost-exim-53", "lhost-exim-54", "lhost-exim-55",
		"lhost-exim-56", "lhost-exim-57", "lhost-exim-58", "lhost-exim-59", "lhost-exim-60",
		"lhost-exim-61", 
	}
	cv := InquireFor[en](nil) 
	cx := 0
	cx++; if cv != nil { t.Errorf("%s(nil) did not return nil", fn) }

	for _, e := range ae {
		ef := "../set-of-emails/maildir/bsd/" + e + ".eml"; eb, _ := os.ReadFile(ef); ee := string(eb)
		eo, _ := mail.ReadMessage(strings.NewReader(ee))
		bo, _ := io.ReadAll(eo.Body)
		bf    := &sis.BeforeFact{
			Headers: rfc5322.Headers(&eo.Header, false),
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

