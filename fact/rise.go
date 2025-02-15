// Copyright (C) 2020-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//   __            _     ____  _          
//  / _| __ _  ___| |_  |  _ \(_)___  ___ 
// | |_ / _` |/ __| __| | |_) | / __|/ _ \
// |  _| (_| | (__| |_ _|  _ <| \__ \  __/
// |_|  \__,_|\___|\__(_)_| \_\_|___/\___|

// Package "fact" provide a function for generating structs keeping decoded bounce mail data
package fact

import "time"
import "strings"
import "net/mail"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/lda"
import "libsisimai.org/sisimai/rhost"
import "libsisimai.org/sisimai/reason"
import "libsisimai.org/sisimai/message"
import "libsisimai.org/sisimai/rfc791"
import "libsisimai.org/sisimai/rfc1123"
import "libsisimai.org/sisimai/rfc1894"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/smtp/reply"
import "libsisimai.org/sisimai/smtp/status"
import "libsisimai.org/sisimai/smtp/command"
import "libsisimai.org/sisimai/smtp/failure"
import sisiaddr "libsisimai.org/sisimai/address"
import sisimoji "libsisimai.org/sisimai/string"

var RetryIndex = reason.Retry()
var RFC822Head = rfc5322.HEADERTABLE()
var ActionList = map[string]bool{ "delayed": true, "delivered": true, "expanded": true, "failed": true, "relayed": true }

// sisimai/fact.Rise() returns []sis.Fact when it successfully decoded bounce messages
func Rise(email *string, origin string, args *sis.DecodingArgs) ([]sis.Fact, []sis.NotDecoded) {
	// @param  *string           email    Entire email message
	// @param  string            origin   Path to the original email file
	// @param  *sis.DecodingArgs args     Arguments for decoding(delivered, vacation, callbacks)
	// @return []sis.Fact               The list of decoded bounce messages
	if email == nil || len(*email) < 1 {
		// The email message is empty
		ce := *sis.MakeNotDecoded("email file is empty", true); ce.Email(origin)
		return []sis.Fact{}, []sis.NotDecoded{ce}
	}

	beforefact := message.Rise(email, args.Callback0); if len((*beforefact).Errors) > 0 {
		// There is some errors while reading the email, decoding the bounce message.
		// Set the email path to sis.NotDecoded.EmailFile
		for j := range (*beforefact).Errors { (*beforefact).Errors[j].Email(origin) }
	}
	if (*beforefact).Void() == true { return []sis.Fact{}, (*beforefact).Errors }

	rfc822data := (*beforefact).RFC822
	listoffact := []sis.Fact{}

	RISEOF: for _, e := range (*beforefact).Digest {
		// Create parameters for sis.Fact
		// - Skip if the value of "recipient" length is 4 or shorter
		// - Skip if the value of "deliverystatus" begins with "2." such as 2.1.5
		// - Skip if the value of "reason" is "vacation"
		if rfc5322.IsEmailAddress(e.Recipient) == false                { continue RISEOF }
		if args.Delivered != true && strings.HasPrefix(e.Status, "2.") { continue RISEOF }
		if args.Vacation  != true && e.Reason == "vaction"             { continue RISEOF }

		addrs := map[string][3]string{} // Addresser, and Recipient
		piece := map[string]string{}    // Each element except email addresses
		thing := sis.Fact{}             // Each sis.Fact struct
		clock := time.Time{}            // The source value of "Timestamp"

		ADDRESSER: for {
			// Detect an email address from message/rfc822 part
			for _, f := range RFC822Head["addresser"] {
				// Check each header in message/rfc822 part
				if len(rfc822data[f])                         == 0  { continue }
				j := sisiaddr.Find(rfc822data[f][0]); if j[0] == "" { continue }
				addrs["addresser"] = j; break ADDRESSER
			}

			if len(addrs["addresser"][0]) == 0 && len((*beforefact).Headers["to"]) > 0 {
				// Fallback: Get the sender address from the header of the bounced email if the address
				// is not set at the loop above.
				j := sisiaddr.Find((*beforefact).Headers["to"][0])
				if j[0] != "" { addrs["addresser"] = j }
			}
			break ADDRESSER
		}
		if len(addrs["addresser"][0]) == 0 { continue RISEOF }

		TIMESTAMP: for {
			// Convert from the value of "Date" or the date string to time.Time
			datevalues := []string{}; if e.Date != "" { datevalues = append(datevalues, e.Date) }

			for _, f := range RFC822Head["date"] {
				// Date information did not exist in message/delivery-status part.
				// Get the value of "Date:" header or other date related headers.
				if len(rfc822data[f]) > 0 { datevalues = append(datevalues, rfc822data[f][0]) }
			}

			if len(datevalues) < 2 && len((*beforefact).Headers["date"]) > 0 {
				// Get the value of "Date:" header of the bounce message
				datevalues = append(datevalues, (*beforefact).Headers["date"][0])
			}
			for _, v := range datevalues {
				// Parse each date string using net/mail.ParseDate()
				if times, nyaan := mail.ParseDate(v); nyaan == nil { clock = times; break }
			}
			if clock.IsZero() {
				// Failed to parse the date string at the previous loop,
				// try to tidy up it using rfc5322.Date() before calling net/mail.ParseDate()
				for _, v := range datevalues {
					// Try to parse the date string tidied by rfc5322.Date()
					j := rfc5322.Date(v); if j != "" {
						// rfc5322.Date() returned a valid date string
						if times, nyaan := mail.ParseDate(j); nyaan == nil { clock = times; break }
					}
				}
			}
			if clock.IsZero() { continue RISEOF }
			break TIMESTAMP
		}

		RECEIVED: for {
			// Try to pick a remote hostname from the error message
			// Scan "Received:" header of the bounce message
			le := len((*beforefact).Headers["received"])
			if e.Rhost == "" {
				// Try to pick a remote hostname from Received: headers of the bounce message
				if cv := rfc1123.Find(e.Diagnosis); rfc1123.IsInternetHost(cv) { e.Rhost = cv }
				if e.Rhost == "" {
					// The remote hostname in the error message did not exist or is not a valid
					// internet hostname
					for ri := le - 1; ri > -1; ri-- {
						// Check the Received: headers backwards and get a remote hostname
						cv := rfc5322.Received((*beforefact).Headers["received"][ri])
						if rfc1123.IsInternetHost(cv[0]) == false { continue }
						e.Rhost = cv[0]; break
					}
				}
			}
			if e.Lhost == e.Rhost { e.Lhost = "" }
			if e.Lhost == "" {
				// Try to pick a local hostname from Received: headers of the bounce message
				for li := 0; li < le; li++ {
					// Check the Received: headers forwards and get a local hostnaame
					cv := rfc5322.Received((*beforefact).Headers["received"][li])
					if rfc1123.IsInternetHost(cv[0]) { e.Lhost = cv[0]; break }
				}
			}

			for _, v := range []*string{&e.Rhost, &e.Lhost} {
				// Check and rewrite each host name
				if *v == "" { continue }

				// Use the domain part as a remote/local host when the value is an email address
				if strings.Contains(*v, "@") { *v = strings.Split(*v, "@")[1] }

				// Remove [], (), \r, and strings before "="
				for _, c := range []string{"(", ")", "[", "]", "\r"} { *v = strings.ReplaceAll(*v, c, "") }
				if strings.Contains(*v, "=") { *v = strings.SplitN(*v, "=", 2)[1] }
				if strings.Contains(*v, " ") {
					// Check a space character in each value and get the first hostname
					ee := strings.Split(*v, " ")
					for _, w := range ee {
						// Get a hostname from the string like "127.0.0.1 x109-20.example.com 192.0.2.20"
						// or "mx.sp.example.jp 192.0.2.135"
						if rfc791.IsIPv4Address(w) == false { *v = w; break }
					}
					if strings.Index(*v, " ") > 0 { *v = ee[0] }
				}
				if strings.HasSuffix(*v, ".") { *v = strings.TrimRight(*v, ".") } // Remove "." at the end of the value
			}
			break RECEIVED
		}

		MESG_ID: for {
			// https://www.rfc-editor.org/rfc/rfc5322#section-3.6.4
			// Leave only string inside of angle brackets(<>)
			if len(rfc822data["message-id"])                                          == 0     { break MESG_ID }
			if sisimoji.Aligned(rfc822data["message-id"][0], []string{"<", "@", ">"}) == false { break MESG_ID }

			piece["messageid"] = strings.Trim(rfc822data["message-id"][0], "<>"); break MESG_ID
		}

		LIST_ID: for {
			// https://www.rfc-editor.org/rfc/rfc2919
			// Get the value of List-Id header: "List name <list-id@example.org>"
			if len(rfc822data["list-id"])                                          == 0     { break LIST_ID }
			if sisimoji.Aligned(rfc822data["list-id"][0], []string{"<", ".", ">"}) == false { break LIST_ID }
			piece["listid"] = sisimoji.Select(rfc822data["list-id"][0], "<", ">", 0);         break LIST_ID
		}

		DIAGNOSTICCODE: for {
			// - Cleanup the value of "Diagnostic-Code:" field
			// - Find and set the SMTP Reply Code
			piece["diagnosticcode"] = e.Diagnosis
			piece["deliverystatus"] = e.Status
			piece["replycode"]      = e.ReplyCode
			if e.Diagnosis == "" { break DIAGNOSTICCODE }

			// Get an SMTP Reply Code and an SMTP Enhanced Status Code
			piece["diagnosticcode"] = strings.ReplaceAll(piece["diagnosticcode"], "\r", "")
			cs := status.Find(piece["diagnosticcode"], "")
			cr := reply.Find(piece["diagnosticcode"], cs)
			piece["deliverystatus"] = status.Prefer(piece["deliverystatus"], cs, cr)

			if len(cr) == 3 {
				// There is an SMTP reply code in the error message 
				if piece["replycode"] == "" { piece["replycode"] = cr }
				if strings.Contains(piece["diagnosticcode"], cr + "-") {
					// 550-5.7.1 [192.0.2.222] Our system has detected that this message is
					// 550-5.7.1 likely unsolicited mail. To reduce the amount of spam sent to Gmail,
					// 550-5.7.1 this message has been blocked. Please visit
					// 550 5.7.1 https://support.google.com/mail/answer/188131 for more information.
					//
					// kijitora@example.co.uk
					//   host c.eu.example.com [192.0.2.3]
					//   SMTP error from remote mail server after end of data:
					//   553-SPF (Sender Policy Framework) domain authentication
					//   553-fail. Refer to the Troubleshooting page at
					//   553-http://www.symanteccloud.com/troubleshooting for more
					//   553 information. (#5.7.1)
					for _, q := range []string{"-", " "} {
						// Remove strings: "550-5.7.1", and "550 5.7.1" from the error message
						cx := cr + q + cs
						piece["diagnosticcode"] = strings.ReplaceAll(piece["diagnosticcode"], cx, "")

						// Remove "553-" and "553 " (SMTP reply code only) from the error message
						cx  = cr + q
						piece["diagnosticcode"] = strings.ReplaceAll(piece["diagnosticcode"], cx, "")
					}

					if strings.Contains(piece["diagnosticcode"], cr) {
						// Add "550 5.1.1" into the head of the error message when the error
						// message does not begin with "550"
						piece["diagnosticcode"] = cr + " " + cs + " " + piece["diagnosticcode"]
					}
				}
			}

			dc := strings.ToLower(piece["diagnosticcode"])
			p1 := strings.Index(dc, "<html>")
			p2 := strings.Index(dc, "</html>")
			if p1 > 0 && p2 > 0 {
				// Remove strings from <html> to </html>
				piece["diagnosticcode"] = piece["diagnosticcode"][:p1] + " " + piece["diagnosticcode"][p2 + 7:]
			}

			piece["diagnosticcode"] = sisimoji.Sweep(piece["diagnosticcode"])
			break DIAGNOSTICCODE
		}

		DIAGNOSTICTYPE: for {
			// Set the value of "diagnostictype" if it is empty
			piece["diagnostictype"] = e.Spec
			piece["reason"]         = e.Reason

			if e.Spec != "" { break DIAGNOSTICTYPE }
			if piece["reason"] == "mailererror"                               { piece["diagnostictype"] = "X-UNIX" }
			if piece["reason"] != "feedback" && piece["reason"] != "vacation" { piece["diagnostictype"] = "SMTP"   }
			break DIAGNOSTICTYPE
		}

		// Set other values returned from sisimai/message.Rise()
		addrs["recipient"] = [3]string{e.Recipient, "", ""}
		piece["subject"]   = strings.ReplaceAll(rfc822data["subject"][0], "\r", "")
		if command.Test(e.Command) { piece["command"] = e.Command }

		CONSTRUCTOR: for {
			// - Create email address object as address.EmailAddress struct
			// - Create decoded bounce mail object as sis.Fact struct
			as := sisiaddr.Rise(addrs["addresser"]); if as.Void() == true { continue RISEOF }
			ar := sisiaddr.Rise(addrs["recipient"]); if ar.Void() == true { continue RISEOF }

			thing.Action         = e.Action
			thing.Addresser      = as
			thing.Alias          = e.Alias; if thing.Alias == "" { thing.Alias = ar.Alias }
			thing.Catch          = (*beforefact).Catch
			thing.DeliveryStatus = piece["deliverystatus"]
			thing.Destination    = ar.Host
			thing.DiagnosticCode = piece["diagnosticcode"]
			thing.DiagnosticType = piece["diagnostictype"]
			thing.FeedbackID     = ""
			thing.FeedbackType   = e.FeedbackType
			thing.HardBounce     = false
			thing.Lhost          = e.Lhost
			thing.ListID         = piece["listid"]
			thing.MessageID      = piece["messageid"]
			thing.Origin         = origin
			thing.Reason         = piece["reason"]
			thing.Rhost          = e.Rhost
			thing.Recipient      = ar
			thing.ReplyCode      = piece["replycode"]; if thing.ReplyCode == "" { reply.Find(piece["diagnosticcode"], "") }
			thing.DecodedBy      = e.Agent
			thing.Command        = piece["command"]
			thing.SenderDomain   = as.Host
			thing.Subject        = piece["subject"]
			thing.Timestamp      = clock
			thing.TimezoneOffset = clock.Format("+0900")
			thing.Token          = sisimoji.Token(as.Address, ar.Address, int(thing.Timestamp.Unix()))

			break CONSTRUCTOR
		}

		ALIAS: for {
			// Look up the Envelope-To address from the Received: header in the original message
			// when the recipient address is same with the value of piece["alias"].
			if thing.Recipient.Address != thing.Alias                { break ALIAS }
			if thing.Alias == "" || len(rfc822data["received"]) == 0 { break ALIAS }

			recv := rfc822data["received"]
			hops := len(recv)
			for i := hops - 1; hops >= 0; hops-- {
				// Search for the string " for " from the Received: header
				if strings.Index(recv[i], " for ") == -1  { continue }
				or := rfc5322.Received(recv[i])

				if len(or) == 0 || or[5] == ""            { continue }
				if rfc5322.IsEmailAddress(or[5]) == false { continue }
				if or[5] == thing.Recipient.Address       { continue }

				thing.Alias = or[5]; break
			}
			break ALIAS
		}
		if thing.Alias == thing.Recipient.Address { thing.Alias = "" }

		REASON: for thing.Reason == "" || RetryIndex[thing.Reason] {
			// Decide the reason of the email bounce
			// The value of thing.Reason is empty or is needed to check with other values again
			re := thing.Reason;        if re == ""              { re = "undefined"                }
			or := lda.Find(&thing);    if reason.IsExplicit(or) { thing.Reason = or; break REASON }
			or  = rhost.Find(&thing);  if reason.IsExplicit(or) { thing.Reason = or; break REASON }
			or  = reason.Find(&thing); if reason.IsExplicit(or) { thing.Reason = or; break REASON }

			if thing.DiagnosticCode != "" { re = "onhold" }
			thing.Reason = re; break REASON
		}

		HARDBOUNCE: for {
			// Set the value of "hardbounce", default value of "bouncebounce" is 0
			if thing.Reason == "delivered" || thing.Reason == "feedback" || thing.Reason == "vacation" {
				// Delete the value of ReplyCode when the Reason is "feedback" or "vacation"
				if thing.Reason != "delivered" { thing.ReplyCode = "" }

			} else {
				// The Reason is not "delivered", or "feedback", or "vacation"
				cv := piece["deliverystatus"] + " " + piece["diagnosticcode"]; if len(cv) < 4 { cv = "" }
				thing.HardBounce = failure.IsHardBounce(thing.Reason, cv)
			}
			break HARDBOUNCE
		}

		DELIVERYSTATUS: for {
			// Set a pseudo status code
			if thing.DeliveryStatus != "" { break DELIVERYSTATUS }

			ce := thing.ReplyCode + " " + piece["diagnosticcode"]; if len(ce) < 4 { ce = "" }
			permanent0 := failure.IsPermanent(ce)
			temporary0 := failure.IsTemporary(ce)
			temporary1 := temporary0; if !permanent0 && !temporary0 { temporary1 = false }
			thing.DeliveryStatus = status.Code(thing.Reason, temporary1)
			break DELIVERYSTATUS
		}

		REPLYCODE: for {
			// Check both of the first digit of "DeliveryStatus" and "ReplyCode"
			cx := [2]string{"", ""}
			if thing.DeliveryStatus != "" { cx[0] = string(thing.DeliveryStatus[0]) }
			if thing.ReplyCode      != "" { cx[1] = string(thing.ReplyCode[0])      }

			if cx[0] != cx[1] {
				// The class of the "Status:" is defer with the first digit of the reply code
				cx[1] = reply.Find(piece["diagnosticcode"], cx[0])
				if strings.HasPrefix(cx[1], cx[0]) {
					// The first digit of cx[1] found by status.Find() is equal to cx[0]
					thing.ReplyCode = cx[1]

				} else {
					// Remove the value of ReplyCode when the 1st digit of the both values are difer
					thing.ReplyCode = ""
				}
			}

			if ActionList[thing.Action] == false {
				// There is an action value that is not described at RFC1894
				if ox := rfc1894.Field("Action: " + thing.Action); len(ox) > 0 {
					// Rewrite the value of "Action:" field to the valid value
					//
					// The syntax for the action-field is:
					//   action-field = "Action" ":" action-value
					//   action-value = "failed" / "delayed" / "delivered" / "relayed" / "expanded"
					thing.Action = ox[2]
				}
			}
			if thing.Reason == "delivered"                          { thing.Action = "delivered" }
			if thing.Reason == "expired"                            { thing.Action = "delayed"   }
			if thing.Action == "" && (cx[0] == "4" || cx[0] == "5") { thing.Action = "failed"    }

			break REPLYCODE
		}

		// Feedback-ID: 1.us-west-2.QHuyeCQrGtIIMGKQfVdUhP9hCQR2LglVOrRamBc+Prk=:AmazonSES
		if len(rfc822data["feedback-id"]) > 0 { thing.FeedbackID = rfc822data["feedback-id"][0] }

		listoffact = append(listoffact, thing)
	}

	if len((*beforefact).Errors) > 0 {
		// There is some errors while reading the email, decoding the bounce message.
		// Set the email path to sis.NotDecoded.EmailFile if it is empty
		for j := range (*beforefact).Errors { (*beforefact).Errors[j].Email(origin) }
	}
	return listoffact, beforefact.Errors
}

