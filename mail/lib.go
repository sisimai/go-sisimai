// Copyright (C) 2020,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _ 
//  _ __ ___   __ _(_) |
// | '_ ` _ \ / _` | | |
// | | | | | | (_| | | |
// |_| |_| |_|\__,_|_|_|

// Package "mail" provides funtions for reading a UNIX mbox, a Maildir, or any email message input
// from Standard-in.
package mail

import "io"
import "os"
import "fmt"
import "bytes"
import "bufio"
import "strings"
import "path/filepath"
import "libsisimai.org/sisimai/v5/eb"

/* EmailEntity struct keeps each parameter of UNIX mbox, Maildir/.
 | FIELD      | UNIX mbox | Maildir/  | Memory    | <STDIN>    |
 |------------|-----------|-----------|-----------|------------|
 | Kind       | o         | o         | o         | o          |
 | Path       | o         | o         | o         | o          |
 | Dir        | o         | o         |           |            |
 | File       | o         | o         |           |            |
 | Size       | o         | o         | o         | o          |
 | newline    | o         |           | o         | o          |
 | offset     | o         | o         | o         | o          |
 | handle     | o         | o         |           |            |
 | payload    |           | o         | o         | o          |
*/
type EmailEntity struct {
	handle  *os.File // https://pkg.go.dev/os#File
	payload [][]byte // Each email message/file name
	Kind    string   // "mailbox", "maildir", "memory" or "stdin"
	Path    string   // Path to the mbox, Maildir/, or "<MEMORY>" or "<STDIN>"
	Dir     string   // Directory name of mbox, Maildir/
	File    string   // File name of the mbox, each file in Maildir/
	Size    int      // Payload size
	offset  int      // Offset position
	newline uint8    // 0 = undefined, 1 = LF, 2 = CR, 3 = CRLF
}

// Rise is a constructor of EmailEntity struct.
//   Arguments:
//     - path (string):  Path to an UNIX mbox, Maildir/, or "STDIN" for standard input.
//   Returns:
//     - (*EmailEntity): Pointer to mail.EmailEntity struct.
//     - (error):        Occurred error.
func Rise(path string) (*EmailEntity, error) {
	ee := EmailEntity{}

	if path == "STDIN" || strings.IndexByte(path, '\n') > -1 {
		// Read from STDIN or Memory(string)
		payload := []byte{}

		if path == "STDIN" {
			// For example, % cat ./bounce.eml | go run sisimai.go STDIN
			ee.Kind = "stdin"
			ee.Path = "<STDIN>"

			// Read all strings from STDIN, and store them to ee.payload
			// TODO: In the case of that the input data is a binary
			stdin, nyaan  := io.ReadAll(os.Stdin); if nyaan != nil { return &ee, nyaan }
			if textlength := len(stdin); textlength == 0 || textlength > eb.XeBYTE {
				// The input text is empty or too large (2GB)
				return &ee, fmt.Errorf("input text is empty or too large: %d bytes", textlength)
			}
			payload = stdin

		} else {
			// Email data is in a string(memory)
			if textlength := len(path); textlength < 3 || textlength > eb.XeBYTE {
				// The input text is empty or too large (2GB)
				return &ee, fmt.Errorf("input text is empty or too large: %d bytes", textlength)
			}
			ee.Kind = "memory"
			ee.Path = "<MEMORY>"
			payload = []byte(path)
		}

		if countUnixMboxFrom(payload) < 2 {
			// There is 1 or 0 "From " line in the payload
			ee.payload = append(ee.payload, payload)
			ee.Size = len(payload)

		} else {
			// There is 2 or more "From " line in the payload
			for _, uf := range bytes.Split(payload, []byte("\nFrom ")) {
				// Split by "From "
				if len(uf) == 0 { continue }
				cv        := append(append([]byte("From "), uf...), '\n')
				ee.payload = append(ee.payload, cv)
				ee.Size   += len(cv)
			}
		}
		ee.setNewLine()

	} else {
		// UNIX mbox or Maildir/
		if filestatus, nyaan:= os.Stat(path); nyaan == nil {
			// the file or the maildir exist
			ee.Path = path

			if filestatus.IsDir() {
				// Maildir/
				ee.Kind = "maildir"
				ee.Dir  = path
				cw, ce := ee.listMaildir(); if ce != nil { return &ee, ce }
				ee.Size = cw

			} else {
				// UNIX mbox
				cw := filestatus.Size(); if cw == 0 || cw > eb.XeBYTE {
					// The mbox is empty or too large (2GB)
					return &ee, fmt.Errorf("%s is empty or too large: %d bytes", path, ee.Size)
				}
				ee.Size = int(cw)
				ee.Kind = "mailbox"
				ee.File = filepath.Base(path)
				ee.Dir  = filepath.Dir(path)
				ee.setNewLine()
			}
		} else {
			// Neither a mailbox nor a maildir exists
			return nil, nyaan
		}
	}
	return &ee, nil
}

// countUnixMboxFrom returns the number of "From " line of the UNIX mbox.
//   Arguments:
//     - mesg ([]byte): Pointer to the entire email message.
//   Returns:
//     - (uint): The number of "From " lines.
func countUnixMboxFrom(mesg []byte) uint {
	if len(mesg) < 5 || bytes.HasPrefix(mesg, []byte("From ")) == false { return 0 }
	return uint(bytes.Count(mesg, []byte("\nFrom ")))
}

// *EmailEntity.Read is an email reader, works like an iterator.
//   Returns:
//     - ([]byte): Each email message one by one.
//     - (error):   Occurred error
func(ee *EmailEntity) Read() ([]byte, error) {
	var email []byte // Email contents: headers and entire message body
	var nyaan  error // Some errors while reading an email file

	switch ee.Kind {
		case "maildir": email, nyaan = ee.readMaildir()
		case "mailbox": email, nyaan = ee.readMailbox()
		case "memory":  email, nyaan = ee.readMemory()
		case "stdin":   email, nyaan = ee.readSTDIN()
	}
	return email, nyaan
}

// *EmailEntity.setNewLine set a new line type(CRLF, CR, LF) to EmailEntity.newline field.
func(ee *EmailEntity) setNewLine() {
	if ee.Kind == "maildir" { return }
	readbuffer := make([]byte, 1000)

	if ee.Kind == "mailbox" || ee.Kind == "stdin" {
		// UNIX mbox or STDIN
		var bufferedio *bufio.Reader

		if ee.Kind == "mailbox" {
			// UNIX mbox
			filep, nyaan := os.Open(ee.Path); if nyaan != nil { return }
			ee.handle  = filep
			bufferedio = bufio.NewReader(ee.handle)

		} else {
			// STDIN
			bufferedio = bufio.NewReader(os.Stdin)
		}

		the1st1000  := make([]byte, 1000)
		if _, nyaan := bufferedio.Read(the1st1000); nyaan != nil && nyaan != io.EOF { return }
		readbuffer   = the1st1000

	} else {
		// Memory
		if len(ee.payload) == 0 || len(ee.payload[0]) == 0 { ee.newline = 0; return }
		readbuffer = ee.payload[0][:min(1000, len(ee.payload[0]))]
	}

	if bytes.Contains(readbuffer, []byte("\r\n")) { ee.newline = 3; return }
	if bytes.IndexByte(readbuffer, '\r') > -1     { ee.newline = 2; return }
	if bytes.IndexByte(readbuffer, '\n') > -1     { ee.newline = 1; return }
	ee.newline = 0
}

