// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//       _               _      ______                  _                        
//  _ __| |__   ___  ___| |_   / / ___| _ __   ___  ___| |_ _ __ _   _ _ __ ___  
// | '__| '_ \ / _ \/ __| __| / /\___ \| '_ \ / _ \/ __| __| '__| | | | '_ ` _ \ 
// | |  | | | | (_) \__ \ |_ / /  ___) | |_) |  __/ (__| |_| |  | |_| | | | | | |
// |_|  |_| |_|\___/|___/\__/_/  |____/| .__/ \___|\___|\__|_|   \__,_|_| |_| |_|
//                                     |_|                                       

package rhost
import "strings"
import "strconv"
import "libsisimai.org/sisimai/sis"
import sisimoji "libsisimai.org/sisimai/string"

func init() {
	// Detect the reason of the bounce returned by this email service
	ReturnedBy["Spectrum"] = func(fo *sis.Fact) string {
		// @param    *sis.Fact fo    Struct to be detected the reason
		// @return   string          Detected bounce reason name
		if fo == nil || fo.DiagnosticCode == "" { return "" }

		errorcodes := [][3]string{
			// https://www.spectrumbusiness.net/support/internet/understanding-email-error-codes
			//   Error codes are placed in one of two categories: incoming or outgoing.
			//   1. If you're trying to send an email to a Charter email address from
			//      a non-Charter email address (such as Gmail, Yahoo, Hotmail, etc.),
			//      you may receive an error that begins with AUP#I, followed by four numbers.
			//
			//   2. If you are trying to send an email from a Charter email address
			//      to an outgoing recipient, you may get an error code beginning with
			//      AUP#O, also followed by four numbers.
			//
			// 1000 Your IP address has been blocked due to suspicious activity. If you're a Spectrum
			//      customer using a Spectrum-issued IP, contact us. If you're using an IP address other
			//      than one provided by Spectrum, blocks will remain in place until they expire.
			[3]string{"1000", "", "blocked"},

			// 1010 This email account has been blocked from sending emails due to suspicious activity.
			//      Blocks will expire based on the nature of the activity. If you're a Spectrum customer,
			//      change all of your Spectrum passwords to secure your account and then contact us.
			[3]string{"1010", "", "rejected"},

			// 1020 This email account has limited access to send emails based on suspicious activity.
			// 1080 Blocks will expire based on the nature of the activity.
			//      If you're a Spectrum customer, contact us to remove the block.
			[3]string{"1020", "1080", "rejected"},

			// 1090 The email you're trying to send can't be processed. Try sending again at a later time.
			[3]string{"1090", "", "systemerror"},

			// 1100 The IP address you're trying to connect from has an issue with the Domain Name System.
			// 1150 Spectrum requires a full circle DNS for emails to be allowed through. Verify the IP
			//      you're connecting from, and check the IP address to ensure a reverse DNS entry exists
			//      for the IP. If the IP address is a Spectrum-provided email address, contact us.
			[3]string{"1100", "1150", "requireptr"},

			// 1160 The email you tried to send goes against your domain's security policies. 
			// 1190 Please contact the email administrators of your domain.
			[3]string{"1160", "1190", "policyviolation"},

			// 1200 The IP address you're trying to send from has been flagged by Cloudmark CSI as
			// 1210 potential spam. Have your IP administrator request a reset. 
			//      Note: Cloudmark has sole discretion whether to remove the sending IP address from
			//            their lists.
			[3]string{"1200", "1210", "blocked"},

			// 1220 Your IP address has been blacklisted by Spamhaus. The owner of the IP address must
			// 1250 contact Spamhaus to be removed from the list.
			//      Note: Spamhaus has sole discretion whether to remove the sending IP address from
			//            their lists.
			[3]string{"1220", "1250", "blokced"},

			// 1260 Spectrum doesn't process IPV6 addresses. Connect with an IPv4 address and try again.
			[3]string{"1260", "", "networkerror"},

			// 1300 Spectrum limits the number of concurrent connections from a sender, as well as the
			// 1340 total number of connections allowed. Limits vary based on the reputation of the IP
			//      address. Reduce your number of connections and try again later.
			[3]string{"1300", "1340", "toomanyconn"},

			// 1350 Spectrum limits emails by the number of messages sent, amount of recipients,
			// 1490 potential for spam and invalid recipients.
			[3]string{"1350", "1490", "speeding"},

			// 1500 Your email was rejected for attempting to send as a different email address than you
			//      signed in under. Check that you're sending emails from the address you signed in with.
			[3]string{"1500", "", "rejected"},

			// 1520 Your email was rejected for attempting to send as a different email address than a
			//      domain that we host. Check the outgoing email address and try again.
			[3]string{"1520", "", "rejected"},

			// 1530 Your email was rejected because it's larger than the maximum size of 20MB.
			[3]string{"1530", "", "mesgtoobig"},

			// 1540 Your emails were deferred for attempting to send too many in a single session.
			//      Reconnect and try reducing the number of emails you send at one time.
			[3]string{"1540", "", "speeding"},

			// 1550 Your email was rejected for having too many recipients in one message. Reduce the
			//      number of recipients and try again later.
			[3]string{"1550", "", "speeding"},

			// 1560 Your email was rejected for having too many invalid recipients. Check your outgoing
			//      email addresses and try again later.
			[3]string{"1560", "", "policyviolation"},

			// 1580 You've tried to send messages to too many recipients in a short period of time.
			//      Wait a little while and try again later.
			[3]string{"1580", "", "speeding"},
		}

		issuedcode := fo.DiagnosticCode
		labelindex := strings.Index(issuedcode, "AUP#"); if labelindex < 0 { return "" }
		codestring := ""

		for _, e := range issuedcode[labelindex + 4:] {
			// Try to get the four digit error code number from the error message
			if len(codestring) == 4 { break }

			// The ASCII code of the character is less than '0' or is greater than '9'
			if e < 48 || e > 57 { codestring = ""; continue }
			codestring += string(e)
		}
		if len(codestring) != 4                              { return "" }
		if sisimoji.ContainsOnlyNumbers(codestring) == false { return "" }

		codenumber, nyaan := strconv.ParseUint(codestring, 10, 16); if nyaan != nil { return "" }
		for _, e := range errorcodes {
			// Try to find an error code matches with the code in the value of fo.DiagnosticCode
			if codestring == e[0] {
				// ["1500", "", "reason"] or ["1500", "1550", "reason"]
				return e[2]

			} else {
				// Check the code number is inlcuded the range like ["1500", "1550", 'reason']
				if e[1] == "" { continue }

				coderange0, nyaan := strconv.ParseUint(e[0], 10, 16); if nyaan != nil { continue }
				coderange1, nyaan := strconv.ParseUint(e[1], 10, 16); if nyaan != nil { continue }
				if codenumber < coderange0 || codenumber > coderange1                 { continue }
				return e[2]
			}
		}
		return ""
	}
}

