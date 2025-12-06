// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____         _      __                              
// |_   _|__  ___| |_   / / __ ___  __ _ ___  ___  _ __  
//   | |/ _ \/ __| __| / / '__/ _ \/ _` / __|/ _ \| '_ \ 
//   | |  __/\__ \ |_ / /| | |  __/ (_| \__ \ (_) | | | |
//   |_|\___||___/\__/_/ |_|  \___|\__,_|___/\___/|_| |_|
import "testing"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"

func TestFind(t *testing.T) {
	fn := "reason.Find"
	cx := 0
	cw := []siba.Fact{
		siba.Fact{},
		siba.Fact{Reason: eb.Re___0},
		siba.Fact{Reason: eb.ReUSER},
		siba.Fact{DeliveryStatus: "2.2.2", ReplyCode: "250"},

		// MailboxFull
		siba.Fact{DiagnosticType: "SMTP", Reason: eb.ReFULL},
		siba.Fact{DiagnosticType: "SMTP", DeliveryStatus: "4.2.2"},
		siba.Fact{DiagnosticType: "SMTP", DiagnosticCode: "Full Mailbox"},

		// MesgTooBig
		siba.Fact{DiagnosticType: "SMTP", Reason: eb.ReSIZE},
		siba.Fact{DiagnosticType: "SMTP", DeliveryStatus: "5.2.3"},
		siba.Fact{DiagnosticType: "SMTP", DiagnosticCode: "Message too big"},

		// ExceedLimit
		siba.Fact{DiagnosticType: "SMTP", Reason: eb.ReXLIM},
		siba.Fact{DiagnosticType: "SMTP", DeliveryStatus: "5.3.4"},
		siba.Fact{DiagnosticType: "SMTP", DiagnosticCode: "Message too large"},

	}

	for j, e := range cw {
		cx++; if cr := Find(&e); cr == "" { t.Errorf("%s(%d) returns %s", fn, j, cr) }
	}
	cx++; if cr := Find(nil); cr != "" { t.Errorf("%s(nil) returns %s", fn, cr) }

	t.Logf("The number of tests = %d", cx)
}

func TestAnotherOne(t *testing.T) {
	fn := "reason.anotherone"
	cx := 0
	cv := anotherone(nil)

	cx++; if cv != "" { t.Errorf("%s(nil) returns %s", fn, cv) }
	t.Logf("The number of tests = %d", cx)
}

