// Copyright (C) 2020-2022,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      

package message
import "bytes"
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc1894"
import "libsisimai.org/sisimai/v5/rfc5322"
import "libsisimai.org/sisimai/v5/rfc5965"

var fieldtable = makefield(rfc1894.FieldIndex, rfc5322.FieldIndex, rfc5965.FieldIndex)
var mediatypes = [][]string{
	{"message/xdelivery-status",                "message/delivery-status"},
	{"message/disposition-notification",        "message/delivery-status"},
	{"message/global-delivery-status",          "message/delivery-status"},
	{"message/global-disposition-notification", "message/delivery-status"},
	{"message/global-delivery-status",          "message/delivery-status"},
	{"message/global-headers",                  "text/rfc822-headers"},
	{"message/global",                          "message/rfc822"},
}

// makefield generates a map including each field name defined in RFC1894, RFC5322, and RFC5965.
//   Arguments:
//     - fs ([]string): Slices including field names.
//   Returns:
//     - (map[string]string): Hash table.
func makefield(fs ...[]string) map[string]string {
	fieldtable := make(map[string]string, 40)
	for _, e := range slices.Concat(fs...) { fieldtable[strings.ToLower(e)] = e }
	return fieldtable
}

// tidy tidies up each field name and format of email headers.
//   Arguments:
//     - head ([]byte): String including fields and values in email headers.
//   Returns:
//     - ([]byte): String tidied up.
func tidy(head []byte) []byte {
	if len(head) == 0 { return head }

	// Find and tidy up fields defined in RFC5322, RFC1894, and RFC5965
	bu := bytes.Buffer{}; bu.Grow(1024)
	el := bytes.Split(head, []byte("\n")); for j, e := range el {
		// 1. Find a field label defined in RFC5322, RFC1894, or RFC5965 from this line
		p0 := bytes.IndexByte(e, ':'); if p0 < 0 { bu.Write(append(e, '\n')); continue }
		cf := string(bytes.ToLower(bytes.TrimRight(e[0:p0], " ")))
		if strings.IndexByte(cf, ' ') > 0 { bu.Write(append(e, '\n')); continue }
		fn := fieldtable[cf]; if fn == "" { bu.Write(append(e, '\n')); continue }

		// 2. Tidy up a sub type of each field defined in RFC1894 such as Reporting-MTA: DNS;...
		ab := make([][]byte, 0, 2)
		bf := e[p0 + 1:]

		// Such as Diagnostic-Code, Remote-MTA, and so on
		// - Before: Diagnostic-Code: SMTP;550 User unknown
		// - After:  Diagnostic-Code: smtp; 550 User unknown
		match := false; for _, ef := range rfc1894.FieldIndex {
			// The field name is not listed in RFC1894
			if fn == ef || fn == "Content-Type" { match = true; break }
		}
		if match == true && bytes.Contains(bf, []byte(";")) == true {
			// The field including one or more ";"
			for _, ef := range bytes.Split(bf, []byte(";")) {
				// 2-1. Trim leading and trailing space characters from the current buffer
				ef = bytes.Trim(ef, " ")

				// 2-2. Convert some parameters to the lower-cased string
				if ps := []byte{}; bytes.IndexByte(ef, ' ') < 1 {
					// For example,
					// - Content-Type: Message/delivery-status => message/delivery-status
					// - Content-Type: Charset=UTF8            => charset=utf8
					// - Reporting-MTA: DNS; ...               => dns
					// - Final-Recipient: RFC822; ...          => rfc822
					if cv := moji.Select(moji.LHS + string(ef), "", "=", 0); cv != "" {
						// charset=, boundary=, and other pairs divided by "="
						ps = []byte(strings.ToLower(cv))
						ef = bytes.Replace(ef, []byte(cv), ps, 1)
					}
					if bytes.Equal(ps, []byte("boundary")) == false { ef = bytes.ToLower(ef) }
					if bytes.Equal(ef, []byte("rfc/822"))  == true  { ef = []byte("rfc822")  }
				}
				ab = append(ab, []byte(ef))
			}

			if fn == "Diagnostic-Code" && len(ab) == 1 && bytes.IndexByte(el[j + 1], ' ') != 0 {
				// Diagnostic-Code: x-unix;
				//   /var/email/kijitora/Maildir/tmp/1000000000.A000000B00000.neko22:
				//   Disk quota exceeded
				ab = append(ab, []byte(""))
			}
			bf = bytes.Join(ab, []byte("; "))
			ab = make([][]byte, 0, 2)

		} else {
			// There is no ";" in the field
			if moji.ContainsAny(fn, []string{"-Date", "-Message-ID"}) == false { bf = bytes.ToLower(bf) }
		}

		// 3. Tidy up a value, and a parameter of Content-Type: field 
		if fn == "Content-Type" {
			// Replace the value of "Content-Type" field
			for _, ef := range mediatypes {
				// - Before: Content-Type: message/xdelivery-status; ...
				// - After:  Content-Type: message/delivery-status; ...
				bf = bytes.Replace(bf, []byte(ef[0]), []byte(ef[1]), 1)
			}
		}

		// 4. Concatenate the field name and the field value
		for _, ef := range bytes.Split(bf, []byte(" ")) {
			// Remove redundant space characters
			if len(ef) > 0 { ab = append(ab, ef) }
		}
		bu.WriteString(fn + ": " + string(bytes.Join(ab, []byte(" "))) + "\n")
	}
	email := bu.Bytes();

	// 5. Convert the lower-cased SMTP command to the upper-cased.
	email = bytes.ReplaceAll(email, []byte("after end of data:"), []byte("after end of DATA:"))
	if bytes.HasSuffix(email, []byte("\n\n")) == false { email = append(email, []byte("\n\n")...) }
	return email
}

