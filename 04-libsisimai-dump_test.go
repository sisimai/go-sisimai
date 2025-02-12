// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sisimai

//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
import "testing"
import "strings"

func TestDump(t *testing.T) {
	fn := "sisimai.Dump"
	cx := 0

	rootdir := "set-of-emails/"
	samples := []string{"mailbox/mbox-0", "mailbox/mbox-1", "maildir/bsd"}
	normals := []string{"maildir/not"}
	sisiarg := Args(); sisiarg.Delivered = true; sisiarg.Vacation = true

	for _, e := range samples {
		ef := "./" + rootdir + e
		cv, _ := Dump(ef, sisiarg)
		cx++; if cv == nil || len(*cv) == 0 { t.Errorf("%s(%s) returns empty", fn, ef) }
		cx++; if strings.HasPrefix(*cv, "[{") == false { t.Errorf("%s(%s) returns invalid JSON string", fn, ef) }
		cx++; if strings.HasSuffix(*cv, "}]") == false { t.Errorf("%s(%s) returns invalid JSON string", fn, ef) }
	}

	for _, e := range normals {
		ef := "./" + rootdir + e
		cv, _ := Dump(ef, sisiarg)
		cx++; if cv != nil { t.Errorf("%s(%s) returns results: %v", fn, ef, *cv) }
	}

	t.Logf("The number of tests = %d", cx)
}

