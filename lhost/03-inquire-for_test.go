// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _____         _      ___ _               _   
// |_   _|__  ___| |_   / / | |__   ___  ___| |_ 
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ |_ 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|
import "testing"
import "libsisimai.org/sisimai/sis"

func TestInquireFor(t *testing.T) {
	fn := "sisimai/lhost.InquireFor"
	cx := 0

	for _, e := range INDEX() {
		bf := &sis.BeforeFact{
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
			Digest:  []sis.DeliveryMatter{},
		}
		cx++; if cv := InquireFor[e](nil); cv.Void() == false { t.Errorf("%s(%s).Void() return false", fn, e) }
		cx++; if cv := InquireFor[e](bf);  cv.Void() == false { t.Errorf("%s(%s).Void() return false", fn, e) }

		bf.Payload = ""
		cx++; if cv := InquireFor[e](bf);  cv.Void() == false { t.Errorf("%s(%s).Void() return false", fn, e) }

		bf.Headers = map[string][]string{}
		cx++; if cv := InquireFor[e](bf);  cv.Void() == false { t.Errorf("%s(%s).Void() return false", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

