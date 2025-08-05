// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/

// Package "address" provide functions related to an email address.
package address
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc5322"

// Undisclosed returns a pseudo recipient or sender address.
//   Arguments:
//     - ctype (bool): Address type; true = recipient, false = sender.
//   Returns:
//     - (string): Generated pseudo recipient or sender address.
func Undisclosed(ctype bool) string {
	p := "recipient"; if ctype == false { p = "sender" }
	return "undisclosed-" + p + "-in-headers@libsisimai.org.invalid"
}

// Final returns a string processed by the ruleset 4 in sendmail.cf file.
//   Arguments:
//     - email (string): String including an email address like "<neko@example.jp>".
//   Returns:
//     - (string): Email address without angle brackets such as "neko@example.jp"
func Final(email string) string {
	if  strings.Count(email, "@") != 1 { return email }
	for strings.HasPrefix(email, "<")  { email = strings.Trim(email, "<") }
	for strings.HasSuffix(email, ">")  { email = strings.Trim(email, ">") }
	return email
}

// IsIncluded returns true if the string includes an email address.
//   Arguments:
//     - text (string): String including an email address like "<neko@example.jp>".
//   Returns:
//     - (bool): true if An email address is included in the given string.
func IsIncluded(text string) bool {
	if len(text) < 5 || strings.IndexByte(text,  '@') < 0 { return false }
	if strings.HasPrefix(text, "<") && strings.HasSuffix(text, ">") {
		// The argument is like "<neko@example.jp>"
		if rfc5322.IsEmailAddress(strings.Trim(text, "<>")) { return true }
		return false

	} else {
		// Such as "nekochan (kijitora) neko@example.jp"
		for _, e := range strings.Split(text, " ") {
			// Is there any email address string in each element?
			if rfc5322.IsEmailAddress(strings.Trim(e, "<>")) { return true }
		}
	}
	return false
}

// IsMailerDaemon checks that the argument is mailer-daemon address or not.
//   Arguments:
//     - email (string): Email address.
//   Returns:
//     - (bool): true if an email address is a mailer-dameon or postmaster address.
func IsMailerDaemon(email string) bool {
	value := strings.ToLower(email)
	names := []string{"mailer-daemon", "postmaster"}
	table := []string{
		"mailer-daemon@", "(mailer-daemon)", "<mailer-daemon>", "mailer-daemon ",
		"postmaster@", "(postmaster)", "<postmaster>",
	}
	return moji.ContainsAny(value, table) || slices.Contains(names, value)
}

