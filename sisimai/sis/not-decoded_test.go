// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      ___   _       _   ____                     _          _ 
// |_   _|__  ___| |_   / / \ | | ___ | |_|  _ \  ___  ___ ___   __| | ___  __| |
//   | |/ _ \/ __| __| / /|  \| |/ _ \| __| | | |/ _ \/ __/ _ \ / _` |/ _ \/ _` |
//   | |  __/\__ \ |_ / / | |\  | (_) | |_| |_| |  __/ (_| (_) | (_| |  __/ (_| |
//   |_|\___||___/\__/_/  |_| \_|\___/ \__|____/ \___|\___\___/ \__,_|\___|\__,_|
import "testing"
import "strings"

func TestMakeNotDecoded(t *testing.T) {
	cc := "NotDecoded"
	fn := "MakeNotDecoded"
	cv := MakeNotDecoded("Test message", true)

	if cv == nil                { t.Fatalf("%s() = nil", fn) }
	if cv.EmailFile != ""       { t.Errorf("%s.EmailFile is not empty: %s", cc, cv.EmailFile) }
	if cv.BecauseOf == ""       { t.Errorf("%s.BecauseOf is empty", cc) }
	if cv.Timestamp.Unix() == 0 { t.Errorf("%s.Timestamp.Unix() is 0", cc) }
	if cv.CalledOff == false    { t.Errorf("%s.CalledOff is false", cc) }
	if cv.DecodedBy != ""       { t.Errorf("%s.DecodedBy is not empty: %s", cc, cv.DecodedBy) }
	if strings.Contains(cv.WhoCalled, fn) == false { t.Errorf("%s.WhoCalled is %s", cc, cv.WhoCalled) }
}

func TestError(t *testing.T) {
	fn := "Error"
	cv := &NotDecoded{BecauseOf: "Test for " + fn, EmailFile: "25.eml"}
	cx := &NotDecoded{BecauseOf: "", EmailFile: "25.eml"}

	if cv == nil                { t.Fatalf("%s() = nil", fn) }
	if cv.Error() == ""         { t.Fatalf("%s() = empty string", fn) }
	if cx.Error() != ""         { t.Fatalf("%s() = string %s", fn, cx.Error()) }
}

func TestLabel(t *testing.T) {
	fn := "Label"
	cv := &NotDecoded{BecauseOf: "Test for " + fn, CalledOff: true}
	cx := &NotDecoded{BecauseOf: "Test for " + fn, CalledOff: false}

	if cv == nil                { t.Fatalf("%s() = nil", fn) }
	if strings.HasPrefix(cv.Label(), " *****error") == false { t.Errorf("%s() = (%s)", fn, cv.Label()) }
	if strings.HasPrefix(cx.Label(), " ***warning") == false { t.Errorf("%s() = (%s)", fn, cx.Label()) }
}

func TestEmail(t *testing.T) {
	fn := "Email"
	cv := &NotDecoded{BecauseOf: "Test for " + fn}

	if cv == nil                { t.Fatalf("%s() = nil", fn) }
	if cv.Email("22.eml") == "" { t.Errorf("%s() = empty string", fn) }
	if cv.EmailFile != "22.eml" { t.Errorf("%s() = (%s)", fn, cv.EmailFile) }
	if cv.Email("25") == "25"   { t.Errorf("%s() = 25", fn) }
}

