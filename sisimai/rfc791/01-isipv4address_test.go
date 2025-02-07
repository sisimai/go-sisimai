// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rfc791

//  _____         _      ______  _____ ____ _____ ___  _ 
// |_   _|__  ___| |_   / /  _ \|  ___/ ___|___  / _ \/ |
//   | |/ _ \/ __| __| / /| |_) | |_ | |      / / (_) | |
//   | |  __/\__ \ |_ / / |  _ <|  _|| |___  / / \__, | |
//   |_|\___||___/\__/_/  |_| \_\_|   \____|/_/    /_/|_|
import "testing"

func TestIsIPv4Address(t *testing.T) {
	fn := "sisimai/rfc791.IsIPv4Address"
	cx := 0
	ae := []struct {text string; expected bool}{
		{"0.0.0.0", true},
		{"1.2.3.4", true},
		{"127.0.0.1", true},
		{"192.0.2.25", true},
		{"255.255.255.255", true},
		{"13597846544569", false},
		{"3.14", false},
		{"3.14", false},
		{"8.17.1", false},
		{"5.6.7.8.9", false},
		{"123.456.78.9", false},
		{"", false},
	}

	for _, e := range ae {
		cx++; if cv := IsIPv4Address(e.text); cv != e.expected {
			t.Errorf("%s(%s) returns %t", fn, e.text, e.expected)
		}
	}
	cx++; if IsIPv4Address("") == true { t.Errorf("%s(nil) returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}

