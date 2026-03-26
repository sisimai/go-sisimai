// Copyright (C) 2025-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____                     _ _                _                  
// |  _ \  ___  ___ ___   __| (_)_ __   __ _   / \   _ __ __ _ ___ 
// | | | |/ _ \/ __/ _ \ / _` | | '_ \ / _` | / _ \ | '__/ _` / __|
// | |_| |  __/ (_| (_) | (_| | | | | | (_| |/ ___ \| | | (_| \__ \
// |____/ \___|\___\___/ \__,_|_|_| |_|\__, /_/   \_\_|  \__, |___/
//                                     |___/             |___/     

package siba

// CfParameter* is an argument of the callback function specified at sisimai.Rise().
type CfParameter0 func(arg *CallbackArg0) (map[string]any, error)
type CfParameter1 func(arg *CallbackArg1) (bool, error)

// DecodingArgs is an argument of the sisimai.Rise() function.
type DecodingArgs struct {
	Callback0 CfParameter0 // [0] The 1st callback function
	Callback1 CfParameter1 // [1] The 2nd callback function
	Delivered bool         // Include siba.Fact{}.Action = "delivered" records in the decoded data
	Vacation  bool         // Include siba.Fact{}.Reason = "Vacation" records in the decoded data
}

