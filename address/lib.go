// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/

// Package "address" provide functions related to an email address
package address

import "strings"
import "libsisimai.org/sisimai/rfc5322"

// Undisclosed returns a pseudo recipient or sender address.
//   Arguments:
//     - argv0 (bool): Address type; true = recipient, false = sender
//   Returns:
//     - (string):     Generated pseudo recipient or sender address
func Undisclosed(argv0 bool) string {
	p := "recipient"; if argv0 == false { p = "sender" }
	return "undisclosed-" + p + "-in-headers@libsisimai.org.invalid"
}

// Final returns a string processed by the ruleset 4 in sendmail.cf file.
//   Arguments:
//     - argv0 (string): String including an email address like "<neko@example.jp>"
//   Returns:
//     - (string):       Email address without angle brackets such as "neko@example.jp"
func Final(argv0 string) string {
	if strings.Count(argv0, "@") != 1  { return argv0 }

	for strings.HasPrefix(argv0, "<") { argv0 = strings.Trim(argv0, "<") }
	for strings.HasSuffix(argv0, ">") { argv0 = strings.Trim(argv0, ">") }

	atmark := strings.LastIndex(argv0, "@")
	useris := argv0[0:atmark]
	hostis := argv0[atmark+1:]

	if rfc5322.IsQuotedAddress(argv0) == false {
		// Remove all the angle brackets from the local part
		useris = strings.ReplaceAll(useris, "<", "")
		useris = strings.ReplaceAll(useris, ">", "")
	}

	// Remove all the angle brackets from the domain part
	hostis = strings.ReplaceAll(hostis, "<", "")
	hostis = strings.ReplaceAll(hostis, ">", "")
	return useris + "@" + hostis
}

// IsIncluded returns true if the string includes an email address.
//   Arguments:
//     - argv0 (string): String including an email address like "<neko@example.jp>"
//   Returns:
//     - (bool):         true if An email address is included in the given string
func IsIncluded(argv0 string) bool {
	if len(argv0) < 5 || strings.IndexByte(argv0,  '@') < 0 { return false }
	if strings.HasPrefix(argv0, "<") && strings.HasSuffix(argv0, ">") {
		// The argument is like "<neko@example.jp>"
		if rfc5322.IsEmailAddress(strings.Trim(argv0, "<>")) { return true }
		return false

	} else {
		// Such as "nekochan (kijitora) neko@example.jp"
		for _, e := range strings.Split(argv0, " ") {
			// Is there any email address string in each element?
			if rfc5322.IsEmailAddress(strings.Trim(e, "<>")) { return true }
		}
	}
	return false
}

// IsMailerDaemon checks that the argument is mailer-daemon address or not.
//   Arguments:
//     - argv0 (string): Email address
//   Returns:
//     - (bool):         true if an email address is a mailer-dameon or postmaster address
func IsMailerDaemon(argv0 string) bool {
	value := strings.ToLower(argv0)
	table := []string{
		"mailer-daemon@", "(mailer-daemon)", "<mailer-daemon>", "mailer-daemon ",
		"postmaster@", "(postmaster)", "<postmaster>",
	}
	for _, e := range table {
		if strings.Contains(value, e) || value == "mailer-daemon" || value == "postmaster" { return true }
	}
	return false
}
