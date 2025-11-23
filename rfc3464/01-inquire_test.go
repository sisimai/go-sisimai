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
import "path/filepath"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/rfc5322"

func TestInquire(t *testing.T) {
	fn := "rfc3464.Inquire"
	ae := []string{
		"rfc3464-01", "rfc3464-03", "rfc3464-04", "rfc3464-06", "rfc3464-07", "rfc3464-08", "rfc3464-09",
		"rfc3464-10", "rfc3464-26", "rfc3464-28", "rfc3464-29", "rfc3464-34", "rfc3464-40", "rfc3464-43",
		"rfc3464-51", "rfc3464-53", "rfc3464-55", "rfc3464-56", "rfc3464-57", "rfc3464-58", "rfc3464-59",
		"rfc3464-60", "rfc3464-61", "rfc3464-62", "rfc3464-63", "rfc3464-64", "rfc3464-65", "rfc3464-66",
		"lhost-powermta-01", "lhost-powermta-02", "lhost-powermta-03",

		"rfc3464-35", "rfc3464-36", "rfc3464-37", "rfc3464-38", "rfc3464-39", "rfc3464-42", "rfc3464-52",
		"rfc3464-54",
	}
	cx := 0
	cv := Inquire(nil) 
	cx++; if cv != nil { t.Errorf("%s(nil) did not return nil", fn) }

	for _, e := range ae {
		ef    := filepath.Join("..", "set-of-emails", "maildir", "bsd", e); ef += ".eml"
		eb, _ := os.ReadFile(ef); ee := string(eb)
		eo, _ := mail.ReadMessage(strings.NewReader(ee))
		bo, _ := io.ReadAll(eo.Body)
		bf    := &siba.BeforeFact{
			Headers: rfc5322.Headers(&eo.Header),
			Payload: string(bo),
		}

		if e == "rfc3464-35" || e == "rfc3464-36" || e == "rfc3464-37" || e == "rfc3464-38" ||
		   e == "rfc3464-39" || e == "rfc3464-42" || e == "rfc3464-52" || e == "rfc3464-54" ||
		   e == "rfc3464-66" {
			// TODO:
			// - rfc3464-35 returns an empty RFC822 part
			// - rfc3464-36 returns an empty RFC822 part
			// - rfc3464-37 returns nil
			// - rfc3464-38 returns nil
			// - rfc3464-39 returns nil
			// - rfc3464-42 returns an empty RFC822 part
			// - rfc3464-52 returns an empty RFC822 part
			// - rfc3464-53 returns an empty RFC822 part
			// - rfc3464-64 returns an empty RFC822 part
			continue
		}

		cv = Inquire(bf)
		cx++; if cv == nil                    { t.Fatalf("%s(%s) returns nil", fn, e) }
		cx++; if len(cv.Digest) < 1           { t.Errorf("%s(%s).Digest is empty", fn, e) }
		cx++; if cv.Digest[0].Agent     != "" { t.Errorf("%s(%s).Digest.Agent is not empty", fn, e) }
		cx++; if cv.Digest[0].Recipient == "" { t.Errorf("%s(%s).Digest.Recipient is empty", fn, e) }
		cx++; if cv.RFC822 == ""              { t.Errorf("%s(%s).RFC822 is empty", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

func TestReturnedBy(t *testing.T) {
	fn := "rfc3464.returnedby"
	cx := 0

	cx++; if cv := returnedby("");  cv != "" { t.Errorf("%s(%s) returns %s", fn, "", cv)  }
	cx++; if cv := returnedby("2"); cv != "" { t.Errorf("%s(%s) returns %s", fn, "1", cv) }
	t.Logf("The number of tests = %d", cx)
}

func TestXfield(t *testing.T) {
	fn := "rfc3464.xfield"
	cx := 0

	cx++; if cv := xfield("");  len(cv) != 0 { t.Errorf("%s(%s) returns %+v", fn, "", cv)  }
	cx++; if cv := xfield("2"); len(cv) != 0 { t.Errorf("%s(%s) returns %+v", fn, "2", cv) }
	t.Logf("The number of tests = %d", cx)
}

