// Copyright (C) 2024 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package lhost

//       _               _      _______          _                            ____   ___   ___ _____ 
//  _ __| |__   ___  ___| |_   / / ____|_  _____| |__   __ _ _ __   __ _  ___|___ \ / _ \ / _ \___ / 
// | '__| '_ \ / _ \/ __| __| / /|  _| \ \/ / __| '_ \ / _` | '_ \ / _` |/ _ \ __) | | | | | | ||_ \ 
// | |  | | | | (_) \__ \ |_ / / | |___ >  < (__| | | | (_| | | | | (_| |  __// __/| |_| | |_| |__) |
// |_|  |_| |_|\___/|___/\__/_/  |_____/_/\_\___|_| |_|\__,_|_| |_|\__, |\___|_____|\___/ \___/____/ 
//                                                                 |___/                             
import "strings"
import "sisimai/sis"

func init() {
	// Decode bounce messages from Microsoft Exchange Server 2003: https://www.microsoft.com/microsoft-365/exchange/email
	InquireFor["Exchange2003"] = func(bf *sis.BeforeFact) sis.RisingUnderway {
		// @param    *sis.BeforeFact bf  Message body of a bounce email
		// @return   RisingUnderway      RisingUnderway structure
		if len(bf.Head) == 0 { return sis.RisingUnderway{} }
		if len(bf.Body) == 0 { return sis.RisingUnderway{} }

		// X-MS-TNEF-Correlator: <00000000000000000000000000000000000000@example.com>
		// X-Mailer: Internet Mail Service (5.5.1960.3)
		// X-MS-Embedded-Report:
		proceedsto := false
		if len(bf.Head["x-ms-embedded-report"]) > 0 { proceedsto = true }
		for proceedsto == false {
			// Check X-Mailer, X-MimeOLE, and Received headers
			tryto := []string{
				"Internet Mail Service (",                           // X-Mailer:
				"Microsoft Exchange Server Internet Mail Connector", // X-Mailer:
				"Produced By Microsoft Exchange",                    // X-MimeOLE:
				" with Internet Mail Service (",                     // Received:
			}
			if len(bf.Head["x-mailer"]) > 0 {
				// X-Mailer:  Microsoft Exchange Server Internet Mail Connector Version 4.0.994.63
				// X-Mailer: Internet Mail Service (5.5.2232.9)
				if strings.HasPrefix(bf.Head["x-mailer"][0], tryto[0]) { proceedsto = true; break }
				if strings.HasPrefix(bf.Head["x-mailer"][0], tryto[1]) { proceedsto = true; break }
			}

			if len(bf.Head["x-mimeole"]) > 0 {
				// X-MimeOLE: Produced By Microsoft Exchange V6.5
				if strings.HasPrefix(bf.Head["x-mimeole"][0],tryto[2]) { proceedsto = true; break }
			}

			for _, e := range bf.Head["received"] {
				// Received: by ***.**.** with Internet Mail Service (5.5.2657.72)
				if strings.Contains(e, tryto[3]) == false { continue }
				proceedsto = true; break
			}
			break
		}
		if proceedsto == false { return sis.RisingUnderway{} }


        return sis.RisingUnderway{}
    }
}

