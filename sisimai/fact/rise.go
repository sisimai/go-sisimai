// Copyright (C) 2020-2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//   __            _   
//  / _| __ _  ___| |_ 
// | |_ / _` |/ __| __|
// |  _| (_| | (__| |_ 
// |_|  \__,_|\___|\__|
import "fmt"
import "time"
import "strings"
import "net/mail"
import "sisimai/sis"
import "sisimai/rhost"
import "sisimai/reason"
import "sisimai/message"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import "sisimai/smtp/command"
import "sisimai/smtp/failure"
import sisiaddr "sisimai/address"
import sisimoji "sisimai/string"

var RetryIndex = reason.Retry()
var RFC822Head = rfc5322.HEADERTABLE()
var ActionList = map[string]bool{ "delayed": true, "delivered": true, "expanded": true, "failed": true, "relayed": true };

// sisimai/fact.Rise() returns []sis.Fact when it successfully decoded bounce messages
func Rise(email *string, origin string, args map[string]bool, hook *func()) []sis.Fact {
	// @param  *string         email    Entire email message
	// @param  string          origin   Path to the original email file
	// @param  map[string]bool args     {"delivered": false, "vacation": false} as the default
	// @param  *func()         hook     The pointer to the callback function
	// @return []sis.Fact               The list of decoded bounce messages
	if len(*email) < 1 { return []sis.Fact{} }

	beforefact := message.Rise(email, hook); if beforefact.Void() == true { return []sis.Fact{} }
	rfc822data := beforefact.RFC822
	listoffact := []sis.Fact{}

	RISEOF: for _, e := range beforefact.Digest {
		// Create parameters for sis.Fact
		// - Skip if the value of "recipient" length is 4 or shorter
		// - Skip if the value of "deliverystatus" begins with "2." such as 2.1.5
		// - Skip if the value of "reason" is "vacation"
		if len(e.Recipient)   < 5                                         { continue RISEOF }
		if args["delivered"] != true && strings.HasPrefix(e.Status, "2.") { continue RISEOF }
		if args["vacation"]  != true && e.Reason == "vaction"             { continue RISEOF }

		addrs := map[string][3]string{} // Addresser, and Recipient
		piece := map[string]string{}    // Each element except email addresses
		thing := sis.Fact{}             // Each sis.Fact struct
		clock := time.Time{}            // The source value of "Timestamp"

		ADDRESSER: for {
			// Detect an email address from message/rfc822 part
			for _, f := range RFC822Head["addresser"] {
				// Check each header in message/rfc822 part
				if len(rfc822data[f])                              == 0 { continue }
				j := sisiaddr.Find(rfc822data[f][0]); if len(j[0]) == 0 { continue }
				addrs["addresser"] = j
				break ADDRESSER
			}

			if len(addrs["addresser"][0]) == 0 && len(beforefact.Head["to"]) > 0 {
				// Fallback: Get the sender address from the header of the bounced email if the address
				// is not set at the loop above.
				j := sisiaddr.Find(beforefact.Head["to"][0])
				if len(j[0]) > 0 { addrs["addresser"] = j }
			}
			break ADDRESSER
		}
		if len(addrs["addresser"][0]) == 0 { continue RISEOF }

		TIMESTAMP: for {
			// Convert from the value of "Date" or the date string to time.Time
			datevalues := []string{}

			if len(e.Date) > 0 { datevalues = append(datevalues, e.Date) }
			for _, f := range RFC822Head["date"] {
				// Date information did not exist in message/delivery-status part.
				// Get the value of "Date:" header or other date related headers.
				if len(rfc822data[f]) == 0 { continue }
				datevalues = append(datevalues, rfc822data[f][0])
			}

			// Get the value of "Date:" header of the bounce message
			if len(datevalues) < 2 { datevalues = append(datevalues, beforefact.Head["date"][0]) }
			for _, v := range datevalues {
				// Parse each date string using time.Parse()
				times, nyaan := mail.ParseDate(v); if nyaan != nil { continue }
				clock = times
				break
			}
			if clock.IsZero() == true { continue RISEOF }

			break TIMESTAMP
		}

		RECEIVED: for {
			// Scan "Received:" header of the original message
			recv := beforefact.Head["received"];
			size := len(recv)

			if size > 0 {
				// Get a local host name and a remote host name from the Received header.
				if len(e.Rhost) == 0  { e.Rhost = rfc5322.Received(recv[size - 1])[0] }
				if e.Lhost == e.Rhost { e.Lhost = "" }
				if len(e.Lhost) == 0  { e.Lhost = rfc5322.Received(recv[0])[0] }
			}
			for _, v := range []*string{&e.Rhost, &e.Lhost} {
				// Check and rewrite each host name
				if len(*v) == 0 { continue }

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
						if sisimoji.IsIPv4Address(w) { continue }
						*v = w; break
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

			piece["messageid"] = strings.Trim(rfc822data["message-id"][0], "<>")
			break MESG_ID
		}

		LIST_ID: for {
			// https://www.rfc-editor.org/rfc/rfc2919
			// Get the value of List-Id header: "List name <list-id@example.org>"
			if len(rfc822data["list-id"])                                          == 0     { break LIST_ID }
			if sisimoji.Aligned(rfc822data["list-id"][0], []string{"<", ".", ">"}) == false { break LIST_ID }
			p0 := strings.Index(rfc822data["list-id"][0], "<");                   if p0 < 0 { break LIST_ID }
			p1 := strings.Index(rfc822data["list-id"][0], ">");                   if p1 < 0 { break LIST_ID }

			piece["listid"] = rfc822data["list-id"][0][p0 + 1:p1]
			break LIST_ID
		}

		DIAGNOSTICCODE: for {
			// - Cleanup the value of "Diagnostic-Code:" field
			// - Find and set the SMTP Reply Code
			piece["diagnosticcode"] = e.Diagnosis
			piece["deliverystatus"] = e.Status
			piece["replycode"]      = e.ReplyCode
			if len(e.Diagnosis) == 0 { break DIAGNOSTICCODE }

			// Get an SMTP Reply Code and an SMTP Enhanced Status Code
			piece["diagnosticcode"] = strings.ReplaceAll(piece["diagnosticcode"], "\r", "")
			cs := status.Find(piece["diagnosticcode"], "")
			cr := reply.Find(piece["diagnosticcode"], cs)
			piece["deliverystatus"] = status.Prefer(piece["deliverystatus"], cs, cr)

			if len(cr) == 3 {
				// There is an SMTP reply code in the error message 
				if len(piece["replycode"]) == 0 { piece["replycode"] = cr }
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
						cx := fmt.Sprintf("%s%s%s", cr, q, cs)
						piece["diagnosticcode"] = strings.ReplaceAll(piece["diagnosticcode"], cx, "")

						// Remove "553-" and "553 " (SMTP reply code only) from the error message
						cx  = fmt.Sprintf("%s%s", cr, q)
						piece["diagnosticcode"] = strings.ReplaceAll(piece["diagnosticcode"], cx, "")
					}

					if strings.Contains(piece["diagnosticcode"], cr) {
						// Add "550 5.1.1" into the head of the error message when the error
						// message does not begin with "550"
						piece["diagnosticcode"] = fmt.Sprintf("%s %s %s", cr, cs, piece["diagnosticcode"])
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
			piece["diagnostictype"] = e.Reason
			piece["reason"]         = e.Reason

			if len(e.Spec) > 0 { break DIAGNOSTICTYPE }
			if piece["reason"] == "mailererror"                               { piece["diagnostictype"] = "X-UNIX" }
			if piece["reason"] != "feedback" && piece["reason"] != "vacation" { piece["diagnostictype"] = "SMTP"   }
			break DIAGNOSTICTYPE
		}

		// Set other values returned from sisimai/message.Rise()
		addrs["recipient"] = [3]string{e.Recipient, "", ""}
		piece["subject"]   = strings.ReplaceAll(rfc822data["subject"][0], "\r", "")
		if command.Test(e.Command) { piece["smtpcommand"] = e.Command }

		CONSTRUCTOR: for {
			// - Create email address object as address.EmailAddress struct
			// - Create decoded bounce mail object as sis.Fact struct
			as := sisiaddr.Rise(addrs["addresser"]); if as.Void() == true { continue RISEOF }
			ar := sisiaddr.Rise(addrs["recipient"]); if ar.Void() == true { continue RISEOF }

			thing.Action         = e.Action
			thing.Addresser      = as
			thing.Alias          = e.Alias; if len(thing.Alias) == 0 { thing.Alias = ar.Alias }
			/* TODO: Implemenet o.Catch = */
			thing.DeliveryStatus = piece["deliverystatus"]
			thing.Destination    = ar.Host
			thing.DiagnosticCode = piece["diagnosticcode"]
			thing.DiagnosticType = piece["diagnostictype"]
			thing.FeedbackType   = e.FeedbackType
			thing.HardBounce     = false
			thing.Lhost          = e.Lhost
			thing.ListID         = piece["listid"]
			thing.MessageID      = piece["messageid"]
			thing.Origin         = origin
			thing.Reason         = piece["reason"]
			thing.Rhost          = e.Rhost
			thing.Recipient      = ar
			thing.ReplyCode      = piece["replycode"]; if len(thing.ReplyCode) == 0 { reply.Find(piece["diagnosticcode"], "") }
			thing.SMTPAgent      = e.Agent
			thing.SMTPCommand    = piece["smtpcommand"]
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
			if len(thing.Alias) == 0                  { break ALIAS }
			if thing.Recipient.Address != thing.Alias { break ALIAS }
			if len(rfc822data["received"]) == 0       { break ALIAS }

			recv := rfc822data["received"]
			hops := len(recv)
			for i := hops - 1; hops >= 0; hops-- {
				// Search for the string " for " from the Received: header
				if strings.Contains(recv[i], " for ") == false { continue }
				or := rfc5322.Received(recv[i])

				if len(or) == 0                            { continue }
				if len(or[5]) == 0                         { continue }
				if sisiaddr.IsEmailAddress(or[5]) == false { continue }
				if or[5] == thing.Recipient.Address        { continue }

				thing.Alias = or[5]; break
			}
			break ALIAS
		}
		if thing.Alias == thing.Recipient.Address { thing.Alias = "" }

		REASON: for {
			// Decide the reason of email bounce
			if len(thing.Reason) == 0 || RetryIndex[thing.Reason] == true {
				// The value of "reason" is empty or is needed to check with other values again
				re := thing.Reason; if re == "" { re = "undefined" }
				thing.Reason = rhost.Find(&thing)
				if thing.Reason == "" { thing.Reason = reason.Find(&thing) }
				if thing.Reason == "" { thing.Reason = re }
			}
			break REASON
		}

		HARDBOUNCE: for {
			// Set the value of "hardbounce", default value of "bouncebounce" is 0
			if thing.Reason == "delivered" || thing.Reason == "feedback" || thing.Reason == "vacation" {
				// Delete the value of ReplyCode when the Reason is "feedback" or "vacation"
				if thing.Reason != "delivered" { thing.ReplyCode = "" }

			} else {
				// The Reason is not "delivered", or "feedback", or "vacation"
				smtperrors := fmt.Sprintf("%s %s", piece["deliverystatus"], piece["diagnosticcode"])
				if len(smtperrors) < 4 { smtperrors = "" }
				thing.HardBounce = failure.IsHardBounce(thing.Reason, smtperrors)
			}
			break HARDBOUNCE
		}

		DELIVERYSTATUS: for {
			// Set a pseudo status code
			if len(thing.DeliveryStatus) > 0 { break DELIVERYSTATUS }

			smtperrors := fmt.Sprintf("%s %s", thing.ReplyCode, piece["diagnosticcode"])
			if len(smtperrors) < 4 { smtperrors = "" }
			permanent0 := failure.IsPermanent(smtperrors)
			temporary0 := failure.IsTemporary(smtperrors)
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
				// There is an action value which is not described at RFC1894
				ox := rfc1894.Field("Action: " + thing.Action)
				if len(ox) > 0 {
					// Rewrite the value of "Action:" field to the valid value
					//
					// The syntax for the action-field is:
					//   action-field = "Action" ":" action-value
					//   action-value = "failed" / "delayed" / "delivered" / "relayed" / "expanded"
					thing.Action = ox[2]
				}
			}
			if thing.Reason == "expired"                            { thing.Action = "delayed" }
			if thing.Action == "" && (cx[0] == "4" || cx[0] == "5") { thing.Action = "failed"  }

			break REPLYCODE
		}
		listoffact = append(listoffact, thing)
	}

	for j, e := range listoffact {
		fmt.Printf("List-Of-Fact[%d] = %##v\n", j, e)
		fmt.Printf("----------------------------------\n")
		fmt.Printf("--[%d]DiagnosticCode = [%s]\n", j, e.DiagnosticCode)
		fmt.Printf("--[%d]DeliveryStatus = [%s]\n", j, e.DeliveryStatus)
		fmt.Printf("--[%d]ReplyCode = [%s]\n", j, e.ReplyCode)
		fmt.Printf("--[%d]Reason = [%s]\n", j, e.Reason)
		fmt.Printf("--[%d]DecodedBy = [%s]\n", j, e.SMTPAgent)
		fmt.Printf("--[%d]Command = [%s]\n", j, e.SMTPCommand)
		fmt.Printf("--[%d]Recipient = [%s]\n", j, e.Recipient.Address)
		fmt.Printf("--[%d]Lhost = [%s]\n", j, e.Lhost)
		fmt.Printf("--[%d]Rhost = [%s]\n", j, e.Rhost)
	}
	return listoffact
}

