// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____ _________  ____  
// |  _ \|  ___/ ___| ___|___ /___ \|___ \ 
// | |_) | |_ | |   |___ \ |_ \ __) | __) |
// |  _ <|  _|| |___ ___) |__) / __/ / __/ 
// |_| \_\_|   \____|____/____/_____|_____|

package rfc5322
import "strings"
import "net/mail"
import "libsisimai.org/sisimai/v5/moji"

// Headers converts a mail.Header struct to a map[string][]string.
//   Arguments:
//     - heads (*mail.Header): Email headers.
//   Returns:
//     - (map[string][]string: Structured email header data.
func Headers(heads *mail.Header) map[string][]string {
	headermaps := map[string][]string{}
	isrequired := []string{"from", "received", "message-id", "content-type", "subject"}

	for e, v := range *heads {
		// Each key name is the lower-cased string, each value is an array ([]string{})
		// The field name of an email header does not contain " "
		f := strings.ToLower(e)
		if moji.ContainsAny(f,  []string{" ", "authentication-"})  { continue }
		if moji.HasPrefixAny(f, []string{"arc-", "dkim-", "-spf"}) { continue }
		headermaps[f] = v
	}

	if cw := len(headermaps["received"]); cw > 0 {
		receivedby := make([]string, 0, cw)
		for _, e := range headermaps["received"] {
			// 1. Exclude the Received header including "(qmail ** invoked from network)".
			// 2. Convert all consecutive spaces and line breaks into a single space character.
			if moji.ContainsAny(e, woReceived) { continue }
			moji.Squeeze(&e, ' ')
			receivedby = append(receivedby, e)
		}
		headermaps["received"] = receivedby
	}

	for _, e := range isrequired {
		// The following fields should be exist
		if len(headermaps[e]) == 0 { headermaps[e] = []string{""} }
	}
	return headermaps
}

