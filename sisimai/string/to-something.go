// Copyright (C) 2020 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string
import "strings"

// ToLF() replace CR and CR/LF to LF.
func ToLF(argv0 *string) *string {
	// @param    [*string] argv0  Text including CR or CR/LF
	// @return   [*string]        LF converted text
	if len(*argv0) == 0 { return argv0 }

	replaceset := [...]string { "\r\n", "\r" }
	lfreplaced := ""

	for _, e := range replaceset {
		// Convert CRLF and CR to LF
		if !strings.Contains(*argv0, e) { continue }
		lfreplaced = strings.Replace(*argv0, e, "\n", -1)
	}
	if len(lfreplaced) == 0 { return argv0 }

	return &lfreplaced
}

// ToPlain() converts given HTML text to a plain text.
func ToPlain(argv0 *string, loose bool) *string {
	// @param    [*string] argv0  HTML text
	// @param    [bool]    loose  Loose check flag
	// @return   [*string]        Plain text
	if len(*argv0) == 0 { return argv0 }

	plain := ""
	xhtml := *argv0
	body0 := strings.Index(strings.ToLower(*argv0), "<body>")

	// Remove HTML header part
	if body0 > -1 { xhtml = xhtml[body0 + len("<body>"):len(xhtml)] }

	for _, e := range strings.Split(xhtml, "</") {
		// Find ">" from HTML element and remove string until ">"
		for {
			if strings.HasPrefix(e, "body>") { break }
			if strings.HasPrefix(e, "html>") { break }

			p := strings.Index(e, "<")
			q := strings.Index(e, ">")

			if p > -1 && p < q { plain += e[0:p - 1] }
			if p == -1 || q == -1 { plain += e }
			e  = e[q + 1:len(e)]
			break
		}
	}
	return &plain
}

// ToUTF8() converts an encoded text to UTF8 text
func ToUTF8(argv0 *string) *string {
	// @param    [*string] argv0  Some encoded text
	// @return   [*string]        UTF-8 text

	// TODO: IMPLEMENT
	utf8 := ""
	return &utf8
}

