// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost
import "strings"

// OrderBySubject() returns an MTA Order decided by the first word of the "Subject": header.
func OrderBySubject(title string) []string {
	// @param   [string] title Subject header string
	// @return  [[]string]     Order of MTA functions
	if len(title) == 0 { return []string {} }

	table := map[string][]string {
		"abuse-report":          []string{"ARF"},
		"auto":                  []string{"RFC3834"},
		"auto-reply":            []string{"RFC3834"},
		"automatic-reply":       []string{"RFC3834"},
		"aws-notification":      []string{"AmazonSES"},
		"complaint-about":       []string{"ARF"},
		"delivery-failure":      []string{"Domino", "X2"},
		"delivery-notification": []string{"MessagingServer" },
		"delivery-report":       []string{"PowerMTA" },
		"delivery-status":       []string{
			"GSuite", "Outlook", "GoogleGroups", "McAfee", "OpenSMTPD", "AmazonSES", "AmazonWorkMail",
			"ReceivingSES", "Gmail", "X3",
		},
		"dmarc-ietf-dmarc": []string{"ARF"},
		"email-feedback":   []string{"ARF"},
		"failed-delivery":  []string{"X2"},
		"failure-delivery": []string{"X2"},
		"failure-notice":   []string{"Yahoo", "qmail", "mFILTER", "Activehunter", "X4"},
		"loop-alert":       []string{"FML"},
		"mail-could":       []string{"InterScanMSS"},
		"mail-delivery":    []string{"Exim", "MailRu", "GMX", "EinsUndEins", "Zoho", "MessageLabs", "MXLogic"},
		"mail-failure":     []string{"Exim"},
		"mail-not":         []string{"X4"},
		"mail-system":      []string{"EZweb"},
		"message-delivery": []string{"MailFoundry"},
		"message-frozen":   []string{"Exim"},
		"message-you":      []string{"Barracuda"},
		"non-recapitabile": []string{"Exchange2007"},
		"non-remis":        []string{"Exchange2007"},
		"notice":           []string{"Courier"},
		"permanent-delivery": []string{"X4"},
		"postmaster-notify":  []string{"Sendmail"},
		"returned-mail": []string{"Sendmail", "Aol", "V5sendmail", "Bigfoot", "Biglobe", "X1"},
		"sorry-your":    []string{"Facebook"},
		"there-was":     []string{"X6"},
		"undeliverable": []string{"Office365", "Exchange2007", "Aol", "Exchange2003"},
		"undeliverable-mail":    []string{"Amavis", "MailMarshalSMTP", "IMailServer"},
		"undeliverable-message": []string{"Notes", "Verizon"},
		"undelivered-mail":      []string{"Postfix", "Aol", "SendGrid", "Zoho"},
		"warning":               []string{"Sendmail", "Exim"},
	}

	// The following order is decided by the first 2 words of Subject: header
	title = strings.Replace(title, "[", " ", -1)
	title = strings.Replace(title, "]", " ", -1)
	title = strings.Replace(title, "_", " ", -1)

	// Squeeze duplicated space characters
	for strings.Contains(title, "  ") { title = strings.ReplaceAll(title, "  ", " ") }

	title  = strings.TrimSpace(title)           // Remove leading space characters
	words := strings.SplitN(strings.ToLower(title), " ", 3)
	first := ""

	if strings.Index(words[0], ":") > 0 {
		// Undeliverable: ..., notify: ...
		first = strings.ToLower(title[0:strings.Index(words[0], ":")])

	} else {
		// Postmaster notify, returned mail, ...
		first = strings.Join(words[0:2], "-")
	}

	first = strings.ReplaceAll(first, `:`, "")
	first = strings.ReplaceAll(first, `,`, "")
	first = strings.ReplaceAll(first, `*`, "")
	first = strings.ReplaceAll(first, `"`, "")
	return table[first]
}

// AnotherOrder() returns MTA functions list as a spare
func AnotherOrder() []string {
	// @param
	// @return  [[]string] Ordered MTA functions list
	return []string {
		// There are another patterns in the value of "Subject:" header of a bounce mail generated by
		// the following MTA/ESP functions
		"MailRu", "Yandex", "Exim", "Sendmail", "Aol", "Office365", "Exchange2007", "Exchange2003",
		"AmazonWorkMail", "AmazonSES", "Barracuda", "InterScanMSS", "KDDI", "SurfControl", "Verizon",
		"ApacheJames", "X2", "X5", "FML",

		// The following is a fallback list
		"Postfix", "GSuite", "Yahoo", "Outlook", "GMX", "MessagingServer", "EinsUndEins", "Domino",
		"Notes", "qmail", "Courier", "OpenSMTPD", "Zoho", "MessageLabs", "MXLogic", "MailFoundry",
		"McAfee", "V5sendmail", "mFILTER", "SendGrid", "ReceivingSES", "Amavis", "PowerMTA", "GoogleGroups",
		"Gmail", "EZweb", "IMailServer", "MailMarshalSMTP", "Activehunter", "Bigfoot", "Biglobe",
		"Facebook", "X4", "X1", "X3", "X6",
	}
}

