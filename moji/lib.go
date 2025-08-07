// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _ 
//  _ __ ___   ___ (_|_)
// | '_ ` _ \ / _ \| | |
// | | | | | | (_) | | |
// |_| |_| |_|\___// |_|
//               |__/   

// Package "moji" provides functions for dealing strings. "moji" stands for "character" or "string" in Japanese.
package moji
import "fmt"
import "strings"
import "crypto/sha1"

const LHS string = "<@>" // The LHS string for Select() function
const RHS string = "<$>" // The RHS string for Select() function

// Token creates the message token from an addresser, and a recipient, and an unix machine time.
//   Arguments:
//     - addre (string): Email address of the sender.
//     - recip (string): Email address of the recipient.
//     - epoch (int):    Machine time of the bounce.
//   Returns:
//     - (string): Message token(SHA1 hex digest) or empty string.
func Token(addre string, recip string, epoch int) string {
	// - http://en.wikipedia.org/wiki/ASCII
	if addre == "" || len(recip) == 0 { return "" }

	// Format: STX(0x02) Sender-Address RS(0x1e) Recipient-Address ETX(0x03)
	plain := fmt.Sprintf("\x02%s\x1e%s\x1e%d\x03", strings.ToLower(addre), strings.ToLower(recip), epoch)
	crypt := sha1.New(); crypt.Write([]byte(plain))
	return fmt.Sprintf("%x", crypt.Sum(nil))
}

// Squeeze remove redundant characters from the given string.
//   Arguments:
//     - text (*string): String including redundant characters like "neko  chan".
//     - char (byte):    Characters to be squeezed, for example ' '.
func Squeeze(text *string, char byte) {
	if text == nil || *text== "" || strings.IndexByte(*text, char) < 0 { return }

	textbuffer := make([]byte, 0, len(*text))
	cb := byte(0); for _, by := range []byte(*text) {
		// Remove a character that is the same character of the previous character
		if by != char || by != cb { textbuffer = append(textbuffer, by); cb = by }
	}
	*text = string(textbuffer)
}

// Sweep clears the string out.
//   Arguments:
//     - text (string): String to be cleaned.
//   Returns:
//     - (string): Cleaned out string.
func Sweep(text string) string {
	if text == "" { return "" }

	text = strings.TrimSpace(strings.ReplaceAll(text, "\t", " ")); Squeeze(&text, ' ')
	if strings.Contains(text, " --") && strings.Contains(text, "-- ") == false {
		// Delete all the string after a boundary string like " --neko-chan"
		text = Select(LHS + text, "", " --", 0)
	}
	return text 
}

// ContainsOnlyNumbers returns true when the given string contain numbers only.
//   Arguments:
//     - text (string): String including only numbers such as "2022"
//   Returns:
//     - (string): true if the string contain only numbers
func ContainsOnlyNumbers(text string) bool {
	if text == "" { return false }
	for _, e := range text { if e < 48 || e > 57 { return false } }
	return true
}

// IsContained checks whether any element in "list" includes "text" or not. (text < list)
//   Arguments:
//     - text (string):   String to be contained as a substring listed in "list".
//     - list ([]string): List of strings.
//   Returns:
//     - (bool): true if one or more string in "text" was found in "list".
func IsContained(text string, list []string) bool {
	if text == "" || len(list) == 0 { return false }
	for _, e := range list { if strings.Contains(e, text) { return true } }
	return false
}

// Aligned checks if each element of the 2nd argument is aligned in the 1st argument or not.
//   Arguments:
//     - text (string):   String to be checked such as "I am a cat. I have, as yet, no name.".
//     - sort ([]string): List including the ordered strings such as `[]string{"cat", "yet"}`.
//   Returns:
//     - (bool): true if the all strings in "sort" are ordered in "text", false otherwise.
func Aligned(text string, sort []string) bool {
	if text == "" || len(sort) == 0 { return false }

	align, right := -1, 0; for _, e := range sort {
		// Get the position of each element in the 1st argument using index()
		if align > 0 { text = text[align + 1:] }
		p := strings.Index(text, e)

		if p < 0 { break }      // Break this loop when there is no string in the 1st argument
		align = len(e) + p - 1  // There is an aligned string in the 1st argument
		right++
	}
	return right == len(sort)
}

// IndexOnTheWay returns the index of the first string of "parts" finding after the start position in "whole".
//   Arguments:
//     - whole (string): The whole string to be searched.
//     - parts (string): The substring to search for.
//     - start (int):    The index from which to start the search.
//   Returns:
//     - (int): The index of parts.
func IndexOnTheWay(whole, parts string, start int) int {
	if start < 0 || start >= len(whole)                    { return -1 }
	fi := strings.Index(whole[start:], parts); if fi == -1 { return -1 }
	return fi + start
}

// Select returns a string selected between the 2nd argument and 3rd argument from the 1st argument.
//   Arguments:
//     - argv0 (string): The string to be searched for example "From: <neko@example.jp>".
//     - begin (string): Substring such as "<".
//     - until (string): Substring such as ">".
//     - start (int):    The index position for seeking.
//   Returns:
//     - (string): Selected string such as "neko@example.jp".
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

