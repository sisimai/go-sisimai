// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package status

//                _           __   _        _             
//  ___ _ __ ___ | |_ _ __   / /__| |_ __ _| |_ _   _ ___ 
// / __| '_ ` _ \| __| '_ \ / / __| __/ _` | __| | | / __|
// \__ \ | | | | | |_| |_) / /\__ \ || (_| | |_| |_| \__ \
// |___/_| |_| |_|\__| .__/_/ |___/\__\__,_|\__|\__,_|___/
//                   |_|                                  
import "strings"

// Prefer() returns the preferred value selected from the arguments
func Prefer(argv0, argv1, argv2 string) string {
	// @param    string argv0  The value of Status: field
	// @param    string argv1  The delivery status value picked from the error message
	// @param    string argv2  The value of An SMTP Reply Code
	// @return   String        The preferred value
	if len(argv0) < 1 { return argv1 }
	if len(argv1) < 1 { return argv0 }

	statuscode := argv0; if len(statuscode) < 5 { return argv1     }
	codeinmesg := argv1; if len(codeinmesg) < 5 { return argv0     }
	esmtpreply := argv2; if len(esmtpreply) < 1 { esmtpreply = "0" }
	the1stchar := [3]byte{
		[]byte(statuscode[0:1])[0], // argv0: The "Status:" field
		[]byte(codeinmesg[0:1])[0], // argv1: The delivery status value in the error message
		[]byte(esmtpreply[0:1])[0], // argv2: SMTP Reply code
	}

	if the1stchar[2] > 0 && the1stchar[0] != the1stchar[1] {
		// There is the 3rd argument (an SMTP Reply Code)
		// Returns the value of argv0 or argv1 which begins with the 1st character of argv2
		if the1stchar[2] == the1stchar[0] { return statuscode }
		if the1stchar[2] == the1stchar[1] { return codeinmesg }
	}
	if statuscode == codeinmesg { return statuscode }

	// Find and check "X.Y.0" and "X.0.0" in the 1st and 2nd argument
	zeroindex1 := [2]int{strings.Index(statuscode, ".0"),   strings.Index(codeinmesg, ".0")}
	zeroindex2 := [2]int{strings.Index(statuscode, ".0.0"), strings.Index(codeinmesg, ".0.0")}

	if zeroindex2[0] > 0 {
		// The "Status:" field is "X.0.0"
		if zeroindex2[1] < 0 { return codeinmesg }
		return statuscode
	}

	if zeroindex1[0] > 0 {
		// The "Status:" field is "X.Y.0" or "X.0.Z"
		if zeroindex1[1] < 0 { return codeinmesg }
	}

	if zeroindex2[1] > 0                      { return statuscode } // The SMTP status code is "X.0.0"
	if statuscode == "4.4.7"                  { return codeinmesg } // "4.4.7" is an ambigous code
	if statuscode == "4.7.0"                  { return codeinmesg } // "4.7.0" indicates "too many errors"
	if strings.Index(statuscode, "5.3.") == 0 { return codeinmesg } // "5.3.Z" is a system error
	if strings.Index(statuscode, ".5.1") == 1 { return codeinmesg } // "X.5.1" indicates an invalid command
	if strings.Index(statuscode, ".5.2") == 1 { return codeinmesg } // "X.5.2" indicates a syntax error
	if strings.Index(statuscode, ".5.4") == 1 { return codeinmesg } // "X.5.4" indicates an invalid command argument
	if strings.Index(statuscode, ".5.5") == 1 { return codeinmesg } // "X.5.5" indicates a wrong protocol version

	if statuscode == "5.1.1" {
		// "5.1.1" is a code of "userunknown"
		if zeroindex1[1] > 0 || strings.Index(codeinmesg, "5.5.") == 0 { return statuscode }
		return codeinmesg

	} else if statuscode == "5.1.3" {
		// "5.1.3"
		if strings.Index(codeinmesg, "5.7.") == 0 { return codeinmesg }
	}
	return statuscode
}

