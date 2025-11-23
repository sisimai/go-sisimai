// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc5322

//  _____         _      ______  _____ ____ ____ _________  ____  
// |_   _|__  ___| |_   / /  _ \|  ___/ ___| ___|___ /___ \|___ \ 
//   | |/ _ \/ __| __| / /| |_) | |_ | |   |___ \ |_ \ __) | __) |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___ ___) |__) / __/ / __/ 
//   |_|\___||___/\__/_/  |_| \_\_|   \____|____/____/_____|_____|
import "testing"
import "os"
import "path/filepath"

func TestLooksLikeEmail(t *testing.T) {
	fn := "rfc5322.LooksLikeEmail"
	cx := 0
	ae := []string{
		"lhost-imailserver-04.eml",
		"lhost-dragonfly-26.eml",
		"lhost-exchange2007-05.eml",
	}
	xe := []string{"", "\x00nekocat", "neko\x7fcat", "neko\xffcat"}

	for _, e := range ae {
		fe    := filepath.Join("..", "set-of-emails", "maildir", "bsd", e)
		by, _ := os.ReadFile(fe); sy := string(by)

		cx++; if len(by) == 0                 { t.Errorf("rfc5322.%s(%s) is empty", fn, e)      }
		cx++; if LooksLikeEmail(&sy) == false { t.Errorf("rfc5322.%s(%s) returns false", fn, e) }
	}
	
	for _, e := range xe {
		cx++; if LooksLikeEmail(&e) == true   { t.Errorf("rfc5322.%s(%s) returns true ", fn, e) }
	}
	for _, e := range []string{filepath.Join("..", ".gitignore"), filepath.Join("..", "set-of-emails", "mailbox", "size-2")} {
		by, _ := os.ReadFile(e); sy := string(by)
		cx++; if LooksLikeEmail(&sy) == true  { t.Errorf("rfc5322.%s(%s) returns true ", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

