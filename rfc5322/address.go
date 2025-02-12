// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____ _________  ____     ___       _     _                   
// |  _ \|  ___/ ___| ___|___ /___ \|___ \   / / \   __| | __| |_ __ ___  ___ ___ 
// | |_) | |_ | |   |___ \ |_ \ __) | __) | / / _ \ / _` |/ _` | '__/ _ \/ __/ __|
// |  _ <|  _|| |___ ___) |__) / __/ / __/ / / ___ \ (_| | (_| | | |  __/\__ \__ \
// |_| \_\_|   \____|____/____/_____|_____/_/_/   \_\__,_|\__,_|_|  \___||___/___/

package rfc5322
import "strings"
import "libsisimai.org/sisimai/rfc1123"

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
	if lasta < 1 || lasta >  64 { return false } // The maximum length of a local part is 64
	if len(email) - lasta > 253 { return false } // The maximum lenght of a domain part is 253

	// "." as the first character of the local part and ".@" are not allowed in a local part when
	// the local part is not quoted by "", but Non-RFC compliant email addresses still persist in
	// the world.
	// if email[0]         == 46 { return false } // '.' at the first character is not allowed in a local part
	// if email[lasta - 1] == 46 { return false } // '.' before the "@" is not allowed in a local part

	quote := IsQuotedAddress(email); if quote == false {
		// The email address is not a quoted address
		if strings.Count(email, "@") > 1 { return false }
		if strings.Contains(email, " ")  { return false }

		// Non-RFC compliant email addresses still persist in the world.
		// if strings.Contains(email, "..") { return false }
		// if strings.Contains(email, ".@") { return false }
	}
	upper := strings.ToUpper(email)
	ipv46 := rfc1123.IsDomainLiteral(email)
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
				// ".." is not allowed in a local part when the local part is not quoted by "" but
				// Non-RFC compliant email addresses still persist in the world.
				// if e == "." && email[j-1] == 46 { match = false; break }

				// The following characters are not allowed in a local part without "..."@example.jp
				if e == "," || e == "@" || e == ":" || e == ";" || e == "(" { match = false; break }
				if e == ")" || e == "<" || e == ">" || e == "[" || e == "]" { match = false; break }
			}
		} else {
			// A domain part of the email address: string after the last "@"
			if email[j] ==  64 {             continue } // '@'
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

	// Check that the domain part is a valid internet host or not
	if match == true && ipv46 == false { match = rfc1123.IsInternetHost(email[lasta + 1:]) }
	return match
}

// IsQuotedAddress() checks that the local part of the argument is quoted
func IsQuotedAddress(email string) bool {
	// @param    string email    Email address string
	// @return   bool            true:  the local part is quoted
	//                           false: the local part is not quoted
	if strings.HasPrefix(email, `"`) == false { return false }
	if strings.Contains(email, `"@`) == false { return false }
	return true
}

// IsComment() returns true if the string starts with "(" and ends with ")"
func IsComment(argv0 string) bool {
	// @param    string argv0    String including an comment in email address like "(neko, cat)"
	// @return   bool            true:  is a comment
	//                           false: is not a comment
	if argv0 == "" || !strings.HasPrefix(argv0, "(") || !strings.HasSuffix(argv0, ")") { return false }
	return true
}

