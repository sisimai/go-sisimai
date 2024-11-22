// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package mail

//                  _ _    __               _ _     _ _      
//  _ __ ___   __ _(_) |  / / __ ___   __ _(_) | __| (_)_ __ 
// | '_ ` _ \ / _` | | | / / '_ ` _ \ / _` | | |/ _` | | '__|
// | | | | | | (_| | | |/ /| | | | | | (_| | | | (_| | | |   
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|\__,_|_|_|\__,_|_|_|   
import "os"
import "io/ioutil"
import "path/filepath"

// readMaildir() is a Maildir/ reader, works as a iterator.
func (this *EmailEntity) readMaildir() (*string, error) {
	// @return   *string  Contents of the each file in the Maildir/
	// @return   error    It has reached to the end of the Maildir/
	if this.handle == nil {
		// Open the Maildir/
		filehandle, nyaan := os.Open(this.Dir); if nyaan != nil { return nil, nyaan }
		this.handle = filehandle // Successfully opened the Maildir/
	}

	emailblock := ""
	if emailfiles, nyaan := this.handle.Readdir(1); nyaan == nil {
		// Read each email file in the Maildir/
		for _, e := range emailfiles {
			// The element is a directory in the Maildir/, OR the size of email file is 0
			this.offset += 1
			if e.IsDir() || e.Size() == 0 { continue }

			this.Size = e.Size()
			this.File = e.Name()
			this.Path = filepath.Clean(filepath.FromSlash(this.Dir + "/" + e.Name()))

			emailbytes := make([]byte, e.Size())
			if readbuffer, nyaan := ioutil.ReadFile(this.Path); nyaan == nil {
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
		return nil, nyaan
	}

	return &emailblock, nil
}

