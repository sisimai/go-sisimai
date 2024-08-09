// Copyright (C) 2020-2021,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reply
import "fmt"
import "strings"

// Find() returns an SMTP reply code found from the given string
func Find(argv1 string, argv2 string) string {
	// @param    string argv1  String including SMTP reply code like 550
	// @param    string argv2  Status code like 5.1.1 or 2 or 4 or 5
	// @return   string        SMTP reply code or empty if the first argument did not include SMTP Reply Code value
	if len(argv1) < 3                                     { return "" }
	if strings.Contains(strings.ToUpper(argv1), "X-UNIX") { return "" }
	if len(argv2) == 0 { argv2 = "0" }

	replycode2 := []uint16{211, 214, 220, 221, 235, 250, 251, 252, 253, 354}
	replycode4 := []uint16{421, 450, 451, 452, 422, 430, 432, 453, 454, 455, 456, 458, 459}
	replycode5 := []uint16{
		550, 552, 553, 551, 521, 525, 502, 520, 523, 524, 530, 533, 534, 535, 538, 551, 555, 556,
		554, 557, 500, 501, 502, 503, 504,
	}
	codeofsmtp := map[string][]uint16{"2": replycode2, "4": replycode4, "5": replycode5}
	statuscode := argv2[0:1]
	replycodes := []uint16{}
	esmtperror := fmt.Sprintf(" %s ", argv1)
	esmtpreply := ""        // SMTP Reply Code
	replyindex := -1        // A position of SMTP reply code found by the strings.Index()
	formerchar := uint8(0)  // A character that is one character before the SMTP reply code
	latterchar := uint8(0)  // A character that is one character  after the SMTP reply code

	if statuscode == "2" || statuscode == "4" || statuscode == "5" {
		// The first character of the 2nd argument is 2 or 4 or 5
		replycodes = codeofsmtp[statuscode]

	} else {
		// The first character of the 2nd argument is 0 or other values
		// TODO: use "slices" package and slices.Concat() avaialble from Go 1.22
		//       https://pkg.go.dev/slices@master
		replycodes = append(replycodes, codeofsmtp["2"]...)
		replycodes = append(replycodes, codeofsmtp["4"]...)
		replycodes = append(replycodes, codeofsmtp["5"]...)
	}

	for _, e := range replycodes {
		// Try to find an SMTP Reply Code from the given string
		f := string(e)
		replyindex = strings.Index(esmtperror, f); if replyindex < 0 { continue }
		formerchar = []byte(esmtperror[replyindex - 1:replyindex])[0]
		latterchar = []byte(esmtperror[replyindex + 3:replyindex + 4])[0]

		if formerchar > 45 && formerchar < 58 { continue }
		if latterchar > 45 && latterchar < 58 { continue }
		esmtpreply = f
		break
	}
	return esmtpreply
}

