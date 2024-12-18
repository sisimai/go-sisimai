// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _ __ ___  __ _ ___  ___  _ __  
// | '__/ _ \/ _` / __|/ _ \| '_ \ 
// | | |  __/ (_| \__ \ (_) | | | |
// |_|  \___|\__,_|___/\___/|_| |_|
import "sisimai/sis"

// Keep each function (pointer) defined in sisimai/reason/*.go to check/detect the bounce reason
var IncludedIn = map[string]func(string) bool {}
var ProbesInto = map[string]func(*sis.Fact) bool {}

var GetRetried = Retry()
var ClassOrder = [][]string{
	[]string{
		"MailboxFull", "MesgTooBig", "ExceedLimit", "Suspend", "HasMoved", "NoRelaying", "AuthFailure",
		"UserUnknown", "Filtered", "RequirePTR", "NotCompliantRFC", "BadReputation", "Rejected",
		"ContentError", "HostUnknown", "SpamDetected", "Speeding", "TooManyConn", "Blocked",
    },
    []string{
		"MailboxFull", "AuthFailure", "BadReputation", "Speeding", "SpamDetected", "VirusDetected",
		"PolicyViolation", "NoRelaying", "SystemError", "NetworkError", "Suspend", "ContentError",
		"SystemFull", "NotAccept", "Expired", "SecurityError", "Suppressed", "MailerError",
    },
    []string{
		"MailboxFull", "MesgTooBig", "ExceedLimit", "Suspend", "UserUnknown", "Filtered", "Rejected",
		"HostUnknown", "SpamDetected", "Speeding", "TooManyConn", "Blocked", "SpamDetected", "AuthFailure",
		"SecurityError", "SystemError", "NetworkError", "Suspend", "Expired", "ContentError", "HasMoved",
		"SystemFull", "NotAccept", "MailerError", "NoRelaying", "Suppressed", "SyntaxError", "OnHold",
	},
}

// Retry() returns the table of reason list which should be checked again
func Retry() map[string]bool {
	return map[string]bool{
		"undefined": true, "onhold": true,  "systemerror": true, "securityerror": true,
		"expired": true, "networkerror": true, "hostunknown": true, "userunknown": true,
	}
}

// IsExplicit() returns false when the argument is empty or is "undefined" or is "onhold"
func IsExplicit(argv1 string) bool {
	// @param    string argv1  Reason name
	// @return   bool          false: The reaosn is not explicit
	if argv1 == "" || argv1 == "undefined" || argv1 == "onhold" { return false }
	return true
}

// Index() returns the list of all the reasons sisimai supoort
func Index() []string {
	return []string{
		"AuthFailure",      // Email rejected due to SPF, DKIM, DMARC failure
		"BadReputation",    // Email rejected due to an IP address reputation
		"Blocked",          // Email rejected due to client IP address or a hostname
		"ContentError",     // Email rejected due to a header format of the email
		// "Delivered",     // Email delivered successfully
		"ExceedLimit",      // Email rejected due to an email exceeded the limit
		"Expired",          // Delivery time has expired due to a connection failure
		"Feedback",         // Email forwarded to the sender as a complaint message from your mailbox provider
		"Filtered",         // Email rejected due to a header content after SMTP DATA command
		"HasMoved",         // Email rejected due to users mailbox has moved and is not forwarded automatically
		"HostUnknown",      // Delivery failed due to a domain part of a recipients email address does not exist
		"MailboxFull",      // Email rejected due to a recipients mailbox is full
		"MailerError",      // Email returned due to a mailer program has not exited successfully
		"MesgTooBig",       // Email rejected due to an email size is too big for a destination mail server
		"NetworkError",     // SMTP connection failed due to DNS look up failure or other network problems
		"NoRelaying",       // Email rejected with error message "Relaying Denied
		"NotAccept",        // Delivery failed due to a destination mail server does not accept any email
		"NotCompliantRFC",  // Email rejected due to non-compliance with RFC
		"OnHold",           // Sisimai could not decided the reason due to there is no (or less) detailed information for judging the reason
		"PolicyViolation",  // Email rejected due to policy violation on a destination host
		"Rejected",         // Email rejected due to a senders email address (envelope from)
		"RequirePTR",       // Email rejected due to missing PTR record or having invalid PTR record
		"SecurityError",    // Email rejected due to security violation was detected on a destination host
		"SpamDetected",     // Email rejected by spam filter running on the remote host
		"Speeding",         // Rejected due to exceeding a rate limit or sending too fast
		"Suppressed",       // Email was not delivered due to being listed in suppression list on MTA
		"Suspend",          // Email rejected due to a recipient account is being suspended
		"SyntaxError",      // Email rejected due to syntax error at sent commands in SMTP session
		"SystemError",      // Email returned due to system error on the remote host
		"SystemFull",       // Email rejected due to a destination mail servers disk is full
		"TooManyConn",      // SMTP connection rejected temporarily due to too many concurrency connections to the remote host
		// "Undefined",     // Sisimai could not detect an error reason
		"UserUnknown",      // Email rejected due to a local part of a recipients email address does not exist
		//"Vacation",       // Email replied automatically due to a recipient is out of office
		"VirusDetected",    // Email rejected due to a virus scanner on a destination host
	}
}

