// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      

package message
import "fmt"
import "strings"
import "net/mail"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/arf"
import "libsisimai.org/sisimai/lhost"
import "libsisimai.org/sisimai/rfc2045"
import "libsisimai.org/sisimai/rfc3464"
import "libsisimai.org/sisimai/rfc3834"
import sisimoji "libsisimai.org/sisimai/string"

// sift() sifts a bounce mail with each MTA module
func sift(bf *sis.BeforeFact, hook sis.CfParameter0) bool {
	// @param  *sis.BeforeFact  bf     Processing message entity.
	// @param  sis.CfParameter0 hook   The callback function for the decoded bounce message
	// @return bool                    true:  Successfully got the results
	//                                 false: Failed to get the results
	if bf == nil || bf.Empty() == true { return false }

	bf.Payload = *(tidy(&bf.Payload)) // Tidy up each field name and value in the entire message body
	mesgformat := ""
	ctencoding := ""

	if len(bf.Headers["content-type"])              > 0 { mesgformat = strings.ToLower(bf.Headers["content-type"][0])              }
	if len(bf.Headers["content-transfer-encoding"]) > 0 { ctencoding = strings.ToLower(bf.Headers["content-transfer-encoding"][0]) }

	if strings.HasPrefix(mesgformat, "text/plain") || strings.HasPrefix(mesgformat, "text/html") {
		// Content-Type: text/plain; charset=UTF-8
		if ctencoding == "base64" {
			// Content-Transfer-Encoding: base64
			cv, nyaan := rfc2045.DecodeB(bf.Payload, ""); bf.Payload = cv
			if nyaan != nil {
				// Something wrong when the function decodes the BASE64 encoded string
				ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), false)
				bf.Errors = append(bf.Errors, ce)
			}
		} else if ctencoding == "quoted-printable" {
			// Content-Transfer-Encoding: quoted-printable
			cv, nyaan := rfc2045.DecodeQ(bf.Payload); bf.Payload = cv
			if nyaan != nil {
				// Something wrong when the function decodes the Quoted-Printable encoded string
				ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), false)
				bf.Errors = append(bf.Errors, ce)
			}
		}
		if strings.HasPrefix(mesgformat, "text/html") { bf.Payload = *(sisimoji.ToPlain(&bf.Payload)) }

	} else if strings.HasPrefix(mesgformat, "multipart/") {
		// In case of Content-Type: multipart/*
		cv, fe := rfc2045.MakeFlat(bf.Headers["content-type"][0], &bf.Payload)
		if cv != nil                { bf.Payload = *cv                      }
		if fe != nil && len(fe) > 0 { bf.Errors  = append(bf.Errors, fe...) }
	}
	bf.Payload  = *(sisimoji.ToLF(&bf.Payload))
	bf.Payload  = strings.ReplaceAll(bf.Payload, "\t", " ") // Replace all the TAB with " "

	if hook != nil {
		// Execute the first callback function
		cvv, nyaan := hook(&sis.CallbackArgs{Headers: bf.Headers, Payload: &bf.Payload}); if nyaan != nil {
			// Something wrong when the 1st callback function executed
			ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), false)
			bf.Errors = append(bf.Errors, ce)
		}
		bf.Catch = cvv
	}

	havecalled := map[string]bool{}
	localhostr := sis.RisingUnderway{}
	modulename := ""

	DECODER: for bf.Empty() == false {
		// 1. MTA Module Candidates to be tried on first, and other sisimai/lhost/*.go
		// 2. sisimai/rfc3464
		// 3. sisimai/arf
		// 4. sisimai/rfc3834
		for _, r := range TryOnFirst {
			// 1. MTA Module Candidates to be tried on first, and other sisimai/lhost/*.go
			if havecalled[r] || r == "ARF" || strings.HasPrefix(r, "RFC") { continue }
			localhostr    = lhost.InquireFor[r](bf)
			havecalled[r] = true
			modulename    = r
			if localhostr.Void() == false { break DECODER }
		}

		if havecalled["rfc3464"] == false {
			// 2. sisimai/rfc3464
			// When the all of sisimai/lhost/*.go modules did not return the decoded data
			localhostr = rfc3464.Inquire(bf)
			havecalled["rfc3464"] = true
			modulename = "RFC3464"
			if localhostr.Void() == false { break DECODER }
		}

		if havecalled["arf"] == false {
			// 3. call sisimai/arf
			// Try to decode the message as a Feedback Loop message
			localhostr = arf.Inquire(bf)
			modulename = "ARF"
			if localhostr.Void() == false { break DECODER }
		}

		if havecalled["rfc3834"] == false {
			// 4. call sisimai/rfc3834
			// Try to sift the message as auto reply message defined in RFC3834
			localhostr = rfc3834.Inquire(bf)
			modulename = "RFC3834"
			if localhostr.Void() == false { break DECODER }
		}
		break // as of now, we have no sample email for coding this block

	} // End of for(DECODER)
	if localhostr.Void() == true { return false }

	for j, _ := range localhostr.Digest {
		// Set the value of "Agent" such as "Postfix", "Sendmail", or "OpenSMTPD"
		if localhostr.Digest[j].Agent != "" { continue }
		localhostr.Digest[j].Agent = modulename
	}

	if strings.Contains(localhostr.RFC822, "\nFrom:") == false && len(bf.Headers["to"]) > 0 {
		// There is no "From:" header, pick the email address from the "To:" header of the
		// bounce message
		localhostr.RFC822 = fmt.Sprintf("From: %s\n%s", bf.Headers["to"][0], localhostr.RFC822)
	}
	di := &(localhostr.Digest[0])
	if strings.Contains(localhostr.RFC822, "\nTo:") == false && di.Recipient != "" {
		// The original message block is empty, insert some values picked from localhostr.Digest as
		// a pseudo header such as "To:", "Date:".
		localhostr.RFC822 = fmt.Sprintf("To: <%s>\n%s", di.Recipient, localhostr.RFC822)
	}

	// Convert headers of the original message to data structure/map[string][]string
	rfc822text := ""; for _, e := range strings.Split(localhostr.RFC822, "\n") {
		// Append each line of localhostr.RFC822 to rfc822text except malformed headers
		if e == "" && rfc822text != ""  { break } // The blank line between the header and the body
		if strings.Index(e, ":") < 1 {            // The line does not contain ":" or begins with ":"
			// The line is not a line continued from the previous line of a long header
			if strings.HasPrefix(e, " ") == false || strings.HasPrefix(e, "\t") == false { continue }
		}
		rfc822text += e + "\n"
	}
	if rfc822text != "" { localhostr.RFC822 = rfc822text + "\n" }

	rfc822part, nyaan := mail.ReadMessage(strings.NewReader(localhostr.RFC822))
	if nyaan != nil {
		// Failed to read the original message part
		ce := *sis.MakeNotDecoded(fmt.Sprintf("%s", nyaan), false)
		bf.Errors = append(bf.Errors, ce)
		return false
	}
	bf.RFC822 = makemap(&rfc822part.Header, false)
	bf.Digest = localhostr.Digest

	return true
}

