// Copyright (C) 2020-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _ _         _     _                 _                  
// | (_) |__  ___(_)___(_)_ __ ___   __ _(_)  ___  _ __ __ _ 
// | | | '_ \/ __| / __| | '_ ` _ \ / _` | | / _ \| '__/ _` |
// | | | |_) \__ \ \__ \ | | | | | | (_| | || (_) | | | (_| |
// |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_(_)___/|_|  \__, |
// https://libsisimai.org/                              |___/

// sisimai (pronounced /ɕi.ɕi.ma.i/) is a library that decodes complex and diverse bounce emails and
// outputs the results of the delivery failure, such as the reason for the bounce and the recipient
// email address, in structured data. It is also possible to output in JSON format.
// More information are available at https://libsisimai.org
package sisimai

import "io"
import "errors"
import "strings"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc5322"
import sisimbox "libsisimai.org/sisimai/v5/mail"
import sisifact "libsisimai.org/sisimai/v5/fact"

const libname string = "sisimai"
const version string = "5.5.0"
const patchlv uint8  = 0
type  CallbackArg0 = siba.CallbackArg0
type  CallbackArg1 = siba.CallbackArg1
type  CfParameter0 = siba.CfParameter0
type  CfParameter1 = siba.CfParameter1

// Version returns the version number of sisimai such as "v5.2.0" or "v5.2.1p22".
func Version() string {
	v := "v" + version; if patchlv > 0 { v += "p" + string(patchlv) }
	return v
}

// Args returns a pointer to siba.DecodingArgs as the 2nd argument of the Rise function.
func Args() *siba.DecodingArgs { return new(siba.DecodingArgs) }

// Rise is a function for decoding bounce mails in a mailbox or a Maildir/.
//   Arguments:
//     - path (string):             Path to an UNIX mbox, Maildir/, or "STDIN" for standard input.
//     - args (*siba.DecodingArgs): Options and callback functions for decoding bounce messages.
//   Returns:
//     - ([]siba.Fact):       List of successfully decoded bounce messages.
//     - ([]siba.NotDecoded): List of occurred errors.
func Rise(path string, args *siba.DecodingArgs) ([]siba.Fact, []siba.NotDecoded) {
	sisidigest := make([]siba.Fact, 0, 2)    // Decoded bounce message structures
	notdecoded := make([]siba.NotDecoded, 0) // List of occurred errors and warnings

	emailthing, nyaan := sisimbox.Rise(path); if nyaan != nil {
		// The file does not exist, or is not a regular file.
		ef := "<STDIN>"; if emailthing != nil { ef = emailthing.Path }
		ce := *siba.MakeNotDecoded(nyaan.Error(), true); ce.Email(ef)
		notdecoded = append(notdecoded, ce)
		return sisidigest, notdecoded
	}

	// The second argument `args` is a pointer to avoid potentially numerous internal struct copies
	// when fact.Rise function is called if the callback functions in siba.DecodingArgs.Callback0
	// or siba.DecodingArgs.Callback1 contain a large amount of code.
	if args == nil { args = new(siba.DecodingArgs) }

	for {
		// Read the email specified with the first argument until io.EOF
		if mesg, nyaan := emailthing.Read(); nyaan != nil {
			// Failed to read the email
			if errors.Is(nyaan, io.EOF) {
				// sisimai has reached to the end of email/directory
				break

			} else {
				// Something wrong, sisimai failed to read the email as a text
				ce := *siba.MakeNotDecoded(nyaan.Error(), true); ce.Email(emailthing.Path)
				notdecoded = append(notdecoded, ce)
				continue
			}
		} else {
			// Read and decode each email file as a string
			if emailthing.Size == 0 || rfc5322.LooksLikeEmail(mesg) == false {
				// Reason for this check:
				//   While mail.Rise() already validates the overall size of the input source, this
				//   specific condition addresses the case where an individual email message extracted
				//   from a multi-message source (like a UNIX mbox) might be empty.
				//   This acts as a necessary secondary validation to prevent issues during further
				//   processing of empty messages.
				ce := *siba.MakeNotDecoded("the file may not be a text file", true); ce.Email(emailthing.Path)
				notdecoded = append(notdecoded, ce)
				continue
			}
			moji.ToLF(mesg); facts, nyaan := sisifact.Rise(mesg, emailthing.Path, args)
			if facts != nil && len(facts) > 0 { sisidigest = append(sisidigest, facts...) }
			if nyaan != nil && len(nyaan) > 0 { notdecoded = append(notdecoded, nyaan...) }

			if args.Callback1 != nil {
				// Run the callback function stored in siba.DecodingArgs.Callback1 specified with
				// the 2nd argument of Sisimai.Rise() after reading each email file every time
				carg := &siba.CallbackArg1{Path: emailthing.Path, Kind: emailthing.Kind, Mail: mesg, Fact: &facts}
				if _, nyaan := args.Callback1(carg); nyaan != nil {
					ce := *siba.MakeNotDecoded(nyaan.Error(), true); ce.Email(emailthing.Path)
					notdecoded = append(notdecoded, ce)
				}
			}
		}
	}

	// TODO: Add warning information of the decoding results into notdecoded as siba.NotDecoded{}
	// when the reason is OnHold or Undefined.
	return sisidigest, notdecoded
}

// Dump returns decoded data as a JSON string.
//   Arguments:
//     - path (string):             Path to an mbox, Maildir/, or "STDIN" for standard input.
//     - args (*siba.DecodingArgs): Options and callback functions for decoding bounce messages.
//   Returns:
//     - (*string):           Decoded data as a JSON string array
//     - ([]siba.NotDecoded): List of occurred errors
func Dump(path string, args *siba.DecodingArgs) (*string, []siba.NotDecoded) {
	sisidigest, notdecoded := Rise(path, args); if len(sisidigest) == 0 { return nil, notdecoded }
	serialized := make([]string, 0)

	for _, e := range sisidigest {
		cj, nyaan := e.Dump(); if nyaan != nil {
			notdecoded = append(notdecoded, *siba.MakeNotDecoded(nyaan.Error(), false))
		}
		if cj != "" { serialized = append(serialized, cj) }
	}
	if len(serialized) == 0 { return nil, notdecoded }

	jsonstring := "[" + strings.Join(serialized, ",") + "]"
	return &jsonstring, notdecoded
}

