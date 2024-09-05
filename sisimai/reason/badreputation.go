// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____            _ ____                  _        _   _             
// | __ )  __ _  __| |  _ \ ___ _ __  _   _| |_ __ _| |_(_) ___  _ __  
// |  _ \ / _` |/ _` | |_) / _ \ '_ \| | | | __/ _` | __| |/ _ \| '_ \ 
// | |_) | (_| | (_| |  _ <  __/ |_) | |_| | || (_| | |_| | (_) | | | |
// |____/ \__,_|\__,_|_| \_\___| .__/ \__,_|\__\__,_|\__|_|\___/|_| |_|
//                             |_|                                     
import "strings"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["BadReputation"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{
			"a poor email reputation score",
			"has been temporarily rate limited due to ip reputation",
			"ip/domain reputation problems",
			"likely suspicious due to the very low reputation",
			"temporarily deferred due to unexpected volume or user complaints", // Yahoo Inc.
			"the sending mta's poor reputation",
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "badreputation" or not
	Truth["BadReputation"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is badreputation, false: is not badreputation
		if fo.Reason == "badreputation" { return true }
		return Match["BadReputation"](strings.ToLower(fo.DiagnosticCode))
	}
}

