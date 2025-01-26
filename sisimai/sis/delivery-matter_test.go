// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      __   _       ____       _ _                      __  __       _   _            
// |_   _|__  ___| |_   / /__(_)___  |  _ \  ___| (_)_   _____ _ __ _   _|  \/  | __ _| |_| |_ ___ _ __ 
//   | |/ _ \/ __| __| / / __| / __| | | | |/ _ \ | \ \ / / _ \ '__| | | | |\/| |/ _` | __| __/ _ \ '__|
//   | |  __/\__ \ |_ / /\__ \ \__ \_| |_| |  __/ | |\ V /  __/ |  | |_| | |  | | (_| | |_| ||  __/ |   
//   |_|\___||___/\__/_/ |___/_|___(_)____/ \___|_|_| \_/ \___|_|   \__, |_|  |_|\__,_|\__|\__\___|_|   
//                                                                  |___/                               
import "testing"
import "strings"

// Action       string     // The value of Action header
// Agent        string     // MTA name
// Alias        string     // The value of alias entry(RHS)
// Command      string     // SMTP command in the message body
// Date         string     // The value of Last-Attempt-Date header
// Diagnosis    string     // The value of Diagnostic-Code header
// FeedbackType string     // Feedback type
// Lhost        string     // The value of Received-From-MTA header
// Reason       string     // Temporary reason of bounce
// Recipient    string     // The value of Final-Recipient header
// ReplyCode    string     // SMTP Reply Code
// Rhost        string     // The value of Remote-MTA header
// Spec         string     // Protocl specification
// Status       string     // The value of Status header
func TestDeliveryMatter(t *testing.T) {
	fn := "DeliveryMatter"
	cv := &DeliveryMatter{
		Action:    "failed",
		Agent:     "Test",
		Alias:     "neko@example.jp",
		Command:   "RCPT",
		Date:      "Sat, 25 Jan 2025 22:22:22 +0900 (JST)",
		Diagnosis: "User unknown: neko@example.jp",
		FeedbackType: "dummy",
		Lhost:     "mta1.example.jp",
		Reason:    "userunknown",
		Recipient: "neko@example.co.jp",
		ReplyCode: "550",
		Rhost:     "mx34.example.co.jp",
		Spec:      "SMTP",
		Status:    "5.1.1",
	}
	cx := 0

	cx++; if cv == nil             { t.Fatalf("%s{} = nil", fn) }
	cx++; if cv.Action       == "" { t.Errorf("%s.Action is empty", fn) }
	cx++; if cv.Agent        == "" { t.Errorf("%s.Agent is empty", fn) }
	cx++; if cv.Alias        == "" { t.Errorf("%s.Alias is empty", fn) }
	cx++; if cv.Command      == "" { t.Errorf("%s.Command is empty", fn) }
	cx++; if cv.Date         == "" { t.Errorf("%s.Date is empty", fn) }
	cx++; if cv.Diagnosis    == "" { t.Errorf("%s.Diagnosis is empty", fn) }
	cx++; if cv.FeedbackType == "" { t.Errorf("%s.FeedbackType is empty", fn) }
	cx++; if cv.Lhost        == "" { t.Errorf("%s.Lhost is empty", fn) }
	cx++; if cv.Reason       == "" { t.Errorf("%s.Reason is empty", fn) }
	cx++; if cv.Recipient    == "" { t.Errorf("%s.Recipient is empty", fn) }
	cx++; if cv.ReplyCode    == "" { t.Errorf("%s.ReplyCode is empty", fn) }
	cx++; if cv.Rhost        == "" { t.Errorf("%s.Rhost is empty", fn) }
	cx++; if cv.Spec         == "" { t.Errorf("%s.Rhost is empty", fn) }
	cx++; if cv.Status       == "" { t.Errorf("%s.Status is empty", fn) }

	// Select()
	fn  = "Select"
	ae := []string{
		"action", "agent", "alias", "command", "date", "diagnosis", "feedbacktype", "lhost",
		"reason", "recipient", "replycode", "rhost", "spec", "status",
	}
	cw := cv.Select("neko")
	cx++; if cw != "" { t.Errorf("%s(neko) is not empty: %s", fn, cw) }

	for _, e := range ae {
		cw = cv.Select(e)
		cx++; if cw == "" { t.Errorf("%s(%s) is empty", fn, e) }
	}

	// Update()
	fn  = "Update"
	ct := cv.Update("", ""); cx++;       if ct == true  { t.Errorf("%s('', '') returns true", fn) }
	ct  = cv.Update("neko", "2"); cx++;  if ct == true  { t.Errorf("%s('neko', '22') returns true", fn) }
	for _, e := range ae {
		// Empty value in the 2nd argument is not allowed
		ct = cv.Update(e, "")
		cx++; if ct == true { t.Errorf("%s(%s, '') returns true", fn, e) }
	}
	for _, e := range ae {
		// The same value in the 2nd argument is not updated
		ct = cv.Update(e, cv.Select(e))
		cx++; if ct == true { t.Errorf("%s(%s, %s) returns true", fn, e, cv.Select(e)) }
	}

	ct  = cv.Update("action", "delayed")
	cx++; if ct == false            { t.Errorf("%s(action, delayed) returns false", fn) }
	cx++; if cv.Action != "delayed" { t.Errorf("%s(action, delayed) did not updated: %s", fn, cv.Action) }

	ct  = cv.Update("agent", "OpenSMTPD")
	cx++; if ct == false             { t.Errorf("%s(agent, OpenSMTPD) returns false", fn) }
	cx++; if cv.Agent != "OpenSMTPD" { t.Errorf("%s(agent, OpenSMTPD) did not updated: %s", fn, cv.Agent) }

	ct  = cv.Update("alias", "neko@example.us")
	cx++; if ct == false                   { t.Errorf("%s(alias, neko@example.us) returns false", fn) }
	cx++; if cv.Alias != "neko@example.us" { t.Errorf("%s(alias, neko@example.us) did not updated: %s", fn, cv.Alias) }

	ct  = cv.Update("command", "STARTTLS")
	cx++; if ct == false                   { t.Errorf("%s(command, STARTTLS) returns false", fn) }
	cx++; if cv.Command != "STARTTLS"      { t.Errorf("%s(command, STARTTLS) did not updated: %s", fn, cv.Command) }

	ct  = cv.Update("date", "Sun, 26 Jan 2025 16:31:50 +0900 (JST)")
	cx++; if ct == false                       { t.Errorf("%s(date, Sun, 26) returns false", fn) }
	cx++; if strings.HasPrefix(cv.Date, "Sat") { t.Errorf("%s(date, Sun, 26) did not updated: %s", fn, cv.Date) }

	ct  = cv.Update("diagnosis", "Mailbox full")
	cx++; if ct == false                    { t.Errorf("%s(diagnosis, Mailbox full) returns false", fn) }
	cx++; if cv.Diagnosis != "Mailbox full" { t.Errorf("%s(diagnosis, Mailbox full) did not updated: %s", fn, cv.Diagnosis) }

	ct  = cv.Update("feedbacktype", "dkim")
	cx++; if ct == false               { t.Errorf("%s(feedbacktype, dkim) returns false", fn) }
	cx++; if cv.FeedbackType != "dkim" { t.Errorf("%s(feedbacktype, dkim) did not updated: %s", fn, cv.FeedbackType) }

	ct  = cv.Update("lhost", "mx2.example.com")
	cx++; if ct == false                   { t.Errorf("%s(lhost, mx2.example.com) returns false", fn) }
	cx++; if cv.Lhost != "mx2.example.com" { t.Errorf("%s(lhost, mx2.example.com) did not updated: %s", fn, cv.Lhost) }

	ct  = cv.Update("reason", "blocked")
	cx++; if ct == false                   { t.Errorf("%s(reason, blocked) returns false", fn) }
	cx++; if cv.Reason != "blocked"        { t.Errorf("%s(reason, blocked) did not updated: %s", fn, cv.Reason) }

	ct  = cv.Update("recipient", "cat@example.org")
	cx++; if ct == false                       { t.Errorf("%s(recipient, cat@example.org) returns false", fn) }
	cx++; if cv.Recipient != "cat@example.org" { t.Errorf("%s(recipient, cat@example.org) did not updated: %s", fn, cv.Recipient) }

	ct  = cv.Update("replycode", "551")
	cx++; if ct == false           { t.Errorf("%s(replycode, 551) returns false", fn) }
	cx++; if cv.ReplyCode != "551" { t.Errorf("%s(replycode, 551) did not updated: %s", fn, cv.ReplyCode) }

	ct  = cv.Update("rhost", "mx4.example.com")
	cx++; if ct == false                   { t.Errorf("%s(rhost, mx4.example.com) returns false", fn) }
	cx++; if cv.Rhost != "mx4.example.com" { t.Errorf("%s(rhost, mx4.example.com) did not updated: %s", fn, cv.Rhost) }

	ct  = cv.Update("spec", "X-UNIX")
	cx++; if ct == false         { t.Errorf("%s(spec, X-UNIX) returns false", fn) }
	cx++; if cv.Spec != "X-UNIX" { t.Errorf("%s(spec, X-UNIX) did not updated: %s", fn, cv.Spec) }

	ct  = cv.Update("status", "5.1.2")
	cx++; if ct == false          { t.Errorf("%s(status, 5.1.2) returns false", fn) }
	cx++; if cv.Status != "5.1.2" { t.Errorf("%s(status, 5.1.2) did not updated: %s", fn, cv.Status) }

	// AsRFC1894
	fn = "AsRFC1894"
	ae = []string{
		"action", "status", "arrival-date", "last-attempt-date", "diagnostic-code", "final-recipient",
		"original-recipient", "x-actual-recipient", "reporting-mta", "remote-mta", "received-from-mta",
	}

	cx++; if cv.AsRFC1894("")    != "" { t.Errorf("%s() did not return an empty string: %s", fn, cv.AsRFC1894("")) }
	cx++; if cv.AsRFC1894("cat") != "" { t.Errorf("%s(cat) did not return an empty string: %s", fn, cv.AsRFC1894("cat")) }
	for _, e := range ae {
		cx++; if cv.AsRFC1894(e) == "" { t.Errorf("%s(%s) returns empty", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

