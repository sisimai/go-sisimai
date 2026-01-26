// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ______          
//  _ __| |__   ___  ___| |_   / / ___|_____  __
// | '__| '_ \ / _ \/ __| __| / / |   / _ \ \/ /
// | |  | | | | (_) \__ \ |_ / /| |__| (_) >  < 
// |_|  |_| |_|\___/|___/\__/_/  \____\___/_/\_\

package rhost
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"

func init() {
	// ReturnedBy[*] detects the reason of the bounce returned by this email service.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (string): Bounce reason name or an empty string.
	ReturnedBy["Cox"] = func(fo *siba.Fact) string {
		// - Email Error Codes: https://www.cox.com/business/support/email-error-codes.html
		// - Feedback Loop Service https://www.cox.com/business/support/feedback-loop-service.html
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		errorcodes := map[string]string{
			// CXBL = Blocked
			// - The sending IP address has been blocked by Cox due to exhibiting spam-like behavior.
			// - Send an email request to Cox to ask for a sending IP address be unblocked.
			//   Note: Cox has sole discretion whether to unblock the sending IP address.
			"CXBL": eb.ReBLOC,

			// CXDNS = RequirePTR
			// - There was an issue with the connecting IP address Domain Name System (DNS).
			// - The Reverse DNS (rDNS) lookup for your IP address is failing. 
			//   - Confirm the IP address that sends your email.
			//   - Check the rDNS of that IP address. If it passes, then wait 24 hours and try resending
			//     your email.
			"CXDNS": eb.ReQPTR,

			// CXSNDR = AuthFailure
			// - There was a problem with the sender's domain.
			// - Your email failed authentication checks against your sending domain's SPF, DomainKeys,
			//   or DKIM policy.
			"CXSNDR": eb.ReAUTH,

			// CXSMTP = Rejected
			// - There was a violation of SMTP protocol.
			// - Your email wasn't delivered because Cox was unable to verify that it came from a
			//   legitimate email sender.
			"CXSMTP": eb.ReFROM,

			// CXCNCT = RateLimited
			// - There was a connection issue from the IP address.
			// - There is a limit to the number of concurrent SMTP connections per IP address to
			//   protect the systems against attack. Ensure that the sending email server is not
			//   opening more than 10 concurrent connections to avoid reaching this limit.
			"CXCNCT": eb.ReRATE,

			// CXMXRT = RateLimited
			//   - The sender has sent email to too many recipients and needs to wait before sending
			//     more email.
			//   - The email sender has exceeded the maximum number of sent email allowed.
			"CXMXRT": eb.ReRATE,

			// CDRBL = Blocked
			// - The sending IP address has been temporarily blocked by Cox due to exhibiting spam-like
			//   behavior.
			// - The block duration varies depending on reputation and other factors, but will not exceed
			//   24 hours. Inspect email traffic for potential spam, and retry email delivery.
			"CDRBL": eb.ReBLOC,

			"CXTHRT":    eb.ReSAFE, // Email sending limited due to suspicious account activity.
			"CXMJ":      eb.ReSAFE, // Email sending blocked due to suspicious account activity on primary Cox account.
			"IPBL0001":  eb.ReBLOC, // The sending IP address is listed in the Spamhaus Zen DNSBL.
			"IPBL0010":  eb.ReBLOC, // The sending IP is listed in the Return Path DNSBL.
			"IPBL0100":  eb.ReBLOC, // The sending IP is listed in the Invaluement ivmSIP DNSBL.
			"IPBL0011":  eb.ReBLOC, // The sending IP is in the Spamhaus Zen and Return Path DNSBLs.
			"IPBL0101":  eb.ReBLOC, // The sending IP is in the Spamhaus Zen and Invaluement ivmSIP DNSBLs.
			"IPBL0110":  eb.ReBLOC, // The sending IP is in the Return Path and Invaluement ivmSIP DNSBLs.
			"IPBL0111":  eb.ReBLOC, // The sending IP is in the Spamhaus Zen, Return Path and Invaluement ivmSIP DNSBLs.
			"IPBL1000":  eb.ReBLOC, // The sending IP address is listed on a CSI blacklist.
			"IPBL1001":  eb.ReBLOC, // The sending IP is listed in the Cloudmark CSI and Spamhaus Zen DNSBLs.
			"IPBL1010":  eb.ReBLOC, // The sending IP is listed in the Cloudmark CSI and Return Path DNSBLs.
			"IPBL1011":  eb.ReBLOC, // The sending IP is in the Cloudmark CSI, Spamhaus Zen and Return Path DNSBLs.
			"IPBL1100":  eb.ReBLOC, // The sending IP is listed in the Cloudmark CSI and Invaluement ivmSIP DNSBLs.
			"IPBL1101":  eb.ReBLOC, // The sending IP is in the Cloudmark CSI, Spamhaus Zen and Invaluement IVMsip DNSBLs.
			"IPBL1110":  eb.ReBLOC, // The sending IP is in the Cloudmark CSI, Return Path and Invaluement ivmSIP DNSBLs.
			"IPBL1111":  eb.ReBLOC, // The sending IP is in the Cloudmark CSI, Spamhaus Zen, Return Path and Invaluement ivmSIP DNSBLs.
			"IPBL00001": eb.ReBLOC, // The sending IP address is listed on a Spamhaus blacklist.
			"URLBL011":  eb.ReSPAM, // A URL within the body of the message was found on blocklists SURBL and Spamhaus DBL.
			"URLBL101":  eb.ReSPAM, // A URL within the body of the message was found on blocklists SURBL and ivmURI.
			"URLBL110":  eb.ReSPAM, // A URL within the body of the message was found on blocklists Spamhaus DBL and ivmURI.
			"URLBL1001": eb.ReSPAM, // The URL is listed on a Spamhaus blacklist.
		}
		messagesof := map[string][]string{
			eb.ReBLOC: []string{ // Blocked
				// - An email client has repeatedly sent bad commands or invalid passwords resulting
				//   in a three-hour block of the client's IP address.
				// - The sending IP address has exceeded the threshold of invalid recipients and has
				//   been blocked.
				"cox too many bad commands from",
				"too many invalid recipients",
			},
			eb.ReQPTR: []string{ // RequirePTR
				// - The reverse DNS check of the sending server IP address has failed.
				// - Cox requires that all connecting email servers contain valid reverse DNS PTR records.
				"dns check failure - try again later",
				"rejected - no rdns",
			},
			eb.ReWONT: []string{ // PolicyViolation
				// - The sending server has attempted to communicate too soon within the SMTP transaction
				// - The message has been rejected because it contains an attachment with one of the
				//   following prohibited file types, which commonly contain viruses: .shb, .shs, .vbe,
				//   .vbs, .wsc, .wsf, .wsh, .pif, .msc, .msi, .msp, .reg, .sct, .bat, .chm, .isp, .cpl,
				//   .js, .jse, .scr, .exe.
				"esmtp no data before greeting",
				"attachment extension is forbidden",
			},
			eb.ReFROM: []string{ // Rejected
				// Cox requires that all sender domains resolve to a valid MX or A-record within DNS.
				"sender rejected",
			},
			eb.RePROC: []string{ // SystemError
				// - Our systems are experiencing an issue which is causing a temporary inability to
				//   accept new email.
				"esmtp server temporarily not available",
			},
			eb.ReRATE: []string{ // RateLimited
				// - The sending IP address has exceeded the five maximum concurrent connection limit.
				// - The SMTP connection has exceeded the 100 email message threshold and was disconnected.
				// - The sending IP address has exceeded one of these rate limits and has been temporarily
				//   blocked.
				// - Cox enforces various rate limits to protect our platform. The sending IP address
				//   has exceeded one of these rate limits and has been temporarily blocked.
				"too many sessions from",
				"requested action aborted: try again later",
				"message threshold exceeded",
			},
			eb.ReUSER: []string{ // UserUnknown
				// - The intended recipient is not a valid Cox Email account.
				"recipient rejected",
			},
		}

		issuedcode := fo.DiagnosticCode + " "
		codenumber := moji.Select(issuedcode, "AUP#", " ", 0)
		reasontext := errorcodes[codenumber]; if reasontext != "" { return reasontext }
		issuedcode  = strings.ToLower(issuedcode)
		for e := range messagesof {
			// Try to find with each error message defined in "messagesof"
			if moji.ContainsAny(issuedcode, messagesof[e]) { return e }
		}
		return ""
	}
}

