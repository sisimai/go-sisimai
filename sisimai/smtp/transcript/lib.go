// Copyright (C) 2022,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package transcript

//                _           ___                                 _       _   
//  ___ _ __ ___ | |_ _ __   / / |_ _ __ __ _ _ __  ___  ___ _ __(_)_ __ | |_ 
// / __| '_ ` _ \| __| '_ \ / /| __| '__/ _` | '_ \/ __|/ __| '__| | '_ \| __|
// \__ \ | | | | | |_| |_) / / | |_| | | (_| | | | \__ \ (__| |  | | |_) | |_ 
// |___/_| |_| |_|\__| .__/_/   \__|_|  \__,_|_| |_|___/\___|_|  |_| .__/ \__|
//                   |_|                                           |_|        
import "strings"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import "sisimai/smtp/command"

/* SMTP Transcript log is like the following ------------------------------------------------------
 Out: 220 mx22.example.com ESMTP MAIL SYSTEM
 In:  EHLO mx22.example.com
 Out: 250-mx22.example.com
 Out: 250-PIPELINING
 Out: 250-SIZE 33554432
 Out: 250-ETRN
 Out: 250-STARTTLS
 Out: 250-AUTH PLAIN LOGIN DIGEST-MD5 CRAM-MD5
 Out: 250-AUTH=PLAIN LOGIN DIGEST-MD5 CRAM-MD5
 Out: 250-XFORWARD NAME ADDR PROTO HELO SOURCE PORT IDENT
 Out: 250-ENHANCEDSTATUSCODES
 Out: 250-8BITMIME
 Out: 250-DSN
 Out: 250 CHUNKING
 In:  XFORWARD NAME=neko2-nyaan3.y.example.co.jp ADDR=230.0.113.2
     PORT=53672
 Out: 250 2.0.0 Ok
 In:  XFORWARD PROTO=SMTP HELO=neko2-nyaan3.y.example.co.jp
     IDENT=2LYC6642BLzFK3MM SOURCE=REMOTE
 Out: 250 2.0.0 Ok
 In:  MAIL
     FROM:<CP85l3ba044547=libsisimai.net@err.y.example.co.jp>
     SIZE=20088
 Out: 250 2.1.0 Ok
 In:  RCPT TO:<kijitora@libsisimai.net> ORCPT=rfc822;kijitora@libsisimai.net
 Out: 250 2.1.5 Ok
 In:  DATA
 Out: 354 End data with <CR><LF>.<CR><LF>
 Out: 451 4.3.0 Error: queue file write error
 In:  QUIT
 Out: 221 2.0.0 Bye
------------------------------------------------------------------------------------------------ */
type ResponseTable struct {
	Reply    string  // SMTP reply code such as 550
	Status   string  // SMTP status code such as 5.1.1
	Text   []string  // Response text string
}

type TranscriptLog struct {
	Command   string            // SMTP Command
	Argument  string            // An argumenet of each SMTP command sent from a client
	Parameter map[string]string // Parameter pairs of the SMTP command
	Response  ResponseTable     // A Response from an SMTP server
}

// Rise() returns the decoded transcript of the SMTP session and makes the structured data
func Rise(argv0, argv1, argv2 string) []TranscriptLog {
	// @param    string  argv0   A transcript text MTA returned
	// @param    string  argv1   A label string of a SMTP cilent
	// @param    string  argv2   A label string of a SMTP server
	// @return   []TranscriptLog Structured data
	if len(argv0) < 1 { return []TranscriptLog{} }
	if len(argv1) < 1 { argv1 = ">>>" } // Label for an SMTP client
	if len(argv2) < 1 { argv2 = "<<<" } // Label for an SMTP server

	// 1. Get the position of ">>>" and "<<<"
	p1 := strings.Index(argv0, argv1); if p1 < 0 { return []TranscriptLog{} }
	p2 := strings.Index(argv0, argv2); if p2 < 0 { return []TranscriptLog{} }

	// 2. Remove the head of the argv0 to the first "<<<" or ">>>"
	sessionlog := []string{}        // Each line of the SMTP transcript log(argv0)
	transcript := []TranscriptLog{} // The list of TranscriptLog{}

	if p2 < p1 {
		// An SMTP server response starting with "<<<" is the first
		argv0 = argv0[p2:]

	} else {
		// An SMTP command starting with ">>>" is the first
		argv0 = argv0[p1:]
	}

	// 3. Remove strings from the first blank line to the tail
	p3 := strings.Index(argv0, "\n\n"); if p3 > 0 { argv0 = argv0[:p3 + 1] }

	// 4. Replace label strings of SMTP client/server at the each line
	for _, e := range strings.Split(argv0, "\n") {
		// Replace the following labels
		e = strings.TrimLeft(e, " ")

		if strings.HasPrefix(e, argv1) || strings.HasPrefix(e, argv2) {
			// - The line starts with ">>>" or the specified label in argv1
			// - The line starts with "<<<" or the specified label in argv2
			if strings.HasPrefix(e, argv1) {
				// 1. argv1 => ">>> " (leading a single space character)
				e = strings.Replace(e, argv1, ">>>", 1)
				for strings.HasPrefix(e, ">>>  ") { e = strings.Replace(e, ">>>  ", ">>> ", 1) }

			} else {
				// 2. argv2 => "<<< " (leading a single space character)
				e = strings.Replace(e, argv2, "<<<", 1)
				for strings.HasPrefix(e, "<<<  ") { e = strings.Replace(e, "<<<  ", "<<< ", 1) }
			}
			sessionlog = append(sessionlog, e)

		} else {
			// The line neither include ">>>" nor "<<<"
			// Concatenate folded lines to each previous line with a single " "(space character)
			ll := len(sessionlog); if ll == 0 { continue }
			sessionlog[ll - 1] += " " + strings.TrimLeft(e, " ")
		}
	}
	if len(sessionlog) == 0 { return transcript }

	// 5. Read each SMTP command and server response
	var cursession *TranscriptLog
	for _, e := range sessionlog {
		// Create []TranscriptLog struct
		if strings.HasPrefix(e, ">>> ") {
			// >>> SMTP-Command Arguments (Sent by the client)
			thecommand := command.Find(e); if len(thecommand) == 0 { continue }
			commandarg := strings.TrimLeft(e[strings.Index(e, thecommand) + len(thecommand):], " ")
			uppercased := strings.ToUpper(commandarg)
			parameters := "" // Command parameters of MAIL, RCPT

			transcript = append(transcript, *(new(TranscriptLog)))
			cursession = &(transcript[len(transcript) - 1])
			cursession.Command = strings.ToUpper(thecommand)
			cursession.Parameter = map[string]string{}

			if thecommand == "MAIL" || thecommand == "RCPT" || thecommand == "XFORWARD" {
				// "MAIL FROM" or "RCPT TO" or "XFORWARD"
				if strings.HasPrefix(uppercased, "FROM:") || strings.HasPrefix(uppercased, "TO:") {
					// >>> MAIL FROM: <neko@example.com> SIZE=65535
					// >>> RCPT TO: <kijitora@example.org>
					p4 := strings.IndexByte(commandarg, '<'); if p4 < 0 { continue }
					p5 := strings.IndexByte(commandarg, '>'); if p5 < 0 { continue }
					cursession.Argument = commandarg[p4 + 1:p5]

					if len(commandarg) > p5 {
						// Store the value of the SMTP command arguments
						parameters = strings.TrimLeft(commandarg[p5 + 1:], " ")
					}
				} else {
					// >>> XFORWARD NAME=neko2-nyaan3.y.example.co.jp ADDR=230.0.113.2 PORT=53672
					// <<< 250 2.0.0 Ok
					// >>> XFORWARD PROTO=SMTP HELO=neko2-nyaan3.y.example.co.jp IDENT=2LYC6642BLzFK3MM SOURCE=REMOTE
					// <<< 250 2.0.0 Ok
					parameters = commandarg
					commandarg = ""
				}

				for _, f := range strings.Split(parameters, " ") {
					// SIZE=22022, PROTO=SMTP, and so on
					if strings.IndexByte(f, '=') < 1             { continue }
					if len(f) < 3                                { continue }
					ee := strings.Split(f, "="); if len(ee) != 2 { continue }
					cursession.Parameter[strings.ToLower(ee[0])] = ee[1]
				}
			}
		} else {
			// <<< SMTP Server Response
			if strings.Index(e, "<<< ") != 0 { continue }
			if len(transcript) == 0 {
				// The first server response
				// Insert "CONN" as a pseudo SMTP command
				transcript = append(transcript, *(new(TranscriptLog)))
				cursession = &(transcript[len(transcript) - 1])
				cursession.Command = "CONN"
			}

			// Out: 220 mx22.example.com ESMTP MAIL SYSTEM
			cursession.Response.Reply  = reply.Find(e, "")
			cursession.Response.Status = status.Find(e, "")
			cursession.Response.Text   = append(cursession.Response.Text, e[4:])
		}
	}
	return transcript
}

// *TranscriptLog.Void() returns true if it does not include any transcript log
func(this *TranscriptLog) Void() bool {
	if len(this.Command) == 0 { return false }
	return true
}

