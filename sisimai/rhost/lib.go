// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _   
//  _ __| |__   ___  ___| |_ 
// | '__| '_ \ / _ \/ __| __|
// | |  | | | | (_) \__ \ |_ 
// |_|  |_| |_|\___/|___/\__|
import "strings"
import "sisimai/sis"

var ReturnedBy = map[string]func(*sis.Fact) bool {}
var RhostClass = map[string][]string{
	"Apple":     []string{".mail.icloud.com", ".apple.com", ".me.com"},
	"Cox":       []string{"cox.net"},
	"FrancePTT": []string{".laposte.net", ".orange.fr", ".wanadoo.fr"},
	"GoDaddy":   []string{"smtp.secureserver.net", "mailstore1.secureserver.net"},
	"Google":    []string{"aspmx.l.google.com", "gmail-smtp-in.l.google.com"},
	"IUA":       []string{".email.ua"},
	"KDDI":      []string{".ezweb.ne.jp", "msmx.au.com"},
	"Microsoft": []string{".prod.outlook.com", ".protection.outlook.com"},
	"Mimecast":  []string{".mimecast.com"},
	"NTTDOCOMO": []string{"mfsmax.docomo.ne.jp"},
	"Spectrum":  []string{"charter.net"},
	"Tencent":   []string{".qq.com"},
	"YahooInc":  []string{".yahoodns.net"},
}

// Find() detects the bounce reason from certain remote hosts
func Find(fo *sis.Fact) string {
	// @param    *sis.Fact fo    Struct to be detected the reason
	// @return   string          Bounce reason name or an empty string
	if fo.DiagnosticCode == "" { return "" }

	remotehost := strings.ToLower(fo.Rhost)
	domainpart := strings.ToLower(fo.Destination)
	rhostclass := ""
	if len(remotehost + domainpart) == 0 { return "" }

	for e := range RhostClass {
		// Try to match the remote host or the domain part with each value of RhostClass
		for _, r := range RhostClass[e] {
			// - Whether the remote host (fo.Rhost) includes "r" or not
			// - Whether "r" includes the domain part of the recipient address or not
			if strings.HasSuffix(remotehost, r) { rhostclass = e; break }
			if strings.HasSuffix(r, domainpart) { rhostclass = e; break }
		}
		if rhostclass != "" { break }
	}
	if rhostclass == "" { return "" }

	return ReturnedBy[rhostclass](sis.Fact)
}

