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
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["Vacation"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
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
	ProbesInto["Vacation"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is vacation, false: is not vacation
		return false
	}
}

