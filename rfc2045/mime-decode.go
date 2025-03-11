// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 

package rfc2045
import "io"
import "fmt"
import "mime"
import "strings"
import "mime/quotedprintable"

// DecodeB decodes Base64 encoded text.
//   Arguments:
//     - argv0 (string): Base64-Encoded text
//     - argv1 (string): Character set name
//   Returns:
//     - (string):       Decoded text
func DecodeB(argv0 string, argv1 string) (string, error) {
	if len(argv0)  < 8 { return argv0, nil }
	if len(argv1) == 0 { argv1 = "utf-8"   }

	decodingif := new(mime.WordDecoder)
	base64text := strings.Join(strings.Split(strings.TrimSpace(argv0), "\n"), "")
	base64text  = fmt.Sprintf("=?%s?B?%s?=", argv1, base64text)

	plain, nyaan := decodingif.Decode(base64text); if nyaan != nil { return "", nyaan }
	return plain, nil
}

// DecodeQ() decodes Quoted-Pritable encdoed text
//   Arguments:
//     - argv0 (string): Quoted-Printable encoded text
//   Returns:
//     - (string):       Decoded text
//     - (error):        Decoding error
func DecodeQ(argv0 string) (string, error) {
	readstring := strings.NewReader(argv0)
	decodingif := quotedprintable.NewReader(readstring)
	plainvalue := ""

	// Failed to decode the quoted-printable text
	plain, nyaan := io.ReadAll(decodingif); if nyaan != nil { plainvalue = argv0 }
	if len(plain) > 0 { plainvalue = string(plain) }

	return plainvalue, nyaan
}

