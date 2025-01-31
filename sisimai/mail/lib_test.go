// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package mail

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   __ _(_) |
//   | |/ _ \/ __| __| / / '_ ` _ \ / _` | | |
//   | |  __/\__ \ |_ / /| | | | | | (_| | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\__,_|_|_|
import "testing"
import "strings"

/*
 | FIELD      | UNIX mbox | Maildir/  | Memory    | <STDIN>    |
 |------------|-----------|-----------|-----------|------------|
 | Kind       | o         | o         | o         | o          |
 | Path       | o         | o         | o         | o          |
 | Dir        | o         | o         |           |            |
 | File       | o         | o         |           |            |
 | Size       | o         | o         | o         | o          |
 | newline    | o         |           | o         | o          |
 | offset     | o         | o         | o         | o          |
 | handle     | o         | o         |           |            |
 | payload    |           | o         | o         | o          |
*/
var Mailbox = []string{
	"../../set-of-emails/mailbox/mbox-0",
	"../../set-of-emails/mailbox/mbox-1",
}
var Maildir = []string{
	"../../set-of-emails/maildir/bsd",
	"../../set-of-emails/maildir/dos",
	"../../set-of-emails/maildir/mac",
	"../../set-of-emails/maildir/err",
}

func TestRise(t *testing.T) {
	fn := "Rise"
	cx := 0

	for _, e := range Mailbox {
		cv, ce := Rise(e)
		cf := "EmailEntity(mailbox)"
		cx++; if cv == nil { t.Errorf("%s(%s) returns nil", fn, e) }
		cx++; if ce != nil { t.Errorf("%s(%s) returns error: %s", fn, e, ce) }
		cx++; if cv.Kind != "mailbox" { t.Errorf("%s.Kind is not mailbox: %s", cf, cv.Kind) }
		cx++; if cv.Path != e         { t.Errorf("%s.Path is not %s", cf, e) }
		cx++; if strings.Contains(cv.File, "mbox-")        == false { t.Errorf("%s.File does not include mbox-: %s", cf, cv.File) }
		cx++; if strings.Contains(cv.Dir, "set-of-emails") == false { t.Errorf("%s.Dir does not include set-of-emails: %s", cf, cv.Dir) }
		cx++; if cv.Size == 0         { t.Errorf("%s.Size is 0", cf) }
		cx++; if cv.newline == 0      { t.Errorf("%s.newline is 0: %d", cf, cv.newline) }
		cx++; if cv.offset > 0        { t.Errorf("%s.offset is not 0: %d", cf, cv.offset) }
		cx++; if len(cv.payload) > 0  { t.Errorf("%s.payload is not 0: %d", cf, len(cv.payload)) }
	}

	for _, e := range Maildir {
		cv, ce:= Rise(e)
		cf := "EmailEntity(maildir)"
		cx++; if cv == nil { t.Errorf("%s(%s) returns nil", fn, e) }
		cx++; if ce != nil { t.Errorf("%s(%s) returns error: %s", fn, e, ce) }
		cx++; if cv.Kind != "maildir" { t.Errorf("%s.Kind is not maildir: %s", cf, cv.Kind) }
		cx++; if cv.Path == ""        { t.Errorf("%s.Path is empty: %s", cf, cv.Path) }
		cx++; if strings.Contains(cv.Dir, "/maildir/") == false { t.Errorf("%s.Dir does not contain /maildir/: %s", cf, cv.Dir) }
		cx++; if cv.Size == 0         { t.Errorf("%s.Size is 0", cf) }
		cx++; if cv.newline != 0      { t.Errorf("%s.newline is not 0: %d", cf, cv.newline) }
		cx++; if cv.offset > 0        { t.Errorf("%s.offset is not 0: %d", cf, cv.offset) }
		cx++; if len(cv.payload) == 0 { t.Errorf("%s.payload is 0", cf) }
	}

	t.Logf("The number of tests = %d", cx)
}

