// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _       _   ____                     _          _ 
// | \ | | ___ | |_|  _ \  ___  ___ ___   __| | ___  __| |
// |  \| |/ _ \| __| | | |/ _ \/ __/ _ \ / _` |/ _ \/ _` |
// | |\  | (_) | |_| |_| |  __/ (_| (_) | (_| |  __/ (_| |
// |_| \_|\___/ \__|____/ \___|\___\___/ \__,_|\___|\__,_|

package sis
import "fmt"
import "time"
import "runtime"

type NotDecoded struct {
	Timestamp time.Time // When the error occurred
	EmailFile string    // An email file name sisimai tried to decoded
	BecauseOf string    // An error message of the failure
	WhoCalled string    // Who called the constructor?
	DecodedBy string    // Copy of sis.Fact.DecodedBy
	CalledOff bool      // Unrecoverable error, the decoding process have called off
}

// MakeNotDecoded is a constructor of sis.NotDecoded struct.
//   Arguments:
//     - argv0 (string): Error message
//     - argv1 (bool):   Unrecoverable error or not
//   Returns:
//     - (*NotDecoded):  Initialized error struct
func MakeNotDecoded(argv0 string, argv1 bool) *NotDecoded {
	p, _, l, _ := runtime.Caller(1); return &NotDecoded{
		BecauseOf: argv0,
		CalledOff: argv1,
		Timestamp: time.Now(),
		WhoCalled: fmt.Sprintf("%s():%d", runtime.FuncForPC(p).Name(), l),
	}
}

// *NotDecoded.Error returns the error message as a string.
func(this *NotDecoded) Error() string {
	if this.BecauseOf == "" { return "" }

	timestring:= this.Timestamp.Format("2006/01/02 15:04:05")
	return timestring + " " + this.EmailFile + " " + this.BecauseOf
}

// *NotDecoded.Label returns a label string for printing error message.
func(this *NotDecoded) Label() string {
	if this.CalledOff == true { return " *****error: " }
	return " ***warning: "
}

// *NotDecoded.Email receives a path to email and set it into EmailFile.
//   Arguments:
//     - argv1 (string): Path to an email being set into the EmailFile
//   Returns:
//     - (string):       Current value of the EmailFile
func(this *NotDecoded) Email(argv1 string) string {
	if argv1          == "" { return this.EmailFile  }
	if this.EmailFile == "" { this.EmailFile = argv1 }
	return this.EmailFile
}

