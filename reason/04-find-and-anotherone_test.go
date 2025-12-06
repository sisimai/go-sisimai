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

		// Delivered
		siba.Fact{DeliveryStatus: "2.2.2", ReplyCode: "250"},

		// Blocked
		siba.Fact{DiagnosticCode: "", Action: eb.CeEHLO},
		siba.Fact{DiagnosticCode: "", Action: eb.CeHELO},

		// MailboxFull
		siba.Fact{DiagnosticType: "SMTP", Reason: eb.ReFULL},
		siba.Fact{DiagnosticType: "SMTP", DeliveryStatus: "4.2.2"},
		siba.Fact{DiagnosticType: "SMTP", DiagnosticCode: "Full Mailbox"},

		// MesgTooBig
		siba.Fact{DiagnosticType: "SMTP", DeliveryStatus: "5.2.3"},
		siba.Fact{DiagnosticType: "SMTP", DiagnosticCode: "Message too big"},

		// ExceedLimit
		siba.Fact{DiagnosticType: "SMTP", DeliveryStatus: "5.3.4"},
		siba.Fact{DiagnosticType: "SMTP", DiagnosticCode: "Message too large"},

		// MailerError
		siba.Fact{DiagnosticType: "X-UNIX", DeliveryStatus: "5.0.0"},
		siba.Fact{DiagnosticType: "X-UNIX", DiagnosticCode: "X-Unix: 127;"},

		// SystemError
		siba.Fact{DiagnosticType: "X-UNIX", Reason: eb.Re___1},
		siba.Fact{DiagnosticType: "", DiagnosticCode: "local configuration error"},

		// Expired
		siba.Fact{Reason: eb.ReEXPR},
		siba.Fact{DiagnosticCode: "Message timed out"},

		// NetworkError
		siba.Fact{Reason: eb.ReNETW},
		siba.Fact{DiagnosticCode: "No route to host"},

		// UserUnknown, Filtered
		siba.Fact{Reason: eb.ReUSER, DiagnosticCode: "User unknown", DeliveryStatus: "5.1.1"},
		siba.Fact{Reason: eb.ReUSER, DiagnosticCode: "User unknown", DeliveryStatus: "5.1.2"},

		// HostUnknown
		siba.Fact{Reason: eb.ReHOST, DiagnosticCode: "Unknown Host", DeliveryStatus: "5.1.0"},

		// FailedSTARTTLS
		siba.Fact{Reason: eb.Re___1, Command: eb.CeTTLS},
		siba.Fact{ReplyCode: "523"},
		siba.Fact{DiagnosticCode: "STARTTLS is required to send mail"},

		// ContentError
		siba.Fact{DiagnosticType: "NEKO", Reason: eb.Re___0},
		siba.Fact{DiagnosticType: "NEKO", Reason: eb.Re___0, DeliveryStatus: "4.6.0"},
		siba.Fact{DiagnosticType: "NEKO", Reason: eb.Re___0, DeliveryStatus: "5.6.0"},

		// SecurityError
		siba.Fact{Reason: eb.ReSECU, DeliveryStatus: "4.7.0"},
		siba.Fact{Reason: eb.ReSECU, DeliveryStatus: "5.7.0"},
		siba.Fact{DiagnosticCode: "Verification failure"},

		// SyntaxError
		siba.Fact{DiagnosticType: "NEKO", Reason: eb.Re___0, ReplyCode: "503"},

		// Undefined
		siba.Fact{DiagnosticType: "NEKO", Reason: eb.Re___0},
		siba.Fact{DiagnosticType: "NEKO", DeliveryStatus: "5.0.900"},
		siba.Fact{DiagnosticType: "NEKO", DiagnosticCode: ""},
		siba.Fact{DiagnosticType: "NEKO", DiagnosticCode: "", Action: eb.AeSTAY},
		siba.Fact{DiagnosticType: "NEKO", DiagnosticCode: "", Action: eb.AeFAIL},

		// OnHold
		siba.Fact{DiagnosticType: "NEKO", Reason: eb.Re___1},
		siba.Fact{DiagnosticType: "NEKO", DeliveryStatus: "5.0.901"},
		siba.Fact{DiagnosticType: "NEKO", Action: eb.AeSTAY},
		siba.Fact{DiagnosticCode: "Nyaaaaaan?"},

		// Vacation
		siba.Fact{DiagnosticType: "SMTP", Reason: eb.ReAWAY},
		siba.Fact{DiagnosticType: "SMTP", DiagnosticCode: "I am out of the office today"},
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

