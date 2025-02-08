// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____         _      __                                ___            _           _          _ ___       
// |_   _|__  ___| |_   / / __ ___  __ _ ___  ___  _ __   |_ _|_ __   ___| |_   _  __| | ___  __| |_ _|_ __  
//   | |/ _ \/ __| __| / / '__/ _ \/ _` / __|/ _ \| '_ \   | || '_ \ / __| | | | |/ _` |/ _ \/ _` || || '_ \ 
//   | |  __/\__ \ |_ / /| | |  __/ (_| \__ \ (_) | | | |_ | || | | | (__| | |_| | (_| |  __/ (_| || || | | |
//   |_|\___||___/\__/_/ |_|  \___|\__,_|___/\___/|_| |_(_)___|_| |_|\___|_|\__,_|\__,_|\___|\__,_|___|_| |_|
import "testing"
import "strings"

func TestIncludedIn(t *testing.T) {
	fn := "sisimai/reason.IncludedIn"
	ae := map[string][]string{
		"AuthFailure":    []string{
			`550 5.1.0 192.0.2.222 is not allowed to send from <example.net> per it's SPF Record`,
			`Unauthenticated email from libsisimai.org is not accepted due to domain's DMARC policy`,
		},
		"BadReputation":  []string{
			"451 4.7.650 The mail server [192.0.2.2] has been temporarily rate limited due to IP reputation.",
			"550 Connections from mx.example.jp (192.0.2.2) are being rejected due to a poor email reputation score.",
			"421 4.7.0 [TSS04] Messages from 192.0.2.25 temporarily deferred due to unexpected volume or user complaints",
		},
		"Blocked":        []string{
			"550 Access from ip address 192.0.2.1 blocked.",
			"Remote host said: 554 INVALID IP FOR SENDING MAIL OF DOMAIN amazonses.com [RCPT_TO]",
			"551 Server access forbidden by your IP 192.0.2.2 websites spamcop.net, mailspike.net for removal",
			"blocked using dnsbl.sorbs.net Please see http://support.mailhostbox.com/email-administrators-guide-error-codes/",
		},
		"ContentError":   []string{
			"550 5.6.0 the headers in this message contain improperly-formatted binary content",
			"554 Transaction failed: Duplicate header 'DKIM-Signature'. (in reply to end of DATA command)",
		},
		"ExceedLimit":    []string{
			"5.2.3 Message too large",
			"permanent failure 5.3.0 - Other mail system problem #5.3.4 message header size exceeds limit",
		},
		"Expired":        []string{
			"421 4.4.7 Delivery time expired",
			"Delivery to the following recipient has been delayed: Message will be retried for 2 more day(s)",
		},
		"FailedSTARTTLS": []string{"538 5.7.10 STARTTLS is required to send mail"},
		"Filtered":       []string{
			"550 5.1.2 User reject",
			"You have been blocked by the recipient",
		},
		"HasMoved":       []string{"550 5.1.6 address neko@cat.cat has been replaced by neko@example.jp"},
		"HostUnknown":    []string{
			"550 5.2.1 Host Unknown",
			"kijitora@neko.example.jp: Domain does not exist",
			"554 The mail could not be delivered to the recipient because the domain is not reachable.",
			"550 5.1.7 No such domain neko.example.com",
		},
		"MailboxFull":    []string{
			"450 4.2.2 Mailbox full",
			"452 Insufficient disk space; try again later",
			"5.2.2 <pseudo-local-part-of-apple-icloud-mail@icloud.com>: user is over quota (in reply to RCPT TO)",
			"550 5.2.2 <kijitora@example.co.jp>... Mailbox Full",
			"save to /mail/spool/q002.kijitora.gol.com/gol.com/22r/cat/n2.nyaan mailbox is full: retry timeout exceeded",
		},
		"MailerError":    []string{
			"X-Unix; 255",
			`554 "|IFS=' ' && exec /usr/local/bin/procmail -f- || exit 75 #kijitora"... Service unavailable`,
			"pipe to |/usr/local/neko/bin/cat kijitora@example.com /home/neko/.cat",
		},
		"MesgTooBig":     []string{
			"400 4.2.3 Message too big",
			"#550 5.2.3 RESOLVER.RST.RecipSizeLimit; message too large for this recipient ##",
			"552 5.2.3 Message size exceeds fixed maximum message size (10485760)",
		},
		"NetworkError":   []string{
			"554 5.4.6 Too many hops",
			"554 5.4.6 Hop count exceeded - possible mail loop",
			"neko.example.com[192.0.2.2]:25: No route to host",
			"Error transferring to neko22.example.org; Maximum hop count exceeded. Message probably in a routing loop.",
		},
		"NoRelaying":     []string{
			"550 5.0.0 Relaying Denied",
			"550 relay not permitted",
			"550 5.7.1 Unable to relay for neko@example.com",
		},
		"NotAccept":      []string{
			"556 SMTP protocol returned a permanent error",
			"550 5.1.2 <nekochan@libsisimai.org>... Host unknown (Name server: .: host not found)",
		},
		"NotCompliantRFC":[]string{
			"550 5.7.1 This message is not RFC 5322 compliant. There are multiple Subject headers.",
			"There are multiple Subject headers. Please visit https://support.google.com/mail/?p=RfcMessageNonCompliant",
		},
	//	"OnHold":         []string{"5.0.901 error"},
		"PolicyViolation":[]string{
			"570 5.7.7 Email not accepted for policy reasons",
			"550 Denied by policy",
			"554 email rejected due to security policies - MCSpamSignature.sa.2.2 (in reply to end of DATA command)",
		},
		"Rejected":       []string{
			"550 5.1.8 Domain of sender address example.org does not exist",
			"5.7.1 Access denied (in reply to MAIL FROM command)",
		},
		"RequirePTR":     []string{
			"550 5.7.25 [192.0.2.25] The IP address sending this message does not have a PTR record setup",
			"571 No PTR Record found. Reverse DNS required:",
			"550 5.7.1 Connections not accepted from servers without a valid sender domain. Fix reverse DNS for 203.0.113.2",
		},
		"SecurityError":  []string{
			"570 5.7.0 Authentication failure",
			"#550 5.7.1 RESOLVER.RST.AuthRequired; authentication required ##rfc822;neko-nyaan@cat.example.jp",
		},
		"SpamDetected":   []string{
			"570 5.7.7 Spam Detected",
			"554 5.7.1 Mail Score (59) over MessageScoringUpperLimit (50) - send error reports to postmaster@example.net",
		},
		"Speeding":       []string{"451 4.7.1 <smtp.example.jp[192.0.2.3]>: Client host rejected: Please try again slower"},
	//	"Suppressed":     []string{"There is no sample email which is returned due to being listed in the suppression list"},
		"Suspend":        []string{
			"550 5.0.0 Recipient suspend the service",
			"550 The domain meangel.net is currently suspended. Try later.",
			"550 5.7.1 <kijitora@example.com>: Recipient address rejected: User kijitora@example.com temporary locked. Please try again later!",
		},
		"SystemError":    []string{
			"500 5.3.5 System config error",
			"554 5.3.5 Local configuration error",
			"X-Postfix; mail for example.jp loops back to myself",
		},
		"SystemFull":     []string{"550 5.0.0 Mail system full"},
		"TooManyConn":    []string{
			"421 Too many connections",
			"452 4.3.2 Connection rate limit exceeded. (in reply to MAIL FROM command)",
		},
		"UserUnknown":    []string{
			"550 5.1.1 Unknown User",
			"550 kijitora@example.com... No such user",
			`5.1.0 - Unknown address error 550-'No Such User Here"' (delivery attempts: 0)`,
			": 550 5.1.1 <kijitora@example.jp>: Recipient address rejected: User unknown in local recipient table",
			"554 delivery error: dd This user doesn't have a yahoo.com account (this-local-part-does-not-exist@yahoo.com)",
			`procmail: Couldn't create \"/var/spool/mail/neko\" id: r.example.org: No such user`,
		},
		"Vacation":       []string{
			"I am away on vacation until December 20th and will return email at that time",
			"Hello, I am away until 05/07/2010 and am unable to read your message.",
			"I am out of the office until 02/02/2018",
			"I will be traveling for work on July 16-30.",
		},
		"VirusDetected":  []string{
			"550 5.7.9 The message was rejected because it contains prohibited virus or spam content",
		},
	}
	cx := 0

	for cr := range ae {
		for _, re := range ae[cr] {
			cx++; if IncludedIn[cr]("")                  == true  { t.Errorf("%s[%s]('') returns true",  fn, cr    ) }
			cx++; if IncludedIn[cr](strings.ToLower(re)) == false { t.Errorf("%s[%s](%s) returns false", fn, cr, re) }
		}
	}
	for _, cr := range []string{"Suppressed", "SyntaxError", "Feedback", "Delivered", "OnHold", "Undefined"} {
		cx++; if IncludedIn[cr]("")     == true  { t.Errorf("%s[%s]('') returns true", fn, cr) }
		cx++; if IncludedIn[cr]("neko") == true  { t.Errorf("%s[%s](neko) returns true", fn, cr) }
	}

	t.Logf("The number of tests = %d", cx)
}

