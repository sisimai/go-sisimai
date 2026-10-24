// Copyright (C) 2020,2022,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _    __         _               
//  _ __ ___   __ _(_) |  / / __ ___ | |__   _____  __
// | '_ ` _ \ / _` | | | / / '_ ` _ \| '_ \ / _ \ \/ /
// | | | | | | (_| | | |/ /| | | | | | |_) | (_) >  < 
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|_.__/ \___/_/\_\

package mail
import "io"
import "os"
import "bytes"
import "bufio"
import "libsisimai.org/sisimai/v5/eb"

// readMailbox is a UNIX mbox reader, works like a iterator.
//   Returns:
//     - (*string): Contents of each email in the UNIX mbox one by one.
//     - (error):   Occurred error.
func(ee *EmailEntity) readMailbox() (*string, error) {
	if ee.offset >= ee.Size { return nil, io.EOF } // The method has been completed to read the mbox
	if ee.handle == nil {
		// Open the UNIX mbox, and read at the offset position
		filehandle, nyaan := os.Open(ee.Path); if nyaan != nil { return nil, nyaan }
		ee.handle = filehandle // Successfully opened the mbox
	}

	seekoffset := int64(ee.offset);              if ee.offset  < 0 { seekoffset = 0    }
	_, nyaan   := ee.handle.Seek(seekoffset, 0); if nyaan != nil   { return nil, nyaan }
	lineending := 0;                             if ee.newline > 2 { lineending = 1    }
	readbuffer := bytes.Buffer{}; readbuffer.Grow(4096)
	emailblock := ""
	thisheight := 0

	unixmboxio := bufio.NewScanner(ee.handle)
	unixmboxio.Buffer(make([]byte, 64 * 1024), eb.XeBYTE); for unixmboxio.Scan() {
		// Read the UNIX mbox until the EOF
		e := unixmboxio.Bytes(); if bytes.HasPrefix(e, []byte("From ")) && readbuffer.Len() > 0 {
			// - The line is a UNIX From line such as "From MAILER-DAEMON Fri Feb  2 18:30:22 2018"
			// - This UNIX From line is the beginning of the second or later email message
			emailblock = readbuffer.String(); readbuffer.Reset()
			ee.offset += len(emailblock) + (thisheight * lineending)
			break
		}
		thisheight += 1
		readbuffer.Write(e); readbuffer.WriteByte('\n')
	}
	if nyaan := unixmboxio.Err(); nyaan != nil {
		// Possible errors captured by unixmboxio.Err():
		// - bufio.ErrTooLong: A single line exceeds the maximum buffer size (eb.XeBYTE)
		//   - "bufio.Scanner: token too long"
		// - I/O errors: File or storage read failures (e.g., EIO, permission, network disconnection)
		// - Stream errors: Corrupted reader state or unexpected stream interruption
		if ee.handle != nil { ee.handle.Close() }
		return nil, nyaan
	}

	if readbuffer.Len() > 0 {
		// The last email message in the UNIX mbox
		emailblock = readbuffer.String()
		ee.offset += readbuffer.Len() + (thisheight * lineending)
		if nyaan  := ee.handle.Close(); nyaan != nil { return nil, nyaan }
	}
	return &emailblock, nil
}

