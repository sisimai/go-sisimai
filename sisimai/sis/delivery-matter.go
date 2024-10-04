// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis
//  ____       _ _                      __  __       _   _            
// |  _ \  ___| (_)_   _____ _ __ _   _|  \/  | __ _| |_| |_ ___ _ __ 
// | | | |/ _ \ | \ \ / / _ \ '__| | | | |\/| |/ _` | __| __/ _ \ '__|
// | |_| |  __/ | |\ V /  __/ |  | |_| | |  | | (_| | |_| ||  __/ |   
// |____/ \___|_|_| \_/ \___|_|   \__, |_|  |_|\__,_|\__|\__\___|_|   
//                                |___/                               
import "strings"
import "sisimai/address"
import "sisimai/rfc1894"
import "sisimai/smtp/status"

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

// Set() substitutes the argv1 as a value into the member related to argv0
func(this *DeliveryMatter) Set(argv0, argv1 string) bool {
	// @param    string argv0  A key name related to the member of DeliveryMatter struct
	// @param    string argv1  The value to be substituted
	// @return   bool          Returns true if it succesufully substituted
	if len(argv0)             == 0 { return false }
	if len(argv1)             == 0 { return false }
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
			// Final-Recipient: RFC822; <kijitora@nyaan.jp>
			v := address.S3S4(argv1); if len(v) == 0 { return false }
			this.Recipient = v

		case "original-recipient", "x-actual-recipient":
			// X-Actual-Recipient: RFC822; kijitora@example.co.jp
			// X-Actual-Recipient: X-Unix; |/var/adm/sm.bin/neko
			v := address.S3S4(argv1); if len(v) == 0 { return false }
			this.Alias = v

		case "received-from-mta":
			// Received-From-MTA: DNS; p225-ix4.kyoto.example.ne.jp
			// Received-From-MTA: dns; sironeko@example.jp
			if strings.Contains(argv1, "@") { argv1 = argv1[strings.Index(argv1, "@") + 1:] }
			this.Lhost = argv1

		case "remote-mta", "reporting-mta":
			// Remote-MTA: dns; mx-aol.mail.gm0.yahoodns.net
			// Remote-MTA: dns; 192.0.2.222 (192.0.2.222, the server for the domain.)
			// Reporting-MTA: dsn; d217-29.smtp-out.amazonses.com
			if strings.Contains(argv1, " ") { argv1 = strings.Split(argv1, " ")[0] }
			this.Rhost = argv1

		case "status":
			// Status: 5.1.1
			// Status: 4.2.2 (Over Quota)
			if strings.Contains(argv1, " ") { argv1 = argv1[0:strings.Index(argv1, " ")] }
			if status.Test(argv1) == false  { return false }
			this.Status = argv1
	}
	return true
}

// Get() returns the value of the member specified at argv0
func(this *DeliveryMatter) Get(argv0 string) string {
	// @param    string argv0  A key name related to the member of DeliveryMatter struct
	// @return   string        The value of the member specified at argv0
	if len(argv0)             == 0 { return "" }
	if len(Fields1894[argv0]) == 0 { return "" }

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
		case "action":    return this.Action
		case "alias":     return this.Alias
		case "date":      return this.Date
		case "diagnosis": return this.Diagnosis
		case "lhost":     return this.Lhost
		case "rhost":     return this.Rhost
		case "recipient": return this.Recipient
		case "status":    return this.Status
		default:          return ""
	}
}

