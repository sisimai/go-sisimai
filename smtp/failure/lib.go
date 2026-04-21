// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                _           ____       _ _                
//  ___ _ __ ___ | |_ _ __   / / _| __ _(_) |_   _ _ __ ___ 
// / __| '_ ` _ \| __| '_ \ / / |_ / _` | | | | | | '__/ _ \
// \__ \ | | | | | |_| |_) / /|  _| (_| | | | |_| | | |  __/
// |___/_| |_| |_|\__| .__/_/ |_|  \__,_|_|_|\__,_|_|  \___|
//                   |_|                                    

// Package "smtp/failure" provides functions related to SMTP errors.
package failure
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/smtp/reply"
import "libsisimai.org/sisimai/v5/smtp/status"

// IsPermanent returns true if the given string indicates a permanent error.
//   Arguments:
//     - text (string): String including SMTP status code.
//   Returns:
//     - (bool): true if it indicates permanent error, false otherwise.
func IsPermanent(text string) bool {
	if text == "" { return false }

	cv := status.Find(text, ""); if cv == "" { cv = reply.Find(text, "") }
	if strings.HasPrefix(cv, "5") || strings.Contains(strings.ToLower(text), " permanent ") { return true }
	return false
}

// IsTemporary returns true if the given string indicates a temporary error.
//   Arguments:
//     - text (string): String including SMTP status code.
//   Returns:
//     - (bool): true if it indicates temporary error, false otherwise.
func IsTemporary(text string) bool {
	if text == "" { return false }

	cv := status.Find(text, ""); if cv== "" { cv = reply.Find(text, "") }
	cc := strings.ToLower(text)

	if strings.HasPrefix(cv, "4")          { return true }
	if strings.Contains(cc, " temporar")   { return true }
	if strings.Contains(cc, " persistent") { return true }
	return false
}

// IsHardBounce checks the reason sisimai detected is a hard bounce or not.
//   Arguments:
//     - name (string): The bounce reason sisimai detected.
//     - code (string): String including SMTP status code.
//   Returns:
//     - (bool): true if it indicates hard bounce, false otherwise.
func IsHardBounce (name, code string) bool {
	if name == eb.Re___0 || name == eb.Re___1 || name == ""        { return false }
	if name == eb.ReSENT || name == eb.ReFEED || name == eb.ReAWAY { return false }
	if name == eb.ReMOVE || name == eb.ReUSER || name == eb.ReHOST { return true  }
	if name != eb.Re00MX                                           { return false }
	if code == ""                                                  { return true  }

	// Check the 2nd argument(a status code or a reply code)
	//   - The SMTP status code or the SMTP reply code starts with "5"
	//   - Deal as a hard bounce when the error message does not indicate a temporary error 
	cv := status.Find(code, ""); if cv == "" { cv = reply.Find(code, "") }
	if strings.HasPrefix(cv, "5") || IsTemporary(code) == false { return true }
	return false
}

// IsSoftBounce checks the reason sisimai detected is a soft bounce or not.
//   Arguments:
//     - name (string): The bounce reason sisimai detected.
//     - code (string): String including SMTP status code.
//   Returns:
//     - (bool): true if it indicates soft bounce, false otherwise.
func IsSoftBounce (name, code string) bool {
	if name == eb.ReSENT || name == eb.ReFEED || name == eb.ReAWAY { return false }
	if name == eb.ReMOVE || name == eb.ReUSER || name == eb.ReHOST { return false }
	if name == eb.Re___0 || name == eb.Re___1                      { return true  }
	if name != eb.Re00MX                                           { return true  }
	if code == ""                                                  { return false }

	// NotAccept: 5xx => hard bounce, 4xx => soft bounce
	// Check the 2nd argument(a status code or a reply code)
	cv := status.Find(code, ""); if cv == "" { cv = reply.Find(code, "") }
	if strings.HasPrefix(cv, "4") { return true }
	return false
}

