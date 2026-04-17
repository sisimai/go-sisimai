// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package moji

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   ___ (_|_)
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \| | |
//   | |  __/\__ \ |_ / /| | | | | | (_) | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\___// |_|
//                                     |__/   
import "testing"
import "bytes"

func TestToPlain(t *testing.T) {
	fn := "moji.ToPlain"
	cx := 0
	cw := []byte(`<html>
        <head><style>h1 { color: red; } </style></head>
        <body>
            <h1>neko</h1>
            <div>
            <a href = "https://libsisimai.org">sisimai</a>
            <a href = "mailto:maketest@libsisimai.org">maketest</a>
            </div>
        </body>
        </html>`)

	cv := ToPlain(cw)
	cx++; if len(cv) == 0                                    { t.Errorf("%s(...) returns empty",   fn) }
	cx++; if bytes.Contains(cv, []byte("<html>"))   == true  { t.Errorf("%s(...) contains <html>", fn) }
	cx++; if bytes.Contains(cv, []byte("<head>"))   == true  { t.Errorf("%s(...) contains <head>", fn) }
	cx++; if bytes.Contains(cv, []byte("<body>"))   == true  { t.Errorf("%s(...) contains <body>", fn) }
	cx++; if bytes.Contains(cv, []byte("<div>"))    == true  { t.Errorf("%s(...) contains <div>",  fn) }
	cx++; if bytes.Contains(cv, []byte("href "))    == true  { t.Errorf("%s(...) contains <div>",  fn) }
	cx++; if bytes.Contains(cv, []byte("sisimai"))  == false { t.Errorf("%s(...) does not contain sisimai",  fn) }
	cx++; if bytes.Contains(cv, []byte("maketest")) == false { t.Errorf("%s(...) does not contain maketest", fn) }

	ce := []byte("<html></html>")
	cx++; if cv = ToPlain(ce); len(cv) == 0 { t.Errorf("%s(%s) returns %s", fn, ce, cv) }

	ce  = []byte("<html><body style = ''></body></html>")
	cx++; if cv = ToPlain(ce); len(cv) > 0 { t.Errorf("%s(%s) returns %s", fn, ce, cv) }

	ce  = []byte("")
	cx++; if cv = ToPlain(nil); cv != nil  { t.Errorf("%s(nil) returns %s", fn, cv) }
	cx++; if cv = ToPlain(ce); len(cv) > 0 { t.Errorf("%s('') returns %s", fn,  cv) }

	t.Logf("The number of tests = %d", cx)
}

