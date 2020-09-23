// Copyright (C) 2020 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package mail

import "os"
import "io/ioutil"
import "path/filepath"

// readMaildir() is a Maildir/ reader, works as a iterator.
func (this *Mail) readMaildir() (*string, error) {
	// @return   [*string]  Contents of the each file in the Maildir/
	// @return   [error]    It has reached to the end of the Maildir/
	if this.handle == nil {
		// Open the Maildir/
		filehandle, oops := os.Open(this.Dir)
		if oops != nil { return nil, oops }
		this.handle = filehandle // Successfully opened the Maildir/
	}

	var emailblock string
	if emailfiles, oops := this.handle.Readdir(1); oops == nil {
		// Read each email file in the Maildir/
		for _, ef := range emailfiles {
			this.offset += 1

			if ef.IsDir() || ef.Size() == 0 {
				// The element is a directory in the Maildir/, OR the size of email file is 0
				continue
			}

			this.Size = ef.Size()
			this.File = ef.Name()
			this.Path = filepath.Clean(filepath.FromSlash(this.Dir + "/" + ef.Name()))
			emailbytes := make([]byte, ef.Size())

			if readbuffer, oops := ioutil.ReadFile(this.Path); oops == nil {
				// No error, successfully opened each email file
				emailbytes = append(emailbytes, readbuffer...)

			} else {
				// Failed to read the email file
				continue
			}
			emailblock = string(emailbytes)
			break
		}
	} else {
		// Completed to read the Maildir/ or Failed to read the Maildir/
		return nil, oops
	}

	return &emailblock, nil
}

