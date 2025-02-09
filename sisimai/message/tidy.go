// Copyright (C) 2020-2022,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package message

//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      
import "fmt"
import "strings"

// tidy() tidies up each field name and format
func tidy(argv0 *string) *string {
	// @param    *string argv0 String including field and value used at an email
	// @return   *string       String tidied up
	email := ""; if len(*argv0) < 1 { return &email }
	lines := strings.Split(*argv0, "\n")

	// Find and tidy up fields defined in RFC5322, RFC1894, and RFC5965
	for i, e := range lines {
		// 1. Find a field label defined in RFC5322, RFC1894, or RFC5965 from this line
		p0 := strings.IndexByte(e, ':'); if p0 < 0                    { email += e + "\n"; continue }
		cf := strings.ToLower(e[0:p0]);  if strings.Contains(cf, " ") { email += e + "\n"; continue }
		fn := FieldTable[cf]

		// There is neither ":" character nor a field listed in FieldTable
		if len(fn) == 0 { email += e + "\n"; continue }

		// 2. Tidy up a sub type of each field defined in RFC1894 such as Reporting-MTA: DNS;...
		ab := []string{}
		bf := e[p0 + 1:]
		p1 := strings.IndexByte(bf, ';')
		for {
			// Such as Diagnostic-Code, Remote-MTA, and so on
			// - Before: Diagnostic-Code: SMTP;550 User unknown
			// - After:  Diagnostic-Code: smtp; 550 User unknown
			match := false
			for _, f := range Fields1894 {
				// The field name is not listed in RFC1894
				if fn == f || fn == "Content-Type" { match = true; break }
			}
			if match == false { break }

			if p1 > 0 {
				// The field including one or more ";"
				for _, f := range strings.Split(bf, ";") {
					// 2-1. Trim leading and trailing space characters from the current buffer
					f   = strings.Trim(f, " ")
					ps := ""

					// 2-2. Convert some parameters to the lower-cased string
					for {
						// For example,
						// - Content-Type: Message/delivery-status => message/delivery-status
						// - Content-Type: Charset=UTF8            => charset=utf8
						// - Reporting-MTA: DNS; ...               => dns
						// - Final-Recipient: RFC822; ...          => rfc822
						if strings.IndexByte(f, ' ') > 0 { break }

						p2 := strings.IndexByte(f, '=')
						if p2 > 0 {
							// charset=, boundary=, and other pairs divided by "="
							ps = strings.ToLower(f[0:p2])
							f  = strings.Replace(f, f[0:p2], ps, 1)
						}
						if ps != "boundary" { f = strings.ToLower(f) }
						break
					}
					ab = append(ab, f)
				}

				for {
					// Diagnostic-Code: x-unix;
					//   /var/email/kijitora/Maildir/tmp/1000000000.A000000B00000.neko22:
					//   Disk quota exceeded
					if fn != "Diagnostic-Code" || len(ab)   != 1 { break }
					if strings.IndexByte(lines[i + 1], ' ') != 0 { break }

					ab = append(ab, ""); break
				}
				bf = strings.Join(ab, "; ")
				ab = []string{}

			} else {
				// There is no ";" in the field
				if strings.Index(fn, "-Date") > 0 || strings.Index(fn, "-Message-ID") > 0 { break }
				bf = strings.ToLower(bf)
			}
			break
		}

		// 3. Tidy up a value, and a parameter of Content-Type: field 
		if len(ReplacesAs[fn]) > 0 {
			// Replace the value of "Content-Type" field
			for _, f := range ReplacesAs[fn] {
				// - Before: Content-Type: message/xdelivery-status; ...
				// - After:  Content-Type: message/delivery-status; ...
				p1 = strings.Index(bf, f[0]); if p1 < 0 { continue }
				bf = strings.Replace(bf, f[0], f[1], 1)
			}
		}

		// 4. Concatenate the field name and the field value
		for _, f := range strings.Split(bf, " ") {
			// Remove redundant space characters
			if len(f) == 0 { continue }
			ab = append(ab, f)
		}
		email += fmt.Sprintf("%s: %s\n", fn, strings.Join(ab, " "))
	}

	if email[len(email) - 2:len(email)] != "\n\n" { email += "\n\n" }
	return &email
}

