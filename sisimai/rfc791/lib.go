// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc791

//  ____  _____ ____ _____ ___  _ 
// |  _ \|  ___/ ___|___  / _ \/ |
// | |_) | |_ | |      / / (_) | |
// |  _ <|  _|| |___  / / \__, | |
// |_| \_\_|   \____|/_/    /_/|_|
import "strings"
import "strconv"

// IsIPv4Address() returns "true" when the given string is an IPv4 address
func IsIPv4Address(argv1 *string) bool {
	// @param    *string argv1 IPv4 address like "192.0.2.25"
	// @return   bool          true:  is an IPv4 address
	//                         false: is not an IPv4 address
	if argv1 == nil || len(*argv1) < 7 { return false }
	if strings.Count(*argv1, ".") != 3 { return false }

	for _, e := range strings.Split(*argv1, ".") {
		// Check each octet is between 0 and 255
		v, nyaan := strconv.Atoi(e)
		if nyaan != nil { return false }
		if v < 0        { return false }
		if v > 255      { return false }
	}
	return true
}

// FindIPv4Address() find IPv4 addresses from the given string
func FindIPv4Address(argv1 *string) []string {
	// @param    *string   argv1 String including an IPv4 address
	// @return   []string        List of IPv4 addresses
	// @since    v5.2.0
	if argv1 == nil || len(*argv1) < 7 { return []string{} }

	// Rewrite: "mx.example.jp[192.0.2.1]" => "mx.example.jp 192.0.2.1"
	argv2 := *argv1
	argv2  = strings.ReplaceAll(argv2, "(", " ",)
	argv2  = strings.ReplaceAll(argv2, ")", " ",)
	argv2  = strings.ReplaceAll(argv2, "[", " ",)
	argv2  = strings.ReplaceAll(argv2, "]", " ",)
	argv2  = strings.ReplaceAll(argv2, ",", " ",)
	ipv4a := []string{}

	for _, e := range strings.Split(argv2, " ") {
		// Find a string including an IPv4 address
		if !strings.Contains(e, ".") { continue }   // IPv4 address must include "." character
		if !IsIPv4Address(&e)        { continue }   // The string is an IPv4 address or not
		ipv4a = append(ipv4a, e)
	}
	return ipv4a
}
