// Copyright (C) 2020-2021,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reply

//                _           __              _       
//  ___ _ __ ___ | |_ _ __   / / __ ___ _ __ | |_   _ 
// / __| '_ ` _ \| __| '_ \ / / '__/ _ \ '_ \| | | | |
// \__ \ | | | | | |_| |_) / /| | |  __/ |_) | | |_| |
// |___/_| |_| |_|\__| .__/_/ |_|  \___| .__/|_|\__, |
//                   |_|               |_|      |___/ 
/* http://www.ietf.org/rfc/rfc5321.txt
//-------------------------------------------------------------------------------------------------
  4.2.1.  Reply Code Severities and Theory
       2yz  Positive Completion reply
       3yz  Positive Intermediate reply
       4yz  Transient Negative Completion reply
       5yz  Permanent Negative Completion reply

       x0z  Syntax: These replies refer to syntax errors, syntactically correct commands that do
            not fit any functional category, and unimplemented or superfluous commands.
       x1z  Information: These are replies to requests for information, such as status or help.
       x2z  Connections: These are replies referring to the transmission channel.
       x3z  Unspecified.
       x4z  Unspecified.
       x5z  Mail system: These replies indicate the status of the receiver mail system vis-a-vis
            the requested transfer or other mail system action.
*/

// 211  System status, or system help reply
// 214  Help message (Information on how to use the receiver or the meaning of a particular
//      non-standard command; this reply is useful only to the human user)
// 220  <domain> Service ready
// 221  <domain> Service closing transmission channel
// 250  Requested mail action okay, completed
// 251  User not local; will forward to <forward-path> (See Section 3.4)
// 252  Cannot VRFY user, but will accept message and attempt delivery (See Section 3.5.3)
// 253  OK, <n> pending messages for node <domain> started (See RFC1985)
// 354  Start mail input; end with <CRLF>.<CRLF>
// 421   <domain> Service not available, closing transmission channel (This may be a reply to
//       any command if the service knows it must shut down)
// 422   (See RFC5248)
// 430   (See RFC5248)
// 432   A password transition is needed (See RFC4954)
// 450   Requested mail action not taken: mailbox unavailable (e.g., mailbox busy or temporarily
//       blocked for policy reasons)
// 451   Requested action aborted: local error in processing
// 452   Requested action not taken: insufficient system storage
// 453   You have no mail (See RFC2645)
// 454   Temporary authentication failure (See RFC4954)
// 455   Server unable to accommodate parameters
// 456   please retry immediately the message over IPv4 because it fails SPF and DKIM (See
//       https://datatracker.ietf.org/doc/html/draft-martin-smtp-ipv6-to-ipv4-fallback-00
// 458   Unable to queue messages for node <domain> (See RFC1985)
// 459   Node <domain> not allowed: <reason> (See RFC51985)
// 500   Syntax error, command unrecognized (This may include errors such as command line too long)
// 501   Syntax error in parameters or arguments
// 502   Command not implemented (see Section 4.2.4)
// 503   Bad sequence of commands
// 504   Command parameter not implemented
// 520   Please use the correct QHLO ID (See https://datatracker.ietf.org/doc/id/draft-fanf-smtp-quickstart-01.txt)
// 521   Host does not accept mail (See RFC7504)
// 523   Encryption Needed (See RFC5248)
// 524   (See RFC5248)
// 525   User Account Disabled (See RFC5248)
// 530   Authentication required (See RFC4954)
// 533   (See RFC5248)
// 534   Authentication mechanism is too weak (See RFC4954)
// 535   Authentication credentials invalid (See RFC4954)
// 538   Encryption required for requested authentication mechanism (See RFC4954)
// 550   Requested action not taken: mailbox unavailable (e.g., mailbox not found, no access, or
//       command rejected for policy reasons)
// 551   User not local; please try <forward-path> (See Section 3.4)
// 552   Requested mail action aborted: exceeded storage allocation
// 553   Requested action not taken: mailbox name not allowed (e.g., mailbox syntax incorrect)
// 554   Transaction failed (Or, in the case of a connection-opening response, "No SMTP service here")
// 555   MAIL FROM/RCPT TO parameters not recognized or not implemented
// 556   Domain does not accept mail (See RFC7504)
// 557   draft-moore-email-addrquery-01
//
import "strconv"
import "strings"
import sisimoji "sisimai/string"

var ReplyCode2 = []string{"211", "214", "220", "221", "235", "250", "251", "252", "253", "354"}
var ReplyCode4 = []string{"421", "450", "451", "452", "422", "430", "432", "453", "454", "455", "456", "458", "459"}
var ReplyCode5 = []string{
	"550", "552", "553", "551", "521", "525", "502", "520", "523", "524", "530", "533", "534", "535", "538",
	"551", "555", "556", "554", "557", "500", "501", "502", "503", "504",
}
var CodeOfSMTP = map[string][]string{"2": ReplyCode2, "4": ReplyCode4, "5": ReplyCode5}

// Test() checks whether a reply code is a valid code or not
func Test(argv0 string) bool {
	// @param    string argv1  Reply Code(DSN)
	// @return   Bool          true = Invalid reply code, false = Valid reply code
	if len(argv0) < 3 { return false }

	reply, nyaan := strconv.Atoi(argv0)
	if nyaan != nil { return false } // Failed to convert from a string to an integer
	if reply <  211 { return false } // The minimum SMTP Reply code is 211
	if reply >  557 { return false } // The maximum SMTP Reply code is 557

	first := reply % 100
	if first >  59 { return false }  // For example, 499 is not an SMTP Reply code

	if first == 2 {
		// 2yz
		if reply == 235                { return true  } // 235 is a valid code for AUTH (RFC4954)
		if reply  > 253                { return false } // The maximum code of 2xy is 253 (RFC5248)
		if reply  > 221 && reply < 250 { return false } // There is no reply code between 221 and 250
		return true
	}

	if first == 3 {
		// 3yz
		if reply != 354 { return false }
		return true
	}
	return true
}

// Find() returns an SMTP reply code found from the given string
func Find(argv1 string, argv2 string) string {
	// @param    string argv1  String including SMTP reply code like 550
	// @param    string argv2  Status code like 5.1.1 or 2 or 4 or 5
	// @return   string        SMTP reply code or empty if the first argument did not include SMTP Reply Code value
	if len(argv1) < 3                                     { return "" }
	if strings.Contains(strings.ToUpper(argv1), "X-UNIX") { return "" }
	if len(argv2) == 0 { argv2 = "0" }

	statuscode := argv2[0:1]
	replycodes := []string{}

	if statuscode == "2" || statuscode == "4" || statuscode == "5" {
		// The first character of the 2nd argument is 2 or 4 or 5
		replycodes = CodeOfSMTP[statuscode]

	} else {
		// The first character of the 2nd argument is 0 or other values
		// TODO: use "slices" package and slices.Concat() avaialble from Go 1.22
		//       https://pkg.go.dev/slices@master
		replycodes = append(replycodes, CodeOfSMTP["5"]...)
		replycodes = append(replycodes, CodeOfSMTP["4"]...)
		replycodes = append(replycodes, CodeOfSMTP["2"]...)
	}

	esmtperror := " " + argv1 + " "
	esmtpreply := ""
	for _, e := range replycodes {
		// Try to find an SMTP Reply Code from the given string
		appearance := strings.Count(esmtperror, e); if appearance == 0 { continue }
		startingat := 1

		for j := 0; j < appearance; j++ {
			// Find all the reply codes in the error message
			replyindex := sisimoji.IndexOnTheWay(esmtperror, e, startingat); if replyindex < 0 { break }
			formerchar := []byte(esmtperror[replyindex - 1:replyindex])[0]
			latterchar := []byte(esmtperror[replyindex + 3:replyindex + 4])[0]

			if formerchar > 45 && formerchar < 58 { startingat += replyindex + 3; continue } // '.' => '9'
			if latterchar > 45 && latterchar < 58 { startingat += replyindex + 3; continue } // '.' => '9'
			esmtpreply = e
			break
		}
		if esmtpreply != "" { break }
	}
	return esmtpreply
}

