// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _           __  __                    _ 
// | | | | __ _ ___|  \/  | _____   _____  __| |
// | |_| |/ _` / __| |\/| |/ _ \ \ / / _ \/ _` |
// |  _  | (_| \__ \ |  | | (_) \ V /  __/ (_| |
// |_| |_|\__,_|___/_|  |_|\___/ \_/ \___|\__,_|
import "strings"
import "sisimai/sis"

func init() {
	// Try to match that the given text and message patterns
	Match["HasMoved"] = func(argv1 string) bool {
		// @param    string argv1 String to be matched with text patterns
		// @return   bool         true: Matched, false: did not match
		index := []string{" has been replaced by "}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "hasmoved" or not
	Truth["HasMoved"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is hasmoved, false: is not hasmoved
		if fo.Reason == "hasmoved" { return true }
		return Match["HasMoved"](strings.ToLower(fo.DiagnosticCode))
	}
}

