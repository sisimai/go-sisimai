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
import "path/filepath"

// % go test -tags=fuzz -fuzz=FuzzPart -fuzztime=1m ./rfc5322
func FuzzPart(f *testing.F) {
	emailfiles := filepath.Join("..", "set-of-emails", "maildir", "bsd", "lhost-postfix-17.eml")
	boundaries := []string{"fffffff.0000000/mx20.example.jp"}
	payload, _ := os.ReadFile(emailfiles); if len(payload) > 0 {
		f.Add(string(payload))
		f.Fuzz(func(t *testing.T, mesg string) { _ = Part(&mesg, boundaries, false) })
	}
}
