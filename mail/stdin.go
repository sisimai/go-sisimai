// Copyright (C) 2020,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _    ______ _____ ____ ___ _   _ 
//  _ __ ___   __ _(_) |  / / ___|_   _|  _ \_ _| \ | |
// | '_ ` _ \ / _` | | | / /\___ \ | | | | | | ||  \| |
// | | | | | | (_| | | |/ /  ___) || | | |_| | || |\  |
// |_| |_| |_|\__,_|_|_/_/  |____/ |_| |____/___|_| \_|

package mail
import "io"

// readSTDIN is an email reader input from the STDIN.
//   Returns:
//     - ([]byte): Contents of each email in the STDIN.
//     - (error):  Occurred error.
func (ee *EmailEntity) readSTDIN() ([]byte, error) {
	if ee.Size == 0 || ee.offset >= len(ee.payload) { return nil, io.EOF }

	emailblock := ee.payload[ee.offset]
	ee.offset++
	return emailblock, nil
}

