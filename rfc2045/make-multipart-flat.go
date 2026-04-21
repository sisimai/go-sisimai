// Copyright (C) 2020-2021,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 

package rfc2045
import "strings"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

// haircut remove unnecessary header fields except Content-Type, Content-Transfer-Encoding from multipart/* block.
//   Arguments:
//     - block (*string): multipart/* block text.
//     - heads (bool):    true if it returns only Content-(Type|Transfer-Encoding) headers.
//   Returns:
//     - ([]string): Two headers and body part of multipart/* block.
func haircut(block *string, heads bool) []string {
	// There is neither "Content-Type:" nor "Content-Transfer-Encoding:" header
	upperchunk, lowerchunk, exists := strings.Cut(*block, "\n\n"); if exists == false { return []string{"", ""} }
	if len(upperchunk) == 0 || strings.Contains(upperchunk, "Content-Type:") == false { return []string{"", ""} }

	headerpart := [2]string{} // {"text/plain; charset=iso-2022-jp; ...", "quoted-printable"}
	for e := range strings.Lines(upperchunk) {
		// Remove fields except Content-Type:, and Content-Transfer-Encoding: in each part of multipart/*
		// block such as the following:
		//   Date: Thu, 29 Apr 2018 22:22:22 +0900
		//   MIME-Version: 1.0
		//   Message-ID: ...
		//   Content-Transfer-Encoding: quoted-printable
		//   Content-Type: text/plain; charset=us-ascii
		if e = strings.TrimRight(e, "\n\r"); strings.HasPrefix(e, "Content-Type:") {
			// Content-Type: ***
			if _, rhs, cut := strings.Cut(e, " "); cut == true && strings.Contains(rhs, "boundary=") {
				// Do not convert to lower-cased when the value of Content-Type include a boundary string
				headerpart[0] = rhs

			} else {
				// The value of Content-Type does not include a boundary string
				headerpart[0] = strings.ToLower(rhs)
			}
		} else if strings.HasPrefix(e, "Content-Transfer-Encoding:") {
			// Content-Transfer-Encodig: ***
			if _, rhs, cut := strings.Cut(e, " "); cut == true { headerpart[1] = strings.ToLower(rhs) }

		} else if strings.Contains(e, "boundary=") || strings.Contains(e, "charset=") {
			// "Content-Type" field has boundary="..." or charset="utf-8"
			if len(headerpart[0]) > 0 {
				// Append parameters
				headerpart[0] += " " + e
				headerpart[0]  = strings.Join(strings.Fields(headerpart[0]), " ")
			}
		}
	}
	if heads { return headerpart[:] }

	mediatable := []string{"/rfc822", "/delivery-status", "/feedback-report"}
	mediatypev := strings.ToLower(headerpart[1])
	ctencoding := headerpart[1]
	multipart1 := [3]string{headerpart[0], headerpart[1], ""}
	multipart2 := strings.Builder{}; multipart2.Grow(len(lowerchunk) / 10 * 15)

	// UPPER CHUNK: Make a body part at the 2nd element of multipart1
	multipart2.WriteString("Content-Type: ")
	multipart2.WriteString(headerpart[0])
	multipart2.WriteByte('\n')

	if moji.ContainsAny(mediatypev, mediatable) == false && ctencoding != "" {
		// Do not append Content-Transfer-Encoding: header when the part is the original message:
		// Content-Type is message/rfc822 or text/rfc822-headers, or message/delivery-status, or
		// message/feedback-report
		multipart2.WriteString("Content-Transfer-Encoding: ")
		multipart2.WriteString(ctencoding)
		multipart2.WriteByte('\n')
	}

	// LOWER CHUNK: Append LF before the lower chunk into the 2nd element of multipart1
	if lowerchunk != "" && lowerchunk[0:1] != "\n" { multipart2.WriteRune('\n') }

	multipart2.WriteString(lowerchunk); multipart1[2] = multipart2.String()
	return multipart1[:]
}

// levelout splits the second argument: multipart/* blocks by a boundary string in the first argument.
//   Arguments:
//     - ctype (string):  value of Content-Type header.
//     - mpart (*string): Pointer to multipart/* message blocks.
//   Returns:
//     - ([][3]string):       List of each part of multipart/*.
//     - ([]siba.NotDecoded): Pointer to an occurred error list.
func levelout(ctype string, mpart *string) ([][3]string, []siba.NotDecoded) {
	if ctype == "" || mpart == nil || *mpart == ""        { return nil, nil }
	boundary01 := Boundary(ctype, 0); if boundary01 == "" { return nil, nil }
	multiparts := strings.Split(*mpart, boundary01 + "\n")
	partstable := make([][3]string, 0, 4)
	notdecoded := make([]siba.NotDecoded, 0)

	// Remove empty or useless preamble and epilogue of multipart/* block
	if len(multiparts[0]) < 8 { multiparts = multiparts[1:] }
	switch cw := len(multiparts); cw {
		case 0:  return nil, nil // There is no valid multipart block
		default: if cw > 2 && len(multiparts[cw - 1]) < 8 { multiparts = multiparts[0:cw - 2] }
	}

	for j, e := range multiparts {
		// Check each part and breaks up internal multipart/* block
		if j > 0 && strings.HasPrefix(e, "Content-") == false {
			// Add "Content-Type: text/plain" field at the head of the part because there is no
			// Content-Type: field; see set-of-emails/maildir/bsd/lhost-x1-01.eml
			e = "Content-Type: text/plain\n\n" + e
		}

		if cf := haircut(&e, false); strings.Contains(cf[0], "multipart/") {
			// There is nested multipart/* block
			boundary02 := Boundary(cf[0], -1);  if len(boundary02) == 0 { continue }
			_, bi, cut := strings.Cut(cf[2], "\n\n");   if cut == false { continue }
			if len(bi) < 8 || strings.Contains(bi, boundary02) == false { continue }

			cv, ce := levelout(cf[0], &bi); if ce != nil && len(ce) > 0 {
				// There is any errors
				notdecoded = append(notdecoded, ce...)
				if cv == nil { continue }
			}
			for _, w := range cv { partstable = append(partstable, [3]string{w[0], w[1], w[2]}) }

		} else {
			// The part is not a multipart/* block
			cw := len(cf)
			ub := e; if len(cf[cw - 1]) > 0 { ub = cf[cw - 1] }
			cv := [3]string{cf[0], cf[1], ub}; for len(cf[0]) > 0 {
				if cf[0] == "" || ub == "" || strings.Contains(ub, "\n\n") == false { break }
				_, cv[2], _ = strings.Cut(ub, "\n\n")
				break
			}
			partstable = append(partstable, cv)
		}
	}
	if len(partstable) == 0 { return nil, notdecoded }

	// Remove `boundary01 + '--'` and strings from the boundary to the end of the body part.
	boundary01 = strings.ReplaceAll(boundary01, "\n", "")
	cw := len(partstable)
	if ls, _, cx := strings.Cut(partstable[cw - 1][2], boundary01 + "--"); cx { partstable[cw - 1][2] = ls }

	return partstable, notdecoded
}

// Makeflat makes multipart/* part blocks flat and decode each part.
//   Arguments:
//     - ctype (string):  Value of Content-Type header.
//     - mpart (*string): Pointer to multipart/* message blocks.
//   Returns:
//     - (*string):           Message body.
//     - ([]siba.NotDecoded): Occurred errors.
func MakeFlat(ctype string, mpart *string) (*string, []siba.NotDecoded) {
	lhead := strings.ToLower(ctype)
	if moji.ContainsAny(lhead, []string{"multipart/", "boundary="}) == false { return nil, nil }

	multiparts, notdecoded := levelout(ctype, mpart)
	flatbuffer := strings.Builder{}; flatbuffer.Grow(len(*mpart) / 2)
	delimiters := []string{"/delivery-status", "/rfc822", "/feedback-report", "/partial"}

	for _, e := range multiparts {
		// Pick only the following parts Sisimai::Lhost will use, and decode each part
		// - text/plain, text/rfc822-headers
		// - message/delivery-status, message/rfc822, message/partial, message/feedback-report
		istexthtml := false
		mediatypev := Parameter(e[0], ""); if len(e[0]) == 0 { mediatypev = "text/plain" }

		// The value of Content-Type: is neither "text/*" nor "message/*"
		if moji.ContainsAny(mediatypev, []string{"text/", "message/"}) == false { continue }
		if mediatypev == "text/html" {
			// Skip text/html part when the value of Content-Type: header in an internal part of
			// multipart/* includes multipart/alternative;
			if strings.Contains(lhead, "multipart/alternative") == false { istexthtml = true }
		}
		bodyinside, bodystring := e[2], "" // Message body of the part, keeps decoded MIME part.

		if ctencoding := e[1]; ctencoding != "" {
			// Check the value of Content-Transfer-Encoding: header.
			var nyaan error; switch ctencoding {
				// - Content-Transfer-Encoding: 8bit, binary, and so on.
				// - sisimai no longer supports multibyte characters except UTF-8.
				// - https://github.com/sisimai/go-sisimai/issues/42
				default:                 bodystring        = bodyinside
				case "base64":           bodystring, nyaan = DecodeB(bodyinside)
				case "quoted-printable": bodystring, nyaan = DecodeQ(bodyinside)
			}
			if nyaan != nil { notdecoded = append(notdecoded, *siba.MakeNotDecoded(nyaan.Error(), false)) }

			switch {
				// Don't pick the decoded part as an error message when the part is
				// - BASE64 encoded.
				// - the value of the charset is not utf-8.
				// - NOT a plain text.
				case     ctencoding != "base64":
				case     moji.Aligned(e[0], []string{"charset", "=", "utf-8"}):
				case     moji.IsText(&bodystring):
				default: continue
			}

			// Try to delete HTML tags inside of text/html part whenever possible.
			if istexthtml { bodystring = *moji.ToPlain(&bodystring) }
			if bodystring == "" { continue }

			// The new-line code in the converted string is CRLF.
			moji.ToLF(&bodystring)

		} else {
			// There is no Content-Transfer-Encoding header in the part.
			bodystring += bodyinside
		}

		// There is no Content-Transfer-Encoding header in the part.
		if moji.ContainsAny(mediatypev, delimiters) {
			// Add Content-Type: header of each part (will be used as a delimiter at Sisimai::Lhost)
			// into the body inside when the value of Content-Type: is message/delivery-status, or
			// message/rfc822, or text/rfc822-headers.
			bodystring = "Content-Type: " + mediatypev + "\n" + bodystring
		}

		if cw := len(bodystring); cw > 1 && bodystring[cw - 2:] != "\n\n" {
			// Append "\n" when the last character of "bodystring" is not LF.
			bodystring += "\n\n"
		}
		flatbuffer.WriteString(bodystring)
	}
	flattenout := flatbuffer.String()
	return &flattenout, notdecoded
}

