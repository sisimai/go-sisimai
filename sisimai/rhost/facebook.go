// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _      ____                _                 _    
//  _ __| |__   ___  ___| |_   / / _| __ _  ___ ___| |__   ___   ___ | | __
// | '__| '_ \ / _ \/ __| __| / / |_ / _` |/ __/ _ \ '_ \ / _ \ / _ \| |/ /
// | |  | | | | (_) \__ \ |_ / /|  _| (_| | (_|  __/ |_) | (_) | (_) |   < 
// |_|  |_| |_|\___/|___/\__/_/ |_|  \__,_|\___\___|_.__/ \___/ \___/|_|\_\
import "strings"
import "sisimai/sis"
import sisimoji "sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["Facebook"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		// @see      https://www.facebook.com/postmaster/response_codes
		if fo.DiagnosticCode                        == ""    { return "" }
		if strings.Contains(fo.DiagnosticCode, "-") == false { return "" }

		errorcodes := map[string][]string{
			// http://postmaster.facebook.com/response_codes
			// NOT TESTD EXCEPT RCP-P2
			"authfailure": []string{
				"POL-P7",   // The message does not comply with Facebook's Domain Authentication requirements.
			},
			"blocked": []string{
				"POL-P1",   // Your mail server's IP Address is listed on the Spamhaus PBL.
				"POL-P2",   // Facebook will no longer accept mail from your mail server's IP Address.
				"POL-P3",   // Facebook is not accepting messages from your mail server. This will persist for 4 to 8 hours.
				"POL-P4",   // Facebook is not accepting messages from your mail server. This will persist for 24 to 48 hours.
				"POL-T1",   // Facebook is not accepting messages from your mail server, but they may be retried later. This will persist for 1 to 2 hours.
				"POL-T2",   // Facebook is not accepting messages from your mail server, but they may be retried later. This will persist for 4 to 8 hours.
				"POL-T3",   // Facebook is not accepting messages from your mail server, but they may be retried later. This will persist for 24 to 48 hours.
			},
			"contenterror": []string{
				"MSG-P2",   // The message contains an attachment type that Facebook does not accept.
			},
			"filtered": []string{
				"RCP-P2",   // The attempted recipient's preferences prevent messages from being delivered.
				"RCP-P3",   // The attempted recipient's privacy settings blocked the delivery.
			},
			"mesgtoobig": []string{
				"MSG-P1",   // The message exceeds Facebook's maximum allowed size.
				"INT-P2",   // The message exceeds Facebook's maximum allowed size.
			},
			"notcompliantrfc": []string{
				"MSG-P3",   // The message contains multiple instances of a header field that can only be present once.
			},
			"rejected": []string{
				"DNS-P1",   // Your SMTP MAIL FROM domain does not exist.
				"DNS-P2",   // Your SMTP MAIL FROM domain does not have an MX record.
				"DNS-T1",   // Your SMTP MAIL FROM domain exists but does not currently resolve.
			},
			"requireptr": []string{
				"DNS-P3",   // Your mail server does not have a reverse DNS record.
				"DNS-T2",   // You mail server's reverse DNS record does not currently resolve.
			},
			"spamdetected": []string{
				"POL-P6",   // The message contains a url that has been blocked by Facebook.
				"POL-P7",   // The message does not comply with Facebook's abuse policies and will not be accepted.
			},
			"suspend": []string{
				"RCP-T4",   // The attempted recipient address is currently deactivated. The user may or may not reactivate it.
			},
			"systemerror": []string{
				"RCP-T1",   // The attempted recipient address is not currently available due to an internal system issue. This is a temporary condition.
			},
			"toomanyconn": []string{
				"CON-T1",   // Facebook's mail server currently has too many connections open to allow another one.
				"CON-T2",   // Your mail server currently has too many connections open to Facebook's mail servers.
				"CON-T3",   // Your mail server has opened too many new connections to Facebook's mail servers in a short period of time.
				"CON-T4",   // Your mail server has exceeded the maximum number of recipients for its current connection.
				"MSG-T1",   // The number of recipients on the message exceeds Facebook's allowed maximum.
			},
			"userunknown": []string{
				"RCP-P1",   // The attempted recipient address does not exist.
				"INT-P1",   // The attempted recipient address does not exist.
				"INT-P3",   // The attempted recpient group address does not exist.
				"INT-P4",   // The attempted recipient address does not exist.
			},
			"virusdetected": []string{
				"POL-P5",   // The message contains a virus.
			},
		}
		errorindex := strings.Index(fo.DiagnosticCode, "-")
		errorlabel := fo.DiagnosticCode[errorindex - 3:errorindex + 3]
		reasontext := ""

		for e := range errorcodes {
			// The key is a bounce reason name
			if sisimoji.EqualsAny(errorlabel, errorcodes[e]) == false { continue }
			reasontext = e; break
		}
		return reasontext
	}
}

