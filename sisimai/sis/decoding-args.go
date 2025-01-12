// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  ____                     _ _                _                  
// |  _ \  ___  ___ ___   __| (_)_ __   __ _   / \   _ __ __ _ ___ 
// | | | |/ _ \/ __/ _ \ / _` | | '_ \ / _` | / _ \ | '__/ _` / __|
// | |_| |  __/ (_| (_) | (_| | | | | | (_| |/ ___ \| | | (_| \__ \
// |____/ \___|\___\___/ \__,_|_|_| |_|\__, /_/   \_\_|  \__, |___/
//                                     |___/             |___/     
// DecodingArgs{} is an argument of the sisimai.Rise() function
type DecodingArgs struct {
	Delivered bool // Include sis.Fact{}.Action = "delivered" records in the decoded data
	Vacation  bool // Include sis.Fact{}.Reason = "vacation" records in the decoded data
	Callback1 func(arg *CallbackArgs) (map[string]interface{}, error) // The 1st callback function
	Callback2 func(arg *CallbackArgs) (bool, error)                   // The 2nd callback function
}

