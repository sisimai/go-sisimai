// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package message
import "strings"
import "sisimai/sis"
import "sisimai/lhost"
import "sisimai/rfc2045"
import sisimoji "sisimai/string"

// sift() sifts a bounce mail with each MTA module
func sift(bf *sis.BeforeFact, hook func()) bool {
	// @param  *sis.BeforeFact bf     Processing message entity.
	// @param  func()          hook   A hook method to be called
	// @return bool                   true:  Successfully got the results
	//                                false: Failed to get the results
	if len(bf.Head) == 0 { return false }
	if len(bf.Body) == 0 { return false }

	// Tidy up each field name and value in the entire message body
	bf.Body = *(tidy(&bf.Body))

	// Decode BASE64 Encoded message body
	mesgformat := ""; if len(bf.Head["content-type"])               > 0 { strings.ToLower(bf.Head["content-type"][0]) }
	ctencoding := ""; if len(bf.Head["content-trtansfer-encoding"]) > 0 { strings.ToLower(bf.Head["content-transfer-encoding"][0]) }

	if strings.HasPrefix(mesgformat, "text/plain") || strings.HasPrefix(mesgformat, "text/html") {
		// Content-Type: text/plain; charset=UTF-8
		if ctencoding == "base64" {
			// Content-Transfer-Encoding: base64
			bf.Body = *(rfc2045.DecodeB(&bf.Body, ""))

		} else {
			// Content-Transfer-Encoding: quoted-printable
			bf.Body = *(rfc2045.DecodeQ(&bf.Body))
		}

		if strings.HasPrefix(mesgformat, "text/html") {
			// Content-Type: text/html;...
			bf.Body = *(sisimoji.ToPlain(&bf.Body, true))
		}
	} else if strings.HasPrefix(mesgformat, "multipart/") {
		// In case of Content-Type: multipart/*
		bf.Body = *(rfc2045.MakeFlat(mesgformat, &bf.Body))
	}
	bf.Body = *(sisimoji.ToLF(&bf.Body))
	bf.Body = strings.ReplaceAll(bf.Body, "\t", " ")

	// TODO: Call the hook funcation
	hook()

	havecalled := map[string]bool{}
	localhostr := sis.RisingUnderway{}
	modulename := ""
	//havesifted := ""
	DECODER: for {
		// 1. MTA Module Candidates to be tried on first, and other sisimai/lhost/*.go
		// 2. sisimai/rfc3464
		// 3. sisimai/arf
		// 4. sisimai/rfc3834
		for _, r := range TryOnFirst {
			// 1. MTA Module Candidates to be tried on first, and other sisimai/lhost/*.go
			if r != "Postfix" { continue } // TODO: Implement all the lhost/*.go modules except postfix.go
			if havecalled[r]  { continue }

			localhostr    = lhost.LhostCode[r](bf)
			havecalled[r] = true
			modulename    = r

			if localhostr.Void() == false { break DECODER }
		}

		if havecalled["rfc3464"] == false {
			// TODO: Implemente sismai/rfc3464.go
			// 2. sisimai/rfc3464
			// When the all of sisimai/lhost/*.go modules did not return the decoded data
			if localhostr.Void() == false { break DECODER }
		}

		if havecalled["arf"] == false {
			// TODO: Implemente sismai/arf.go
			// 3. call sisimai/arf
			// Try to decode the message as a Feedback Loop message
			if localhostr.Void() == false { break DECODER }
		}

		if havecalled["rfc3834"] == false {
			// TODO: Implemente sismai/rfc3834.go
			// 4. call sisimai/rfc3834
			// Try to sift the message as auto reply message defined in RFC3834
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

	return true
}

