// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _            _            _ 
// | __ )| | ___   ___| | _____  __| |
// |  _ \| |/ _ \ / __| |/ / _ \/ _` |
// | |_) | | (_) | (__|   <  __/ (_| |
// |____/|_|\___/ \___|_|\_\___|\__,_|

package reason
import "strings"
import "libsisimai.org/sisimai/v5/sis"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - argv1 (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn["Blocked"] = func(argv1 string) bool {
		if argv1 == "" { return false }

		index := []string{
			" said: 550 blocked",
			"//www.spamcop.net/bl.",
			"bad sender ip address",
			"banned sending ip", // Office365
			"blacklisted by",
			"blocked using ",
			"blocked - see http",
			"dnsbl:attrbl",
			"client host rejected: abus detecte gu_eib_02", // SFR
			"client host rejected: abus detecte gu_eib_04", // SFR
			"client host rejected: may not be mail exchanger",
			"client host rejected: was not authenticated",  // Microsoft
			"confirm this mail server",
			"connection dropped",
			"connection refused by",
			"connection reset by peer",
			"connection was dropped by remote host",
			"connections not accepted from ip addresses on spamhaus xbl",
			"currently sending spam see: ",
			"domain does not exist:",
			"dynamic/zombied/spam ips blocked",
			"error: no valid recipients from ",
			"esmtp not accepting connections", // icloud.com
			"extreme bad ip profile",
			"go away",
			"helo command rejected:",
			"host network not allowed",
			"hosts with dynamic ip",
			"invalid ip for sending mail of domain",
			"is in a black list",
			"is not allowed to send mail from",
			"no access from mail server",
			"no matches to nameserver query",
			"not currently accepting mail from your ip", // Microsoft
			"part of their network is on our block list",
			"please use the smtp server of your isp",
			"refused - see http",
			"rejected - multi-blacklist", // junkemailfilter.com
			"rejected because the sending mta or the sender has not passed validation",
			"rejecting open proxy", // Sendmail(srvrsmtp.c)
			"sender ip address rejected",
			"server access forbidden by your ip ",
			"service not available, closing transmission channel",
			"smtp error from remote mail server after initial connection:", // Exim
			"sorry, that domain isn't in my list of allowed rcpthosts",
			"sorry, your remotehost looks suspiciously like spammer",
			"temporarily deferred due to unexpected volume or user complaints",
			"to submit messages to this e-mail system has been rejected",
			"too many spams from your ip", // free.fr
			"too many unwanted messages have been sent from the following ip address above",
			"was blocked by ",
			"we do not accept mail from dynamic ips", // @mail.ru
			"you are not allowed to connect",
			"you are sending spam",
			"your ip address is listed in the rbl",
			"your network is temporary blacklisted",
			"your server requires confirmation",
		}
		pairs := [][]string{
			[]string{"(", "@", ":blocked)"},
			[]string{"access from ip address ", " blocked"},
			[]string{"client host ", " blocked using"},
			[]string{"connections will not be accepted from ", " because the ip is in spamhaus's list"},
			[]string{"dnsbl:rbl ", ">_is_blocked"},
			[]string{"email blocked by ", ".barracudacentral.org"},
			[]string{"email blocked by ", "spamhaus"},
			[]string{"host ", " refused to talk to me: ", " blocked"},
			[]string{"ip ", " is blocked by earthlink"}, // Earthlink
			[]string{"is in an ", "rbl on "},
			[]string{"mail server at ", " is blocked"},
			[]string{"mail from "," refused:"},
			[]string{"mail from "," refused, see ", "orbs"},
			[]string{"message from ", " rejected based on blacklist"},
			[]string{"messages from ", " temporarily deferred due to user complaints"}, // Yahoo!
			[]string{"server ip ", " listed as abusive"},
			[]string{"sorry! your ip address", " is blocked by rbl"}, // junkemailfilter.com
			[]string{"the domain ", " is blacklisted"},
			[]string{"the email ", " is blacklisted"},
			[]string{"the ip", " is blacklisted"},
			[]string{"veuillez essayer plus tard. service refused, please try later. ", "103"},
			[]string{"veuillez essayer plus tard. service refused, please try later. ", "510"},
			[]string{"your access ip", " has been rejected"},
			[]string{"your sender's ip address is listed at ", ".abuseat.org"},
		}
		return moji.ContainsAny(argv1, index) || moji.AlignedAny(argv1, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*sis.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto["Blocked"] = func(fo *sis.Fact) bool {
		if fo == nil                                   { return false }
		if fo.Reason == "blocked"                      { return true  }
		if status.Name(fo.DeliveryStatus) == "blocked" { return true  }
		return IncludedIn["Blocked"](strings.ToLower(fo.DiagnosticCode))
	}
}


