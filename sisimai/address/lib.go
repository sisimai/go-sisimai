// Copyright (C) 2020-2021 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address
import "fmt"
import "strings"

// Undisclosed() returns a pseudo recipient address or a pseudo sender address
func Undisclosed(a string) string {
	// @param    [string] a   Address type: 'r' or 's'
	// @return   [string]     Pseudo recipient address or sender address when the "a" is neither "r" nor "s"
	addr := ""
	if a == "s" { addr = "sender"    }
	if a == "r" { addr = "recipient" }
	return fmt.Sprintf("undisclosed-%s-in-headers@libsisimai.org.invalid", addr)
}

// IsQuotedAddress() checks that the local part of the argument is quoted
func IsQuotedAddress(email string) bool {
	// @param    [string] email  Email address string
	// @return   [bool]          true:  the local part is quoted
	//                           false: the local part is not quoted
	if strings.HasPrefix(email, `"`) == false { return false }
	if strings.Contains(email, `"@`) == false { return false }
	return true
}

// IsEmailAddress() checks that the argument is an email address or not
func IsEmailAddress(email string) bool {
	// @param    [string] email  Email address string
	// @return   [bool]          true:  is an email address
	//                           false: is not an email address
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

			if email[j] > 57 && email[j] < 65 { match = false; break } // ":" to "@"
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
	// @param    [string] email  Email address
	// @return   [bool]          true:  is a mailer-daemon
	//                           false: is not a mailer-daemon
	match := false
	value := strings.ToLower(email)
	table := []string {
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

// S3S4() runs like ruleset 3,4 of sendmail.cf
func S3S4(argv0 string) string {
	// @param    [string] input  Text including an email address
	// @return   [string]        Email address without comment, brackets
	if len(argv0) == 0 { return "" }

	list := Find(argv0)
	if len(list) == 0 { return argv0 }
	return list[0]
}

// ExpandVERP() gets the original recipient address from VERP
func ExpandVERP(email string) string {
	// @param    [string] email  VERP Address
	// @return   [string]        Email address
	if len(email) == 0 { return "" }
	local := strings.SplitN(email, "@", 2)[0]

	// bounce+neko=example.org@example.org => neko@example.org
	if strings.Index(local, "+") < 1 { return "" }
	if strings.Index(local, "=") < 1 { return "" }
	if strings.Index(local, "+") > len(local) - 1 { return "" }
	if strings.Index(local, "=") > len(local) - 1 { return "" }

	verp1 := strings.Replace(strings.SplitN(local, "+", 2)[1], "=", "@", 1)
	if IsEmailAddress(verp1) { return verp1 }
	return ""
}

// ExpandAlias() removes string from "+" to "@" at a local part
func ExpandAlias(email string) string {
	// @param    [string] email  Email alias string
	// @return   [string]        Expanded email address
	if len(email) == 0 { return "" }
	if !IsEmailAddress(email) { return "" }
	if !strings.Contains(email, "+") { return "" }
	if strings.Index(email, "+") < 1 { return "" }

	// neko+straycat@example.org => neko@example.org
	lpart := email[0:strings.Index(email, "+")]
	dpart := strings.SplitN(email, "@", 2)[1]
	return lpart + "@" + dpart
}

