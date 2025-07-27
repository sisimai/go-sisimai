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

// IsContained checks whether any element in argv2 includes argv1 or not. (argv1 < argv2)
//   Arguments:
//     - argv1 (string):   String to be contained as a substring listed in argv2.
//     - argv2 ([]string): List of strings.
//   Returns:
//     - (bool): true if one or more string in argv1 was found in argv2.
func IsContained(argv1 string, argv2 []string) bool {
	if argv1 == "" || len(argv2) == 0 { return false }
	for _, e := range argv2 { if strings.Contains(e, argv1) { return true } }
	return false
}

// ContainsAny checks whether any element in argv2 is included in argv1 or not. (argv1 > argv2)
//   Arguments:
//     - argv1 (string):   String containing any substring listed in argv2.
//     - argv2 ([]string): List of strings to find in argv1.
//   Returns:
//     - (bool): true if one or more string in argv2 was found in argv1.
func ContainsAny(argv1 string, argv2 []string) bool {
	if argv1 == "" || len(argv2) == 0 { return false }

	// It works like `grep { index($e, $_) > -1 } @list` in Perl
	for _, e := range argv2 { if strings.Contains(argv1, e) { return true } }
	return false
}

// HasPrefixAny checks whether any alement in argv2 starts with the argv1 or not.
//   Arguments:
//     - argv1 (string):   String containing any substring listed in argv2.
//     - argv2 ([]string): List of strings to find in argv1.
//   Returns:
//     - (bool): true if the string in argv1 starts with any string listed in argv2.
func HasPrefixAny(argv1 string, argv2 []string) bool {
	if argv1 == "" || len(argv2) == 0 { return false }

	// It works like `grep { index($e, $_) == 0 } @list` in Perl
	for _, e := range argv2 { if strings.HasPrefix(argv1, e) { return true } }
	return false
}

// AlignedAny checks if each slice of the 2nd argument is aligned in the 1st argument or not.
//   Arguments:
//     - argv1 (string):     String to be checked such as "I am a cat. I have, as yet, no name.".
//     - argv2 ([][]string): List including the ordered strings such as `[][]string{[]striing{"cat", "yet"}}`.
//   Returns:
//     - (bool): true if the all strings are ordered in argv1, false otherwise.
func AlignedAny(argv1 string, argv2 [][]string) bool {
	if argv1 == "" || len(argv2) == 0 { return false }
	for _, e := range argv2 { if p := Aligned(argv1, e); p == true { return true } }
	return false
}

