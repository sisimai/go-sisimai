// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/
import "strings"

// Undisclosed() returns a pseudo recipient address or a pseudo sender address
func Undisclosed(f bool) string {
	// @param    string f   Address type: true = recipient, false = sender
	// @return   string     Pseudo recipient address or sender address when the "a" is neither "r" nor "s"
	p := "recipient"; if f == false { p = "sender" }
	return "undisclosed-" + p + "-in-headers@libsisimai.org.invalid"
}

// Final() returns a string processed by Ruleset 4 in sendmail.cf
func Final(argv0 string) string {
	// @param    string argv0  String including an email address like "<neko@example.jp>"
	// @return   string        String without angle brackets: "neko@example.jp"
	if argv0                     == "" { return "" }
	if strings.Count(argv0, "@") != 1  { return argv0 }

	for strings.HasPrefix(argv0, "<") { argv0 = strings.Trim(argv0, "<") }
	for strings.HasSuffix(argv0, ">") { argv0 = strings.Trim(argv0, ">") }

	atmark := strings.LastIndex(argv0, "@")
	useris := argv0[0:atmark]
	hostis := argv0[atmark+1:]

	if IsQuotedAddress(argv0) == false {
		// Remove all the angle brackets from the local part
		useris = strings.ReplaceAll(useris, "<", "")
		useris = strings.ReplaceAll(useris, ">", "")
	}

	// Remove all the angle brackets from the domain part
	hostis = strings.ReplaceAll(hostis, "<", "")
	hostis = strings.ReplaceAll(hostis, ">", "")
	return useris + "@" + hostis
}

