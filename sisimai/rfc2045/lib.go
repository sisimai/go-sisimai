// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045

//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 
import "strings"

// Parameter() finds a value of specified parameter name from Content-Type header.
func Parameter(argv0 string, argv1 string) string {
	// @param    string argv0  The value of Content-Type: header
	// @param    string argv1  Lower-cased attribute name of the parameter
    // @return   string        The value of the parameter
	if argv0 == "" { return "" }
	cv := ""
	ci := 0

	if len(argv1) > 0 {
		// There is a parameter name in the second argument
		cv = strings.ToLower(argv1) + "="
		ci = strings.Index(strings.ToLower(argv0), cv); if ci == -1 { return "" }
	}

	// Find the value of the parameter name specified in "argv1"
	cf := strings.Split(argv0[ci + len(cv):], ";")[0]; if argv1 != "boundary" { cf = strings.ToLower(cf) }
	cf  = strings.Replace(cf, `'`, "", -1)
	cf  = strings.Replace(cf, `"`, "", -1)

	return cf
}

// CharacterSet() returns "ISO-2022-JP" as a character set name from "=?ISO-2022-JP?B?...?="
func CharacterSet(argv0 string) string {
	// @param    string argv0  Base64 or Quoted-Printable encoded text
	// @return   string        A character set name like "iso-2022-jp"
	if strings.HasPrefix(argv0, "=?") == false { return "" }
	if strings.HasSuffix(argv0, "?=") == false { return "" }

	argv1 := strings.ToUpper(argv0)
	index := strings.Index(argv1, "?B?"); if index < 0 { index = strings.Index(argv1, "?Q?") }

	if index < 0 { return "" }
	return argv1[2:index]
}

// Boundary() finds a boundary string from the value of Content-Type header.
func Boundary(argv0 string, start int) string {
	// @param    string  argv0    The value of Content-Type header
	// @param    int     start    -1: boundary string itself
	//                             0: Start of boundary: "--boundary"
	//                             1: End of boundary" "--boundary--"
	// @return   string            Boundary string
	if argv0 == "" { return "" }; btext := Parameter(argv0, "boundary")
	if btext == "" { return "" }

	// Content-Type: multipart/mixed; boundary=Apple-Mail-5--931376066
	// Content-Type: multipart/report; report-type=delivery-status;
	//    boundary="n6H9lKZh014511.1247824040/mx.example.jp"
	if start > -1 { btext  = "--" + btext }
	if start >  0 { btext += "--" }
	return btext
}

