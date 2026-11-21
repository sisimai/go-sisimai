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

// % go test -tags=fuzz -fuzz=FuzzIsIncluded -fuzztime=1m ./address
func FuzzIsIncluded(f *testing.F) {
	emailseeds := []string{
		"<neko@example.jp>",
		"<n@e.jp>",
		"Kijitora neko@example.jp (Nekochan)",
		"Sironeko <siro@example.jp> (Meow)",
		"<mailer-daemon>",
		"<neko@chan>",
		"<neko@localhost>",
		"<cat@localhost6>",
		"",
	}
	for _, e := range emailseeds { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, addr string) { _ = IsIncluded(addr) })
}

// % go test -tags=fuzz -fuzz=FuzzIsMailerDaemon -fuzztime=1m ./address
func FuzzIsMailerDaemon(f *testing.F) {
	emailseeds := []string{
		"http://postmaster.1and1.com/en/error-messages?ip=%1s",
		"http://postmaster.example.net/guidelines/",
		"http://postmaster.example.net/waters/fbl_change_form.html",
		"license assigned to them. The recipient's email admin can use the Office 365=",
		"recipient's email admin can fix.<ul><li><b>Check the error for information abo=",
		"likely a problem that only the recipient's email admin can fix.",
		"below. It's likely that the recipient's email admin is the only one who c=",
		"of your organization. Contact your email admin for assistance.",
		"postmaster@cx.libsisimai.org: Domain does not exist",
		"postmaster@mxr.biglobe.ne.jp",
		"421 4.7.0 [TSS04] Messages from 192.0.2.25 temporarily deferred due to unexpected volume or user complaints",
		"re an email admin, refer to the More Info for Email Admins section below.",
		"contact your email admin for assistance.<br>",
		"tell their email admin that it appears that their email system is refusing =",
		"this problem.</p><p>For Email Admins<br>No connection could be made because=",
		"for administrators:</b></p>",
		"ues, forward this message to your email admin.",
	}
	for _, e := range emailseeds { f.Add(e + "\n") }
	f.Fuzz(func(t *testing.T, addr string) { _ = IsMailerDaemon(addr) })
}

