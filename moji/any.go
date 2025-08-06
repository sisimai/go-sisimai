// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _ 
//  _ __ ___   ___ (_|_)
// | '_ ` _ \ / _ \| | |
// | | | | | | (_) | | |
// |_| |_| |_|\___// |_|
//               |__/   

package moji
import "strings"

// ContainsAny checks whether any element in "list" is included in "text" or not. (text > list)
//   Arguments:
//     - text (string):   String containing any substring listed in "list".
//     - list ([]string): List of strings to find in "text".
//   Returns:
//     - (bool): true if one or more string in "list" was found in "text".
func ContainsAny(text string, list []string) bool {
	if text == "" || len(list) == 0 { return false }

	// It works like `grep { index($e, $_) > -1 } @list` in Perl
	for _, e := range list { if strings.Contains(text, e) { return true } }
	return false
}

// HasPrefixAny checks whether any alement in "list" starts with the "text" or not.
//   Arguments:
//     - text (string):   String containing any substring listed in "list".
//     - list ([]string): List of strings to find in "text".
//   Returns:
//     - (bool): true if the string in "text" starts with any string listed in "list".
func HasPrefixAny(text string, list []string) bool {
	if text == "" || len(list) == 0 { return false }

	// It works like `grep { index($e, $_) == 0 } @list` in Perl
	for _, e := range list { if strings.HasPrefix(text, e) { return true } }
	return false
}

// AlignedAny checks if each slice of the 2nd argument is aligned in the 1st argument or not.
//   Arguments:
//     - text (string):     String to be checked such as "I am a cat. I have, as yet, no name.".
//     - list ([][]string): List including the ordered strings such as `[][]string{[]striing{"cat", "yet"}}`.
//   Returns:
//     - (bool): true if the all strings are ordered in "text", false otherwise.
func AlignedAny(text string, list [][]string) bool {
	if text == "" || len(list) == 0 { return false }
	for _, e := range list { if p := Aligned(text, e); p == true { return true } }
	return false
}

