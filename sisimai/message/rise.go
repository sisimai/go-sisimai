// Copyright (C) 2020-2022,2024 azumakuniyuki and sisimai development team, All rights reserved.
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
import "sisimai/sis"
import "sisimai/lhost"
import "sisimai/rfc1894"
import "sisimai/rfc2045"
import "sisimai/rfc5322"
import "sisimai/rfc5965"
import sisimoji "sisimai/string"

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
func Rise(mesg *string, hook interface{}) sis.BeforeFact {
	// @param   *string     mesg  Entire email message
	// @param   interface{} hook  callback method
	// @return  Message           Structured email data
	if mesg == nil || len(*mesg) < 1 { return sis.BeforeFact{} }

	mesg        = sisimoji.ToLF(mesg)
	parseagain := 0
	beforefact := new(sis.BeforeFact)

	RISE: for parseagain < 2 {
		// 1. Split email data to headers and a body part.
		email, nyaan := mail.ReadMessage(strings.NewReader(*mesg))
		if nyaan != nil {
			// Failed to read the message as an email
			fmt.Printf("ERROR = %s\n", nyaan)
			break

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
				beforefact.Headers["subject"][0] = rfc2045.DecodeH(rawsubject)

			} else {
				// THe header is not mime-encoded
				beforefact.Headers["subject"][0] = rawsubject
			}

			// TODO: Remove "Fwd:" string from the "Subject:" header
			// TODO: Delete quoted strings, quote symbols(>)
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

		// TODO Implement this block
		// 4. Try to sift again
		//    There is a bounce message inside of mutipart/*, try to sift the first message/rfc822
		//    part as a entire message body again.
		//parseagain++
		break RISE
	}

	// TODO Implement this block
	// 5. Rewrite headers of the original message in the body part
	return *beforefact
}

