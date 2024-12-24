// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//   ____      _ _ _                _        _                  
//  / ___|__ _| | | |__   __ _  ___| | __   / \   _ __ __ _ ___ 
// | |   / _` | | | '_ \ / _` |/ __| |/ /  / _ \ | '__/ _` / __|
// | |__| (_| | | | |_) | (_| | (__|   <  / ___ \| | | (_| \__ \
//  \____\__,_|_|_|_.__/ \__,_|\___|_|\_\/_/   \_\_|  \__, |___/
//                                                    |___/     
// CallbackArgs{} is an argument of the callback functions that are called at sisimai.Rise() and 
// sisimai/message.sift()
type CallbackArgs struct {
	Headers map[string][]string
	Payload *string
}

