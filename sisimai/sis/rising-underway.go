// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis
//  ____  _     _             _   _           _                               
// |  _ \(_)___(_)_ __   __ _| | | |_ __   __| | ___ _ ____      ____ _ _   _ 
// | |_) | / __| | '_ \ / _` | | | | '_ \ / _` |/ _ \ '__\ \ /\ / / _` | | | |
// |  _ <| \__ \ | | | | (_| | |_| | | | | (_| |  __/ |   \ V  V / (_| | |_| |
// |_| \_\_|___/_|_| |_|\__, |\___/|_| |_|\__,_|\___|_|    \_/\_/ \__,_|\__, |
//                      |___/                                           |___/ 

// Each MTA function in sisimai/lhost returns RisingUnderway{}
type RisingUnderway struct {
	Digest []DeliveryMatter // List of DeliveryMatter structs
	RFC822 string           // The original message
}

func(this *RisingUnderway) Void() bool {
	// @param    NONE
	// @return   bool   Returns true if RisingUnderway.Digest is empty
	if len(this.Digest) == 0 { return true }
	return false
}
