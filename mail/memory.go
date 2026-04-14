// Copyright (C) 2020,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _    __                                         
//  _ __ ___   __ _(_) |  / / __ ___   ___ _ __ ___   ___  _ __ _   _ 
// | '_ ` _ \ / _` | | | / / '_ ` _ \ / _ \ '_ ` _ \ / _ \| '__| | | |
// | | | | | | (_| | | |/ /| | | | | |  __/ | | | | | (_) | |  | |_| |
// |_| |_| |_|\__,_|_|_/_/ |_| |_| |_|\___|_| |_| |_|\___/|_|   \__, |
//                                                              |___/ 

package mail
import "io"

// readMemory is an email reader stored in a variable as a string.
//   Returns:
//     - ([]byte): Contents of each email in EmailEntity.payload field.
//     - (error):  Occurred error.
func (ee *EmailEntity) readMemory() ([]byte, error) {
	if ee.Size == 0 || ee.offset >= len(ee.payload) { return nil, io.EOF }

	ee.offset++
	return []byte(ee.payload[ee.offset]), nil
}

