// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//       _               _      _______          _                            ____   ___   ___ _____ 
//  _ __| |__   ___  ___| |_   / / ____|_  _____| |__   __ _ _ __   __ _  ___|___ \ / _ \ / _ \___  |
// | '__| '_ \ / _ \/ __| __| / /|  _| \ \/ / __| '_ \ / _` | '_ \ / _` |/ _ \ __) | | | | | | | / / 
// | |  | | | | (_) \__ \ |_ / / | |___ >  < (__| | | | (_| | | | | (_| |  __// __/| |_| | |_| |/ /  
// |_|  |_| |_|\___/|___/\__/_/  |_____/_/\_\___|_| |_|\__,_|_| |_|\__, |\___|_____|\___/ \___//_/   
//                                                                 |___/                             
import "strings"
import "sisimai/sis"

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
			// Content-Language: en-US, fr-FR
			if strings.HasPrefix(bf.Head["subject"][0], "Undeliverable")    { proceedsto = 1; break }
			if strings.HasPrefix(bf.Head["subject"][0], "Non_remis_")       { proceedsto = 1; break }
			if strings.HasPrefix(bf.Head["subject"][0], "Non recapitabile") { proceedsto = 1; break }
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

        return sis.RisingUnderway{}
    }
}

