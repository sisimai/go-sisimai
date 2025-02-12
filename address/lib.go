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

// IsIncluded() returns true if the string include an email address
func IsIncluded(argv0 string) bool {
	// @param    string argv0    String including an email address like "<neko@example.jp>"
	// @return   bool            true:  is including an email address
	//                           false: is not including an email address
	if len(argv0) < 5 || strings.Contains(argv0,  "@") == false { return false }
	if strings.HasPrefix(argv0, "<") && strings.HasSuffix(argv0, ">") {
		// The argument is like "<neko@example.jp>"
		if rfc5322.IsEmailAddress(strings.Trim(argv0, "<>")) { return true }
		return false

	} else {
		// Such as "nekochan (kijitora) neko@example.jp"
		for _, e := range strings.Split(argv0, " ") {
			// Is there any email address string in each element?
			e = strings.Trim(e, "<>"); if rfc5322.IsEmailAddress(e) { return true }
		}
	}
	return false
}

// IsMailerDaemon() checks that the argument is mailer-daemon or not
func IsMailerDaemon(email string) bool {
	// @param    string email    Email address
	// @return   bool            true:  is a mailer-daemon
	//                           false: is not a mailer-daemon
	match := false
	value := strings.ToLower(email)
	table := []string{
		"mailer-daemon@", "(mailer-daemon)", "<mailer-daemon>", "mailer-daemon ",
		"postmaster@", "(postmaster)", "<postmaster>",
	}
	for _, e := range table {
		if strings.Contains(value, e) || value == "mailer-daemon" || value == "postmaster" { return true }
	}
	return match
}
