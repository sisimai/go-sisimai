// Copyright (C) 2020-2022,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _   
// | | |__   ___  ___| |_ 
// | | '_ \ / _ \/ __| __|
// | | | | | (_) \__ \ |_ 
// |_|_| |_|\___/|___/\__|
import "sisimai/sis"

// Keep each function for decoding a bounce mail
var InquireFor = map[string]func(*sis.BeforeFact) sis.RisingUnderway {}

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
		"Activehunter", "AmazonSES", "ApacheJames", "Biglobe", "Courier", "Domino",
		"DragonFly", "EZweb", "EinsUndEins", "Exchange2003", "Exchange2007", "Exim", "FML", "GMX",
		"GoogleGroups", "Gmail", "GoogleWorkspace", "IMailServer", "InterScanMSS", "KDDI", "MailFoundry", "MailMarshalSMTP",
		"MessagingServer", "Notes", "Office365", "OpenSMTPD", "Postfix", "Sendmail", "V5sendmail",
		"Verizon", "X1", "X2", "X3", "X6", "Zoho", "mFILTER", "qmail",
	}
}

