// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc3834

//  ____  _____ ____ _____  ___ _____ _  _   
// |  _ \|  ___/ ___|___ / ( _ )___ /| || |  
// | |_) | |_ | |     |_ \ / _ \ |_ \| || |_ 
// |  _ <|  _|| |___ ___) | (_) |__) |__   _|
// |_| \_\_|   \____|____/ \___/____/   |_|  
import "fmt"
import "strings"
import "sisimai/sis"
import "sisimai/rfc2045"
import sisimoji "sisimai/string"
import sisiaddr "sisimai/address"

// Inquire() decodes a bounce message which have fields defined in RFC3464
func Inquire(bf *sis.BeforeFact) sis.RisingUnderway {
	// @param    *sis.BeforeFact bf  Message body of a bounce email
	// @return   RisingUnderway      RisingUnderway structure
	// @see      https://tools.ietf.org/html/rfc3464
	if len(bf.Headers) == 0 { return sis.RisingUnderway{} }
	if len(bf.Payload) == 0 { return sis.RisingUnderway{} }

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
		cv := sisiaddr.S3S4(bf.Headers[e][0]); if sisiaddr.IsEmailAddress(cv) == false { continue }
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

	v.Reason    = "vacaion"
	v.Date      = bf.Headers["date"][0]
	rfc822part += fmt.Sprintf("To: <%s>\n", dscontents[0].Recipient)
	return sis.RisingUnderway{ Digest: dscontents, RFC822: rfc822part }
}

