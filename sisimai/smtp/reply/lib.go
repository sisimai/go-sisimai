// Copyright (C) 2020-2021 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reply
import "strings"
import "strconv"
import sisimoji "sisimai/string"

/*
 http://www.ietf.org/rfc/rfc5321.txt
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

  4.2.3.  Reply Codes in Numeric Order
       211  System status, or system help reply
       214  Help message (Information on how to use the receiver or the meaning of a particular
            non-standard command; this reply is useful only to the human user)
       220  <domain> Service ready
       221  <domain> Service closing transmission channel
       250  Requested mail action okay, completed
       251  User not local; will forward to <forward-path> (See Section 3.4)
       252  Cannot VRFY user, but will accept message and attempt delivery (See Section 3.5.3)
       354  Start mail input; end with <CRLF>.<CRLF>
       421  <domain> Service not available, closing transmission channel (This may be a reply to
            any command if the service knows it must shut down)
       450  Requested mail action not taken: mailbox unavailable (e.g., mailbox busy or temporarily
            blocked for policy reasons)
       451  Requested action aborted: local error in processing
       452  Requested action not taken: insufficient system storage
       455  Server unable to accommodate parameters
       500  Syntax error, command unrecognized (This may include errors such as command line too long)
       501  Syntax error in parameters or arguments
       502  Command not implemented (see Section 4.2.4)
       503  Bad sequence of commands
       504  Command parameter not implemented
       550  Requested action not taken: mailbox unavailable (e.g., mailbox not found, no access,
            or command rejected for policy reasons)
       551  User not local; please try <forward-path> (See Section 3.4)
       552  Requested mail action aborted: exceeded storage allocation
       553  Requested action not taken: mailbox name not allowed (e.g., mailbox syntax incorrect)
       554  Transaction failed (Or, in the case of a connection-opening response, "No SMTP service here")
       555  MAIL FROM/RCPT TO parameters not recognized or not implemented
*/

// Find() returns an SMTP reply code found from the given string
func Find(argv0 string) string {
	// @param    [string] argv0  String including an SMTP reply code
	// @return   [string]        Found SMTP reply code or an empty string
	if len(argv0) < 3                     { return "" }
	if strings.Contains(argv0, "X-UNIX;") { return "" }

	argv0  = strings.ReplaceAll(argv0, "-", " ") // "550-5.1.1" => "550 5.1.1"
	found := ""

	for _, e := range strings.Fields(argv0) {
		// Find an SMTP reply code from each field
		e = strings.TrimSpace(e)  // Strip space characters
		e = strings.Trim(e, "[]") // Strip square brackets
		e = strings.Trim(e, "()") // Strip parentheses

		if len(e) < 3                               { continue } // Minimun length is 3: "550"
		if sisimoji.ContainsOnlyNumbers(e) == false { continue }

		if e[0:1] == "2" || e[0:1] == "4" || e[0:1] == "5" {
			// The first letter of an SMTP reply code is 2,4 or 5
			i, oops := strconv.Atoi(e)
			if oops != nil { continue }
			if i < 200     { continue }
			if i > 579     { continue }

			switch {
				case i > 199 && i < 254:
					// 200 OK
					found = e
					break

				case i > 399 && i < 480:
					// 421 Deferrerd
					found = e
					break

				case i > 499 && i < 580:
					// 550 User unknown
					found = e
					break

				default:
					continue
			}
		}
	}
	return found
}

