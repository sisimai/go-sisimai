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
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/rfc5322"

func TestInquire22(t *testing.T) {
	en := "DragonFly"
	fn := "lhost.InquireFor[" + en + "]"
	ae := []string{
		"lhost-dragonfly-01", "lhost-dragonfly-02", "lhost-dragonfly-03", "lhost-dragonfly-04",
		"lhost-dragonfly-05", "lhost-dragonfly-06", "lhost-dragonfly-07", "lhost-dragonfly-08",
		"lhost-dragonfly-09", "lhost-dragonfly-10", "lhost-dragonfly-11", "lhost-dragonfly-12",
		"lhost-dragonfly-13", "lhost-dragonfly-14", "lhost-dragonfly-15", "lhost-dragonfly-16",
		"lhost-dragonfly-17", "lhost-dragonfly-18", "lhost-dragonfly-19", "lhost-dragonfly-20",
		"lhost-dragonfly-21", "lhost-dragonfly-22", "lhost-dragonfly-23", "lhost-dragonfly-24",
		"lhost-dragonfly-25", "lhost-dragonfly-26", "lhost-dragonfly-27", "lhost-dragonfly-28",
		"lhost-dragonfly-29", "lhost-dragonfly-30", 
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

