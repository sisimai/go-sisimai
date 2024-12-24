// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package message

//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      
import "strings"
import "net/mail"
import "sisimai/sis"
import "sisimai/arf"
import "sisimai/lhost"
import "sisimai/rfc2045"
import "sisimai/rfc3464"
import "sisimai/rfc3834"
import sisimoji "sisimai/string"

// sift() sifts a bounce mail with each MTA module
func sift(bf *sis.BeforeFact, hook interface{}) bool {
	// @param  *sis.BeforeFact bf     Processing message entity.
	// @param  interface{}     hook   The callback function for the decoded bounce message
	// @return bool                   true:  Successfully got the results
	//                                false: Failed to get the results
	if len(bf.Headers) == 0 { return false }
	if len(bf.Payload) == 0 { return false }

	// Tidy up each field name and value in the entire message body
	bf.Payload = *(tidy(&bf.Payload))

	// Decode BASE64 Encoded message body
	mesgformat := ""; if len(bf.Headers["content-type"]) > 0 {
		// Content-Type: text/plain; charset=utf8
		mesgformat = strings.ToLower(bf.Headers["content-type"][0])
	}
	ctencoding := ""; if len(bf.Headers["content-transfer-encoding"]) > 0 {
		// Content-Transfer-Encoding: base64
		ctencoding = strings.ToLower(bf.Headers["content-transfer-encoding"][0])
	}

	if strings.HasPrefix(mesgformat, "text/plain") || strings.HasPrefix(mesgformat, "text/html") {
		// Content-Type: text/plain; charset=UTF-8
		if ctencoding == "base64" {
			// Content-Transfer-Encoding: base64
			bf.Payload = rfc2045.DecodeB(bf.Payload, "")

		} else if ctencoding == "quoted-printable" {
			// Content-Transfer-Encoding: quoted-printable
			bf.Payload = rfc2045.DecodeQ(bf.Payload)
		}

		if strings.HasPrefix(mesgformat, "text/html") {
			// Content-Type: text/html;...
			bf.Payload = *(sisimoji.ToPlain(&bf.Payload))
		}
	} else if strings.HasPrefix(mesgformat, "multipart/") {
		// In case of Content-Type: multipart/*
		if cv := rfc2045.MakeFlat(bf.Headers["content-type"][0], &bf.Payload); cv != nil { bf.Payload = *cv }
	}
	bf.Payload = *(sisimoji.ToLF(&bf.Payload))
	bf.Payload = strings.ReplaceAll(bf.Payload, "\t", " ")

	cfargument := &sis.CallbackArgs{Headers: bf.Headers, Payload: &bf.Payload}
	cfreturned := hook.(func(*sis.CallbackArgs) map[string]interface{})(cfargument)
	havecalled := map[string]bool{}
	localhostr := sis.RisingUnderway{}
	modulename := ""
	DECODER: for {
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
		if len(localhostr.Digest[j].Agent) > 0 { continue }
		localhostr.Digest[j].Agent = modulename
	}

	// Convert headers of the original message to data structure/map[string][]string
	rfc822text := ""
	for _, e := range strings.Split(localhostr.RFC822, "\n") {
		// Append each line of localhostr.RFC822 to rfc822text except malformed headers
		if e == "" && rfc822text != ""  { break } // The blank line between the header and the body
		if strings.Index(e, ":") < 1 {            // The line does not contain ":" or begins with ":"
			// The line is not a line continued from the previous line of a long header
			if strings.HasPrefix(e, " ") == false || strings.HasPrefix(e, "\t") == false { continue }
		}
		rfc822text += e + "\n"
	}
	if rfc822text != "" { localhostr.RFC822 = rfc822text + "\n" }

	rfc822part, nyaan := mail.ReadMessage(strings.NewReader(localhostr.RFC822)); if nyaan != nil { return false }
	bf.RFC822 = makemap(&rfc822part.Header, false)
	bf.Digest = localhostr.Digest
	bf.Catch  = cfreturned

	return true
}

