// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package failure
//                _           ____       _ _                
//  ___ _ __ ___ | |_ _ __   / / _| __ _(_) |_   _ _ __ ___ 
// / __| '_ ` _ \| __| '_ \ / / |_ / _` | | | | | | '__/ _ \
// \__ \ | | | | | |_| |_) / /|  _| (_| | | | |_| | | |  __/
// |___/_| |_| |_|\__| .__/_/ |_|  \__,_|_|_|\__,_|_|  \___|
//                   |_|                                    
import "strings"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"

// IsPermanent() returns true if the given string indicates a permanent error
func IsPermanent(argv1 string) bool {
	// @param   string argv1  String including SMTP Status code
	// @return  bool          true(permanet error), false(is not a permanent error)
	if len(argv1) == 0 { return false }

	statuscode := status.Find(argv1, "")
	if statuscode == "" { statuscode = reply.Find(argv1, "") }

	if strings.HasPrefix(statuscode, "5")                      { return true }
	if strings.Contains(strings.ToLower(argv1), " permanent ") { return true }
	return false
}

// IsTemporary() returns true if the given string indicates a temporary error
func IsTemporary(argv1 string) bool {
	// @param   string argv1  String including SMTP Status code
	// @return  bool          true(temporary error), false(is not a temporary error)
	if len(argv1) == 0 { return false }

	statuscode := status.Find(argv1, "")
	if statuscode == "" { statuscode = reply.Find(argv1, "") }

	if strings.HasPrefix(statuscode, "4")                      { return true }
	if strings.Contains(strings.ToLower(argv1), " temporar")   { return true }
	if strings.Contains(strings.ToLower(argv1), " persistent") { return true }
	return false
}

// IsHardBounce() checks the reason sisimai detected is a hard bounce or not
func isHardBounce (argv1, argv2 string) bool {
	// @param   string argv1  The bounce reason sisimai detected
	// @param   string argv2  String including SMTP Status code
	// @return  bool          true: is a hard bounce
	if argv1 == "undefined" || argv1 == "onhold"      || argv1 == ""            { return false }
	if argv1 == "deliverd"  || argv1 == "feedback"    || argv1 == "vacation"    { return false }
	if argv1 == "hasmoved"  || argv1 == "userunknown" || argv1 == "hostunknown" { return true  }
	if argv1 != "notaccept"                                                     { return false }

	// NotAccept: 5xx => hard bounce, 4xx => soft bounce
	hardbounce := false
	if len(argv2) > 0 {
		// Check the 2nd argument(a status code or a reply code)
		statuscode := status.Find(argv2, "")
		if statuscode == "" { statuscode = reply.Find(argv2, "") }
		if strings.HasPrefix(statuscode, "5") { hardbounce = true }

	} else {
		// Deal "NotAccept" as a hard bounce when the 2nd argument is empty
		hardbounce = true
	}
	return hardbounce
}

// IsSoftBounce() checks the reason sisimai detected is a soft bounce or not
func isSoftBounce (argv1, argv2 string) bool {
	// @param   string argv1  The bounce reason sisimai detected
	// @param   string argv2  String including SMTP Status code
	// @return  bool          true: is a soft bounce
	if argv1 == "deliverd"  || argv1 == "feedback"    || argv1 == "vacation"    { return false }
	if argv1 == "hasmoved"  || argv1 == "userunknown" || argv1 == "hostunknown" { return false }
	if argv1 == "undefined" || argv1 == "onhold"                                { return true  }
	if argv1 != "notaccept"                                                     { return true  }

	// NotAccept: 5xx => hard bounce, 4xx => soft bounce
	softbounce := false
	if len(argv2) > 0 {
		// Check the 2nd argument(a status code or a reply code)
		statuscode := status.Find(argv2, "")
		if statuscode == "" { statuscode = reply.Find(argv2, "") }
		if strings.HasPrefix(statuscode, "4") { softbounce = true }
	}
	return softbounce
}

