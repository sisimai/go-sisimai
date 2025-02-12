// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      __   _       _____                 _ _    _       _     _                   
// |_   _|__  ___| |_   / /__(_)___  | ____|_ __ ___   __ _(_) |  / \   __| | __| |_ __ ___  ___ ___ 
//   | |/ _ \/ __| __| / / __| / __| |  _| | '_ ` _ \ / _` | | | / _ \ / _` |/ _` | '__/ _ \/ __/ __|
//   | |  __/\__ \ |_ / /\__ \ \__ \_| |___| | | | | | (_| | | |/ ___ \ (_| | (_| | | |  __/\__ \__ \
//   |_|\___||___/\__/_/ |___/_|___(_)_____|_| |_| |_|\__,_|_|_/_/   \_\__,_|\__,_|_|  \___||___/___/
import "testing"
import "strings"

func TestEmailAddress(t *testing.T) {
	cc := "EmailAddress"
	cx := 0
	cv := &EmailAddress{
		Address: "neko@example.jp",
		User:    "neko",
		Host:    "example.jp",
		Verp:    "neko+cat=example.org@example.jp",
		Alias:   "neko+cat@example.jp",
		Name:    "nekochan",
		Comment: "(meow)",
	}

	cx++; if cv == nil                        { t.Fatalf("%s{} = nil", cc) }
	cx++; if cv.Address == "" { t.Errorf("%s.Address is empty", cc) }
	cx++; if cv.User    == "" { t.Errorf("%s.User is empty", cc) }
	cx++; if cv.Host    == "" { t.Errorf("%s.Host is empty", cc) }
	cx++; if cv.Verp    == "" { t.Errorf("%s.Verp is empty", cc) }
	cx++; if cv.Alias   == "" { t.Errorf("%s.Alias is empty", cc) }
	cx++; if cv.Name    == "" { t.Errorf("%s.Name is empty", cc) }
	cx++; if cv.Comment == "" { t.Errorf("%s.Comment is empty", cc) }
	cx++; if cv.Void()  == true { t.Errorf("%s.Void() returns true", cc) }

	cx++; if strings.Contains(cv.Address, "@") == false { t.Errorf("%s.Address does not include @: %s", cc, cv.Address) }
	cx++; if strings.Contains(cv.User, "@")    == true  { t.Errorf("%s.User includes @: %s", cc, cv.User) }
	cx++; if strings.Contains(cv.Host, "@")    == true  { t.Errorf("%s.Host includes @: %s", cc, cv.Host) }
	cx++; if strings.Contains(cv.Verp, "@")    == false { t.Errorf("%s.Verp does not include @: %s", cc, cv.Verp) }
	cx++; if strings.Contains(cv.Alias, "@")   == false { t.Errorf("%s.Alias does not include @: %s", cc, cv.Alias) }

	t.Logf("The number of tests = %d", cx)
}

