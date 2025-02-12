// Copyright (C) 2020-2022,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package message

//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      
import "io"
import "fmt"
import "strings"
import "net/mail"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/lhost"
import "libsisimai.org/sisimai/rfc1894"
import "libsisimai.org/sisimai/rfc2045"
import "libsisimai.org/sisimai/rfc5322"
import "libsisimai.org/sisimai/rfc5965"
import sisimoji "libsisimai.org/sisimai/string"

var Fields1894 = rfc1894.FIELDINDEX()
var Fields5322 = rfc5322.FIELDINDEX()
var Fields5965 = rfc5965.FIELDINDEX()
var FieldTable = makefield(Fields1894, Fields5322, Fields5965)
var TryOnFirst = []string{}
var DefaultSet = lhost.AnotherOrder()
var Boundaries = []string{"Content-Type: message/rfc822", "Content-Type: text/rfc822-headers"};
var ReplacesAs = map[string][][]string{
    "Content-Type": [][]string{
		{"message/xdelivery-status",         "message/delivery-status"},
		{"message/disposition-notification", "message/delivery-status"},
	},
}

// makefield() generates a map including each field name defined in RFC1894, RFC5322, and RFC5965
func makefield(argv1 []string, argv2 []string, argv3 []string) map[string]string {
	fieldtable := map[string]string{}
	for _, e := range argv1 { fieldtable[strings.ToLower(e)] = e }
	for _, e := range argv2 { fieldtable[strings.ToLower(e)] = e }
	for _, e := range argv3 { fieldtable[strings.ToLower(e)] = e }
	return fieldtable
}

// Rise() works as a constructor of Sisimai::Message
func Rise(mesg *string, hook sis.CfParameter0) *sis.BeforeFact {
	// @param   *string          mesg  Entire email message
	// @param   sis.CfParameter0 hook  Callback Function
	// @return  Message                Structured email data
	if mesg == nil || len(*mesg) < 1 { return &sis.BeforeFact{} }

	mesg        = sisimoji.ToLF(mesg)
	retryagain := 0
	beforefact := new(sis.BeforeFact)

	RISE: for retryagain < 2 {
		// 1. Split email data to headers and a body part.
		email, nyaan := mail.ReadMessage(strings.NewReader(*mesg))
		if nyaan != nil {
			// Failed to read the message as an email
			ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), true)
			beforefact.Errors = append(beforefact.Errors, ce)
			return beforefact

		} else {
			// Build "Message" struct
			if strings.HasPrefix(*mesg, "From ") {
				// The message has Unix From line (MAILER-DAEMON Tue Feb 11 00:00:00 2014)
				beforefact.Sender = (*mesg)[0:strings.Index(*mesg, "\n")]

			} else {
				// Set pseudo UNIX From line
				beforefact.Sender = "MAILER-DAEMON Fri Feb  2 18:30:22 2018"
			}

			// Build "Head", "Body" members of BeforeFact
			beforefact.Headers = makemap(&email.Header, false)
			bodystring, nyaan := io.ReadAll(email.Body); if nyaan != nil { break RISE }
			beforefact.Payload = string(bodystring)
		}

		// 2. Decode and rewrite the "Subject:" header for deciding the order of MTA functions
		rawsubject := strings.TrimSpace(beforefact.Headers["subject"][0])
		if len(rawsubject) > 0 {
			// Decode MIME-Encoded "Subject:" header
			if rfc2045.IsEncoded(rawsubject) {
				// The header is mime-encoded
				cv, nyaan := rfc2045.DecodeH(rawsubject); beforefact.Headers["subject"][0] = cv
				if nyaan != nil {
					// Something wrong when the function decodes the MIME-Encoded Subejct header
					ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), false)
					beforefact.Errors = append(beforefact.Errors, ce)
				}
			} else {
				// THe header is not mime-encoded
				beforefact.Headers["subject"][0] = rawsubject
			}
			if cv := strings.ToLower(rawsubject); strings.HasPrefix(cv, "fwd:") || strings.HasPrefix(cv, "fw:") {
				// - Remove "Fwd:" string from the "Subject:" header
				// - Delete quoted strings, quote symbols(>)
				rawsubject = strings.TrimSpace(rawsubject[strings.Index(cv, ":") + 1:])
				beforefact.Payload = strings.ReplaceAll(beforefact.Payload, "\n> ", "\n")
				beforefact.Payload = strings.ReplaceAll(beforefact.Payload, "\n>\n", "\n\n")
			}
		}

		// 3. Rewrite message body for detecting the bounce reason
		TryOnFirst  = lhost.OrderBySubject(beforefact.Headers["subject"][0])
		TryOnFirst  = append(TryOnFirst, DefaultSet...)
		siftstatus := sift(beforefact, hook); if siftstatus == true { break RISE }

		for _, e := range Boundaries {
			// Check the message body contains "message/rfc822" or "message/delivery-status" for
			// decoding the bounce message in the forwarded email
			if strings.Contains(beforefact.Payload, e) { break RISE }
		}

		// 4. Try to sift again
		//    There is a bounce message inside of mutipart/*, try to sift the first message/rfc822
		//    part as a entire message body again. rfc3464/1086-a847b090.eml is the email but the
		//    results decoded by sisimai are unstable.
		retryagain++
		cv := rfc5322.Part(&beforefact.Payload, Boundaries, true)[1]; if len(cv) < 128 { break RISE }
		mesg = &cv
	}
	if beforefact.Void() == true { return &sis.BeforeFact{} }
	return beforefact
}

