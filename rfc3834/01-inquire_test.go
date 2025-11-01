// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc3834

//  _____         _      ______  _____ ____ _____  ___ _____ _  _   
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___ / ( _ )___ /| || |  
//   | |/ _ \/ __| __| / /| |_) | |_ | |     |_ \ / _ \ |_ \| || |_ 
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) | (_) |__) |__   _|
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/ \___/____/   |_|  
import "testing"
import "os"
import "io"
import "strings"
import "net/mail"
import "path/filepath"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/rfc5322"

func TestInquire(t *testing.T) {
	fn := "rfc3834.Inquire"
	ae := []string{
		"rfc3834-01", "rfc3834-02", "rfc3834-03", "rfc3834-04", "rfc3834-05", "rfc3834-06",
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

		cv = Inquire(bf)
		cx++; if cv == nil                    { t.Fatalf("%s(%s) returns nil", fn, e) }
		cx++; if len(cv.Digest) < 1           { t.Errorf("%s(%s).Digest is empty", fn, e) }
		cx++; if cv.Digest[0].Agent     != "" { t.Errorf("%s(%s).Digest.Agent is not empty", fn, e) }
		cx++; if cv.Digest[0].Recipient == "" { t.Errorf("%s(%s).Digest.Recipient is empty", fn, e) }
		cx++; if cv.RFC822 == ""              { t.Errorf("%s(%s).RFC822 is empty", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

