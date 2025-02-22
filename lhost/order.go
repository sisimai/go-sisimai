// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _               _      _____          _           
// | | |__   ___  ___| |_   / / _ \ _ __ __| | ___ _ __ 
// | | '_ \ / _ \/ __| __| / / | | | '__/ _` |/ _ \ '__|
// | | | | | (_) \__ \ |_ / /| |_| | | | (_| |  __/ |   
// |_|_| |_|\___/|___/\__/_/  \___/|_|  \__,_|\___|_|   

package lhost
import "strings"

// OrderBySubject() returns an MTA Order decided by the first word of the "Subject": header.
func OrderBySubject(title string) []string {
	// @param   [string] title Subject header string
	// @return  [[]string]     Order of MTA functions
	if title == "" { return []string{} }

	table := map[string][]string {
		"abuse-report":          []string{"ARF"},
		"auto":                  []string{"RFC3834"},
		"auto-reply":            []string{"RFC3834"},
		"automatic-reply":       []string{"RFC3834"},
		"aws-notification":      []string{"AmazonSES"},
		"complaint-about":       []string{"ARF"},
		"delivery-failure":      []string{"Domino", "X2"},
		"delivery-notification": []string{"MessagingServer"},
		"delivery-status":       []string{"OpenSMTPD", "GoogleWorkspace", "Gmail", "GoogleGroups", "AmazonSES", "X3"},
		"dmarc-ietf-dmarc":      []string{"ARF"},
		"email-feedback":        []string{"ARF"},
		"failed-delivery":       []string{"X2"},
		"failure-delivery":      []string{"X2"},
		"failure-notice":        []string{"qmail", "mFILTER", "Activehunter"},
		"loop-alert":            []string{"FML"},
		"mail-could":            []string{"InterScanMSS"},
		"mail-delivery":         []string{"Exim", "DragonFly", "GMX", "Zoho", "EinsUndEins"},
		"mail-failure":          []string{"Exim"},
		"mail-system":           []string{"EZweb"},
		"message-delivery":      []string{"MailFoundry"},
		"message-frozen":        []string{"Exim"},
		"non-recapitabile":      []string{"Exchange2007"},
		"non-remis":             []string{"Exchange2007"},
		"notice":                []string{"Courier"},
		"postmaster-notify":     []string{"Sendmail"},
		"returned-mail":         []string{"Sendmail", "Biglobe", "V5sendmail", "X1"},
		"there-was":             []string{"X6"},
		"undeliverable":         []string{"Exchange2007", "Exchange2003"},
		"undeliverable-mail":    []string{"MailMarshalSMTP", "IMailServer"},
		"undeliverable-message": []string{"Notes", "Verizon"},
		"undelivered-mail":      []string{"Postfix", "Zoho"},
		"warning":               []string{"Sendmail", "Exim"},
	}

	// The following order is decided by the first 2 words of Subject: header
	for _, e := range []string{"[", "]", "_"} { title = strings.Replace(title, e, " ", -1) }

	// Squeeze duplicated space characters
	for strings.Contains(title, "  ") { title = strings.ReplaceAll(title, "  ", " ") }

	title  = strings.TrimSpace(title) // Remove leading space characters
	words := strings.SplitN(strings.ToLower(title), " ", 3)
	first := ""

	if word0 := strings.IndexByte(words[0], ':'); word0 > 0 {
		// Undeliverable: ..., notify: ...
		first = strings.ToLower(title[:word0])

	} else {
		// Postmaster notify, returned mail, ...
		first = strings.Join(words[0:2], "-")
	}
	for _, e := range []string{`:`, `,`, `*`, `"`} { first = strings.ReplaceAll(first, e, "") }
	return table[first]
}

// AnotherOrder() returns MTA functions list as a spare
func AnotherOrder() []string {
	// @param
	// @return  [[]string] Ordered MTA functions list
	return []string{
		// There are another patterns in the value of "Subject:" header of a bounce mail generated by
		// the following MTA/ESP functions
		"Exim", "Sendmail", "Exchange2007", "Exchange2003", "AmazonSES", "InterScanMSS", "KDDI",
		"Verizon", "ApacheJames", "FML", "X2",

		// The following is a fallback list
		"Postfix", "OpenSMTPD", "Courier", "qmail", "MessagingServer", "MailMarshalSMTP", "Domino",
		"Notes", "Gmail", "Zoho", "GMX", "GoogleGroups", "MailFoundry", "V5sendmail", "IMailServer", 
		"mFILTER", "Activehunter", "EZweb", "Biglobe", "EinsUndEins", "X1", "X3", "X6",
	}
}

