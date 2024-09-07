// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

// __     __              _   _             
// \ \   / /_ _  ___ __ _| |_(_) ___  _ __  
//  \ \ / / _` |/ __/ _` | __| |/ _ \| '_ \ 
//   \ V / (_| | (_| (_| | |_| | (_) | | | |
//    \_/ \__,_|\___\__,_|\__|_|\___/|_| |_|
//                                          
import "strings"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["Vacation"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			"i am away on vacation",
			"i am away until",
			"i am out of the office",
			"i will be traveling for work on",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "vacation" or not
	Truth["Vacation"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is vacation, false: is not vacation
		return false
	}
}

