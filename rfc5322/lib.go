// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____ _________  ____  
// |  _ \|  ___/ ___| ___|___ /___ \|___ \ 
// | |_) | |_ | |   |___ \ |_ \ __) | __) |
// |  _ <|  _|| |___ ___) |__) / __/ / __/ 
// |_| \_\_|   \____|____/____/_____|_____|

// Package "rfc5322" provides functions for email addresses, Date: header, Received: headers, and
// other headers and messages related to RFC5322 https://datatracker.ietf.org/doc/html/rfc5322
package rfc5322

var FieldIndex = []string{
	// The following fields are not referred in Sisimai
	// Resent-From Resent-Sender Resent-Cc Cc Bcc Resent-Bcc In-Reply-To References Comments Keywords
	"Resent-Date", "From", "Sender", "Reply-To", "To", "Message-ID", "Subject", "Return-Path",
	"Received", "Date", "X-Mailer", "Content-Type", "Content-Transfer-Encoding", "Content-Description",
	"Content-Disposition",
}
var HeaderTable = map[string][]string{
	"messageid": []string{"message-id"},
	"subject":   []string{"subject"},
	"listid":    []string{"list-id"},
	"date":      []string{"date", "osted-date", "posted", "resent-date"},
	"addresser": []string{
		"from", "return-path", "reply-to", "errors-to", "reverse-path", "x-postfix-sender",
		"envelope-from", "x-envelope-from",
	},
	"recipient": []string{
		"to", "delivered-to", "forward-path", "envelope-to", "x-envelope-to", "resent-to", "x-apparently-to",
	},
}

