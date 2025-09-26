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
//     - (int):   The number of email files in the Maildir/.
//     - (error): Occurred error.
func(ee *EmailEntity) listMaildir() (int, error) {
	if ee.handle == nil {
		// Open the Maildir/
		filehandle, nyaan := os.Open(ee.Dir);  if nyaan != nil { return 0, nyaan }
		ee.handle = filehandle // Successfully opened the Maildir/
	}
	direntries, nyaan := ee.handle.Readdir(0); if nyaan != nil { return 0, nyaan }
	for _, e := range direntries {
		// Read each email file in the Maildir/
		if e.IsDir() == false || e.Size() > 0 { ee.payload = append(ee.payload, e.Name()) }
	}
	nyaan = ee.handle.Close(); ee.handle = nil
	return len(ee.payload), nyaan
}

// readMaildir is an email reader in the Maildir/, works like a iterator.
//   Returns:
//     - (*string): Contents of each email file in the Maildir/ one by one.
//     - (error):   Occurred error.
func(ee *EmailEntity) readMaildir() (*string, error) {
	if ee.Size == 0         { return nil, fmt.Errorf("there is no email file in %s", ee.Dir) }
	if ee.Size <= ee.offset { return nil, io.EOF }

	for {
		// Try to read the email file
		ee.File = ee.payload[ee.offset]; ee.offset++
		ee.Path = filepath.Join(ee.Dir, ee.File)
		b, nyaan := os.ReadFile(ee.Path); if nyaan != nil || len(b) == 0 {
			// Failed to read the email file or the email file is empty
			if ee.offset >= ee.Size { return nil, io.EOF }
			continue
		}
		cv := string(b); return &cv, nil
	}
}

