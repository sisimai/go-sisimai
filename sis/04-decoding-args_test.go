// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package sis

//  _____         _      __   _       ____                     _ _                _                  
// |_   _|__  ___| |_   / /__(_)___  |  _ \  ___  ___ ___   __| (_)_ __   __ _   / \   _ __ __ _ ___ 
//   | |/ _ \/ __| __| / / __| / __| | | | |/ _ \/ __/ _ \ / _` | | '_ \ / _` | / _ \ | '__/ _` / __|
//   | |  __/\__ \ |_ / /\__ \ \__ \_| |_| |  __/ (_| (_) | (_| | | | | | (_| |/ ___ \| | | (_| \__ \
//   |_|\___||___/\__/_/ |___/_|___(_)____/ \___|\___\___/ \__,_|_|_| |_|\__, /_/   \_\_|  \__, |___/
//                                                                       |___/             |___/     
import "testing"

// Delivered bool // Include sis.Fact{}.Action = "delivered" records in the decoded data
// Vacation  bool // Include sis.Fact{}.Reason = "vacation" records in the decoded data
// Callback0 CfParameter0 // [0] The 1st callback function
// Callback1 CfParameter1 // [1] The 2nd callback function
func TestDecodingArgs(t *testing.T) {
	fn := "sis.DecodingArgs"
	c1 := func(arg *CallbackArgs) (map[string]interface{}, error) {
		data := make(map[string]interface{}); data["nekochan"] = []string{"kijitora", "nyaaaan"}
		return data, nil
	}
	c2 := func(arg *CallbackArgs) (bool, error) { return true, nil }
	cv := &DecodingArgs{
		Delivered: false,
		Vacation:  false,
		Callback0: c1,
		Callback1: c2,
	}
	cx := 0

	cx++; if cv           == nil   { t.Fatalf("%s{} = nil", fn) }
	cx++; if cv.Delivered != false { t.Errorf("%s.Delivered is not false", fn) }
	cx++; if cv.Vacation  != false { t.Errorf("%s.Vacation is not nil", fn) }
	cx++; if cv.Callback0 == nil   { t.Errorf("%s.Callback0 is nil", fn) }
	cx++; if cv.Callback1 == nil   { t.Errorf("%s.Callback1 is nil", fn) }

	t.Logf("The number of tests = %d", cx)
}

