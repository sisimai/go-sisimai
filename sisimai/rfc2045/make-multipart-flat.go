// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045
import "fmt"
import "strings"
import sisimoji "sisimai/string"

// haircut() remove unnecessary header fields except Content-Type, Content-Transfer-Encoding from
// multipart/* block.
func haircut(block *string, heads bool) []string {
	// @param    *string  block  multipart/* block text
	// @param    bool     heads  true: Returns only Content-(Type|Transfer-Encoding) headers
	// @return   []string        Two headers and body part of multipart/* block
	textchunks := strings.SplitN(*block, "\n\n", 2)
	upperchunk := textchunks[0]
	lowerchunk := textchunks[1]

	if len(upperchunk) == 0 || strings.Contains(upperchunk, "Content-Type:") == false {
		// There is neither "Content-Type:" nor "Content-Transfer-Encoding:" header
		return []string { "", "" }
	}

	var headerpart[2] string = [2]string{"", ""} // {"text/plain; charset=iso-2022-jp; ...", "quoted-printable"}
	for _, e := range strings.Split(upperchunk, "\n") {
		// Remove fields except Content-Type:, and Content-Transfer-Encoding: in each part of multipart/*
		// block such as the following:
		//   Date: Thu, 29 Apr 2018 22:22:22 +0900
		//   MIME-Version: 1.0
		//   Message-ID: ...
		//   Content-Transfer-Encoding: quoted-printable
		//   Content-Type: text/plain; charset=us-ascii
		if strings.HasPrefix(e, "Content-Type:") {
			// Content-Type: ***
			v := strings.SplitN(e, " ", 2)
			if strings.Contains(v[1], "boundary=") {
				// Do not convert to lower-cased when the value of Content-Type include a boundary string
				headerpart[0] = v[1]

			} else {
				// The value of Content-Type does not include a boundary string
				headerpart[0] = strings.ToLower(v[1])
			}
		} else if strings.HasPrefix(e, "Content-Transfer-Encoding:") {
			// Content-Transfer-Encodig: ***
			v := strings.SplitN(e, " ", 2)
			headerpart[1] = strings.ToLower(v[1])

		} else if strings.Contains(e, "boundary=") || strings.Contains(e, "charset=") {
			// "Content-Type" field has boundary="..." or charset="utf-8"
			if len(headerpart[0]) > 0 {
				// Append parameters
				headerpart[0] += " " + e
				headerpart[0]  = sisimoji.Squeeze(headerpart[0], " ")
			}
		}
	}
	if heads { return headerpart[:] }

	mediatypev := strings.ToLower(headerpart[1])
	ctencoding := headerpart[1]
	multipart1 := [...]string{headerpart[0], headerpart[1], ""}

	for {
		// UPPER CHUNK: Make a body part at the 2nd element of multipart1
		multipart1[2] = fmt.Sprintf("Content-Type: %s\n", headerpart[0])

		// Do not append Content-Transfer-Encoding: header when the part is the original message:
		// Content-Type is message/rfc822 or text/rfc822-headers, or message/delivery-status, or
		// message/feedback-report
		if strings.Contains(mediatypev, "/rfc822")          { break }
		if strings.Contains(mediatypev, "/delivery-status") { break }
		if strings.Contains(mediatypev, "/feedback-report") { break }
		if len(ctencoding) == 0                             { break }

		multipart1[2] += fmt.Sprintf("Content-Transfer-Encoding: %s\n", ctencoding)
		break
	}

	for {
		// LOWER CHUNK: Append LF before the lower chunk into the 2nd element of multipart1
		if len(lowerchunk) == 0    { break }
		if lowerchunk[0:1] == "\n" { break }

		multipart1[2] += "\n"
		break
	}
	multipart1[2] += lowerchunk
	return multipart1[:]
}

// levelout() splits the second argument: multipart/* blocks by a boundary string in the first argument.
func levelout(argv0 string, argv1 *string) [][3]string {
	// @param    string      argv0  The value of Content-Type header
	// @param    *string     argv1  A pointer to multipart/* message blocks
	// @return   [][3]string        List of each part of multipart/*
	if len(argv0)  == 0 { return nil }
	if len(*argv1) == 0 { return nil }

	boundary01 := Boundary(argv0, 0); if len(boundary01) == 0 { return nil }
	multiparts := strings.Split(*argv1, boundary01 + "\n")
	partstable := [][3]string{}

	// Remove empty or useless preamble and epilogue of multipart/* block
	if len(multiparts[0])                   < 8 { multiparts = multiparts[1:] }
	if len(multiparts[len(multiparts) - 1]) < 8 { multiparts = multiparts[0:len(multiparts) - 2] }

	for _, e := range multiparts {
		// Check each part and breaks up internal multipart/* block
		f := haircut(&e, false)
		if strings.Contains(f[0], "multipart/") {
			// There is nested multipart/* block
			boundary02 := Boundary(f[0], -1); if len(boundary02) == 0 { continue }
			bodyinside := strings.SplitN(f[2], "\n\n", 2)[1]

			if len(bodyinside) < 8                               { continue }
			if strings.Contains(bodyinside, boundary02) == false { continue }
			v := levelout(f[0], &bodyinside); if v == nil        { continue }

			for _, w := range v {
				partstable = append(partstable, [3]string{w[0], w[1], w[2]})
			}
		} else {
			// The part is not a multipart/* block
			b := e; if len(f[len(f) - 1]) > 0 { b = f[len(f) - 1] }
			v := [3]string{f[0], f[1], b}
			if len(f[0]) > 0 { v[2] = strings.SplitN(b, "\n\n", 2)[1] }
			partstable = append(partstable, v)
		}
	}
	if len(partstable) == 0 { return nil }

	// Remove `boundary01 + '--'` and strings from the boundary to the end of the body part.
	boundary01 = strings.Replace(boundary01, "\n", "", -1)
	b := partstable[len(partstable) - 1][2]
	p := strings.Index(b, boundary01 + "--")
	if p > -1 { partstable[len(partstable) - 1][2] = strings.SplitN(b, boundary01 + "--", 2)[0] }

	return partstable
}

// Makeflat() makes multipart/* part blocks flat and decode each part.
func MakeFlat(argv0 string, argv1 *string) *string {
	// @param    string  argv0  The value of Content-Type header
	// @param    *string argv1  A pointer to multipart/* message blocks
	// @return   *string        Message body
	if strings.Index(argv0, "multipart/") == -1 { return nil }
	if strings.Index(argv0, "boundary=")  == -1 { return nil }

	// Some bounce messages include lower-cased "content-type:" field such as the followings:
	//   - content-type: message/delivery-status        => Content-Type: message/delivery-status
	//   - content-transfer-encoding: quoted-printable  => Content-Transfer-Encoding: quoted-printable
	//   - CHARSET=, BOUNDARY=                          => charset-, boundary=
	//   - message/xdelivery-status                     => message/delivery-status
	for _, e := range [...]string{ "CONTENT-TYPE", "Content-type", "content-type" } {
		// Transform the Content-Type header name to the camel cased
		// TODO: These fields have been transformed by sisimai.message.Tidy() function ...?
		*argv1 = strings.Replace(*argv1, e + ":", "Content-Type:", -1)
	}

	for _, e := range [...]string { "CONTENT-TRANSFER-ENCODING", "content-transfer-encoding" } {
		// Transform the Content-Transfer-Encoding header name to the camel cased
		// TODO: These fields have been transformed by sisimai.message.Tidy() function ...?
		*argv1 = strings.Replace(*argv1, e + ":", "Content-Transfer-Encoding:", -1)
	}

	for _, e := range [...]string { "CHARSET", "CharSet", "Charset", "BOUNDARY", "Boundary" } {
		// Transform each parameter field name to the lower cased
		// TODO: These parameters have been transformed by sisimai.message.Tidy() function ...?
		*argv1 = strings.Replace(*argv1, e + "=", strings.ToLower(e) + "=", -1)
	}
	*argv1 = strings.Replace(*argv1, "message/xdelivery-status", "message/delivery-status", -1)

	multiparts := levelout(argv0, argv1)
	flattenout := ""

	for _, e := range multiparts {
		// Pick only the following parts Sisimai::Lhost will use, and decode each part
		// - text/plain, text/rfc822-headers
		// - message/delivery-status, message/rfc822, message/partial, message/feedback-report
		istexthtml := false
		mediatypev := Parameter(e[0], ""); if len(e[0]) == 0 { mediatypev = "text/plain" }

		// The value of Content-Type: is neither "text/*" nor "message/*"
		if strings.Index(mediatypev, "text/") + strings.Index(mediatypev, "message/") == -2 { continue }
		if mediatypev == "text/html" {
			// Skip text/html part when the value of Content-Type: header in an internal part of
			// multipart/* includes multipart/alternative;	
			if strings.Index(argv0, "multipart/alternative") > -1 { continue }
			istexthtml = true
		}
		ctencoding := e[1] // The value of Content-Transfer-Encoding header
		bodyinside := e[2] // Message body of the part
		bodystring := ""

		if len(ctencoding) > 0 {
			// Check the value of Content-Transfer-Encoding: header
			if ctencoding == "base64" {
				// Content-Transfer-Encoding: base64
				bodystring = *DecodeB(&bodyinside, "")

			} else if ctencoding == "quoted-printable" {
				// Content-Transfer-Encoding: quoted-printable
				bodystring = *DecodeQ(&bodyinside)

			} else if ctencoding == "7bit" {
				// Content-Transfer-Encoding: 7bit
				ctxcharset := Parameter(e[0], "charset")
				if strings.Index(ctxcharset, "iso-2022-") > -1 && len(ctxcharset) == 11 {
					// Content-Type: text/plain; charset=ISO-2022-JP
					//
					// TODO: Convert the string to UTF-8
					//       $bodystring = ${ Sisimai::String->to_utf8(\$bodyinside, $1) };
					bodystring = bodyinside

				} else {
					// No "charset" parameter in the value of Content-Type: header
					bodystring = bodyinside
				}
			} else {
				// Content-Transfer-Encoding: 8bit, binary, and so on
				bodystring = bodyinside
			}

			// Try to delete HTML tags inside of text/html part whenever possible
			if istexthtml { bodystring = *sisimoji.ToPlain(&bodystring, false) }
			if len(bodystring) == 0 { continue }

			// The new-line code in the converted string is CRLF
			if strings.Index(bodystring, "\r\n") > -1 { bodystring = *sisimoji.ToLF(&bodystring) }

		} else {
			// There is no Content-Transfer-Encoding header in the part 
			if strings.HasSuffix(mediatypev, "/delivery-status") ||
			   strings.HasSuffix(mediatypev, "/feedback-report") ||
			   strings.HasSuffix(mediatypev, "/rfc822") {

				// Add Content-Type: header of each part (will be used as a delimiter at Sisimai::Lhost)
				// into the body inside when the value of Content-Type: is message/delivery-status, or
				// message/rfc822, or text/rfc822-headers
				bodystring += fmt.Sprintf("Content-Type: %s\n", mediatypev)
			}
			bodystring += bodyinside
		}

		// Append "\n" when the last character of $bodystring is not LF
		if bodystring[len(bodystring) - 2:len(bodystring)] != "\n\n" { bodystring += "\n\n" }
		flattenout += bodystring
	}
	return &flattenout
}

