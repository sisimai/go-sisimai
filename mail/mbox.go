// Copyright (C) 2020,2022,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _    __         _               
//  _ __ ___   __ _(_) |  / / __ ___ | |__   _____  __
// | '_ ` _ \ / _` | | | / / '_ ` _ \| '_ \ / _ \ \/ /
// | | | | | | (_| | | |/ /| | | | | | |_) | (_) >  < 
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|_.__/ \___/_/\_\

package mail
import "io"
import "os"
import "bufio"
import "strings"

// readMailbox is a UNIX mbox reader, works like a iterator.
//   Returns:
//     - (*string): Contents of each email in the UNIX mbox one by one.
//     - (error):   Occurred error.
func(this *EmailEntity) readMailbox() (*string, error) {
	if this.offset >= this.Size { return nil, io.EOF } // The method has been completed to read the mbox
	if this.handle == nil {
		// Open the UNIX mbox, and read at the offset position
		filehandle, nyaan := os.Open(this.Path); if nyaan != nil { return nil, nyaan }
		this.handle = filehandle // Successfully opened the mbox
	}

	seekoffset := int64(this.offset);              if this.offset  < 0 { seekoffset = 0 }
	_, nyaan   := this.handle.Seek(seekoffset, 0); if nyaan != nil { return nil, nyaan  }
	lineending := 0;                               if this.newline > 2 { lineending = 1 }
	readbuffer := strings.Builder{}; readbuffer.Grow(4096)
	emailblock := ""
	thisheight := 0

	unixmboxio := bufio.NewScanner(this.handle); for unixmboxio.Scan() {
		// Read the UNIX mbox until the EOF
		e := unixmboxio.Text()
		if strings.HasPrefix(e, "From ") && readbuffer.Len() > 0 {
			// - The line is a UNIX From line such as "From MAILER-DAEMON Fri Feb  2 18:30:22 2018"
			// - This UNIX From line is the beginning of the second or later email message
			emailblock   = readbuffer.String(); readbuffer.Reset()
			this.offset += len(emailblock) + (thisheight * lineending)
			break
		}
		thisheight += 1
		readbuffer.WriteString(e + "\n")
	}

	if readbuffer.Len() > 0 {
		// The last email message in the UNIX mbox
		emailblock   = readbuffer.String()
		this.offset += readbuffer.Len() + (thisheight * lineending)
		if nyaan := this.handle.Close(); nyaan != nil { return nil, nyaan }
	}
	return &emailblock, nil
}

