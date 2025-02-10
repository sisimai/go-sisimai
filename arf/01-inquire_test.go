// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package arf

//  _____         _      ___    ____  _____ 
// |_   _|__  ___| |_   / / \  |  _ \|  ___|
//   | |/ _ \/ __| __| / / _ \ | |_) | |_   
//   | |  __/\__ \ |_ / / ___ \|  _ <|  _|  
//   |_|\___||___/\__/_/_/   \_\_| \_\_|    
import "testing"

func TestIsARF(t *testing.T) {
	fn := "sisimai/arf.isARF"
	cx := 0
	cx++; if isARF(nil) == true { t.Errorf("%s(nil) returns true", fn) }

	t.Logf("The number of tests = %d", cx)
}

func TestInquire(t *testing.T) {
	fn := "sisimai/arf.Inquire"
	cx := 0
	cv := Inquire(nil) 

	cx++; if cv.Void() == false { t.Errorf("%s(nil).Void() returns false", fn) }
	t.Logf("The number of tests = %d", cx)
}

