// Copyright (C) 2020-2021 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string
import "strings"
import "strconv"

// Squeeze() remove redundant characters
func Squeeze(argv0 string, chars string) string {
	// @param    [string] argvs  Email address, name, and other elements
	// @return   [EmailAddress]  EmailAddress struct when the email address was not valid
	if len(argv0) == 0 { return ""    }
	if len(chars) == 0 { return argv0 }

	for _, e := range strings.Split(chars, "") {
		// Remove redundant characters from "argv1"
		for strings.Count(argv0, e + e) > 1 {
			// e + e => e
			e = strings.ReplaceAll(argv0, e + e, e)
		}
	}
	return argv0
}

// Sweep() clears the string out
func Sweep(argv0 string) string {
	// @param    [string] argv1  String to be cleaned
	// @return   [string]        Cleaned out string
	if len(argv0) == 0 { return "" }
	argv0 = Squeeze(argv0, " ")
	argv0 = strings.ReplaceAll(argv0, "\t", "")
	argv0 = strings.TrimSpace(argv0)

	if strings.Contains(argv0, " --") {
		// Delete all the string after a boundary string like " --neko-nyaan"
		for {
			if strings.Contains(argv0, "-- ")  { break }
			if strings.Contains(argv0, "--\t") { break }

			argv0 = argv0[0:strings.Index(argv0, " --")-1]
			break
		}
	}
	return argv0
}

// ContainsOnlyNumbers() returns true when the given string contain numbers only
func ContainsOnlyNumbers(argv0 string) bool {
	// @param    [string] argv0  String
	// @return   [bool]          true, false
	if len(argv0) == 0 { return false }

	match := true
	for _, e := range argv0 { if e < 48 || e > 57 { match = false; break } }

	return match
}

// IsIPv4Address() returns true when the given string is an IPv4 address
func IsIPv4Address(argv0 string) bool {
	// @param    [string] argv0  IPv4 address like "192.0.2.25"
	// @return   [bool]          true: is an IPv4 address, false: is not an IPv4 address
	if len(argv0) < 8                 { return false }
	if strings.Count(argv0, ".") != 3 { return false }

	match := true
	for _, e := range strings.Split(argv0, ".") {
		// Check each octet is between 0 and 255
		v, oops := strconv.Atoi(e)
		if oops != nil { match = false; break }
		if v < 0       { match = false; break }
		if v > 255     { match = false; break }
	}
	return match
}

