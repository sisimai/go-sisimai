// Copyright (C) 2020-2022,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      

// Package "message" provides functions to read email message as a string, to tidy up each line.
package message

import "io"
import "strings"
import "net/mail"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc5322"

var pseudofrom = "MAILER-DAEMON Fri Feb  2 18:30:22 2018"
var boundaries = []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"};

// Rise decode and structure various formats of bounce emails.
//   Arguments:
//     - mesg (*string):          Entire email message.
//     - hook (sis.CfParameter0): The first callback function.
//   Returns:
//     - (*sis.BeforeFact): Decoded and structured bounce email data.
func Rise(mesg *string, hook sis.CfParameter0) *sis.BeforeFact {
	if mesg == nil || len(*mesg) < 1 { return new(sis.BeforeFact) }

	retryagain := 0
	beforefact := new(sis.BeforeFact)

	RISE: for retryagain < 2 {
		// 1. Split email data to headers and a body part.
		moji.ToLF(mesg)
		if email, nyaan := mail.ReadMessage(strings.NewReader(*mesg)); nyaan != nil {
			// Failed to read the message as an email
			ce := *sis.MakeNotDecoded(nyaan.Error(), true)
			beforefact.Errors = append(beforefact.Errors, ce)
			return beforefact

		} else {
			// Build "Message" struct
			if strings.HasPrefix(*mesg, "From ") {
				// The message has Unix From line (MAILER-DAEMON Tue Feb 11 00:00:00 2014)
				beforefact.Sender = moji.Select(moji.LHS + *mesg, "", "\n", 0)

			} else {
				// Set pseudo UNIX From line
				beforefact.Sender = pseudofrom
			}

			// Build "Head", "Body" members of BeforeFact
			beforefact.Headers = rfc5322.Headers(&email.Header, false)
			bodystring, nyaan := io.ReadAll(email.Body); if nyaan != nil { break RISE }
			beforefact.Payload = string(bodystring)
		}

		// 2. Rewrite the Subject header and the entire message body of the forwarded message
		if rawsubject := strings.TrimSpace(beforefact.Headers["subject"][0]); rawsubject != "" {
			// There used to be code in this block to decode MIME-encoded Subject headers, but it
			// was completely removed in v5.2.1.
			// See https://github.com/sisimai/go-sisimai/issues/42
			if cv := strings.ToLower(rawsubject); moji.HasPrefixAny(cv, []string{"fwd:", "fw:"}) {
				// - Remove "Fwd:" string from the "Subject:" header
				// - Delete quoted strings, quote symbols(>)
				rawsubject = strings.TrimSpace(moji.Select(cv + moji.RHS, ":", "", 0))
				beforefact.Payload = strings.ReplaceAll(beforefact.Payload, "\n> ", "\n")
				beforefact.Payload = strings.ReplaceAll(beforefact.Payload, "\n>\n", "\n\n")
			}
			beforefact.Headers["subject"][0] = rawsubject
		}

		// 3. Rewrite message body for detecting the bounce reason
		if siftstatus := sift(beforefact, hook); siftstatus == true { break RISE }
		for _, e := range boundaries {
			// Check the message body contains "message/rfc822" or "message/delivery-status" for
			// decoding the bounce message in the forwarded email
			if strings.Contains(beforefact.Payload, e) { break RISE }
		}

		// 4. Try to sift again
		//    There is a bounce message inside of mutipart/*, try to sift the first message/rfc822
		//    part as a entire message body again. rfc3464/1086-a847b090.eml is the email but the
		//    results decoded by sisimai are unstable.
		retryagain++
		cv := rfc5322.Part(&beforefact.Payload, boundaries, true)[1]; if len(cv) < 128 { break RISE }
		mesg = &cv
	}
	if beforefact.HasDone() == false { return new(sis.BeforeFact) }
	return beforefact
}

