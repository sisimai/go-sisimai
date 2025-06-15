// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package moji

//  _____         _      __               _ _ 
// |_   _|__  ___| |_   / / __ ___   ___ (_|_)
//   | |/ _ \/ __| __| / / '_ ` _ \ / _ \| | |
//   | |  __/\__ \ |_ / /| | | | | | (_) | | |
//   |_|\___||___/\__/_/ |_| |_| |_|\___// |_|
//                                     |__/   
import "testing"

func TestAligned(t *testing.T) {
	fn := "moji.Aligned"
	cw := "Final-Recipient: rfc822; <neko@example.jp>"
	cx := 0

	cx++; if Aligned(cw, []string{"rfc822", "<", "@", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", " ", "@", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", ";", "<", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"Final-", ":", ";", ">"}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", "[", "@", ">"}) == true  { t.Errorf("%s(%s) returns true",  fn, cw) }
	cx++; if Aligned(cw, []string{"rfc822", "<", "@", " "}) == true  { t.Errorf("%s(%s) returns true",  fn, cw) }

	cx++; if Aligned("", []string{})          == true { t.Errorf("%s(%s) returns true",  fn, "") }
	cx++; if Aligned("neko", []string{})      == true { t.Errorf("%s(%s) returns true",  fn, "neko") }
	cx++; if Aligned("", []string{"neko"})    == true { t.Errorf("%s(%s) returns true",  fn, "") }
	cx++; if Aligned("cat", []string{"neko"}) == true { t.Errorf("%s(%s) returns true",  fn, "cat") }

	t.Logf("The number of tests = %d", cx)
}

func TestAlignedAny(t *testing.T) {
	fn := "moji.AlignedAny"
	cw := "Final-Recipient: rfc822; <neko@example.jp>"
	cx := 0
	s0 := []string{"NEKO", "CAT"}

	cx++; if AlignedAny(cw, [][]string{s0, []string{"rfc822", "<", "@", ">"}}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if AlignedAny(cw, [][]string{s0, []string{"rfc822", " ", "@", ">"}}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if AlignedAny(cw, [][]string{s0, []string{"rfc822", ";", "<", ">"}}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if AlignedAny(cw, [][]string{s0, []string{"Final-", ":", ";", ">"}}) == false { t.Errorf("%s(%s) returns false", fn, cw) }
	cx++; if AlignedAny(cw, [][]string{s0, []string{"rfc822", "[", "@", ">"}}) == true  { t.Errorf("%s(%s) returns true",  fn, cw) }
	cx++; if AlignedAny(cw, [][]string{s0, []string{"rfc822", "<", "@", " "}}) == true  { t.Errorf("%s(%s) returns true",  fn, cw) }

	cx++; if AlignedAny("", [][]string{})          == true   { t.Errorf("%s(%s) returns true",  fn, "") }
	cx++; if AlignedAny("neko", [][]string{})      == true   { t.Errorf("%s(%s) returns true",  fn, "neko") }
	cx++; if AlignedAny("", [][]string{[]string{"neko"}})    { t.Errorf("%s(%s) returns true",  fn, "") }
	cx++; if AlignedAny("cat", [][]string{[]string{"neko"}}) { t.Errorf("%s(%s) returns true",  fn, "cat") }

	t.Logf("The number of tests = %d", cx)

}
