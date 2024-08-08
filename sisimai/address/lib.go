// Copyright (C) 2020-2021 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package address
import "fmt"
import "strings"

// Undisclosed() returns a pseudo recipient address or a pseudo sender address
func Undisclosed(a string) string {
	// @param    [string] a   Address type: 'r' or 's'
	// @return   [string]     Pseudo recipient address or sender address when the "a" is neither "r" nor "s"
	addr := ""
	if a == "s" { addr = "sender"    }
	if a == "r" { addr = "recipient" }
	return fmt.Sprintf("undisclosed-%s-in-headers@libsisimai.org.invalid", addr)
}

// S3S4() runs like ruleset 3,4 of sendmail.cf
func S3S4(argv0 string) string {
	// @param    [string] input  Text including an email address
	// @return   [string]        Email address without comment, brackets
	if len(argv0) == 0 { return "" }

	list := Find(argv0)
	if len(list) == 0 { return argv0 }
	return list[0]
}

// Final() returns a string processed by Ruleset 4 in sendmail.cf
func Final(argv0 string) string {
	// @param    [string] argv0  String including an email address like "<neko@nyaan.jp>"
	// @return   [string]        String without angle brackets: "neko@nyaan.jp"
	for strings.HasPrefix(argv0, "<") { argv0 = strings.Trim(argv0, "<") }
	for strings.HasSuffix(argv0, ">") { argv0 = strings.Trim(argv0, ">") }

	atmark := strings.LastIndex(argv0, "@")
	useris := argv0[0:atmark]
	hostis := argv0[atmark+1:]

	if IsQuotedAddress(argv0) == false {
		// Remove all the angle brackets from the local part
		useris = strings.ReplaceAll(useris, "<", "")
		useris = strings.ReplaceAll(useris, ">", "")
	}
	// Remove all the angle brackets from the domain part
	hostis = strings.ReplaceAll(hostis, "<", "")
	hostis = strings.ReplaceAll(hostis, ">", "")
	return useris + "@" + hostis
}

// ExpandVERP() gets the original recipient address from VERP
func ExpandVERP(email string) string {
	// @param    [string] email  VERP Address
	// @return   [string]        Email address
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
	// @param    [string] email  Email alias string
	// @return   [string]        Expanded email address
	if len(email) == 0 { return "" }
	if !IsEmailAddress(email) { return "" }
	if !strings.Contains(email, "+") { return "" }
	if strings.Index(email, "+") < 1 { return "" }

	// neko+straycat@example.org => neko@example.org
	lpart := email[0:strings.Index(email, "+")]
	dpart := strings.SplitN(email, "@", 2)[1]
	return lpart + "@" + dpart
}

