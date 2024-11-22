// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package mail

//                  _ _ 
//  _ __ ___   __ _(_) |
// | '_ ` _ \ / _` | | |
// | | | | | | (_| | | |
// |_| |_| |_|\__,_|_|_|
// sisimai/mail is a package for reading a UNIX mbox, a Maildir, or any email message input from Standard-in
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
 | Size       | o         |           | o         | o          |
 | NewLine    | o         |           | o         | o          |
*/
type EmailEntity struct {
	Kind    string   // "mailbox", "maildir", "memory" or "stdin"
	Path    string   // Path to the mbox, Maildir/, or "<MEMORY>" or "<STDIN>"
	Dir     string   // Directory name of mbox, Maildir/
	File    string   // File name of the mbox, each file in Maildir/
	NewLine uint8    // 0 = undefined, 1 = LF, 2 = CR, 3 = CRLF
	Size    int64    // Payload size
	offset  int64    // Offset position
	handle  *os.File // https://pkg.go.dev/os#File
	payload []string // Each email message
}

// Rise() is a constructor of EmailEntity struct
func Rise(argv0 string) (*EmailEntity, error) {
	// @param    string     argv0  Path to mbox or Maildir/
	// @return   *mail.EmailEntity Pointer to mail.EmailEntity struct
	thing := EmailEntity{}

	if argv0 == "STDIN" {
		// Read from STDIN
		thing.Kind = "stdin"
		thing.Path = "<STDIN>"

	} else if strings.Contains(argv0, "\n") {
		// Email data is in a string
		thing.Kind = "memory"
		thing.Path = "<MEMORY>"

	} else {
		// UNIX mbox or Maildir/
		if filestatus, nyaan:= os.Stat(argv0); nyaan == nil {
			// the file or the maildir exist
			thing.Path = argv0
			thing.Size = filestatus.Size()

			if filestatus.IsDir() {
				// Maildir/
				thing.Kind = "maildir"
				thing.Dir  = argv0

			} else {
				// UNIX mbox
				thing.Kind = "mailbox"
				thing.File = filepath.Base(argv0)
				thing.setNewLine()
			}
		} else {
			// Neither a mailbox nor a maildir exists
			return nil, nyaan
		}
	}
	return &thing, nil
}

// *EmailEntity.Read() is an email reader, works as an iterator.
func(this *EmailEntity) Read() (*string, error) {
	// @param    NONE
	// @return   *string Contents of mbox/Maildir
	var email *string // Email contents: headers and entire message body
	var nyaan  error  // Some errors while reading an email file

	switch this.Kind {
		case "mailbox":
			email, nyaan = this.readMailbox()
		case "maildir":
			email, nyaan = this.readMaildir()
/**
	TODO: IMPLEMENT
		case "memory":
			email, nyaan = this.readMemory()
		case "stdin":
			email, nyaan = this.readSTDIN()
**/
	}
	return email, nyaan
}

// *EmailEntity.setNewLine() returns true if the newline code is CRLF or CR or LF
func(this *EmailEntity) setNewLine() bool {
	// @param    NONE
	// @return   bool true if the newline code is CRLF or CR or LF
	var bufferedio *bufio.Reader
	var the1st1000 []byte
	var readbuffer string

	if this.Kind == "mailbox" {
		// UNIX mbox
		if filep, nyaan := os.Open(this.Path); nyaan != nil {
			// Failed to open the file
			fmt.Fprintf(os.Stderr, " *****error: %s\n", nyaan)
			this.NewLine = 0
			return false

		} else {
			// Successfully opened the mbox
			this.handle = filep
		}

		bufferedio = bufio.NewReader(this.handle)
		the1st1000 = make([]byte, 1000)
		_, nyaan := bufferedio.Read(the1st1000)
		if nyaan != nil && nyaan != io.EOF {
			// Failed to read the 1st 1000 bytes
			fmt.Fprintf(os.Stderr, " *****error: %s\n", nyaan)
			this.NewLine = 0
			return false
		}
	} else {
		// TODO: Maildir, STDIN, or memory
	}

	readbuffer = string(the1st1000)
	if strings.Contains(readbuffer, "\r\n") { this.NewLine = 3; return true }
	if strings.Contains(readbuffer, "\r")   { this.NewLine = 2; return true }
	if strings.Contains(readbuffer, "\n")   { this.NewLine = 1; return true }
	this.NewLine = 0; return false
}

