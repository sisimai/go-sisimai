// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  ____        __                _____          _   
// | __ )  ___ / _| ___  _ __ ___|  ___|_ _  ___| |_ 
// |  _ \ / _ \ |_ / _ \| '__/ _ \ |_ / _` |/ __| __|
// | |_) |  __/  _| (_) | | |  __/  _| (_| | (__| |_ 
// |____/ \___|_|  \___/|_|  \___|_|  \__,_|\___|\__|
// sisimai/message.Rise() returns BeforeFact{}
type BeforeFact struct {
	Sender  string              // Unix FROM line ("From ")
	Headers map[string][]string // Email headers
	Payload string              // Email body
	RFC822  map[string][]string // Email headers of the original message
	Digest  []DeliveryMatter    // Decoded results returned from sisimai/lhost/*
	Catch   interface{}         // Any data structure returned by the callback function
	Errors  []NotDecoded        // All the errors and warnings
}

func(this *BeforeFact) Void() bool {
	// @param    NONE
	// @return   bool   Returns true if BeforeFact.Digest or RFC822 is empty
	if len(this.Digest) == 0 { return true }
	if len(this.RFC822) == 0 { return true }
	return false
}

