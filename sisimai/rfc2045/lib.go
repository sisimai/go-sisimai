// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc2045
import "strings"

// Parameter() finds a value of specified parameter name from Content-Type header.
func Parameter(argv0 string, argv1 string) string {
	// @param    [string] argv0  The value of Content-Type: header
	// @param    [string] argv1  Lower-cased attribute name of the parameter
    // @return   [string]        The value of the parameter
	if len(argv0) == 0 { return "" }

	parameterq := ""
	paramindex := 0

	if len(argv1) > 0 {
		// There is a parameter name in the second argument
		parameterq = strings.ToLower(argv1) + "="
		paramindex = strings.Index(argv0, parameterq)
	}
	if paramindex == -1 { return "" }

	// Find the value of the parameter name specified in $argv1
	foundtoken := strings.Split(argv0[paramindex + len(parameterq):], ";")[0]
	if argv1 != "boundary" { foundtoken = strings.ToLower(foundtoken) }
	foundtoken  = strings.Replace(foundtoken, `'`, "", -1)
	foundtoken  = strings.Replace(foundtoken, `"`, "", -1)

	return foundtoken
}

// Boundary() finds a boundary string from the value of Content-Type header.
func Boundary(argv0 string, start int) string {
	// @param    [string]  argv0 The value of Content-Type header
	// @param    [int]     start -1: boundary string itself
	//                            0: Start of boundary: "--boundary"
	//                            1: End of boundary" "--boundary--"
	// @return   [string]        Boundary string
	if len(argv0) == 0 { return "" }

	btext := Parameter(argv0, "boundary")
	if len(btext) == 0 { return "" }

	// Content-Type: multipart/mixed; boundary=Apple-Mail-5--931376066
	// Content-Type: multipart/report; report-type=delivery-status;
	//    boundary="n6H9lKZh014511.1247824040/mx.example.jp"
	if start > -1 { btext = "--" + btext }
	if start >  0 { btext += "--" }
	return btext
}

