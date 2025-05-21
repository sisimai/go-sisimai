// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____         _      __                              
// |_   _|__  ___| |_   / / __ ___  __ _ ___  ___  _ __  
//   | |/ _ \/ __| __| / / '__/ _ \/ _` / __|/ _ \| '_ \ 
//   | |  __/\__ \ |_ / /| | |  __/ (_| \__ \ (_) | | | |
//   |_|\___||___/\__/_/ |_|  \___|\__,_|___/\___/|_| |_|
import "testing"
import "strings"
import "slices"

var ae = []string{
	"authfailure", "badreputation", "blocked", "contenterror", "exceedlimit", "expired", "failedstarttls",
	"feedback", "filtered", "hasmoved", "hostunknown", "mailboxfull", "mailererror", "mesgtoobig",
	"networkerror", "norelaying", "notaccept", "notcompliantrfc", "onhold", "policyviolation",
	"rejected", "requireptr", "securityerror", "spamdetected", "speeding", "suppressed", "suspend",
	"syntaxerror", "systemerror", "systemfull", "toomanyconn", "userunknown", "virusdetected",
	"undefined", "delivered", "vacation",
}

func TestAvailables(t *testing.T) {
	fn := "reason.Availables"
	cx := 0
	cv := Availables

	cx++; if len(cv) ==  0 { t.Errorf("%s is empty", fn) }
	cx++; if len(cv) != 36 { t.Errorf("%s includes invalid elements: %d", fn, len(cv)) }
	for e := range cv {
		cx++; if e == ""     { t.Errorf("%s returned an empty key", fn) }
		cx++; if cv[e] == "" { t.Errorf("%s[%s] is empty", fn, cv[e]) }
		cx++; if slices.Contains(ae, strings.ToLower(e)) == false {
			t.Errorf("%s() returns invalid reason name: %s", fn, e)
		}
		cx++; if ProbesInto[e](nil) == true { t.Errorf("ProbesInto[%s](nil) returns true", e) }
		cx++; if IncludedIn[e](" ") == true { t.Errorf("IncludedIn[%s](' ') returns true", e) }
	}

	t.Logf("The number of tests = %d", cx)
}

func TestIsExplicit(t *testing.T) {
	fn := "reason.IsExplicit"
	cx := 0

	for _, e := range ae {
		if e == "onhold" || e == "undefined" { continue }
		cx++; if cv := IsExplicit(e); cv == false { t.Errorf("%s(%s) returns false", fn, e) }
	}
	cx++; if IsExplicit("")          == true { t.Errorf("%s() returns true", fn) }
	cx++; if IsExplicit("onhold")    == true { t.Errorf("%s(onhold) returns true", fn) }
	cx++; if IsExplicit("undefined") == true { t.Errorf("%s(undefined) returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}
