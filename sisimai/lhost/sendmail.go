// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost
//  _ _               _      __                  _                 _ _ 
// | | |__   ___  ___| |_   / /__  ___ _ __   __| |_ __ ___   __ _(_) |
// | | '_ \ / _ \/ __| __| / / __|/ _ \ '_ \ / _` | '_ ` _ \ / _` | | |
// | | | | | (_) \__ \ |_ / /\__ \  __/ | | | (_| | | | | | | (_| | | |
// |_|_| |_|\___/|___/\__/_/ |___/\___|_| |_|\__,_|_| |_| |_|\__,_|_|_|
//                                                                     
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// V8Sendmail: /usr/sbin/sendmail
	InquireFor["Sendmail"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Body)            == 0 { return sis.RisingUnderway{} }
		if len(bf.Head["x-aol-ip"]) > 0 { return sis.RisingUnderway{} } // X-AOL-IP is a header defined in AOL

		match := false
		if strings.HasPrefix(bf.Head["subject"][0], "Warning: ") ||
		   strings.HasSuffix(bf.Head["subject"][0], "see transcript for details") {
			// Subject: Warning: could not send message for past 4 hours
			// Subject: Returned mail: see transcript for details
			match = true
		}
		if match == false { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"}
		startingof := map[string][]string{
			// savemail.c:1040|if (printheader && !putline("   ----- Transcript of session follows -----\n",
			// savemail.c:1041|          mci))
			// savemail.c:1042|  goto writeerr;
			// savemail.c:1360|if (!putline(
			// savemail.c:1361|    sendbody
			// savemail.c:1362|    ? "   ----- Original message follows -----\n"
			// savemail.c:1363|    : "   ----- Message header follows -----\n",
			"message": []string{"   ----- Transcript of session follows -----"},
			"error":   []string{"... while talking to "},
		}
		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		readslices := []string{""}        // Copy each line for later reference
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		thecommand := ""                  // An SMTP command name begins with the string ">>>"
		esmtpreply := []string{}          // Reply messages from the remote server on an SMTP session
		sessionerr := false               // Flag, true if it is an SMTP session error
		anotherset := map[string]string{} // Another error information
		v          := &(dscontents[len(dscontents) - 1])

		for j, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			readslices = append(readslices, e) // Save the current line for the next loop

			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }

			f := rfc1894.Match(e); if f > 0 {
				// "e" matched with any field defined in RFC3464
				o := rfc1894.Field(e); if len(o) == 0 { continue }
				z := fieldtable[o[0]]
				v  = &(dscontents[len(dscontents) - 1])

				if o[3] == "addr" {
					// Final-Recipient: rfc822; kijitora@example.jp
					// X-Actual-Recipient: rfc822; kijitora@example.co.jp
					if o[0] == "final-recipient" {
						// Final-Recipient: rfc822; kijitora@example.jp
						if len(v.Recipient) > 0 {
							// There are multiple recipient addresses in the message body.
							dscontents = append(dscontents, sis.DeliveryMatter{})
							v = &(dscontents[len(dscontents) - 1])
						}
						v.Recipient = o[2]
						recipients += 1

					} else {
						// X-Actual-Recipient: rfc822; kijitora@example.co.jp
						v.Alias = o[2]
					}
				} else if o[3] == "code" {
					// # Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
					v.Spec = o[1]
					v.Diagnosis = o[2]

				} else {
					// Other DSN fields defined in RFC3464
					v.Set(o[0], o[2]); if f != 1 { continue }

					// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
					permessage[z] = o[2]
					if sisimoji.EqualsAny(z, keystrings) == false { keystrings = append(keystrings, z) }
				}
			} else {
				// The line does not begin with a DSN field defined in RFC3464
				//
				// ----- Transcript of session follows -----
				// ... while talking to mta.example.org.:
				// >>> DATA
				// <<< 550 Unknown user recipient@example.jp
				// 554 5.0.0 Service unavailable
				if strings.HasPrefix(e, " ") == false {
					// Other error messages which does not start with " "
					if strings.HasPrefix(e, ">>> ") {
						// >>> DATA (Client Command)
						thecommand = command.Find(e)

					} else if strings.HasPrefix(e, "<<< ") {
						// <<< Response from SMTP Server
						isincluded := false
						for _, r := range esmtpreply {
							// Whether the response is included in the "estmpreply" or not
							if strings.Contains(r, e[4:]) == true { isincluded = true; break }
						}
						if isincluded == false { esmtpreply = append(esmtpreply, e[4:]) }

					} else {
						// Detect an SMTP session error or a connection error
						if sessionerr == true { continue }
						if strings.HasPrefix(e, startingof["error"][0]) {
							// ----- Transcript of session follows -----
							// ... while talking to mta.example.org.:
							sessionerr = true
							continue
						}

						if strings.HasPrefix(e, "<") && sisimoji.Aligned(e, []string{"@", ">.", " "}) {
							// <kijitora@example.co.jp>... Deferred: Name server: example.co.jp.: host name lookup failure
							anotherset["recipient"] = sisiaddr.S3S4(e[0:strings.Index(e, ">")])
							anotherset["diagnosis"] = e[strings.Index(e," ") + 1:]

						} else {
							// ----- Transcript of session follows -----
							// Message could not be delivered for too long
							// Message will be deleted from queue
							cr := reply.Find(e, "")
							cs := status.Find(e, "")

							if len(cr + cs) > 7 {
								// 550 5.1.2 <kijitora@example.org>... Message
								//
								// DBI connect('dbname=...')
								// 554 5.3.0 unknown mailer error 255
								anotherset["status"]     = cs
								anotherset["diagnosis"] += " " + e

							} else if strings.HasPrefix(e, "Message: ") || strings.HasPrefix(e, "Warning: ") {
								// Message could not be delivered for too long
								// Warning: message still undelivered after 4 hours
								anotherset["diagnosis"] += " " + e
							}
						}
					}
				} else {
					// Get the error message continued from the previous line
					if strings.HasPrefix(e, " ")                            == false { continue }
					if strings.HasPrefix(readslices[j], "Diagnostic-Code:") == false { continue }
					v.Diagnosis += " " + sisimoji.Sweep(e)
				}
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Set default values stored in "permessage" if each value in "dscontents" is empty.
			e := &(dscontents[j])
			for _, z := range keystrings {
				// Do not set an empty string into each member of DeliveryMatter{}
				if len(v.Get(z))       > 0 { continue }
				if len(permessage[z]) == 0 { continue }
				e.Set(z, permessage[z])
			}

			if len(anotherset["diagnosis"]) > 0 {
				// Copy alternative error message to e.Diagnosis
				if strings.HasPrefix(e.Diagnosis, " ")       { e.Diagnosis = anotherset["diagnosis"] }
				if sisimoji.ContainsOnlyNumbers(e.Diagnosis) { e.Diagnosis = anotherset["diagnosis"] } 
				if len(e.Diagnosis) == 0                     { e.Diagnosis = anotherset["diagnosis"] } 
			}

			for {
				// Replace or append the error message in "diagnosis" with the ESMTP Reply Code
				// when the following conditions have matched
				if len(esmtpreply) == 0 { break }
				if recipients != 1      { break }

				e.Diagnosis = fmt.Sprintf("%s %s", strings.Join(esmtpreply, " "), e.Diagnosis)
				break
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)
			if e.Command == "" { e.Command = thecommand }
			if e.Command == "" { e.Command = command.Find(e.Diagnosis) }
			if e.Command == "" { if len(esmtpreply) > 0 { e.Command = "EHLO" }}

			for {
				// Check alternative status code and override it
				if len(anotherset["status"]) == 0 { break }
				if status.Test(e.Status) == true  { break }

				e.Status = anotherset["status"]
				break
			}

			if strings.HasSuffix(e.Recipient, "@") {
				// @example.jp, no local part
				// Get the email address from the value of Diagnostic-Code field
				cv := sisiaddr.Find(e.Diagnosis)
				if cv[0] != "" { e.Recipient = cv[0] }
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
	}
}

