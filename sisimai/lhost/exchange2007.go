// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//  _ _               _      _______          _                            ____   ___   ___ _____ 
// | | |__   ___  ___| |_   / / ____|_  _____| |__   __ _ _ __   __ _  ___|___ \ / _ \ / _ \___  |
// | | '_ \ / _ \/ __| __| / /|  _| \ \/ / __| '_ \ / _` | '_ \ / _` |/ _ \ __) | | | | | | | / / 
// | | | | | (_) \__ \ |_ / / | |___ >  < (__| | | | (_| | | | | (_| |  __// __/| |_| | |_| |/ /  
// |_|_| |_|\___/|___/\__/_/  |_____/_/\_\___|_| |_|\__,_|_| |_|\__, |\___|_____|\___/ \___//_/   
//                                                              |___/                             
import "strings"
import "sisimai/sis"
import "sisimai/rfc5322"
import "sisimai/smtp/reply"
import "sisimai/smtp/status"
import sisiaddr "sisimai/address"
import sisimoji "sisimai/string"
import "fmt"

func init() {
	// Decode bounce messages from Microsoft Exchange Server 2007: https://www.microsoft.com/microsoft-365/exchange/email
	InquireFor["Exchange2007"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		// These headers exist only a bounce mail from Office365
		if len(bf.Head["x-ms-exchange-crosstenant-originalarrivaltime"]) > 0 { return sis.RisingUnderway{} }
		if len(bf.Head["x-ms-exchange-crosstenant-fromentityheader"])    > 0 { return sis.RisingUnderway{} }

		proceedsto := uint8(0); for {
			// Content-Language: en-US, fr-FR, sv-SE
			if strings.HasPrefix(bf.Head["subject"][0], "Undeliverable")    { proceedsto = 1; break }
			if strings.HasPrefix(bf.Head["subject"][0], "Non_remis_")       { proceedsto = 1; break }
			if strings.HasPrefix(bf.Head["subject"][0], "Non recapitabile") { proceedsto = 1; break }
			if strings.HasPrefix(bf.Head["subject"][0], "Olevererbart")     { proceedsto = 1; break }
			break
		}
		if proceedsto                       == 0 { return sis.RisingUnderway{} }
		if len(bf.Head["content-language"]) == 0 { return sis.RisingUnderway{} }

		for {
			// Content-Laugage: JP or ja-JP
			if len(bf.Head["content-language"][0]) == 2 { proceedsto++; break }
			if len(bf.Head["content-language"][0]) == 5 { proceedsto++; break }
			break
		}
		if proceedsto < 2 { return sis.RisingUnderway{} }
fmt.Printf("BODY = (%s)\n", bf.Body)
		indicators := INDICATORS()
		boundaries := []string{
			"Original message headers:",             // en-US
			"tes de message d'origine :",            // fr-FR/En-têtes de message d'origine
			"Intestazioni originali del messaggio:", // it-CH
			"Ursprungshuvuden:",                     // sv-SE
		}
		startingof := map[string][]string{
			"message": []string{
				"Diagnostic information for administrators:",           // en-US
				"Informations de diagnostic pour les administrateurs",  // fr-FR
				"Informazioni di diagnostica per gli amministratori",   // it-CH
				"Diagnostisk information f",                            // sv-SE
			},
			"error": []string{" RESOLVER.", " QUEUE."},
			"rhost": []string{
				"Generating server",        // en-US
				"Serveur de g",             // fr-FR/Serveur de génération
				"Server di generazione",    // it-CH
				"Genererande server",       // sv-SE
			},
		}
		ndrsubject := map[string]string{
			"SMTPSEND.DNS.NonExistentDomain": "hostunknown",   // 554 5.4.4 SMTPSEND.DNS.NonExistentDomain
			"SMTPSEND.DNS.MxLoopback":        "networkerror",  // 554 5.4.4 SMTPSEND.DNS.MxLoopback
			"RESOLVER.ADR.BadPrimary":        "systemerror",   // 550 5.2.0 RESOLVER.ADR.BadPrimary
			"RESOLVER.ADR.RecipNotFound":     "userunknown",   // 550 5.1.1 RESOLVER.ADR.RecipNotFound
			"RESOLVER.ADR.ExRecipNotFound":   "userunknown",   // 550 5.1.1 RESOLVER.ADR.ExRecipNotFound
			"RESOLVER.ADR.RecipLimit":        "toomanyconn",   // 550 5.5.3 RESOLVER.ADR.RecipLimit
			"RESOLVER.ADR.InvalidInSmtp":     "systemerror",   // 550 5.1.0 RESOLVER.ADR.InvalidInSmtp
			"RESOLVER.ADR.Ambiguous":         "systemerror",   // 550 5.1.4 RESOLVER.ADR.Ambiguous, 420 4.2.0 RESOLVER.ADR.Ambiguous
			"RESOLVER.RST.AuthRequired":      "securityerror", // 550 5.7.1 RESOLVER.RST.AuthRequired
			"RESOLVER.RST.NotAuthorized":     "rejected",      // 550 5.7.1 RESOLVER.RST.NotAuthorized
			"RESOLVER.RST.RecipSizeLimit":    "mesgtoobig",    // 550 5.2.3 RESOLVER.RST.RecipSizeLimit
			"QUEUE.Expired":                  "expired",       // 550 4.4.7 QUEUE.Expired
		}

		dscontents := []sis.DeliveryMatter{{}}
		emailparts := rfc5322.Part(&bf.Body, boundaries, false)
		readcursor := uint8(0)              // Points the current cursor position
		recipients := uint8(0)              // The number of 'Final-Recipient' header
		connvalues := 0                     // Counter, 3 if it has got the all values of connheader
		connheader := [1]string{""}         // [rhost]
		v          := &(dscontents[len(dscontents) - 1])

		for _, e := range(strings.Split(emailparts[0], "\n")) {
			// Read error messages and delivery status lines from the head of the email to the
			// previous line of the beginning of the original message.
			if readcursor == 0 {
				// Beginning of the bounce message or message/delivery-status part
				for _, f := range startingof["message"] {
					if strings.HasPrefix(e, f) { readcursor |= indicators["deliverystatus"]; break }
				}
				continue
			}
			if readcursor & indicators["deliverystatus"] == 0 { continue }

			if connvalues == len(connheader) {
				// Diagnostic information for administrators:
				//
				// Generating server: mta2.neko.example.jp
				//
				// kijitora@example.jp
				// //550 5.1.1 RESOLVER.ADR.RecipNotFound; not found ////
				//
				// Original message headers:
				if strings.Contains(e, " ") == false && strings.Index(e, "@") > 1 {
					// This line includes an email address only
					if len(v.Recipient) > 0 {
						// There are multiple recipient addresses in the message body.
						dscontents = append(dscontents, sis.DeliveryMatter{})
						v = &(dscontents[len(dscontents) - 1])
					}
					v.Recipient = sisiaddr.S3S4(e)
					recipients += 1

				} else {
					cr := reply.Find(e, "")
					cs := status.Find(e, "")
					if cr != "" || cs != "" {
						// #550 5.1.1 RESOLVER.ADR.RecipNotFound; not found ##
						// #550 5.2.3 RESOLVER.RST.RecipSizeLimit; message too large for this recipient ##
						// Remote Server returned '550 5.1.1 RESOLVER.ADR.RecipNotFound; not found'
						// 3/09/2016 8:05:56 PM - Remote Server at mydomain.com (10.1.1.3) returned '550 4.4.7 QUEUE.Expired; message expired'
						v.ReplyCode = cr
						v.Status    = cs
						v.Diagnosis = e

					} else {
						// This line is a continued line of the error message
						if v.Diagnosis == ""                            { continue }
						if strings.HasSuffix(v.Diagnosis, "=") == false { continue }
						v.Diagnosis = strings.TrimRight(v.Diagnosis, "=")
					}
				}
			} else {
				// Diagnostic information for administrators:
				//
				// Generating server: mta22.neko.example.org
				if sisimoji.HasPrefixAny(e, startingof["rhost"]) == false { continue }
				if connheader[0] != ""                                    { continue }
				connheader[0] = strings.Trim(e[strings.Index(e, ":"):], " ")
				connvalues++
			}
		}
		if recipients == 0  { return sis.RisingUnderway{} }

		for j, _ := range dscontents {
			// Tidy up the error message
			e := &(dscontents[j])
			e.Diagnosis = sisimoji.Sweep(e.Diagnosis)

			p0 := -1
			p1 := strings.Index(e.Diagnosis, ";")
			for _, r := range startingof["error"] {
				// Find an error message and an error code 
				p0 = strings.Index(e.Diagnosis, r)
				if p0 > -1 { break }
			}
			if p0 < 0 || p1 < 0 { continue }

			// #550 5.1.1 RESOLVER.ADR.RecipNotFound; not found ##
			if len(ndrsubject[e.Diagnosis[p0:p1]]) > 0 { e.Reason = e.Diagnosis[p0:p1] }
		}

		return sis.RisingUnderway{ Digest: dscontents, RFC822: emailparts[1] }
    }
}

