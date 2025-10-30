// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _   _       _   ____                     _          _ 
// | \ | | ___ | |_|  _ \  ___  ___ ___   __| | ___  __| |
// |  \| |/ _ \| __| | | |/ _ \/ __/ _ \ / _` |/ _ \/ _` |
// | |\  | (_) | |_| |_| |  __/ (_| (_) | (_| |  __/ (_| |
// |_| \_|\___/ \__|____/ \___|\___\___/ \__,_|\___|\__,_|

package siba
import "fmt"
import "time"
import "runtime"

type NotDecoded struct {
	Timestamp time.Time // When the error occurred
	EmailFile string    // An email file name sisimai tried to decoded
	BecauseOf string    // An error message of the failure
	WhoCalled string    // Who called the constructor?
	DecodedBy string    // Copy of siba.Fact.DecodedBy
	CalledOff bool      // Unrecoverable error, the decoding process have called off
}

// MakeNotDecoded is a constructor of siba.NotDecoded struct.
//   Arguments:
//     - mesg (string): Error message.
//     - flag (bool):   Unrecoverable error or not.
//   Returns:
//     - (*NotDecoded):  Initialized error struct.
func MakeNotDecoded(mesg string, flag bool) *NotDecoded {
	p, _, l, _ := runtime.Caller(1); return &NotDecoded{
		BecauseOf: mesg,
		CalledOff: flag,
		Timestamp: time.Now(),
		WhoCalled: fmt.Sprintf("%s():%d", runtime.FuncForPC(p).Name(), l),
	}
}

// *NotDecoded.Error returns the error message as a string.
//   Returns:
//     - (string): Formatted error message with a timestamp.
func(no *NotDecoded) Error() string {
	if no.BecauseOf == "" { return "" }

	timestring:= no.Timestamp.Format("2006/01/02 15:04:05")
	return timestring + " " + no.EmailFile + " " + no.BecauseOf
}

// *NotDecoded.Label returns a label string for printing error message.
//   Returns:
//     - (string): Label string
func(no *NotDecoded) Label() string {
	if no.CalledOff == true { return " *****error: " }
	return " ***warning: "
}

// *NotDecoded.Email receives a path to email and set it into EmailFile.
//   Arguments:
//     - path (string): Path to an email being set into the EmailFile.
//   Returns:
//     - (string): Current value of the EmailFile.
func(no *NotDecoded) Email(path string) string {
	if path         == "" { return no.EmailFile }
	if no.EmailFile == "" { no.EmailFile = path }
	return no.EmailFile
}

