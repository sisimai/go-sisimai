// Copyright (C) 2020,2024-2026 azumakuniyuki and sisimai development team, All
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
//     - text (*string): Text including CR or CR/LF.
//   Returns:
//     - (*string): Text converted to UTF8.
func ToLF(text *string) *string {
	if text == nil || *text == "" || strings.IndexByte(*text, '\r') < 0 { return nil }

	readbuffer := []byte(*text)
	bytelength := len(readbuffer)
	tolinefeed := make([]byte, 0, len(readbuffer))

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
	*text = string(tolinefeed)
	return nil
}

// ToPlain converts given HTML text to a plain text.
//   Arguments:
//     - htmle (*string): Text including HTML elements.
//   Returns:
//     - (*string): Converted plain text.
func ToPlain(htmle *string) *string {
	if htmle == nil || *htmle == "" { return htmle }

	lower := strings.ToLower(*htmle); if strings.Contains(lower, "<body") == false { return htmle }
	xhtml := *htmle
	buffr := strings.Builder{}; buffr.Grow(len(xhtml) / 4)
	for _, e := range []string{">", " ", "\t", "\n"} {
		// Find the position of <body?, and remove the HTML header part
		body0 := strings.Index(lower, "<body" + e); if body0 < 0 { continue }
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

		if p0 >  0 { buffr.WriteString(xhtml[0:p0]); buffr.WriteByte(' ')      }
		if p0 > p1 { buffr.WriteString(xhtml[p1 + 1:p0]); buffr.WriteByte(' ') }

		xhtml = xhtml[p1 + 1:]
	}

	// Remove or replace entity references
	table := map[string]string{"lt": "<", "gt": ">", "quot": `"`, "nbsp": " ", "copy": "(C)", "amp": "&"}
	plain := ""
	for _, e := range table { plain = strings.ReplaceAll(buffr.String(), "&" + e + ";", table[e]) }
	plain = strings.Join(strings.Fields(strings.ReplaceAll(plain, "\n", " ")), " ")
	return &plain
}

