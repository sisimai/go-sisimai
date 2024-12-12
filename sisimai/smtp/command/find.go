// Copyright (C) 2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package command

//                _           __                                            _ 
//  ___ _ __ ___ | |_ _ __   / /__ ___  _ __ ___  _ __ ___   __ _ _ __   __| |
// / __| '_ ` _ \| __| '_ \ / / __/ _ \| '_ ` _ \| '_ ` _ \ / _` | '_ \ / _` |
// \__ \ | | | | | |_| |_) / / (_| (_) | | | | | | | | | | | (_| | | | | (_| |
// |___/_| |_| |_|\__| .__/_/ \___\___/|_| |_| |_|_| |_| |_|\__,_|_| |_|\__,_|
//                   |_|                                                      
import "strings"
import sisimoji "sisimai/string"

func Find(argv0 string) string {
	// @param    string argv0  Text including SMTP command
	// @return   string        Found SMTP command
	if Test(argv0) == false { return "" }

	commandset := []string{}
	commandmap := map[string]string{"STAR": "STARTTLS", "XFOR": "XFORWARD"}
	issuedcode := " " + argv0 + " "

	for _, e := range Detectable {
		// Find an SMTP command from the given string
		p0 := strings.Index(argv0, e); if p0 < 0 { continue }
		if strings.Contains(e, " ") == false {
			// For example, "RCPT T" does not appear in an email address or a domain name
			cx := true; for {
				// Exclude an SMTP command in the part of an email address, a domain name, such as
				// DATABASE@EXAMPLE.JP, EMAIL.EXAMPLE.COM, and so on.
				cw := len(e) + 1
				ca := []byte(issuedcode[p0:p0 + 1])[0]
				cz := []byte(issuedcode[p0 + cw:p0 + cw + 1])[0]

				if ca > 47 && ca <  58 || cz > 47 && cz <  58 { break } // 0-9
				if ca > 63 && ca <  91 || cz > 63 && cz <  91 { break } // @-Z
				if ca > 96 && ca < 123 || cz > 96 && cz < 123 { break } // `-z
				cx = false; break
			}
			if cx == true { continue }
		}
		smtpc := e[0:4] // The first 4 characters of SMTP command found in the argument

		if sisimoji.HasPrefixAny(smtpc, commandset) { continue }
		if smtpc == "STAR" || smtpc == "XFOR" { smtpc = commandmap[smtpc] }
		commandset = append(commandset, smtpc)
	}
	if len(commandset) == 0 { return "" }
	return commandset[len(commandset)-1]
}

