// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package arf

//     _    ____  _____ 
//    / \  |  _ \|  ___|
//   / _ \ | |_) | |_   
//  / ___ \|  _ <|  _|  
// /_/   \_\_| \_\_|    
import "fmt"
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/lhost"
import "libsisimai.org/sisimai/rfc1894"
import "libsisimai.org/sisimai/rfc5322"
import sisimoji "libsisimai.org/sisimai/string"
import sisiaddr "libsisimai.org/sisimai/address"

// isARF() returns true if the bounce mail is Abuse Feedback Reporting Format
func isARF(bf *sis.BeforeFact) bool {
	// @param    *sis.BeforeFact bf  Message body of a bounce email
	// @return   bool                true if the mail is ARF
	// @see      https://tools.ietf.org/html/rfc5965
	if bf == nil || bf.Empty() == true { return false }

	abuse := []string{"staff@hotmail.com", "complaints@email-abuse.amazonses.com"}
	ctype := bf.Headers["content-type"][0]

	// Content-Type: multipart/report; report-type=feedback-report; ...
	if sisimoji.Aligned(ctype, []string{"report-type=", "feedback-report"}) { return true }
	if strings.Contains(ctype, "multipart/mixed") {
		// Microsoft (Hotmail, MSN, Live, Outlook) uses its own report format.
		// Amazon SES Complaints bounces
		if strings.Contains(bf.Headers["subject"][0], "complaint about message from ") {
			// From: staff@hotmail.com
			// From: complaints@email-abuse.amazonses.com
			// Subject: complaint about message from 192.0.2.1
			if sisimoji.ContainsAny(bf.Headers["from"][0], abuse) { return true }
		}
	}

	APPLE: for {
		// X-Apple-Unsubscribe: true
		if len(bf.Headers["x-apple-unsubscribe"]) == 0      { break APPLE }
		if bf.Headers["x-apple-unsubscribe"][0]   == "true" { return true }
		break APPLE
	}

	return false
}

// Inquire() decodes a bounce message that is ARF: Abuse Feedback Reporting Formatted email
func Inquire(bf *sis.BeforeFact) sis.RisingUnderway {
	// @param    *sis.BeforeFact bf  Message body of a bounce email
	// @return   RisingUnderway      RisingUnderway structure
	// @see      https://tools.ietf.org/html/rfc5965
	if bf == nil || bf.Empty() || isARF(bf) == false { return sis.RisingUnderway{} }

	// http://tools.ietf.org/html/rfc5965
	// http://en.wikipedia.org/wiki/Feedback_loop_(email)
	// http://en.wikipedia.org/wiki/Abuse_Reporting_Format
	//
	// Netease DMARC uses:    This is a spf/dkim authentication-failure report for an email message received from IP
	// OpenDMARC 1.3.0 uses:  This is an authentication failure report for an email message received from IP
	// Abusix ARF uses:       this is an autogenerated email abuse complaint regarding your network.
	indicators := lhost.INDICATORS()
	boundaries := []string{
		"Content-Type: message/rfc822",
		"Content-Type: text/rfc822-headers",
		"Content-Type: text/rfc822-header",  // ?
	}
	reportpart := false
	reportfrom := "Content-Type: message/feedback-report"
	arfpreface := [][]string{
		[]string{"this is a", "abuse report"},
		[]string{"this is a", "authentication", "failure report"},
		[]string{"this is a", " report for"},
		[]string{"this is an authentication", "failure report"},
		[]string{"this is an autogenerated email abuse complaint"},
		[]string{"this is an email abuse report"},
	}

	dscontents := []sis.DeliveryMatter{{}}
	emailparts := rfc5322.Part(&bf.Payload, boundaries, false)
	readcursor := uint8(0)            // Points the current cursor position
	recipients := uint8(0)            // The number of "Final-Recipient" header
	timestamp0 := ""                  // The value of "Arrival-Date" or "Received-Date"
	remotehost := ""                  // The value of "Source-IP" field
	reportedby := ""                  // The value of "Reporting-MTA" field
	anotherone := ""                  // Other fields(append to Diagnosis)
	v          := &(dscontents[len(dscontents) - 1])

    // 3.1.  Required Fields
    //
    //   The following report header fields MUST appear exactly once:
    //
    //   o  "Feedback-Type" contains the type of feedback report (as defined
    //      in the corresponding IANA registry and later in this memo).  This
    //      is intended to let report parsers distinguish among different
    //      types of reports.
    //
    //   o  "User-Agent" indicates the name and version of the software
    //      program that generated the report.  The format of this field MUST
    //      follow section 14.43 of [HTTP].  This field is for documentation
    //      only; there is no registry of user agent names or versions, and
    //      report receivers SHOULD NOT expect user agent names to belong to a
    //      known set.
    //
    //   o  "Version" indicates the version of specification that the report
    //      generator is using to generate the report.  The version number in
    //      this specification is set to "1".
	for _, e := range(strings.Split(emailparts[0], "\n")) {
		// Read error messages and delivery status lines from the head of the email to the
		// previous line of the beginning of the original message.
		if readcursor == 0 {
			// Beginning of the bounce message or message/delivery-status part
			r := strings.ToLower(e); for _, f := range arfpreface {
				// Hello,
				// this is an autogenerated email abuse complaint regarding your network.
				if sisimoji.Aligned(r, f) == false { continue }
				readcursor |= indicators["deliverystatus"]
				v.Diagnosis += " " + e
				break
			}
			continue
		}
		if readcursor & indicators["deliverystatus"] == 0 { continue }
		if len(e) == 0                                    { continue }
		if e == reportfrom             { reportpart = true; continue }

		if reportpart {
			// Content-Type: message/feedback-report
			// MIME-Version: 1.0
			//
			// Feedback-Type: abuse
			// User-Agent: SomeGenerator/1.0
			// Version: 0.1
			// Original-Mail-From: <somespammer@example.net>
			// Original-Rcpt-To: <kijitora@example.jp>
			// Received-Date: Thu, 29 Apr 2009 00:00:00 JST
			// Source-IP: 192.0.2.1
			if strings.HasPrefix(e, "Original-Rcpt-To: ") || strings.HasPrefix(e, "Removal-Recipient: ") {
				// Original-Rcpt-To header field is optional and may appear any number of times as appropriate:
				// Original-Rcpt-To: <kijitora@example.jp>
				// Removal-Recipient: user@example.com
				cv := sisiaddr.S3S4(e[strings.Index(e, " ") + 1:]); if sisiaddr.IsEmailAddress(cv) == false       { continue }
				cw := len(dscontents);                              if cw > 0 && cv == dscontents[cw-1].Recipient { continue }

				if len(v.Recipient) > 0 {
					// There are multiple recipient addresses in the message body.
					dscontents = append(dscontents, sis.DeliveryMatter{})
					v = &(dscontents[len(dscontents) - 1])
				}
				v.Recipient = cv
				recipients += 1

			} else if strings.HasPrefix(e, "Feedback-Type: ") {
				// The header field MUST appear exactly once.
				// Feedback-Type: abuse
				v.FeedbackType = e[strings.Index(e, " ") + 1:]

			} else if strings.HasPrefix(e, "Authentication-Results: ") {
				// "Authentication-Results" indicates the result of one or more authentication checks
				// run by the report generator.
				//
				// Authentication-Results: mail.example.com;
				//   spf=fail smtp.mail=somespammer@example.com
				anotherone += e + ", "

			} else if strings.HasPrefix(e, "User-Agent: ") {
				// The header field MUST appear exactly once.
				// User-Agent: SomeGenerator/1.0
				anotherone += e + ", "

			} else if strings.HasPrefix(e, "Received-Date: ") || strings.HasPrefix(e, "Arrival-Date: ") {
				// Arrival-Date header is optional and MUST NOT appear more than once.
				// Received-Date: Thu, 29 Apr 2010 00:00:00 JST
				// Arrival-Date: Thu, 29 Apr 2010 00:00:00 +0000
				timestamp0 = e[strings.Index(e, " ") + 1:]

			} else if strings.HasPrefix(e, "Reporting-MTA: ") {
				// The header is optional and MUST NOT appear more than once.
				// Reporting-MTA: dns; mx.example.jp
				cv := rfc1894.Field(e); if len(cv) == 0 { continue }
				reportedby = cv[2]

			} else if strings.HasPrefix(e, "Source-IP: ") {
				// The header is optional and MUST NOT appear more than once.
				// Source-IP: 192.0.2.45
				remotehost = e[strings.Index(e, " ") + 1:]

			} else if strings.HasPrefix(e, "Original-Mail-From: ") {
				// the header is optional and MUST NOT appear more than once.
				// Original-Mail-From: <somespammer@example.net>
				anotherone += e + ", "
			}
		} else {
			// Messages before "Content-Type: message/feedback-report" part
			v.Diagnosis += " " + e
		}
	}

	for recipients == 0 {
		// There is no recipient address in the message
		if len(bf.Headers["x-apple-unsubscribe"]) > 0 {
			// X-Apple-Unsubscribe: true
			if bf.Headers["x-apple-unsubscribe"][0]         != "true" { break }
			if strings.Contains(bf.Headers["from"][0], "@") == false  { break }
			dscontents[0].Recipient    = bf.Headers["from"][0]
			dscontents[0].Diagnosis    = sisimoji.Sweep(emailparts[0])
			dscontents[0].FeedbackType = "opt-out"

			// Addpend To: field as a pseudo header
			if emailparts[1] == "" { emailparts[1] = fmt.Sprintf("To: <%s>\n", bf.Headers["from"][0]) }

		} else {
			// Pick it from the original message part
			p1 := strings.Index(emailparts[1], "\nTo:");               if p1  < 0  { break }
			p2 := sisimoji.IndexOnTheWay(emailparts[1], "\n", p1 + 4); if p2  < 0  { break }
			cv := sisiaddr.S3S4(emailparts[1][p1 + 4:p2])

			// There is no valid email address in the To: header of the original message such as
			// To: <Undisclosed Recipients>
			if cv == "" { cv = sisiaddr.Undisclosed(true) }
			dscontents[0].Recipient = cv
		}
		recipients++
	}
	if recipients == 0 { return sis.RisingUnderway{} }

	if anotherone != "" { anotherone = ": " + strings.TrimRight(sisimoji.Sweep(anotherone), ",") }
	for j, _ := range dscontents {
		// Tidy up the error message in e.Diagnosis, Try to detect the bounce reason.
		e := &(dscontents[j])
		e.Diagnosis = sisimoji.Sweep(e.Diagnosis + anotherone)
		e.Reason    = "feedback"
		e.Rhost     = remotehost
		e.Lhost     = reportedby
		e.Date      = timestamp0

		// Copy some values from the previous element when the report have 2 or more email address
		if j == 0 || len(dscontents) == 1 { continue }
		p := &(dscontents[j - 1])
		if e.Diagnosis    == "" { e.Diagnosis    = p.Diagnosis    }
		if e.FeedbackType == "" { e.FeedbackType = p.FeedbackType }
	}

	return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
}

