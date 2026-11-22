//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
package rfc5322
import "testing"
import "os"
import "strings"
import "net/mail"
import "path/filepath"

// % go test -tags=fuzz -fuzz=FuzzHeaders -fuzztime=1m ./rfc5322
func FuzzHeaders(f *testing.F) {
	emailfiles := filepath.Join("..", "set-of-emails", "maildir", "bsd", "lhost-opensmtpd-17.eml")
	payload, _ := os.ReadFile(emailfiles); if len(payload) > 0 {
		message := string(payload)
		headers := message

		if p1 := strings.Index(message, "\n\n"); p1 != -1 {
			headers = message[:p1]

		} else if p1 := strings.Index(message, "\r\n\r\n"); p1 != -1 {
			headers = message[:p1]
		}
		f.Add(headers)
		f.Fuzz(
			func(t *testing.T, heads string) {
				mesg, e := mail.ReadMessage(strings.NewReader(heads + "\n\nNyaaan?"))
				if e != nil { return }
				_ = Headers(&mesg.Header)
			},
		)
	}
}

