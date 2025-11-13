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
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/moji"

var availables = []string{
	eb.CeHELO, eb.CeEHLO, eb.CeMAIL, eb.CeRCPT, eb.CeDATA, eb.CeQUIT, eb.CeRSET, eb.CeNOOP,
	eb.CeVRFY, eb.CeETRN, eb.CeEXPN, eb.CeHELP, eb.CeAUTH, eb.CeTTLS, eb.CeXFWD, eb.CeCONN,
}
var detectable = []string{
	eb.CeHELO, eb.CeEHLO, eb.CeTTLS, eb.CeAUTH + " PLAIN", eb.CeAUTH + " LOGIN",
	eb.CeAUTH + " CRAM-", eb.CeAUTH + " DIGEST-", eb.CeMAIL + " F", eb.CeRCPT, eb.CeRCPT + " T",
	eb.CeDATA, eb.CeQUIT, eb.CeXFWD,
}
var ExceptDATA = []string{eb.CeCONN, eb.CeEHLO, eb.CeHELO, eb.CeMAIL, eb.CeRCPT}

// Test checks that an SMTP command in the argument is valid or not.
//   Arguments:
//     - comm (string): An SMTP command.
//   Returns:
//     - (bool): true if the argument is a valid SMTP command.
func Test(comm string) bool {
	if len(comm) < 4                      { return false }
	if moji.ContainsAny(comm, availables) { return true  }
	return false
}

// Find returns an SMTP command found in the argument.
//   Arguments:
//     - text (string): Text including SMTP command.
//   Returns:
//     - (string): Found SMTP command.
func Find(text string) string {
	if Test(text) == false { return "" }

	commandset := make([]string, 0, 4)
	commandmap := map[string]string{"STAR": eb.CeTTLS, "XFOR": eb.CeXFWD}
	issuedcode := " " + text + " "

	for _, e := range detectable {
		// Find an SMTP command from the given string
		p0 := strings.Index(text, e); if p0 < 0 { continue }
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

