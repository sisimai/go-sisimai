// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      __   _       ____      _ _ _                _        _                  
// |_   _|__  ___| |_   / /__(_)___  / ___|__ _| | | |__   __ _  ___| | __   / \   _ __ __ _ ___ 
//   | |/ _ \/ __| __| / / __| / __|| |   / _` | | | '_ \ / _` |/ __| |/ /  / _ \ | '__/ _` / __|
//   | |  __/\__ \ |_ / /\__ \ \__ \| |__| (_| | | | |_) | (_| | (__|   <  / ___ \| | | (_| \__ \
//   |_|\___||___/\__/_/ |___/_|___(_)____\__,_|_|_|_.__/ \__,_|\___|_|\_\/_/   \_\_|  \__, |___/
//                                                                                     |___/     
import "testing"

func TestCallbackArg0(t *testing.T) {
	cc := "CallbackArg0"
	fn := "sis.CallbackArg0"
	cv := &CallbackArg0{
		Headers: map[string][]string{"Nekochan": []string{"Kijitora", "Michistuna"}},
		Payload: &cc,
	}
	cx := 0

	cx++; if cv == nil                        { t.Fatalf("%s{} = nil", fn) }
	cx++; if len(cv.Headers["Nekochan"]) != 2 { t.Errorf("%s.Headers have not 2 elements", fn) }
	cx++; if len(*cv.Payload)            == 0 { t.Errorf("%s.Payload is empty", fn) }
	t.Logf("The number of tests = %d", cx)
}

func TestCallbackArg1(t *testing.T) {
	cc := "CallbackArg1"
	fn := "sis.CallbackArg1"
	cv := &CallbackArg1{
		Path: "/var/mail/root",
		Kind: "mailbox",
		Mail: &cc,
		Fact: nil,
	}
	cx := 0

	cx++; if cv == nil                      { t.Fatalf("%s{} = nil", fn) }
	cx++; if cv.Path    != "/var/mail/root" { t.Errorf("%s.Path is %s", fn, cv.Path) }
	cx++; if cv.Kind    != "mailbox"        { t.Errorf("%s.Kind is %s", fn, cv.Kind) }
	cx++; if (*cv.Mail) == ""               { t.Errorf("%s.Mail is empty", fn) }
	cx++; if (*cv).Fact != nil              { t.Errorf("%s.Fact is not nil", fn) }

	t.Logf("The number of tests = %d", cx)
}
