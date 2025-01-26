// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      __   _       _   _       _   ____                     _          _ 
// |_   _|__  ___| |_   / /__(_)___  | \ | | ___ | |_|  _ \  ___  ___ ___   __| | ___  __| |
//   | |/ _ \/ __| __| / / __| / __| |  \| |/ _ \| __| | | |/ _ \/ __/ _ \ / _` |/ _ \/ _` |
//   | |  __/\__ \ |_ / /\__ \ \__ \_| |\  | (_) | |_| |_| |  __/ (_| (_) | (_| |  __/ (_| |
//   |_|\___||___/\__/_/ |___/_|___(_)_| \_|\___/ \__|____/ \___|\___\___/ \__,_|\___|\__,_|
import "testing"
import "strings"

func TestMakeNotDecoded(t *testing.T) {
	cc := "NotDecoded"
	fn := "MakeNotDecoded"
	cv := MakeNotDecoded("Test message", true)
	cx := 0

	cx++; if cv == nil                { t.Fatalf("%s() = nil", fn) }
	cx++; if cv.EmailFile != ""       { t.Errorf("%s.EmailFile is not empty: %s", cc, cv.EmailFile) }
	cx++; if cv.BecauseOf == ""       { t.Errorf("%s.BecauseOf is empty", cc) }
	cx++; if cv.Timestamp.Unix() == 0 { t.Errorf("%s.Timestamp.Unix() is 0", cc) }
	cx++; if cv.CalledOff == false    { t.Errorf("%s.CalledOff is false", cc) }
	cx++; if cv.DecodedBy != ""       { t.Errorf("%s.DecodedBy is not empty: %s", cc, cv.DecodedBy) }
	cx++; if strings.Contains(cv.WhoCalled, fn) == false { t.Errorf("%s.WhoCalled is %s", cc, cv.WhoCalled) }
	t.Logf("The number of tests = %d", cx)
}

func TestError(t *testing.T) {
	fn := "Error"
	cv := &NotDecoded{BecauseOf: "Test for " + fn, EmailFile: "25.eml"}
	cw := &NotDecoded{BecauseOf: "", EmailFile: "25.eml"}
	cx := 0

	cx++; if cv == nil        { t.Fatalf("%s() = nil", fn) }
	cx++; if cv.Error() == "" { t.Fatalf("%s() = empty string", fn) }
	cx++; if cw.Error() != "" { t.Fatalf("%s() = string %s", fn, cw.Error()) }
	t.Logf("The number of tests = %d", cx)
}

func TestLabel(t *testing.T) {
	fn := "Label"
	cv := &NotDecoded{BecauseOf: "Test for " + fn, CalledOff: true}
	cw := &NotDecoded{BecauseOf: "Test for " + fn, CalledOff: false}
	cx := 0

	cx++; if cv == nil                { t.Fatalf("%s() = nil", fn) }
	cx++; if strings.HasPrefix(cv.Label(), " *****error") == false { t.Errorf("%s() = (%s)", fn, cv.Label()) }
	cx++; if strings.HasPrefix(cw.Label(), " ***warning") == false { t.Errorf("%s() = (%s)", fn, cw.Label()) }
	t.Logf("The number of tests = %d", cx)
}

func TestEmail(t *testing.T) {
	fn := "Email"
	cv := &NotDecoded{BecauseOf: "Test for " + fn}
	cx := 0

	cx++; if cv == nil                { t.Fatalf("%s() = nil", fn) }
	cx++; if cv.Email("22.eml") == "" { t.Errorf("%s() = empty string", fn) }
	cx++; if cv.EmailFile != "22.eml" { t.Errorf("%s() = (%s)", fn, cv.EmailFile) }
	cx++; if cv.Email("25") == "25"   { t.Errorf("%s() = 25", fn) }
	t.Logf("The number of tests = %d", cx)
}

