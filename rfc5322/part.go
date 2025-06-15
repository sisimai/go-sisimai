// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____ _________  ____  
// |  _ \|  ___/ ___| ___|___ /___ \|___ \ 
// | |_) | |_ | |   |___ \ |_ \ __) | __) |
// |  _ <|  _|| |___ ___) |__) / __/ / __/ 
// |_| \_\_|   \____|____/____/_____|_____|

package rfc5322
import "strings"
import "libsisimai.org/sisimai/v5/moji"

// Part splits the entire message body given as the 1st argument into error message lines and the
// original message part only include email headers.
//   Arguments:
//     - email (*string):  Entire message body
//     - cutby ([]string): String list of the message/rfc822 or the beginning of the original message part
//     - keeps (bool):     Flag for keeping strings after "\n\n"
//   Returns:
//     - ([2]string):      [2]string{"Error message lines", "The original message"}
func Part(email *string, cutby []string, keeps bool) [2]string {
	if email == nil || *email == "" || len(cutby) == 0 { return [2]string{} }

	positionor := -1 // A position of the boundary string
	formerbuff := strings.Builder{}; formerbuff.Grow(len(*email) / 2) // The error message part
	latterbuff := strings.Builder{}; latterbuff.Grow(len(*email) / 2) // The original message part

	for _, e := range cutby {
		// Find a boundary string(2nd argument)] from the 1st argument
		positionor = strings.Index(*email, e); if positionor > 0 { break }
	}

	if positionor > 0 {
		// There is the boundary string in the message body
		formerbuff.WriteString((*email)[:positionor])
		rfc822part := strings.Split((*email)[positionor:], "\n\n")

		for _, e := range rfc822part {
			// Find a part including "Received:", "From:" header
			if moji.ContainsAny(e, []string{"Received: ", "From: "}) == false { continue }
			latterbuff.WriteString(e); break
		}
		if latterbuff.Len() == 0 { latterbuff.WriteString((*email)[positionor:]) }

	} else {
		// Substitute the entire message to the former part when the boundary string is not included
		// in the 1st argument
		formerbuff.WriteString(*email)
	}

	latterpart := latterbuff.String(); if latterpart != "" {
		// Remove blank lines, the message body of the original message, and append "\n" at the end
		// of the original message headers
		// 1. Remove leading blank lines
		// 2. Remove text after the first blank line: \n\n
		// 3. Append "\n" at the end of test block when the last character is not "\n"
		for _, e := range strings.Split(latterpart, "") {
			// Remove leading blank lines
			if e == " " || e == "\n" || e == "\r" { continue }
			latterpart = e + moji.Select(latterpart + moji.RHS, e, "", 0)
			break
		}

		if keeps == true && strings.Contains(latterpart, "\n\n") {
			// Remove text after the first blank line when "keeps" is true
			latterpart = moji.Select(moji.LHS + latterpart, "", "\n\n", 0) + "\n"
		}

		// Append "\n" at the end of the original message
		if strings.HasSuffix(latterpart, "\n") == false { latterpart += "\n" }
	}
	return [2]string{formerbuff.String(), latterpart}
}

