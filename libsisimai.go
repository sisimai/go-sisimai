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
import sisimbox "libsisimai.org/sisimai/mail"
import sisifact "libsisimai.org/sisimai/fact"
import sisimoji "libsisimai.org/sisimai/string"

const libname string = "sisimai"
const version string = "5.2.0"
const patchlv uint8  = 0
type  CallbackArgs = sis.CallbackArgs
type  CfParameter0 = sis.CfParameter0
type  CfParameter1 = sis.CfParameter1

// Version() returns the version number of sisimai
func Version() string {
	// @param   NONE
	// @return  string  version number like "v5.2.0p22"
	v := "v" + version; if patchlv > 0 { v += "p" + string(patchlv) }
	return v
}

// Args() returns the pointer to sis.DecodingArgs{} as the 2nd argument of Rise() function
func Args() *sis.DecodingArgs {
	// @param   NONE
	// @return  *sis.DecodingArgs
	return &sis.DecodingArgs{
		Delivered: false, // Include sis.Fact{}.Action = "delivered" records in the decoded data
		Vacation:  false, // Include sis.Fact{}.Reason = "vacation" records in the decoded data
		Callback0: nil,   // [0] The 1st callback function
		Callback1: nil,   // [1] The 2nd callback function
	}
}

// sisimai.Rise() is a function for decoding bounce mails in a mailbox or a Maildir/
func Rise(path string, args *sis.DecodingArgs) (*[]sis.Fact, *[]sis.NotDecoded) {
	// @param   string            path  Path to mbox or Maildir/ or "STDIN"
	// @param   *sis.DecodingArgs args  Arguments for decoding
	// @return  []sis.Fact
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
			mesg = sisimoji.ToLF(mesg)
			fact, nyaan := sisifact.Rise(mesg, emailthing.Path, args)
			if len(fact)  > 0 { sisidigest = append(sisidigest, fact...)  }
			if len(nyaan) > 0 { notdecoded = append(notdecoded, nyaan...) }

			if args.Callback1 != nil {
				// Run the callback function stored in sis.DecodingArgs.Callback1 specified with the
				// 2nd argument of Sisimai.Rise() after reading each email file every time
				carg := &sis.CallbackArgs{
					Headers: map[string][]string{
						"path": []string{emailthing.Path},
						"dir":  []string{emailthing.Dir},
						"file": []string{emailthing.File},
						"kind": []string{emailthing.Kind},
					},
					Payload: mesg,
				}
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

// sisimai.Dump() returns decoded data as a JSON string
func Dump(path string, args *sis.DecodingArgs) (*string, *[]sis.NotDecoded) {
	// @param   string            path  Path to mbox or Maildir/ or "STDIN"
	// @param   *sis.DecodingArgs args  Arguments for decoding
	// @return  *string
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

