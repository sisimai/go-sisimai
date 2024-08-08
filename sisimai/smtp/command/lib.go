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

func list() []string { return []string{"HELO", "EHLO", "AUTH", "MAIL", "RCPT", "DATA", "QUIT"} }
func Find(argv0 string) string {
	// @param    string argv0  Text including SMTP command
	// @return   string        Found SMTP command
	if len(argv0) == 0 { return "" }
	found := ""
	words := strings.Split(argv0, " ")

	seeks: for _, e := range list() {
		// Find an SMTP command from the string
		for _, w := range words {
			// splitted by " "
			if e == w { found = e; break seeks }
		}
	}
	return found
}

