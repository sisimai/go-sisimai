// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//       _               _      ___                          _     
//  _ __| |__   ___  ___| |_   / / \   _ __ ___   __ ___   _(_)___ 
// | '__| '_ \ / _ \/ __| __| / / _ \ | '_ ` _ \ / _` \ \ / / / __|
// | |  | | | | (_) \__ \ |_ / / ___ \| | | | | | (_| |\ V /| \__ \
// |_|  |_| |_|\___/|___/\__/_/_/   \_\_| |_| |_|\__,_| \_/ |_|___/
import "slices"
import "strings"
import "sisimai/sis"
import "sisimai/rfc1894"
import "sisimai/rfc5322"
import sisimoji "sisimai/string"

func init() {
	// Decode bounce messages from amavisd-new: https://www.amavis.org/
	InquireFor["Amavis"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		// From: "Content-filter at neko1.example.jp" <postmaster@neko1.example.jp>
		// Subject: Undeliverable mail, MTA-BLOCKED
		if strings.Index(bf.Head["from"][0], `"Content-filter at `) != 0 { return sis.RisingUnderway{} }

		indicators := INDICATORS()
		boundaries := []string{"Content-Type: text/rfc822-headers"}
		startingof := map[string][]string{ "message": []string{"The mesage "} }
		messagesof := map[string][]string{
			// amavisd-new-2.11.1/amavisd:1840|%smtp_reason_by_ccat = (
			// amavisd-new-2.11.1/amavisd:1840|  # currently only used for blocked messages only, status 5xx
			// amavisd-new-2.11.1/amavisd:1840|  # a multiline message will produce a valid multiline SMTP response
			// amavisd-new-2.11.1/amavisd:1840|  CC_VIRUS,       'id=%n - INFECTED: %V',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BANNED,      'id=%n - BANNED: %F',
			// amavisd-new-2.11.1/amavisd:1840|  CC_UNCHECKED.',1', 'id=%n - UNCHECKED: encrypted',
			// amavisd-new-2.11.1/amavisd:1840|  CC_UNCHECKED.',2', 'id=%n - UNCHECKED: over limits',
			// amavisd-new-2.11.1/amavisd:1840|  CC_UNCHECKED,      'id=%n - UNCHECKED',
			// amavisd-new-2.11.1/amavisd:1840|  CC_SPAM,        'id=%n - spam',
			// amavisd-new-2.11.1/amavisd:1840|  CC_SPAMMY.',1', 'id=%n - spammy (tag3)',
			// amavisd-new-2.11.1/amavisd:1840|  CC_SPAMMY,      'id=%n - spammy',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH.',1',   'id=%n - BAD HEADER: MIME error',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH.',2',   'id=%n - BAD HEADER: nonencoded 8-bit character',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH.',3',   'id=%n - BAD HEADER: contains invalid control character',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH.',4',   'id=%n - BAD HEADER: line made up entirely of whitespace',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH.',5',   'id=%n - BAD HEADER: line longer than RFC 5322 limit',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH.',6',   'id=%n - BAD HEADER: syntax error',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH.',7',   'id=%n - BAD HEADER: missing required header field',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH.',8',   'id=%n - BAD HEADER: duplicate header field',
			// amavisd-new-2.11.1/amavisd:1840|  CC_BADH,        'id=%n - BAD HEADER',
			// amavisd-new-2.11.1/amavisd:1840|  CC_OVERSIZED,   'id=%n - Message size exceeds recipient\'s size limit',
			// amavisd-new-2.11.1/amavisd:1840|  CC_MTA.',1',    'id=%n - Temporary MTA failure on relaying',
			// amavisd-new-2.11.1/amavisd:1840|  CC_MTA.',2',    'id=%n - Rejected by next-hop MTA on relaying',
			// amavisd-new-2.11.1/amavisd:1840|  CC_MTA,         'id=%n - Unable to relay message back to MTA',
			// amavisd-new-2.11.1/amavisd:1840|  CC_CLEAN,       'id=%n - CLEAN',
			// amavisd-new-2.11.1/amavisd:1840|  CC_CATCHALL,    'id=%n - OTHER',  # should not happen
			// ...
			// amavisd-new-2.11.1/amavisd:15289|my $status = setting_by_given_contents_category(
			// amavisd-new-2.11.1/amavisd:15289|  $blocking_ccat,
			// amavisd-new-2.11.1/amavisd:15289|  { CC_VIRUS,       "554 5.7.0",
			// amavisd-new-2.11.1/amavisd:15289|    CC_BANNED,      "554 5.7.0",
			// amavisd-new-2.11.1/amavisd:15289|    CC_UNCHECKED,   "554 5.7.0",
			// amavisd-new-2.11.1/amavisd:15289|    CC_SPAM,        "554 5.7.0",
			// amavisd-new-2.11.1/amavisd:15289|    CC_SPAMMY,      "554 5.7.0",
			// amavisd-new-2.11.1/amavisd:15289|    CC_BADH.",2",   "554 5.6.3",  # nonencoded 8-bit character
			// amavisd-new-2.11.1/amavisd:15289|    CC_BADH,        "554 5.6.0",
			// amavisd-new-2.11.1/amavisd:15289|    CC_OVERSIZED,   "552 5.3.4",
			// amavisd-new-2.11.1/amavisd:15289|    CC_MTA,         "550 5.3.5",
			// amavisd-new-2.11.1/amavisd:15289|    CC_CATCHALL,    "554 5.7.0",
			// amavisd-new-2.11.1/amavisd:15289|  });
			// ...
			// amavisd-new-2.11.1/amavisd:15332|my $response = sprintf("%s %s%s%s", $status,
			// amavisd-new-2.11.1/amavisd:15333|  ($final_destiny == D_PASS     ? "Ok" :
			// amavisd-new-2.11.1/amavisd:15334|   $final_destiny == D_DISCARD  ? "Ok, discarded" :
			// amavisd-new-2.11.1/amavisd:15335|   $final_destiny == D_REJECT   ? "Reject" :
			// amavisd-new-2.11.1/amavisd:15336|   $final_destiny == D_BOUNCE   ? "Bounce" :
			// amavisd-new-2.11.1/amavisd:15337|   $final_destiny == D_TEMPFAIL ? "Temporary failure" :
			// amavisd-new-2.11.1/amavisd:15338|                                  "Not ok ($final_destiny)" ),
			"spamdetected":  []string{" - spam"},
			"virusdetected": []string{" - infected"},
			"contenterror":  []string{" - bad header:"},
			"exceedlimit":   []string{" - message size exceeds recipient"},
			"systemerror":   []string{
				" - temporary mta failure on relaying",
				" - rejected by next-hop mta on relaying",
				" - unable to relay message back to mta",
			},
		}

		fieldtable := rfc1894.FIELDTABLE()
		permessage := map[string]string{} // Store values of each Per-Message field
		keystrings := []string{}          // Key list of permessage
		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		recipients := uint8(0)     // The number of 'Final-Recipient' header
		readcursor := uint8(0)     // Points the current cursor position
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				if strings.HasPrefix(e, startingof["message"][0]) { readcursor |= indicators["deliverystatus"] }
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }
			if len(e)                                    == 0 { continue }
			f := rfc1894.Match(e);             if f      == 0 { continue }
			o := rfc1894.Field(e);             if len(o) == 0 { continue }
			z := fieldtable[o[0]]
			v  = &(dscontents[len(dscontents) - 1])

			if o[3] == "addr" {
				// Final-Recipient: rfc822; kijitora@example.jp
				// X-Actual-Recipient: rfc822; kijitora@example.co.jp
				if o[0] == "final-recipient" {
					// Final-Recipient: rfc822; kijitora@example.jp
					if len(v.Recipient) > 0 {
						// There are multiple recipient addresses in the message body.
						dscontents = append(dscontents, sis.DeliveryMatter{})
						v = &(dscontents[len(dscontents) - 1])
					}
					v.Recipient = o[2]
					recipients += 1

				} else {
					// X-Actual-Recipient: rfc822; kijitora@example.co.jp
					v.Alias = o[2]
				}
			} else if o[3] == "code" {
				// Diagnostic-Code: SMTP; 550 5.1.1 <userunknown@example.jp>... User Unknown
				v.Spec = o[1]
				if strings.ToUpper(o[1]) == "X-POSTFIX" { v.Spec = "SMTP" }
				v.Diagnosis = o[2]

			} else {
				// Other DSN fields defined in RFC3464
				v.Set(o[0], o[2]); if f != 1 { continue }

				// Copy the lower-cased member name of DeliveryMatter{} for "permessage"
				permessage[z] = o[2]
				if slices.Contains(keystrings, z) == false { keystrings = append(keystrings, z) }
			}
		}
		if recipients == 0 { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Set default values stored in "permessage" if each value in "dscontents" is empty.
			e := &(dscontents[j])
			for _, z := range keystrings {
				// Do not set an empty string into each member of DeliveryMatter{}
				if len(v.Get(z))       > 0 { continue }
				if len(permessage[z]) == 0 { continue }
				e.Set(z, permessage[z])
			}

			if e.Diagnosis == "" { sisimoji.Sweep(e.Diagnosis) }
			ce := strings.ToLower(e.Diagnosis)

			FINDREASON: for p := range messagesof {
				// The key name is a bounce reason
				for _, r := range messagesof[p] {
					// Try to find an error message including lower-cased string listed in messagesof
					if strings.Contains(ce, r) == false { continue }
					e.Reason = p; break FINDREASON
				}
			}
		}
		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

