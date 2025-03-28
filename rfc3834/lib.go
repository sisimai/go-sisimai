// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _____  ___ _____ _  _   
// |  _ \|  ___/ ___|___ / ( _ )___ /| || |  
// | |_) | |_ | |     |_ \ / _ \ |_ \| || |_ 
// |  _ <|  _|| |___ ___) | (_) |__) |__   _|
// |_| \_\_|   \____|____/ \___/____/   |_|  

// Package "rfc3834" provides functions like a MTA module in "lhost" package for decoding automatic
// responded messages formatted according to RFC3834; Recommendations for Automatic Responses to 
// Electronic Mail https://datatracker.ietf.org/doc/html/rfc3834
package rfc3834
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc2045"
import "libsisimai.org/sisimai/v5/rfc5322"
import "libsisimai.org/sisimai/v5/address"

// Inquire() decodes a bounce message that includes a vacation message.
//   Arguments:
//     - bf (*sis.BeforeFact):  Message entity in progress
//   Returns:
//     - (*sis.RisingUnderway): A structure as a staging data that is processed in message.sift() function
//   See:
//     - https://datatracker.ietf.org/doc/html/rfc3834
func Inquire(bf *sis.BeforeFact) *sis.RisingUnderway {
	if bf == nil || bf.IsEmpty() == true { return nil }

	boundaries := []string{"__SISIMAI_PSEUDO_BOUNDARY__"}
	lowerlabel := []string{"from", "to", "subject", "auto-submitted", "precedence", "x-apple-action"}
	lowervalue := map[string]string{}
	dontdecode := map[string][]string{
		"from":    []string{"root@", "postmaster@", "mailer-daemon@"},
		"to":      []string{"root@"},
		"subject": []string{
			"security information for", // sudo(1)
			"mail failure -",           // Exim
		},
    }
	autoreply0 := map[string][]string{
		// http://www.iana.org/assignments/auto-submitted-keywords/auto-submitted-keywords.xhtml
		"auto-submitted": []string{"auto-generated", "auto-replied", "auto-notified"},
		"precedence":     []string{"auto_reply"},
		"subject":        []string{"auto:", "auto response:", "automatic reply:", "out of office:", "out of the office:"},
		"x-apple-action": []string{"vacation"},
    };
	proceedsto := true

	for _, e := range lowerlabel {
		// Set lower-cased value of each header related to auto-response
		if len(bf.Headers[e]) > 0 { lowervalue[e] = strings.ToLower(bf.Headers[e][0]) }
	}

	DETECT_EXCLUSION_MESSAGE: for e := range dontdecode {
		// Exclude messages from root@
		if moji.ContainsAny(lowervalue[e], dontdecode[e]) == false { continue }
		proceedsto = false; break DETECT_EXCLUSION_MESSAGE
	}
	if proceedsto == false { return nil }

	proceedsto = false
	DETECT_AUTOREPLY_MESSAGE: for e := range autoreply0 {
		// Check Auto-Submitted field defined in RFC3834 and other headers
		if moji.HasPrefixAny(lowervalue[e], autoreply0[e]) == false { continue }
		proceedsto = true; break DETECT_AUTOREPLY_MESSAGE
	}
	if proceedsto == false { return nil }

	dscontents := []sis.DeliveryMatter{{}}
	recipients := uint8(0)            // The number of recipients
	v          := &(dscontents[len(dscontents) - 1])

	RECIPIENT_ADDRESS: for _, e := range []string{"from", "return-path"} {
		// Try to get the recipient adddress from some headers
		if len(bf.Headers[e]) == 0 { continue }
		cv := address.S3S4(bf.Headers[e][0]); if rfc5322.IsEmailAddress(cv) == false { continue }
		v.Recipient = cv
		recipients += 1
		break RECIPIENT_ADDRESS
	}
	if recipients == 0 { return nil }

	moji.Squeeze(&bf.Payload, '\n') // Squeeze continuous "\n" in the message body
	bf.Payload  = strings.Trim(bf.Payload, "\n")
	bodyslices := strings.Split(bf.Payload, "\n")
	rfc822part := ""

	if bf.Headers["content-type"][0] != "" {
		// Get the boundary string and set regular expression for matching with the boundary string.
		if cv := rfc2045.Boundary(bf.Headers["content-type"][0], 0); cv != "" { boundaries[0] = cv }
	}

	if len(bodyslices) < 5 {
		// There is vacation message only in the message body
		bf.Payload  = strings.ReplaceAll(bf.Payload, "\n", " ")
		v.Diagnosis = moji.Sweep(bf.Payload)

	} else {
		for _, e := range bodyslices {
			// Read vacation messages from the head of the email
			if e != "" && strings.HasPrefix(e, "--") == false { v.Diagnosis += e + " " }
		}
	}

	if p1 := strings.Index(bf.Headers["subject"][0], ": "); p1 > -1 {
		// Pick the original Subject: value from the bounce message
		if moji.ContainsAny(lowervalue["subject"], autoreply0["subject"]) {
			rfc822part += "Subject: " + moji.Sweep(bf.Headers["subject"][0][p1 + 2:]) + "\n"
		}
	}

	v.Reason    = "vacation"
	v.Date      = bf.Headers["date"][0]
	rfc822part += "To: <" + dscontents[0].Recipient + ">\n"
	return &sis.RisingUnderway{Digest: dscontents, RFC822: rfc822part}
}

