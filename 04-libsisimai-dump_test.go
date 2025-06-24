// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sisimai

//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
import "testing"
import "os"
import "os/exec"
import "strings"

func TestDump(t *testing.T) {
	fn := "sisimai.Dump"
	cx := 0

	rootdir := "set-of-emails/"
	samples := []string{"mailbox/mbox-0", "mailbox/mbox-1", "maildir/bsd"}
	normals := []string{"maildir/not"}
	notfile := []string{"/dev/null", "/dev/neko"}
	isempty := "/tmp/empty-file-for-test-of-sisimai"
	sisiarg := Args(); sisiarg.Delivered = true; sisiarg.Vacation = true

	for _, e := range samples {
		ef := "./" + rootdir + e
		cv, _ := Dump(ef, sisiarg)
		cx++; if cv == nil || len(*cv) == 0 { t.Errorf("%s(%s) returns empty", fn, ef) }
		cx++; if strings.HasPrefix(*cv, "[{") == false { t.Errorf("%s(%s) returns invalid JSON string", fn, ef) }
		cx++; if strings.HasSuffix(*cv, "}]") == false { t.Errorf("%s(%s) returns invalid JSON string", fn, ef) }

		// When the 2nd argument is nil
		cv, _  = Dump(ef, nil)
		cx++; if cv == nil || len(*cv) == 0 { t.Errorf("%s(%s, nil) returns empty", fn, ef) }
	}

	for _, e := range normals {
		ef := "./" + rootdir + e
		cv, _ := Dump(ef, sisiarg)
		cx++; if cv != nil { t.Errorf("%s(%s) returns results: %v", fn, ef, *cv) }

		// When the 2nd argument is nil
		cv, _  = Dump(ef, nil)
		cx++; if cv != nil { t.Errorf("%s(%s, nil) returns results: %v", fn, ef, *cv) }
	}

	for _, e := range notfile {
		cv, ce := Rise(e, sisiarg)
		cx++; if len(cv) != 0 { t.Errorf("%s(%s) returns results: %v", fn, e, cv) }
		cx++; if len(ce) == 0 { t.Errorf("%s(%s) returns an empty error", fn, e) }

		cv, _   = Rise(e, nil)
		cx++; if len(cv) != 0 { t.Errorf("%s(%s, nil) returns results: %v", fn, e, cv) }
	}

	comm := exec.Command("touch", isempty); nyaan := comm.Run()
	if nyaan == nil {
		cv, ce := Rise(isempty, sisiarg)
		cx++; if len(cv) != 0 { t.Errorf("%s(%s) returns results: %v", fn, isempty, cv) }
		cx++; if len(ce) == 0 { t.Errorf("%s(%s) returns an empty error", fn, isempty) }

		cv, _   = Rise(isempty, nil)
		cx++; if len(cv) != 0 { t.Errorf("%s(%s, nil) returns results: %v", fn, isempty, cv) }

		os.Remove(isempty)
	}

	t.Logf("The number of tests = %d", cx)
}

