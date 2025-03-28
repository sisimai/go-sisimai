// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 

package rfc2045
import "fmt"
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"

// haircut remove unnecessary header fields except Content-Type, Content-Transfer-Encoding from multipart/* block.
//   Arguments:
//     - block (*string): multipart/* block text
//     - heads (bool):    true if it returns only Content-(Type|Transfer-Encoding) headers
//   Returns:
//     - ([]string):      Two headers and body part of multipart/* block
func haircut(block *string, heads bool) []string {
	textchunks := strings.SplitN(*block, "\n\n", 2); if len(textchunks) < 2 { return []string{"", ""} }
	upperchunk := textchunks[0]
	lowerchunk := textchunks[1]

	// There is neither "Content-Type:" nor "Content-Transfer-Encoding:" header
	if len(upperchunk) == 0 || strings.Contains(upperchunk, "Content-Type:") == false { return []string{"", ""} }

	var headerpart[2] string = [2]string{} // {"text/plain; charset=iso-2022-jp; ...", "quoted-printable"}
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
			if cv := strings.SplitN(e, " ", 2); strings.Contains(cv[1], "boundary=") {
				// Do not convert to lower-cased when the value of Content-Type include a boundary string
				headerpart[0] = cv[1]

			} else {
				// The value of Content-Type does not include a boundary string
				headerpart[0] = strings.ToLower(cv[1])
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
				moji.Squeeze(&headerpart[0], ' ')
			}
		}
	}
	if heads { return headerpart[:] }

	mediatable := []string{"/rfc822", "/delivery-status", "/feedback-report"}
	mediatypev := strings.ToLower(headerpart[1])
	ctencoding := headerpart[1]
	multipart1 := [3]string{headerpart[0], headerpart[1], ""}
	multipart2 := strings.Builder{}; multipart2.Grow(len(lowerchunk) / 10 * 15)

	for {
		// UPPER CHUNK: Make a body part at the 2nd element of multipart1
		multipart2.WriteString("Content-Type: " + headerpart[0] + "\n")

		// Do not append Content-Transfer-Encoding: header when the part is the original message:
		// Content-Type is message/rfc822 or text/rfc822-headers, or message/delivery-status, or
		// message/feedback-report
		if moji.ContainsAny(mediatypev, mediatable) || ctencoding == "" { break }
		multipart2.WriteString("Content-Transfer-Encoding: " + ctencoding + "\n")
		break
	}

	// LOWER CHUNK: Append LF before the lower chunk into the 2nd element of multipart1
	if lowerchunk != "" && lowerchunk[0:1] != "\n" { multipart2.WriteRune('\n') }

	multipart2.WriteString(lowerchunk); multipart1[2] = multipart2.String()
	return multipart1[:]
}

// levelout splits the second argument: multipart/* blocks by a boundary string in the first argument.
//   Arguments:
//     - arvg0 (string):      value of Content-Type header
//     - arvg1 (*string):     Pointer to multipart/* message blocks
//   Returns:
//     - ([][3]string):       List of each part of multipart/*
//     - (*[]sis.NotDecoded): Pointer to an occurred error list
func levelout(argv0 string, argv1 *string) ([][3]string, *[]sis.NotDecoded) {
	if argv0 == "" || argv1 == nil || *argv1 == ""        { return nil, nil }
	boundary01 := Boundary(argv0, 0); if boundary01 == "" { return nil, nil }
	multiparts := strings.Split(*argv1, boundary01 + "\n")
	partstable := make([][3]string, 0, 4)
	notdecoded := []sis.NotDecoded{}

	// Remove empty or useless preamble and epilogue of multipart/* block
	if len(multiparts[0])                   < 8 { multiparts = multiparts[1:] }
	if len(multiparts[len(multiparts) - 1]) < 8 { multiparts = multiparts[0:len(multiparts) - 2] }

	for j, e := range multiparts {
		// Check each part and breaks up internal multipart/* block
		if j > 0 && strings.HasPrefix(e, "Content-") == false {
			// Add "Content-Type: text/plain" field at the head of the part because there is no
			// Content-Type: field; see set-of-emails/maildir/bsd/lhost-x1-01.eml
			e = "Content-Type: text/plain\n\n" + e
		}

		if cf := haircut(&e, false); strings.Contains(cf[0], "multipart/") {
			// There is nested multipart/* block
			boundary02 := Boundary(cf[0], -1); if len(boundary02) == 0 { continue }
			bodyinside := strings.SplitN(cf[2], "\n\n", 2)[1]
			if len(bodyinside) < 8 || strings.Contains(bodyinside, boundary02) == false { continue }

			cv, ce := levelout(cf[0], &bodyinside)
			if ce != nil && len(*ce) > 0 {
				// There is any errors
				notdecoded = append(notdecoded, *ce...)
				if cv == nil { continue }
			}
			for _, w := range cv { partstable = append(partstable, [3]string{w[0], w[1], w[2]}) }

		} else {
			// The part is not a multipart/* block
			cw := len(cf)
			ub := e; if len(cf[cw - 1]) > 0 { ub = cf[cw - 1] }
			cv := [3]string{cf[0], cf[1], ub}; for len(cf[0]) > 0 {
				if cf[0] == "" || ub == "" || strings.Contains(ub, "\n\n") == false { break }
				cv[2] = strings.SplitN(ub, "\n\n", 2)[1]
				break
			}
			partstable = append(partstable, cv)
		}
	}
	if len(partstable) == 0 { return nil, &notdecoded }

	// Remove `boundary01 + '--'` and strings from the boundary to the end of the body part.
	boundary01 = strings.Replace(boundary01, "\n", "", -1)
	cw := len(partstable)
	bo := partstable[cw - 1][2]
	p1 := strings.Index(bo, boundary01 + "--")
	if p1 > -1 { partstable[cw - 1][2] = strings.SplitN(bo, boundary01 + "--", 2)[0] }

	return partstable, &notdecoded
}

// Makeflat makes multipart/* part blocks flat and decode each part.
//   Arguments:
//     - argv0 (string):      Value of Content-Type header
//     - argv1 (*string):     Pointer to multipart/* message blocks
//   Returns:
//     - (*string):           Message body
//     - (*[]sis.NotDecoded): Occurred errors
func MakeFlat(argv0 string, argv1 *string) (*string, *[]sis.NotDecoded) {
	lhead := strings.ToLower(argv0)
	if moji.ContainsAny(lhead, []string{"multipart/", "boundary="}) == false { return nil, nil }

	// Some bounce messages include lower-cased "content-type:" field such as the followings:
	//   - content-type: message/delivery-status        => Content-Type: message/delivery-status
	//   - content-transfer-encoding: quoted-printable  => Content-Transfer-Encoding: quoted-printable
	//   - CHARSET=, BOUNDARY=                          => charset-, boundary=
	//   - message/xdelivery-status                     => message/delivery-status
	for _, e := range []string{"CONTENT-TYPE", "Content-type", "content-type"} {
		// Transform the Content-Type header name to the camel cased
		// TODO: These fields have been transformed by sisimai.message.Tidy() function ...?
		*argv1 = strings.Replace(*argv1, e + ":", "Content-Type:", -1)
	}

	for _, e := range []string{"CONTENT-TRANSFER-ENCODING", "content-transfer-encoding"} {
		// Transform the Content-Transfer-Encoding header name to the camel cased
		// TODO: These fields have been transformed by sisimai.message.Tidy() function ...?
		*argv1 = strings.Replace(*argv1, e + ":", "Content-Transfer-Encoding:", -1)
	}

	for _, e := range []string{"CHARSET", "CharSet", "Charset", "BOUNDARY", "Boundary"} {
		// Transform each parameter field name to the lower cased
		// TODO: These parameters have been transformed by sisimai.message.Tidy() function ...?
		*argv1 = strings.Replace(*argv1, e + "=", strings.ToLower(e) + "=", -1)
	}
	*argv1 = strings.Replace(*argv1, "message/xdelivery-status", "message/delivery-status", -1)

	multiparts, notdecoded := levelout(argv0, argv1)
	flatbuffer := strings.Builder{}; flatbuffer.Grow(len(*argv1) / 2)
	delimiters := []string{"/delivery-status", "/rfc822", "/feedback-report", "/partial"}

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
			if strings.Contains(lhead, "multipart/alternative") { continue }
			istexthtml = true
		}
		bodyinside := e[2] // Message body of the part
		bodystring := ""

		if ctencoding := e[1]; len(ctencoding) > 0 {
			// Check the value of Content-Transfer-Encoding: header
			if ctencoding == "base64" {
				// Content-Transfer-Encoding: base64
				cv, nyaan := DecodeB(bodyinside, ""); bodystring = cv
				if nyaan != nil {
					// Something wrong when the function decodes the BASE64 encoded string
					ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), false)
					*notdecoded = append(*notdecoded, ce)
				}
			} else if ctencoding == "quoted-printable" {
				// Content-Transfer-Encoding: quoted-printable
				cv, nyaan := DecodeQ(bodyinside); bodystring = cv
				if nyaan != nil {
					// Something wrong when the function decodes the Quoted-Printable encoded string
					ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), false)
					*notdecoded = append(*notdecoded, ce)
				}
			} else {
				// - Content-Transfer-Encoding: 8bit, binary, and so on
				// - sisimai no longer supports multibyte characters except UTF-8
				// - https://github.com/sisimai/go-sisimai/issues/42
				bodystring = bodyinside
			}

			// Try to delete HTML tags inside of text/html part whenever possible
			if istexthtml { bodystring = *moji.ToPlain(&bodystring) }
			if len(bodystring) == 0 { continue }

			// The new-line code in the converted string is CRLF
			moji.ToLF(&bodystring)

		} else {
			// There is no Content-Transfer-Encoding header in the part 
			bodystring += bodyinside
		}

		// There is no Content-Transfer-Encoding header in the part 
		if moji.ContainsAny(mediatypev, delimiters) {
			// Add Content-Type: header of each part (will be used as a delimiter at Sisimai::Lhost)
			// into the body inside when the value of Content-Type: is message/delivery-status, or
			// message/rfc822, or text/rfc822-headers
			bodystring = "Content-Type: " + mediatypev + "\n" + bodystring
		}

		// Append "\n" when the last character of $bodystring is not LF
		if bodystring[len(bodystring) - 2:len(bodystring)] != "\n\n" { bodystring += "\n\n" }
		flatbuffer.WriteString(bodystring)
	}
	flattenout := flatbuffer.String()
	return &flattenout, notdecoded
}

