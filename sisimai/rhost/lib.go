// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package rhost

//       _               _   
//  _ __| |__   ___  ___| |_ 
// | '__| '_ \ / _ \/ __| __|
// | |  | | | | (_) \__ \ |_ 
// |_|  |_| |_|\___/|___/\__|

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

