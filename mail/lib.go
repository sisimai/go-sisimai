// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
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
import "bufio"
import "strings"
import "path/filepath"

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
	Kind    string   // "mailbox", "maildir", "memory" or "stdin"
	Path    string   // Path to the mbox, Maildir/, or "<MEMORY>" or "<STDIN>"
	Dir     string   // Directory name of mbox, Maildir/
	File    string   // File name of the mbox, each file in Maildir/
	Size    int      // Payload size
	offset  int      // Offset position
	newline uint8    // 0 = undefined, 1 = LF, 2 = CR, 3 = CRLF
	handle  *os.File // https://pkg.go.dev/os#File
	payload []string // Each email message/file name
}
const maximumSize = 2000 * 1024 * 1024 * 1024

// Rise() is a constructor of EmailEntity struct
func Rise(argv0 string) (*EmailEntity, error) {
	// @param    string     argv0  Path to mbox or Maildir/
	// @return   *mail.EmailEntity Pointer to mail.EmailEntity struct
	ee := EmailEntity{}

	if argv0 == "STDIN" || strings.IndexByte(argv0, '\n') > -1 {
		// Read from STDIN or Memory(string)
		payload := ""

		if argv0 == "STDIN" {
			// For example, % cat ./bounce.eml | go run sisimai.go STDIN
			ee.Kind = "stdin"
			ee.Path = "<STDIN>"

			for {
				// Read all strings from STDIN, and store them to ee.payload
				// TODO: In the case of that the input data is a binary
				stdin, nyaan  := io.ReadAll(os.Stdin); if nyaan != nil { return &ee, nyaan }
				if textlength := len(stdin); textlength == 0 || textlength > maximumSize {
					// The input text is empty or too large (2GB)
					return &ee, fmt.Errorf("input text is empty or too large: %d bytes", textlength)
				}
				payload = string(stdin)
				break
			}
		} else {
			// Email data is in a string(memory)
			if textlength := len(argv0); textlength == 0 || textlength > maximumSize {
				// The input text is empty or too large (2GB)
				return &ee, fmt.Errorf("input text is empty or too large: %d bytes", textlength)
			}
			ee.Kind = "memory"
			ee.Path = "<MEMORY>"
			payload = argv0
		}

		if cw := CountUnixMboxFrom(&payload); cw < 2 {
			// There is 1 or 0 "From " line in the payload
			ee.payload = append(ee.payload, payload)
			ee.Size = len(payload)

		} else {
			// There is 2 or more "From " line in the payload
			for _, uf := range strings.Split(payload, "\nFrom ") {
				// Split by "From "
				if uf == "" { continue }
				cv := fmt.Sprintf("From %s\n", uf)
				ee.payload = append(ee.payload, cv)
				ee.Size   += len(cv)
			}
		}
		ee.setNewLine() // TODO: Receive and check the return values

	} else {
		// UNIX mbox or Maildir/
		if filestatus, nyaan:= os.Stat(argv0); nyaan == nil {
			// the file or the maildir exist
			ee.Path = argv0

			if filestatus.IsDir() {
				// Maildir/
				ee.Kind = "maildir"
				ee.Dir  = argv0
				cw, ce := ee.listMaildir(); if ce != nil { return &ee, ce }
				ee.Size = cw

			} else {
				// UNIX mbox
				cw := filestatus.Size(); if cw == 0 || cw > maximumSize {
					// The mbox is empty or too large (2GB)
					return &ee, fmt.Errorf("%s is empty or too large: %d bytes", argv0, ee.Size)
				}
				ee.Size = int(cw)
				ee.Kind = "mailbox"
				ee.File = filepath.Base(argv0)
				ee.Dir  = filepath.Dir(argv0)
				ee.setNewLine() // TODO: Receive and check the return values
			}
		} else {
			// Neither a mailbox nor a maildir exists
			return nil, nyaan
		}
	}
	return &ee, nil
}

// CountUnixMboxFrom() returns the number of "From " line of the Unix mbox
func CountUnixMboxFrom(argv0 *string) uint {
	// @param    *string argv0  A pointer to the entire email message
	// @return    unit          The number of "From " lines
	if len(*argv0) < 5 || strings.HasPrefix(*argv0, "From ") == false { return 0 }
	cw := strings.Count(*argv0, "\nFrom ")
	return uint(cw)
}

// *EmailEntity.Read() is an email reader, works as an iterator.
func(this *EmailEntity) Read() (*string, error) {
	// @return   *string Contents of mbox/Maildir
	var email *string // Email contents: headers and entire message body
	var nyaan  error  // Some errors while reading an email file

	switch this.Kind {
		case "maildir": email, nyaan = this.readMaildir()
		case "mailbox": email, nyaan = this.readMailbox()
		case "memory":  email, nyaan = this.readMemory()
		case "stdin":   email, nyaan = this.readSTDIN()
	}
	return email, nyaan
}

// *EmailEntity.setNewLine() returns true if the newline code is CRLF or CR or LF
func(this *EmailEntity) setNewLine() (bool, error) {
	// @param    NONE
	// @return   bool true if the newline code is CRLF or CR or LF
	if this.Kind == "maildir" { return false, nil }
	var bufferedio *bufio.Reader
	var readbuffer string

	if this.Kind == "mailbox" || this.Kind == "stdin" {
		// UNIX mbox or STDIN
		if this.Kind == "mailbox" {
			// UNIX mbox
			if filep, nyaan := os.Open(this.Path); nyaan != nil {
				// Failed to open the file
				this.newline = 0
				return false, nyaan

			} else {
				// Successfully opened the mbox
				this.handle = filep
			}
			bufferedio = bufio.NewReader(this.handle)

		} else {
			// STDIN
			bufferedio = bufio.NewReader(os.Stdin)
		}

		the1st1000 := make([]byte, 1000)
		_, nyaan := bufferedio.Read(the1st1000)
		if nyaan != nil && nyaan != io.EOF {
			// Failed to read the 1st 1000 bytes
			this.newline = 0
			return false, nyaan
		}
		readbuffer = string(the1st1000)

	} else {
		// Memory
		if len(this.payload) ==  0 { this.newline = 0; return false, nil }
		if this.payload[0]   == "" { this.newline = 0; return false, nil }
		readbuffer = this.payload[0][:1000]
	}

	if strings.Contains(readbuffer, "\r\n")     { this.newline = 3; return true, nil }
	if strings.IndexByte(readbuffer, '\r') > -1 { this.newline = 2; return true, nil }
	if strings.IndexByte(readbuffer, '\n') > -1 { this.newline = 1; return true, nil }
	this.newline = 0; return false, nil
}

