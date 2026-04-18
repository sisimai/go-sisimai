// Copyright (C) 2020-2021,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _ 
//  _ __ ___   ___ (_|_)
// | '_ ` _ \ / _ \| | |
// | | | | | | (_) | | |
// |_| |_| |_|\___// |_|
//               |__/   

// Package "moji" provides functions for dealing strings. "moji" stands for "character" or "string" in Japanese.
package moji
import "strings"

const LHS string = "<@>" // The LHS string for Select() function
const RHS string = "<$>" // The RHS string for Select() function

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

// IsText checks that the text is a plain text or not.
//   Arguments:
//     - text (*string):  String to be checked that the text is a plain text or not.
//   Returns:
//     - (bool): true if the first 10 bytes indicates that the text is a text.
func IsText(text *string) bool {
	if text == nil || len(*text) == 0 { return false }

	cw := len(*text); for j := 0; j < 10; j++ {
		// - Disallow DEL(0x7F:127) or later.
		// - Allow HT(0x09:9), LF(0x0A:10), VT(0x0B:11), FF(0x0C:12), CR(0x0D:13)
		if j >= cw                        { break        }
		cv := (*text)[j]; if cv > 126     { return false }
		if cv < 32 && (cv < 9 || cv > 13) { return false }
	}
	return true
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
//     - whole (string): The whole string to be searched for example "From: <neko@example.jp>".
//     - begin (string): Substring such as "<".
//     - until (string): Substring such as ">".
//     - start (int):    The index position for seeking.
//   Returns:
//     - (string): Selected string such as "neko@example.jp".
func Select(whole, begin, until string, start int) string {
	if whole == "" || start < 0 { return ""   }
	if start > (len(whole) - 2) { return ""   }
	if begin == "" /* <@> */    { begin = LHS }
	if until == "" /* <$> */    { until = RHS }

	cv := whole[start:]
	cw := [3]int{len(cv), len(begin), len(until)}
	if cw[0] < 3 || cw[0] <= (cw[1] + cw[2]) { return "" }

	ci    := [3]int{0, -1, -1}
	ci[1]  = strings.Index(cv, begin);                     if ci[1] < 0 { return "" }
	ci[2]  = strings.Index(cv[ci[1] + cw[1] + 1:], until); if ci[2] < 0 { return "" }
	ci[2] += ci[1] + cw[1] + 1
	return cv[ci[1] + cw[1]:ci[2]]
}

