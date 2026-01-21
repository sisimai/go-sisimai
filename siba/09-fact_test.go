// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package siba

//  _____         _      __   _ _             _____          _   
// |_   _|__  ___| |_   / /__(_) |__   __ _  |  ___|_ _  ___| |_ 
//   | |/ _ \/ __| __| / / __| | '_ \ / _` | | |_ / _` |/ __| __|
//   | |  __/\__ \ |_ / /\__ \ | |_) | (_| |_|  _| (_| | (__| |_ 
//   |_|\___||___/\__/_/ |___/_|_.__/ \__,_(_)_|  \__,_|\___|\__|
import "testing"
import "libsisimai.org/sisimai/v5/eb"

func TestFact(t *testing.T) {
	cc := "Fact"
	cx := 0
	cv := &Fact{
		Action: eb.AeFAIL,
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
		Reason: eb.ReUSER,
		Rhost: "",
		Recipient: EmailAddress{Address: "cat@example.org"},
		ReplyCode: "550",
		Command: eb.CeQUIT,
		SenderDomain: "example.jp",
		Token: "",
		Toxic: true,
	}

	cx++; if cv == nil             { t.Fatalf("%s{} = nil", cc) }
	cx++; if cv.Action    == ""    { t.Errorf("%s.Action is empty", cc) }
	cx++; if cv.Alias     == ""    { t.Errorf("%s.Alias is empty", cc) }
	cx++; if cv.DecodedBy == ""    { t.Errorf("%s.DecodedBy is empty", cc) }
	cx++; if cv.Lhost     == ""    { t.Errorf("%s.Lhost is empty", cc) }
	cx++; if cv.Reason    == ""    { t.Errorf("%s.Reason is empty", cc) }
	cx++; if cv.Command   == ""    { t.Errorf("%s.Command is empty", cc) }
	cx++; if cv.Toxic     == false { t.Errorf("%s.Toxic is false", cc) }
	cx++; if cv.IsToxic() == false { t.Errorf("%s.IsToxic() returns false", cc) }

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
		cx++; if cv.IsToxic() == true { t.Errorf("%s.IsToxic() returns true", cc) }

		dj, de := cv.MarshalJSON()
		cx++; if string(dj) == ""    { t.Errorf("%s.MarshalJSON() returns empty", cc) }
		cx++; if de         != nil   { t.Errorf("%s.MarshalJSON() returns error: %s", cc, de) }

		dx, de := cv.Dump()
		cx++; if dx == ""  { t.Errorf("%s.Dump() returns empty", cc) }
		cx++; if de != nil { t.Errorf("%s.Dump() returns error: %s", cc, de) }
	}
	t.Logf("The number of tests = %d", cx)
}

func TestIsToxic(t *testing.T) {
	fn := "IsToxic"
	cx := 0
	cw := []Fact{
		Fact{},
		Fact{DeliveryStatus: "5.0.0", ReplyCode: "550", Reason: eb.Re___0, Command: eb.CeCONN},
		Fact{DeliveryStatus: "4.0.0", ReplyCode: "421", Reason: eb.Re___1, Command: eb.CeCONN},
		Fact{DeliveryStatus: "4.2.2", ReplyCode: "450", Reason: eb.ReFULL, Command: eb.CeRCPT},
		Fact{DeliveryStatus: "5.9.999", ReplyCode: "",  Reason: eb.RePASS, Command: eb.CeMAIL},
		Fact{DeliveryStatus: "", ReplyCode: "", Reason: eb.ReFEED, Command: "", FeedbackType: "auth-failure"},
	}
	cv := []Fact{
		Fact{DeliveryStatus: "5.1.0", ReplyCode: "550", Reason: eb.ReHOST, Command: eb.CeCONN},
		Fact{DeliveryStatus: "5.1.1", ReplyCode: "550", Reason: eb.ReUSER, Command: eb.CeRCPT},
		Fact{DeliveryStatus: "5.1.6", ReplyCode: "556", Reason: eb.ReMOVE, Command: eb.CeRCPT},
		Fact{DeliveryStatus: "5.0.1", ReplyCode: "500", Reason: eb.Re00MX, Command: eb.CeCONN},
		Fact{DeliveryStatus: "5.7.0", ReplyCode: "550", Reason: eb.ReQUIT, Command: eb.CeDATA},
		Fact{DeliveryStatus: "5.7.1", ReplyCode: "550", Reason: eb.ReSTOP, Command: eb.CeCONN},
		Fact{DeliveryStatus: "5.1.2", ReplyCode: "501", Reason: eb.ReFILT, Command: eb.CeRCPT},
		Fact{DeliveryStatus: "5.2.2", ReplyCode: "552", Reason: eb.ReFULL, Command: eb.CeRCPT},
		Fact{DeliveryStatus: "5.7.3", ReplyCode: "550", Reason: eb.RePASS, Command: eb.CeRCPT},
		Fact{DeliveryStatus: "5.7.4", ReplyCode: "",    Reason: eb.RePASS, Command: eb.CeMAIL},
		Fact{DeliveryStatus: "", ReplyCode: "", Reason: eb.ReFEED, Command: "", FeedbackType: "abuse"},
	}

	for _, e := range cw {
		cx++; if e.IsToxic() == true  { t.Errorf("%s(%s) returns true",  fn, e.DeliveryStatus) }
	}
	for _, e := range cv {
		cx++; if e.IsToxic() == false { t.Errorf("%s(%s) returns false", fn, e.DeliveryStatus) }
	}
	t.Logf("The number of tests = %d", cx)
}




