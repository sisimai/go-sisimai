// Copyright (C) 2020-2021,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/

package address
import "strings"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc1123"
import "libsisimai.org/sisimai/v5/rfc5322"

const (
	HereIsEmailAddress = 1 << (iota + 1) // <neko@example.org>
	HereIsQuotedString                   // "Neko, Nyaan"
	HereIsCommentBlock                   // (nekochan)
)

// S3S4 runs like the ruleset 3 and 4 of /etc/sendmail.cf file.
//   Arguments:
//     - email (string): String including an email address.
//   Returns:
//     - (string): Email address picked from the given string.
func S3S4(email string) string {
	if len(email)   == 0 { return "" }; list := Find(email)
	if len(list[0]) == 0 { return "" }
	return list[0]
}

// Find is an email address parser with a name and comment.
//   Arguments:
//     - text (string): String including an email address.
//   Returns:
//     - ([3]string): Email address table such as `[3]string{"address", "name", "comment"}`.
func Find(text string) [3]string {
	if len(text) < 5 { return [3]string{} }

	delimiters := `<>(),"`
	groupindex := 0 // Group index: 0=undefined, 1=address, 2=name, 3=comment
	readcursor := 0 // Points the current cursor position
	readbuffer := [3]strings.Builder{}; readbuffer[0].Grow(32); readbuffer[1].Grow(32); readbuffer[2].Grow(8)

	for _, e := range text {
		// Check each character
		if strings.ContainsRune(delimiters, e) {
			// The character is a delimiter
			switch e {
			case ',': // The "," is a email address separator or a character in a "name"
				if IsIncluded(readbuffer[0].String()) {
					// The email address has already been picked
					if readcursor & HereIsCommentBlock > 0 {
						// The cursor is in the comment block (Neko, Nyaan)
						readbuffer[2].WriteRune(e)

					} else if readcursor & HereIsQuotedString > 0 {
						// "Neko, Nyaan"
						readbuffer[1].WriteRune(e)

					} else {
						// The cursor is not in neither the quoted-string nor the comment block.
						readcursor, groupindex = 0, 0
					}
				} else {
					// "," is in the display name or the quoted local part of the email address
					// "Neko, Nyaan" <neko@cat.example.org> OR <"neko,cat"@example.org>
					if groupindex == 0 || groupindex == 2 {
						// Deal as a character of the display name
						readbuffer[1].WriteRune(e)

					} else {
						// Append "e" to "address" readbuffer[0] or "comment" readbuffer[2]
						readbuffer[groupindex - 1].WriteRune(e)
					}
				}
			case '<': // The beginning of an email address or a character in the display name or the comment
				if readbuffer[0].Len() == 0 {
					// The 1st character of the email address: <neko@cat.example.jp>
					readcursor |= HereIsEmailAddress
					readbuffer[0].Reset(); readbuffer[0].WriteRune(e)
					groupindex = 1

				} else if IsIncluded(readbuffer[0].String()) {
					// Check that readbuffer[0] already has a valid email address or not
					// The value of readbuffer[0] is a valid email address
					if rfc5322.IsComment(readbuffer[2].String()) {
						// "e" is a part of the comment
						readbuffer[2].WriteRune(e)

					} else {
						// "e" is a part of the display name
						readbuffer[1].WriteRune(e)
					}
				}
			case '>': // The end of an email address or a character in the display name or the comment
				if readcursor & HereIsEmailAddress > 0 {
					// The email address in readbuffer[0] has been successfully constructed
					readcursor &= ^HereIsEmailAddress
					readbuffer[0].WriteRune(e)
					groupindex = 0

				} else {
					// ">" is a part of the comment block or the display name
					if rfc5322.IsComment(readbuffer[2].String()) {
						// "e" is a part of the comment
						readbuffer[2].WriteRune(e)

					} else {
						// "e" is a part of the display name
						readbuffer[1].WriteRune(e)
					}
				}
			case '(': // The beginning of a comment block or a character in the display name or the comment
				if readcursor & HereIsEmailAddress > 0 {
					// An email address including a comment like the followings:
					// <"neko(cat)"@example.org> or <neko(cat)@example.org>
					if strings.IndexByte(readbuffer[0].String(), '"') > -1 {
						// Quoted local part in the email address like <"neko(cat)"@example.org>
						readbuffer[0].WriteRune(e)

					} else {
						// A comment in the email address like <neko(cat)@example.org>
						readcursor |= HereIsCommentBlock
						if strings.HasSuffix(readbuffer[2].String(), ")") { readbuffer[2].WriteRune(' ') }
						readbuffer[2].WriteRune(e)
						groupindex = 3
					}
				} else if readcursor & HereIsCommentBlock > 0 {
					// Comment at the outside of an email address (...(...)
					if strings.HasSuffix(readbuffer[2].String(), ")") { readbuffer[2].WriteRune(' ') }
					readbuffer[2].WriteRune(e)

				} else if readcursor & HereIsQuotedString > 0 {
					// "Neko, Nyaan(cat)", Deal as a display name
					readbuffer[1].WriteRune(e)

				} else {
					// The beginning of the comment block
					readcursor |= HereIsCommentBlock
					if strings.HasSuffix(readbuffer[2].String(), ")") { readbuffer[2].WriteRune(' ') }
					readbuffer[2].WriteRune(e)
					groupindex = 3
				}
			case ')': // The end of a comment block or a character in the display name or the comment
				if readcursor & HereIsEmailAddress > 0 {
					// An email address including a comment like the followings:
					// <"neko(cat)"@example.org> or <neko(cat)@example.org>
					if strings.IndexByte(readbuffer[0].String(), '"') > -1 {
						// Quoted local part in the email address like <"neko(cat)"@example.org>
						readbuffer[0].WriteRune(e)

					} else {
						// A comment in the email address like <neko(cat)@example.org>
						readcursor &= ^HereIsCommentBlock
						readbuffer[2].WriteRune(e)
						groupindex = 1
					}
				} else if readcursor & HereIsCommentBlock > 0 {
					// Comment at the outside of an email address (...(...)
					readcursor &= ^HereIsCommentBlock
					readbuffer[2].WriteRune(e)
					groupindex = 0

				} else {
					// Deal as a display name
					readbuffer[1].WriteRune(e)
					groupindex = 0
				}
			case '"': // The beginning|end of the quoted string block or a part of an email address.
				if groupindex == 0 {
					// The beginning of the quoted-string block
					readbuffer[1].WriteRune(e)
					readcursor |= HereIsQuotedString
					groupindex  = 2

				} else if groupindex == 2 {
					// The end of the quoted-string block
					readcursor &= ^HereIsQuotedString
					groupindex  = 0

				} else if groupindex > 0 {
					// A part of the email address or the comment block
					readbuffer[groupindex - 1].WriteRune(e)

				} else {
					// The display name lke "Neko, Nyaan"
					readbuffer[1].WriteRune(e)
					if readcursor & HereIsQuotedString == 0            { continue }
					if strings.HasSuffix(readbuffer[1].String(), `\"`) { continue } // "Neko, Nyaan \"...
					readcursor &= ^HereIsQuotedString
					groupindex = 0
				} 
			}
		} else {
			// The character is not a delimiter
			switch groupindex {
				// - Deal as a character of the display name
				// - Append "e" to "address" readbuffer[0] or "comment" readbuffer[2]
				case 0, 2: readbuffer[1].WriteRune(e) 
				default:   readbuffer[groupindex - 1].WriteRune(e)
			}
		}
	} // End of the loop(for)

	layoutbuff := [3]string{readbuffer[0].String(), readbuffer[1].String(), readbuffer[2].String()}
	emailtable := [3]string{} // [0]Address, [1]Name, [2]Comment

	if len(layoutbuff[0]) == 0 {
		// There is no email address
		if rfc5322.IsEmailAddress(layoutbuff[1]) == true {
			// The display name part is an email address like "neko@example.jp"
			// TODO: Implement this block in p5-sisimai, rb-sisimai
			layoutbuff[0] = "<" + strings.TrimSpace(layoutbuff[1]) + ">"

		} else if IsIncluded(layoutbuff[1]) == true {
			// Try to use the string like an email address in the display name
			for _, e := range strings.Split(layoutbuff[1], " ") {
				// Find an email address
				if rfc5322.IsEmailAddress(e) { layoutbuff[0] = e; break }
			}
		} else if IsMailerDaemon(layoutbuff[1]) == true {
			// Allow if the string is MAILER-DAEMON
			layoutbuff[0] = strings.TrimSpace(layoutbuff[1])
		}
	}

	for moji.Aligned(layoutbuff[0], []string{"(", ")"}) {
		// Remove the comment block from the email address
		// - (cat)nekochan@example.org
		// - nekochan(cat)cat@example.org
		// - nekochan(cat)@example.org
		ce := "(" + moji.Select(layoutbuff[0], "(", ")", 0) + ")"
		layoutbuff[0] = strings.Replace(layoutbuff[0], ce, "", 1)
		if len(layoutbuff[2]) == 0 { layoutbuff[2] = ce } else { layoutbuff[2] += " " + ce }
	}

	if IsIncluded(layoutbuff[0]) || IsMailerDaemon(layoutbuff[0]) {
		// The email address must not include any character except from 0x20 to 0x7e.
		// - Remove angle brackets, other brackets, and quotations: []<>{}'` except a domain part is
		//   an IP address like neko@[192.0.2.222]
		// - Remove angle brackets, other brackets, and quotations: ()[]<>{}'`;. and `"`
		if rfc1123.IsDomainLiteral(layoutbuff[0]) == false { layoutbuff[0] = strings.Trim(layoutbuff[0], "[]{}()`';.") }
		                                                     layoutbuff[0] = Final(strings.Trim(layoutbuff[0], "<>"))
		if rfc5322.IsQuotedAddress(layoutbuff[0]) == false { layoutbuff[0] = strings.Trim(layoutbuff[0], `"`)          }
		emailtable[0] = layoutbuff[0]
	}

	if layoutbuff[1] != "" {
		// Remove trailing spaces at the display name and the comment block
		layoutbuff[1] = strings.TrimSpace(layoutbuff[1])

		if strings.HasPrefix(layoutbuff[1], `"`) == false || strings.HasSuffix(layoutbuff[1], `"`) == false {
			// Remove redundant spaces from the display name when the value is not a "quoted-string"
			moji.Squeeze(&layoutbuff[1], ' ')
		}
		if rfc5322.IsQuotedAddress(layoutbuff[1]) == false {
			// Trim `"` from the display name when the value is not like "neko-cat"@libsisimai.org
			layoutbuff[1] = strings.Trim(layoutbuff[1], `"`)
		}
		emailtable[1] = layoutbuff[1]
	}

	E0: for emailtable[0] == "" {
		// There is no email address in emailtable[0]
		for _, e := range []string{layoutbuff[1], layoutbuff[2]} {
			// Try to pick an email address from each element in layoutbuff
			for _, f := range strings.Split(e, " ") {
				// Find an email address like string from each element splitted by " "
				if f == "" || strings.IndexByte(f, '@') < 1      { continue }
				f = strings.Trim(f, "{}()[]`';."); if len(f) < 5 { continue }
				f = Final(f)

				if rfc5322.IsQuotedAddress(f) == false { e = strings.Trim(e, `"`)    }
				if rfc5322.IsEmailAddress(f)  == true  { emailtable[0] = f; break E0 }
			}
		}
		break E0
	}

	// Remove "." at the end of the email address such as "neko@example.jp."
	emailtable[0] = strings.TrimRight(emailtable[0], ".")

	// Check and tidy up the comment block
	if rfc5322.IsComment(layoutbuff[2]) { emailtable[2] = strings.TrimSpace(layoutbuff[2]) }

	return emailtable
}

