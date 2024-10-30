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

func Find(argv0 string) string {
	// @param    string argv0  Text including SMTP command
	// @return   string        Found SMTP command
	if Test(argv0) == false { return "" }

	commandset := []string{}
	commandmap := map[string]string{"STAR": "STARTTLS", "XFOR": "XFORWARD"}
	for _, e := range Detectable {
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

