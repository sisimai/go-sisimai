// Copyright (C) 2020-2022,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost
import "sisimai/sis"

// Keep each function for parsing a bounce mail
var LhostCode = map[string]func(*sis.BeforeFact) sis.RisingUnderway {}

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
/*
func Rise(mhead map[string][]string, mbody *string) LhostRR {

	var localhostr LhostRR
	var lhostorder []string = OrderBySubject(mhead["subject"][0])
	    lhostorder          = append(lhostorder, AnotherOrder() ...)

	for _, e := range lhostorder {
		// Call init() function of each MTA module in sisimai/lhost
		localhostr = LhostCode[e]
		if localhostr.DS
		_, ok := LhostCode[e]
		if !ok { continue } // TODO: Remove this line after we've implement sisimai/lhost pakcage

		q := LhostCode[e](mhead, mbody); if oops != nil { continue }
		if len(q.DS) == 0 { continue }
		rr = q
	}

	return rr, nil
}
*/

