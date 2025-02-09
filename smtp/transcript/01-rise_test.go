// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package transcript

//  _____         _      __             _           ___                                 _       _   
// |_   _|__  ___| |_   / /__ _ __ ___ | |_ _ __   / / |_ _ __ __ _ _ __  ___  ___ _ __(_)_ __ | |_ 
//   | |/ _ \/ __| __| / / __| '_ ` _ \| __| '_ \ / /| __| '__/ _` | '_ \/ __|/ __| '__| | '_ \| __|
//   | |  __/\__ \ |_ / /\__ \ | | | | | |_| |_) / / | |_| | | (_| | | | \__ \ (__| |  | | |_) | |_ 
//   |_|\___||___/\__/_/ |___/_| |_| |_|\__| .__/_/   \__|_|  \__,_|_| |_|___/\___|_|  |_| .__/ \__|
//                                         |_|                                           |_|        
import "testing"
import "strings"
import "os"
import "libsisimai.org/sisimai/smtp/reply"
import "libsisimai.org/sisimai/smtp/status"
import sisimoji "libsisimai.org/sisimai/string"

func TestRise(t *testing.T) {
	fn := "sisimai/smtp/transcript.Rise"
	ef := "../../../set-of-emails/maildir/bsd/lhost-postfix-75.eml"
	cc := []string{"CONN", "HELO", "EHLO", "AUTH", "MAIL", "RCPT", "DATA", "QUIT", "RSET", "XFORWARD"}
	cx := 0

	bx, ce := os.ReadFile(ef); if len(bx) == 0 || ce != nil {
		cx++; t.Fatalf("os.ReadFile(%s) returns error: %s", ef, ce)
	}
	cw := string(bx); cw = cw[strings.Index(cw, "\n\n") + 2:]
	ct := Rise(cw, "In:", "Out:")

	cx++; if len(ct) == 0 { t.Errorf("%s() returns empty", fn) }
	fn = "TanscriptLog"; for _, e := range ct {
		if e.Command == "" {
			cx++; if e.Void() == true { t.Errorf("%s.Void() returns true", fn) }
		} else {
			cx++; if cv := e.Command; sisimoji.EqualsAny(cv, cc) == false {
				t.Errorf("%s.Command(%s) is not listed in %v", fn, cv, cc)
			}
		}

		if e.Command == "MAIL" || e.Command == "RCPT" {
			cx++; if strings.Contains(e.Argument, "@") == false { t.Errorf("%s.Argument does not include @: %s", fn, e.Argument) }
		} else {
			cx++; if e.Argument != "" { t.Errorf("%s.Argument s not empty: %s", fn, e.Argument) }
		}

		for f := range e.Parameter {
			cx++; if e.Parameter[f] == "" { t.Errorf("%s.Parameter[%s] is empty", fn, f) }
		}

		cx++; if e.Response.Reply != "" && reply.Test(e.Response.Reply)   == false {
			t.Errorf("%s.Response.Reply is invalid: (%s)", fn, e.Response.Reply)
		}
		cx++; if e.Response.Status != "" && status.Test(e.Response.Status) == false {
			t.Errorf("%s.Response.Status is invalid: (%s)", fn, e.Response.Status)
		}

		for j, f := range e.Response.Text {
			cx++; if cv := reply.Find(f, ""); cv == "" {
				t.Errorf("%s.Response.Text[%d] is empty", fn, j)
			}
		}
	}

	t.Logf("The number of tests = %d", cx)
}

