// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lda

//  _     ____    _    
// | |   |  _ \  / \   
// | |   | | | |/ _ \  
// | |___| |_| / ___ \ 
// |_____|____/_/   \_\
import "strings"
import "libsisimai.org/sisimai/sis"
import sisimoji "libsisimai.org/sisimai/string"

var LocalAgent = map[string][]string{
	// Each error message should be a lower-cased string
	// dovecot/src/deliver/deliver.c
	// 11: #define DEFAULT_MAIL_REJECTION_HUMAN_REASON \
	// 12: "Your message to <%t> was automatically rejected:%n%r"
	"dovecot":    []string{"your message to <", "> was automatically rejected:"},
	"mail.local": []string{"mail.local: "},
	"procmail":   []string{"procmail: ", "/procmail "},
	"maildrop":   []string{"maildrop: "},
	"vpopmail":   []string{"vdelivermail: "},
	"vmailmgr":   []string{"vdeliver: "},
}
var MessagesOf = map[string]map[string][]string{
	// Each error message should be a lower-cased string
	"dovecot": map[string][]string{
		"mailboxfull": []string{
			"not enough disk space",
			"quota exceeded", // Dovecot 1.2 dovecot/src/plugins/quota/quota.c
			"quota exceeded (mailbox for user is full)", // dovecot/src/plugins/quota/quota.c
		},
		"userunknown": []string{"mailbox doesn't exist: "},
	},
	"mail.local": map[string][]string{
		"mailboxfull": []string{
			"disc quota exceeded",
			"mailbox full or quota exceeded",
		},
		"systemerror": []string{"temporary file write error"},
		"userunknown": []string{
			": invalid mailbox path",
			": unknown user:",
			": user missing home directory",
			": user unknown",
		},
	},
	"procmail": map[string][]string{
		"mailboxfull": []string{"quota exceeded while writing", "user over quota"},
		"systemerror": []string{"service unavailable"},
		"systemfull":  []string{"no space left to finish writing"},
	},
	"maildrop": map[string][]string{
		"mailboxfull": []string{"maildir over quota."},
		"userunknown": []string{
			"cannot find system user",
			"invalid user specified.",
		},
	},
	"vpopmail": map[string][]string{
		"filtered":    []string{"user does not exist, but will deliver to "},
		"mailboxfull": []string{"domain is over quota", "user is over quota"},
		"suspend":     []string{"account is locked email bounced"},
		"userunknown": []string{"sorry, no mailbox here by that name."},
	},
	"vmailmgr": map[string][]string{
		"mailboxfull": []string{"delivery failed due to system quota violation"},
		"userunknown": []string{
			"invalid or unknown base user or domain",
			"invalid or unknown virtual user",
			"user name does not refer to a virtual user",
		},
	},
}

// Find() detects the bounce reason from the error message generated by Local Delivery Agent
func Find(fo *sis.Fact) string {
	// @param    *sis.Fact fo    Struct to be detected the reason
	// @return   string          Bounce reason name or an empty string
	if fo == nil || fo.DiagnosticCode == ""     { return "" }
	if fo.Command != "" && fo.Command != "DATA" { return "" }

	deliversby := "" // LDA; Local Delivery Agent name
	reasontext := "" // Detected bounce reason
	issuedcode := strings.ToLower(fo.DiagnosticCode)

	for e := range LocalAgent {
		// Find a local delivery agent name from the lower-cased error message
		if sisimoji.ContainsAny(issuedcode, LocalAgent[e]) == false { continue }
		deliversby = e; break
	}
	if deliversby == "" { return "" }

	for e := range MessagesOf[deliversby] {
		// The key nane is a bounce reason name
		if sisimoji.ContainsAny(issuedcode, MessagesOf[deliversby][e]) == false { continue }
		reasontext = e; break
	}

	// procmail: Couldn't create "/var/mail/tmp.nekochan.22"
	if reasontext == "" { reasontext = "mailererror" }
	return reasontext
}

