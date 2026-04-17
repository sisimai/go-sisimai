// Copyright (C) 2020-2021,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 

package rfc2045
import "bytes"
import "strings"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

// haircut remove unnecessary header fields except Content-Type, Content-Transfer-Encoding from multipart/* block.
//   Arguments:
//     - block ([]byte): multipart/* block text.
//     - heads (bool):   true if it returns only Content-(Type|Transfer-Encoding) headers.
//   Returns:
//     - ([][]byte): Two headers and body part of multipart/* block.
func haircut(block []byte, heads bool) [][]byte {
	// There is neither "Content-Type:" nor "Content-Transfer-Encoding:" header
	upperchunk, lowerchunk, exists := bytes.Cut(block, []byte("\n\n"));  if exists == false { return [][]byte{nil, nil} }
	if len(upperchunk) == 0 || bytes.Contains(upperchunk, []byte("Content-Type:")) == false { return [][]byte{nil, nil} }

	headerpart := [2][]byte{} // {"text/plain; charset=iso-2022-jp; ...", "quoted-printable"}
	for e := range bytes.Lines(upperchunk) {
		// Remove fields except Content-Type:, and Content-Transfer-Encoding: in each part of multipart/*
		// block such as the following:
		//   Date: Thu, 29 Apr 2018 22:22:22 +0900
		//   MIME-Version: 1.0
		//   Message-ID: ...
		//   Content-Transfer-Encoding: quoted-printable
		//   Content-Type: text/plain; charset=us-ascii
		if e = bytes.TrimRight(e, "\n\r"); bytes.HasPrefix(e, []byte("Content-Type:")) {
			// Content-Type: ***
			if _, rhs, cut := bytes.Cut(e, []byte(" ")); cut == true && bytes.Contains(rhs, []byte("boundary=")) {
				// Do not convert to lower-cased when the value of Content-Type include a boundary string
				headerpart[0] = rhs

			} else {
				// The value of Content-Type does not include a boundary string
				headerpart[0] = bytes.ToLower(rhs)
			}
		} else if bytes.HasPrefix(e, []byte("Content-Transfer-Encoding:")) {
			// Content-Transfer-Encodig: ***
			if _, rhs, cut := bytes.Cut(e, []byte(" ")); cut == true { headerpart[1] = bytes.ToLower(rhs) }

		} else if bytes.Contains(e, []byte("boundary=")) || bytes.Contains(e, []byte("charset=")) {
			// "Content-Type" field has boundary="..." or charset="utf-8"
			if len(headerpart[0]) > 0 {
				// Append parameters
				headerpart[0] = append(append(headerpart[0], ' '), e...)
				headerpart[0] = bytes.Join(bytes.Fields(headerpart[0]), []byte(" "))
			}
		}
	}
	if heads { return headerpart[:] }

	mediatable := [][]byte{[]byte("/rfc822"), []byte("/delivery-status"), []byte("/feedback-report")}
	mediatypev := bytes.ToLower(headerpart[1])
	ctencoding := headerpart[1]
	multipart1 := [3][]byte{headerpart[0], headerpart[1], []byte("")}
	multipart2 := bytes.Buffer{}; multipart2.Grow(len(lowerchunk) / 10 * 15)

	// UPPER CHUNK: Make a body part at the 2nd element of multipart1
	multipart2.Write([]byte("Content-Type: "))
	multipart2.Write(headerpart[0])
	multipart2.WriteByte('\n')

	if len(ctencoding) > 0 {
		// Do not append Content-Transfer-Encoding: header when the part is the original message:
		cf := false; for _, e := range mediatable {
			// Content-Type is message/rfc822 or text/rfc822-headers, or message/delivery-status,
			// or message/feedback-report
			if bytes.Contains(mediatypev, e) == true { cf = true; break }
		}
		if cf == false {
			// Append Content-Transfer-Encoding: field when the part does not include delivery
			// status information.
			multipart2.Write([]byte("Content-Transfer-Encoding: "))
			multipart2.Write(append(ctencoding, '\n'))
		}
	}

	// LOWER CHUNK: Append LF before the lower chunk into the 2nd element of multipart1
	if len(lowerchunk) > 0 && bytes.Equal(lowerchunk[0:1], []byte("\n")) == false { multipart2.WriteByte('\n') }

	multipart2.Write(lowerchunk); multipart1[2] = multipart2.Bytes()
	return multipart1[:]
}

// levelout splits the second argument: multipart/* blocks by a boundary string in the first argument.
//   Arguments:
//     - ctype (string):  value of Content-Type header.
//     - mpart ([]byte):  multipart/* message blocks.
//   Returns:
//     - ([][3][]byte):       List of each part of multipart/*.
//     - ([]siba.NotDecoded): Pointer to an occurred error list.
func levelout(ctype string, mpart []byte) ([][3][]byte, []siba.NotDecoded) {
	if ctype == "" || len(mpart) == 0                     { return nil, nil }
	boundary01 := Boundary(ctype, 0); if boundary01 == "" { return nil, nil }
	multiparts := bytes.Split(mpart, append([]byte(boundary01), '\n'))
	partstable := make([][3][]byte, 0, 4)
	notdecoded := make([]siba.NotDecoded, 0)

	// Remove empty or useless preamble and epilogue of multipart/* block
	if len(multiparts[0]) < 8 { multiparts = multiparts[1:] }
	switch cw := len(multiparts); cw {
		case 0:  return nil, nil // There is no valid multipart block
		default: if cw > 2 && len(multiparts[cw - 1]) < 8 { multiparts = multiparts[0:cw - 2] }
	}

	for j, e := range multiparts {
		// Check each part and breaks up internal multipart/* block
		if j > 0 && bytes.HasPrefix(e, []byte("Content-")) == false {
			// Add "Content-Type: text/plain" field at the head of the part because there is no
			// Content-Type: field; see set-of-emails/maildir/bsd/lhost-x1-01.eml
			e = append([]byte("Content-Type: text/plain\n\n"), e...)
		}

		if cf := haircut(e, false); bytes.Contains(cf[0], []byte("multipart/")) {
			// There is nested multipart/* block
			boundary02 := Boundary(string(cf[0]), -1);if len(boundary02) == 0 { continue }
			_, bi, cut := bytes.Cut(cf[2], []byte("\n\n"));   if cut == false { continue }
			if len(bi) < 8 || bytes.Contains(bi, []byte(boundary02)) == false { continue }

			cv, ce := levelout(string(cf[0]), bi); if ce != nil && len(ce) > 0 {
				// There is any errors
				notdecoded = append(notdecoded, ce...)
				if cv == nil { continue }
			}
			for _, w := range cv { partstable = append(partstable, [3][]byte{w[0], w[1], w[2]}) }

		} else {
			// The part is not a multipart/* block
			cw := len(cf)
			ub := e; if len(cf[cw - 1]) > 0 { ub = cf[cw - 1] }
			cv := [3][]byte{cf[0], cf[1], ub}; for len(cf[0]) > 0 {
				if len(cf[0]) < 1 || len(ub) < 1 || bytes.Contains(ub, []byte("\n\n")) == false { break }
				_, cv[2], _ = bytes.Cut(ub, []byte("\n\n"))
				break
			}
			partstable = append(partstable, cv)
		}
	}
	if len(partstable) == 0 { return nil, notdecoded }

	// Remove `boundary01 + '--'` and strings from the boundary to the end of the body part.
	boundary01 = strings.ReplaceAll(boundary01, "\n", "")
	cw := len(partstable)
	cb := []byte(boundary01 + "--")
	if ls, _, cx := bytes.Cut(partstable[cw - 1][2], cb); cx { partstable[cw - 1][2] = ls }

	return partstable, notdecoded
}

// Makeflat makes multipart/* part blocks flat and decode each part.
//   Arguments:
//     - ctype (string):  Value of Content-Type header.
//     - mpart ([]byte):  multipart/* message blocks.
//   Returns:
//     - ([]byte):            Message body.
//     - ([]siba.NotDecoded): Occurred errors.
func MakeFlat(ctype string, mpart []byte) ([]byte, []siba.NotDecoded) {
	lhead := strings.ToLower(ctype)
	if moji.ContainsAny(lhead, []string{"multipart/", "boundary="}) == false { return nil, nil }

	multiparts, notdecoded := levelout(ctype, mpart)
	flatbuffer := bytes.Buffer{}; flatbuffer.Grow(len(mpart) / 2)
	delimiters := []string{"/delivery-status", "/rfc822", "/feedback-report", "/partial"}

	for _, e := range multiparts {
		// Pick only the following parts Sisimai::Lhost will use, and decode each part
		// - text/plain, text/rfc822-headers
		// - message/delivery-status, message/rfc822, message/partial, message/feedback-report
		istexthtml := false
		mediatypev := Parameter(string(e[0]), ""); if len(e[0]) == 0 { mediatypev = "text/plain" }

		// The value of Content-Type: is neither "text/*" nor "message/*"
		if moji.ContainsAny(mediatypev, []string{"text/", "message/"}) == false { continue }
		if mediatypev == "text/html" {
			// Skip text/html part when the value of Content-Type: header in an internal part of
			// multipart/* includes multipart/alternative;
			if strings.Contains(lhead, "multipart/alternative") == false { istexthtml = true }
		}
		bodyinside, bodystring := e[2], []byte("") // Message body of the part, keeps decoded MIME part.

		if ctencoding := string(e[1]); ctencoding != "" {
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
				case     moji.Aligned(string(e[0]), []string{"charset", "=", "utf-8"}):
				case     moji.IsText(bodystring):
				default: continue
			}

			// Try to delete HTML tags inside of text/html part whenever possible.
			if istexthtml { bodystring = moji.ToPlain(bodystring) }
			if len(bodystring) == 0 { continue }

			// The new-line code in the converted string is CRLF.
			bodystring = moji.ToLF(bodystring)

		} else {
			// There is no Content-Transfer-Encoding header in the part.
			bodystring = append(bodystring, bodyinside...)
		}

		// There is no Content-Transfer-Encoding header in the part.
		if moji.ContainsAny(mediatypev, delimiters) {
			// Add Content-Type: header of each part (will be used as a delimiter at Sisimai::Lhost)
			// into the body inside when the value of Content-Type: is message/delivery-status, or
			// message/rfc822, or text/rfc822-headers.
			partheader := make([]byte, 0, len(bodystring) + 50)
			partheader  = append(partheader, []byte("Content-Type: ")...)
			partheader  = append(partheader, []byte(mediatypev)...)
			partheader  = append(partheader, '\n')
			partheader  = append(partheader, bodystring...)
			bodystring  = partheader
		}

		if cw := len(bodystring); cw > 1 && bytes.HasSuffix(bodystring, []byte("\n\n")) == false {
			// Append "\n" when the last character of "bodystring" is not LF.
			bodystring = append(bodystring, []byte("\n\n")...)
		}
		flatbuffer.Write(bodystring)
	}
	return flatbuffer.Bytes(), notdecoded
}

