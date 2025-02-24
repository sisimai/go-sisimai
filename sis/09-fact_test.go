// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      __   _       _____          _   
// |_   _|__  ___| |_   / /__(_)___  |  ___|_ _  ___| |_ 
//   | |/ _ \/ __| __| / / __| / __| | |_ / _` |/ __| __|
//   | |  __/\__ \ |_ / /\__ \ \__ \_|  _| (_| | (__| |_ 
//   |_|\___||___/\__/_/ |___/_|___(_)_|  \__,_|\___|\__|
import "testing"

func TestFact(t *testing.T) {
	cc := "Fact"
	cx := 0
	cv := &Fact{
		Action: "failed",
		Addresser: EmailAddress{Address: "neko@example.jp"},
		Alias: "meumeu@example.org",
		DecodedBy: "Test",
		DeliveryStatus: "5.1.1",
		Destination: "example.org",
		DiagnosticCode: "Test",
		DiagnosticType: "SMTP",
		FeedbackID: "",
		FeedbackType: "",
		HardBounce: true,
		Lhost: "localhost",
		Reason: "userunknown",
		Rhost: "",
		Recipient: EmailAddress{Address: "cat@example.org"},
		ReplyCode: "550",
		Command: "QUIT",
		SenderDomain: "example.jp",
		Token: "",
	}

	cx++; if cv == nil           { t.Fatalf("%s{} = nil", cc) }
	cx++; if cv.Action    == ""  { t.Errorf("%s.Action is empty", cc) }
	cx++; if cv.Alias     == ""  { t.Errorf("%s.Alias is empty", cc) }
	cx++; if cv.DecodedBy == ""  { t.Errorf("%s.DecodedBy is empty", cc) }
	cx++; if cv.Lhost     == ""  { t.Errorf("%s.Lhost is empty", cc) }
	cx++; if cv.Reason    == ""  { t.Errorf("%s.Reason is empty", cc) }
	cx++; if cv.Command   == ""  { t.Errorf("%s.Command is empty", cc) }

	if cv != nil {
		dj, de := cv.MarshalJSON()
		cx++; if string(dj) == ""    { t.Errorf("%s.MarshalJSON() returns empty", cc) }
		cx++; if de         != nil   { t.Errorf("%s.MarshalJSON() returns error: %s", cc, de) }

		dx, de := cv.Dump()
		cx++; if dx == ""  { t.Errorf("%s.Dump() returns empty", cc) }
		cx++; if de != nil { t.Errorf("%s.Dump() returns error: %s", cc, de) }
	}

	cv  = &Fact{}
	if cv != nil {
		dj, de := cv.MarshalJSON()
		cx++; if string(dj) == ""    { t.Errorf("%s.MarshalJSON() returns empty", cc) }
		cx++; if de         != nil   { t.Errorf("%s.MarshalJSON() returns error: %s", cc, de) }

		dx, de := cv.Dump()
		cx++; if dx == ""  { t.Errorf("%s.Dump() returns empty", cc) }
		cx++; if de != nil { t.Errorf("%s.Dump() returns error: %s", cc, de) }

	}
	t.Logf("The number of tests = %d", cx)
}

