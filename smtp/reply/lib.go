// Copyright (C) 2020-2021,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                _           __              _       
//  ___ _ __ ___ | |_ _ __   / / __ ___ _ __ | |_   _ 
// / __| '_ ` _ \| __| '_ \ / / '__/ _ \ '_ \| | | | |
// \__ \ | | | | | |_| |_) / /| | |  __/ |_) | | |_| |
// |___/_| |_| |_|\__| .__/_/ |_|  \___| .__/|_|\__, |
//                   |_|               |_|      |___/ 

// Package "smtp/reply" provides funtions related to SMTP reply codes such as 421, 550.
package reply

// RFC 1870: SMTP Service Extension for Message Size Declaration (SIZE)
// RFC 1985: SMTP Service Extension for Remote Message Queue Starting (ETRN)
// RFC 2645: Authenticated TURN for On-Demand Mail Relay (ATRN)
// RFC 3207: SMTP Service Extension for Secure SMTP over Transport Layer Security (STARTTLS) 1
// RFC 3030: SMTP Service Extension for Command Pipelining (CHUNKING)
// RFC 3461: Delivery Status Notifications (DSN)
// RFC 4954: SMTP Service Extension for Authentication (AUTH)
// RFC 5321: Simple Mail Transfer Protocol
// RFC 5336: SMTP Extension for Internationalized Email (UTF8SMTP)
// RFC 6531: SMTP Extension for Internationalized Email (SMTPUTF8)
// RFC 7504: SMTP 521 and 556 Reply Codes 
// RFC 9422: The LIMITS SMTP Service Extension
//-------------------------------------------------------------------------------------------------
// 4.2.1.  Reply Code Severities and Theory
//   2yz  Positive Completion reply
//   3yz  Positive Intermediate reply
//   4yz  Transient Negative Completion reply
//   5yz  Permanent Negative Completion reply

//   x0z  Syntax: These replies refer to syntax errors, syntactically correct commands that do not
//        fit any functional category, and unimplemented or superfluous commands.
//   x1z  Information: These are replies to requests for information, such as status or help.
//   x2z  Connections: These are replies referring to the transmission channel.
//   x3z  Unspecified.
//   x4z  Unspecified.
//   x5z  Mail system: These replies indicate the status of the receiver mail system vis-a-vis the
//        requested transfer or other mail system action.
//
// 211  System status, or system help reply
// 214  Help message (Information on how to use the receiver or the meaning of a particular
//      non-standard command; this reply is useful only to the human user)
// 220  <domain> Service ready
// 221  <domain> Service closing transmission channel
// 235  This response to the AUTH command indicates that the authentication was successful (RFC4954)
// 250  Requested mail action okay, completed
// 251  User not local; will forward to <forward-path> (See Section 3.4)
// 252  Cannot VRFY user, but will accept message and attempt delivery (See Section 3.5.3)
// 253  OK, <n> pending messages for node <domain> started (See RFC1985)
// 334  A server challenge is sent as a 334 reply with the text part containing the [BASE64] encoded
//      string supplied by the SASL mechanism.  This challenge MUST NOT contain any text other
//      than the BASE64 encoded challenge. (RFC4954)
// 354  Start mail input; end with <CRLF>.<CRLF>
// 421  <domain> Service not available, closing transmission channel (This may be a reply to
//      any command if the service knows it must shut down)
// 422  (See RFC5248)
// 430  (See RFC5248)
// 432  A password transition is needed (See RFC4954)
// 450  Requested mail action not taken: mailbox unavailable (e.g., mailbox busy or temporarily
//      blocked for policy reasons)
// 451  Requested action aborted: local error in processing
// 452  Requested action not taken: insufficient system storage
// 453  You have no mail (See RFC2645)
// 454  Temporary authentication failure (See RFC4954)
// 455  Server unable to accommodate parameters
// 458  Unable to queue messages for node <domain> (See RFC1985)
// 459  Node <domain> not allowed: <reason> (See RFC51985)
// 500  Syntax error, command unrecognized (This may include errors such as command line too long)
// 501  Syntax error in parameters or arguments
// 502  Command not implemented (see Section 4.2.4)
// 503  Bad sequence of commands
// 504  Command parameter not implemented
// 521  Host does not accept mail (See RFC7504)
// 523  Encryption Needed (See RFC5248)
// 524  (See RFC5248)
// 525  User Account Disabled (See RFC5248)
// 530  Authentication required (See RFC4954)
// 533  (See RFC5248)
// 534  Authentication mechanism is too weak (See RFC4954)
// 535  Authentication credentials invalid (See RFC4954)
// 538  Encryption required for requested authentication mechanism (See RFC4954)
// 550  Requested action not taken: mailbox unavailable (e.g., mailbox not found, no access, or
//      command rejected for policy reasons)
// 551  User not local; please try <forward-path> (See Section 3.4)
// 552  Requested mail action aborted: exceeded storage allocation
// 553  Requested action not taken: mailbox name not allowed (e.g., mailbox syntax incorrect)
// 554  Transaction failed (Or, in the case of a connection-opening response, "No SMTP service here")
// 555  MAIL FROM/RCPT TO parameters not recognized or not implemented
// 556  Domain does not accept mail (See RFC7504)
//
import "bytes"
import "slices"
import "strconv"
import "strings"
import "libsisimai.org/sisimai/v5/eb"

var replycode2 = []string{"211", "214", "220", "221", "235", "250", "251", "252", "253", "334", "354"}
var replycode4 = []string{"421", "450", "451", "452", "422", "430", "432", "453", "454", "455", "458", "459"}
var replycode5 = []string{
	"550", "552", "553", "551", "521", "525", "523", "524", "530", "533", "534", "535", "538", "555",
	"556", "554", "500", "501", "502", "503", "504",
}
var codeofsmtp = map[string][]string{"2": replycode2, "4": replycode4, "5": replycode5}
var associated = map[string][]string{
	"422": []string{eb.CeAUTH,  "4.7.12",  eb.ReSAFE}, // RFC5238
	"432": []string{eb.CeAUTH,  "4.7.12",  eb.ReSAFE}, // RFC4954, RFC5321
	"451": []string{"",         "",        eb.RePROC}, // RFC2465, RFC5321
	"452": []string{"",         "",        eb.ReDISK}, // RFC5321
	"454": []string{eb.CeAUTH,  "4.7.0",   eb.ReSAFE}, // RFC3207, RFC4954
	"455": []string{"",         "",        eb.ReCOMM}, // RFC5321
	"500": []string{"",         "",        eb.ReCOMM}, // RFC5321
	"501": []string{"",         "",        eb.ReCOMM}, // RFC5321
	"502": []string{"",         "",        eb.ReCOMM}, // RFC5321
	"503": []string{"",         "",        eb.ReCOMM}, // RFC5321
	"504": []string{"",         "",        eb.ReCOMM}, // RFC5321
	"521": []string{eb.CeCONN,  "",        eb.Re00MX}, // RFC7504
	"523": []string{eb.CeAUTH,  "5.7.10",  eb.ReSAFE}, // RFC5248
	"524": []string{eb.CeAUTH,  "5.7.11",  eb.ReSAFE}, // RFC5248
	"525": []string{eb.CeAUTH,  "5.7.13",  eb.ReSAFE}, // RFC5248
	"534": []string{eb.CeAUTH,  "5.7.9",   eb.ReSAFE}, // RFC4954, RFC5248
	"535": []string{eb.CeAUTH,  "5.7.8",   eb.ReSAFE}, // RFC4954, RFC5248
	"538": []string{eb.CeAUTH,  "5.7.11",  eb.ReSAFE}, // RFC4954, RFC5248
	"551": []string{"",         "",        eb.ReMOVE}, // RFC5321, RFC5336, RFC6531
	"552": []string{"",         "",        eb.ReFULL}, // RFC5321
	"555": []string{"",         "",        eb.ReCOMM}, // RFC5321
	"556": []string{eb.CeRCPT,  "",        eb.Re00MX}, // RFC7504
}

// AssociatedWith returns a slice associated with the SMTP reply code of the argument
//   Arguments:
//     - reply (string): SMTP reply code
//   Returns:
//     - ([]string):     ["SMTP Command", "DSN", "Reason"]
//   Since:
//     - 5.2.2
func AssociatedWith(reply string) []string { return associated[reply] }

// Test checks whether a reply code is a valid code or not.
//   Arguments:
//     - code (string): SMTP reply code.
//   Returns:
//     - (bool): true if the argument is a valid SMTP reply code, false otherwise.
func Test(code string) bool {
	if len(code) < 3 { return false }

	reply, nyaan := strconv.Atoi(code)
	if nyaan != nil     { return false } // Failed to convert from a string to an integer
	if reply <  211     { return false } // The minimum SMTP Reply code is 211
	if reply >  556     { return false } // The maximum SMTP Reply code is 556 (RFC7504)
	if reply % 100 > 59 { return false } // For example, 499 is not an SMTP Reply code

	if first := reply / 100; first == 2 {
		// 2yz
		if reply == 235                { return true  } // 235 is a valid code for AUTH (RFC4954)
		if reply  > 253                { return false } // The maximum code of 2xy is 253 (RFC5248)
		if reply  > 221 && reply < 250 { return false } // There is no reply code between 221 and 250
		return true

	} else {
		// 3yz is 334 or 354 only
		if first == 3 && reply != 334 && reply != 354 { return false }
	}
	return true
}

// Find returns an SMTP reply code found from the given string.
//   Arguments:
//     - logs (string): String including SMTP reply code like 550.
//     - hint (string): SMTP status code like "5.1.1", or the 1st digit of the code like "2", "4", or "5".
//   Returns:
//     - (string): SMTP reply code found in the 1st argument.
func Find(logs, hint string) string {
	if len(logs) < 3 || strings.Contains(strings.ToUpper(logs), "X-UNIX") { return "" }
	if len(hint) == 0 { hint = "0" }

	esmtperror := []byte(" " + logs + " "); cw := len(esmtperror)
	replycodes := make([]string, 0, 50)
	if cr := hint[0:1]; cr == "2" || cr == "4" || cr == "5" {
		// The first character of the 2nd argument is 2 or 4 or 5
		replycodes = codeofsmtp[cr]

	} else {
		// The first character of the 2nd argument is 0 or other values
		replycodes = slices.Concat(replycodes, codeofsmtp["5"], codeofsmtp["4"], codeofsmtp["2"])
	}

	for _, e := range replycodes {
		// Try to find an SMTP Reply Code from the given string
		characterb := []byte(e)
		appearance := bytes.Count(esmtperror, characterb); if appearance == 0 { continue }
		startingat := 1

		for j := 0; j < appearance; j++ {
			// Find all the reply codes in the error message
			if cw <= startingat { break }

			replyindex := bytes.Index(esmtperror[startingat:], characterb); if replyindex < 0 { break }
			replyindex += startingat
			formerchar := esmtperror[replyindex - 1:replyindex][0]
			latterchar := esmtperror[replyindex + 3:replyindex + 4][0]

			if formerchar > 45 && formerchar < 58 { startingat += replyindex + 3; continue } // '.' => '9'
			if latterchar > 45 && latterchar < 58 { startingat += replyindex + 3; continue } // '.' => '9'
			return e
		}
	}
	return ""
}

