// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc3464

//  ____  _____ ____ _____ _  _    __   _  _   
// |  _ \|  ___/ ___|___ /| || |  / /_ | || |  
// | |_) | |_ | |     |_ \| || |_| '_ \| || |_ 
// |  _ <|  _|| |___ ___) |__   _| (_) |__   _|
// |_| \_\_|   \____|____/   |_|  \___/   |_|  
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/lhost"
import "sisimai/rfc1894"
import "sisimai/rfc2045"
import "sisimai/rfc5322"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

// Inquire() decodes a bounce message which have fields defined in RFC3464
func Inquire(bf *sis.BeforeFact) sis.RisingUnderway {
	// @param    *sis.BeforeFact bf  Message body of a bounce email
	// @return   RisingUnderway      RisingUnderway structure
	// @see      https://tools.ietf.org/html/rfc3464
	if len(bf.Head) == 0 { return sis.RisingUnderway{} }
	if len(bf.Body) == 0 { return sis.RisingUnderway{} }

	indicators := lhost.INDICATORS()
	boundaries := []string{
		// When the new value added, the part of the value should be listed in "delimiters" variable
		// defined at MakeFlat() function in sisimai/rfc2045/make-multipart-flat.go
		"Content-Type: message/rfc822",
		"Content-Type: text/rfc822-headers",
		"Content-Type: message/partial",
		"Content-Disposition: inline", // See lhost-amavis-*.eml, lhost-facebook-*.eml
	}
	startingof := map[string][]string{"message": []string{"Content-Type: message/delivery-status"}}

	for sisimoji.ContainsAny(bf.Body, boundaries) == false {
		// There is no "Content-Type: message/rfc822" line in the message body
		// Insert "Content-Type: message/rfc822" before "Return-Path:" of the original message
		cv := "\n\nReturn-Path:"; if strings.Contains(bf.Body, cv) == false { break }
		bf.Body = strings.Replace(bf.Body, cv, "\n\n" + boundaries[0] + cv, 1)
		break
	}
	fieldtable := rfc1894.FIELDTABLE()
	permessage := map[string]string{} // Store values of each Per-Message field
	keystrings := []string{}          // Key list of permessage
	dscontents := []sis.DeliveryMatter{{}}
	alternates := sis.DeliveryMatter{}
	emailparts := rfc5322.Part(&bf.Body, boundaries, false)
	readcursor := uint8(0)            // Points the current cursor position
	readslices := []string{""}        // Copy each line for later reference
	recipients := uint8(0)            // The number of 'Final-Recipient' header
	beforemesg := ""                  // String before startingof["message"]
	goestonext := false               // Flag: do not append the line into "beforemesg"
	isboundary := []string{rfc2045.Boundary(bf.Head["content-type"][0], 0)}
	v          := &(dscontents[len(dscontents) - 1])

	for strings.Contains(emailparts[0], "@") == false {
		// There is a bounce message inside of message/rfc822 part at lhost-x5-*
		cv := ""
		he := true
		p0 := strings.Index(bf.Body, boundaries[0] + "\n"); if p0 < 0 { break }
		for _, e := range strings.Split(bf.Body[p0 + 32:], "\n") {
			// Remove headers before the first "\n\n" after "Content-Type: message/rfc822" line
			if he { if e == "" { he = false }; continue }
			if strings.HasPrefix(e, "--")    { continue }
			cv += e + "\n"
		}
		emailparts = rfc5322.Part(&cv, boundaries, false)
		break;
	}

	for strings.Contains(emailparts[0], startingof["message"][0]) == false {
		// There is no "Content-Type: message/delivery-status" line in the message body
		// Insert "Content-Type: message/delivery-status" before "Reporting-MTA:" field
		cv := "\n\nReporting-MTA:"; if strings.Contains(emailparts[0], cv) == false { break }
		emailparts[0] = strings.Replace(emailparts[0], cv, "\n\n" + startingof["message"][0] + cv, 1)
		break
	}

	for j, e := range(strings.Split(emailparts[0], "\n")) {
		// Read error messages and delivery status lines from the head of the email to the
		// previous line of the beginning of the original message.
		readslices = append(readslices, e) // Save the current line for the next loop

		if readcursor == 0 {
			// Beginning of the bounce message or message/delivery-status part
			if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }

			for {
				// Append each string before startingof["message"][0] except the following patterns
				// for the later reference
				if e == ""    { break } // Blank line
				if goestonext { break } // Skip if the part is text/html, image/icon in multipart/*

				// This line is a boundary kept in "multiparts" as a string, when the end of
				// the boundary appeared, the condition above also returns true.
				if sisimoji.HasPrefixAny(e, isboundary) { goestonext = false; break }
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
				if strings.HasPrefix(e, "Content-")        { break } // Content-Disposition, ...
				if strings.HasPrefix(e, "This is a MIME")  { break } // This is a MIME-formatted message.
				if strings.HasPrefix(e, "This is a multi") { break } // This is a multipart message in MIME format
				if strings.HasPrefix(e, "This is an auto") { break } // This is an automatically generated ...
				if strings.HasPrefix(e, "This multi-part") { break } // This multi-part MIME message contains...
				if strings.HasPrefix(e, "###")             { break } // A frame like #####
				if strings.HasPrefix(e, "***")             { break } // A frame like *****
				if strings.Contains(e, "---- The follow")  { break } // ----- The following addresses had delivery problems -----
				if strings.Contains(e, "---- Transcript")  { break } // ----- Transcript of session follows -----
				beforemesg += e + " "; break
			}
			continue
		}
		if readcursor & indicators["deliverystatus"] == 0 { continue }
		if len(e) == 0                                    { continue }

		if f := rfc1894.Match(e); f > 0 {
			// This line matched with any field defined in RFC3464
			o := rfc1894.Field(e); if len(o) == 0 { continue }
			z := fieldtable[o[0]]
			v  = &(dscontents[len(dscontents) - 1])

			if o[3] == "addr" {
				// Final-Recipient: rfc822; kijitora@example.jp
				// X-Actual-Recipient: rfc822; kijitora@example.co.jp
				if o[0] == "final-recipient" {
					// Final-Recipient: rfc822; kijitora@example.jp
					// Final-Recipient: x400; /PN=...
					cv := sisiaddr.S3S4(o[2]); if sisiaddr.IsEmailAddress(cv) == false   { continue }
					cw := len(dscontents); if cw > 0 && cv == dscontents[cw-1].Recipient { continue }

					if len(v.Recipient) > 0 {
						// There are multiple recipient addresses in the message body.
						dscontents = append(dscontents, sis.DeliveryMatter{})
						v = &(dscontents[len(dscontents) - 1])
					}
					v.Recipient = cv
					recipients += 1

				} else {
					// X-Actual-Recipient: rfc822; kijitora@example.co.jp
					v.Alias = o[2]
				}
			} else if o[3] == "code" {
				// Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
				v.Spec       = o[1]
				v.Diagnosis += o[2] + " "

			} else {
				// Other DSN fields defined in RFC3464
				if o[4] != "" {
					// There are other error messages as a comment such as the following:
					// Status: 5.0.0 (permanent failure)
					// Status: 4.0.0 (cat.example.net: host name lookup failure)
					v.Diagnosis += " " + o[4]
				}
				v.Update(v.AsRFC1894(o[0]), o[2]); if f != 1 { continue }

				// Copy the lower-cased member name of sis.DeliveryMatter{} for "permessage" for
				// the later reference
				permessage[z] = o[2]
				if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }
			}
		} else {
			// Check that the line is a continued line of the value of Diagnostic-Code: field or not
			if strings.HasPrefix(e, "X-") && strings.Contains(e, ": ") {
				// This line is a MTA-Specific fields begins with "X-"
				if is3rdparty(e) == false { continue }
				if cv := xfield(e); len(cv) > 0 && rfc1894.Match(cv[0]) == 0 {
					// Check the first element is a field defined in RFC1894 or not
					if strings.HasPrefix(cv[4], "reason:") {
						// cv[4] is a string line "reason:mailboxfull"
						v.Update("reason", cv[4][strings.Index(cv[4], ":") + 1:])
					}
				} else {
					// Set the value picked from "X-*" field to the member of sis.DeliveryMatter
					// when the current value is empty
					z := fieldtable[strings.ToLower(cv[0])]; if len(z) < 1 { continue }
					if v.Select(z) == "" { v.Update(z, cv[2]) }
				}
			} else {
				// The line may be a continued line of the value of the Diagnostic-Code: field
				if strings.HasPrefix(readslices[j], "Diagnostic-Code:") == false {
					// In the case of multiple "message/delivery-status" line
					if strings.HasPrefix(e, "Content-") { continue } // Content-Disposition, ...
					beforemesg += e + " "
					continue
				}

				// Diagnostic-Code: SMTP; 550-5.7.26 The MAIL FROM domain [email.example.jp]
				//    has an SPF record with a hard fail
				if strings.HasPrefix(e, " ") == false { continue }
				v.Diagnosis += " " + sisimoji.Sweep(e)
			}
		}
	}
	for recipients == 0 {
		// There is no valid recipient address, Try to use the alias addaress as a final recipient
		if dscontents[0].Alias == ""                             { break }
		if sisiaddr.IsEmailAddress(dscontents[0].Alias) == false { break }
		dscontents[0].Recipient = dscontents[0].Alias; recipients++
	}
	if recipients == 0 { return sis.RisingUnderway{} }

	if beforemesg != "" {
		// Pick some values of []sis.DeliveryMatte{} from the string before startingof["message"]
		beforemesg           = sisimoji.Sweep(beforemesg)
		alternates.Command   = command.Find(beforemesg)
		alternates.ReplyCode = reply.Find(beforemesg, dscontents[0].Status)
		alternates.Status    = status.Find(beforemesg, alternates.ReplyCode)
	}
	issuedcode := strings.ToLower(beforemesg)

	for j, _ := range dscontents {
		// Set default values stored in "permessage" if each value in "dscontents" is empty.
		e := &(dscontents[j]); for _, z := range keystrings {
			// Do not set an empty string into each member of sis.DeliveryMatter{}
			if len(v.Select(z))    > 0 { continue }
			if len(permessage[z]) == 0 { continue }
			e.Update(z, permessage[z])
		}

		e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
		lowercased := strings.ToLower(e.Diagnosis)
		if strings.Contains(issuedcode, lowercased) == true {
			// "beforemesg" contains the entire strings of e.Diagnosis
			e.Diagnosis = beforemesg

		} else {
			// The value of e.Diagnosis is not contained in "beforemesg"
			// There may be an important error message in "beforemesg"
			e.Diagnosis = sisimoji.Sweep(beforemesg + " " + e.Diagnosis)
		}

		e.Command   = command.Find(e.Diagnosis);         if e.Command   == "" { e.Command   = alternates.Command   }
		e.ReplyCode = reply.Find(e.Diagnosis, e.Status); if e.ReplyCode == "" { e.ReplyCode = alternates.ReplyCode }

		if e.Status == "" { e.Status = status.Find(e.Diagnosis, e.ReplyCode) }
		if e.Status == "" { e.Status = alternates.Status                     }
	}

	// Set the recipient address as To: header in the original message part
	if emailparts[1] == "" { emailparts[1] = fmt.Sprintf("To: <%s>\n", dscontents[0].Recipient) }

	return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
}

