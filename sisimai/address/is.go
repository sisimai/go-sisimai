// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address
//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/
//                                      
import "strings"

// IsQuotedAddress() checks that the local part of the argument is quoted
func IsQuotedAddress(email string) bool {
	// @param    string email    Email address string
	// @return   bool            true:  the local part is quoted
	//                           false: the local part is not quoted
	if strings.HasPrefix(email, `"`) == false { return false }
	if strings.Contains(email, `"@`) == false { return false }
	return true
}

// IsIncluded() returns true if the string include an email address
func IsIncluded(argv0 string) bool {
	// @param    string argv0    String including an email address like "<neko@nyaan.jp>"
	// @return   bool            true:  is including an email address
	//                           false: is not including an email address
	if len(argv0) == 0                           { return false }
	if strings.HasPrefix(argv0, "<") == false    { return false }
	if strings.HasSuffix(argv0, ">") == false    { return false }
	if strings.Contains(argv0,  "@") == false    { return false }
	if IsEmailAddress(strings.Trim(argv0, "<>")) { return true  }
	return false
}

// IsComment() returns true if the string starts with "(" and ends with ")"
func IsComment(argv0 string) bool {
	// @param    string argv0    String including an comment in email address like "(neko, nyaan)"
	// @return   bool            true:  is a comment
	//                           false: is not a comment
	if len(argv0)                    == 0     { return false }
	if strings.HasPrefix(argv0, "(") == false { return false }
	if strings.HasSuffix(argv0, ")") == false { return false }
	return true
}

// IsEmailAddress() checks that the argument is an email address or not
func IsEmailAddress(email string) bool {
	// @param    string email    Email address string
	// @return   bool            true:  is an email address
	//                           false: is not an email address
	if len(email) < 5 { return false } // n@e.e

	// See http://www.ietf.org/rfc/rfc5322.txt
	//   or http://www.ex-parrot.com/pdw/Mail-RFC822-Address.html ...
	//   addr-spec       = local-part "@" domain
	//   local-part      = dot-atom / quoted-string / obs-local-part
	//   domain          = dot-atom / domain-literal / obs-domain
	//   domain-literal  = [CFWS] "[" *([FWS] dcontent) [FWS] "]" [CFWS]
	//   dcontent        = dtext / quoted-pair
	//   dtext           = NO-WS-CTL /     ; Non white space controls
	//                     %d33-90 /       ; The rest of the US-ASCII
	//                     %d94-126        ;  characters not including "[",
	//                                     ;  "]", or "\"
	email  = strings.Trim(email, " \t")
	lasta := strings.LastIndex(email, "@")
	lastd := strings.LastIndex(email, ".")

	if len(email)         > 254 { return false } // The maximum length of an email address is 254
	if lasta < 1 || lasta >  64 { return false } // The maximum lenght of a local part is 64
	if len(email) - lasta > 253 { return false } // The maximum lenght of a domain part is 253
	if email[0]           == 46 { return false } // "." at the first character is not allowed in a local part
	if email[lasta - 1]   == 46 { return false } // "." before the "@" is not allowed in a local part

	upper := strings.ToUpper(email)
	quote := IsQuotedAddress(email)
	match := true

	for j, e := range(strings.Split(email, "")) {
		// 31 < The ASCII code of each character < 127
		if j < lasta {
			// A local part of the email address: string before the last "@"
			if email[j] <  32 { match = false; break } // Before " "
			if email[j] > 126 { match = false; break } // After  "~"
			if j == 0 { continue }

			if quote {
				// The email address has quoted local part like "neko@nyaan"@example.org
				if email[j-1] != 34 {
					// When the previous character is not "\", '"', " ", and "\t" is not allowed
					if email[j] ==  9 { match = false; break } // "\t"
					if email[j] == 32 { match = false; break } // " "
					if email[j] == 34 { match = false; break } // "\"
					if email[j] == 92 { match = false; break } // '"'
				}
			} else {
				// The local part is not quoted
				// ".." is not allowed in a local part when the local part is not quoted by ""
				if e == "." && email[j-1] == 46 { match = false; break }

				// The following characters are not allowed in a local part without "..."@example.jp
				if e == "," || e == "@" { match = false; break }
				if e == ":" || e == ";" { match = false; break }
				if e == "(" || e == ")" { match = false; break }
				if e == "<" || e == ">" { match = false; break }
				if e == "p" || e == "]" { match = false; break }
			}
		} else {
			// A domain part of the email address: string after the last "@"
			if email[j] <   45 { match = false; break } // Before "-"
			if email[j] ==  47 { match = false; break } // Equals "/"
			if email[j] >  122 { match = false; break } // After  "z"

			if email[j] > 57 && email[j] < 64 { match = false; break } // ":" to "?"
			if email[j] > 90 && email[j] < 97 { match = false; break } // "[" to "`"

			if j > lastd {
				// *TLD of the domain part: string after the last "."
				if upper[j] < 65 { match = false; break } // Before "A"
				if upper[j] > 90 { match = false; break } // After  "Z"
			}
		}
	}

	return match
}

// IsMailerDaemon() checks that the argument is mailer-daemon or not
func IsMailerDaemon(email string) bool {
	// @param    string email    Email address
	// @return   bool            true:  is a mailer-daemon
	//                           false: is not a mailer-daemon
	match := false
	value := strings.ToLower(email)
	table := []string{
		"mailer-daemon@", "(mailer-daemon)", "<mailer-daemon>", "mailer-daemon ",
		"postmaster@", "(postmaster)", "<postmaster>",
	}
	for _, e := range table {
		if strings.Contains(value, e) || value == "mailer-daemon" || value == "postmaster" {
			match = true
			break
		}
	}
	return match
}

