// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _    __               _ _     _ _      
//  _ __ ___   __ _(_) |  / / __ ___   __ _(_) | __| (_)_ __ 
// | '_ ` _ \ / _` | | | / / '_ ` _ \ / _` | | |/ _` | | '__|
// | | | | | | (_| | | |/ /| | | | | | (_| | | | (_| | | |   
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|\__,_|_|_|\__,_|_|_|   

package mail
import "io"
import "os"
import "fmt"
import "path/filepath"

// listMaildir is a Maildir/ reader, works like a iterator.
//   Returns:
//     - (int):   The number of email files in the Maildir/
//     - (error): Occurred error
func(this *EmailEntity) listMaildir() (int, error) {
	if this.handle == nil {
		// Open the Maildir/
		filehandle, nyaan := os.Open(this.Dir);  if nyaan != nil { return 0, nyaan }
		this.handle = filehandle // Successfully opened the Maildir/
	}
	direntries, nyaan := this.handle.Readdir(0); if nyaan != nil { return 0, nyaan }
	for _, e := range direntries {
		// Read each email file in the Maildir/
		if e.IsDir() == false || e.Size() > 0 { this.payload = append(this.payload, e.Name()) }
	}
	this.handle.Close(); this.handle = nil
	return len(this.payload), nil
}

// readMaildir is an email reader in the Maildir/, works like a iterator.
//   Returns:
//     - (*string): Contents of each email file in the Maildir/ one by one
//     - (error):   Occurred error
func(this *EmailEntity) readMaildir() (*string, error) {
	if this.Size == 0           { return nil, fmt.Errorf("there is no email file in %s", this.Dir) }
	if this.Size <= this.offset { return nil, io.EOF }

	for {
		// Try to read the email file
		this.File = this.payload[this.offset]; this.offset++
		this.Path = filepath.Clean(filepath.FromSlash(this.Dir + "/" + this.File))
		b, nyaan := os.ReadFile(this.Path); if nyaan != nil || len(b) == 0 {
			// Failed to read the email file or the email file is empty
			if this.offset >= this.Size { return nil, io.EOF }
			continue
		}
		cv := string(b); return &cv, nil
	}
}

