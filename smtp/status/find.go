// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                _           __   _        _             
//  ___ _ __ ___ | |_ _ __   / /__| |_ __ _| |_ _   _ ___ 
// / __| '_ ` _ \| __| '_ \ / / __| __/ _` | __| | | / __|
// \__ \ | | | | | |_| |_) / /\__ \ || (_| | |_| |_| \__ \
// |___/_| |_| |_|\__| .__/_/ |___/\__\__,_|\__|\__,_|___/
//                   |_|                                  

package status
import "fmt"
import "sort"
import "strings"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc791"

// Find returns a delivery status code found from the given string.
//   Arguments:
//     - text (string): String including DSN; SMTP status code.
//     - code (string): SMTP Reply Code like "550" or the 1st digit of code such as "2", "4", or "5".
//   Returns:
//     - (string): SMTP status code found in the 1st argument.
func Find(text string, code string) string {
	if len(text) < 7 { return ""   }
	if len(code) < 1 { code = " " }

	eestatuses := make([]string, 0, 3)
	esmtperror := " " + text + "   " // Why 3 space characters? see https://github.com/sisimai/p5-sisimai/issues/574
	lookingfor := make(map[string]string, 10)
	indextable := make([]int, 0, 10)
	givenclass := code[0:1]; switch givenclass {
		case "2", "4", "5": eestatuses = append(eestatuses, givenclass + ".")
		default:            eestatuses = append(eestatuses, []string{"5.", "4.", "2."}...)
	}

	// Rewrite an IPv4 address in the given string(text) with '***.***.***.***'
	ip4address := rfc791.FindIPv4Address(esmtperror)
	for _, e := range ip4address { esmtperror = strings.ReplaceAll(esmtperror, e, "***.***.***.***") }
	for _, e := range eestatuses {
		// Count the number of "5.", "4.", and "2." in the error message
		p0, p1 := 0, 0
		for p0 > -1 {
			// Find all of the "5." and "4." string and store its postion
			p0 = moji.IndexOnTheWay(esmtperror, e, p1); if p0 < 0 { break }
			p1 = p0 + 5
			lookingfor[fmt.Sprintf("%04d", p0)] = e
			indextable = append(indextable, p0)
		}
	}
	if len(lookingfor) == 0 { return "" }

	statuscode := make([]string, 0, 2) // List of SMTP Status Code, Keep the order of appearances
	anotherone := ""                   // Alternative code
	stringsize := len(esmtperror)
	readbuffer := strings.Builder{}; readbuffer.Grow(5)

	sort.Slice(indextable, func(a, b int) bool { return indextable[a] < indextable[b] })
	for _, e := range indextable {
		// Try to find an SMTP Status Code from the given string
		cu := fmt.Sprintf("%04d", e)
		ci := moji.IndexOnTheWay(esmtperror, lookingfor[cu], e); if ci < 0 { continue }
		cx := [8]byte{0, 0, 0, 0, 0, 0, 0, 0}

		if stringsize > ci     { cx[0] = []byte(esmtperror[ci - 1:ci])[0]     } // [0] The previous character of the status
		if stringsize > ci + 3 { cx[1] = []byte(esmtperror[ci + 2:ci + 3])[0] } // [1] The value of the "Subject", "5.[7].261"
		if stringsize > ci + 4 { cx[2] = []byte(esmtperror[ci + 3:ci + 4])[0] } // [2] "." chacater, a separator of the Subject and the Detail

		if cx[0]  > 45 && cx[0]  <  58 { continue } // Previous character is a number
		if cx[0] == 86 || cx[0] == 118 { continue } // Avoid a version number("V" or "v")
		if cx[1]  < 48 || cx[1]  >  55 { continue } // The value of the subject is not a number(0-7)
		if cx[2] != 46                 { continue } // It is not a "." character: a separator

		if readbuffer.Len() > 0 { readbuffer.Reset() }
		readbuffer.WriteString(lookingfor[cu])
		readbuffer.WriteByte(cx[1])
		readbuffer.WriteByte('.')

		if stringsize > ci + 5 { cx[3] = []byte(esmtperror[ci + 4:ci + 5])[0] } // [3] The 1st digit of the detail
		if stringsize > ci + 6 { cx[4] = []byte(esmtperror[ci + 5:ci + 6])[0] } // [4] The 2nd digit of the detail
		if stringsize > ci + 7 { cx[5] = []byte(esmtperror[ci + 6:ci + 7])[0] } // [5] The 3rd digit of the detail
		if stringsize > ci + 8 { cx[6] = []byte(esmtperror[ci + 7:ci + 8])[0] } // [6] The next character

		if cx[3] < 48 || cx[3] > 57 { continue } // The 1st digit of the detail is not a number
		readbuffer.WriteByte(cx[3])

		if cv := readbuffer.String(); strings.Index(cv, ".0.0") == 1 || cv == "4.4.7" {
			// Find another status code except *.0.0, 4.4.7
			anotherone = cv; continue
		}

		// The 2nd digit of the detail is not a number
		if cx[4] < 48 || cx[4] > 57 { statuscode = append(statuscode, readbuffer.String()); continue }
		readbuffer.WriteByte(cx[4]) // The 2nd digit of the detail is a number

		// The 3rd digit of the detail is not a number
		if cx[5] < 48 || cx[5] > 57 { statuscode = append(statuscode, readbuffer.String()); continue }
		readbuffer.WriteByte(cx[5]) // The 3rd digit of the detail is a number

		if cx[6] > 47 && cx[6] < 58 { continue }
		statuscode = append(statuscode, readbuffer.String())
	}

	if len(anotherone) > 0 { statuscode = append(statuscode, anotherone) }
	if len(statuscode) < 1 { return "" }

	cv := ""; for j, e := range statuscode {
		// Select one from picked status codes
		if j == 0 { cv = e; continue }
		cv = Prefer(cv, e, "");
	}
	return cv
}

