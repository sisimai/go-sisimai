// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis
import "sisimai/rfc1894"

var Fields1894 = rfc1894.FIELDTABLE()
type DeliveryMatter struct {
	Action       string     // The value of Action header
	Agent        string     // MTA name
	Alias        string     // The value of alias entry(RHS)
	Command      string     // SMTP command in the message body
	Date         string     // The value of Last-Attempt-Date header
	Diagnosis    string     // The value of Diagnostic-Code header
	FeedbackType string     // Feedback type
	Lhost        string     // The value of Received-From-MTA header
	Reason       string     // Temporary reason of bounce
	Recipient    string     // The value of Final-Recipient header
	ReplyCode    string     // SMTP Reply Code
	Rhost        string     // The value of Remote-MTA header
	HardBounce   bool       // Hard bounce or not
	Spec         string     // Protocl specification
	Status       string     // The value of Status header
}

// Set()
func(this *DeliveryMatter) Set(argv0, argv1 string) bool {
	if len(Fields1894[argv0]) == 0 { return false }
	switch argv0 {
		// Available values are the followings:
		// - "action":             Action    (list)
		// - "arrival-date":       Date      (date)
		// - "diagnostic-code":    Diagnosis (code)
		// - "final-recipient":    Recipient (addr)
		// - "last-attempt-date":  Date      (date)
		// - "original-recipient": Alias     (addr)
		// - "received-from-mta":  Lhost     (host)
		// - "remote-mta":         Rhost     (host)
		// - "reporting-mta":      Rhost     (host)
		// - "status":             Status    (stat)
		// - "x-actual-recipient": Alias     (addr)
		default: return false
		case "action":
			// Action: failed
			this.Action = argv1

		case "arrival-date", "last-attempt-date":
			// Arrival-Date: Thu, 29 Apr 2019 23:34:45 +0900 (JST)
			// Last-Attempt-Date: Mon, 21 May 2018 16:10:00 +0900
			this.Date = argv1

		case "diagnostic-code":
			// Diagnostic-Code: smtp; 550 DMARC check failed.
			this.Diagnosis = argv1

		case "final-recipient":
			// Final-Recipient: RFC822; kijitora@nyaan.jp
			this.Recipient = argv1

		case "original-recipient", "x-actual-recipient":
			// X-Actual-Recipient: RFC822; kijitora@example.co.jp
			// X-Actual-Recipient: X-Unix; |/var/adm/sm.bin/neko
			this.Alias = argv1

		case "received-from-mta":
			// Received-From-MTA: DNS; p225-ix4.kyoto.example.ne.jp
			this.Lhost = argv1
		
		case "remote-mta", "reporting-mta":
			// Remote-MTA: dns; mx-aol.mail.gm0.yahoodns.net
			// Reporting-MTA: dsn; d217-29.smtp-out.amazonses.com
			this.Rhost = argv1

		case "status":
			// Status: 5.1.1
			this.Status = argv1
	}
	return true
}

// Get()
func(this *DeliveryMatter) Get(argv0 string) string {
	return ""
}

