// Copyright (C) 2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _            _            _ 
// | __ )| | ___   ___| | _____  __| |
// |  _ \| |/ _ \ / __| |/ / _ \/ _` |
// | |_) | | (_) | (__|   <  __/ (_| |
// |____/|_|\___/ \___|_|\_\___|\__,_|

package reason
import "strings"
import "libsisimai.org/sisimai/v5/eb"
import "libsisimai.org/sisimai/v5/siba"
import "libsisimai.org/sisimai/v5/moji"
import "libsisimai.org/sisimai/v5/smtp/status"

func init() {
	// IncludedIn[*] Try to check the argument string includes any of the strings in the error message pattern.
	//   Arguments:
	//     - mesg (string): Does the string include any of the strings listed in the pattern?
	//   Returns:
	//     - (bool): true if the argument includes one or more error message pattern.
	IncludedIn[eb.ReBLOC] = func(mesg string) bool {
		if mesg == "" { return false }

		index := []string{
			"bad sender ip address",
			"banned sending ip", // Office365
			"blacklisted by",
			"dnsbl:attrbl",
			"client host rejected: abus detecte gu_eib_02", // SFR
			"client host rejected: abus detecte gu_eib_04", // SFR
			"client host rejected: may not be mail exchanger",
			"connection refused by",
			"currently sending spam see: ",
			"domain does not exist:",
			"domain isn't in my list of allowed rcpthosts",
			"error: no valid recipients from ",
			"esmtp not accepting connections", // icloud.com
			"extreme bad ip profile",
			"helo command rejected:",
			"host network not allowed",
			"invalid ip for sending mail of domain",
			"is in a black list",
			"is not allowed to send mail from",
			"no access from mail server",
			"no matches to nameserver query",
			"part of their network is on our block list",
			"please use the smtp server of your isp",
			"rejected - multi-blacklist", // junkemailfilter.com
			"rejected because the sending mta or the sender has not passed validation",
			"rejecting open proxy", // Sendmail(srvrsmtp.c)
			"sender ip address rejected",
			"server access forbidden by your ip ",
			"service not available, closing transmission channel",
			"smtp error from remote mail server after initial connection:", // Exim
			"temporarily deferred due to unexpected volume or user complaints",
			"you are not allowed to connect",
			"you are sending spam",
			"your ip address is listed in the rbl",
			"your network is temporary blacklisted",
			"your remotehost looks suspiciously like spammer",
			"your server requires confirmation",
		}
		pairs := [][]string{
			[]string{"(", "@", ":blocked)"},
			[]string{"access from ip address ", " blocked"},
			[]string{"blocked by ", " dnsbl"},
			[]string{"client ", " blocked using"},
			[]string{"connection ", "dropped"},
			[]string{"connections will not be accepted from ", " because the ip is in spamhaus's list"},
			[]string{"dnsbl:rbl ", ">_is_blocked"},
			[]string{"dynamic", " ip"},
			[]string{"email blocked by ", ".barracudacentral.org"},
			[]string{"email blocked by ", "spamhaus"},
			[]string{"from ", " ip address"},
			[]string{"host ", " said: ", "550 blocked"},
			[]string{"host ", " refused to talk to me: ", " blocked"},
			[]string{"ip ", " is blocked by earthlink"}, // Earthlink
			[]string{"is in an ", "rbl on "},
			[]string{"mail server at ", " is blocked"},
			[]string{"mail from "," refused"},
			[]string{"message from ", " rejected based on blacklist"},
			[]string{"messages from ", " temporarily deferred due to user complaints"}, // Yahoo!
			[]string{"server ip ", " listed as abusive"},
			[]string{"sorry! your ip address", " is blocked by rbl"}, // junkemailfilter.com
			[]string{"the ", " is blacklisted"}, // the email, the domain, the ip
			[]string{"veuillez essayer plus tard. service refused, please try later. ", "103"},
			[]string{"veuillez essayer plus tard. service refused, please try later. ", "510"},
			[]string{"your access ip", " has been rejected"},
			[]string{"your sender's ip address is listed at ", ".abuseat.org"},
		}
		return moji.ContainsAny(mesg, index) || moji.AlignedAny(mesg, pairs)
	}

	// ProbesInto[*] checks the bounce reason is the reason defined in this file or not.
	//   Arguments:
	//     - fo (*siba.Fact): Decoded data in progress.
	//   Returns:
	//     - (bool): true if a reason is the reason defined in this file.
	ProbesInto[eb.ReBLOC] = func(fo *siba.Fact) bool {
		if fo == nil                                   { return false }
		if fo.Reason == eb.ReBLOC                      { return true  }
		if status.Name(fo.DeliveryStatus) == eb.ReBLOC { return true  }
		return IncludedIn[eb.ReBLOC](strings.ToLower(fo.DiagnosticCode))
	}
}

