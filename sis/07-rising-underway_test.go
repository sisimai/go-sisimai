// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      __   _       ____  _     _             _   _           _                               
// |_   _|__  ___| |_   / /__(_)___  |  _ \(_)___(_)_ __   __ _| | | |_ __   __| | ___ _ ____      ____ _ _   _ 
//   | |/ _ \/ __| __| / / __| / __| | |_) | / __| | '_ \ / _` | | | | '_ \ / _` |/ _ \ '__\ \ /\ / / _` | | | |
//   | |  __/\__ \ |_ / /\__ \ \__ \_|  _ <| \__ \ | | | | (_| | |_| | | | | (_| |  __/ |   \ V  V / (_| | |_| |
//   |_|\___||___/\__/_/ |___/_|___(_)_| \_\_|___/_|_| |_|\__, |\___/|_| |_|\__,_|\___|_|    \_/\_/ \__,_|\__, |
//                                                        |___/                                           |___/ 
import "testing"

// Digest []DeliveryMatter // List of DeliveryMatter structs
// RFC822 string           // The original message
// Errors []NotDecoded     // Errors occurred in sisimai/lhost/*
func TestRisingUnderWay(t *testing.T) {
	fn := "RisingUnderway"
	cv := &RisingUnderway{
		RFC822: "Dummy message",
		Digest: []DeliveryMatter{
			DeliveryMatter{
				Action:    "failed",
				Agent:     "Test",
				Alias:     "neko@example.jp",
				Command:   "RCPT",
				Date:      "Sat, 25 Jan 2025 22:22:22 +0900 (JST)",
				Diagnosis: "User unknown: neko@example.jp",
				FeedbackType: "dummy",
				Lhost:     "mta1.example.jp",
				Reason:    "userunknown",
				Recipient: "neko@example.co.jp",
				ReplyCode: "550",
				Rhost:     "mx34.example.co.jp",
				Spec:      "SMTP",
				Status:    "5.1.1",
			},
		},
		Errors: []NotDecoded{
			NotDecoded{BecauseOf: "Test for " + fn, EmailFile: "25.eml"},
		},
	}
	cx := 0

	cx++; if cv == nil            { t.Fatalf("%s{} = nil", fn) }
	cx++; if cv.RFC822      == "" { t.Errorf("%s.RFC822 is empty", fn) }
	cx++; if len(cv.Digest) == 0  { t.Errorf("%s.RFC822 is empty", fn) }
	cx++; if len(cv.Errors) == 0  { t.Errorf("%s.RFC822 is empty", fn) }
	cx++; if cv.Void()            { t.Errorf("%s.Void() returns true", fn) }

	fn += ".Digest"
	cw := cv.Digest[0]
	cx++; if cw.Action       == "" { t.Errorf("%s.Action is empty", fn) }
	cx++; if cw.Agent        == "" { t.Errorf("%s.Agent is empty", fn) }
	cx++; if cw.Alias        == "" { t.Errorf("%s.Alias is empty", fn) }
	cx++; if cw.Command      == "" { t.Errorf("%s.Command is empty", fn) }
	cx++; if cw.Date         == "" { t.Errorf("%s.Date is empty", fn) }
	cx++; if cw.Diagnosis    == "" { t.Errorf("%s.Diagnosis is empty", fn) }
	cx++; if cw.FeedbackType == "" { t.Errorf("%s.FeedbackType is empty", fn) }
	cx++; if cw.Lhost        == "" { t.Errorf("%s.Lhost is empty", fn) }
	cx++; if cw.Reason       == "" { t.Errorf("%s.Reason is empty", fn) }
	cx++; if cw.Recipient    == "" { t.Errorf("%s.Recipient is empty", fn) }
	cx++; if cw.ReplyCode    == "" { t.Errorf("%s.ReplyCode is empty", fn) }
	cx++; if cw.Rhost        == "" { t.Errorf("%s.Rhost is empty", fn) }
	cx++; if cw.Spec         == "" { t.Errorf("%s.Rhost is empty", fn) }
	cx++; if cw.Status       == "" { t.Errorf("%s.Status is empty", fn) }

	t.Logf("The number of tests = %d", cx)
}

