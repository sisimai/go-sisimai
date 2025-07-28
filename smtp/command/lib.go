// Copyright (C) 2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                _           __                                            _ 
//  ___ _ __ ___ | |_ _ __   / /__ ___  _ __ ___  _ __ ___   __ _ _ __   __| |
// / __| '_ ` _ \| __| '_ \ / / __/ _ \| '_ ` _ \| '_ ` _ \ / _` | '_ \ / _` |
// \__ \ | | | | | |_| |_) / / (_| (_) | | | | | | | | | | | (_| | | | | (_| |
// |___/_| |_| |_|\__| .__/_/ \___\___/|_| |_| |_|_| |_| |_|\__,_|_| |_|\__,_|
//                   |_|                                                      

// Package "smtp/command" provides functions related to SMTP commands.
package command
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/moji"

var availables = []string{
	"HELO", "EHLO", "MAIL", "RCPT", "DATA", "QUIT", "RSET", "NOOP", "VRFY", "ETRN",
	"EXPN", "HELP", "AUTH", "STARTTLS", "XFORWARD",
	"CONN", // CONN is a pseudo SMTP command used only in Sisimai
}
var detectable = []string{
	"HELO", "EHLO", "STARTTLS", "AUTH PLAIN", "AUTH LOGIN", "AUTH CRAM-", "AUTH DIGEST-",
	"MAIL F", "RCPT", "RCPT T", "DATA", "QUIT", "XFORWARD",
}
var ExceptDATA = []string{"CONN", "EHLO", "HELO", "MAIL", "RCPT"}

// Test checks that an SMTP command in the argument is valid or not.
//   Arguments:
//     - argv0 (string): An SMTP command.
//   Returns:
//     - (bool): true if the argument is a valid SMTP command.
func Test(argv0 string) bool {
	if len(argv0) < 4                      { return false }
	if moji.ContainsAny(argv0, availables) { return true  }
	return false
}

// Find returns an SMTP command found in the argument.
//   Arguments:
//     - argv0 (string): Text including SMTP command.
//   Returns:
//     - (string): Found SMTP command.
func Find(argv0 string) string {
	if Test(argv0) == false { return "" }

	commandset := make([]string, 0, 4)
	commandmap := map[string]string{"STAR": "STARTTLS", "XFOR": "XFORWARD"}
	issuedcode := " " + argv0 + " "

	for _, e := range detectable {
		// Find an SMTP command from the given string
		p0 := strings.Index(argv0, e); if p0 < 0 { continue }
		if strings.IndexByte(e, ' ') < 0 {
			// For example, "RCPT T" does not appear in an email address or a domain name
			cx, cw := true, len(e) + 1
			ca, cz := []byte(issuedcode[p0:p0 + 1])[0], []byte(issuedcode[p0 + cw:p0 + cw + 1])[0]
			switch {
				// Exclude an SMTP command in the part of an email address, a domain name, such as
				// DATABASE@EXAMPLE.JP, EMAIL.EXAMPLE.COM, and so on.
				case ca > 47 && ca <  58 || cz > 47 && cz <  58: // 0-9
				case ca > 63 && ca <  91 || cz > 63 && cz <  91: // @-Z
				case ca > 96 && ca < 123 || cz > 96 && cz < 123: // `-z
				default: cx = false
			}
			if cx == true { continue }
		}
		smtpc := e[0:4] // The first 4 characters of SMTP command found in the argument

		if moji.HasPrefixAny(smtpc, commandset) { continue }
		if slices.Contains([]string{"STAR", "XFOR"}, smtpc) { smtpc = commandmap[smtpc] }
		commandset = append(commandset, smtpc)
	}
	if len(commandset) == 0 { return "" }
	return commandset[len(commandset)-1]
}

