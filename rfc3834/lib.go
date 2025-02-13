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
import "fmt"
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/rfc2045"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"
import sisiaddr "libsisimai.org/sisimai/address"

// Inquire() decodes a bounce message that includes a vacation message
func Inquire(bf *sis.BeforeFact) sis.RisingUnderway {
	// @param    *sis.BeforeFact bf  Message body of a bounce email
	// @return   RisingUnderway      RisingUnderway structure
	// @see      https://tools.ietf.org/html/rfc3834
	if bf == nil || len(bf.Headers) == 0 || bf.Payload == "" { return sis.RisingUnderway{} }

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
		if len(lowervalue[e]) == 0 { continue }
		if sisimoji.ContainsAny(lowervalue[e], dontdecode[e]) == false { continue }
		proceedsto = false; break DETECT_EXCLUSION_MESSAGE
	}
	if proceedsto == false { return sis.RisingUnderway{} }

	proceedsto = false
	DETECT_AUTOREPLY_MESSAGE: for e := range autoreply0 {
		// Check Auto-Submitted field defined in RFC3834 and other headers
		if len(lowervalue[e]) == 0 { continue }
		if sisimoji.HasPrefixAny(lowervalue[e], autoreply0[e]) == false { continue }
		proceedsto = true; break DETECT_AUTOREPLY_MESSAGE
	}
	if proceedsto == false { return sis.RisingUnderway{} }

	dscontents := []sis.DeliveryMatter{{}}
	recipients := uint8(0)            // The number of recipients
	v          := &(dscontents[len(dscontents) - 1])

	RECIPIENT_ADDRESS: for _, e := range []string{"from", "return-path"} {
		// Try to get the recipient adddress from some headers
		if len(bf.Headers[e]) == 0 { continue }
		cv := sisiaddr.S3S4(bf.Headers[e][0]); if rfc5322.IsEmailAddress(cv) == false { continue }
		v.Recipient = cv
		recipients += 1
		break RECIPIENT_ADDRESS
	}
	if recipients == 0 { return sis.RisingUnderway{} }

	// Squeeze continuous "\n" in the message body
	bf.Payload  = strings.Trim(strings.ReplaceAll(bf.Payload, "\n\n", "\n"), "\n")
	bodyslices := strings.Split(bf.Payload, "\n")
	rfc822part := ""

	if bf.Headers["content-type"][0] != "" {
		// Get the boundary string and set regular expression for matching with the boundary string.
		cv := rfc2045.Boundary(bf.Headers["content-type"][0], 0)
		if cv != "" { boundaries[0] = cv }
	}

	if len(bodyslices) < 5 {
		// There is vacation message only in the message body
		bf.Payload  = strings.ReplaceAll(bf.Payload, "\n", " ")
		v.Diagnosis = sisimoji.Sweep(bf.Payload)

	} else {
		for _, e := range bodyslices {
			// Read vacation messages from the head of the email
			if len(e) == 0 || strings.HasPrefix(e, "--") { continue }
			v.Diagnosis += e + " "
		}
	}

	for {
		// Pick the original Subject: value from the bounce message
		p1 := strings.Index(bf.Headers["subject"][0], ": "); if p1 < 0 { break }
		if sisimoji.ContainsAny(lowervalue["subject"], autoreply0["subject"]) == false { break }
		cv := sisimoji.Sweep(bf.Headers["subject"][0][p1 + 2:])
		rfc822part += fmt.Sprintf("Subject: %s\n", cv)
		break
	}

	v.Reason    = "vacation"
	v.Date      = bf.Headers["date"][0]
	rfc822part += fmt.Sprintf("To: <%s>\n", dscontents[0].Recipient)
	return sis.RisingUnderway{ Digest: dscontents, RFC822: rfc822part }
}

