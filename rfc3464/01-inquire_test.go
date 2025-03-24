// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc3464

//  _____         _      ______  _____ ____ _____ _  _    __   _  _   
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___ /| || |  / /_ | || |  
//   | |/ _ \/ __| __| / /| |_) | |_ | |     |_ \| || |_| '_ \| || |_ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__   _| (_) |__   _|
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/   |_|  \___/   |_|  
import "testing"
import "os"
import "io"
import "strings"
import "net/mail"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc5322"

func TestInquire(t *testing.T) {
	fn := "rfc3464.Inquire"
	ae := []string{
		"rfc3464-01", "rfc3464-03", "rfc3464-04", "rfc3464-06", "rfc3464-07", "rfc3464-08", "rfc3464-09",
		"rfc3464-10", "rfc3464-26", "rfc3464-28", "rfc3464-29", "rfc3464-34", "rfc3464-40", "rfc3464-43",
		"rfc3464-51", "rfc3464-53", "rfc3464-55", "rfc3464-56", "rfc3464-57", "rfc3464-58", "rfc3464-59",
		"rfc3464-60", "rfc3464-61", "rfc3464-62", "rfc3464-63", "rfc3464-64", "rfc3464-65", 
		"lhost-powermta-01", "lhost-powermta-02", "lhost-powermta-03",
	}
	cx := 0
	cv := Inquire(nil) 
	cx++; if cv != nil { t.Errorf("%s(nil) did not return nil", fn) }

	for _, e := range ae {
		ef := "../set-of-emails/maildir/bsd/" + e + ".eml"; eb, _ := os.ReadFile(ef); ee := string(eb)
		eo, _ := mail.ReadMessage(strings.NewReader(ee))
		bo, _ := io.ReadAll(eo.Body)
		bf    := &sis.BeforeFact{
			Headers: rfc5322.Headers(&eo.Header, false),
			Payload: string(bo),
		}

		cv = Inquire(bf)
		cx++; if cv == nil                    { t.Errorf("%s(%s) returns nil", fn, e) }
		cx++; if len(cv.Digest) < 1           { t.Errorf("%s(%s).Digest is empty", fn, e) }
		cx++; if cv.Digest[0].Agent     != "" { t.Errorf("%s(%s).Digest.Agent is not empty", fn, e) }
		cx++; if cv.Digest[0].Recipient == "" { t.Errorf("%s(%s).Digest.Recipient is empty", fn, e) }
		cx++; if cv.RFC822 == ""              { t.Errorf("%s(%s).RFC822 is empty", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

