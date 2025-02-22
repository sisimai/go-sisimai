// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string

//  _____         _      __   _        _             
// |_   _|__  ___| |_   / /__| |_ _ __(_)_ __   __ _ 
//   | |/ _ \/ __| __| / / __| __| '__| | '_ \ / _` |
//   | |  __/\__ \ |_ / /\__ \ |_| |  | | | | | (_| |
//   |_|\___||___/\__/_/ |___/\__|_|  |_|_| |_|\__, |
//                                             |___/ 
import "testing"
import "strings"

func TestToPlain(t *testing.T) {
	fn := "sisimai/string.ToPlain"
	cx := 0
	cw := `<html>
        <head><style>h1 { color: red; } </style></head>
        <body>
            <h1>neko</h1>
            <div>
            <a href = "https://libsisimai.org">sisimai</a>
            <a href = "mailto:maketest@libsisimai.org">maketest</a>
            </div>
        </body>
        </html>`

	cv := ToPlain(&cw)
	cx++; if *cv == ""                                  { t.Errorf("%s(...) returns empty",   fn) }
	cx++; if strings.Contains(*cv, "<html>")   == true  { t.Errorf("%s(...) contains <html>", fn) }
	cx++; if strings.Contains(*cv, "<head>")   == true  { t.Errorf("%s(...) contains <head>", fn) }
	cx++; if strings.Contains(*cv, "<body>")   == true  { t.Errorf("%s(...) contains <body>", fn) }
	cx++; if strings.Contains(*cv, "<div>")    == true  { t.Errorf("%s(...) contains <div>",  fn) }
	cx++; if strings.Contains(*cv, "href ")    == true  { t.Errorf("%s(...) contains <div>",  fn) }
	cx++; if strings.Contains(*cv, "sisimai")  == false { t.Errorf("%s(...) does not contain sisimai",  fn) }
	cx++; if strings.Contains(*cv, "maketest") == false { t.Errorf("%s(...) does not contain maketest", fn) }

	ce := ""
	cx++; if cv = ToPlain(nil); cv != nil { t.Errorf("%s(nil) returns %s", fn, *cv) }
	cx++; if cv = ToPlain(&ce); *cv != "" { t.Errorf("%s('') returns %s", fn, *cv) }

	t.Logf("The number of tests = %d", cx)
}

