// Copyright (C) 2020 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045
import "strings"
import "mime"
import "fmt"
import "log"
import "mime/quotedprintable"
import "io/ioutil"

// IsEncoded() checks that the argument is MIME encoded string or not.
func IsEncoded(argv0 string) bool {
	// @param    string    argv0   String to be checked that it is MIME encoded or not
	// @return   bool              true: Not MIME encoded string
    //                             false: MIME encoded string
	argv0  = strings.ToUpper(argv0)
	match := false
	for {
		// =?UTF-8?B?44OL44Oj44O844Oz?=
		if !strings.Contains(argv0, "=?") { break } // Begins with "=?"
		if !strings.Contains(argv0, "?=") { break } // Ends with "?="
		if strings.Count(argv0, "?") != 4 { break } // "?" appears 4 times
		if len(argv0) < 8                 { break } // String length should be 8 or more
		if strings.Contains(argv0, "?B?") || strings.Contains(argv0, "?Q?") { match = true }
		break
	}
	return match
}

// DecodeH() decodes the value of email header which is MIME-Encoded string.
func DecodeH(argv0 *string) *string {
	// @param    *string    argvs  A pointer to MIME-Encoded text
	// @return   *string           MIME-Decoded text
	hasdecoded := ""
	stringlist := []string{}
	decodingif := new(mime.WordDecoder)
	characters := []string{".", "[", "]"}

	if strings.Contains(*argv0, " ") {
		// The argument string include 1 or more space characters
		stringlist = strings.Split(*argv0, " ")

	} else {
		// The argument string does not contain any space characters
		stringlist = append(stringlist, *argv0)
	}

	for index, value := range stringlist {
		// Check and decode each part of the string
		if IsEncoded(value) {
			if ! strings.HasPrefix(value, "=?") {
				// For example, "[=?UTF-8?B?...]"
				for _, c := range characters { value = strings.Replace(value, c + "=?", "=?", -1) }
			}

			if ! strings.HasSuffix(value, "?=") {
				// For example, "=?UTF-8?B?....?=."
				for _, c := range characters { value = strings.Replace(value, "?=" + c, "?=", -1) }
			}

			if de, oops := decodingif.DecodeHeader(value); oops == nil {
				// Successfully decoded
				if index > 0 { hasdecoded += " " }
				hasdecoded += de

			} else {
				// Failed to decode
				if index > 0 { hasdecoded += " " }
				hasdecoded += value
			}
		} else {
			// Is not MIME-Encoded text part
			if index > 0 { hasdecoded += " " }
			hasdecoded += value
		}
	}
	return &hasdecoded
}

// DecodeB() decodes Base64 encoded text.
func DecodeB(argv0 *string, argv1 string) *string {
	// @param    *string    argv0  Base64 Encoded text
	// @param    string     argv1  Character set name
	// @return   *string           MIME-Decoded text
	if len(*argv0) < 8 { return argv0 }
	if len(argv1) == 0 { argv1 = "utf-8" }

	decodingif := new(mime.WordDecoder)
	base64text := strings.TrimSpace(*argv0)
	base64text  = strings.Join(strings.Split(base64text, "\n"), "")
	base64text  = fmt.Sprintf("=?%s?B?%s?=", argv1, base64text)
	plainvalue := ""

	if plain, nyaan := decodingif.Decode(base64text); nyaan != nil {
		// Failed to decode the base64-encoded text
		log.Fatal(nyaan)

	} else {
		// Successfully decoded
		plainvalue = plain
	}
	return &plainvalue
}

// DecodeQ() decodes Quoted-Pritable encdoed text
func DecodeQ(argv0 *string) *string {
	// @param    *string    argv0 Quoted-Printable Encoded text
	// @return   *string          MIME-Decoded text
	readstring := strings.NewReader(*argv0)
	decodingif := quotedprintable.NewReader(readstring)
	plainvalue := ""

	if plain, nyaan := ioutil.ReadAll(decodingif); nyaan != nil {
		// Failed to decode the quoted-printable text
		log.Fatal(nyaan)

	} else {
		// Successfully decoded
		plainvalue = string(plain)
	}
	return &plainvalue
}

