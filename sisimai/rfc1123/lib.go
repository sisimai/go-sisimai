// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc1123

//  ____  _____ ____ _ _ ____  _____ 
// |  _ \|  ___/ ___/ / |___ \|___ / 
// | |_) | |_ | |   | | | __) | |_ \ 
// |  _ <|  _|| |___| | |/ __/ ___) |
// |_| \_\_|   \____|_|_|_____|____/ 
import "strings"

// IsValidHostname() returns "true" when the given string is a valid hostname
func IsValidHostname(argv1 string) bool {
	// @param    string argv1  Hostname
	// @return   bool          true:  is a valid hostname
	//                         false: is not a valid hostname
	// @see https://datatracker.ietf.org/doc/html/rfc1123
	if len(argv1) > 255                      { return false }
	if strings.Contains(argv1, ".") == false { return false }
	if strings.Contains(argv1, "..") == true { return false }
	if strings.HasPrefix(argv1, "-") == true { return false }
	if strings.HasPrefix(argv1, ".") == true { return false }
	if strings.HasSuffix(argv1, "-") == true { return false }
	if strings.HasSuffix(argv1, ".") == true { return false }

	hostnameok := true
	for _, e := range strings.Split(strings.ToUpper(argv1), "") {
		// Check each octet is between 0 and 255
		if e[0] <  45              { hostnameok = false; break } //  45 = '-'
		if e[0] == 47              { hostnameok = false; break } //  47 = '/'
		if e[0] >  57 && e[0] < 65 { hostnameok = false; break } //  57 = '9', 65 = 'A'
		if e[0] >  90              { hostnameok = false; break } //  90 = 'Z'
	}
	if hostnameok == false { return false }

	for _, e := range strings.Split(strings.Split(argv1, ".")[strings.Count(argv1, ".")], "") {
		// The top level domain should not include a number
		if e[0] > 47 && e[0] < 58  { hostnameok = false; break }
	}
	return hostnameok
}

