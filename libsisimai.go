// Copyright (C) 2020-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _ _         _     _                 _                  
// | (_) |__  ___(_)___(_)_ __ ___   __ _(_)  ___  _ __ __ _ 
// | | | '_ \/ __| / __| | '_ ` _ \ / _` | | / _ \| '__/ _` |
// | | | |_) \__ \ \__ \ | | | | | | (_| | || (_) | | | (_| |
// |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_(_)___/|_|  \__, |
// https://libsisimai.org/                              |___/

// sisimai is a library that decodes complex and diverse bounce emails and outputs the results of the
// delivery failure, such as the reason for the bounce and the recipient email address, in structured
// data. It is also possible to output in JSON format. 
// More information are available at https://libsisimai.org
package sisimai

import "io"
import "fmt"
import "errors"
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/moji"
import sisimbox "libsisimai.org/sisimai/mail"
import sisifact "libsisimai.org/sisimai/fact"

const libname string = "sisimai"
const version string = "5.2.1"
const patchlv uint8  = 0
type  CallbackArg0 = sis.CallbackArg0
type  CallbackArg1 = sis.CallbackArg1
type  CfParameter0 = sis.CfParameter0
type  CfParameter1 = sis.CfParameter1

// Version returns the version number of sisimai such as "v5.2.0" or "v5.2.1p22".
func Version() string {
	v := "v" + version; if patchlv > 0 { v += "p" + string(patchlv) }
	return v
}

// Args returns a pointer to sis.DecodingArgs as the 2nd argument of the Rise function.
func Args() *sis.DecodingArgs { return new(sis.DecodingArgs) }

// Rise is a function for decoding bounce mails in a mailbox or a Maildir/.
//   Arguments:
//     - path (string):            Path to an UNIX mbox, Maildir/, or "STDIN" for standard input.
//     - args (*sis.DecodingArgs): Options and callback functions for decoding bounce messages
//   Returns:
//     - (*[]sis.Fact):            List of successfully decoded bounce messages
//     - (*[]sis.NotDecoded):      List of occurred errors
func Rise(path string, args *sis.DecodingArgs) (*[]sis.Fact, *[]sis.NotDecoded) {
	sisidigest := []sis.Fact{}       // Decoded bounce message structures
	notdecoded := []sis.NotDecoded{} // List of occurred errors and warnings

	emailthing, nyaan := sisimbox.Rise(path); if nyaan != nil {
		// The file does not exist, or is not a regular file.
		ef := "<STDIN>"; if emailthing != nil { ef = emailthing.Path }
		ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), true); ce.Email(ef)
		notdecoded = append(notdecoded, ce)
		return &sisidigest, &notdecoded
	}

	for {
		// Read the email specified with the first argument until io.EOF
		if mesg, nyaan := emailthing.Read(); nyaan != nil {
			// Failed to read the email
			if errors.Is(nyaan, io.EOF) {
				// sisimai has reached to the end of email/directory
				break

			} else {
				// Something wrong, sisimai failed to read the email as a text
				ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), true); ce.Email(emailthing.Path)
				notdecoded = append(notdecoded, ce)
				continue
			}
		} else {
			// Read and decode each email file as a string
			if emailthing.Size == 0 {
				// The email file was empty
				ce := *sis.MakeNotDecoded("the email file is empty", true); ce.Email(emailthing.Path)
				notdecoded = append(notdecoded, ce)
				continue
			}
			moji.ToLF(mesg); facts, nyaan := sisifact.Rise(mesg, emailthing.Path, args)
			if facts != nil && len(*facts) > 0 { sisidigest = append(sisidigest, *facts...) }
			if nyaan != nil && len(*nyaan) > 0 { notdecoded = append(notdecoded, *nyaan...) }

			if args.Callback1 != nil {
				// Run the callback function stored in sis.DecodingArgs.Callback1 specified with the
				// 2nd argument of Sisimai.Rise() after reading each email file every time
				carg := &sis.CallbackArg1{Path: emailthing.Path, Kind: emailthing.Kind, Mail: mesg, Fact: facts}
				if _, nyaan := args.Callback1(carg); nyaan != nil {
					ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), true); ce.Email(emailthing.Path)
					notdecoded = append(notdecoded, ce)
				}
			}
		}
	}

	// TODO: Add warning information of the decoding results into notdecoded as sis.NotDecoded{}
	// when the reason is "onhold" or "undefined"
	return &sisidigest, &notdecoded
}

// Dump returns decoded data as a JSON string.
//   Arguments:
//     - path (string):            Path to an mbox, Maildir/, or "STDIN" for standard input.
//     - args (*sis.DecodingArgs): Options and callback functions for decoding bounce messages
//   Returns:
//     - (*string):                Decoded data as a JSON string array
//     - (*[]sis.NotDecoded):      List of occurred errors
func Dump(path string, args *sis.DecodingArgs) (*string, *[]sis.NotDecoded) {
	sisidigest, notdecoded := Rise(path, args); if len(*sisidigest) == 0 { return nil, notdecoded }
	serialized := []string{}

	for _, e := range *sisidigest {
		cj, nyaan := e.Dump(); if nyaan != nil {
			*notdecoded = append(*notdecoded, *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), false))
		}
		if cj != "" { serialized = append(serialized, cj) }
	}
	if len(serialized) == 0 { return nil, notdecoded }

	jsonstring := "[" + strings.Join(serialized, ",") + "]"
	return &jsonstring, notdecoded
}

