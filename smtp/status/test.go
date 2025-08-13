// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                _           __   _        _             
//  ___ _ __ ___ | |_ _ __   / /__| |_ __ _| |_ _   _ ___ 
// / __| '_ ` _ \| __| '_ \ / / __| __/ _` | __| | | / __|
// \__ \ | | | | | |_| |_) / /\__ \ || (_| | |_| |_| \__ \
// |___/_| |_| |_|\__| .__/_/ |___/\__\__,_|\__|\__,_|___/
//                   |_|                                  

package status
import "strings"
import "strconv"

// Test checks whether an SMTP status code is a valid code or not.
//   Arguments:
//     - code (string): SMTP status code to be checked.
//   Returns:
//     - (bool): true if the argument is a valid SMTP status code.
func Test(code string) bool {
	if len(code) < 5 || len(code) > 7 { return false }

	token := make([]int16, 0, 3) // Each digit like [5,7,26] converted from "5.7.26"
	for _, e := range strings.Split(code, ".") {
		digit, nyaan := strconv.Atoi(e); if nyaan == nil { token = append(token, int16(digit)) }
	}
	if len(token) != 3 { return false } // The number of elements should be 3 like [5,1,1]

	if token[0] < 2 || token[0] == 3 || token[0] > 5 { return false } // Status: [136].y.z does not exist
	if token[1] < 0 || token[1]  > 7 || token[2] < 0 { return false }
	return true
}

