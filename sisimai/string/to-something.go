// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string

//      _        _             
//  ___| |_ _ __(_)_ __   __ _ 
// / __| __| '__| | '_ \ / _` |
// \__ \ |_| |  | | | | | (_| |
// |___/\__|_|  |_|_| |_|\__, |
//                       |___/ 
import "strings"

// ToLF() replace CR and CR/LF to LF.
func ToLF(argv0 *string) *string {
	// @param    *string argv0  Text including CR or CR/LF
	// @return   *string        LF converted text
	if len(*argv0) == 0 { return argv0 }

	characters := [2]string{"\r\n", "\r"}
	crreplaced := *argv0
	for _, e := range characters {
		// Convert CRLF and CR to LF
		if strings.Contains(crreplaced, e) == false { continue }
		crreplaced = strings.ReplaceAll(crreplaced, e, "\n")
	}

	if len(crreplaced) == 0 { return argv0 }
	return &crreplaced
}

// ToPlain() converts given HTML text to a plain text.
func ToPlain(argv0 *string, loose bool) *string {
	// @param    [*string] argv0  HTML text
	// @param    [bool]    loose  Loose check flag
	// @return   [*string]        Plain text
	if len(*argv0) == 0 { return argv0 }

	xhtml := *argv0
	lower := strings.ToLower(*argv0)
	plain := "" // Plain text (including no HTML element)
	body0 := -1 // Index of the beginning of the <body> element

	for _, e := range []string{">", " ", "\t", "\n"} {
		// Find the position of <body?, and remove the HTML header part
		body0  = strings.Index(lower, "<body" + e); if body0 < 0 { continue }
		body0 += len("<body>") + 1

		if e != ">" { body0 = IndexOnTheWay(lower, ">", body0) + 1 }
		xhtml = xhtml[body0:]
		lower = strings.ToLower(xhtml)

		// Remove string from <style> to </style>
		p0 := strings.Index(lower, "<style");  if p0 < 0 { break }
		p1 := strings.Index(lower, "</style"); if p1 < 0 { break }
		xhtml = xhtml[:p0] + xhtml[p1 + 8:]
	}

	for strings.Contains(xhtml, "<") || strings.Contains(xhtml, ">") {
		// Find "<" from HTML element and remove string until ">"
		p0 := strings.Index(xhtml, "<");         if p0 < 0 { break }
		p1 := IndexOnTheWay(xhtml, ">", p0 + 2); if p1 < 0 { break }

		if p0 >  0 { plain += xhtml[0:p0] + " "      }
		if p0 > p1 { plain += xhtml[p1 + 1:p0] + " " }

		xhtml = xhtml[p1 + 1:]
	}

	// Remove or replace entity references
	plain = strings.ReplaceAll(plain, "&lt;",   "<")
	plain = strings.ReplaceAll(plain, "&gt;",   ">")
	plain = strings.ReplaceAll(plain, "&quot;", `"`)
	plain = strings.ReplaceAll(plain, "&nbsp;", " ")
	plain = strings.ReplaceAll(plain, "&copy;", "(C)")
	plain = strings.ReplaceAll(plain, "&amp;",  "&")
	plain = Sweep(strings.ReplaceAll(plain, "\n", " "))
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

