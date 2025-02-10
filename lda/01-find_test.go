// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lda

//  _____         _      ___     ____    _    
// |_   _|__  ___| |_   / / |   |  _ \  / \   
//   | |/ _ \/ __| __| / /| |   | | | |/ _ \  
//   | |  __/\__ \ |_ / / | |___| |_| / ___ \ 
//   |_|\___||___/\__/_/  |_____|____/_/   \_\
import "testing"

func TestFind(t *testing.T) {
	fn := "sisimai/lda.Find"
	cx := 0

	cx++; if Find(nil) != "" { t.Errorf("%s(nil) returns true", fn) }
	t.Logf("The number of tests = %d", cx)
}

