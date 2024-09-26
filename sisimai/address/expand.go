// Copyright (C) 2020-2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address

//            _     _                   
//   __ _  __| | __| |_ __ ___  ___ ___ 
//  / _` |/ _` |/ _` | '__/ _ \/ __/ __|
// | (_| | (_| | (_| | | |  __/\__ \__ \
//  \__,_|\__,_|\__,_|_|  \___||___/___/
import "strings"

// ExpandVERP() gets the original recipient address from VERP
func ExpandVERP(email string) string {
	// @param    string email  VERP Address
	// @return   string        Email address
	if len(email) == 0 { return "" }
	local := strings.SplitN(email, "@", 2)[0]

	// bounce+neko=example.org@example.org => neko@example.org
	if strings.Index(local, "+") < 1 { return "" }
	if strings.Index(local, "=") < 1 { return "" }
	if strings.Index(local, "+") > len(local) - 1 { return "" }
	if strings.Index(local, "=") > len(local) - 1 { return "" }

	verp1 := strings.Replace(strings.SplitN(local, "+", 2)[1], "=", "@", 1)
	if IsEmailAddress(verp1) { return verp1 }
	return ""
}

// ExpandAlias() removes string from "+" to "@" at a local part
func ExpandAlias(email string) string {
	// @param    string email  Email alias string
	// @return   string        Expanded email address
	if len(email)                   == 0     { return "" }
	if IsEmailAddress(email)        == false { return "" }
	if strings.Contains(email, "+") == false { return "" }
	if strings.Index(email, "+")     < 1     { return "" }

	// neko+straycat@example.org => neko@example.org
	lpart := email[0:strings.Index(email, "+")]
	dpart := strings.SplitN(email, "@", 2)[1]
	return lpart + "@" + dpart
}

