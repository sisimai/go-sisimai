// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _____ ___  _ 
// |  _ \|  ___/ ___|___  / _ \/ |
// | |_) | |_ | |      / / (_) | |
// |  _ <|  _|| |___  / / \__, | |
// |_| \_\_|   \____|/_/    /_/|_|

// Package "rfc791" provides functions related to IPv4 address described in RFC791.
// https://datatracker.ietf.org/doc/html/rfc791
package rfc791

// IsIPv4Address returns "true" when the given string is an IPv4 address.
//   Arguments:
//     - addr (string): IPv4 address like "192.0.2.25".
//   Returns:
//     - (bool): true if the argument is a valid IPv4 Address.
//    See:
//     - https://datatracker.ietf.org/doc/html/rfc791
func IsIPv4Address(addr string) bool {
	if size := len(addr); size < 7 || size > 15 { return false }

	co, ci := 0, -1; for j := 0; j < len(addr); j++ {
		// Check that each octed is between 0 and 255.
		if cv := addr[j]; cv < '0' || cv > '9' {
			// Is not a numeric character
			if cv != '.' || ci == -1 { return false }
			co +=  1
			ci  = -1

		} else {
			// Is a numeric character
			if ci == -1 { ci = int(cv - '0'); continue }

			ci = (ci * 10) + int(cv - '0')
			if ci > 255 { return false }
		}
	}
	return co == 3 && ci > -1
}

// FindIPv4Address finds IPv4 addresses from the given string.
//   Arguments:
//     - text (string): String including an IPv4 address.
//   Returns:
//     - ([]string): List of IPv4 addresses found and picked from the argument.
func FindIPv4Address(text string) []string {
	if len(text) < 7 { return []string{} }

	// Rewrite: "mx.example.jp[192.0.2.1]" => "mx.example.jp 192.0.2.1"
	ipv4a := make([]string, 0, 1); cw := -1; for j := 0; j < len(text); j++ {
		switch text[j] {
			// Find a string including an IPv4 address
			case '(', ')', '[', ']', ',', ' ':
				if cw < 0 { continue }
				if cv := text[cw:j]; IsIPv4Address(cv) == true { ipv4a = append(ipv4a, cv) }
				cw = -1
			default: if cw < 0 { cw = j }
		}
	}
	if cw == -1 { return ipv4a }
	if cv := text[cw:]; IsIPv4Address(cv) == true { ipv4a = append(ipv4a, cv) }
	return ipv4a
}

