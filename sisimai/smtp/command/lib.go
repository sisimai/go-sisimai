// Copyright (C) 2021 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package command
import "strings"

func list() []string { return []string { "HELO", "EHLO", "AUTH", "MAIL", "RCPT", "DATA", "QUIT" } }
func Find(argv0 string) string {
	// @param    [string] argv0  Text including SMTP command
	// @return   [string]        Found SMTP command
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

