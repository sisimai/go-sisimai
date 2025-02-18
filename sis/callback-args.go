// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//   ____      _ _ _                _        _                  
//  / ___|__ _| | | |__   __ _  ___| | __   / \   _ __ __ _ ___ 
// | |   / _` | | | '_ \ / _` |/ __| |/ /  / _ \ | '__/ _` / __|
// | |__| (_| | | | |_) | (_| | (__|   <  / ___ \| | | (_| \__ \
//  \____\__,_|_|_|_.__/ \__,_|\___|_|\_\/_/   \_\_|  \__, |___/
//                                                    |___/     

package sis

// CallbackArg0 is an argument of the first callback function that are called at message.sift().
// It is aliased to sisimai.CallbackArg0 at the libsisimai.go
type CallbackArg0 struct {
	Headers map[string][]string // Email headers of the bounce mail
	Payload *string             // Entire message body of the bounce mail
}

// CallbackArg1 is an argument of the callback functions that are called at sisimai.Rise(). It is
// aliased to sisimai.CallbackArg1 at the libsisimai.go
type CallbackArg1 struct {
	Path  string // Path to the original email file or "<STDIN>" or "<MEMORY>"
	Kind  string // Kind of the original email file or "stdin" or "memory"
	Mail *string // Entire message body of the bounce mail including all the headers
	Fact *[]Fact // Decoded results
}

