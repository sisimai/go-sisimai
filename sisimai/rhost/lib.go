// Copyright (C) 2024-2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _   
//  _ __| |__   ___  ___| |_ 
// | '__| '_ \ / _ \/ __| __|
// | |  | | | | (_) \__ \ |_ 
// |_|  |_| |_|\___/|___/\__|
import "strings"
import "sisimai/sis"

var ReturnedBy = map[string]func(*sis.Fact) string {}
var RhostClass = map[string][]string{
	"Aol":         []string{".mail.aol.com", ".mx.aol.com"},
	"Apple":       []string{".mail.icloud.com", ".apple.com", ".me.com"},
	"Cox":         []string{"cox.net"},
	"Facebook":    []string{".facebook.com"},
	"FrancePTT":   []string{".laposte.net", ".orange.fr", ".wanadoo.fr"},
	"GoDaddy":     []string{"smtp.secureserver.net", "mailstore1.secureserver.net"},
	"Google":      []string{"aspmx.l.google.com", "gmail-smtp-in.l.google.com"},
	"GSuite":      []string{"googlemail.com"},
	"IUA":         []string{".email.ua"},
	"KDDI":        []string{".ezweb.ne.jp", "msmx.au.com"},
	"MessageLabs": []string{".messagelabs.com"},
	"Microsoft":   []string{".prod.outlook.com", ".protection.outlook.com", ".onmicrosoft.com", ".exchangelabs.com"},
	"Mimecast":    []string{".mimecast.com"},
	"NTTDOCOMO":   []string{"mfsmax.docomo.ne.jp"},
	"Outlook":     []string{".hotmail.com"},
	"Spectrum":    []string{"charter.net"},
	"Tencent":     []string{".qq.com"},
	"YahooInc":    []string{".yahoodns.net"},
}

// Name() returns the rhost class name
func Name(fo *sis.Fact) string {
	// @param    *sis.Fact fo    Decoded data
	// @return   string          rhost class name or an empty string
	if fo == nil { return "" }

	// Try to match the hostname patterns with the following order:
	// 1. destination: The domain part of the recipient address
	// 2. rhost: remote hostname
	// 3. lhost: local MTA hostname
	clienthost := strings.ToLower(fo.Lhost)
	remotehost := strings.ToLower(fo.Rhost)
	domainpart := strings.ToLower(fo.Destination)
	for e := range RhostClass {
		// Try to match the domain part of the recipient address with each value of RhostClass
		for _, r := range RhostClass[e] {
			// - Whether "r" includes the domain part of the recipient address or not
			if strings.HasSuffix(r, domainpart) { return e }
		}
	}

	for e := range RhostClass {
		// Try to match the remote host with each value of RhostClass
		for _, r := range RhostClass[e] {
			// - Whether the remote host (fo.Rhost) includes "r" or not
			if strings.HasSuffix(remotehost, r) { return e }
		}
	}

	// Neither the remote host nor the destination did not matched with any value of RhostClass
	for e := range RhostClass {
		// Try to match the client host with each value of RhostClass
		for _, r := range RhostClass[e] {
			// - Whether the local MTA host (fo.Lhost) includes "r" or not
			if strings.HasSuffix(clienthost, r) { return e }
		}
	}
	return ""
}

// Find() detects the bounce reason from certain remote hosts
func Find(fo *sis.Fact) string {
	// @param    *sis.Fact fo    Decoded data
	// @return   string          Bounce reason name or an empty string
	rhostclass := Name(fo); if rhostclass == "" { return "" }
	return ReturnedBy[rhostclass](fo)
}

