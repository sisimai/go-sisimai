// Copyright (C) 2020,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _    __               _ _     _ _      
//  _ __ ___   __ _(_) |  / / __ ___   __ _(_) | __| (_)_ __ 
// | '_ ` _ \ / _` | | | / / '_ ` _ \ / _` | | |/ _` | | '__|
// | | | | | | (_| | | |/ /| | | | | | (_| | | | (_| | | |   
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|\__,_|_|_|\__,_|_|_|   

package mail
import "io"
import "os"
import "errors"
import "strings"
import "strconv"
import "path/filepath"
import "libsisimai.org/sisimai/v5/eb"

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

	xe := []string{} // Stack occurred errors while reading the Maildir/
	for _, e := range direntries {
		// Read each email file in the Maildir/ except the following emails:
		by := e.Size(); switch {
			// 1. Is a directory
			// 2. Is an empty file
			// 3. Is too large: eb.XeBYTE(64MB)
			case e.IsDir():      { xe = append(xe, e.Name() + " is a directory")   }
			case by < 1:         { xe = append(xe, e.Name() + " is an empty file") }
			case by > eb.XeBYTE: { xe = append(xe, e.Name() + " is larger than " + strconv.Itoa(eb.XeBYTE)) } 
			default:               ee.payload = append(ee.payload, e.Name())
		}
	}
	nyaan = ee.handle.Close(); if nyaan != nil { xe = append(xe, nyaan.Error()) }
	ee.handle = nil

	// There is one or more errors occurred.
	if len(xe) > 0 { return len(ee.payload), errors.New(strings.Join(xe, "\n")) }
	return len(ee.payload), nil
}

// readMaildir is an email reader in the Maildir/, works like a iterator.
//   Returns:
//     - (*string): Contents of each email file in the Maildir/ one by one.
//     - (error):   Occurred error.
func(ee *EmailEntity) readMaildir() (*string, error) {
	if ee.Size == 0         { return nil, errors.New("there is no email file in " + ee.Dir) }
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

