// Copyright (C) 2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
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

func Find(argv0 string) string {
	// @param    string argv0  Text including SMTP command
	// @return   string        Found SMTP command
	if Test(argv0) == false { return "" }

	commandset := []string{}
	commandmap := map[string]string{"STAR": "STARTTLS", "XFOR": "XFORWARD"}
	detectable := []string{
		"HELO", "EHLO", "STARTTLS", "AUTH PLAIN", "AUTH LOGIN", "AUTH CRAM-", "AUTH DIGEST-",
		"MAIL F", "RCPT", "RCPT T", "DATA", "QUIT", "XFORWARD",
	}

	for _, e := range detectable {
		// Find an SMTP command from the given string
		if strings.Contains(argv0, e) == false { continue }

		smtpc := e[0:4] // The first 4 characters of SMTP command found in the argument
		found := false  // There is the same SMTP command in "commandset" or not
		for _, c := range commandset {
			// Check that the command found in the argument is already included in "commandset"
			if strings.HasPrefix(c, smtpc) == false { continue }
			found = true
			break
		}
		if found { continue } // There is the same SMTP command in "commandset"

		if smtpc == "STAR" || smtpc == "XFOR" { smtpc = commandmap[smtpc] }
		commandset = append(commandset, smtpc)
	}
	if len(commandset) == 0 { return "" }
	return commandset[len(commandset)-1]
}

