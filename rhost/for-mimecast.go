// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ____  __ _                              _   
//  _ __| |__   ___  ___| |_   / /  \/  (_)_ __ ___   ___  ___ __ _ ___| |_ 
// | '__| '_ \ / _ \/ __| __| / /| |\/| | | '_ ` _ \ / _ \/ __/ _` / __| __|
// | |  | | | | (_) \__ \ |_ / / | |  | | | | | | | |  __/ (_| (_| \__ \ |_ 
// |_|  |_| |_|\___/|___/\__/_/  |_|  |_|_|_| |_| |_|\___|\___\__,_|___/\__|

package rhost
import "strings"
import "libsisimai.org/sisimai/sis"
import "libsisimai.org/sisimai/smtp/reply"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["Mimecast"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		if fo == nil || fo.DiagnosticCode == "" { return "" }
		if reply.Test(fo.ReplyCode) == false    { return "" }

		messagesof := map[string][][2]string{
			// https://community.mimecast.com/s/article/Mimecast-SMTP-Error-Codes-842605754
			// https://community.mimecast.com/s/article/email-security-cloud-gateway-mimecast-smtp-error-codes
			"authfailure": [][2]string{
				// - The inbound message has been rejected because the originated IP address isn"t list-
				//   ed in the published SPF records for the sending domain.
				// - Ensure all the IP addresses for your mail servers are listed in your SPF records.
				//   Alternatively, create a DNS Authentication (Inbound / Outbound) policy with the
				//   "Inbound SPF" or "Reject on Hard Fail" option disabled. Messages that fail our SPF
				//   checks are subjected to spam and RBL checks, instead of being rejected.
				[2]string{"550", "spf sender invalid - envelope rejected"},

				// - The DKIM key for the outbound message is broken and doesn't match the DNS record of
				//   the registered sender.
				// - Check your organization's DNS record is populated with the right public key as part
				//   of the DNS Authentication Outbound Signing definition. The private key of the key-
				//   pair must be populated in the DNS Authentication policy, along with the domain and
				//   selector of that record.
				[2]string{"550", "dkim sender invalid - envelope rejected"},

				// - The inbound message has been rejected because the originated IP address isn't list-
				//   ed in the published SPF records for the sending domain.
				// - Ensure all the IP addresses for your mail servers are listed in your SPF records.
				[2]string{"550", "dmarc sender invalid - envelope rejected"},
			},
			"badreputation": [][2]string{
				// - The sending mail server is subjected to Greylisting. This requires the server to
				//   retry the connection, between one minute and 12 hours. Alternatively, the sender's
				//   IP address has a poor reputation.
				// - These reputation checks can be bypassed with an Auto Allow or Permitted Senders
				//   policy. If it's legitimate traffic, amend your Greylisting policy.
				[2]string{"451", "internal resources are temporarily unavailable"},

				// - Ongoing reputation checks have resulted in the message being rejected due to poor
				//   IP reputation. This could occur after a 4xx error.
				// - Create an Auto Allow or Permitted Senders policy.
				//   Note:
				//     You can request a review of your source IP ranges by completing our online form.
				[2]string{"550", "local ct ip reputation - (reject)"},
			},
			"blocked": [][2]string{
				// - The sender"s IP address has been blocked by a Blocked Senders Policy.
				// - Remove the entry from the policy.
				[2]string{"421", "sender address blocked"},

				// - The Sender's IP address has been placed on the block list due to too many invalid
				//   connections.
				// - The sender's mail server must retry the connection. The mail server performing the
				//   connection says the recipient address validation isn't responding.
				[2]string{"451", "recipient temporarily unavailable"},

				// - You've reached your mail server's limit.
				// - Wait and try again. The mail server won't accept any messages until you're under
				//   the limit.
				[2]string{"451", "ip temporarily blacklisted"},

				// - The sender's IP address is listed in an RBL. The text displayed is specific to the
				//   RBL which lists the sender's IP address.
				// - Bypass the RBL with an Auto Allow or Permitted Senders policy. Additionally request
				//   the associated IP address from the RBL.
				//[550, '< details of RBL >'}, NEED AN ACTUAL ERROR MESSAGE STRING
			},
			"expired": [][2]string{
				// - Journal messages past the expiration
				// - Attempts are being made to journal mail past the set expiry threshold.
				//   A retry response will replace the failure because the message is marked for retry
				//   if rejected, causing the journal queue to grow.
				// - Check to confirm there are no significant time discrepancies on the mail server.
				//   Discontinue journaling old messages past the expiry threshold.
				[2]string{"550", "Journal messages past the expiration"},
			},
			"failedstarttls": [][2]string{
				// - This email has been sent using SMTP, but TLS is required by policy.
				// - Delete or change the Secure Receipt or Secure Delivery policy enforcing TLS.
				//   Alternatively, ensure the certificates on the mail server haven't expired. If using
				//   a proxy server, ensure it isn't intercepting the traffic and modifying encryption
				//   parameters.
				[2]string{"553", "this route requires encryption (tls)"},

				// - A TLS connection has been attempted using a TLS version that is lower than TLS 1.2.
				// - Delete or change the Secure Receipt or Secure Delivery policy enforcing TLS.
				//   Alternatively, ensure the mail server attempting to connect is using the appropri-
				//   ate version of TLS.
				[2]string{"553", "this route requires tls version 1.2 or greater"},

				// - A secure connection was attempted using ciphers that do not meet the configured ci-
				//   pher strength.
				// - Delete or change the Secure Receipt or Secure Delivery policy enforcing TLS. Alter-
				//   natively, ensure the certificates on the mail server haven't expired. If using a
				//   proxy server, ensure it isn't intercepting the traffic and modifying encryption
				//   parameters.
				[2]string{"553", "this route requires high-strength ciphers"},

				// - Validation on your umbrella account's domain name does not conform to your DNS.
				// - Check you DNS has the required umbrella accounts listed as comma-separated values.
				[2]string{"554", "configuration is invalid for this certificate"},
			},
			"mesgtoobig": [][2]string{
				// - The email size either exceeds an Email Size Limit policy or is larger than the
				//   Mimecast service limit. The default is 100 MB for the Legacy MTA, and 200 MB for
				//   "the Latest MTA".
				// - Resend the message ensuring it's smaller than the limitation set. The transmission
				//   and content-encoding can add significantly to the total message size (e.g. a mes-
				//   sage with a 70 MB attachment, can have an overall size larger than 100 MB).
				[2]string{"554", "maximum email size exceeded"},
			},
			"networkerror": [][2]string{
				// - The recipients' domains have MX records configured incorrectly
				// - Check and remove any MX records that point to hostnames with outbound references.
				//   Only Inbound smart hosts are supported on MX records.
				[2]string{"451", "the incorrect hostname used for inbounds"},
				[2]string{"550", "the incorrect hostname used for inbounds"},

				// - The message has too many "received headers" as it has been forwarded across multi-
				//   ple hops. Once 25 hops have been reached, the email is rejected.
				// - Investigate the email addresses in the communication pairs, to see what forwarders
				//   are configured on the mail servers.
				[2]string{"554", "mail loop detected"},
			},
			"norelaying": [][2]string{
				// - Both the sender and recipient domains specified in the transmission are external to
				//   Mimecast, and aren't allowed to relay through the Mimecast service and/or the con-
				//   necting IP address isn't recognized as authorized.
				// - Mimecast customers should contact Mimecast Support to add the Authorized Outbound
				//   address, or to take other remedial action.
				[2]string{"451", "open relay not allowed"},
				[2]string{"451", "open relay is not allowed"},
			},
			"notaccept": [][2]string{
				// - The customer account Inbound emails are disabled in the Administration Console.
				// - Contact Mimecast Support if the account's inbound traffic should be allowed.
				[2]string{"451", "account inbounds disabled"},
			},
			"onhold": [][2]string{
				// - The customer account outbound emails are disabled in the Administration Console.
				// - Contact Mimecast Support if the account's outbound traffic should be allowed.
				[2]string{"451", "account outbounds disabled"},

				// - Omni Directional hostnames are enabled.
				// - Disable Omni Directional hostnames.
				[2]string{"451", "hostname is not authorized"},

				// - Attempts are being made to journal mail that is past the set expiry threshold. The
				//   failure will be replaced by a retry response because the message is marked for re-
				//   try if rejected, causing the journal queue to grow.
				// - Check to confirm there are no significant time discrepancies on the mail server.
				//   Discontinue journaling old messages past the expiry threshold.
				[2]string{"550", "journal message past expiration"},
			},
			"policyviolation": [][2]string{
				// - The message has triggered an Anti-Spoofing policy.
				// - Create an Anti-Spoofing policy to take no action for the sender's address or IP ad-
				//   dress.
				[2]string{"550", "anti-spoofing policy - inbound not allowed"},
				[2]string{"550", "rejected by header-based anti-spoofing policy"},

				// - The message has triggered a Content Examination policy.
				// - Either create a Content Examination Bypass policy or adjust the Content Examination
				//   policy as required.
				[2]string{"550Z", "message bounced due to content examination policy"},

				// - The message has triggered a Geographical Restriction policy.
				// - Delete or amend the policy.
				[2]string{"554", "host network not allowed"},
				[2]string{"554", "host network, not allowed"},
			},
			"rejected": [][2]string{
				// - The sender's email address or domain has triggered a Blocked Senders Policy or
				//   there's an SPF hard rejection.
				// - Delete or modify the Blocked Senders policy to exclude the sender address.
				[2]string{"550", "administrative prohibition envelope blocked"},

				// - A personal block policy is in place for the email address/domain.
				// - Remove the email address/domain from the Managed Senders list.
				[2]string{"550", "envelope blocked - user entry"},
				[2]string{"550", "envelope blocked - user domain entry"},
				[2]string{"550", "rejected by header-based manually blocked senders - block for manual block"},

				// - A Block Sender Policy has been applied to reject emails based on the Header From or
				//   Envelope From address.
				// - Delete or change the Blocked Senders policy.
				[2]string{"550", "rejected by header-based blocked senders - block policy for header from"},
				[2]string{"550", "envelope rejected - block policy for envelope from address"},
			},
			"securityerror": [][2]string{
				// - Messages submitted to SMTP port 587 require authentication. This error indicates
				//   the authentication details provided were incorrect.
				// - Check your authentication details match an internal email address in Mimecast, with
				//   a corresponding Mimecast cloud password. Alternatively, consider sending the mes-
				//   sage on SMTP port 25.
				[2]string{"535", "incorrect authentication data"},
				[2]string{"550", "submitter failed to disabled"},
				[2]string{"550", "submitter failed to authenticate"},
			},
			"systemerror": [][2]string{
				// - The Mimecast server is under maximum load.
				// - No action is required from the end-user. The message will retry 30 times and when
				//   server resources are available, the message is processed.
				[2]string{"451", "unable to process connection at this time"},

				// - The message was incorrectly terminated. This can be caused by:
				//   - Files that previously contained a virus, but haven't been cleaned by an anti-virus
				//     product, leaving traces in the message.
				//   - Firewall issues on the sender's side.
				//   - Incorrectly configured content rules on a security device.
				// - Investigate the Intrusion Detection software or other SMTP protocol analyzers. If
				//   running a Cisco Firewall, ensure the Mail- guard or SMTP Fixup module is disabled.
				[2]string{"451", "message ended early"},

				// - Generic error if the reason is unknown
				// - Contact Mimecast Support.
				[2]string{"451", "unable to process command"},

				// - Generic error if the reason is unknown
				// - Contact Mimecast Support.
				[2]string{"451", "unable to process an email at this time"},
			},
			"toomanyconn": [][2]string{
				// - There are too many concurrent inbound connections for the account. The default is 20.
				// - The IP address is automatically removed from the block list after five minutes.
				//   Continued invalid connections result in the IP being readded to the block list. En-
				//   sure you don't route outbound or journal messages to Mimecast from an IP address
				//   that hasn't been authorized to do so.
				[2]string{"451", "account service is temporarily unavailable"},

				// - The sending server issues more than 100 RCPT TO entries. By default, Mimecast only
				//   accepts 100 RCPT TO entries per message body (DATA). The error triggers the sending
				//   mail server to provide the DATA for the first 100 recipients before it provides the
				//   next batch of RCPT TO entries.
				// - Most mail servers respect the transient error and treat it as a "truncation request".
				//   If your mail server, firewall, or on-site solution doesn't respect the error, you
				//   must ensure that no more than 100 recipients are submitted.
				//   Note:
				//       Solutions like SMTP Fix-Up / MailGuard and ESMTP inspection on Cisco Pix and
				//       ASA Firewalls are known not to respect the transient error. We advise you to
				//       disable this functionality.
				[2]string{"452", "too many recipients"},

				// - There are too many concurrent outbound connections for the account.
				// - Send the messages in smaller chunks to recipients.
				[2]string{"550", "exceeding outbound thread limit"},
			},
			"userunknown": [][2]string{
				// - The email address isn't a valid SMTP address.
				// - The sender must resend the message to a valid internal email address.
				[2]string{"501", "invalid address"},

				// - The server has encountered a bad sequence of commands, or it requires an authenti-
				//   cation.
				// - In case of a "bad sequence", the server has pulled off its commands in the wrong
				//   order, usually because of a broken connection. If authentication is needed, enter
				//   your username and password.
				[2]string{"503", "user unknown"},

				// - Known recipient, LDAP, or SMTP call forwarding recipient validation checks haven't
				//   returned a valid internal user.
				// - The sender must resend the message to a valid internal recipient address.
				[2]string{"550", "invalid recipient"},
			},
			"virusdetected": [][2]string{
				// - A signature was detected that could either be a virus, or a spam score over the
				//   maximum threshold. The spam score isn't available in the Administration Console. If
				//   you aren't a Mimecast customer but have emails rejected with this error code, con-
				//   tact the recipient to adjust their configuration and permit your address. If unsuc-
				//   cessful, your IT department can submit a request to review these email rejections
				//   via our Sender Feedback form.
				// - Anti-virus checks cannot be bypassed. Contact the sender to see if they can stop
				//   these messages from being blocked. Anti-spam checks can be bypassed using a Per-
				//   mitted Senders or Auto Allow policy. Rejected emails can be viewed in your Outbound
				//   Activity and searching for the required email address.
				[2]string{"554", "email rejected due to security policies"},
			},
		}

		issuedcode := strings.ToLower(fo.DiagnosticCode); for e := range messagesof {
			// The key name is a bounce reason name
			for _, f := range messagesof[e] {
				// Try to find each SMTP reply code, error message
				if fo.ReplyCode              != f[0]  { continue }
				if strings.Contains(issuedcode, f[1]) { return e }
			}
		}
		return ""
	}
}

