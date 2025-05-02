// Copyright (C) 2020-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/

package address
import "strings"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc5322"

// ExpandVERP gets the original recipient address from a VERP address.
//   Arguments:
//     - email (string): VERP address such as "bounce+neko=example.jp@example.org"
//   Returns:
//     - (string):       Expanded email address such as "neko@example.jp"
func ExpandVERP(email string) string {
	if email                          == ""   { return "" }
	if strings.IndexByte(email, '@')  == -1   { return "" }
	if rfc5322.IsQuotedAddress(email) == true { return "" } // Do not expand "neko+cat=example.jp"@example.org

	cv := moji.Select(email, "+", "@", 0);  if cv == ""                   { return "" }
	cw := strings.Replace(cv, "=", "@", 1); if rfc5322.IsEmailAddress(cw) { return cw }
	return ""
}

// ExpandAlias removes string from "+" to "@" at a local part.
//   Arguments:
//     - email (string): Email alias such as "neko+straycat@example.jp"
//   Returns:
//     - (string):       Email address "neko@example.jp"
func ExpandAlias(email string) string {
	if email == "" || strings.IndexByte(email, '+') < 1 { return "" }
	if rfc5322.IsEmailAddress(email)  == false          { return "" }
	if rfc5322.IsQuotedAddress(email) == true           { return "" } // Do not expand "neko+cat"@example.org
	return moji.Select(moji.LHS + email, "", "+", 0) + "@" + moji.Select(email + moji.RHS, "@", "", 1)
}

