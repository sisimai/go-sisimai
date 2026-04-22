// Copyright (C) 2020,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____ _________  ____     ______               _               _   
// |  _ \|  ___/ ___| ___|___ /___ \|___ \   / /  _ \ ___  ___ ___(_)_   _____  __| |_ 
// | |_) | |_ | |   |___ \ |_ \ __) | __) | / /| |_) / _ \/ __/ _ \ \ \ / / _ \/ _` (_)
// |  _ <|  _|| |___ ___) |__) / __/ / __/ / / |  _ <  __/ (_|  __/ |\ V /  __/ (_| |_ 
// |_| \_\_|   \____|____/____/_____|_____/_/  |_| \_\___|\___\___|_| \_/ \___|\__,_(_)

package rfc5322
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc791"

// Received convert Received headers to a structured data.
//   Arguments:
//     - rhead (string): Received header.
//   Returns:
//     - ([]string):     Each item in the Received header order by the following:
//                       - 0: (from)   "hostname"
//                       - 1: (by)     "hostname"
//                       - 2: (via)    "protocol/tcp"
//                       - 3: (with)   "protocol/smtp"
//                       - 4: (id)     "queue-id"
//                       - 5: (for)    "envelope-to address"
//   See:
//     - https://datatracker.ietf.org/doc/html/rfc5322#section-3.6.7
func Received(rhead string) [6]string {
	//   received        =   "Received:" *received-token ";" date-time CRLF
	//   received-token  =   word / angle-addr / addr-spec / domain
	//
	// - Appendix A.4. Message with Trace Fields
	//   Received:
	//       from x.y.test
	//       by example.net
	//       via TCP
	//       with ESMTP
	//       id ABC12345
	//       for <mary@example.net>;  21 Nov 1997 10:05:43 -0600
	if strings.IndexByte(rhead, ' ') < 0 || moji.ContainsAny(rhead, woReceived) { return [6]string{} }

	recvd := strings.Split(rhead, " ")
	label := [6]string{"from", "by", "via", "with", "id", "for"}
	skips := []string{"unknown", "localhost", "[127.0.0.1]", "[IPv6:::1]"}
	chars := []string{"(", ")", ";"} // Removed by strings.ReplaceAll()
	token := make(map[string]string, 6)
	other := make([]string, 6)
	alter := make([]string, 6)
	right := false

	for j, e := range recvd {
		// Look up each label defined in label from Received header
		cf := strings.ToLower(e)
		cb := false

		for _, v := range label { if cf == v { cb = true; break } }
		if cb == false || j + 1 > len(recvd) - 1 { continue }

		token[cf] = strings.ToLower(recvd[j + 1]);
		for _, f := range chars { token[cf] = strings.ReplaceAll(token[cf], f, "") }

		if cf != "from"                              { continue }
		if j + 2 > len(recvd) - 1                    { break    }
		if strings.IndexByte(recvd[j + 2], '(') != 0 { continue }

		// Get and keep a hostname in the comment as follows:
		// from mx1.example.com (c213502.kyoto.example.ne.jp [192.0.2.135]) by mx.example.jp (V8/cf)
		// []string{
		//  "from",                         // index + 0
		//  "mx1.example.com",              // index + 1
		//  "(c213502.kyoto.example.ne.jp", // index + 2
		//  "[192.0.2.135])",               // index + 3
		//  "by",
		//  "mx.example.jp",
		//  "(V8/cf)",
		//  ...
		// }
		// The 2nd element after the current element is NOT a continuation of the current element
		// such as "(c213502.kyoto.example.ne.jp)"
		other = append(other, recvd[j + 2])
		for _, f := range chars { other[0] = strings.ReplaceAll(other[0], f, "") }

		// The 2nd element after the current element is a continuation of the current element.
		// such as "(c213502.kyoto.example.ne.jp", "[192.0.2.135])"
		if j + 3 > len(recvd) - 1 { break }
		other = append(other, recvd[j + 3])
		for _, f := range chars { other[1] = strings.ReplaceAll(other[1], f, "") }
	}

	for _, e := range other {
		// Check alternatives in "other", and then delete uninformative values.
		if len(e) < 4 || slices.Contains(skips, e) { continue }
		if strings.IndexByte(e, '.') == -1         { continue }
		if strings.IndexByte(e, '=')  >  1         { continue }
		alter = append(alter, e)
	}

	for _, e := range []string{"from", "by"} {
		// Remove square brackets from the IP address such as "[192.0.2.25]"
		if token[e] == "" || strings.IndexByte(token[e], '[') != 0 { continue }

		cv := rfc791.FindIPv4Address(token[e])
		if len(cv) > 0 { token[e] = cv[0] } else { token[e] = "" }
	}
	if _, e := token["from"]; e == false { token["from"] = "" }

	switch {
		// Prefer hostnames over IP addresses, except for localhost.localdomain and similar.
		case token["from"] == "localhost" || token["from"] == "localhost.localdomain":
		case strings.IndexByte(token["from"], '.') < 0:
		case len(rfc791.FindIPv4Address(token["from"])) > 0:
		default: right = true // No need to rewrite token["from"]
	}

	for right == false && len(alter) > 0 {
		// Try to rewrite uninformative hostnames and IP addresses in token["from"]
		if strings.Contains(alter[0], token["from"]) { break }
		if strings.IndexByte(token["from"], '.') == -1 {
			// A hostname without a domain name such as "mail", "mx", or "mbox"
			if strings.IndexByte(alter[0], '.') > 0 { token["from"] = alter[0] }

		} else {
			// An IPv4 address
			token["from"] = alter[0]
		}
		break
	}
	if len(token["by"])   == 0 { delete(token, "by")   }
	if len(token["from"]) == 0 { delete(token, "from") }
	if len(token["for"])   > 0 { token["for"] = strings.Trim(token["for"], "<>") }

	for _, e := range label {
		// Delete an invalid value
		if token[e] == ""                        { token[e] = ""; continue }
		if strings.IndexByte(token[e], ' ') > -1 { token[e] = ""; continue }
		if strings.IndexByte(token[e], '[') > -1 { token[e] = strings.Replace(token[e], "[", "", 1) }
		if strings.IndexByte(token[e], ']') > -1 { token[e] = strings.Replace(token[e], "]", "", 1) }
	}

	return [6]string{token["from"], token["by"], token["via"], token["with"], token["id"], token["for"]}
}

