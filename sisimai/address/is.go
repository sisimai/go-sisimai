// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/
import "strings"
import "sisimai/rfc791"

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
	// @param    string argv0    String including an email address like "<neko@example.jp>"
	// @return   bool            true:  is including an email address
	//                           false: is not including an email address
	if len(argv0)                     < 5     { return false }
	if strings.Contains(argv0,  "@") == false { return false }

	if strings.HasPrefix(argv0, "<") && strings.HasSuffix(argv0, ">") {
		// The argument is like "<neko@example.jp>"
		if IsEmailAddress(strings.Trim(argv0, "<>")) { return true }
		return false

	} else {
		// Such as "nekochan (kijitora) neko@example.jp"
		for _, e := range strings.Split(argv0, " ") {
			// Is there any email address string in each element?
			if IsEmailAddress(e) { return true }
		}
	}
	return false
}

// IsComment() returns true if the string starts with "(" and ends with ")"
func IsComment(argv0 string) bool {
	// @param    string argv0    String including an comment in email address like "(neko, cat)"
	// @return   bool            true:  is a comment
	//                           false: is not a comment
	if len(argv0)                    == 0     { return false }
	if strings.HasPrefix(argv0, "(") == false { return false }
	if strings.HasSuffix(argv0, ")") == false { return false }
	return true
}

// IsDomainLiteral() returns true if the domain part is [IPv4:...] or [IPv6:...]
func IsDomainLiteral(email string) bool {
	// @param    string email    Email address string
	// @return   bool            true:  is an domain-literal
	//                           false: is not an domain-literal
	email = strings.Trim(email, "<>")
	if len(email)                     < 16    { return false } // e@[IPv4:0.0.0.0] is 16 characters
	if strings.HasSuffix(email, "]") == false { return false }

	if strings.Contains(email, "@[IPv4:") {
		// neko@[IPv4:192.0.2.25]
		p1 := strings.Index(email, "@[IPv4:")
		cv := email[p1 + 7:len(email) - 1]
		return rfc791.IsIPv4Address(&cv)

	} else if strings.Contains(email, "@[IPv6:") {
		// neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]
		p1 := strings.Index(email, "@[IPv6:")
		cv := email[p1 + 7:len(email) - 1]
		if len(cv) == 39 && strings.Count(cv, ":") == 7 { return true }
	}
	return false
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
	if email[0]           == 46 { return false } // '.' at the first character is not allowed in a local part
	if email[lasta - 1]   == 46 { return false } // '.' before the "@" is not allowed in a local part

	quote := IsQuotedAddress(email); if quote == false {
		// The email address is not a quoted address
		if strings.Contains(email, " ")  { return false }
		if strings.Contains(email, "..") { return false }
		if strings.Contains(email, ".@") { return false }
		if strings.Count(email, "@") > 1 { return false }
	}
	upper := strings.ToUpper(email)
	ipv46 := IsDomainLiteral(email)
	match := true

	for j, e := range(strings.Split(email, "")) {
		// 31 < The ASCII code of each character < 127
		if j < lasta {
			// A local part of the email address: string before the last "@"
			if email[j]  <  32 { match = false; break } // Before ' '
			if email[j]  > 126 { match = false; break } // After  '~'
			if j        ==   0 {             continue } // The character is the first character

			if jp := email[j - 1]; quote == true {
				// The email address has quoted local part like "neko@cat"@example.org
				if jp == 92 { // 92 = '\'
					// When the previous character IS '\', only the followings are allowed: '\', '"'
					if email[j] != 92 && email[j] != 34 { match = false; break }

				} else {
					// When the previous character IS NOT '\'
					if email[j] == 34 && j + 1 < lasta  { match = false; break } // `"` is allowed only immediately before the `@`
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
				if e == "[" || e == "]" { match = false; break }
			}
		} else {
			// A domain part of the email address: string after the last "@"
			if email[j] <   45 { match = false; break } // Before '-'
			if email[j] ==  47 { match = false; break } // Equals '/'
			if email[j] ==  92 { match = false; break } // Equals '\'
			if email[j] >  122 { match = false; break } // After  'z'

			if ipv46 == false {
				// Such as "example.jp", "neko.example.org"
				if email[j] >  57 && email[j] <  64 { match = false; break } // ':' to '?'
				if email[j] >  90 && email[j] <  97 { match = false; break } // '[' to '`'

			} else {
				// Such as "[IPv4:192.0.2.25]"
				if email[j] >  59 && email[j] <  64 { match = false; break } // ';' to '?'
				if email[j] >  93 && email[j] <  97 { match = false; break } // '^' to '`'
			}

			if j > lastd && ipv46 == false {
				// *TLD of the domain part: string after the last '.'
				if upper[j] < 65 { match = false; break } // Before 'A'
				if upper[j] > 90 { match = false; break } // After  'Z'
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

