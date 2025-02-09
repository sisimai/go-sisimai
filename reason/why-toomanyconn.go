// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____           __  __                    ____                  
// |_   _|__   ___ |  \/  | __ _ _ __  _   _ / ___|___  _ __  _ __  
//   | |/ _ \ / _ \| |\/| |/ _` | '_ \| | | | |   / _ \| '_ \| '_ \ 
//   | | (_) | (_) | |  | | (_| | | | | |_| | |__| (_) | | | | | | |
//   |_|\___/ \___/|_|  |_|\__,_|_| |_|\__, |\____\___/|_| |_|_| |_|
//                                     |___/                        
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/smtp/status"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["TooManyConn"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }

		index := []string{
			"all available ips are at maximum connection limit", // SendGrid
			"connection rate limit exceeded",
			"exceeds per-domain connection limit for",
			"has exceeded the max emails per hour ",
			"throttling failure: daily message quota exceeded",
			"throttling failure: maximum sending rate exceeded",
			"too many connections",
			"too many connections from your host.", // Microsoft
			"too many concurrent smtp connections", // Microsoft
			"too many errors from your ip",         // Free.fr
			"too many recipients",                  // ntt docomo
			"too many smtp sessions for this host", // Sendmail(daemon.c)
			"trop de connexions, ",
			"we have already made numerous attempts to deliver this message",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "toomanyconn" or not
	ProbesInto["TooManyConn"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is toomanyconn, false: is not toomanyconn
		if fo == nil                                       { return false }
		if fo.Reason == "toomanyconn"                      { return true  }
		if status.Name(fo.DeliveryStatus) == "toomanyconn" { return true  }
		return IncludedIn["TooManyConn"](strings.ToLower(fo.DiagnosticCode))
	}
}

