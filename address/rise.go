// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/

package address
import "strings"
import "libsisimai.org/sisimai/v5/siba"

// Rise is a constructor of siba.EmailAddress.
//   Arguments:
//     - addrs ([3]string): Address slice such as `[3]string{"email address", "display name", "comment"}`.
//   Returns:
//     - (*siba.EmailAddress): EmailAddress struct when the email address is valid.
func Rise(addrs [3]string) *siba.EmailAddress {
	if addrs[0] == "" { return nil }

	thing := new(siba.EmailAddress)
	email := Final(addrs[0])

	if lasta := strings.LastIndex(email, "@"); lasta > 0 {
		// Get the local part and the domain part from the email address
		// - Local part of the address:  "neko"
		// - Domain part of the address: "example.jp"
		lpart, dpart := email[:lasta], email[lasta + 1:]

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
		if IsMailerDaemon(addrs[0]) == false || strings.IndexByte(addrs[0], ' ') > -1 { return nil }

		// The argument does not include " "
		thing.User    = addrs[0]
		thing.Address = thing.User
	}

	thing.Name    = addrs[1]
	thing.Comment = addrs[2]
	return thing
}

