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
import "golang.org/x/net/html/charset"
import "libsisimai.org/sisimai/moji"

// IsEncoded checks that the argument is MIME encoded string or not.
//   Arguments:
//     - argv0 (string): String to be checked that it is MIME encoded or not
//   Returns:
//     - (bool):         true if the argument is a MIME encoded string
func IsEncoded(argv0 string) bool {
	// For example, the argument is like "=?UTF-8?B?44OL44Oj44O844Oz?="
	argv0  = strings.ToUpper(argv0)

	if strings.Contains(argv0, "=?") == false          { return false } // Should begin with "=?"
	if strings.Contains(argv0, "?=") == false          { return false } // Should end with "?="
	if len(argv0) < 8                                  { return false } // Should be 8 or more length
	if moji.ContainsAny(argv0, []string{"?B?", "?Q?"}) { return true  }
	return false
}

// DecodeH decodes the value of email header which is a MIME-Encoded string.
//   Arguments:
//     - argv0 (string): MIME-Encoded text
//   Returns:
//     - (string):       Decoded text
func DecodeH(argv0 string) (string, error) {
	if argv0 == "" { return "", nil }

	decodingif := new(mime.WordDecoder); if CharacterSet(argv0) != "UTF-8" {
		// The character set is not UTF-8
		decodingif.CharsetReader = func(c string, v io.Reader) (io.Reader, error) {
			eo, _ := charset.Lookup(c)
			return eo.NewDecoder().Reader(v), nil
		}
	}

	toreadable := "" // Human readble text (has decoded)
	stringlist := []string{}
	replacingc := []string{".", "[", "]"}

	if strings.IndexByte(argv0, ' ') > 0 {
		// The argument string include 1 or more space characters
		stringlist = strings.Split(argv0, " ")

	} else {
		// The argument string does not contain any space characters
		stringlist = append(stringlist, argv0)
	}

	for j, e := range stringlist {
		// Check and decode each part of the string
		if IsEncoded(e) == false {
			// Is not MIME-Encoded text part
			if j > 0 { toreadable += " " }; toreadable += e
			continue
		}

		// MIME-Encoded text part such as "=?UTF-8?B?44OL44Oj44O844Oz?="
		if strings.HasPrefix(e, "=?") == false {
			// For example, "[=?UTF-8?B?...]"
			for _, c := range replacingc { e = strings.Replace(e, c + "=?", "=?", -1) }
		}

		if strings.HasSuffix(e, "?=") == false {
			// For example, "=?UTF-8?B?....?=."
			for _, c := range replacingc { e = strings.Replace(e, "?=" + c, "?=", -1) }
		}

		if cv, nyaan := decodingif.DecodeHeader(e); nyaan == nil {
			// Successfully decoded
			if j > 0 { toreadable += " " }; toreadable += cv

		} else {
			// Failed to decode
			if j > 0 { toreadable += " " }; toreadable += e
		}
	}
	return toreadable, nil
}

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

