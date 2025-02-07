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
import sisimoji "sisimai/string"

func TestIndex(t *testing.T) {
	fn := "sisimai/reason.Index"
	ae := []string{
		"authfailure", "badreputation", "blocked", "contenterror", "exceedlimit", "expired", "failedstarttls",
		"feedback", "filtered", "hasmoved", "hostunknown", "mailboxfull", "mailererror", "mesgtoobig",
		"networkerror", "norelaying", "notaccept", "notcompliantrfc", "onhold", "policyviolation",
		"rejected", "requireptr", "securityerror", "spamdetected", "speeding", "suppressed", "suspend",
		"syntaxerror", "systemerror", "systemfull", "toomanyconn", "userunknown", "virusdetected",
	}
	cx := 0
	cv := Index()

	cx++; if len(cv) ==  0 { t.Errorf("%s() returns empty", fn) }
	cx++; if len(cv) != 33 { t.Errorf("%s() returns empty", fn) }
	for _, e := range cv {
		cx++; if e == "" { t.Errorf("%s() includes an empty string", fn) }
		cx++; if sisimoji.EqualsAny(strings.ToLower(e), ae) == false {
			t.Errorf("%s() returns invalid reason name: %s", fn, e)
		}
	}

	t.Logf("The number of tests = %d", cx)
}

func TestRetry(t *testing.T) {
	fn := "sisimai/reason.Retry"
	cx := 0
	cv := Retry()

	cx++; if len(cv) ==  0 { t.Errorf("%s() returns empty", fn) }
	for e := range cv {
		cx++; if cv[e] == false { t.Errorf("%s() returns a value which is false: %s", fn, e) }
	}

	t.Logf("The number of tests = %d", cx)
}

