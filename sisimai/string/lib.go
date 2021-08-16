// Copyright (C) 2020-2021 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string
import "strings"

// Squeeze() remove redundant characters
func Squeeze(argv0 string, chars string) string {
	// @param    [string] argvs  Email address, name, and other elements
	// @return   [EmailAddress]    EmailAddress struct when the email address was not valid
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

