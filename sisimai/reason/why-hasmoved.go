// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _   _           __  __                    _ 
// | | | | __ _ ___|  \/  | _____   _____  __| |
// | |_| |/ _` / __| |\/| |/ _ \ \ / / _ \/ _` |
// |  _  | (_| \__ \ |  | | (_) \ V /  __/ (_| |
// |_| |_|\__,_|___/_|  |_|\___/ \_/ \___|\__,_|
import "strings"
import "libsisimai.org/sisimai/sis"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["HasMoved"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		if argv1 == "" { return false }
		index := []string{" has been replaced by "}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "hasmoved" or not
	ProbesInto["HasMoved"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is hasmoved, false: is not hasmoved
		if fo        == nil        { return false }
		if fo.Reason == "hasmoved" { return true  }
		return IncludedIn["HasMoved"](strings.ToLower(fo.DiagnosticCode))
	}
}

