// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  ____        __                _____          _   
// | __ )  ___ / _| ___  _ __ ___|  ___|_ _  ___| |_ 
// |  _ \ / _ \ |_ / _ \| '__/ _ \ |_ / _` |/ __| __|
// | |_) |  __/  _| (_) | | |  __/  _| (_| | (__| |_ 
// |____/ \___|_|  \___/|_|  \___|_|  \__,_|\___|\__|
// Rise() in sisimai/message returns BeforeFact{}
type BeforeFact struct {
	From    string              // Unix FROM line ("From ")
	Head    map[string][]string // Email headers
	Body    string              // Email body
	Digest  []DeliveryMatter    // Decoded results returned from sisimai/lhost/*
	RFC822  string              // The original message
	Catch   func()              // Callback
}

func(this *BeforeFact) Void() bool {
	// @param    NONE
	// @return   bool Returns true if BeforeFact.Digest is empty
	if len(this.Digest) == 0 { return true }
	return false
}

