// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____        __                _____          _   
// | __ )  ___ / _| ___  _ __ ___|  ___|_ _  ___| |_ 
// |  _ \ / _ \ |_ / _ \| '__/ _ \ |_ / _` |/ __| __|
// | |_) |  __/  _| (_) | | |  __/  _| (_| | (__| |_ 
// |____/ \___|_|  \___/|_|  \___|_|  \__,_|\___|\__|

package siba

// message.Rise() returns BeforeFact{}.
type BeforeFact struct {
	Headers map[string][]string // Email headers of the bounce mail
	RFC822  map[string][]string // Email headers of the original message
	Digest  []DeliveryMatter    // Decoded results returned from lhost/via-*.go
	Errors  []NotDecoded        // All the errors and warnings
	Catch   interface{}         // Any data structure returned by the callback function [0]
	Sender  string              // Unix FROM line ("From ")
	Payload string              // Entire message body of the bounce mail
}

// *BeforeFact.IsEmpty returns true when Headers or body is empty.
func(be *BeforeFact) IsEmpty() bool {
	if len(be.Headers) == 0 || be.Payload == "" { return true }
	return false
}

// *BeforeFact.HasDone returns false when Digest or RFC822 is empty.
func(be *BeforeFact) HasDone() bool {
	if len(be.Digest) == 0 || len(be.RFC822) == 0 { return false }
	return true
}

