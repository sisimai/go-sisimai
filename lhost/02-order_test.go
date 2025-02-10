// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _____         _      __            _           
// |_   _|__  ___| |_   / /__  _ __ __| | ___ _ __ 
//   | |/ _ \/ __| __| / / _ \| '__/ _` |/ _ \ '__|
//   | |  __/\__ \ |_ / / (_) | | | (_| |  __/ |   
//   |_|\___||___/\__/_/ \___/|_|  \__,_|\___|_|   
import "testing"

func TestOrderBySubject(t *testing.T) {
	fn := "sisimai/lhost.OrderBySubject"
	cx := 0
	ae := []string{
		"Undeliverable Mail",
		"complaint about message from 192.0.2.222",
		"AWS Notification Message",
		"Abuse Report",
		"Auto reply: Nyaan",
		"Automatic reply: Nyaan",
		"DELIVERY FAILURE: ",
		"DELIVERY FAILURE: User Kijitoranyan (kijitora@example.jp) not listed in",
		"Delivery Notification: Delivery has been delayed",
		"Delivery Notification: Delivery has failed",
		"Delivery Status",
		"Delivery Status Notification (Delay)",
		"Delivery Status Notification (Failure)",
		"Delivery Status Notification (Mail Delivery Delayed)",
		"Delivery failure",
		"Delivery status notification",
		"Delivery status notification: delayed",
		"Delivery status notification: error",
		"Delivery status notification: failed",
		"Delivery status notification: warning",
		"Email Feedback Report for IP 192.0.2.25",
		"FAILURE NOTICE : Nyaan",
		"Failure Notice",
		"Mail Delivery Failure",
		"Mail Delivery Status Notification (Delay)",
		"Mail Delivery Status Report",
		"Mail System Error - Returned Mail",
		"Mail could not be delivered",
		"Mail delivery failed",
		"Mail delivery failed: returning message to sender",
		"Mail failure - malformed recipient address",
		"Message delivery has failed",
		"NOTICE: mail delivery status.",
		"Non recapitabile: Neko Nyaan",
		"Postmaster notify: see transcript for details",
		"Returned Mail: User unknown",
		"Returned mail: Cannot send message for 5 days",
		"Returned mail: Deferred: Connection timed out during user open with example.org",
		"Returned mail: Deferred: Host Name Lookup Failure",
		"Returned mail: Host unknown",
		"Returned mail: Nyaan",
		"Returned mail: Requested action not taken: mailbox name not allowed",
		"Returned mail: Service unavailable",
		"Returned mail: User unknown",
		"Returned mail: see transcript for details",
		"There was an error sending your mail",
		"Undeliverable Mail",
		"Undeliverable Mail ",
		`Undeliverable Mail: "Nyaan"`,
		"Undeliverable mail, MTA-BLOCKED",
		"Undeliverable Message",
		"Undeliverable message",
		"Undeliverable: Kijitora Cat",
		"Undeliverable: Nyaaaaan",
		"Undeliverable: Nyaan",
		"Undeliverable: Nyaan, Neko Nyaan",
		"Undeliverable: email bounce",
		"Undeliverable: kijitora@nyaan.example.net",
		"Undelivered Mail Returned to Sender",
		"Warning: could not send message for past 4 hours",
		"Warning: message neko222-nyaaan-22 delayed 24 hours",
		"[dmarc-ietf] DMARC test message",
		"failed delivery",
		"failure delivery",
		"failure notice",
	}

	cv := OrderBySubject("")
	cx++; if len(cv) != 0 { t.Errorf("%s() is not empty: %s", fn, cv) }
	for _, e := range ae {
		cv = OrderBySubject(e)
		cx++; if len(cv) == 0 { t.Errorf("%s(%s) is empty", fn, e) }
	}
	t.Logf("The number of tests = %d", cx)
}

func TestAnotherOrder(t *testing.T) {
	fn := "sisimai/lhost.OrderBySubject"
	cx := 0
	cv := AnotherOrder()

	cx++; if len(cv) == 0 { t.Errorf("%s() is empty", fn) }
	t.Logf("The number of tests = %d", cx)
}


