// Copyright (C) 2020,2024-2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  ____  _____ ____ ____   ___  _  _  ____  
// |  _ \|  ___/ ___|___ \ / _ \| || || ___| 
// | |_) | |_ | |     __) | | | | || ||___ \ 
// |  _ <|  _|| |___ / __/| |_| |__   _|__) |
// |_| \_\_|   \____|_____|\___/   |_||____/ 

package rfc2045
import "io"
import "bytes"
import "encoding/base64"
import "mime/quotedprintable"

// DecodeB decodes Base64 encoded text.
//   Arguments:
//     - text ([]byte): Base64-Encoded text.
//   Returns:
//     - ([]byte): Decoded text.
func DecodeB(text []byte) ([]byte, error) {
	if len(text) < 8 { return text, nil }

	base64text := bytes.ReplaceAll(bytes.TrimSpace(text), []byte("\n"), []byte(""))
	base64buff := make([]byte, base64.StdEncoding.DecodedLen(len(base64text)))
	cw, nyaan  := base64.StdEncoding.Decode(base64buff, base64text); if nyaan != nil { return []byte(""), nyaan }
	return base64buff[:cw], nil
}

// DecodeQ decodes Quoted-Pritable encdoed text.
//   Arguments:
//     - text ([]byte): Quoted-Printable encoded text.
//   Returns:
//     - ([]byte): Decoded text.
//     - (error):  Decoding error.
func DecodeQ(text []byte) ([]byte, error) {
	if len(text) < 8 { return text, nil }
	decodingif := quotedprintable.NewReader(bytes.NewReader(text))

	var readbuffer bytes.Buffer
	_, nyaan := io.Copy(&readbuffer, decodingif); if nyaan != nil { return []byte(""), nyaan }
	return readbuffer.Bytes(), nil
}

