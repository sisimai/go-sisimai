// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      _______ _           _   _           _ _____ _           
// | | |__   ___  ___| |_   / / ____(_)_ __  ___| | | |_ __   __| | ____(_)_ __  ___ 
// | | '_ \ / _ \/ __| __| / /|  _| | | '_ \/ __| | | | '_ \ / _` |  _| | | '_ \/ __|
// | | | | | (_) \__ \ |_ / / | |___| | | | \__ \ |_| | | | | (_| | |___| | | | \__ \
// |_|_| |_|\___/|___/\__/_/  |_____|_|_| |_|___/\___/|_| |_|\__,_|_____|_|_| |_|___/
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import "sisimai/smtp/status"
import "sisimai/smtp/command"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

func init() {
	// Decode bounce messages from 1&1: https://www.1und1.de/
	InquireFor["EinsUndEins"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		if strings.HasPrefix(bf.Head["from"][0], `"Mail Delivery System"`) == false  { return sis.RisingUnderway{} }
		if bf.Head["from"][0] != "Mail delivery failed: returning message to sender" { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"--- The header of the original message is following. ---"}
		startingof := map[string][]string{
			"message": []string{"This message was created automatically by mail delivery software"},
			"error":   []string{"For the following reason:"},
		}
		messagesof := map[string][]string{"mesgtoobig": []string{"Mail size limit exceeded"}}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)            // Points the current cursor position
		recipients := uint8(0)            // The number of 'Final-Recipient' header
		alternates := ""                  // Alternative error message
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e) == 0                                    { continue }

			// The following address failed:
			//
			// general@example.eu
			//
			// For the following reason:
			//
			// Mail size limit exceeded. For explanation visit
			// http://postmaster.1and1.com/en/error-messages?ip=%1s
			if sisimoji.Aligned(e, []string{"@", "."}) {
				// general@example.eu OR
				// the line begin with 4 space characters, end with ":" like "    neko@example.eu:"
				ce := sisiaddr.S3S4(e); if sisiaddr.IsEmailAddress(ce) == false { continue }
				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = ce
				recipients += 1

			} else if strings.HasPrefix(e, startingof["error"][0]) {
				// For the following reason:
				v.Diagnosis = e

			} else {
				//  Get error message and append the error message strings
				if v.Diagnosis != "" { v.Diagnosis += " " + e; continue }

				// OR the following format:
				//   neko@example.fr:
				//   SMTP error from remote server for TEXT command, host: ...
				alternates += " " + e
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Get and set other values into sis.DeliveryMatter{}
			e := &(dscontents[j])

			if e.Diagnosis == "" { e.Diagnosis = alternates }
			e.Command = command.Find(e.Diagnosis)

			if sisimoji.Aligned(e.Diagnosis, []string{"host: ", " reason:"}) {
				// SMTP error from remote server for TEXT command,
				//   host: smtp-in.orange.fr (193.252.22.65)
				//   reason: 550 5.2.0 Mail rejete. Mail rejected. ofr_506 [506]
				p1 := strings.Index(e.Diagnosis, "host: ")
				e.Rhost  = sisimoji.Sweep(strings.Split(e.Diagnosis[p1:], " ")[1])
				e.Status = status.Find(e.Diagnosis, "")

				if strings.Contains(e.Diagnosis, "for TEXT command") { e.Command = "DATA" }
				if strings.Contains(e.Diagnosis, "SMTP error")       { e.Spec    = "SMTP" }

			} else {
				// Remvoe "For the following reason:" string from e.Diagnosis
				p1 := strings.Index(e.Diagnosis, startingof["error"][0])
				p2 := len(startingof["error"][0])
				if p1 > -1 { e.Diagnosis = e.Diagnosis[p1 + p2:] }
			}
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			FINDREASON: for r := range messagesof {
				// The key name is a bounce reason name
				for _, f := range messagesof[r] {
					// Try to find an error message including lower-cased string listed in messagesof
					if strings.Contains(e.Diagnosis, f) == false { continue }
					e.Reason = r; break FINDREASON
				}
			}
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

