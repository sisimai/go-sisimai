// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _ __ ___   ___  ___ ___  __ _  __ _  ___ 
// | '_ ` _ \ / _ \/ __/ __|/ _` |/ _` |/ _ \
// | | | | | |  __/\__ \__ \ (_| | (_| |  __/
// |_| |_| |_|\___||___/___/\__,_|\__, |\___|
//                                |___/      

package message
import "strings"
import "net/mail"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/arf"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/lhost"
import "libsisimai.org/sisimai/v5/rfc2045"
import "libsisimai.org/sisimai/v5/rfc3464"
import "libsisimai.org/sisimai/v5/rfc3834"
import "libsisimai.org/sisimai/v5/rfc5322"

// sift is called from the Rise function and decode and structure various formats of bounce emails.
//   Arguments:
//     - bf (*sis.BeforeFact):    Message entity in progress.
//     - hook (sis.CfParameter0): The first callback function.
//   Returns:
//     - (bool): true = successfully decoded and structured the bounce emails, false = failed to decode.
func sift(bf *sis.BeforeFact, hook sis.CfParameter0) bool {
	if bf == nil || bf.IsEmpty() == true { return false }

	bf.Payload  = *(tidy(&bf.Payload)) // Tidy up each field name and value in the entire message body
	mesgformat := ""
	ctencoding := ""

	if len(bf.Headers["content-type"])              > 0 { mesgformat = strings.ToLower(bf.Headers["content-type"][0])              }
	if len(bf.Headers["content-transfer-encoding"]) > 0 { ctencoding = strings.ToLower(bf.Headers["content-transfer-encoding"][0]) }

	if moji.HasPrefixAny(mesgformat, []string{"text/plain", "text/html"}) {
		// Content-Type: text/plain; charset=UTF-8
		var nyaan error
		switch ctencoding {
			case "base64":           bf.Payload, nyaan = rfc2045.DecodeB(bf.Payload, "")
			case "quoted-printable": bf.Payload, nyaan = rfc2045.DecodeQ(bf.Payload)
		}
		if nyaan != nil {
			// Something wrong when the function decodes the Quoted-Printable encoded string
			ce := *sis.MakeNotDecoded(nyaan.Error(), false)
			bf.Errors = append(bf.Errors, ce)
		}
		if strings.HasPrefix(mesgformat, "text/html") { bf.Payload = *(moji.ToPlain(&bf.Payload)) }

	} else if strings.HasPrefix(mesgformat, "multipart/") {
		// In case of Content-Type: multipart/*
		cv, fe := rfc2045.MakeFlat(bf.Headers["content-type"][0], &bf.Payload)
		if cv != nil                 { bf.Payload = *cv                      }
		if fe != nil && len(fe) > 0  { bf.Errors  = append(bf.Errors, fe...) }
	}
	moji.ToLF(&bf.Payload)
	bf.Payload = strings.ReplaceAll(bf.Payload, "\t", " ") // Replace all the TAB with " "

	if hook != nil {
		// Execute the first callback function
		cvv, nyaan := hook(&sis.CallbackArg0{Headers: bf.Headers, Payload: &bf.Payload}); if nyaan != nil {
			// Something wrong when the 1st callback function executed
			ce := *sis.MakeNotDecoded(nyaan.Error(), false)
			bf.Errors = append(bf.Errors, ce)
		}
		bf.Catch = cvv
	}

	orders := lhost.OrderBySubject(bf.Headers["subject"][0])
	called := make(map[string]bool, 40)
	rising := &sis.RisingUnderway{}
	module := ""

	DECODER: for bf.IsEmpty() == false {
		// 1. MTA Module Candidates to be tried on first, and other lhost.InquireFor[*]
		// 2. rfc3464.Inquire()
		// 3. arf.Inquire()
		// 4. rfc3834.Inqquire()
		for _, r := range orders {
			// 1. MTA Module candidates to be tried on first, and other lhost.InquireFor[*]
			if called[r] || r == "ARF" || r == "RFC3834" { continue }
			called[r] = true
			if rising = lhost.InquireFor[r](bf); rising != nil { module = r; break DECODER }
		}
		if rising = rfc3464.Inquire(bf); rising != nil { module = "RFC3464"; break DECODER }
		if rising = arf.Inquire(bf);     rising != nil { module = "ARF";     break DECODER }
		if rising = rfc3834.Inquire(bf); rising != nil { module = "RFC3834"; break DECODER }

		break DECODER // as of now, we have no sample email for coding this block

	} // End of for(DECODER)
	if rising == nil || len(rising.Digest) == 0 { return false }

	for j := range rising.Digest {
		// Set the value of "Agent" such as "Postfix", "Sendmail", or "OpenSMTPD"
		if rising.Digest[j].Agent == "" { rising.Digest[j].Agent = module }
	}

	if strings.Contains(rising.RFC822, "\nFrom:") == false && len(bf.Headers["to"]) > 0 {
		// There is no "From:" header, pick the email address from the "To:" header of the
		// bounce message
		rising.RFC822 = "From: " + bf.Headers["to"][0] + "\n" + rising.RFC822
	}

	if di := &(rising.Digest[0]); strings.Contains(rising.RFC822, "\nTo:") == false && di.Recipient != "" {
		// The original message block is empty, insert some values picked from rising.Digest as
		// a pseudo header such as "To:", "Date:".
		rising.RFC822 = "To: <" + di.Recipient + ">\n" + rising.RFC822
	}

	// Convert headers of the original message to data structure/map[string][]string
	rfc822buff := strings.Builder{}; rfc822buff.Grow(len(rising.RFC822))
	for e := range strings.Lines(rising.RFC822) {
		// Append each line of rising.RFC822 to rfc822buff except malformed headers
		// The blank line between the header and the body
		if e = strings.TrimRight(e, "\n\r"); e == "" && rfc822buff.Len() > 0 { break }
		if strings.IndexByte(e, ':') < 1 {           // The line does not contain ":" or begins with ":"
			// The line is not a line continued from the previous line of a long header
			if strings.HasPrefix(e, " ") == false || strings.HasPrefix(e, "\t") == false { continue }
		}
		rfc822buff.WriteString(e + "\n")
	}
	if rfc822buff.Len() > 0 { rising.RFC822 = rfc822buff.String() + "\n" }

	rfc822part, nyaan := mail.ReadMessage(strings.NewReader(rising.RFC822))
	if nyaan != nil {
		// Failed to read the original message part
		ce := *sis.MakeNotDecoded(nyaan.Error(), false)
		bf.Errors = append(bf.Errors, ce)
		return false
	}
	bf.RFC822 = rfc5322.Headers(&rfc822part.Header, false)
	bf.Digest = rising.Digest

	return true
}

