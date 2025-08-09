// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 

// Package "rfc2045" provides functions for RFC2045; Multipurpose Internet Mail Extensions (MIME)
// Part One: Format of Internet Message Bodies. https://datatracker.ietf.org/doc/html/rfc2045
package rfc2045
import "strings"
import "libsisimai.org/sisimai/v5/moji"

// Parameter finds a value of specified parameter name from Content-Type header.
//   Arguments:
//     - ctype (string): Value of Content-Type: header.
//     - lower (string): Lower-cased attribute name of the parameter.
//   Returns:
//     - (string): The value of the parameter.
//   See:
//     - https://datatracker.ietf.org/doc/html/rfc2045
func Parameter(ctype string, lower string) string {
	if ctype == "" { return "" }

	cv := ""; ci := 0; if len(lower) > 0 {
		// There is a parameter name in the second argument
		cv = strings.ToLower(lower) + "="
		ci = strings.Index(strings.ToLower(ctype), cv); if ci == -1 { return "" }
	}

	// Find the value of the parameter name specified in "lower"
	cf := strings.Split(ctype[ci + len(cv):], ";")[0]; if lower != "boundary" { cf = strings.ToLower(cf) }
	for _, e := range []string{`'`, `"`} { cf = strings.ReplaceAll(cf, e, "") }

	return cf
}

// CharacterSet returns "ISO-2022-JP" as a character set name from "=?ISO-2022-JP?B?...?=".
//   Arguments:
//     - text (string): Base64 or Quoted-Printable encoded text.
//   Returns:
//     - (string): Character set name like "iso-2022-jp".
func CharacterSet(text string) string {
	if strings.HasPrefix(text, "=?") == false || strings.HasSuffix(text, "?=") == false { return "" }
	return moji.Select(strings.ToUpper(text), "=?", "?", 0)
}

// Boundary finds a boundary string from the value of Content-Type header.
//   Arguments:
//     - ctype (string): Value of Content-Type header.
//     - start (int): 
//        - -1: boundary string itself
//        -  0: Start of boundary: "--boundary"
//        -  1: End of boundary" "--boundary--"
//   Returns:
//     - (string): Boundary string.
func Boundary(ctype string, start int) string {
	if ctype == "" { return "" }; btext := Parameter(ctype, "boundary")
	if btext == "" { return "" }

	// Content-Type: multipart/mixed; boundary=Apple-Mail-5--931376066
	// Content-Type: multipart/report; report-type=delivery-status;
	//    boundary="n6H9lKZh014511.1247824040/mx.example.jp"
	if start > -1 { btext  = "--" + btext }
	if start >  0 { btext += "--" }
	return btext
}

