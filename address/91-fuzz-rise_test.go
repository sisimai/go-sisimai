//go:build fuzz

// Copyright (C) 2026 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
//  _____         _      ___       _     _                   
// |_   _|__  ___| |_   / / \   __| | __| |_ __ ___  ___ ___ 
//   | |/ _ \/ __| __| / / _ \ / _` |/ _` | '__/ _ \/ __/ __|
//   | |  __/\__ \ |_ / / ___ \ (_| | (_| | | |  __/\__ \__ \
//   |_|\___||___/\__/_/_/   \_\__,_|\__,_|_|  \___||___/___/
//                                                           
package address
import "testing"

// % go test -tags=fuzz -fuzz=FuzzRise -fuzztime=1m ./address
func FuzzRise(f *testing.F) {
	emailseeds := [][3]string{
		[3]string{"neko@example.jp", "Nyaan?", "(kijitora)"},
		[3]string{"<neko@example.jp>", "<neko>", "<kijitora>"},
	}
	for _, e := range emailseeds { f.Add(e[0], e[1], e[2]) }
	f.Fuzz(func(t *testing.T, addr, name, comm string) {
		args := [3]string{addr, name, comm}
		_     = Rise(args) 
	})
}

