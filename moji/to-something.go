// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All
// rights reserved. This software is distributed under The BSD 2-Clause License.
//                  _ _ 
//  _ __ ___   ___ (_|_)
// | '_ ` _ \ / _ \| | |
// | | | | | | (_) | | |
// |_| |_| |_|\___// |_|
//               |__/   

package moji
import "strings"

// ToLF replace CR and CR/LF with LF.
//   Arguments:
//     - argv0 (*string): Text including CR or CR/LF
//   Returns:
//     - (error):         Always nil
func ToLF(argv0 *string) *string {
	if argv0 == nil || *argv0 == "" || strings.IndexByte(*argv0, '\r') < 0 { return nil }

	readbuffer := []byte(*argv0)
	tolinefeed := make([]byte, 0, len(readbuffer))
	bytelength := len(readbuffer)

	for j := 0; j < bytelength; j++ {
		// Replace '\r' and '\r\n' with '\n'
		if readbuffer[j] != '\r' { tolinefeed = append(tolinefeed, readbuffer[j]); continue }
		if j + 1 < bytelength && readbuffer[j + 1] == '\n' {
			// The next character is not the last character, and the next character is '\n'
			tolinefeed = append(tolinefeed, '\n'); j++

		} else {
			// The next character is the last character, or the next character is not '\n'
			tolinefeed = append(tolinefeed, '\n')
		}
	}
	*argv0 = string(tolinefeed)
	return nil
}

// ToPlain converts given HTML text to a plain text.
//   Arguments:
//     - argv0 (*string): Text including HTML elements
//   Returns:
//     - (*string):       Converted plain text
func ToPlain(argv0 *string) *string {
	if argv0 == nil || *argv0 == "" { return argv0 }

	xhtml := *argv0
	lower := strings.ToLower(*argv0); if strings.Contains(lower, "<body") == false { return argv0 }
	plain := ""
	table := map[string]string{"lt": "<", "gt": ">", "quot": `"`, "nbsp": " ", "copy": "(C)", "amp": "&"}
	body0 := -1; for _, e := range []string{">", " ", "\t", "\n"} {
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

	for strings.IndexByte(xhtml, '<') > -1 || strings.IndexByte(xhtml, '>') > -1 {
		// Find "<" from HTML element and remove string until ">"
		p0 := strings.IndexByte(xhtml, '<');     if p0 < 0 { break }
		p1 := IndexOnTheWay(xhtml, ">", p0 + 2); if p1 < 0 { break }

		if p0 >  0 { plain += xhtml[0:p0] + " "      }
		if p0 > p1 { plain += xhtml[p1 + 1:p0] + " " }

		xhtml = xhtml[p1 + 1:]
	}

	// Remove or replace entity references
	for _, e := range table { plain = strings.ReplaceAll(plain, "&" + e + ";", table[e]) }
	plain = Sweep(strings.ReplaceAll(plain, "\n", " "))
	return &plain
}

