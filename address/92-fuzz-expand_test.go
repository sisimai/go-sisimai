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

// % go test -tags=fuzz -fuzz=FuzzExpandVERP -fuzztime=1m ./address
func FuzzExpandVERP(f *testing.F) {
	emailseeds := []string{
		"cat+neko=example.jp@example.org",
		"+@example.org",
		"=@example.org",
		"+=@example.org",
		"=+@example.org",
		"+neko!example.jp",
		"=neko!example.jp",
		"cat+neko=meow@example@jp",
		"neko",
	}
	for _, e := range emailseeds { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, verp string) { _ = ExpandVERP(verp) })
}

// % go test -tags=fuzz -fuzz=FuzzExpandAlias -fuzztime=1m ./address
func FuzzExpandAlias(f *testing.F) {
	emailseeds := []string{
		"<neko>",
		"<neko@example.jp>",
		"+neko@example.jp",
		"neko+example+jp",
		`"neko+example+jp"@example.org`,
		"neko+cat@example.jp",
		"neko@example.jp",
	}
	for _, e := range emailseeds { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, addr string) { _ = ExpandAlias(addr) })
}
