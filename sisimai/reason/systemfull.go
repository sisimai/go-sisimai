// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____            _                 _____      _ _ 
// / ___| _   _ ___| |_ ___ _ __ ___ |  ___|   _| | |
// \___ \| | | / __| __/ _ \ '_ ` _ \| |_ | | | | | |
//  ___) | |_| \__ \ ||  __/ | | | | |  _|| |_| | | |
// |____/ \__, |___/\__\___|_| |_| |_|_|   \__,_|_|_|
//        |___/                                      
import "strings"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["SystemFull"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			"mail system full",
			"requested mail action aborted: exceeded storage allocation", // MS Exchange
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "systemfull" or not
	Truth["SystemFull"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is systemfull, false: is not systemfull
		return false
	}
}

