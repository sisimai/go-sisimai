// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string

//      _        _             
//  ___| |_ _ __(_)_ __   __ _ 
// / __| __| '__| | '_ \ / _` |
// \__ \ |_| |  | | | | | (_| |
// |___/\__|_|  |_|_| |_|\__, |
//                       |___/ 
import "fmt"
import "strings"
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

// Is8Bit() checks the argument is including an 8-bit character or not
func Is8Bit(argv1 *string) bool {
	// @param    *string argv1  Any string to be checked
	// @return   bool           false:  ASCII Characters only
	//                          true:  Including an 8-bit character
	eight := false
	for _, e := range *argv1 {
		if e < 128 { continue }
		eight = true
		break
	}
	return eight
}

// Squeeze() remove redundant characters
func Squeeze(argv1 string, chars string) string {
	// @param    string argv1  String including redundant characters like "neko  nyaan"
	// @param    string chars  Characters to be squeezed 
	// @return   string        Squeezed string like "neko nyaan"
	if len(argv1) == 0 { return ""    }
	if len(chars) == 0 { return argv1 }

	for strings.Contains(argv1, chars + chars) {
		// Remove redundant characters from "argv1"
		argv1 = strings.ReplaceAll(argv1, chars + chars, chars)
	}
	return argv1
}

// Sweep() clears the string out
func Sweep(argv1 string) string {
	// @param    string argv1  String to be cleaned
	// @return   string        Cleaned out string
	if len(argv1) == 0 { return "" }

	argv1 = strings.ReplaceAll(argv1, "\t", " ")
	argv1 = strings.TrimSpace(argv1)
	argv1 = Squeeze(argv1, " ")

	for strings.Contains(argv1, " --") {
		// Delete all the string after a boundary string like " --neko-nyaan"
		if strings.Contains(argv1, "-- ")  { break }
		if strings.Contains(argv1, "--\t") { break }
		argv1 = argv1[0:strings.Index(argv1, " --") - 1]
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

		if p < 0 { break }      // Break this loop when there is no string in the 1st argument
		align = len(e) + p - 1  // There is an aligned string in the 1st argument
		right++
	}

	if right == len(argv2) { return true }
	return false
}

// IndexOnTheWay() returns the index of the first instance of argv1 after argv2 in argv0
func IndexOnTheWay(argv0, argv1 string, start int) int {
	// @param    string argv0  The string to be searched
	// @param    string argv1  The substring to search for
	// @param    int    start  The index from which to start the search
	// @return   string        The index of argv1
	if start < 0 || start >= len(argv0) { return -1 }
	return strings.Index(argv0[start:], argv1) + start
}

