// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

// Retry() returns the table of reason list which should be checked again
func Retry() map[string]bool {
	return map[string]bool{
		"undefined": true, "onhold": true,  "systemerror": true, "securityerror": true, "expired": true,
		"suspend": true, "networkerror": true, "hostunknown": true, "userunknown": true,
    }
}

// Index() returns the list of all the reasons sisimai supoort
func Index() []string {
	return []string{
        "AuthFailure", "BadReputation", "Blocked", "ContentError", "ExceedLimit", "Expired", "Filtered",
		"HasMoved", "HostUnknown", "MailboxFull", "MailerError", "MesgTooBig", "NetworkError", "NotAccept",
		"NotCompliantRFC", "OnHold", "Rejected", "NoRelaying", "SpamDetected", "VirusDetected", "PolicyViolation",
		"SecurityError", "Speeding", "Suspend", "RequirePTR", "SystemError", "SystemFull", "TooManyConn",
		"UserUnknown", "SyntaxError",
    }
}

