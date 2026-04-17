// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ _____  ___ _____ _  _   
// |  _ \|  ___/ ___|___ / ( _ )___ /| || |  
// | |_) | |_ | |     |_ \ / _ \ |_ \| || |_ 
// |  _ <|  _|| |___ ___) | (_) |__) |__   _|
// |_| \_\_|   \____|____/ \___/____/   |_|  

// Package "rfc3834" provides functions like a MTA module in "lhost" package for decoding automatic
// responded messages formatted according to RFC3834; Recommendations for Automatic Responses to 
// Electronic Mail. https://datatracker.ietf.org/doc/html/rfc3834
package rfc3834
import "bytes"
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/rfc5322"
import "libsisimai.org/sisimai/v5/address"

// Inquire decodes a bounce message that includes a vacation message.
//   Arguments:
//     - bf (*siba.BeforeFact): Message entity in progress.
//   Returns:
//     - (*siba.RisingUnderway): A structure as a staging data that is processed in message.sift() function.
//   See:
//     - https://datatracker.ietf.org/doc/html/rfc3834
func Inquire(bf *siba.BeforeFact) *siba.RisingUnderway {
	if bf == nil || bf.IsEmpty() == true { return nil }

	proceedsto := true
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
    }
	suspending := [][]string{
		[]string{"this email inbox", " is no longer in use."},
	}

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

	dscontents := make([]siba.DeliveryMatter, 1); v := &dscontents[0]
	recipients := uint8(0)
	v.Reason    = eb.ReAWAY

	RECIPIENT_ADDRESS: for _, e := range []string{"reply-to", "from", "return-path"} {
		// Try to get the recipient adddress from some headers
		if len(bf.Headers[e]) == 0 { continue }
		cv := address.S3S4(bf.Headers[e][0]); if rfc5322.IsEmailAddress(cv) == false { continue }
		v.Recipient = cv
		recipients += 1
		break RECIPIENT_ADDRESS
	}
	if recipients == 0 { return nil }

	bodystring := string(bytes.Trim(bytes.ReplaceAll(bf.Payload, []byte("\n\n"), []byte("\n")), "\n"))
	bodyslices := strings.Split(bodystring, "\n")
	rfc822part := strings.Builder{}; rfc822part.Grow(512)

	if len(bodyslices) < 5 {
		// There is vacation message only in the message body
		bf.Payload  = []byte(strings.Join(strings.Fields(bodystring), " "))
		v.Diagnosis = string(bf.Payload)

	} else {
		messagebuf := strings.Builder{}; messagebuf.Grow(len(bf.Payload))
		for _, e := range bodyslices {
			// Read vacation messages from the head of the email
			if len(e) < 1 || strings.HasPrefix(e, "--") { continue }
			messagebuf.WriteString(e); messagebuf.WriteByte(' ')
		}
		v.Diagnosis = messagebuf.String()
	}

	if p1 := strings.Index(bf.Headers["subject"][0], ": "); p1 > -1 {
		// Pick the original Subject: value from the bounce message
		if moji.ContainsAny(lowervalue["subject"], autoreply0["subject"]) {
			// 
			rfc822part.WriteString("Subject: ")
			rfc822part.WriteString(bf.Headers["subject"][0][p1 + 2:])
			rfc822part.WriteByte('\n')
		}
	}

	cv := strings.ToLower(v.Diagnosis); for _, e := range suspending {
		// Check that the auto-replied message indicates the "Suspend" reason or not.
		if moji.Aligned(cv, e) { v.Reason = eb.ReQUIT; break }
	}

	v.Date = bf.Headers["date"][0]
	rfc822part.WriteString("To: <")
	rfc822part.WriteString(dscontents[0].Recipient)
	rfc822part.WriteString(">\n")
	return &siba.RisingUnderway{Digest: dscontents, RFC822: rfc822part.String()}
}

