// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package command

//                _           __                                            _ 
//  ___ _ __ ___ | |_ _ __   / /__ ___  _ __ ___  _ __ ___   __ _ _ __   __| |
// / __| '_ ` _ \| __| '_ \ / / __/ _ \| '_ ` _ \| '_ ` _ \ / _` | '_ \ / _` |
// \__ \ | | | | | |_| |_) / / (_| (_) | | | | | | | | | | | (_| | | | | (_| |
// |___/_| |_| |_|\__| .__/_/ \___\___/|_| |_| |_|_| |_| |_|\__,_|_| |_|\__,_|
//                   |_|                                                      
import sisimoji "sisimai/string"

// Test() checks that an SMTP command in the argument is valid or not
func Test(argv0 string) bool {
	// @param    string argv0  An SMTP command
	// @return   bool          false: Is not a valid SMTP command
	//                         true:  Is a valid SMTP command
	// @since v5.2.0
	if len(argv0) < 4                          { return false }
	if sisimoji.ContainsAny(argv0, Availables) { return true  }
	return false
}

