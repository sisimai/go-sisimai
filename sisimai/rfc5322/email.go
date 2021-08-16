// Copyright (C) 2020 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322
import "strings"
import "fmt"

// Received() convert Received headers to a structured data
func Received(argv0 string) []string {
	// @param    [string] argv0  Received header
	// @return   [[]string]      Received header as a structured data
	return []string {}
}

// Fillet() split given entire message body into error message lines and the original message part only
// include email headers.
func Fillet(mbody *string, bones []string) []string {
	// @param    [*string]  mbody  Entire message body
	// @param    [[]string] bones  String list of the message/rfc822 or the beginning of the original
	//                             message part
	// @return   [[]string]        { "Error message lines", "The original message" }
	if len(*mbody) == 0 { return nil }
	if len(bones)  == 0 { return nil }

	bone1 := ""
	for _, b := range bones {
		// Try to find the separator string
		if !strings.Contains(*mbody, b) { continue }
		bone1 = b
	}

	parts := strings.SplitN(*mbody, bone1, 2)
	if len(parts) == 1 { parts = append(parts, "") }
	if len(parts[1]) > 0 {
		// Remove blank lines, the message body of the original message, and append "\n" at the end
		// of the original message headers
		for _, e := range strings.Split(parts[1], "") {
			// Remove leading blank lines
			if e == " " || e == "\n" || e == "\r" { continue }
			p := strings.Index(parts[1], e)
			if p > 0 {
				// There is leading space characters at the head of parts[1]
				parts[1] = parts[1][p:len(parts[1])]
			}
			break
		}

		if strings.Contains(parts[1], "\n\n") {
			// Remove text after the first blank line
			parts[1] = parts[1][0:strings.Index(parts[1], "\n\n") + 1]
		}

		if !strings.HasSuffix(parts[1], "\n") {
			// Append "\n" at the end of the original message
			parts[1] += "\n"
		}
	}
		fmt.Printf("LENGTH-OF-EEEE = [%d]\n", len(parts[1]))
	return parts
}

