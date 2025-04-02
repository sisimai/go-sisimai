// Copyright (C) 2020-2022,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      

package message
import "strings"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc1894"
import "libsisimai.org/sisimai/v5/rfc5322"
import "libsisimai.org/sisimai/v5/rfc5965"

var fieldtable = makefield(rfc1894.FieldIndex, rfc5322.FieldIndex, rfc5965.FieldIndex)
var replacesas = map[string][][]string{
    "Content-Type": [][]string{
		{"message/xdelivery-status",         "message/delivery-status"},
		{"message/disposition-notification", "message/delivery-status"},
	},
}

// makefield generates a map including each field name defined in RFC1894, RFC5322, and RFC5965.
func makefield(argv1 []string, argv2 []string, argv3 []string) map[string]string {
	fieldtable := map[string]string{}
	for _, e := range argv1 { fieldtable[strings.ToLower(e)] = e }
	for _, e := range argv2 { fieldtable[strings.ToLower(e)] = e }
	for _, e := range argv3 { fieldtable[strings.ToLower(e)] = e }
	return fieldtable
}

// tidy tidies up each field name and format of email headers.
//   Arguments:
//     - argv0 (*string): String including fields and values in email headers
//   Returns:
//     - (*string):       String tidied up
func tidy(argv0 *string) *string {
	if argv0 == nil || *argv0 == "" { return nil }

	lines := strings.Split(*argv0, "\n")
	email := ""; bu := strings.Builder{}; bu.Grow(1024)

	// Find and tidy up fields defined in RFC5322, RFC1894, and RFC5965
	for j, e := range lines {
		// 1. Find a field label defined in RFC5322, RFC1894, or RFC5965 from this line
		p0 := strings.IndexByte(e, ':'); if p0 < 0                         { bu.WriteString(e + "\n"); continue }
		cf := strings.ToLower(e[0:p0]);  if strings.IndexByte(cf, ' ') > 0 { bu.WriteString(e + "\n"); continue }
		fn := fieldtable[cf];            if fn == ""                       { bu.WriteString(e + "\n"); continue }

		// 2. Tidy up a sub type of each field defined in RFC1894 such as Reporting-MTA: DNS;...
		ab := make([]string, 0, 2)
		bf := e[p0 + 1:]
		p1 := strings.IndexByte(bf, ';')
		for {
			// Such as Diagnostic-Code, Remote-MTA, and so on
			// - Before: Diagnostic-Code: SMTP;550 User unknown
			// - After:  Diagnostic-Code: smtp; 550 User unknown
			match := false; for _, f := range rfc1894.FieldIndex {
				// The field name is not listed in RFC1894
				if fn == f || fn == "Content-Type" { match = true; break }
			}
			if match == false { break }

			if p1 > 0 {
				// The field including one or more ";"
				for _, f := range strings.Split(bf, ";") {
					// 2-1. Trim leading and trailing space characters from the current buffer
					f = strings.Trim(f, " ")

					// 2-2. Convert some parameters to the lower-cased string
					ps := ""; for strings.IndexByte(f, ' ') < 1 {
						// For example,
						// - Content-Type: Message/delivery-status => message/delivery-status
						// - Content-Type: Charset=UTF8            => charset=utf8
						// - Reporting-MTA: DNS; ...               => dns
						// - Final-Recipient: RFC822; ...          => rfc822
						if p2 := strings.IndexByte(f, '='); p2 > 0 {
							// charset=, boundary=, and other pairs divided by "="
							ps = strings.ToLower(f[0:p2])
							f  = strings.Replace(f, f[0:p2], ps, 1)
						}
						if ps != "boundary" { f = strings.ToLower(f) }
						break
					}
					ab = append(ab, f)
				}

				if fn == "Diagnostic-Code" && len(ab) == 1 && strings.IndexByte(lines[j + 1], ' ') != 0 {
					// Diagnostic-Code: x-unix;
					//   /var/email/kijitora/Maildir/tmp/1000000000.A000000B00000.neko22:
					//   Disk quota exceeded
					ab = append(ab, "")
				}
				bf = strings.Join(ab, "; ")
				ab = make([]string, 0, 2)

			} else {
				// There is no ";" in the field
				if moji.ContainsAny(fn, []string{"-Date", "-Message-ID"}) == false { bf = strings.ToLower(bf) }
			}
			break
		}

		// 3. Tidy up a value, and a parameter of Content-Type: field 
		if len(replacesas[fn]) > 0 {
			// Replace the value of "Content-Type" field
			for _, f := range replacesas[fn] {
				// - Before: Content-Type: message/xdelivery-status; ...
				// - After:  Content-Type: message/delivery-status; ...
				p1 = strings.Index(bf, f[0]); if p1 > -1 { bf = strings.Replace(bf, f[0], f[1], 1) }
			}
		}

		// 4. Concatenate the field name and the field value
		for _, f := range strings.Split(bf, " ") {
			// Remove redundant space characters
			if f != "" { ab = append(ab, f) }
		}
		bu.WriteString(fn + ": " + strings.Join(ab, " ") + "\n")
	}

	email = bu.String();
	if email[len(email) - 2:len(email)] != "\n\n" { email += "\n\n" }
	return &email
}

