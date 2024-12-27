// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _   _       _   ____                     _          _ 
// | \ | | ___ | |_|  _ \  ___  ___ ___   __| | ___  __| |
// |  \| |/ _ \| __| | | |/ _ \/ __/ _ \ / _` |/ _ \/ _` |
// | |\  | (_) | |_| |_| |  __/ (_| (_) | (_| |  __/ (_| |
// |_| \_|\___/ \__|____/ \___|\___\___/ \__,_|\___|\__,_|
import "fmt"
import "time"

// NotDecoded{} is a structure keeping a decoding error at sisimai.Rise()
type NotDecoded struct {
	EmailFile string    // An email file name sisimai tried to decoded
	BecauseOf string    // An error message of the failure
	Addresser string    // Copy of sis.Fact.Addresser.Address
	MessageID string    // Copy of sis.Fact.MessageID
	DecodedBy string    // Copy of sis.Fact.DecodedBy
	Timestamp time.Time // When the error occurred
}

// Error() returns the error message as a string
func(this *NotDecoded) Error() string {
	// @param    NONE
	// @return   string  an error message
	if this.EmailFile == "" || this.BecauseOf == "" { return "" }

	timestring:= this.Timestamp.Format("2006/01/02 15:04:05")
	return fmt.Sprintf("%s %s %s", timestring, this.EmailFile, this.BecauseOf)
}

