// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//  _____         _      __    _               _   
// |_   _|__  ___| |_   / / __| |__   ___  ___| |_ 
//   | |/ _ \/ __| __| / / '__| '_ \ / _ \/ __| __|
//   | |  __/\__ \ |_ / /| |  | | | | (_) \__ \ |_ 
//   |_|\___||___/\__/_/ |_|  |_| |_|\___/|___/\__|
import "testing"
import "libsisimai.org/sisimai/sis"

var TestRhosts = []struct {lhost string; rhost string; destination string; expected string}{
	{"mx.example.com", "mx2.mail.aol.com", "aol.com", "Aol"},
	{"mx2.mail.aol.com", "mx0.example.com", "aol.com", "Aol"},
	{"mx.example.com", "mx.example.org", "aol.com", "Aol"},

	{"mx.example.com", "mx2.mail.icloud.com", "apple.com", "Apple"},
	{"mx2.mail.icloud.com", "mx0.example.com", "me.com", "Apple"},
	{"mx.example.com", "mx.example.org", "icloud.com", "Apple"},

	{"mx.example.com", "relay1.mx.cloudflare.net", "example.com", "Cloudflare"},
	{"relay2.mx.cloudflare.net", "mx0.example.com", "example.org", "Cloudflare"},
	{"mx.example.com", "mx.example.org", "cloudflare.net", "Cloudflare"},

	{"mx.example.com", "mx.cox.net", "example.com", "Cox"},
	{"mx.cox.net", "mx0.example.com", "example.org", "Cox"},
	{"mx.example.com", "mx.example.org", "cox.net", "Cox"},

	{"mx.example.com", "mx.facebook.com", "example.com", "Facebook"},
	{"mx.facebook.com", "mx0.example.com", "example.org", "Facebook"},
	{"mx.example.com", "mx.example.org", "facebook.com", "Facebook"},

	{"mx.example.com", "mx.laposte.net", "example.com", "FrancePTT"},
	{"mx.orange.fr", "mx0.example.com", "example.org", "FrancePTT"},
	{"mx.example.com", "mx.example.org", "wanadoo.fr", "FrancePTT"},

	{"mx.example.com", "smtp.secureserver.net", "example.com", "GoDaddy"},
	{"mailstore1.secureserver.net", "mx0.example.com", "example.org", "GoDaddy"},
	{"mx.example.com", "mx.example.org", "secureserver.net", "GoDaddy"},

	{"mx.example.com", "aspmx.l.google.com", "example.com", "Google"},
	{"gmail-smtp-in.l.google.com", "mx0.example.com", "example.org", "Google"},
	{"mx.example.com", "mx.example.org", "google.com", "Google"},

	{"mx.example.com", "mx.l.googlemail.com", "example.com", "GSuite"},
	{"gmail-smtp-in.l.googlemail.com", "mx0.example.com", "example.org", "GSuite"},
	{"mx.example.com", "mx.example.org", "googlemail.com", "GSuite"},

	{"mx.example.com", "mx.email.ua", "example.com", "IUA"},
	{"smtp.email.ua", "mx0.example.com", "example.org", "IUA"},
	{"mx.example.com", "mx.example.org", "email.ua", "IUA"},

	{"mx.example.com", "lsean.ezweb.ne.jp", "example.com", "KDDI"},
	{"msmx.au.com", "mx0.example.com", "example.org", "KDDI"},
	{"mx.example.com", "mx.example.org", "au.com", "KDDI"},

	{"mx.example.com", "mx.messagelabs.com", "example.com", "MessageLabs"},
	{"mx1.messagelabs.com", "mx0.example.com", "example.org", "MessageLabs"},
	{"mx.example.com", "mx.example.org", "messagelabs.com", "MessageLabs"},

	{"mx.example.com", "mx.prod.outlook.com", "example.com", "Microsoft"},
	{"mx1.protection.outlook.com", "mx0.example.com", "example.org", "Microsoft"},
	{"mx.example.com", "mx.example.org", "onmicrosoft.com", "Microsoft"},
	{"mx.example.com", "mx.example.org", "exchangelabs.com", "Microsoft"},

	{"mx.example.com", "mx.mimecast.com", "example.com", "Mimecast"},
	{"mx.mimecast.com", "mx0.example.com", "example.org", "Mimecast"},
	{"mx.example.com", "mx.example.org", "mimecast.com", "Mimecast"},

	{"mx.example.com", "mfsmax.docomo.ne.jp", "example.com", "NTTDOCOMO"},
	{"mfsmax.docomo.ne.jp", "mx0.example.com", "example.org", "NTTDOCOMO"},
	{"mx.example.com", "mx.example.org", "docomo.ne.jp", "NTTDOCOMO"},

	{"mx.example.com", "mx.hotmail.com", "example.com", "Outlook"},
	{"mx.hotmail.com", "mx0.example.com", "example.org", "Outlook"},
	{"mx.example.com", "mx.example.org", "hotmail.com", "Outlook"},

	{"mx.example.com", "mx.charter.net", "example.com", "Spectrum"},
	{"mx.charter.net", "mx0.example.com", "example.org", "Spectrum"},
	{"mx.example.com", "mx.example.org", "charter.net", "Spectrum"},

	{"mx.example.com", "mx.qq.com", "example.com", "Tencent"},
	{"mx.qq.com", "mx0.example.com", "example.org", "Tencent"},
	{"mx.example.com", "mx.example.org", "qq.com", "Tencent"},

	{"mx.example.com", "mx.yahoodns.net", "example.com", "YahooInc"},
	{"mx.yahoodns.net", "mx0.example.com", "example.org", "YahooInc"},
	{"mx.example.com", "mx.example.org", "yahoodns.net", "YahooInc"},
}

func TestName(t *testing.T) {
	fn := "sisimai/rhost.Name"
	cv := Name(nil)
	cx := 0

	cx++; if cv != "" { t.Errorf("%s(nil) returns %s", fn, cv) }
	for _, e := range TestRhosts {
		ae := &sis.Fact{
			Lhost: e.lhost,
			Rhost: e.rhost,
			Destination: e.destination,
		}
		cx++; if cv = Name(ae); cv == ""         { t.Errorf("%s(%s) returns an empty string", fn, e.destination) }
		cx++; if cv = Name(ae); cv != e.expected { t.Errorf("%s(%s) returns %s", fn, e.destination, e.expected)  }
	}
	t.Logf("The number of tests = %d", cx)
}

func TestFind(t *testing.T) {
	fn := "sisimai/rhost.Find"
	cv := Name(nil)
	cx := 0

	cx++; if cv != "" { t.Errorf("%s(nil) returns %s", fn, cv) }
	for _, e := range TestRhosts {
		ae := &sis.Fact{
			Lhost: e.lhost,
			Rhost: e.rhost,
			Destination: e.destination,
			DiagnosticCode: "nekochan-nyaan",
			DeliveryStatus: "5.0.0",
			ReplyCode: "550",
			Reason: "",
		}
		cx++; if cv = Find(ae); cv != "" { t.Errorf("%s(%s) returns %s", fn, e.destination, e.expected)  }

		ae.DiagnosticCode = ""
		cx++; if cv = Find(ae); cv != "" { t.Errorf("%s(%s) returns %s", fn, e.destination, e.expected)  }
	}
	t.Logf("The number of tests = %d", cx)
}

