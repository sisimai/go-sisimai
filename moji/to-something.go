// Copyright (C) 2020,2024-2026 azumakuniyuki and sisimai development team, All
// rights reserved. This software is distributed under The BSD 2-Clause License.
//                  _ _ 
//  _ __ ___   ___ (_|_)
// | '_ ` _ \ / _ \| | |
// | | | | | | (_) | | |
// |_| |_| |_|\___// |_|
//               |__/   

package moji
import "bytes"

// ToLF replace CR and CR/LF with LF.
//   Arguments:
//     - text ([]byte): Text including CR or CR/LF.
//   Returns:
//     - (*string): Text converted to UTF8.
func ToLF(text []byte) []byte {
	if len(text) < 1 || bytes.IndexByte(text, '\r') < 0 { return text }

	bytelength := len(text)
	tolinefeed := make([]byte, 0, bytelength)

	for j := 0; j < bytelength; j++ {
		// Replace '\r' and '\r\n' with '\n'
		if text[j] != '\r' { tolinefeed = append(tolinefeed, text[j]); continue }
		if j + 1 < bytelength && text[j + 1] == '\n' {
			// The next character is not the last character, and the next character is '\n'
			tolinefeed = append(tolinefeed, '\n'); j++

		} else {
			// The next character is the last character, or the next character is not '\n'
			tolinefeed = append(tolinefeed, '\n')
		}
	}
	return tolinefeed
}

// ToPlain converts given HTML text to a plain text.
//   Arguments:
//     - htmle ([]byte): Text including HTML elements.
//   Returns:
//     - ([]byte): Converted plain text.
func ToPlain(htmle []byte) []byte {
	if len(htmle) == 0 { return htmle }

	// Find the position of "<body?", and remove the HTML header part
	lower := bytes.ToLower(htmle)
	body0 := bytes.Index(lower, []byte("<body")); if body0 < 0 { return htmle }
	body1 := bytes.IndexByte(lower[body0:], '>'); if body1 < 1 { return htmle }
	htmle  = htmle[body0 + body1 + 1:]

	for {
		// Remove string from <style> to </style>
		lower = bytes.ToLower(htmle)
		p0   := bytes.Index(lower, []byte("<style"));  if p0 < 0 { break }
		p1   := bytes.Index(lower, []byte("</style")); if p1 < 0 { break }

		if p1 < p0 { break }
		buffr := make([]byte, p0 + len(htmle) - (p1 + 8))
		copy(buffr, htmle[:p0])
		copy(buffr[p0:], htmle[p1 + 8:])
		htmle  = buffr
	}

	var buffr bytes.Buffer; buffr.Grow(len(htmle))
	xhtml := htmle; for {
		// Find "<" from HTML element and remove string until ">"
		p0 := bytes.IndexByte(xhtml, '<'); if p0 < 0 { buffr.Write(xhtml); break }
		buffr.Write(xhtml[:p0]); buffr.WriteByte(' ')

		p1 := bytes.IndexByte(xhtml[p0:], '>'); if p1 < 0 { break }
		xhtml = xhtml[p0 + p1 + 1:]
	}

	table := map[string]string{"lt": "<", "gt": ">", "quot": `"`, "nbsp": " ", "copy": "(C)", "amp": "&"}
	plain := buffr.Bytes()
	for k, e := range table {
		// Remove or replace entity references
		deref := bytes.Join([][]byte{[]byte("&"), []byte(k), []byte(";")}, nil) 
		plain  = bytes.ReplaceAll(plain, deref, []byte(e))
	}

	plain = bytes.ReplaceAll(plain, []byte("\n"), []byte(" "))
    return bytes.Join(bytes.Fields(plain), []byte(" "))
}

