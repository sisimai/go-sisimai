// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sisimai

//  _____         _      ___ _ _         _     _                 _ 
// |_   _|__  ___| |_   / / (_) |__  ___(_)___(_)_ __ ___   __ _(_)
//   | |/ _ \/ __| __| / /| | | '_ \/ __| / __| | '_ ` _ \ / _` | |
//   | |  __/\__ \ |_ / / | | | |_) \__ \ \__ \ | | | | | | (_| | |
//   |_|\___||___/\__/_/  |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_|
import "testing"

func TestFactor(t *testing.T) {
	fn := "sisimai.Factor"
	cx := 0
	cv := Factor()

	cx++; if cv                == nil  { t.Fatalf("%s() returned nil", fn) }
	cx++; if cv.Action         != ""   { t.Errorf("%s().Action is not empty: %s", fn, cv.Action) }
	cx++; if cv.Addresser.User != ""   { t.Errorf("%s().Addresser.User is not empty: %s", fn, cv.Addresser.User) }
	cx++; if cv.Alias          != ""   { t.Errorf("%s().Alias is not empty: %s", fn, cv.Alias) }
	cx++; if cv.Catch          != nil  { t.Errorf("%s().Catch is not nil: %v", fn, cv.Catch) }
	cx++; if cv.Command        != ""   { t.Errorf("%s().Command is not empty: %s", fn, cv.Command) }
	cx++; if cv.DecodedBy      != ""   { t.Errorf("%s().DecodedBy is not empty: %s", fn, cv.DecodedBy) }
	cx++; if cv.DeliveryStatus != ""   { t.Errorf("%s().DeliveryStatus is not empty: %s", fn, cv.DeliveryStatus) }
	cx++; if cv.Destination    != ""   { t.Errorf("%s().Destination is not empty: %s", fn, cv.Destination) }
	cx++; if cv.DiagnosticCode != ""   { t.Errorf("%s().DiagnosticCode is not empty: %s", fn, cv.DiagnosticCode) }
	cx++; if cv.DiagnosticType != ""   { t.Errorf("%s().DiagnosticType is not empty: %s", fn, cv.DiagnosticType) }
	cx++; if cv.FeedbackID     != ""   { t.Errorf("%s().FeedbackID is not empty: %s", fn, cv.FeedbackID) }
	cx++; if cv.FeedbackType   != ""   { t.Errorf("%s().FeedbackType is not empty: %s", fn, cv.FeedbackType) }
	cx++; if cv.HardBounce     == true { t.Errorf("%s().HardBounce is not false: %t", fn, cv.HardBounce) }
	cx++; if cv.Lhost          != ""   { t.Errorf("%s().Lhost is not empty: %s", fn, cv.Lhost) }
	cx++; if cv.ListID         != ""   { t.Errorf("%s().ListID is not empty: %s", fn, cv.ListID) }
	cx++; if cv.MessageID      != ""   { t.Errorf("%s().MessageID is not empty: %s", fn, cv.MessageID) }
	cx++; if cv.Origin         != ""   { t.Errorf("%s().Origin is not empty: %s", fn, cv.Origin) }
	cx++; if cv.Reason         != ""   { t.Errorf("%s().Reason is not empty: %s", fn, cv.Reason) }
	cx++; if cv.Rhost          != ""   { t.Errorf("%s().Rhost is not empty: %s", fn, cv.Rhost) }
	cx++; if cv.Recipient.User != ""   { t.Errorf("%s().Recipient.User is not empty: %s", fn, cv.Recipient.User) }
	cx++; if cv.ReplyCode      != ""   { t.Errorf("%s().ReplyCode is not empty: %s", fn, cv.ReplyCode) }
	cx++; if cv.SenderDomain   != ""   { t.Errorf("%s().SenderDomain is not empty: %s", fn, cv.SenderDomain) }
	cx++; if cv.Subject        != ""   { t.Errorf("%s().Subject is not empty: %s", fn, cv.Subject) }
	cx++; if cv.TimezoneOffset != ""   { t.Errorf("%s().TimezoneOffset is not empty: %s", fn, cv.TimezoneOffset) }
	cx++; if cv.Token          != ""   { t.Errorf("%s().Token is not empty: %s", fn, cv.Token) }

	t.Logf("The number of tests = %d", cx)
}

