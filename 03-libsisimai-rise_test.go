// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sisimai

//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
import "testing"
import "strings"
import sisimoji "libsisimai.org/sisimai/string"

func TestRise(t *testing.T) {
	fn := "sisimai.Rise"
	cx := 0

	rootdir := "set-of-emails/"
	samples := []string{"mailbox/mbox-0", "mailbox/mbox-1", "maildir/bsd"}
	normals := []string{"maildir/not"}
	sisiarg := Args(); sisiarg.Delivered = true; sisiarg.Vacation = true
	errorat := []string{"lhost-office365-13.eml"}

	for _, e := range samples {
		ef := "./" + rootdir + e
		cv, ce := Rise(ef, sisiarg)

		cx++; if len(*cv) == 0 { t.Errorf("%s(%s) returns empty", fn, ef) }
		cx++; if len(*ce) != 0 {
			for _, e := range *ce {
				fe := strings.Split(e.EmailFile, "/");
				cx++; if sisimoji.EqualsAny(fe[len(fe) - 1], errorat) { continue }
				cx++; if strings.Contains(e.BecauseOf, "iso-2022-jp") { continue }

				t.Errorf("%s(%s) returns error: %v", fn, ef, ce)
			}
		}

		for j, e := range *cv {
			cx++; if e.Addresser.Void() == true { t.Errorf("[%04d] Addresser.Void() is true", j) }
			cx++; if e.Recipient.Void() == true { t.Errorf("[%04d] Recipient.Void() is true", j) }
			cx++; if e.Catch            != nil  { t.Errorf("[%04d] Catch inlcude data: %v", j, e.Catch) }
			cx++; if e.DecodedBy        == ""   { t.Errorf("[%04d] DecodedBy is empty", j) }
			cx++; if e.Reason           == ""   { t.Errorf("[%04d] Reason is empty", j) }
			cx++; if e.Token            == ""   { t.Errorf("[%04d] Token is empty", j) }

			cx++; if e.SenderDomain != e.Addresser.Host    { t.Errorf("[%04d] Invalid SenderDomain: %s", j, e.SenderDomain) }
			cx++; if e.Destination  != e.Recipient.Host    { t.Errorf("[%04d] Invalid Destination: %s", j, e.Destination) }
			cx++; if e.Alias        == e.Recipient.Address { t.Errorf("[%04d] Invalid Alias: %s", j, e.Alias) }
			cx++; if e.Timestamp.IsZero() == true          { t.Errorf("[%04d] Invalid Timestamp: %v", j, e.Timestamp) }
			cx++; if ! strings.Contains(e.Origin, rootdir) { t.Errorf("[%04d] Invalid Origin: %s ", j, e.Origin) }
		}
	}

	for _, e := range normals {
		ef := "./" + rootdir + e
		cv, ce := Rise(ef, sisiarg)

		cx++; if len(*cv) != 0 { t.Errorf("%s(%s) returns results: %v", fn, ef, *cv) }
		cx++; if len(*ce) != 0 { t.Errorf("%s(%s) returns error: %v", fn, ef, ce) }
	}

	t.Logf("The number of tests = %d", cx)
}

