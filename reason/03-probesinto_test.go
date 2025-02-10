// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package reason

//  _____         _      __                                ____            _               ___       _        
// |_   _|__  ___| |_   / / __ ___  __ _ ___  ___  _ __   |  _ \ _ __ ___ | |__   ___  ___|_ _|_ __ | |_ ___  
//   | |/ _ \/ __| __| / / '__/ _ \/ _` / __|/ _ \| '_ \  | |_) | '__/ _ \| '_ \ / _ \/ __|| || '_ \| __/ _ \ 
//   | |  __/\__ \ |_ / /| | |  __/ (_| \__ \ (_) | | | |_|  __/| | | (_) | |_) |  __/\__ \| || | | | || (_) |
//   |_|\___||___/\__/_/ |_|  \___|\__,_|___/\___/|_| |_(_)_|   |_|  \___/|_.__/ \___||___/___|_| |_|\__\___/ 
import "testing"

func TestProbesInto(t *testing.T) {
	fn := "sisimai/reason.ProbesInto"
	cx := 0

	for _, cr := range Index() {
		cx++; if ProbesInto[cr](nil) == true { t.Errorf("%s[%s](nil) returns true", fn, cr) }
	}
	t.Logf("The number of tests = %d", cx)
}

