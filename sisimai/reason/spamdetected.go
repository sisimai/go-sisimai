// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  ____                        ____       _            _           _ 
// / ___| _ __   __ _ _ __ ___ |  _ \  ___| |_ ___  ___| |_ ___  __| |
// \___ \| '_ \ / _` | '_ ` _ \| | | |/ _ \ __/ _ \/ __| __/ _ \/ _` |
//  ___) | |_) | (_| | | | | | | |_| |  __/ ||  __/ (__| ||  __/ (_| |
// |____/| .__/ \__,_|_| |_| |_|____/ \___|\__\___|\___|\__\___|\__,_|
//       |_|                                                          
import "strings"
import "sisimai/sis"
import "sisimai/smtp/status"
import sisimoji "sisimai/string"

func init() {
	// Try to check the argument string includes any of the strings in the error message pattern
	IncludedIn["SpamDetected"] = func(argv1 string) bool {
		// @param    string argv1 Does the string include any of the strings listed in the pattern?
		// @return   bool         true: Included, false: did not include
		index := []string{
			" - spam",
			"//www.spamhaus.org/help/help_spam_16.htm",
			"//dsbl.org/help/help_spam_16.htm",
			"//mail.163.com/help/help_spam_16.htm",
			"554 5.7.0 reject, id=",
			"appears to be unsolicited",
			"blacklisted url in message",
			"block for spam",
			"blocked by policy: no spam please",
			"blocked by spamassassin",                      // rejected by SpamAssassin
			"blocked for abuse. see http://att.net/blocks", // AT&T
			"cannot be forwarded because it was detected as spam",
			"considered unsolicited bulk e-mail (spam) by our mail filters",
			"content filter rejection",
			"cyberoam anti spam engine has identified this email as a bulk email",
			"denied due to spam list",
			"high probability of spam",
			"is classified as spam and is rejected",
			"listed in work.drbl.imedia.ru",
			"the mail server detected your message as spam and has prevented delivery.", // CPanel/Exim with SA rejections on
			"mail appears to be unsolicited", // rejected due to spam
			"mail content denied",            // http://service.mail.qq.com/cgi-bin/help?subtype=1&&id=20022&&no=1000726
			"may consider spam",
			"message considered as spam or virus",
			"message contains spam or virus",
			"message content rejected",
			"message detected as spam",
			"message filtered",
			"message filtered. please see the faqs section on spam",
			"message filtered. refer to the troubleshooting page at ",
			"message looks like spam",
			"message is being rejected as it seems to be a spam",
			"message refused by mailmarshal spamprofiler",
			"message refused by trustwave seg spamprofiler",
			"message rejected as spam",
			"message rejected because of unacceptable content",
			"message rejected due to suspected spam content",
			"message rejected for policy reasons",
			"message was rejected for possible spam/virus content",
			"our email server thinks this email is spam",
			"our system has detected that this message is ",
			"probable spam",
			"reject bulk.advertising",
			"rejected: spamassassin score ",
			"rejected - bulk email",
			"rejecting banned content",
			"rejecting mail content",
			"related to content with spam-like characteristics",
			"sender domain listed at ",
			"sending address not accepted due to spam filter",
			"spam blocked",
			"spam check",
			"spam content matched",
			"spam detected",
			"spam email",
			"spam email not accepted",
			"spam message rejected.", // mail.ru
			"spam not accepted",
			"spam refused",
			"spam rejection",
			"spam score ",
			"spambouncer identified spam", // SpamBouncer identified SPAM
			"spamming not allowed",
			"too many spam complaints",
			"too much spam.",              // Earthlink
			"the email message was detected as spam",
			"the message has been rejected by spam filtering engine",
			"the message was rejected due to classification as bulk mail",
			"the content of this message looked like spam", // SendGrid
			"this message appears to be spam",
			"this message has been identified as spam",
			"this message has been scored as spam with a probability",
			"this message was classified as spam",
			"this message was rejected by recurrent pattern detection system",
			"transaction failed spam message not queued",   // SendGrid
			"we dont accept spam",
			"your email appears similar to spam we have received before",
			"your email breaches local uribl policy",
			"your email had spam-like ",
			"your email is considered spam",
			"your email is probably spam",
			"your email was detected as spam",
			"your message as spam and has prevented delivery",
			"your message has been temporarily blocked by our filter",
			"your message has been rejected because it appears to be spam",
			"your message has triggered a spam block",
			"your message may contain the spam contents",
			"your message failed several antispam checks",
		}
		pairs := [][]string{
			[]string{"greylisted", " please try again in"},
			[]string{"mail rejete. mail rejected. ", "506"},
			[]string{"our filters rate at and above ", " percent probability of being spam"},
			[]string{"rejected by ", " (spam)"},
			[]string{"rejected due to spam ", "classification"},
			[]string{"rejected due to spam ", "content"},
			[]string{"rule imposed as ", " is blacklisted on"},
			[]string{"spam ", " exceeded"},
			[]string{"this message scored ", " spam points"},
		}

		for _, v := range index { if strings.Contains(argv1, v) { return true }}
		for _, v := range pairs { if sisimoji.Aligned(argv1, v) { return true }}
		return false
	}

	// The bounce reason is "spamdetected" or not
	ProbesInto["SpamDetected"] = func(fo *sis.Fact) bool {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   bool            true: is spamdetected, false: is not spamdetected
		if fo.Reason == "spamdetected"                      { return true }
		if status.Name(fo.DeliveryStatus) == "spamdetected" { return true }
		if fo.Command == "CONN" || fo.Command == "EHLO" || fo.Command == "HELO" ||
		   fo.Command == "MAIL" || fo.Command == "RCPT" { return false }
		return IncludedIn["SpamDetected"](strings.ToLower(fo.DiagnosticCode))
	}
}

