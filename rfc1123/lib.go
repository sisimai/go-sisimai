// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _ _ ____  _____ 
// |  _ \|  ___/ ___/ / |___ \|___ / 
// | |_) | |_ | |   | | | __) | |_ \ 
// |  _ <|  _|| |___| | |/ __/ ___) |
// |_| \_\_|   \____|_|_|_____|____/ 

// package "rfc1123" provides functions related to Internet hosts described in RFC1123
// https://datatracker.ietf.org/doc/html/rfc1123
package rfc1123
import "strings"
import "libsisimai.org/sisimai/rfc791"
import sisimoji "libsisimai.org/sisimai/string"

var Sandwiched = [][]string{
	// (Postfix) postfix/src/smtp/smtp_proto.c: "host %s said: %s (in reply to %s)",
	// - <kijitora@example.com>: host re2.example.com[198.51.100.2] said: 550 ...
	// - <kijitora@example.org>: host r2.example.org[198.51.100.18] refused to talk to me:
	[]string{"host ", " said: "},
	[]string{"host ", " talk to me: "},
	[]string{"while talking to ", ":"}, // (Sendmail) ... while talking to mx.bouncehammer.jp.:
	[]string{"host ", " ["},            // (Exim) host mx.example.jp [192.0.2.20]: 550 5.7.0 
	[]string{" by ", ". ["},            // (Gmail) ...for the recipient domain example.jp by mx.example.jp. [192.0.2.1].

	// (MailFoundry)
	// - Delivery failed for the following reason: Server mx22.example.org[192.0.2.222] failed with: 550...
	// - Delivery failed for the following reason: mail.example.org[192.0.2.222] responded with failure: 552..
	[]string{"delivery failed for the following reason: ", " with"},
	[]string{"remote system: ", "("}, // (MessagingServer) Remote system: dns;mx.example.net (mx. -- 
	[]string{"smtp server <", ">"},   // (X6) SMTP Server <smtpd.libsisimai.org> rejected recipient ...
	[]string{"-mta: ", ">"},          // (MailMarshal) Reporting-MTA:      <rr1.example.com>
	[]string{" : ", "["},             // (SendGrid) cat:000000:<cat@example.jp> : 192.0.2.1 : mx.example.jp:[192.0.2.2]...
}
var StartAfter = []string{
	"generating server: ",   // (Exchange2007) Generating server: mta4.example.org
	"serveur de g",          // fr-FR/Serveur de g辿n辿ration
	"server di generazione", // it-CH
	"genererande server",    // sv-SE
}
var ExistUntil = []string{
	" did not like our ",  // (Dragonfly) mail-inbound.libsisimai.net [192.0.2.25] did not like our DATA: ...
}
var Prefix0x32 = []string{"(", "[", "<"}
var Suffix0x32 = []string{")", "]", ">", ":", ";"}

// IsInternetHost() returns "true" when the given string is a valid Internet hostname
func IsInternetHost(argv1 string) bool {
	// @param    string argv1  Hostname
	// @return   bool          true:  is a valid Internet hostname
	//                         false: is not a valid Internet hostname
	// @see https://datatracker.ietf.org/doc/html/rfc1123
	if len(argv1) < 4 || len(argv1) > 255 { return false }

	// Deal "localhost", "localhost6" as a valid hostname
	if argv1 == "localhost" || argv1 == "localhost6"    { return true  }
	if strings.IndexByte(argv1, '.') == -1              { return false }
	if strings.Contains(argv1, "..") == true            { return false }
	if sisimoji.HasPrefixAny(argv1, []string{".", "-"}) { return false }
	if strings.HasSuffix(argv1, "-") == true            { return false }

	// Allow the hostname starting with A-Label: "xn--" of IDN(Internationalized Domain Name)
	if strings.Contains(argv1, "--") == true && strings.HasPrefix(argv1, "xn--") == false { return false }

	hostnameok := true
	for _, e := range strings.Split(strings.ToUpper(argv1), "") {
		// Check each characater is a number or an alphabet
		if e[0] <  45              { hostnameok = false; break } //  45 = '-'
		if e[0] == 47              { hostnameok = false; break } //  47 = '/'
		if e[0] >  57 && e[0] < 65 { hostnameok = false; break } //  57 = '9', 65 = 'A'
		if e[0] >  90              { hostnameok = false; break } //  90 = 'Z'
	}
	if hostnameok == false { return false }

	cv := argv1[strings.LastIndex(argv1, ".") + 1:]; if len(cv) > 63 { return false }
	for _, e := range strings.Split(cv, "") {
		// The top level domain should not include a number
		if e[0] > 47 && e[0] < 58  { hostnameok = false; break }
	}
	return hostnameok
}

// IsDomainLiteral() returns true if the domain part is [IPv4:...] or [IPv6:...]
func IsDomainLiteral(email string) bool {
	// @param    string email    Email address string
	// @return   bool            true:  is an domain-literal
	//                           false: is not an domain-literal
	email = strings.Trim(email, "<>")
	if len(email)                     < 16    { return false } // e@[IPv4:0.0.0.0] is 16 characters
	if strings.HasSuffix(email, "]") == false { return false }

	if strings.Contains(email, "@[IPv4:") {
		// neko@[IPv4:192.0.2.25]
		return rfc791.IsIPv4Address(sisimoji.Select(email, "@[IPv4:", "]", 0))

	} else if strings.Contains(email, "@[IPv6:") {
		// neko@[IPv6:2001:0DB8:0000:0000:0000:0000:0000:0001]
		cv := sisimoji.Select(email, "@[IPv6:", "]", 0)
		if len(cv) == 39 && strings.Count(cv, ":") == 7 { return true }
	}
	return false
}

// Find() returns a valid internet hostname found from the argument
func Find(argv1 string) string {
	// @param    string argv1  String including hostnames
	// @return   string        A valid internet hostname found in the argument
	if argv1 == "" { return "" }

	// Replace some string for splitting by " "
	// - mx.example.net[192.0.2.1] => mx.example.net [192.0.2.1]
	// - mx.example.jp:[192.0.2.1] => mx.example.jp :[192.0.2.1]
	sourcetext := strings.ToLower(argv1)
	for _, e := range Prefix0x32 { sourcetext = strings.ReplaceAll(sourcetext, e, " " + e) }
	for _, e := range Suffix0x32 { sourcetext = strings.ReplaceAll(sourcetext, e, e + " ") }

	sourcelist := []string{}
	foundtoken := []string{}
	thelongest := uint8(0)
	hostnameis := ""

	MAKELIST: for {
		for _, e := range Sandwiched {
			// Check a hostname exists between the e[0] and e[1] at slice "Sandwiched"
			// Each slice in Sandwich have 2 elements
			if sisimoji.Aligned(sourcetext, e) == false { continue }
			p1 := strings.Index(sourcetext, e[0])
			p2 := strings.Index(sourcetext, e[1]); cw := len(e[0]); if p1 + cw >= p2 { continue }

			sourcelist = strings.Split(sourcetext[p1 + cw:p2], " ")
			break MAKELIST
		}

		// Check other patterns which are not sandwiched
		for _, e := range StartAfter {
			// StartAfter have some strings, not a slice([]string).
			if strings.Contains(sourcetext, e) == false { continue }
			p1 := strings.Index(sourcetext, e)

			sourcelist = strings.Split(sourcetext[p1 + len(e):], " ")
			break MAKELIST
		}

		for _, e := range ExistUntil {
			// ExistUntil have some strings, not a slice([]string).
			if strings.Contains(sourcetext, e) == false { continue }
			p1 := strings.Index(sourcetext, e)

			sourcelist = strings.Split(sourcetext[0:p1], " ")
			break MAKELIST
		}

		if len(sourcelist) == 0 { sourcelist = strings.Split(sourcetext, " ") }
		break MAKELIST
	}

	for _, e := range sourcelist {
		// Pick some strings which have 4 or more length, is including "." character
		e = strings.TrimRight(e, ".") // Remove "." at the end of the string
		for _, f := range Prefix0x32 { e = strings.ReplaceAll(e, f, "") }
		for _, f := range Suffix0x32 { e = strings.ReplaceAll(e, f, "") }

		if len(e) < 4 || strings.IndexByte(e, '.') < 0 || IsInternetHost(e) == false { continue }
		foundtoken = append(foundtoken, e)
	}
	if len(foundtoken) == 0 { return ""            }
	if len(foundtoken) == 1 { return foundtoken[0] }

	for _, e := range foundtoken {
		// Returns the longest hostname
		cw := uint8(len(e)); if thelongest >= cw { continue }
		hostnameis = e
		thelongest = cw
	}
	return hostnameis
}

