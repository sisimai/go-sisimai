// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____          _   
// |  ___|_ _  ___| |_ 
// | |_ / _` |/ __| __|
// |  _| (_| | (__| |_ 
// |_|  \__,_|\___|\__|
//                     
import "time"
import sisiaddr "sisimai/address"

// sisimai/fact.Rise() returns []sis.Fact
type Fact struct {
	Action          string                  // The value of "Action:" field
	Addresser       sisiaddr.EmailAddress   // The sender address of the original message
	Alias           string                  // The alias of the recipient address
	Catch           interface{}             // Results generated by user-defined callback function
	DeliveryStatus  string                  // Delivery Status such as "5.2.2"
	Destination     string                  // The domain part of the "Recipinet"
	DiagnosticCode  string                  // The value of "Diagnostic-Code:" field
	DiagnosticType  string                  // The 1st part of "Diagnostic-Code:" field
	FeedbackType    string                  // Feedback Type
	HardBounce      bool                    // Hard bounce or not
	ListID          string                  // The value of "List-Id" field of the original message
	Lhost           string                  // local host name/Local MTA
	MessageID       string                  // The value of "Message-Id:" header of the original message
	Origin          string                  // The email path as a data source
	Reason          string                  // Bounce reason
	Rhost           string                  // Remote host name/Remote MTA
	Recipient       sisiaddr.EmailAddress   // The recipient address of the original message
	ReplyCode       string                  // SMTP Reply Code such as "421"
	SMTPAgent       string                  // Module(Engine) name
	SMTPCommand     string                  // The last SMTP command
	SenderDomain    string                  // The domain part of the "Addresser"
	Subject         string                  // UTF-8 Subject text
	TimeStamp       time.Time               // The value of "Date:" header in the original message
	TimezoneOffset  int32                   // Time zone offset(seconds)
	Token           string                  // Message token/MD5 Hex digest value
}
