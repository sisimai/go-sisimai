// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _  ___  ___  _  _   
// |  _ \|  ___/ ___/ |( _ )/ _ \| || |  
// | |_) | |_ | |   | |/ _ \ (_) | || |_ 
// |  _ <|  _|| |___| | (_) \__, |__   _|
// |_| \_\_|   \____|_|\___/  /_/   |_|  

// Package "rfc1894" provides funtions related to RFC1894; An Extensible Message Format for Delivery
// Status Notifications (https://datatracker.ietf.org/doc/html/rfc1894)
package rfc1894
import "strings"
import sisimoji "libsisimai.org/sisimai/string"

func FIELDINDEX() []string {
	return []string{
		"Action", "Arrival-Date", "Diagnostic-Code", "Final-Recipient", "Last-Attempt-Date",
		"Original-Recipient", "Received-From-MTA", "Remote-MTA", "Reporting-MTA", "Status",
		"X-Actual-Recipient", "X-Original-Message-ID",
	}
}

// FIELDTABLE() return pairs that a field name and key name defined in sisimai/lhost package
func FIELDTABLE() map[string]string {
	return map[string]string {
		"action":             "action",
		"arrival-date":       "date",
		"diagnostic-code":    "diagnosis",
		"final-recipient":    "recipient",
		"last-attempt-date":  "date",
		"original-recipient": "alias",
		"received-from-mta":  "rhost",
		"remote-mta":         "rhost",
		"reporting-mta":      "lhost",
		"status":             "status",
		"x-actual-recipient": "alias",
    }
}

// Label() returns a lower-cased field name
func Label(argv0 string) string {
	// @param    string  argv0 A line including field and value defined in RFC3464
	// @return   string        Lower-cased field name as a label
	if argv0 == "" || strings.Index(argv0, ":") < 1 { return "" }
	return strings.ToLower(strings.SplitN(argv0, ":", 2)[0])
}

// Match() checks that the argument matches with a field defined in RFC3464 or not
func Match(argv0 string) uint8 {
	// @param    string argv0 A line inlcuding field and value defined in RFC3464
	// @return   uint8        0: not matched, 1: matched with per-message field, 2 is per-recipient
	fieldname0 := map[string]string{
		// https://tools.ietf.org/html/rfc3464#section-2.2
		//   Some fields of a DSN apply to all of the delivery attempts described by that DSN. At
		//   most, these fields may appear once in any DSN. These fields are used to correlate the
		//   DSN with the original message transaction and to provide additional information which
		//   may be useful to gateways.
		//
		//   The following fields (not defined in RFC 3464) are used in Sisimai
		//     - X-Original-Message-ID: <....> (GSuite)
		//
		//   The following fields are not used in Sisimai:
		//     - Original-Envelope-Id
		//     - DSN-Gateway
		"arrival-date":          ":",
		"received-from-mta":     ";",
		"reporting-mta":         ";",
		"x-original-message-id": "@",
	}
	fieldname1 := map[string]string{
		// https://tools.ietf.org/html/rfc3464#section-2.3
		//   A DSN contains information about attempts to deliver a message to one or more
		//   recipients. The delivery information for any particular recipient is contained in a
		//   group of contiguous per-recipient fields. Each group of per-recipient fields is
		//   preceded by a blank line.
		//
		//   The following fields (not defined in RFC 3464) are used in Sisimai
		//     - X-Actual-Recipient: RFC822; ....
		//
		//   The following fields are not used in Sisimai:
		//     - Will-Retry-Until
		//     - Final-Log-ID
		"action":                "e", // fail[e]d, d[e]livered, d[e]layed, [e]xpired, r[e]layed
		"diagnostic-code":       ";",
		"final-recipient":       ";",
		"last-attempt-date":     ":",
		"original-recipient":    ";",
		"remote-mta":            ";",
		"status":                ".",
		"x-actual-recipient":    ";",
	}
	cx := uint8(0)
	cv := Label(argv0)

	for e := range fieldname0 {
		// Per message fields
		if cv != e || strings.Contains(argv0, fieldname0[e]) == false { continue }
		cx = 1; break
	}
	if cx > 0 { return cx }

	for e := range fieldname1 {
		// Per recipient fields
		if cv != e || strings.Contains(argv0, fieldname1[e]) == false { continue }
		cx = 2; break
	}
	return cx
}

// Field() checks that the argument is including field defined in RFC3464 or not and return values
func Field(argv0 string) []string {
	// @param    string   argv0 A line inlcuding field and value defined in RFC3464
	// @return   []string       []string{"field-name", "value-type", "value", "field-group", "comment"}
	if len(argv0) < 7 { return []string{} }

	fieldgroup := map[string]string{
		"original-recipient":    "addr",
		"final-recipient":       "addr",
		"x-actual-recipient":    "addr",
		"diagnostic-code":       "code",
		"arrival-date":          "date",
		"last-attempt-date":     "date",
		"received-from-mta":     "host",
		"remote-mta":            "host",
		"reporting-mta":         "host",
		"action":                "list",
		"status":                "stat",
		"x-original-message-id": "text",
    }
	correction := map[string]string{
		"deliverable": "delivered",
		"expired":     "delayed",
		"failure":     "failed",
    }
	actionlist := []string{"failed", "delayed", "delivered", "relayed", "expanded"}
	captureson := map[string][]string{
		"addr": []string{"Final-Recipient", "Original-Recipient", "X-Actual-Recipient"},
		"code": []string{"Diagnostic-Code"},
		"date": []string{"Arrival-Date", "Last-Attempt-Date"},
		"host": []string{"Received-From-MTA", "Remote-MTA", "Reporting-MTA"},
		"list": []string{"Action"},
		"stat": []string{"Status"},
	//  "text": []string{"X-Original-Message-ID", "Final-Log-ID", "Original-Envelope-ID"}
	}

	parts := strings.SplitN(argv0, ":", 2) // []string{"Final-Recipient", " rfc822; <neko@example.jp>"}
	label := strings.ToLower(parts[0])     // "final-recipient"
	group, nyaan := fieldgroup[label]      // "addr"
	if nyaan == false || len(captureson[group]) == 0 { return []string{} }

	match := false
	for _, e := range captureson[group] {
		// Try to match with each pattern of Per-Message field, Per-Recipient field
		if label == strings.ToLower(e) { match = true; break }
	}
	if match == false { return []string{} }

	// - 0: Field-Name
	// - 1: Sub Type: RFC822, DNS, X-Unix, and so on)
	// - 2: Value
	// - 3: Field Group(addr, code, date, host, stat, text)
	// - 4: Comment
	table   := []string{label, "", "", group, ""}
	parts[1] = strings.TrimSpace(parts[1])

	if group == "addr" || group == "code" || group == "host" {
		// - Final-Recipient: RFC822; kijitora@example.jp
		// - Diagnostic-Code: SMTP; 550 5.1.1 <kijitora@example.jp>... User Unknown
		// - Remote-MTA: DNS; mx.example.jp
		if strings.Contains(parts[1], ";") {
			// There is a valid sub type (including ";")
			v := strings.SplitN(parts[1], ";", 2)
			if len(v) > 0 { table[1] = strings.ToUpper(strings.TrimSpace(v[0])) }
			if len(v) > 1 { table[2] = strings.TrimSpace(v[1])                  }

		} else {
			// There is no sub type like "Diagnostic-Code: 550 5.1.1 <kijitora@example.jp>..."
			table[2] = strings.TrimSpace(parts[1])
			switch group {
				case "addr": table[1] = "RFC822"
				case "code": table[1] = "SMTP"
				case "host": table[1] = "DNS"
			}
		}

		if group == "host" { table[2] = strings.ToLower(table[2]) }
		if len(strings.ReplaceAll(table[2], " ", "")) == 0 { table[2] = "" }

	} else if group == "list" {
		// Action: failed
		// Check that the value is an available value defined in "actionlist" or not.
		// When the value is invalid, convert to an available value defined in "correction"
		v := strings.ToLower(parts[1])
		if sisimoji.EqualsAny(v, actionlist) { table[2] = v }
		if table[2] == "" && len(correction[v]) > 0 { table[2] = correction[v] }

	} else {
		// Other groups such as Status:, Arrival-Date:, or X-Original-Message-ID:.
		// There is no ";" character in the field.
		// - Status: 5.2.2
		// - Arrival-Date: Mon, 21 May 2018 16:09:59 +0900
		table[2] = parts[1]
		if group != "date" { table[2] = strings.ToLower(parts[1]) }
	}

	if sisimoji.Aligned(table[2], []string{" (", ")"}) {
		// Extract text enclosed in parentheses as comments
		// Reporting-MTA: dns; mr21p30im-asmtp004.me.example.com (tcp-daemon)
		p1 := strings.LastIndex(table[2], " (")
		p2 := strings.LastIndex(table[2], ")" )
		table[4] = table[2][p1 + 2:p2]
		table[2] = table[2][0:p1]
	}

	return table
}

