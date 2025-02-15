// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/

package address
import "strings"
import "libsisimai.org/sisimai/sis"

// Rise() is a constructor of sis.EmailAddress
func Rise(argvs [3]string) sis.EmailAddress {
	// @param    [3]string argvs  ["Email address", "name", "comment"]
	// @return   sis.EmailAddress EmailAddress struct when the email address was not valid
	if argvs[0] == "" { return sis.EmailAddress{} }

	thing := new(sis.EmailAddress)
	email := Final(argvs[0])

	if lasta := strings.LastIndex(email, "@"); lasta > 0 {
		// Get the local part and the domain part from the email address
		lpart := email[:lasta]     // Local part of the address:  "neko"
		dpart := email[lasta + 1:] // Domain part of the address: "example.jp"

		if other := ExpandVERP(email); other != "" {
			// The email address is a VERP address such as "neko+cat=example.jp@example.org"
			thing.Verp = other

		} else if other := ExpandAlias(email); other != "" {
			// The email address is an alias address such as "neko+cat@example.jp"
			thing.Alias = other
		}

		// Remove the folowing characters: "<", ">", ",", ".", and ";" from the email address
		lpart = strings.TrimLeft(lpart, "<");     thing.User = lpart
		dpart = strings.TrimRight(dpart, ">,.;"); thing.Host = dpart
		thing.Address = lpart + "@" + dpart

	} else {
		// The argument does not include "@"
		if IsMailerDaemon(argvs[0]) == false { return sis.EmailAddress{} }
		if strings.Contains(argvs[0], " ")   { return sis.EmailAddress{} }

		// The argument does not include " "
		thing.User    = argvs[0]
		thing.Address = thing.User
	}

	thing.Name    = argvs[1]
	thing.Comment = argvs[2]
	return *thing
}

