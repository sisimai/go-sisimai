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
//     - text (string): Base64-Encoded text.
//   Returns:
//     - (string): Decoded text.
func DecodeB(text string) (string, error) {
	if len(text) < 8 { return text, nil }

	base64text := strings.ReplaceAll(strings.TrimSpace(text), "\n", "")
	cv, nyaan  := base64.StdEncoding.DecodeString(base64text); if nyaan != nil { return "", nyaan }
	return string(cv), nil
}

// DecodeQ decodes Quoted-Pritable encdoed text.
//   Arguments:
//     - text (string): Quoted-Printable encoded text.
//   Returns:
//     - (string): Decoded text.
//     - (error):  Decoding error.
func DecodeQ(text string) (string, error) {
	if len(text) < 8 { return text, nil }
	decodingif := quotedprintable.NewReader(bytes.NewReader([]byte(text)))

	var readbuffer bytes.Buffer
	_, nyaan := io.Copy(&readbuffer, decodingif); if nyaan != nil { return "", nyaan }
	return readbuffer.String(), nil
}

