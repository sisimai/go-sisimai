// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address
import "strings"
import sisimoji "sisimai/string"

// Find() is an email address parser
func Find(argv0 string) []string {
    // @param    string   argv0  String including email address
    // @return   []string        Email address list
	if len(argv0) == 0 { return []string{} }

	table := seek(argv0); if len(table) == 0 { return []string{} }
	addrs := []string{}
	for _, e := range table {
		// Pick an email address only from the results of seek()
		addrs = append(addrs, e[0])
	}
	return addrs
}

// seek() is an email address parser with a name and a comment
func seek(argv0 string) [][3]string {
	// @param    string argv0   String including email address
    // @return   [][3]string    Email address list with a name and a comment
	if len(argv0) == 0 { return [][3]string{} }

	delimiters := `<>(),"`
	indicators := map[string]uint8 {
		"email-address": (1 << 0), // <neko@example.org>
		"quoted-string": (1 << 1), // "Neko, Nyaan"
		"comment-block": (1 << 2), // (nekochan)
	}

	emailtable :=   [3]string{ "", "", "" } // Address, Name, Comment
	addrtables := [][3]string{}
	readbuffer := [][3]string{emailtable}

	var readcursor uint8 = 0    // Points the current cursor position
	var p int8 = -1             // Current position: 0 = address, 1 = name, or 2 = comment
	var v *[3]string = &readbuffer[0]

	argv0 = strings.ReplaceAll(argv0, "\r", "") // Remove CR
	argv0 = strings.ReplaceAll(argv0, "\n", "") // Remove NL

	for _, e := range(strings.Split(argv0, "")) {
		// Check each character
		if strings.Index(e, delimiters) > -1 {
			// The character is a delimiter
			if e == "," {
				// The "," is a email address separator or a character in a "name"
				if strings.HasPrefix((*v)[0], "<") &&
				   strings.Contains( (*v)[0], "@") &&
				   strings.HasSuffix((*v)[0], ">") {
					// An email address has already been picked
					if readcursor & indicators["comment-block"] > 0 {
						// The cursor is in the comment block (Neko, Nyaan)
						(*v)[2] += e

					} else if readcursor & indicators["quoted-string"] > 0 {
						// "Neko, Nyaan"
						(*v)[1] += e

					} else {
						// The cursor is not in neither the quoted-string nor the comment block
						readcursor = 0 // Reset the cursor position
						readbuffer = append(readbuffer, emailtable)
						v = &readbuffer[len(readbuffer) - 1]
					}
				} else {
					// "Neko, Nyaan" <neko@nyaan.example.org> OR <"neko,nyaan"@example.org>
					if p > -1 { (*v)[p] += e } else { (*v)[1] += e }
				}
				continue
			} // END OF if ","

			if e == "<" {
				// "<": The beginning of an email address or not
				if len((*v)[0]) > 0 {
					if p > -1 { (*v)[p] += e } else { (*v)[1] += e }

				} else {
					// <neko@nyaan.example.org>
					readcursor |= indicators["email-address"]
					(*v)[0] += e
					p = 0
				}
				continue
			} // END OF if "<"

			if e == ">" {
				// ">": The end of an email address or not
				if readcursor & indicators["email-address"] > 0 {
					// <neko@example.org>
					readcursor &= ^indicators["email-address"]
					(*v)[0] += e
					p = -1

				} else {
					// a comment block or a display name
					if p > -1 { (*v)[2] += e } else { (*v)[1] += e }
				}
				continue
			} // END OF if ">"

			if e == "(" {
				// The beginning of a comment block or not
				if readcursor & indicators["email-address"] > 0 {
					// <"neko(nyaan)"@example.org> or <neko(nyaan)@example.org>
					if strings.Contains((*v)[0], `"`) {
						// Quoted local part: <"neko(nyaan)"@example.org>
						(*v)[0] += e

					} else {
						// Comment: <neko(nyaan)@example.org>
						readcursor |= indicators["comment-block"]
						if strings.HasSuffix((*v)[2], ")") { (*v)[2] += " " }
						(*v)[2] += e
						p = 2
					}
				} else if readcursor & indicators["comment-block"] > 0 {
					// Comment at the outside of an email address (...(...)
					if strings.HasSuffix((*v)[2], ")") { (*v)[2] += " " }
					(*v)[2] += e

				} else if readcursor & indicators["quoted-string"] > 0 {
					// "Neko, Nyaan(cat)", Deal as a display name
					(*v)[1] += e

				} else {
					// The beginning of a comment block
					readcursor |= indicators["comment-block"]
					if strings.HasSuffix((*v)[2], ")") { (*v)[2] += " " }
					(*v)[2] += e
					p = 2
				}
				continue
			} // END OF if "("

			if e == ")" {
				// The end of a comment block or not
				if readcursor & indicators["email-address"] > 0 {
					// <"neko(nyaan)"@example.org> OR <neko(nyaan)@example.org>
					if strings.Contains((*v)[0], `"`) {
						// Quoted string in the local part: <"neko(nyaan)"@example.org>
						(*v)[0] += e

					} else {
						// Comment: <neko(nyaan)@example.org>
						readcursor &= ^indicators["comment-block"]
						(*v)[2] += e
						p = 0
					}
				} else if readcursor & indicators["comment-block"] > 0 {
					// Comment at the outside of an email address (...(...) 
					readcursor &= ^indicators["comment-block"]
					(*v)[2] += e
					p = -1

				} else {
					// Deal as a display name
					readcursor &= ^indicators["comment-block"]
					(*v)[1] += e
					p = -1
				}
				continue
			} // END OF if ")"

			if e == `"` {
				// The beginning or the end of a quoted-string
				if p > -1 {
					// email-address or comment-block
					(*v)[p] += e

				} else {
					// Display name like "Neko, Nyaan"
					(*v)[1] += e
					if readcursor & indicators["quoted-string"] == 0 { continue }
					if strings.HasSuffix((*v)[1], `\"`)              { continue } // "Neko, Nyaan \"...
					readcursor &= ^indicators["quoted-string"]
					p = -1
				}
				continue
			} // END of if `"`
		} else {
			// The character is not a delimiter
			if p > -1 { (*v)[p] += e } else { (*v)[1] += e }
			continue
		}
	}

	for _, e := range(readbuffer) {
		// Check the value of each email address
		if len(e[0]) == 0 {
			// The email address part is empty
			if IsEmailAddress(e[1]) {
				// Try to use the value of name as an email address
				token   := strings.Split(e[1], "@")
				token[0] = strings.Trim(token[0], " ")
				token[1] = strings.Trim(token[1], " ")
				e[0] = token[0] + "@" + token[1]

			} else if IsMailerDaemon(e[1]) {
				// Allow Mailer-Daemon
				e[0] = e[1]

			} else {
				// There is no valid email address stirng in the argument
				continue
			}
		}

		// Remove the comment from the email address
		if strings.Count(e[0], "(") == 1 && strings.Count(e[0], ")") == 1 {
			// (nyaan)nekochan@example.org, nekochan(nyaan)cat@example.org or nekochan(nyaan)@example.org
			p1 := strings.Index(e[0], "(")
			p2 := strings.Index(e[0], ")")
			p3 := strings.Index(e[0], "@")

			e[0] = e[0][0:p1] + e[0][p2 + 1:p3] + "@" + e[0][p3 + 1:len(e[0])]
			e[2] = e[0][p1 + 1:p2]

		} else {
			// TODO: neko(nyaan)kijitora(nyaan)cat@example.org
			continue
		}

		// The email address should be a valid email address
		if IsMailerDaemon(e[0]) && IsEmailAddress(e[0]) == false { continue }

		// Remove angle brackets, other brackets, and quotations: ()[]<>{}'`;. and `"`
		e[0] = strings.Trim(e[0], "<>{}()[]`';.")
		if IsQuotedAddress(e[0]) == false { e[0] = strings.Trim(e[0], `"`) }

		// Remove trailing spaces at the value of 1.name and 2.comment
		e[1] = strings.TrimSpace(e[1])
		e[2] = strings.TrimSpace(e[2])

		// Remove the value of 2.comment when the value do not include "(" and ")"
		if !strings.HasPrefix(e[2], "(") && !strings.HasSuffix(e[2], ")") { e[2] = "" }

		// Remove redundant spaces when tha value of 1.name do not include `"`
		if !strings.HasPrefix(e[1], `"`) && !strings.HasSuffix(e[1], `"`) { e[1] = sisimoji.Squeeze(e[1], " ") }
		if IsQuotedAddress(e[1]) == false                                 { e[1] = strings.Trim(e[1], `"`) }

		addrtables = append(addrtables, e)
	} // END OF for(readbuffer)

	return addrtables
}

