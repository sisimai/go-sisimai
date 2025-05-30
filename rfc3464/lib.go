// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _____ _  _    __   _  _   
// |  _ \|  ___/ ___|___ /| || |  / /_ | || |  
// | |_) | |_ | |     |_ \| || |_| '_ \| || |_ 
// |  _ <|  _|| |___ ___) |__   _| (_) |__   _|
// |_| \_\_|   \____|____/   |_|  \___/   |_|  

// Package "rfc3464" provides functions like a MTA module in "lhost" package for decoding bounce
// messages formatted according to RFC3464; An Extensible Message Format for Delivery Status Notifications
// https://datatracker.ietf.org/doc/html/rfc3464
package rfc3464
import "slices"
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/lhost"
import "libsisimai.org/sisimai/v5/address"
import "libsisimai.org/sisimai/v5/rfc1894"
import "libsisimai.org/sisimai/v5/rfc2045"
import "libsisimai.org/sisimai/v5/rfc5322"
import "libsisimai.org/sisimai/v5/smtp/reply"
import "libsisimai.org/sisimai/v5/smtp/status"
import "libsisimai.org/sisimai/v5/smtp/command"

// Inquire decodes a bounce message that have fields defined in RFC3464.
//   Arguments:
//     - bf (*sis.BeforeFact):  Message entity in progress
//   Returns:
//     - (*sis.RisingUnderway): A structure as a staging data that is processed in message.sift() function
//   See:
//     - https://datatracker.ietf.org/doc/html/rfc3464
func Inquire(bf *sis.BeforeFact) *sis.RisingUnderway {
	if bf == nil || bf.IsEmpty() == true { return nil }

	boundaries := []string{
		// When the new value added, the part of the value should be listed in "delimiters" variable
		// defined at MakeFlat() function in rfc2045/make-multipart-flat.go
		"Content-Type: message/rfc822",
		"Content-Type: text/rfc822-headers",
		"Content-Type: message/partial",
		"Content-Disposition: inline", // See lhost-amavis-*.eml, lhost-facebook-*.eml
	}
	startingof := map[string][]string{"message": []string{"Content-Type: message/delivery-status"}}
	dontappend := []string{
		"Content-", "This is a MIME", "This is a multi", "This is an auto", "This multi-part", "###", "***", "--",
	}

	for moji.ContainsAny(bf.Payload, boundaries) == false {
		// There is no "Content-Type: message/rfc822" line in the message body
		// Insert "Content-Type: message/rfc822" before "Return-Path:" of the original message
		cv := "\n\nReturn-Path:"; if strings.Contains(bf.Payload, cv) == false { break }
		bf.Payload = strings.Replace(bf.Payload, cv, "\n\n" + boundaries[0] + cv, 1)
		break
	}
	permessage := map[string]string{}   // Store values of each Per-Message field
	keystrings := make([]string, 0, 4)  // Key list of permessage
	dscontents := make([]sis.DeliveryMatter, 1); v := &dscontents[0]
	alternates := new(sis.DeliveryMatter)
	emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
	readcursor := uint8(0)              // Points the current cursor position
	readslices := make([]string, 1, 32) // Copy each line for later reference
	recipients := uint8(0)              // The number of 'Final-Recipient' header
	goestonext := false                 // Flag: do not append the line into "leadinbuff"
	leadinbuff := strings.Builder{}; leadinbuff.Grow(len(emailparts[0]) / 2)
	eachbuffer := []strings.Builder{{}}; b := &eachbuffer[0]; b.Grow(128)
	isboundary := []string{rfc2045.Boundary(bf.Headers["content-type"][0], 0)}

	for strings.IndexByte(emailparts[0], '@') == -1 {
		// There is no email address in the first element of emailparts
		// There is a bounce message inside of message/rfc822 part at lhost-x5-*, rfc3464/1311
		p0 := -1 // The index of the boundary string found first
		p1 :=  0 // Offset position of the message body after the boundary string
		ct := "" // Boundary string found first such as "Content-Type: message/rfc822"
		for _, e := range boundaries {
			// Look for a boundary string from the message body
			p0 = strings.Index(bf.Payload, e + "\n"); if p0 < 0 { continue }
			p1 = p0 + len(e) + 2
			ct = e; break
		}
		if p0 < 0 { break } // There is no boundary string

		cx := bf.Payload[p1:]
		p2 := strings.Index(cx, "\n\n")
		cv := cx[p2 + 2:]
		emailparts = rfc5322.Part(&cv, []string{ct}, false); break
	}

	for strings.Contains(emailparts[0], startingof["message"][0]) == false {
		// There is no "Content-Type: message/delivery-status" line in the message body
		// Insert "Content-Type: message/delivery-status" before "Reporting-MTA:" field
		cv := "\n\nReporting-MTA:"; if strings.Contains(emailparts[0], cv) == false { break }
		emailparts[0] = strings.Replace(emailparts[0], cv, "\n\n" + startingof["message"][0] + cv, 1)
		break
	}

	for _, e := range []string{"Final-Recipient", "Original-Recipient"} {
		// Fix the malformed field "Final-Recipient: <kijitora@example.jp>"
		cv := "\n" + e + ": "
		cx := cv + "<"; if strings.Contains(emailparts[0], cx) == false { continue }

		// Insert "rfc822; " just after the field name
		emailparts[0] = strings.Replace(emailparts[0], cv + "<", cv + "rfc822; ", 1)
		p0 := strings.Index(emailparts[0], cv)
		p1 := moji.IndexOnTheWay(emailparts[0], ">\n", p0 + 1)
		emailparts[0] = emailparts[0][:p1] + emailparts[0][p1 + 1:]
	}

	j := -1; for e := range strings.Lines(emailparts[0]) {
		// Read error messages and delivery status lines from the head of the email to the
		// previous line of the beginning of the original message.
		e = strings.TrimRight(e, "\n\r")
		j++; readslices = append(readslices, e) // Save the current line for the next loop

		if readcursor == 0 {
			// Beginning of the bounce message or message/delivery-status part
			if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= lhost.Indicators["deliverystatus"] }

			for {
				// Append each string before startingof["message"][0] except the following patterns
				// for the later reference
				if e == "" || goestonext { break } // Skip if the line is empty or the part is text/html, image/icon in multipart/*

				// This line is a boundary kept in "multiparts" as a string, when the end of
				// the boundary appeared, the condition above also returns true.
				if moji.HasPrefixAny(e, isboundary) { goestonext = false; break }
				if strings.HasPrefix(e, "Content-Type:") {
					// Content-Type: field in multipart/*
					if strings.Contains(e, "multipart/") {
						// Content-Type: multipart/alternative; boundary=aa00220022222222ffeebb
						// Pick the boundary string and store it into "isboucdary"
						isboundary = append(isboundary, rfc2045.Boundary(e, 0))

					} else if strings.Contains(e, "text/plain") {
						// Content-Type: "text/plain"
						goestonext = false

					} else {
						// Other types: for example, text/html, image/jpg, and so on
						goestonext = true
					}
					break
				}
				if moji.HasPrefixAny(e, dontappend)      { break }
				if strings.Contains(e, "--- The follow") { break } // ----- The following addresses had delivery problems -----
				if strings.Contains(e, "--- Transcript") { break } // ----- Transcript of session follows -----
				leadinbuff.WriteString(e + " ");           break
			}
			continue
		}
		if readcursor & lhost.Indicators["deliverystatus"] == 0 || e == "" { continue }

		if f := rfc1894.Match(e); f > 0 {
			// This line matched with any field defined in RFC3464
			o := rfc1894.Field(e); if len(o) == 0 { continue }
			z := rfc1894.FieldTable[o[0]]
			v  = sis.TailDeliveryMatter(&dscontents)
			b  = &(eachbuffer[len(eachbuffer) - 1]); b.Grow(128)

			if o[3] == "addr" {
				// Final-Recipient: rfc822; kijitora@example.jp
				// X-Actual-Recipient: rfc822; kijitora@example.co.jp
				if o[0] == "final-recipient" {
					// Final-Recipient: rfc822; kijitora@example.jp
					// Final-Recipient: x400; /PN=...
					cv := address.S3S4(o[2]); if rfc5322.IsEmailAddress(cv) == false     { continue }
					cw := len(dscontents); if cw > 0 && cv == dscontents[cw-1].Recipient { continue }

					if len(v.Recipient) > 0 {
						// There are multiple recipient addresses in the message body.
						v = sis.NextDeliveryMatter(&dscontents)
						eachbuffer = append(eachbuffer, strings.Builder{})
						b = &(eachbuffer[len(eachbuffer) - 1]); b.Grow(128)
					}
					v.Recipient = cv
					recipients += 1

				} else {
					// X-Actual-Recipient: rfc822; kijitora@example.co.jp
					v.Alias = o[2]
				}
			} else if o[3] == "code" {
				// Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
				v.Spec = o[1]
				b.WriteString(o[2] + " ")

			} else {
				// Other DSN fields defined in RFC3464
				// There are other error messages as a comment such as the following:
				// Status: 5.0.0 (permanent failure)
				// Status: 4.0.0 (cat.example.net: host name lookup failure)
				if o[4] != "" { b.WriteString(" " + o[4] + " ") }
				v.Update(v.AsRFC1894(o[0]), o[2]); if f != 1 { continue }

				// Copy the lower-cased member name of sis.DeliveryMatter{} for "permessage" for
				// the later reference
				permessage[z] = o[2]
				if slices.Contains(keystrings, z) == false { keystrings = append(keystrings, z) }
			}
		} else {
			// Check that the line is a continued line of the value of Diagnostic-Code: field or not
			if strings.HasPrefix(e, "X-") && strings.Contains(e, ": ") {
				// This line is a MTA-Specific fields begins with "X-"
				if is3rdparty(e) == false { continue }
				if cv := xfield(e); len(cv) > 0 && len(rfc1894.FieldTable[strings.ToLower(cv[0])]) == 0 {
					// Check the first element is a field defined in RFC1894 or not
					// cv[4] is a string line "reason:mailboxfull"
					v.Reason = moji.Select(cv[4] + moji.RHS, "reason:", "", 0)

				} else {
					// Set the value picked from "X-*" field to the member of sis.DeliveryMatter
					// when the current value is empty
					z := rfc1894.FieldTable[strings.ToLower(cv[0])]; if len(z) < 1 { continue }
					if v.Select(z) == "" { v.Update(z, cv[2]) }
				}
			} else {
				// The line may be a continued line of the value of the Diagnostic-Code: field
				if strings.HasPrefix(readslices[j], "Diagnostic-Code:") == false {
					// In the case of multiple "message/delivery-status" line
					if strings.HasPrefix(e, "Content-") { continue } // Content-Disposition, ...
					if strings.HasPrefix(e, "--")       { continue } // Boundary string
					leadinbuff.WriteString(e + " ");      continue
				}

				// Diagnostic-Code: SMTP; 550-5.7.26 The MAIL FROM domain [email.example.jp]
				//    has an SPF record with a hard fail
				if strings.HasPrefix(e, " ") { b.WriteString(" " + moji.Sweep(e)) }
			}
		}
	}
	for recipients == 0 {
		// There is no valid recipient address, Try to use the alias addaress as a final recipient
		if rfc5322.IsEmailAddress(dscontents[0].Alias) == false { break }
		dscontents[0].Recipient = dscontents[0].Alias; recipients++
	}
	if recipients == 0 { return nil }

	beforemesg := ""; if leadinbuff.Len() > 0 {
		// Pick some values of []sis.DeliveryMatter{} from the string before startingof["message"]
		beforemesg           = moji.Sweep(leadinbuff.String())
		alternates.Command   = command.Find(beforemesg)
		alternates.ReplyCode = reply.Find(beforemesg, dscontents[0].Status)
		alternates.Status    = status.Find(beforemesg, alternates.ReplyCode)
	}
	issuedcode := strings.ToLower(beforemesg)

	for j := range dscontents {
		// Set default values stored in "permessage" if each value in "dscontents" is empty.
		e := &dscontents[j]; for _, z := range keystrings {
			// Do not set an empty string into each member of sis.DeliveryMatter{}
			if len(v.Select(z)) > 0 || len(permessage[z]) == 0 { continue }
			e.Update(z, permessage[z])
		}

		e.Diagnosis = moji.Sweep(eachbuffer[j].String())
		if recipients == 1 {
			// Do not mix the error message of each recipient with "beforemesg" when there is
			// multiple recipient addresses in the bounce message
			lowercased := strings.ToLower(e.Diagnosis)
			if strings.Contains(issuedcode, lowercased) == true {
				// "beforemesg" contains the entire strings of e.Diagnosis
				e.Diagnosis = beforemesg

			} else {
				// The value of e.Diagnosis is not contained in "beforemesg"
				// There may be an important error message in "beforemesg"
				e.Diagnosis = moji.Sweep(beforemesg + " " + e.Diagnosis)
			}
		}
		e.Command   = command.Find(e.Diagnosis);         if e.Command   == "" { e.Command   = alternates.Command   }
		e.ReplyCode = reply.Find(e.Diagnosis, e.Status); if e.ReplyCode == "" { e.ReplyCode = alternates.ReplyCode }

		if e.Status == "" { e.Status = status.Find(e.Diagnosis, e.ReplyCode) }
		if e.Status == "" { e.Status = alternates.Status                     }
	}
	return &sis.RisingUnderway{Digest: dscontents, RFC822: emailparts[1]}
}

