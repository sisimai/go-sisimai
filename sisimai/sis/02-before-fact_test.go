// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      __   _       ____        __                _____          _   
// |_   _|__  ___| |_   / /__(_)___  | __ )  ___ / _| ___  _ __ ___|  ___|_ _  ___| |_ 
//   | |/ _ \/ __| __| / / __| / __| |  _ \ / _ \ |_ / _ \| '__/ _ \ |_ / _` |/ __| __|
//   | |  __/\__ \ |_ / /\__ \ \__ \_| |_) |  __/  _| (_) | | |  __/  _| (_| | (__| |_ 
//   |_|\___||___/\__/_/ |___/_|___(_)____/ \___|_|  \___/|_|  \___|_|  \__,_|\___|\__|
import "testing"

// Sender  string              // Unix FROM line ("From ")
// Headers map[string][]string // Email headers
// Payload string              // Email body
// RFC822  map[string][]string // Email headers of the original message
// Digest  []DeliveryMatter    // Decoded results returned from sisimai/lhost/*
// Catch   interface{}         // Any data structure returned by the callback function
// Errors  []NotDecoded        // All the errors and warnings
func TestBeforeFact(t *testing.T) {
	fn := "sis.BeforeFact"
	cv := &BeforeFact{
		Sender:  "From <mailer-daemon@example.jp>",
		Headers: map[string][]string{"Subject": []string{"Delivery Failure"}},
		Payload: "Sorry, the email delivery failed",
		RFC822:  map[string][]string{"To": []string{"<postmaster@example.org>"}},
		Digest:  []DeliveryMatter{DeliveryMatter{Action: "failed"}},
		Catch:   nil,
		Errors:  []NotDecoded{*(MakeNotDecoded("Test message", true))},
	}
	cx := 0

	cx++; if cv == nil                         { t.Fatalf("%s{} = nil", fn) }
	cx++; if len(cv.Sender)             == 0   { t.Errorf("%s.Sender is empty", fn) }
	cx++; if len(cv.Headers["Subject"]) != 1   { t.Errorf("%s.Headers[Subject] have not 1 element", fn) }
	cx++; if len(cv.Payload)            == 0   { t.Errorf("%s.Payload is empty", fn) }
	cx++; if len(cv.RFC822["To"])       != 1   { t.Errorf("%s.RFC822[To] have not 1 element", fn) }
	cx++; if len(cv.Digest[0].Action)   == 0   { t.Errorf("%s.Digest.Action is empty", fn) }
	cx++; if cv.Catch                   != nil { t.Errorf("%s.Catch is not nil", fn) }
	cx++; if len(cv.Errors)             != 1   { t.Errorf("%s.Errors have not 1 element", fn) }
	cx++; if cv.Void() == true                 { t.Errorf("%s.Digest returns true", fn) }
	t.Logf("The number of tests = %d", cx)
}

