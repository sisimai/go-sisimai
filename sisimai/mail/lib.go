// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.

// sisimai/mail is a package for reading a UNIX mbox, a Maildir, or any email message input from Standard-in
package mail
import "os"
import "strings"
import "path/filepath"

/* Mail struct keeps each parameter of UNIX mbox, Maildir/.
  | FIELD      | UNIX mbox | Maildir/  | Memory    | <STDIN>    |
  |------------|-----------|-----------|-----------|------------|
  | Kind       | o         | o         | o         | o          |
  | Path       | o         | o         | o         | o          |
  | Dir        | o         | o         |           |            |
  | File       | o         | o         |           |            |
  | Size       | o         |           | o         | o          |
*/
type Mail struct {
	Kind    string   // "mailbox", "maildir", "memory" or "stdin"
	Path    string   // Path to the mbox, Maildir/, or "<MEMORY>" or "<STDIN>"
	Dir     string   // Directory name of mbox, Maildir/
	File    string   // File name of the mbox, each file in Maildir/
	Size    int64    // Payload size
	offset  int64    // Offset position
	handle  *os.File // os.Open()
	payload []string // Each email message
}

// Rise() is a constructor of Mail struct
func Rise(argv0 string) (*Mail, error) {
	// @param    string     argv0  Path to mbox or Maildir/
	// @return   *mail.Mail        Pointer to mail.Mail struct
	thing := Mail{}

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
		if filestatus, aargh := os.Stat(argv0); aargh == nil {
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
			}
		} else {
			// Neither a mailbox nor a maildir exists
			return nil, aargh
		}
	}
	return &thing, nil
}

// *Mail.Read() is an email reader, works as an iterator.
func(this *Mail) Read() (*string, error) {
	// @param    NONE
	// @return   *string Contents of mbox/Maildir
	var mail *string // Email contents: headers and entire message body
	var oops  error  // Some errors while reading an email file

	switch this.Kind {
		case "mailbox":
			mail, oops = this.readMailbox()
		case "maildir":
			mail, oops = this.readMaildir()
/**
	TODO: IMPLEMENT
		case "memory":
			mail, oops = this.readMemory()
		case "stdin":
			mail, oops = this.readSTDIN()
**/
	}
	return mail, oops
}

