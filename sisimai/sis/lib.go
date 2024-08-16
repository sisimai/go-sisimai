// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

// Sisimai
//  ___       _                        _   ____  _                   _                       
// |_ _|_ __ | |_ ___ _ __ _ __   __ _| | / ___|| |_ _ __ _   _  ___| |_ _   _ _ __ ___  ___ 
//  | || '_ \| __/ _ \ '__| '_ \ / _` | | \___ \| __| '__| | | |/ __| __| | | | '__/ _ \/ __|
//  | || | | | ||  __/ |  | | | | (_| | |  ___) | |_| |  | |_| | (__| |_| |_| | | |  __/\__ \
// |___|_| |_|\__\___|_|  |_| |_|\__,_|_| |____/ \__|_|   \__,_|\___|\__|\__,_|_|  \___||___/
type DeliveryStatus struct {
	Action       string     // The value of Action header
	Agent        string     // MTA name
	Alias        string     // The value of alias entry(RHS)
	Command      string     // SMTP command in the message body
	Date         string     // The value of Last-Attempt-Date header
	Diagnosis    string     // The value of Diagnostic-Code header
	FeedbackType string     // Feedback type
	Lhost        string     // The value of Received-From-MTA header
	Reason       string     // Temporary reason of bounce
	Recipient    string     // The value of Final-Recipient header
	ReplyCode    string     // SMTP Reply Code
	Rhost        string     // The value of Remote-MTA header
	HardBounce   bool       // Hard bounce or not
	Spec         string     // Protocl specification
	Status       string     // The value of Status header
}

// Each MTA function in sisimai/lhost returns RisingUnderway{}
type RisingUnderway struct {
	Digest []DeliveryStatus // List of DeliveryStatus structs
	RFC822 string           // The original message
}

func(this *RisingUnderway) Void() bool {
	// @param    NONE
	// @return   bool Returns true if RisingUnderway.Digest is empty
	if len(this.Digest) == 0 { return true }
	return false
}

// Rise() in sisimai/message returns BeforeFact{}
type BeforeFact struct {
	From    string              // Unix FROM line ("From ")
	Head    map[string][]string // Email headers
	Body    string              // Email body
	Digest  []DeliveryStatus    // Decoded results returned from sisimai/lhost/*
	RFC822  string              // The original message
	Catch   func()              // Callback
}

