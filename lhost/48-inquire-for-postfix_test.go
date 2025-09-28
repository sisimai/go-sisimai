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

func TestInquire48(t *testing.T) {
	en := "Postfix"
	fn := "lhost.InquireFor[" + en + "]"
	ae := []string{
		"lhost-postfix-01", "lhost-postfix-02", "lhost-postfix-03", "lhost-postfix-04", "lhost-postfix-05",
		"lhost-postfix-06", "lhost-postfix-07", "lhost-postfix-08", "lhost-postfix-09", "lhost-postfix-10",
		"lhost-postfix-11", "lhost-postfix-13", "lhost-postfix-14", "lhost-postfix-15", "lhost-postfix-16",
		"lhost-postfix-17", "lhost-postfix-28", "lhost-postfix-29", "lhost-postfix-30", "lhost-postfix-31",
		"lhost-postfix-32", "lhost-postfix-33", "lhost-postfix-34", "lhost-postfix-35", "lhost-postfix-36",
		"lhost-postfix-37", "lhost-postfix-38", "lhost-postfix-39", "lhost-postfix-40", "lhost-postfix-41",
		"lhost-postfix-42", "lhost-postfix-43", "lhost-postfix-44", "lhost-postfix-45", "lhost-postfix-46",
		"lhost-postfix-47", "lhost-postfix-48", "lhost-postfix-49", "lhost-postfix-50", "lhost-postfix-51",
		"lhost-postfix-52", "lhost-postfix-53", "lhost-postfix-54", "lhost-postfix-55", "lhost-postfix-56",
		"lhost-postfix-57", "lhost-postfix-58", "lhost-postfix-59", "lhost-postfix-60", "lhost-postfix-61",
		"lhost-postfix-62", "lhost-postfix-63", "lhost-postfix-64", "lhost-postfix-65", "lhost-postfix-66",
		"lhost-postfix-67", "lhost-postfix-68", "lhost-postfix-69", "lhost-postfix-70", "lhost-postfix-71",
		"lhost-postfix-72", "lhost-postfix-73", "lhost-postfix-74", "lhost-postfix-75", "lhost-postfix-76",
		"lhost-postfix-77", "lhost-postfix-78", 
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
		cx++; if cv == nil                    { t.Fatalf("%s(%s) returns nil", fn, e) }
		cx++; if len(cv.Digest) < 1           { t.Errorf("%s(%s).Digest is empty", fn, e) }
		cx++; if cv.Digest[0].Agent     != "" { t.Errorf("%s(%s).Digest.Agent is not empty", fn, e) }
		cx++; if cv.Digest[0].Recipient == "" { t.Errorf("%s(%s).Digest.Recipient is empty", fn, e) }
		cx++; if cv.RFC822 == ""              { t.Errorf("%s(%s).RFC822 is empty (%d)", fn, e, len(cv.RFC822)) }
	}

	t.Logf("The number of tests = %d", cx)
}

