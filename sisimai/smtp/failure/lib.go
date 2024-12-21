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

	statuscode := status.Find(argv1, "");  if statuscode == "" { statuscode = reply.Find(argv1, "") }
	if strings.HasPrefix(statuscode, "5")                      { return true }
	if strings.Contains(strings.ToLower(argv1), " permanent ") { return true }
	return false
}

// IsTemporary() returns true if the given string indicates a temporary error
func IsTemporary(argv1 string) bool {
	// @param   string argv1  String including SMTP Status code
	// @return  bool          true(temporary error), false(is not a temporary error)
	if len(argv1) == 0 { return false }

	statuscode := status.Find(argv1, ""); if statuscode == "" { statuscode = reply.Find(argv1, "") }
	issuedcode := strings.ToLower(argv1)

	if strings.HasPrefix(statuscode, "4")          { return true }
	if strings.Contains(issuedcode, " temporar")   { return true }
	if strings.Contains(issuedcode, " persistent") { return true }
	return false
}

// IsHardBounce() checks the reason sisimai detected is a hard bounce or not
func IsHardBounce (argv1, argv2 string) bool {
	// @param   string argv1  The bounce reason sisimai detected
	// @param   string argv2  String including SMTP Status code
	// @return  bool          true: is a hard bounce
	if argv1 == "undefined" || argv1 == "onhold"      || argv1 == ""            { return false }
	if argv1 == "deliverd"  || argv1 == "feedback"    || argv1 == "vacation"    { return false }
	if argv1 == "hasmoved"  || argv1 == "userunknown" || argv1 == "hostunknown" { return true  }
	if argv1 != "notaccept"                                                     { return false }
	if argv2 == ""                                                              { return true  }

	// Check the 2nd argument(a status code or a reply code)
	//   - The SMTP status code or the SMTP reply code starts with "5"
	//   - Deal as a hard bounce when the error message does not indicate a temporary error 
	cv := status.Find(argv2, ""); if cv == "" { cv = reply.Find(argv2, "") }
	if strings.HasPrefix(cv, "5") || IsTemporary(argv2) == false { return true }
	return false
}

// IsSoftBounce() checks the reason sisimai detected is a soft bounce or not
func IsSoftBounce (argv1, argv2 string) bool {
	// @param   string argv1  The bounce reason sisimai detected
	// @param   string argv2  String including SMTP Status code
	// @return  bool          true: is a soft bounce
	if argv1 == "deliverd"  || argv1 == "feedback"    || argv1 == "vacation"    { return false }
	if argv1 == "hasmoved"  || argv1 == "userunknown" || argv1 == "hostunknown" { return false }
	if argv1 == "undefined" || argv1 == "onhold"                                { return true  }
	if argv1 != "notaccept"                                                     { return true  }
	if argv2 == ""                                                              { return false }

	// NotAccept: 5xx => hard bounce, 4xx => soft bounce
	// Check the 2nd argument(a status code or a reply code)
	cv := status.Find(argv2, ""); if cv == "" { cv = reply.Find(argv2, "") }
	if strings.HasPrefix(cv, "4") { return true }
	return false
}

