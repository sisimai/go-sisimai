// Copyright (C) 2020-2021,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____                 _ _    _       _     _                   
// | ____|_ __ ___   __ _(_) |  / \   __| | __| |_ __ ___  ___ ___ 
// |  _| | '_ ` _ \ / _` | | | / _ \ / _` |/ _` | '__/ _ \/ __/ __|
// | |___| | | | | | (_| | | |/ ___ \ (_| | (_| | | |  __/\__ \__ \
// |_____|_| |_| |_|\__,_|_|_/_/   \_\__,_|\__,_|_|  \___||___/___/

package siba
type EmailAddress struct {
	Address string // Email address
	User    string // Local part of the email address
	Host    string // Domain part of the email address
	Verp    string // Expanded VERP address
	Alias   string // Expanded Alias of the email address
	Name    string // Display name
	Comment string // (Comment)
}

