// Copyright (C) 2020,2022,2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package mail

//                  _ _    __         _               
//  _ __ ___   __ _(_) |  / / __ ___ | |__   _____  __
// | '_ ` _ \ / _` | | | / / '_ ` _ \| '_ \ / _ \ \/ /
// | | | | | | (_| | | |/ /| | | | | | |_) | (_) >  < 
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|_.__/ \___/_/\_\
import "io"
import "os"
import "bufio"
import "strings"

// readMailbox() is a UNIX mbox reader, works as a iterator.
func(this *EmailEntity) readMailbox() (*string, error) {
	// @return   *string  Contents of the mbox
	// @return   error    It has reached to the end of the mbox
	if this.offset >= this.Size { return nil, io.EOF } // The method has been completed to read the mbox
	if this.handle == nil {
		// Open the UNIX mbox, and read at the offset position
		filehandle, nyaan := os.Open(this.Path); if nyaan != nil { return nil, nyaan }
		this.handle = filehandle // Successfully opened the mbox
	}

	seekoffset := this.offset;                     if this.offset  < 0 { seekoffset = 0 }
	_, nyaan   := this.handle.Seek(seekoffset, 0); if nyaan != nil { return nil, nyaan  }
	lineending := 0;                               if this.NewLine > 2 { lineending = 1 }
	unixmboxio := bufio.NewScanner(this.handle)
	readbuffer := ""
	emailblock := ""
	thisheight := 0

	for unixmboxio.Scan() {
		// Read the UNIX mbox until the EOF
		e := unixmboxio.Text()
		if strings.HasPrefix(e, "From ") && readbuffer != "" {
			// The line is a UNIX From line such as "From MAILER-DAEMON Fri Feb  2 18:30:22 2018"
			// This UNIX From line is the beginning of the second or later email message
			emailblock  = readbuffer; readbuffer = ""
			this.offset += int64(len(emailblock) + (thisheight * lineending))
			break
		}
		thisheight += 1
		readbuffer += e + "\n"
	}

	if readbuffer != "" {
		// The last email message in the UNIX mbox
		emailblock  = readbuffer
		this.offset += int64(len(readbuffer) + (thisheight * lineending))
		this.handle.Close()
	}
	return &emailblock, nil
}

