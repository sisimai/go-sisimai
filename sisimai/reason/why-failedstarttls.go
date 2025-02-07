// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____     _ _          _ ____ _____  _    ____ _____ _____ _     ____  
// |  ___|_ _(_) | ___  __| / ___|_   _|/ \  |  _ \_   _|_   _| |   / ___| 
// | |_ / _` | | |/ _ \/ _` \___ \ | | / _ \ | |_) || |   | | | |   \___ \ 
// |  _| (_| | | |  __/ (_| |___) || |/ ___ \|  _ < | |   | | | |___ ___) |
// |_|  \__,_|_|_|\___|\__,_|____/ |_/_/   \_\_| \_\|_|   |_| |_____|____/ 
import "strings"
import "sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["FailedSTARTTLS"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		index := []string{
			"starttls is required to send mail",
			"tls required but not supported", // SendGrid:the recipient mailserver does not support TLS or have a valid certificate
		}
		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "FailedSTARTTLS" or not
	ProbesInto["FailedSTARTTLS"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is FailedSTARTTLS, false: is not FailedSTARTTLS
		if fo.Reason == "failedstarttls"                                           { return true }
		if fo.Command == "STARTTLS"                                                { return true }
		if fo.ReplyCode == "523" || fo.ReplyCode == "524" || fo.ReplyCode == "538" { return true }
		return IncludedIn["FailedSTARTTLS"](strings.ToLower(fo.DiagnosticCode))
	}
}

