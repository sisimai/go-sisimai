// Copyright (C) 2020,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package mail

//                  _ _    __               _ _     _ _      
//  _ __ ___   __ _(_) |  / / __ ___   __ _(_) | __| (_)_ __ 
// | '_ ` _ \ / _` | | | / / '_ ` _ \ / _` | | |/ _` | | '__|
// | | | | | | (_| | | |/ /| | | | | | (_| | | | (_| | | |   
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|\__,_|_|_|\__,_|_|_|   
import "os"
import "fmt"
import "io/ioutil"
import "path/filepath"

// readMaildir() is a Maildir/ reader, works as a iterator.
func (this *EmailEntity) readMaildir() (*string, error) {
	// @return   *string  Contents of the each file in the Maildir/
	// @return   error    It has reached to the end of the Maildir/
	if this.handle == nil {
		// Open the Maildir/
		filehandle, nyaan := os.Open(this.Dir);  if nyaan != nil { return nil, nyaan }
		this.handle = filehandle // Successfully opened the Maildir/
	}

	emailblock := ""
	emailfiles, nyaan := this.handle.Readdir(1); if nyaan != nil { return &emailblock, nyaan }

	for _, e := range emailfiles {
		// Read each email file in the Maildir/
		this.offset += 1; if e.IsDir() { continue }
		this.Size    = e.Size()
		this.File    = e.Name()
		this.Path    = filepath.Clean(filepath.FromSlash(this.Dir + "/" + e.Name()))
		if this.Size == 0 { return &emailblock, fmt.Errorf("%s in Maildir/ is empty", e.Name()) }

		emailbytes := make([]byte, e.Size())
		readbuffer, nyaan := ioutil.ReadFile(this.Path); if nyaan != nil { return &emailblock, nyaan }
		emailbytes  = append(emailbytes, readbuffer...)
		emailblock  = string(emailbytes)
		break
	}
	return &emailblock, nil
}

