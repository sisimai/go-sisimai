// Copyright (C) 2020 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package mail

import "io"
import "os"
import "strings"
import sisimoji "sisimai/string"

const MboxHeader = "\nFrom "
const BufferSize = 512

// readMailbox() is a UNIX mbox reader, works as a iterator.
func (this *Mail) readMailbox() (*string, error) {
	// @return   [*string]  Contents of the mbox
	// @return   [error]    It has reached to the end of the mbox

	// The method has been completed to read the mbox
	if this.offset >= this.Size { return nil, io.EOF }
	if this.handle == nil {
		// Open the mbox, and read at the offset position
		filehandle, oops := os.Open(this.Path)
		if oops != nil { return nil, oops }
		this.handle = filehandle // Successfully opened the mbox
	}

	var emailblock string
	for {
		// Read mbox until the EOF
		loopbuffer := ""
		readbuffer := make([]byte, BufferSize)

		if by, oops := this.handle.ReadAt(readbuffer, this.offset); oops == nil {
			// No error returned at reading the mbox, append the read buffer into the loopbuffer
			loopbuffer += string(readbuffer)
			loopbuffer  = *(sisimoji.ToLF(&loopbuffer))
			fromlindex := strings.Index(loopbuffer, MboxHeader)

			if fromlindex > 0 {
				// From line is included in the read buffer
				emailblock += loopbuffer[:fromlindex + 1]
				this.offset = int64(len(emailblock))
				break

			} else {
				// There is no "From " string in the loopbuffer, try to find "From " string in the previous
				// loop and the latest buffer
				tempbuffer := emailblock + loopbuffer
				fromlindex  = strings.Index(tempbuffer, MboxHeader)

				if fromlindex > 0 {
					// "From " string exists across the emailblock and the loopbuffer
					emailblock  = tempbuffer[:len(tempbuffer) - fromlindex + 1]
					this.offset = int64(len(emailblock))
					break

				} else {
					// "From " string dit not appeared yet
					emailblock  += loopbuffer
					this.offset += int64(len(loopbuffer))
				}
			}
		} else {
			// There is any failure on reading the mbox
			if oops == io.EOF {
				// Reached to the end of the mbox
				tempbuffer  := string(readbuffer[:by])
				emailblock  += *(sisimoji.ToLF(&tempbuffer))
				this.offset += int64(len(tempbuffer))
				this.handle.Close()
				break

			} else {
				// Something wrong
				return nil, oops
			}
		}
	} // The end of the loop(for)

	return &emailblock, nil
}

