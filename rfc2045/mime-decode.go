// Copyright (C) 2020,2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 

package rfc2045
import "io"
import "bytes"
import "strings"
import "encoding/base64"
import "mime/quotedprintable"

// DecodeB decodes Base64 encoded text.
//   Arguments:
//     - argv0 (string): Base64-Encoded text.
//     - argv1 (string): Character set name.
//   Returns:
//     - (string): Decoded text.
func DecodeB(argv0 string, argv1 string) (string, error) {
	if len(argv0) < 8 { return argv0, nil }

	base64text := strings.ReplaceAll(strings.TrimSpace(argv0), "\n", "")
	cv, nyaan  := base64.StdEncoding.DecodeString(base64text); if nyaan != nil { return "", nyaan }
	return string(cv), nil
}

// DecodeQ decodes Quoted-Pritable encdoed text.
//   Arguments:
//     - argv0 (string): Quoted-Printable encoded text.
//   Returns:
//     - (string): Decoded text.
//     - (error):  Decoding error.
func DecodeQ(argv0 string) (string, error) {
	if len(argv0)  < 8 { return argv0, nil }
	decodingif := quotedprintable.NewReader(bytes.NewReader([]byte(argv0)))

	var readbuffer bytes.Buffer
	_, nyaan := io.Copy(&readbuffer, decodingif); if nyaan != nil { return "", nyaan }
	return readbuffer.String(), nil
}

