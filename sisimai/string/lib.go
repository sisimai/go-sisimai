// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string
import "fmt"
import "strings"
import "strconv"
import "crypto/sha1"

// Token() creates the message token from an addresser, and a recipient, and an unix machine time
func Token(argv1 string, argv2 string, epoch int) string {
	// @param    string addr1  A sender's email address
	// @param    string addr2  A recipient's email address
	// @param    int    epoch  Machine time of the email bounce
	// @return   string        Message token(MD5 hex digest) or empty string if the any argument is missing
	// @see      http://en.wikipedia.org/wiki/ASCII
	if len(argv1) == 0 { return "" }
	if len(argv2) == 0 { return "" }

	// Format: STX(0x02) Sender-Address RS(0x1e) Recipient-Address ETX(0x03)
	plain := fmt.Sprintf("\x02%s\x1e%s\x1e%d\x03", strings.ToLower(argv1), strings.ToLower(argv2), epoch)
	crypt := sha1.New(); crypt.Write([]byte(plain))
	token := crypt.Sum(nil)

	return fmt.Sprintf("%x", token)
}

// Squeeze() remove redundant characters
func Squeeze(argv1 string, chars string) string {
	// @param    string argv1  Email address, name, and other elements
	// @return   EmailAddress  EmailAddress struct when the email address was not valid
	if len(argv1) == 0 { return ""    }
	if len(chars) == 0 { return argv1 }

	for _, e := range strings.Split(chars, "") {
		// Remove redundant characters from "argv1"
		for strings.Count(argv1, e + e) > 1 {
			// e + e => e
			e = strings.ReplaceAll(argv1, e + e, e)
		}
	}
	return argv1
}

// Sweep() clears the string out
func Sweep(argv1 string) string {
	// @param    string argv1  String to be cleaned
	// @return   string        Cleaned out string
	if len(argv1) == 0 { return "" }
	argv1 = Squeeze(argv1, " ")
	argv1 = strings.ReplaceAll(argv1, "\t", "")
	argv1 = strings.TrimSpace(argv1)

	if strings.Contains(argv1, " --") {
		// Delete all the string after a boundary string like " --neko-nyaan"
		for {
			if strings.Contains(argv1, "-- ")  { break }
			if strings.Contains(argv1, "--\t") { break }

			argv1 = argv1[0:strings.Index(argv1, " --")-1]
			break
		}
	}
	return argv1
}

// ContainsOnlyNumbers() returns true when the given string contain numbers only
func ContainsOnlyNumbers(argv1 string) bool {
	// @param    string argv1  String
	// @return   bool          true, false
	if len(argv1) == 0 { return false }

	match := true
	for _, e := range argv1 { if e < 48 || e > 57 { match = false; break } }

	return match
}

// IsIPv4Address() returns "true" when the given string is an IPv4 address
func IsIPv4Address(argv1 string) bool {
	// @param    string argv1  IPv4 address like "192.0.2.25"
	// @return   bool          true:  is an IPv4 address
	//                         false: is not an IPv4 address
	if len(argv1) < 7                 { return false }
	if strings.Count(argv1, ".") != 3 { return false }

	match := true
	for _, e := range strings.Split(argv1, ".") {
		// Check each octet is between 0 and 255
		v, oops := strconv.Atoi(e)
		if oops != nil { match = false; break }
		if v < 0       { match = false; break }
		if v > 255     { match = false; break }
	}
	return match
}

// FindIPv4Address() find IPv4 addresses from the given string
func FindIPv4Address(argv1 string) []string {
	// @param    string   argv1  String including an IPv4 address
	// @return   []string        List of IPv4 addresses
	// @since    v5.2.0
	if len(argv1) < 7 { return []string{} }

	// Rewrite: "mx.example.jp[192.0.2.1]" => "mx.example.jp 192.0.2.1"
	strings.Replace(argv1, "(", " ", 1)
	strings.Replace(argv1, ")", " ", 1)
	strings.Replace(argv1, "[", " ", 1)
	strings.Replace(argv1, "]", " ", 1)

	ipv4a := []string{}
	for _, e := range strings.Split(argv1, " ") {
		// Find a string including an IPv4 address
		if !strings.Contains(e, ".") { continue }	// IPv4 address must include "." character
		if !IsIPv4Address(e)         { continue }	// The string is an IPv4 address or not
		ipv4a = append(ipv4a, e)
	}
	return ipv4a
}

// Aligned() checks if each element of the 2nd argument is aligned in the 1st argument or not
func Aligned(argv1 string, argv2 []string) bool {
	// @param    string   argv1  String to be checked
	// @param    []string argv2  List including the ordered strings
	// @return   bool
	// @since    v5.2.0
	if len(argv1) == 0 { return false }
	if len(argv2) == 0 { return false }

	align := -1
	right :=  0

	for _, e := range argv2 {
		// Get the position of each element in the 1st argument using index()
		if align > 0 { argv1 = argv1[align + 1:] }
		p := strings.Index(argv1, e)

		if p < 0 { break }		// Break this loop when there is no string in the 1st argument
		align = len(e) + p - 1	//  There is an aligned string in the 1st argument
		right++
	}

	if right == len(argv2) { return true }
	return false
}

