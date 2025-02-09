// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package status

//                _           __   _        _             
//  ___ _ __ ___ | |_ _ __   / /__| |_ __ _| |_ _   _ ___ 
// / __| '_ ` _ \| __| '_ \ / / __| __/ _` | __| | | / __|
// \__ \ | | | | | |_| |_) / /\__ \ || (_| | |_| |_| \__ \
// |___/_| |_| |_|\__| .__/_/ |___/\__\__,_|\__|\__,_|___/
//                   |_|                                  
import "strings"
import "strconv"

// Test() checks whether a status code is a valid code or not
func Test(argv1 string) bool {
	// @param    string argv1  Status code(DSN)
	// @return   bool          false = Invalid status code, true = Valid status code
	if len(argv1) < 5 || len(argv1) > 7 { return false }

	token := []int16{} // Each digit like [5,7,26] converted from "5.7.26"
	for _, e := range strings.Split(argv1, ".") {
		digit, nyaan := strconv.Atoi(e); if nyaan != nil { break }
		token = append(token, int16(digit))
	}

	if len(token) != 3 { return false } // The number of elements should be 3 like [5,1,1]
	if token[0]    < 2 { return false } // Status: 1.x.y does not exist
	if token[0]   == 3 { return false } // Status: 3.x.y does not exist
	if token[0]    > 5 { return false } // Status: 6.x.y does not exist
	if token[1]    < 0 { return false }
	if token[1]    > 7 { return false }
	if token[2]    < 0 { return false }
	return true
}

