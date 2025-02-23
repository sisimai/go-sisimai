// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____        __                _____          _   
// | __ )  ___ / _| ___  _ __ ___|  ___|_ _  ___| |_ 
// |  _ \ / _ \ |_ / _ \| '__/ _ \ |_ / _` |/ __| __|
// | |_) |  __/  _| (_) | | |  __/  _| (_| | (__| |_ 
// |____/ \___|_|  \___/|_|  \___|_|  \__,_|\___|\__|

package sis

// sisimai/message.Rise() returns BeforeFact{}
type BeforeFact struct {
	Sender  string              // Unix FROM line ("From ")
	Headers map[string][]string // Email headers of the bounce mail
	Payload string              // Entire message body of the bounce mail
	RFC822  map[string][]string // Email headers of the original message
	Digest  []DeliveryMatter    // Decoded results returned from lhost/via-*.go
	Catch   interface{}         // Any data structure returned by the callback function [0]
	Errors  []NotDecoded        // All the errors and warnings
}

// Empty() returns true when Headers or body is empty
func(this *BeforeFact) Empty() bool {
	if len(this.Headers) == 0 || this.Payload == "" { return true }
	return false
}

// Void() returns true when Digest or RFC822 is empty
func(this *BeforeFact) Void() bool {
	// @param    NONE
	// @return   bool   Returns true if BeforeFact.Digest or RFC822 is empty
	if len(this.Digest) == 0 || len(this.RFC822) == 0 { return true }
	return false
}

