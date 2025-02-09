// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package string

//      _        _             
//  ___| |_ _ __(_)_ __   __ _ 
// / __| __| '__| | '_ \ / _` |
// \__ \ |_| |  | | | | | (_| |
// |___/\__|_|  |_|_| |_|\__, |
//                       |___/ 
import "strings"

// ContainsAny() checks whether any element in argv2 is included in argv1 or not
func ContainsAny(argv1 string, argv2 []string) bool {
	// @param    string   argv1 A string containing any substring listed in argv2
	// @param    []string argv2 A list of strings to find in argv1
	// @return   bool
	if argv1 == "" || len(argv2) == 0 { return false }

	for _, e := range argv2 {
		// It works like `grep { index($e, $_) > -1 } @list` in Perl
		if strings.Contains(argv1, e) { return true  }
	}
	return false
}

// EqualsAny() checks whether any element in argv2 is equal to the argv1 or not
func EqualsAny(argv1 string, argv2 []string) bool {
	// @param    string   argv1 A string containing any substring listed in argv2
	// @param    []string argv2 A list of strings to find in argv1
	// @return   bool
	if argv1 == "" || len(argv2) == 0 { return false }

	for _, e := range argv2 {
		// It works like `grep { $e eq $_ } @list` in Perl
		if argv1 == e { return true }
	}
	return false
}

// HasPrefixAny() checks whether any alement in argv2 starts with the argv1 or not
func HasPrefixAny(argv1 string, argv2 []string) bool {
	// @param    string   argv1 A string containing any substring listed in argv2
	// @param    []string argv2 A list of strings to find in argv1
	// @return   bool
	if argv1 == "" || len(argv2) == 0 { return false }

	for _, e := range argv2 {
		// It works like `grep { index($e, $_) == 0 } @list` in Perl
		if strings.HasPrefix(argv1, e) { return true }
	}
	return false
}

