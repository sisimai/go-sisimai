// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package siba

//  _____         _      __   _ _             ____        __                _____          _   
// |_   _|__  ___| |_   / /__(_) |__   __ _  | __ )  ___ / _| ___  _ __ ___|  ___|_ _  ___| |_ 
//   | |/ _ \/ __| __| / / __| | '_ \ / _` | |  _ \ / _ \ |_ / _ \| '__/ _ \ |_ / _` |/ __| __|
//   | |  __/\__ \ |_ / /\__ \ | |_) | (_| |_| |_) |  __/  _| (_) | | |  __/  _| (_| | (__| |_ 
//   |_|\___||___/\__/_/ |___/_|_.__/ \__,_(_)____/ \___|_|  \___/|_|  \___|_|  \__,_|\___|\__|
import "testing"
import "libsisimai.org/sisimai/v5/eb"

// Sender  string              // Unix FROM line ("From ")
// Headers map[string][]string // Email headers
// Payload string              // Email body
// RFC822  map[string][]string // Email headers of the original message
// Digest  []DeliveryMatter    // Decoded results returned from lhost/*
// Catch   interface{}         // Any data structure returned by the callback function
// Errors  []NotDecoded        // All the errors and warnings
func TestBeforeFact(t *testing.T) {
	fn := "siba.BeforeFact"
	cv := &BeforeFact{
		Sender:  "From <mailer-daemon@example.jp>",
		Headers: map[string][]string{"Subject": []string{"Delivery Failure"}},
		Payload: []byte("Sorry, the email delivery failed"),
		RFC822:  map[string][]string{"To": []string{"<postmaster@example.org>"}},
		Digest:  []DeliveryMatter{DeliveryMatter{Action: eb.AeFAIL}},
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
	cx++; if cv.HasDone() == false { t.Errorf("%s.HasDone() returns false", fn) }
	cx++; if cv.IsEmpty() == true  { t.Errorf("%s.IsEmpty() returns true", fn) }

	cv.Headers = nil
	cx++; if cv.IsEmpty() == false { t.Errorf("%s.IsEmpty() returns false", fn) }
	cx++; if cv.HasDone() == false { t.Errorf("%s.HasDone() returns false", fn) }

	cv.Payload = []byte{}
	cx++; if cv.IsEmpty() == false { t.Errorf("%s.IsEmpty() returns false", fn) }
	cx++; if cv.HasDone() == false { t.Errorf("%s.HasDone() returns false", fn) }

	cv.Digest = nil
	cx++; if cv.IsEmpty() == false { t.Errorf("%s.IsEmpty() returns false", fn) }
	cx++; if cv.HasDone() == true  { t.Errorf("%s.HasDone() returns true", fn) }

	cv.RFC822 = nil
	cx++; if cv.IsEmpty() == false { t.Errorf("%s.IsEmpty() returns false", fn) }
	cx++; if cv.HasDone() == true  { t.Errorf("%s.HasDone() returns false", fn) }

	t.Logf("The number of tests = %d", cx)
}

