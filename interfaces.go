// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ _ _         _     _                 _                     ___  __ 
// | (_) |__  ___(_)___(_)_ __ ___   __ _(_)  ___  _ __ __ _   / (_)/ _|
// | | | '_ \/ __| / __| | '_ ` _ \ / _` | | / _ \| '__/ _` | / /| | |_ 
// | | | |_) \__ \ \__ \ | | | | | | (_| | || (_) | | | (_| |/ / | |  _|
// |_|_|_.__/|___/_|___/_|_| |_| |_|\__,_|_(_)___/|_|  \__, /_/  |_|_|  
// https://libsisimai.org/                             |___/            

package sisimai
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/reason"

// Reason returns the list of bounce reasons sisimai can detect.
func Reason() map[string]string { return reason.Availables }

// Factor retuns empty (initialized with zero values) sis.Fact instance.
func Factor() *sis.Fact { return new(sis.Fact) }

