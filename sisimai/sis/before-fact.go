// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

// Sisimai
//  ___       _                        _   ____  _                   _                       
// |_ _|_ __ | |_ ___ _ __ _ __   __ _| | / ___|| |_ _ __ _   _  ___| |_ _   _ _ __ ___  ___ 
//  | || '_ \| __/ _ \ '__| '_ \ / _` | | \___ \| __| '__| | | |/ __| __| | | | '__/ _ \/ __|
//  | || | | | ||  __/ |  | | | | (_| | |  ___) | |_| |  | |_| | (__| |_| |_| | | |  __/\__ \
// |___|_| |_|\__\___|_|  |_| |_|\__,_|_| |____/ \__|_|   \__,_|\___|\__|\__,_|_|  \___||___/

// Rise() in sisimai/message returns BeforeFact{}
type BeforeFact struct {
	From    string              // Unix FROM line ("From ")
	Head    map[string][]string // Email headers
	Body    string              // Email body
	Digest  []DeliveryMatter    // Decoded results returned from sisimai/lhost/*
	RFC822  string              // The original message
	Catch   func()              // Callback
}

