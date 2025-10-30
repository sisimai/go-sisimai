// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _____         _      ___ _               _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|
import "testing"
import "libsisimai.org/sisimai/v5/siba"

func TestInquireFor(t *testing.T) {
	fn := "lhost.InquireFor"
	cx := 0
	ae := []string{
		"Activehunter", "AmazonSES", "ApacheJames", "Biglobe", "Courier", "Domino", "DragonFly", "EZweb",
		"EinsUndEins", "Exchange2003", "Exchange2007", "Exim", "FML", "GMX", "GoogleGroups", "Gmail",
		"GoogleWorkspace", "IMailServer", "KDDI", "MailFoundry", "MailMarshalSMTP", "MessagingServer",
		"Notes", "OpenSMTPD", "Postfix", "Sendmail", "TrendMicro", "V5sendmail", "Verizon",
		"X1", "X2", "X3", "X6", "Zoho", "mFILTER", "qmail",
	}

	for _, e := range ae {
		bf := &siba.BeforeFact{
			Sender:  "MAILER-DAEMON",
			Headers: map[string][]string{
				"from": []string{"<postmaster@example.jp>"},
				"received": []string{"via localhost"},
				"message-id": []string{"<22.02@example.jp>"},
				"content-type": []string{"text/plain"},
				"subject": []string{"Delivery failure"},
			},
			Payload: "Nekochan",
			RFC822:  map[string][]string{},
			Digest:  []siba.DeliveryMatter{},
		}
		cx++; if cv := InquireFor[e](nil); cv != nil { t.Errorf("%s[%s]() did not return nil", fn, e) }
		cx++; if cv := InquireFor[e](bf);  cv != nil { t.Errorf("%s[%s]() did not return nil", fn, e) }

		bf.Payload = ""
		cx++; if cv := InquireFor[e](bf);  cv != nil { t.Errorf("%s[%s]() did not return nil", fn, e) }

		bf.Payload = "nekochan"
		bf.Headers = map[string][]string{}
		cx++; if cv := InquireFor[e](bf);  cv != nil { t.Errorf("%s[%s]() did not return nil", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

