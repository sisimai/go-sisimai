// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  ____  _____ ____ ____ _________  ____     ______               _               _   
// |  _ \|  ___/ ___| ___|___ /___ \|___ \   / /  _ \ ___  ___ ___(_)_   _____  __| |_ 
// | |_) | |_ | |   |___ \ |_ \ __) | __) | / /| |_) / _ \/ __/ _ \ \ \ / / _ \/ _` (_)
// |  _ <|  _|| |___ ___) |__) / __/ / __/ / / |  _ <  __/ (_|  __/ |\ V /  __/ (_| |_ 
// |_| \_\_|   \____|____/____/_____|_____/_/  |_| \_\___|\___\___|_| \_/ \___|\__,_(_)
import "strings"
import "sisimai/address"
import sisimoji "sisimai/string"

// Received() convert Received headers to a structured data
func Received(argv1 string) [6]string {
	// @param    string    argv1  Received header
	// @return   []string         Each item in the Received header order by the following:
	//                            0: (from)   "hostname"
	//                            1: (by)     "hostname"
	//                            2: (via)    "protocol/tcp"
	//                            3: (with)   "protocol/smtp"
	//                            4: (id)     "queue-id"
	//                            5: (for)    "envelope-to address"
	// Received: (qmail 10000 invoked by uid 999); 24 Apr 2013 00:00:00 +0900
	if strings.IndexByte(argv1, ' ') < 0                { return [6]string{} }
	if strings.Contains(argv1, " invoked by uid")       { return [6]string{} }
	if strings.Contains(argv1, " invoked from network") { return [6]string{} }

	// - https://datatracker.ietf.org/doc/html/rfc5322
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
	recvd := strings.Split(argv1, " ")
	label := [6]string{"from", "by", "via", "with", "id", "for"}
	token := make(map[string]string)
	other := []string{}
	alter := []string{}
	right := false

	for i, e := range recvd {
		// Look up each label defined in label from Received header
		f := strings.ToLower(e)
		p := false
		for _, v := range label { if f == v { p = true; break } }

		if p == false             { continue }
		if i + 1 > len(recvd) - 1 { continue }

		token[f] = strings.ToLower(recvd[i + 1]);
		token[f] = strings.ReplaceAll(token[f], "(", "")
		token[f] = strings.ReplaceAll(token[f], ")", "")
		token[f] = strings.ReplaceAll(token[f], ";", "")

		if f != "from"                           { continue }
		if i + 2 > len(recvd) - 1                { break    }
		if strings.Index(recvd[i + 2], "(") != 0 { continue }

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
		other    = append(other, recvd[i + 2])
		other[0] = strings.ReplaceAll(other[0], "(", "")
		other[0] = strings.ReplaceAll(other[0], ")", "")
		other[0] = strings.ReplaceAll(other[0], ";", "")

		// The 2nd element after the current element is a continuation of the current element.
		// such as "(c213502.kyoto.example.ne.jp", "[192.0.2.135])"
		if i + 3 > len(recvd) - 1 { break }
		other    = append(other, recvd[i + 3])
		other[1] = strings.ReplaceAll(other[1], "(", "")
		other[1] = strings.ReplaceAll(other[1], ")", "")
		other[1] = strings.ReplaceAll(other[1], ";", "")
	}

	for _, e := range other {
		// Check alternatives in "other", and then delete uninformative values.
		if len(e) < 4         { continue }
		if e == "unknown"     { continue }
		if e == "localhost"   { continue }
		if e == "[127.0.0.1]" { continue }
		if e == "[IPv6:::1]"  { continue }
		if strings.IndexByte(e, '.') == -1 { continue }
		if strings.IndexByte(e, '=')  >  1 { continue }
		alter = append(alter, e)
	}

	for _, e := range []string{"from", "by"} {
		// Remove square brackets from the IP address such as "[192.0.2.25]"
		if len(token[e]) == 0 { continue }
		if strings.IndexByte(token[e], '[') != 0 { continue }

		p := sisimoji.FindIPv4Address(token[e])
		if len(p) > 0 { token[e] = p[0] } else { token[e] = "" }
	}
	_, e := token["from"]; if e == false { token["from"] = "" }

	for {
		// Prefer hostnames over IP addresses, except for localhost.localdomain and similar.
		if token["from"] == "localhost"             { break }
		if token["from"] == "localhost.localdomain" { break }
		if strings.Index(token["from"], ".") < 0    { break } // A hostname without a domain name
		if len(sisimoji.FindIPv4Address(token["from"])) > 0 { break }

		// No need to rewrite token["from"]
		right = true
		break
	}

	for {
		// Try to rewrite uninformative hostnames and IP addresses in token["from"]
		if right == true   { break } // There is no need to rewrite
		if len(alter) == 0 { break } // There is no alternative for rewriting
		if strings.Contains(alter[0], token["from"]) { break }

		if strings.HasPrefix(token["from"], "localhost") {
			// "localhost" or "localhost.localdomain"
		} else if strings.IndexByte(token["from"], '.') == -1 {
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
	if len(token["for"])   > 0 { token["for"] = address.S3S4(token["for"]) }

	for _, e := range label {
		// Delete an invalid value
		if len(token[e]) == 0              { token[e] = ""; continue }
		if strings.Contains(token[e], " ") { token[e] = ""; continue }
		if strings.Contains(token[e], "[") { strings.Replace(token[e], "[", "", 1) }
		if strings.Contains(token[e], "]") { strings.Replace(token[e], "]", "", 1) }
	}

	return [6]string{token["from"], token["by"], token["via"], token["with"], token["id"], token["for"]}
}

