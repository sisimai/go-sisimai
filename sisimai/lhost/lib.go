// Copyright (C) 2020-2022,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost
import "net/mail"

type LhostDS struct {
	Action       string // The value of Action header
	Agent        string // MTA name
	Alias        string // The value of alias entry(RHS)
	Command      string // SMTP command in the message body
	Date         string // The value of Last-Attempt-Date header
	Diagnosis    string // The value of Diagnostic-Code header
	FeedbackType string // Feedback type
	Lhost        string // The value of Received-From-MTA header
	Reason       string // Temporary reason of bounce
	Recipient    string // The value of Final-Recipient header
	ReplyCode    uint   // SMTP Reply Code
	Rhost        string // The value of Remote-MTA header
	HardBounce   bool   // Hard bounce or not
	Spec         string // Protocl specification
	Status       string // The value of Status header
}

type LhostRR struct {
	DS     []LhostDS    // List of LhostDS structs
	RFC822 string       // The original message
}

// Keep each function for parsing a bounce mail
var LhostCode = map[string]func(mail.Header, *string) (LhostRR, error) {}

// DELIVERYSTATUS() returns a data structure for a bounce message.
func DELIVERYSTATUS() map[string]string {
	return map[string]string {
		"action": "",       // The value of Action header
		"agent": "",        // MTA name
		"alias": "",        // The value of alias entry(RHS)
		"command": "",      // SMTP command in the message body
		"date": "",         // The value of Last-Attempt-Date header
		"diagnosis": "",    // The value of Diagnostic-Code header
		"feedbacktype": "", // Feedback Type
		"lhost": "",        // The value of Received-From-MTA header
		"reason": "",       // Temporary reason of bounce
		"recipient": "",    // The value of Final-Recipient header
		"replycode": "",    // SMTP Reply Code
		"rhost": "",        // The value of Remote-MTA header
		"hardbounce": "",   // Hard bounce or not
		"spec": "",         // Protocl specification
		"status": "",       // The value of Status header
	}
}

// INDICATORS() returns flags for position variables used at MTA functions in sisimai/lhost.
func INDICATORS() map[string]uint8 {
	return map[string]uint8 {
		"deliverystatus": (1 << 1),
		"message-rfc822": (1 << 2),
	}
}

// INDEX() returns MTA functions list in sisimai/lhost sorted by Alphabetical order.
func INDEX() []string {
	return []string{
		"Activehunter", "Amavis", "AmazonSES", "AmazonWorkMail", "Aol", "ApacheJames", "Barracuda",
		"Bigfoot", "Biglobe", "Courier", "Domino", "DragonFly", "EZweb", "EinsUndEins", "Exchange2003",
		"Exchange2007", "Exim", "FML", "Facebook", "GMX", "GSuite", "GoogleGroups", "Gmail",
		"IMailServer", "InterScanMSS", "KDDI", "MXLogic", "MailFoundry", "MailMarshalSMTP", "MailRu",
		"McAfee", "MessageLabs", "MessagingServer", "Notes", "Office365", "OpenSMTPD", "Outlook",
		"Postfix", "PowerMTA", "ReceivingSES", "SendGrid", "Sendmail", "SurfControl", "V5sendmail",
		"Verizon", "X1", "X2", "X3", "X4", "X5", "X6", "Yahoo", "Yandex", "Zoho", "mFILTER", "qmail",
	}
}

// Rise() is a wrapper function for calling each MTA functions in sisimai/lhost.
func Rise(mhead mail.Header, mbody *string) (LhostRR, error) {
	rr := LhostRR{}

	lhostorder := OrderBySubject(mhead["Subject"][0])
	lhostorder  = append(lhostorder, AnotherOrder() ...)

	for _, e := range lhostorder {
		_, ok := LhostCode[e]
		if !ok { continue } // TODO: Remove this line after we've implement sisimai/lhost pakcage

		q, oops := LhostCode[e](mhead, mbody); if oops != nil { continue }
		if len(q.DS) == 0 { continue }
		rr = q
	}

	return rr, nil
}

