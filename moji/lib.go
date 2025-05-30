// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _ 
//  _ __ ___   ___ (_|_)
// | '_ ` _ \ / _ \| | |
// | | | | | | (_) | | |
// |_| |_| |_|\___// |_|
//               |__/   

// Package "moji" provides functions for dealing strings
// "moji" stands for "character" or "string" in Japanese
package moji
import "fmt"
import "strings"
import "crypto/sha1"

const LHS string = "<@>" // The LHS string for Select() function
const RHS string = "<$>" // The RHS string for Select() function

// Token creates the message token from an addresser, and a recipient, and an unix machine time.
//   Arguments:
//     - argv1 (string): Email address of the sender
//     - argv2 (string): Email address of the recipient
//     - epoch (int):    Machine time of the bounce
//   Returns:
//     - (string):       Message token(SHA1 hex digest) or empty string
func Token(argv1 string, argv2 string, epoch int) string {
	// - http://en.wikipedia.org/wiki/ASCII
	if argv1 == "" || len(argv2) == 0 { return "" }

	// Format: STX(0x02) Sender-Address RS(0x1e) Recipient-Address ETX(0x03)
	plain := fmt.Sprintf("\x02%s\x1e%s\x1e%d\x03", strings.ToLower(argv1), strings.ToLower(argv2), epoch)
	crypt := sha1.New(); crypt.Write([]byte(plain))
	return fmt.Sprintf("%x", crypt.Sum(nil))
}

// Squeeze remove redundant characters from the given string
//   Arguments:
//     - argv0 (*string): String including redundant characters like "neko  chan"
//     - argv1 (byte):    Characters to be squeezed, for example ' '
func Squeeze(argv0 *string, argv1 byte) {
	if argv0 == nil || *argv0 == "" || strings.IndexByte(*argv0, argv1) < 0 { return }

	textbuffer := make([]byte, 0, len(*argv0))
	cb := byte(0); for _, by := range []byte(*argv0) {
		// Remove a character that is the same character of the previous character
		if by == argv1 && by == cb { continue }
		textbuffer = append(textbuffer, by)
		cb = by
	}
	*argv0 = string(textbuffer)
}

// Sweep clears the string out.
//   Arguments:
//     - argv1 (string): String to be cleaned
//   Returns:
//     - (string):       Cleaned out string
func Sweep(argv1 string) string {
	if argv1 == "" { return "" }

	argv1 = strings.TrimSpace(strings.ReplaceAll(argv1, "\t", " ")); Squeeze(&argv1, ' ')
	if strings.Contains(argv1, " --") && strings.Contains(argv1, "-- ") == false {
		// Delete all the string after a boundary string like " --neko-chan"
		argv1 = Select(LHS + argv1, "", " --", 0)
	}
	return argv1
}

// ContainsOnlyNumbers returns true when the given string contain numbers only.
func ContainsOnlyNumbers(argv1 string) bool {
	if argv1 == "" { return false }
	for _, e := range argv1 { if e < 48 || e > 57 { return false } }
	return true
}

// Aligned checks if each element of the 2nd argument is aligned in the 1st argument or not.
//   Arguments:
//     - argv1 (string):   String to be checked such as "I am a cat. I have, as yet, no name."
//     - argv2 ([]string): List including the ordered strings such as []string{"cat", "yet"}
//   Returns:
//     - (bool):           true if the all strings are ordered in argv1, false otherwise.
func Aligned(argv1 string, argv2 []string) bool {
	if argv1 == "" || len(argv2) == 0 { return false }

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

	return right == len(argv2)
}

// IndexOnTheWay returns the index of the first string of argv1 finding after the start position in argv0
//   Arguments:
//     - argv0 (string): The string to be searched
//     - argv1 (string): The substring to search for
//     - start (int):    The index from which to start the search
//   Returns:
//     - (int):          The index of argv1
func IndexOnTheWay(argv0, argv1 string, start int) int {
	if start < 0 || start >= len(argv0)                    { return -1 }
	fi := strings.Index(argv0[start:], argv1); if fi == -1 { return -1 }
	return fi + start
}

// Select returns a string selected between the 2nd argument and 3rd argument from the 1st argument.
//   Arguments:
//     - argv0 (string): The string to be searched for example "From: <neko@example.jp>"
//     - begin (string): Substring such as "<"
//     - until (string): Substring such as ">"
//     - start (int):    The index position for seeking
//   Returns:
//     - (string):       Selected string such as "neko@example.jp"
func Select(argv0, begin, until string, start int) string {
	if argv0 == "" || start < 0 { return ""   }
	if begin == "" /* <@> */    { begin = LHS }
	if until == "" /* <$> */    { until = RHS }

	textlength := [3]int{len(argv0), len(begin), len(until)}
	sourcetext := argv0

	if start > 0 {
		if start > textlength[0] - 2 { return "" }
		sourcetext = argv0[start:]
		textlength[0] = len(sourcetext)
	}

	if textlength[0] < 3 || textlength[0] <= (textlength[1] + textlength[2]) { return "" }
	indextable   := [3]int{0, -1, -1}
	indextable[1] = strings.Index(sourcetext, begin); if indextable[1] == -1 { return "" }
	indextable[2] = strings.Index(sourcetext[indextable[1] + textlength[1] + 1:], until)

	if indextable[2] < 0 { return "" }; indextable[2] += indextable[1] + textlength[1] + 1
	return sourcetext[indextable[1] + textlength[1]:indextable[2]]
}

