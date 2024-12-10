// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reply

//                _           __              _       
//  ___ _ __ ___ | |_ _ __   / / __ ___ _ __ | |_   _ 
// / __| '_ ` _ \| __| '_ \ / / '__/ _ \ '_ \| | | | |
// \__ \ | | | | | |_| |_) / /| | |  __/ |_) | | |_| |
// |___/_| |_| |_|\__| .__/_/ |_|  \___| .__/|_|\__, |
//                   |_|               |_|      |___/ 
import "strings"
import sisimoji "sisimai/string"

// Find() returns an SMTP reply code found from the given string
func Find(argv1 string, argv2 string) string {
	// @param    string argv1  String including SMTP reply code like 550
	// @param    string argv2  Status code like 5.1.1 or 2 or 4 or 5
	// @return   string        SMTP reply code or empty if the first argument did not include SMTP Reply Code value
	if len(argv1) < 3                                     { return "" }
	if strings.Contains(strings.ToUpper(argv1), "X-UNIX") { return "" }
	if len(argv2) == 0 { argv2 = "0" }

	statuscode := argv2[0:1]
	replycodes := []string{}

	if statuscode == "2" || statuscode == "4" || statuscode == "5" {
		// The first character of the 2nd argument is 2 or 4 or 5
		replycodes = CodeOfSMTP[statuscode]

	} else {
		// The first character of the 2nd argument is 0 or other values
		// TODO: use "slices" package and slices.Concat() avaialble from Go 1.22
		//       https://pkg.go.dev/slices@master
		replycodes = append(replycodes, CodeOfSMTP["2"]...)
		replycodes = append(replycodes, CodeOfSMTP["4"]...)
		replycodes = append(replycodes, CodeOfSMTP["5"]...)
	}

	esmtperror := " " + argv1 + " "
	esmtpreply := ""
	for _, e := range replycodes {
		// Try to find an SMTP Reply Code from the given string
		//replyindex := strings.Index(esmtperror, e); if replyindex < 0 { continue }
		appearance := strings.Count(esmtperror, e); if appearance < 1 { continue }
		startingat := 1

		for j := 0; j < appearance; j++ {
			// Find all the reply codes in the error message
			replyindex := sisimoji.IndexOnTheWay(esmtperror, e, startingat); if replyindex < 0 { break }
			formerchar := []byte(esmtperror[replyindex - 1:replyindex])[0]
			latterchar := []byte(esmtperror[replyindex + 3:replyindex + 4])[0]

			if formerchar > 45 && formerchar < 58 { startingat += replyindex + 3; continue }
			if latterchar > 45 && latterchar < 58 { startingat += replyindex + 3; continue }
			esmtpreply = e
			break
		}
		if esmtpreply != "" { break }
	}
	return esmtpreply
}

