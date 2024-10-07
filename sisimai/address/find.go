// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/
import "strings"
import sisimoji "sisimai/string"

// S3S4() runs like ruleset 3,4 of sendmail.cf
func S3S4(argv1 string) string {
	// @param    string argv1  Text including an email address
	// @return   string        Email address without comment, brackets
	if len(argv1)   == 0 { return "" }; list := Find(argv1)
	if len(list[0]) == 0 { return "" }
	return list[0]
}

// Find() is an email address parser with a name and a comment
func Find(argv1 string) [3]string {
	// @param    string argv1  String including email address
	// @return   [3]string{}          Email address table: {"address", "name", "comment"}
	if len(argv1) < 5 { return [3]string{} }
		argv1 = strings.ReplaceAll(argv1, "\r", "") // Remove CR
		argv1 = strings.ReplaceAll(argv1, "\n", "") // Remove NL
	if len(argv1) < 5 { return [3]string{} }

	delimiters := `<>(),"`
	groupindex := uint8(0)              // Group index: 0=undefined, 1=address, 2=name, 3=comment
	readcursor := uint8(0)              // Points the current cursor position
	readbuffer := [3]string{"", "", ""} // Read buffer for emailtable
	emailtable := [3]string{"", "", ""} // [0]Address, [1]Name, [2]Comment
	indicators := map[string]uint8{
		"email-address": (1 << 0), // <neko@example.org>
		"quoted-string": (1 << 1), // "Neko, Nyaan"
		"comment-block": (1 << 2), // (nekochan)
	}

	for _, e := range strings.Split(argv1, "") {
		// Check each character
		if strings.Contains(delimiters, e) {
			// The character is a delimiter
			if e == "," {
				// The "," is a email address separator or a character in a "name"
				if IsIncluded(readbuffer[0]) {
					// The email address has already been picked
					if readcursor & indicators["comment-block"] > 0 {
						// The cursor is in the comment block (Neko, Nyaan)
						readbuffer[2] += e

					} else if readcursor & indicators["quoted-string"] > 0 {
						// "Neko, Nyaan"
						readbuffer[1] += e

					} else {
						// The cursor is not in neither the quoted-string nor the comment block
						readcursor = 0 // Reset the current position
						groupindex = 0
					}
				} else {
					// "," is in the display name or the quoted local part of the email address
					// "Neko, Nyaan" <neko@nyaan.example.org> OR <"neko,nyaan"@example.org>
					if groupindex == 0 || groupindex == 2 {
						// Deal as a character of the display name
						readbuffer[1] += e

					} else {
						// Append "e" to "address" readbuffer[0] or "comment" readbuffer[2]
						readbuffer[groupindex - 1] += e
					}
				}
				continue
			}   // End of if(",")

			if e == "<" {
				// "<": The beginning of an email address or a character in the display name or the comment
				if len(readbuffer[0]) == 0 {
					// The 1st character of the email address: <neko@nyaan.example.jp>
					readcursor |= indicators["email-address"]
					readbuffer[0] = e
					groupindex = 1

				} else {
					// Check that readbuffer[0] already has a valid email address or not
					if IsIncluded(readbuffer[0]) {
						// The value of readbuffer[0] is a valid email address
						// "e" is a part of the display name or the comment
						if IsComment(readbuffer[2]) { readbuffer[2] += e } else { readbuffer[1] += e }
					}
				}
				continue
			}   // End of if("<")

			if e == ">" {
				// ">": The end of an email address or a character in the display name or the comment
				if readcursor & indicators["email-address"] > 0 {
					// The email address in readbuffer[0] has been successfully constructed
					readcursor &= ^indicators["email-address"]
					readbuffer[0] += e
					groupindex = 0

				} else {
					// ">" is a part of the comment block or the display name
					if IsComment(readbuffer[2]) { readbuffer[2] += e } else { readbuffer[1] += e }
				}
				continue
			}   // End of if(">")

			if e == "(" {
				// "(": The beginning of a comment block or a character in the display name or the comment
				if readcursor & indicators["email-address"] > 0 {
					// An email address including a comment like the followings:
					// <"neko(nyaan)"@example.org> or <neko(nyaan)@example.org>
					if strings.Contains(readbuffer[0], `"`) {
						// Quoted local part in the email address like <"neko(nyaan)"@example.org>
						readbuffer[0] += e

					} else {
						// A comment in the email address like <neko(nyaan)@example.org>
						readcursor |= indicators["comment-block"]
						if strings.HasSuffix(readbuffer[2], ")") { readbuffer[2] += " " }
						readbuffer[2] += e
						groupindex = 2
					}
				} else if readcursor & indicators["comment-block"] > 0 {
					// Comment at the outside of an email address (...(...)
					if strings.HasSuffix(readbuffer[2], ")") { readbuffer[2] += " " }
					readbuffer[2] += e

				} else if readcursor & indicators["quoted-string"] > 0 {
					// "Neko, Nyaan(cat)", Deal as a display name
					readbuffer[1] += e

				} else {
					// The beginning of the comment block
					readcursor |= indicators["comment-block"]
					if strings.HasSuffix(readbuffer[2], ")") { readbuffer[2] += " " }
					readbuffer[2] += e
					groupindex = 2
				}
				continue
			}   // End of if("(")

			if e == ")" {
				// "(": The end of a comment block or a character in the display name or the comment
				if readcursor & indicators["email-address"] > 0 {
					// An email address including a comment like the followings:
					// <"neko(nyaan)"@example.org> or <neko(nyaan)@example.org>
					if strings.Contains(readbuffer[0], `"`) {
						// Quoted local part in the email address like <"neko(nyaan)"@example.org>
						readbuffer[0] += e

					} else {
						// A comment in the email address like <neko(nyaan)@example.org>
						readcursor &= ^indicators["comment-block"]
						readbuffer[2] += e
						groupindex = 1
					}
				} else if readcursor & indicators["comment-block"] > 0 {
					// Comment at the outside of an email address (...(...)
					readcursor &= ^indicators["comment-block"]
					readbuffer[2] += e
					groupindex = 0

				} else {
					// Deal as a display name
					readbuffer[1] += e
					groupindex = 0
				}
				continue
			}   // End of if(")")

			if e == `"` {
				// The beginning or the end of a quoted-string
				if groupindex > 0 {
					// A part of the email address or the comment block
					readbuffer[groupindex - 1] += e

				} else {
					// The display name lke "Neko, Nyaan"
					readbuffer[1] += e
					if readcursor & indicators["quoted-string"] == 0 { continue }
					if strings.HasSuffix(readbuffer[1], `\"`)        { continue } // "Neko, Nyaan \"...
					readcursor &= ^indicators["quoted-string"]
					groupindex = 0
				}
				continue
			}   // End of if(`"`)
		} else {
			// The character is not a delimiter
			if groupindex == 0 || groupindex == 2 {
				// Deal as a character of the display name
				readbuffer[1] += e

			} else {
				// Append "e" to "address" readbuffer[0] or "comment" readbuffer[2]
				readbuffer[groupindex - 1] += e
			}
			continue
		}
	} // End of the loop(for)

	if len(readbuffer[0]) == 0 {
		// There is no email address
		if IsEmailAddress(readbuffer[1]) {
			// The display name part is an email address like "neko@example.jp"
			// TODO: Implement this block in p5-sisimai, rb-sisimai
			readbuffer[0] = "<" + readbuffer[1] + ">"

		} else if IsIncluded(readbuffer[1]) {
			// Try to use the string like an email address in the display name
			for _, e := range strings.Split(readbuffer[1], " ") {
				// Find an email address
				if IsEmailAddress(e) == false { continue }
				readbuffer[0] = e; break
			}
		} else if IsMailerDaemon(readbuffer[1]) {
			// Allow if the string is MAILER-DAEMON
			readbuffer[0] = readbuffer[1]
		}
	}

	for sisimoji.Aligned(readbuffer[0], []string{"(", ")"}) {
		// Remove the comment block from the email address
		// - (nyaan)nekochan@example.org
		// - nekochan(nyaan)cat@example.org
		// - nekochan(nyaan)@example.org
		p1 := strings.Index(readbuffer[0], "(")
		p2 := strings.Index(readbuffer[0], ")")
		ce := readbuffer[0][p1:p2 + 1]

		readbuffer[0] = strings.Replace(readbuffer[0], ce, "", 1)
		if len(readbuffer[2]) == 0 { readbuffer[2] = ce } else { readbuffer[2] += " " + ce }
	}

	if IsIncluded(readbuffer[0]) || IsMailerDaemon(readbuffer[0]) {
		// The email address must not include any character except from 0x20 to 0x7e.

		// - Remove angle brackets, other brackets, and quotations: []<>{}'` except a domain part is
		//   an IP address like neko@[192.0.2.222]
		// - Remove angle brackets, other brackets, and quotations: ()[]<>{}'`;. and `"`
		readbuffer[0] = strings.Trim(readbuffer[0], "<>{}()[]`';.")
		if IsQuotedAddress(readbuffer[0]) == false { readbuffer[0] = strings.Trim(readbuffer[0], `"`) }
		emailtable[0] = readbuffer[0]
	}

	if len(readbuffer[1]) > 0 {
		// Remove trailing spaces at the display name and the comment block
		readbuffer[1] = strings.TrimSpace(readbuffer[1])

		for {
			// Remove redundant spaces from the display name when the value is not a "quoted-string"
			if strings.HasPrefix(readbuffer[1], `"`) == false { break }
			if strings.HasSuffix(readbuffer[1], `"`) == false { break }

			readbuffer[1] = sisimoji.Squeeze(readbuffer[1], " ")
			break
		}
		if IsQuotedAddress(readbuffer[1]) == false {
			// Trim `"` from the display name when the value is not like "neko-nyaan"@libsisimai.org
			readbuffer[1] = strings.Trim(readbuffer[1], `"`)
		}
		emailtable[1] = readbuffer[1]
	}

	for _, e := range readbuffer {
		// There is no email address, try to pick an email address from each element in readbuffer
		e = strings.Trim(e, "<>{}()[]`';.")
		if IsQuotedAddress(e) == false { e = strings.Trim(e, `"`) }
		if IsEmailAddress(e)  == true  { emailtable[0] = e; break }
		if emailtable[0] != "" { break }
	}

	// Check and tidy up the comment block
	if IsComment(readbuffer[2]) { emailtable[2] = strings.TrimSpace(readbuffer[2]) }

	return emailtable
}

