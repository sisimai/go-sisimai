// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//                  _ _    ______ _____ ____ ___ _   _ 
//  _ __ ___   __ _(_) |  / / ___|_   _|  _ \_ _| \ | |
// | '_ ` _ \ / _` | | | / /\___ \ | | | | | | ||  \| |
// | | | | | | (_| | | |/ /  ___) || | | |_| | || |\  |
// |_| |_| |_|\__,_|_|_/_/  |____/ |_| |____/___|_| \_|

package mail
import "io"

func (this *EmailEntity) readSTDIN() (*string, error) {
	// @return   *string  Contents of the mbox input from STDIN
	// @return   error    It has reached to the end of the mbox
	if this.Size == 0 || this.offset >= len(this.payload) { return nil, io.EOF }

	emailblock := this.payload[this.offset]
	this.offset++
	return &emailblock, nil
}

