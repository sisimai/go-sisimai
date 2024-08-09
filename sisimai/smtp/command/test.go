// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package command
import "strings"

// Test() checks that an SMTP command in the argument is valid or not
func Test(argv0 string) bool {
	// @param    string argv0  An SMTP command
	// @return   bool          false: Is not a valid SMTP command
	//                         true:  Is a valid SMTP command
	// @since v5.0.0
	if len(argv0) < 3 { return false }

	match := false
	table := []string{
		"HELO", "EHLO", "MAIL", "RCPT", "DATA", "QUIT", "RSET", "NOOP", "VRFY", "ETRN", "EXPN",
		"HELP", "AUTH", "STARTTLS", "XFORWARD",
		"CONN", // CONN is a pseudo SMTP command used only in Sisimai
	}
	for _, e := range table {
		// Find any SMTP command from the given argument
		if strings.Contains(argv0, e) == false { continue }
		match = true
		break
	}
	if match { return true }
	return false
}

